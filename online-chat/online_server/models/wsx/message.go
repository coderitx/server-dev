package wsx

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"net"
	"net/http"
	"online-chat/common/errorx"
	"online-chat/global"
	"sync"
	"time"
)

const PublishKey = "websocket-chat"

type WSMessage struct {
	gorm.Model
	FormID   int64 // 发送着
	TargetID int64 // 接受者
	Type     int   // 发送类型： 群聊，私聊，广播
	Media    int   // 消息类型： 文字，图片，音频
	Message  string
	Pri      string // 图片相关
	Url      string // url相关
	Desc     string // 描述
	Amount   int    // 其他数字统计
}

func (*WSMessage) TableName() string {
	return "message"
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

var (
	clientMap = make(map[int64]*Node, 0)
	lock      sync.RWMutex
)

// Publish 发布消息到redis
func Publish(ctx context.Context, channel string, msg WSMessage) error {
	err := global.RDB.Publish(ctx, channel, msg).Err()
	if err != nil {
		zap.S().Errorf("publish to redis error: %v", err)
		return err
	}
	return nil
}

// Subscribe 订阅redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := global.RDB.Subscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		zap.S().Errorf("subscribe redis msg error: %v", err)
		//return "", err
	} else {
		zap.S().Info("msg: ", msg.Payload)
	}
	fmt.Println("msg.Payload: ", msg.Payload)
	return msg.Payload, nil
}

func (w *WSMessage) SendMsg(ctx *gin.Context, upgrade websocket.Upgrader) int {
	ws, err := upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return errorx.ServerErrorCode
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	for {
		msg, err := Subscribe(ctx, PublishKey)
		if err != nil {
			fmt.Println("error: ", err)
			zap.S().Errorf("send msg error: %v", err)
		}

		t := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s][%s]", t, msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			zap.S().Errorf("write msg error: %v", err)
			return errorx.ServerErrorCode
		}
	}
	return errorx.SuccessCode
}

func (w *WSMessage) Chat(write http.ResponseWriter, request *http.Request) int {
	// 1. 合法性校验
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(write, request, nil)
	if err != nil {
		zap.S().Errorf("upgrader websocket error: %v", err)
		return errorx.ServerErrorCode
	}
	// 2. 获取连接
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	// 3. 用户关系

	// 4. userid 与 node 绑定
	lock.Lock()
	clientMap[w.FormID] = node
	lock.Unlock()
	// 5. 发送
	go sendProc(node)
	// 6. 接收
	go recvProc(node)

	// 进入聊天室初始化一个消息
	sendMsg(w.FormID, w.TargetID, []byte("hello server"))

	return errorx.SuccessCode
}

func sendProc(n *Node) {
	for {
		select {
		case data := <-n.DataQueue:
			err := n.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				zap.S().Errorf("write message to node error: %v", err)
				return
			}
		}
	}
}

func recvProc(n *Node) {
	for {
		_, data, err := n.Conn.ReadMessage()
		if err != nil {
			zap.S().Errorf("read message to node error: %v", err)
			return
		}
		broadMessage(data)
		fmt.Println("[ws] recv message: ", string(data))
	}
}

var udpsendChan chan []byte = make(chan []byte, 1024)

func broadMessage(data []byte) {
	udpsendChan <- data
}

func init() {
	go udpSendProc()
	go udpRecvProc()
}

// udp发送消息协程
func udpSendProc() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 10, 255),
		Port: 3000,
	})
	if err != nil {
		zap.S().Errorf("udp send message error: %v", err)
		return
	}
	defer conn.Close()

	for {
		select {
		case data := <-udpsendChan:
			_, err := conn.Write(data)
			if err != nil {
				zap.S().Errorf("write message to node error: %v", err)
				return
			}
			broadMessage(data)
			fmt.Println("[ws] udp message: ", string(data))
			return
		}
	}

}

// udp接收消息协程
func udpRecvProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		zap.S().Errorf("udp recv message error: %v", err)
		return
	}
	defer con.Close()
	for {
		var buf [512]byte
		n, err := con.Read(buf[:])
		if err != nil {
			zap.S().Errorf("read message to node error: %v", err)
			return
		}
		dispatch(buf[:n])
	}
}

// 后段调度逻辑
func dispatch(b []byte) {
	msg := WSMessage{}
	json.Unmarshal(b, &msg)
	// 根据msg.Type 调用不同的发送方法
	switch msg.Type {
	case 1:
		sendMsg(msg.FormID, msg.TargetID, b)
	case 2:
		//sendGroupMsg()
	case 3:
		//sendAllMsg()

	}
}

func sendMsg(fromID int64, targetID int64, msg []byte) {
	lock.Lock()
	node, ok := clientMap[targetID]
	lock.Unlock()
	if ok {
		node.DataQueue <- msg
	}
}
