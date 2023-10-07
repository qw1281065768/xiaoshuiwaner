package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

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

	// 传入request和responseWriter
	server := officialAccount.GetServer(req, rw)
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
		fmt.Println(err)
		return
	}
	fmt.Println(server.ResponseMsg)
	//发送回复的消息
	err = server.Send()
	fmt.Println(err)
}
