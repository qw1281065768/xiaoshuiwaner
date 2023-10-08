package service

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Message struct {
	MsgId        int64  `json:"MsgId"`
	Content      string `json:"Content"`
	MsgType      string `json:"MsgType"`
	CreateTime   int    `json:"CreateTime"`
	FromUserName string `json:"FromUserName"`
	ToUserName   string `json:"ToUserName"`
}

// WXRepTextMsg 微信回复文本消息结构体
type WXRepTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	// 若不标记XMLName, 则解析后的xml名为该结构体的名称
	XMLName xml.Name `xml:"xml"`
}

// ReplyMessageHandler 回复消息接口
func ReplyMessageHandler(rw http.ResponseWriter, req *http.Request) {
	wc := wechat.NewWechat()
	//这里本地内存保存access_token，也可选择redis，memcache或者自定cache
	memory := cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:     "wx4c448f0d8141b6e6",
		AppSecret: "AAQ9G7sEAAABAAAAAADYEl7EIFIdKbnYY0AhZSAAAAAraAdzKm8i8JwFJ68cDOJBtIHsmv3F8e00LCv7f+tYviC+CNGw8id1pIHomelyP7+eaEYJe3Gq577QcgzWXoTs5Ak+WXMnG/pB5INCH3ZqUuGmKsY9XhzYleN6sVUQ2qyeOqHtN81coOUizp61GIVr7LtT347AcTgRHisbQZGQ",
		//Token:     "qw7843251",
		//EncodingAESKey: "xxxx",
		Cache: memory,
	}
	officialAccount := wc.GetOfficialAccount(cfg)

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(reqBody))
	input := &Message{}
	err = json.Unmarshal(reqBody, input)
	if err != nil {
		fmt.Println("111", err)
		return
	}

	repTextMsg := WXRepTextMsg{
		ToUserName:   input.FromUserName,
		FromUserName: input.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      fmt.Sprintf("[消息回复] - %s", time.Now().Format("2006-01-02 15:04:05")),
	}

	msg, err := xml.Marshal(&repTextMsg)
	if err != nil {
		fmt.Println("222", err)
		return
	}
	server := officialAccount.GetServer(req, rw)
	server.XML(msg)

	// 传入request和responseWriter

	/*server := officialAccount.GetServer(req, rw)

	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		//TODO
		//回复消息：演示回复用户发送的消息
		fmt.Println(msg.Content)
		fmt.Println(msg)
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	server.SkipValidate(true)
	err = server.Serve()
	if err != nil {
		fmt.Println("111", err)
		return
	}
	fmt.Println("999", server)
	//发送回复的消息
	err = server.Send()
	fmt.Println("222", err)*/
}
