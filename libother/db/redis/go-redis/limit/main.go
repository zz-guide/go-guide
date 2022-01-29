package main

import (
	"context"
	"log"
	"time"

	"go-guide/libother/db/redis/go-redis/common"
	. "go-guide/libother/db/redis/go-redis/limit/limit"
	xrate "golang.org/x/time/rate"
)

func main() {
	//PeriodLimitTest()
	//TokenLimitTest()
	//xTokenBucketLimit()
	//xTokenBucketLimit1()
	//xTokenBucketLimit2()
}

func PeriodLimitTest() {
	rdb := common.Rdb()
	defer rdb.Close()
	const (
		seconds = 100 // 100s
		quota   = 5   // 允许5次请求
		total   = 100 // 一共发起100次请求
	)

	l := NewPeriodLimit(seconds, quota, rdb, "periodlimit:")
	var allowed, hitQuota, overQuota int
	for i := 0; i < total; i++ {
		val, err := l.Take("userId:1")
		if err != nil {
			log.Println(err)
			return
		}

		switch val {
		case Allowed:
			allowed++
		case HitQuota:
			hitQuota++
		case OverQuota:
			overQuota++
		default:
			log.Println("unknown status")
		}
	}

	log.Println("quota-1 == allowed", quota-1, allowed)
	log.Println("hitQuota == 1", hitQuota)
	log.Println("overQuota == total-quota", overQuota, total-quota)
}

func TokenLimitTest() {
	rdb := common.Rdb()
	defer rdb.Close()
	const (
		total = 100
		rate  = 5
		burst = 10
	)

	l := NewTokenLimiter(rate, burst, rdb, "tokenlimit:")
	var allowed int
	for i := 0; i < 200; i++ {
		time.Sleep(time.Second / time.Duration(total))
		if l.Allow() {
			allowed++
		}
	}

	log.Println("allowed >= burst+rate", allowed, burst+rate)
}

func xTokenBucketLimit() {
	// 初始化一个限速器，每秒产生10个令牌，桶的大小为100个
	// 初始化状态桶是满的
	limiter := xrate.NewLimiter(10, 100)
	for i := 0; i < 10; i++ {
		if limiter.AllowN(time.Now(), 25) {
			log.Printf("%03d Ok  %s\n", i, time.Now().Format("2006-01-02 15:04:05.000"))
		} else {
			log.Printf("%03d Err %s\n", i, time.Now().Format("2006-01-02 15:04:05.000"))
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func xTokenBucketLimit1() {
	// 指定令牌桶大小为5，每秒补充3个令牌
	limiter := xrate.NewLimiter(3, 5)

	// 指定超时时间为5秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	for i := 0; ; i++ {
		log.Printf("%03d %s\n", i, time.Now().Format("2006-01-02 15:04:05.000"))

		// 每次消费2个令牌
		err := limiter.WaitN(ctx, 2)
		if err != nil {
			log.Printf("timeout: %s\n", err.Error())
			return
		}
	}
}

func xTokenBucketLimit2() {
	// 指定令牌桶大小为5，每秒补充3个令牌
	limiter := xrate.NewLimiter(3, 5)

	// 指定超时时间为5秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	for i := 0; ; i++ {
		log.Printf("%03d %sn", i, time.Now().Format("2006-01-02 15:04:05.000"))
		reserve := limiter.Reserve()
		if !reserve.OK() {
			//返回是异常的，不能正常使用
			log.Println("Not allowed to act! Did you remember to set lim.burst to be > 0 ?")
			return
		}
		delayD := reserve.Delay()
		log.Println("休眠:", delayD)
		time.Sleep(delayD)
		select {
		case <-ctx.Done():
			log.Println("超时退出")
			return
		default:
		}
	}
}
