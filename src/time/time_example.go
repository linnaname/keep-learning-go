package main

import (
	"fmt"
	"time"
)

func main() {

	/**
	时间点
	*/
	now := time.Now() //获取当前时间
	fmt.Println(now)
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	fmt.Println(now.Unix())

	loc, _ := time.LoadLocation("America/Los_Angeles")
	fmt.Println(time.Now().In(loc))

	/**
	format
	*/
	//这个固定时间真的是个吃错药的设计
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format(time.UnixDate))

	tm, _ := time.Parse("2006-01-02 15:04:05", "2018-04-23 12:24:51")
	fmt.Println(tm)

	/**
	Duration
	*/

	fmt.Println(time.Second)
	tp, _ := time.ParseDuration("1.5s")
	fmt.Println(tp.Truncate(1000), tp.Seconds(), tp.Nanoseconds())

	/**
	时间运算
	*/
	//time.Sleep(time.Duration(10) * time.Second)
	//time.After(time.Duration(10) * time.Second)
	start := time.Now()
	fmt.Println(time.Since(start))
	fmt.Println(start.Add(time.Duration(10) * time.Second)) // 加
	fmt.Println(start.AddDate(1, 1, 1))

	dt1 := time.Date(2018, 1, 10, 0, 0, 1, 100, time.Local)
	dt2 := time.Date(2018, 1, 9, 23, 59, 22, 100, time.Local)
	// 不用关注时区，go会转换成时间戳进行计算
	fmt.Println(dt1.Sub(dt2))

	dt := time.Date(2018, 1, 10, 0, 0, 1, 100, time.Local)
	fmt.Println(time.Now().After(dt))  // true
	fmt.Println(time.Now().Before(dt)) // false

	// 是否相等 判断两个时间点是否相等时推荐使用 Equal 函数
	fmt.Println(dt.Equal(time.Now()))

	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	fmt.Println(time.Date(2018, 1, 2, 0, 0, 0, 0, beijing)) // 2018-01-02 00:00:00 +0800 Beijing Time

	// 当前时间转为指定时区时间
	fmt.Println(time.Now().In(beijing))

	//定时任务
	//ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	//for i := range ticker {
	//	fmt.Println(i)//每秒都会执行的任务
	//}

	timer := time.NewTimer(time.Second * 2)
	//获取当前时间
	t1 := time.Now()
	//接收通道数据
	t2 := <-timer.C

	fmt.Println(t1)
	fmt.Println(t2)

	timer3 := time.AfterFunc(time.Second*2, func() {
		fmt.Println("2秒函数执行了")
	})
	timer3.Reset(time.Second * 2)

	time.Sleep(time.Second * 3)

}
