package timewheel

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestTimeWheel(t *testing.T) {
	begin := time.Now().UnixMilli()
	times := 0
	tw := New(1*time.Second, 3600, func(i interface{}) error {
		times++
		fmt.Printf("%s: %dth tick -- %d\n", i, times, time.Now().UnixMilli()-begin)
		if i.(string) != "LOOP" && i.(string) != "NOOP" {
			return errors.New("fail")
		}
		return nil
	})

	tw.Start()

	tw.AddTimer(true, 1*time.Second, 9028, "LOOP")   // 周期性任务
	tw.AddTimer(true, 1*time.Second, 9029, "LOOP#")  // 周期性任务
	tw.AddTimer(false, 1*time.Second, 9376, "NOOP")  // 单次任务
	tw.AddTimer(false, 1*time.Second, 9377, "NOOP#") // 单次任务

	//time.Sleep(5 * time.Second)
	//
	//tw.RemoveTimer(9028) // 根据索引删除任务
	//fmt.Println("删除")
	//
	//time.Sleep(3 * time.Second)
	//
	//tw.AddTimer(true, 3*time.Second, 9028, "LOOP")
	//fmt.Println("再次添加")
	//
	//time.Sleep(3 * time.Second)
	//
	//tw.Stop()
	//fmt.Println("暂停")
	//
	//time.Sleep(1 * time.Second)
	//
	//tw.Start()
	//fmt.Println("继续")

	time.Sleep(30 * time.Second)
}

func TestHighConcurrent(t *testing.T) {
	tw := New(time.Second, 3600, func(i interface{}) error {
		if (rand.Intn(6)) != 0 {
			fmt.Println("!")
			return errors.New("xxx")
		}
		fmt.Println("~")
		return nil
	})

	tw.Start()

	for i := 0; i < 10000; i++ {
		fmt.Println(i)
		go func() {
			tw.AddTimer(true, time.Second, time.Now().Unix(), "xxx")
		}()
	}

	time.Sleep(time.Minute)
}
