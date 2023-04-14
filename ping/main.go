package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var (
	timeout int64
	size    int
	count   int
	typ     uint8 = 8
	code    uint8 = 0
)

type ICMP struct {
	Type        uint8
	Code        uint8
	CheckSum    uint16
	ID          uint16
	SequenceNum uint16
}

func getCommandArgs() {
	flag.Int64Var(&timeout, "w", 10000, "请求超时时间：毫秒")
	flag.IntVar(&size, "l", 32, "请求缓冲区长度")
	flag.IntVar(&count, "n", 4, "回穿次数")
	flag.Parse()
}

func checkSum(data []byte) uint16 {
	length := len(data)
	index := 0
	var sum uint32 = 0
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		length -= 2
		index += 2
	}
	if length != 0 {
		sum += uint32(data[index])
	}
	hig16 := sum >> 16
	for hig16 != 0 {
		sum = hig16 + uint32(uint16(sum))
		hig16 = sum >> 16
	}

	return uint16(^sum)
}
func main() {
	getCommandArgs()
	destIp := os.Args[len(os.Args)-1]
	conn, err := net.DialTimeout("ip:icmp", destIp, time.Duration(size)*time.Millisecond)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer conn.Close()
	fmt.Printf("PING %s (%s): %d data bytes\n", destIp, conn.RemoteAddr(), size)
	for i := 0; i < count; i++ {
		t := time.Now()
		data := make([]byte, 32)
		icmp := &ICMP{
			Type:        typ,
			Code:        code,
			CheckSum:    0,
			ID:          1,
			SequenceNum: 1,
		}
		var buf bytes.Buffer
		binary.Write(&buf, binary.BigEndian, icmp)
		data = buf.Bytes()
		checkSum := checkSum(data)
		data[2] = byte(checkSum >> 8)
		data[3] = byte(checkSum)
		conn.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Millisecond))
		n, err := conn.Write(data)
		if err != nil {
			log.Fatalln(err)
			continue
		}
		bufRead := make([]byte, 65535)
		n, err = conn.Read(bufRead)
		if err != nil {
			log.Fatalln(err)
			continue
		}
		ts := time.Since(t).Milliseconds()
		fmt.Printf("%d bytes from %d.%d.%d.%d: icmp_seq=%d ttl=%d time=%d ms\n", n, bufRead[12], bufRead[13], bufRead[14], bufRead[15], n-28, bufRead[8], ts)
		time.Sleep(time.Second)
	}
}
