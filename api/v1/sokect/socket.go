package sokect

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
	"os/exec"
	"time"
	"web-server/global"
)

type SocketApi struct {
	upgrader websocket.Upgrader
	client   []*websocket.Conn // 客户端连接池
}

type Message struct {
	Content interface{} `json:"content"`
	Sender  string      `json:"sender"`
	Type    string      `json:"type"`
}

// Websocket 创建 WebSocket 连接 Create WebSocket connection
func (s *SocketApi) Websocket(c *gin.Context) {
	// 将 HTTP 连接升级为 WebSocket 连接
	s.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil) // 使用 upgrader 将 HTTP 连接升级为 WebSocket 连接
	if err != nil {
		global.PalLog.Error("Failed to upgrade to WebSocket: ", zap.Error(err))
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			global.PalLog.Error("Failed to close connection: ", zap.Error(err))
		}
	}(conn) // 在函数返回时关闭连接 defer to close the connection when the function returns

	// 如果客户端连接成功，将连接添加到连接池中 但是如果连接池中已经存在该连接，则不再添加
	// If the client connection is successful, add the connection to the connection pool, but if the connection already exists in the connection pool, do not add it again
	s.client = append(s.client, conn)

	fmt.Println("客户端数量", len(s.client))

	for {

		_, message, err := conn.ReadMessage() // 读取消息 Read message
		if err != nil {
			s.removeClient(conn)
			// 如果发生错误，直接退出循环 If an error occurs, exit the loop directly
			break
		}

		go s.FormatMessageWithCmd(conn, message)

		fmt.Println("Received message: ", string(message))

		// 广播消息
		s.SendMessage(conn, "Hello, Client!")
	}
}

// SendMessage 广播消息 Broadcast message
func (s *SocketApi) SendMessage(conn *websocket.Conn, message string) {
	for _, c := range s.client {
		if c != conn { // 不给发送者发送消息 Do not send messages to the sender
			err := c.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				global.PalLog.Error("Failed to send message: ", zap.Error(err))
				return
			}
		}
	}
}

// FormatMessageWithCmd 格式化cmd消息 Format cmd message
func (s *SocketApi) FormatMessageWithCmd(conn *websocket.Conn, message []byte) {
	var messageData Message
	err := json.Unmarshal(message, &messageData)
	if err != nil {
		return
	}
	// 执行终端命令 Execute terminal command
	if messageData.Type == "cmd" {
		output, err := exec.Command("bash", "-c", messageData.Content.(string)).Output()
		if err != nil {
			global.PalLog.Error("Failed to execute command: ", zap.Error(err))
			return
		}
		// 模拟打字机效果 并返回给发送的客户端 Simulate typewriter effect and return to the sending client
		s.TypewriterEffectWithConn(conn, string(output))
	}
}

// removeClient removes a connection from the client slice
func (s *SocketApi) removeClient(conn *websocket.Conn) {
	for i, c := range s.client { // 遍历连接池，找到要删除的连接 Traverse the connection pool and find the connection to be deleted
		if c == conn { // 删除连接 Remove connection
			// 删除连接 Remove connection
			s.client = append(s.client[:i], s.client[i+1:]...)
			break
		}
	}
}

// TypewriterEffectWithNotConn 打字机效果发送消息 Typewriter effect to send messages
func (s *SocketApi) TypewriterEffectWithNotConn(message string) {
	for _, c := range s.client {
		err := c.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			global.PalLog.Error("Failed to send message: ", zap.Error(err))
			return
		}
		for _, v := range message {
			err := c.WriteMessage(websocket.TextMessage, []byte(string(v)))
			if err != nil {
				global.PalLog.Error("Failed to send message: ", zap.Error(err))
				return
			}
			// 间隔 0.1 秒发送消息 Send message every 0.1 second
			time.Sleep(30 * time.Millisecond)
		}
	}
}

// TypewriterEffectWithConn 打字机效果发送消息 Typewriter effect to send messages
func (s *SocketApi) TypewriterEffectWithConn(conn *websocket.Conn, message string) {
	for _, c := range s.client {
		if c != nil && c == conn {
			err := c.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				global.PalLog.Error("Failed to send message: ", zap.Error(err))
				return
			}
			for _, v := range message {
				err := c.WriteMessage(websocket.TextMessage, []byte(string(v)))
				if err != nil {
					global.PalLog.Error("Failed to send message: ", zap.Error(err))
					return
				}
				// 间隔 0.1 秒发送消息 Send message every 0.1 second
				time.Sleep(30 * time.Millisecond)
			}
		}
	}
}
