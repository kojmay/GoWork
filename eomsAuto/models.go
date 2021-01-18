package main

import (
	"time"
)

// 定义一些model

// Ticket 工单模型
type Ticket struct {
	ticketID    string    //工单号
	ticketType  string    //工单类型
	url         string    //链接
	group       string    //所属组
	deadline    time.Time //处理时限
	arrivedTime time.Time //分派时间
	title       string    //工单名称
	status      string    // 工单状态
}
