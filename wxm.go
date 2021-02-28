package wxm

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

const URL = "http://www.bangnikanzhe.com/wx/sendstockmsg"

type WxMsg struct {
	Userid  string `json:"userid"`
	Userpwd string `json:"userpwd"`
	Mark    string `json:"mark"`
	First   string `json:"first"`
	Kw1     string `json:"kw1"`
	Kw2     string `json:"kw2"`
	Kw3     string `json:"kw3"`
}

var (
	userName string = ""
	userPwd  string = ""
)

func Init(name string, pwd string) {
	userName = name
	userPwd = pwd
}

func SendMsg(txt1, txt2, txt3 string) (string, error) {
	if userName == "" {
		return "", errors.New("not init user name, call Init(name, pwd)")
	}
	msgData := makeData(userName, userPwd, txt1, txt2, txt3)
	result, err := post(msgData)
	return result, err

}

func SendMsgToUser(name, pwd, txt1, txt2, txt3 string) (string, error) {
	msgData := makeData(name, pwd, txt1, txt2, txt3)
	return post(msgData)

}

///////////////////////////////////////////////////////////////////////
// 内部函数

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func post(data interface{}) (string, error) { //url string,  , contentType string

	contentType := "application/json"
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(URL, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result), nil
}

func makeData(name, pwd, txt1, txt2, txt3 string) *WxMsg {
	msg := &WxMsg{Userid: name, Userpwd: pwd, First: txt1, Kw2: txt2, Kw3: txt3}
	return msg
}
