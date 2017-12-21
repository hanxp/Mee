package main

import (
	"time"
	"fmt"
)

const (
	// TimeFormatDefault 时间默认格式
	TimeFormatDefault = "2006-01-02 15:04:05"
	// TimeFormatDate 日期格式
	TimeFormatDate = "2006-01-02"
)

func main() {

	tm1:="2017-11-30T14:00:13+08:00"
	tm2:="2017-11-30 14:00:13"

	// 本地时间
	local, _ := time.LoadLocation("Local")
	// 截止日期
	t1, _ := time.ParseInLocation(TimeFormatDefault, tm1, local)
	t2, _ := time.ParseInLocation(TimeFormatDefault, tm2, local)
	fmt.Println(t1)
	fmt.Println(t2)

}
