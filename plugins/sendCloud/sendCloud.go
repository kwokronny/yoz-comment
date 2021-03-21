package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	//"strings"
)

type mailInfo struct {
	To string
}

func SendMail(info mailInfo) {
	// json.Marshal()
	RequestURI := "http://api.sendcloud.net/apiv2/mail/sendtemplate"
	PostParams := url.Values{
		"apiUser":            {"kwokronny_notice"},
		"apiKey":             {"3d6fa4576e3c60e943e0cf4e9fbdb049"},
		"from":               {"no-reply@notice.kwokronny.top"},
		"fromName":           {"KwokRonny"},
		"xsmtpapi":           {"{\"sub\":{\"%you%\"}}"},
		"to":                 {info.To}, //to is address list
		"subject":            {"你在 KwokRonny 博客上的留言有回复啦"},
		"templateInvokeName": {"kwok_comment_notice"},
		"useAddressList":     {"true"},
	}
	PostBody := bytes.NewBufferString(PostParams.Encode())
	ResponseHandler, err := http.Post(RequestURI, "application/x-www-form-urlencoded", PostBody)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ResponseHandler.Body.Close()
	BodyByte, err := ioutil.ReadAll(ResponseHandler.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(BodyByte))
}
