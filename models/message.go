package models

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	FormId   uint   //发送者
	TargetID uint   // 接受者
	Type     string // 传播类型 群聊 私聊 广播
	Media    int    // 消息类型 文字 图片 音频
	Content  string // 消息内容
	Pic      string
	Url      string
	Desc     string
	Amount   int // 其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

// 映射关系
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

// 发送者ID 接受者ID 消息类型 发生的内容 发送类型
func Chat(writer http.ResponseWriter, request *http.Request) {
	// 1. 获取参数并校验合法性
	// token := query.Get("token")
	query := request.URL.Query()
	Id := query.Get("userId")
	userId, _ := strconv.ParseInt(Id, 10, 64)
	msgType := query.Get("type")
	targetId := query.Get("targetId")
	context := query.Get("context")
	isvalid := true

	conn, err := (&websocket.Upgrader{
		//token校验
		CheckOrigin: func(r *http.Request) bool {
			return isvalid
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 2. 获取conn
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}

	// 3. 用户关系
	// 4. userid 与node绑定并且加锁
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()

	// 5. 完成发送的逻辑
	go sendProc(node)
	// 6. 完成接受的逻辑
	go recvProc(node)
}

func sendProc(node *Node) {

}

func recvProc(node *Node) {
}
