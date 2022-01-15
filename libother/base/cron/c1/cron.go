package c1

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func TestTimeAfter() {
	tChan := time.After(time.Second * 3)
	fmt.Printf("tchan type=%T\n", tChan)
	fmt.Println("mark 1")
	fmt.Println("tChan=", <-tChan)
	fmt.Println("mark 2")
}

func TestCron() {
	i := 0
	finishChan := make(chan struct{})

	var queueId cron.EntryID
	var err error

	nyc, _ := time.LoadLocation("Asia/Shanghai")
	var c = cron.New(cron.WithSeconds(), cron.WithLocation(nyc))

	queueId, err = c.AddFunc("*/2 * * * * *", func() {
		i++
		log.Println("cron running:", i)

		if i > 3 {
			//c.Remove(queueId)
			finishChan <- struct{}{}
			log.Println("程序结束")
		}
	})

	fmt.Println(queueId)

	if err != nil {
		log.Println(err)
		return
	}

	c.Start()
	fmt.Println("开始")
	<-finishChan
	fmt.Println("结束")
}
