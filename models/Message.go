package models

import (
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"net"
	"net/http"
	"strconv"
	"sync"
)

type Message struct {
	gorm.Model
	FromId   string `json:"from_id"`   //发送者
	TargetId string `json:"target_id"` //接受者
	Type     string `json:"type"`      //消息来源类型 群聊 私聊 广播
	Media    int    `json:"media"`     //消息类型 文字 图片 音频 视频
	Content  string `json:"content"`
	Pic      string `json:"pic"`
	Desc     string `json:"desc"`
	Url      string `json:"url"`
	Amount   int    `json:"amount"` //其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

// 放映射关系
var ClientMap map[int64]*Node = make(map[int64]*Node, 0)
var lock sync.RWMutex

func Chat(wrt http.ResponseWriter, req *http.Request) {
	//校验token
	query := req.URL.Query()
	Id := query.Get("userId")
	UserId, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		return
	}
	//Token := query.Get("token")
	//TargetId := query.Get("targetId")
	//Context := query.Get("context")
	isvalid := true //checkToken()

	conn, err := (&websocket.Upgrader{
		//token校验
		CheckOrigin: func(r *http.Request) bool {
			return isvalid
		},
	}).Upgrade(wrt, req, nil)
	if err != nil {
		fmt.Println(err)
	}
	//获取conn
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 1024),
		GroupSets: set.New(set.ThreadSafe),
	}
	//用户关系
	//userId和node绑定并且加锁
	lock.Lock()
	ClientMap[UserId] = node
	lock.Unlock()

	//完成发送逻辑
	go SendProc(node)
	//完成接收逻辑
	go RecvProc(node)
}

func RecvProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func SendProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
		}
		broadMsg(data)
		fmt.Println("[ws] <<<<<< " + string(data))
	}
}

var updsendChan chan []byte = make(chan []byte, 1024)

func broadMsg(data []byte) {
	updsendChan <- data
}

func init() {
	go udSendProc()
	go udRecvProc()
}

// 完成upd数据发送的协程
func udRecvProc() {
	conn, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   net.IPv4(10, 33, 56, 186),
		Port: 3000,
	})
	if err != nil {
		return
	}
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)
	for {
		select {
		case data := <-updsendChan:
			_, err := conn.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func udSendProc() {
	udp, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		return
	}
	defer func(upd *net.UDPConn) {
		err := udp.Close()
		if err != nil {
			return
		}
	}(udp)
	for {
		var buffer [512]byte
		n, err := udp.Read(buffer[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		dispatch(buffer[:n])
	}
}

// 后端调度逻辑处理
func dispatch(bytes []byte) {

}
