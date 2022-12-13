package timewheel

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeWheel(t *testing.T) {
	begin := time.Now().UnixMilli()
	times := 0
	tw := New(1*time.Second, 3600, func(i interface{}) {
		times++
		fmt.Printf("%s: %dth tick -- %d\n", i, times, time.Now().UnixMilli()-begin)
	})

	tw.Start()

	tw.AddTimer(true, 3*time.Second, 9028, "LOOP")  // 周期性任务
	tw.AddTimer(false, 5*time.Second, 9376, "NOOP") // 单次任务

	time.Sleep(5 * time.Second)

	tw.RemoveTimer(9028) // 根据索引删除任务
	fmt.Println("删除")

	time.Sleep(3 * time.Second)

	tw.AddTimer(true, 3*time.Second, 9028, "LOOP")
	fmt.Println("再次添加")

	time.Sleep(3 * time.Second)

	tw.Stop()
	fmt.Println("暂停")

	time.Sleep(1 * time.Second)

	tw.Start()
	fmt.Println("继续")

	time.Sleep(10 * time.Second)
}
