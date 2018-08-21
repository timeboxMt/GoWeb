package controllers

import (
	"time"
)

//NoticeReqest 消息体
type NoticeReqest struct {
	Receiver   string
	Content    string
	CreateDate string
}

//Notice 接收消息Api
func Notice(info NoticeReqest) bool {
	info.CreateDate = time.Now().Format("")
	return true
}

//VerifyTheData 验证消息数据
func VerifyTheData(info NoticeReqest) bool {
	if info.Receiver == "" {
		return false
	}
	if info.Content == "" {
		return false
	}
	return true
}
