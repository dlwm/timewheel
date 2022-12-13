# timewheel
Golang实现的时间轮

[![Go Report Card](https://goreportcard.com/badge/github.com/ouqiang/timewheel)](https://goreportcard.com/report/github.com/ouqiang/timewheel)

![时间轮](https://raw.githubusercontent.com/ouqiang/timewheel/master/timewheel.jpg)

# 原理
[延迟消息的实现](http://www.10tiao.com/html/249/201703/2651959961/1.html)

# 安装

```shell
go get -u github.com/dlwm/timewheel
```

# 使用

```go
package main

import (
	"fmt"
	"github.com/dlwm/timewheel"
	"time"
)

func main() {
	// 初始化时间轮
	// 第一个参数为tick刻度，即时间轮多久转动一次，最小为1秒，跨度1秒
	// 第二个参数为时间轮槽slot数量，即每轮有多少个记录点
	// 第三个参数为回调函数
	tw := timewheel.New(1*time.Second, 3600, func(i interface{}) {
		fmt.Println(i)
	})

	// 开启时间轮
	tw.Start()

	// 添加定时器 
	// 第一个参数为周期性开启标记
	// 第二个参数为延迟时间
	// 第三个参数为定时器唯一标识，删除定时器需传递此参数
	// 第四个参数为用户自定义数据，此参数将会传递给回调函数, 类型为interface{}
	tw.AddTimer(true, 3*time.Second, 9028, "LOOP")  // 周期性任务
	tw.AddTimer(false, 5*time.Second, 9376, "NOOP") // 单次任务
	
	// 根据索引删除任务，参数为添加定时器传递的唯一标识
	tw.RemoveTimer(9028)
	
	// 暂停时间轮
	tw.Stop()

	// 继续时间轮
	tw.Start()

	select {}
}
```

# 添加
