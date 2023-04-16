package wsx

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"online-chat/common/errorx"
	"online-chat/global"
	"time"
)

const PublishKey = "websocket-chat"

type WSMessage struct {
	gorm.Model
	FormID   string // 发送着
	TargetID string // 接受者
	Type     string // 消息类型： 群聊，私聊，广播
	Media    int    // 消息类型： 文字，图片，音频
	Message  string
	Pri      string // 图片相关
	Url      string // url相关
	Desc     string // 描述
	Amount   int    // 其他数字统计
}

func (*WSMessage) TableName() string {
	return "message"
}

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
