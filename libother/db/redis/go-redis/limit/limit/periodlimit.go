package limit

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

//0：表示错误，比如可能是redis故障、过载
//1：允许
//2：允许但是当前窗口内已到达上限，如果是跑批业务的话此时可以休眠sleep一下等待下个窗口（作者考虑的非常细致）
//3：拒绝

//--KYES[1]:限流器key
//--ARGV[1]:qos,单位时间内最多请求次数
//--ARGV[2]:单位限流窗口时间
//--请求最大次数,等于p.quota
//local limit = tonumber(ARGV[1])
//--窗口即一个单位限流周期,这里用过期模拟窗口效果,等于p.permit
//local window = tonumber(ARGV[2])
//--请求次数+1,获取请求总数
//local current = redis.call("INCRBY",KYES[1],1)
//--如果是第一次请求,则设置过期时间并返回 成功
//if current == 1 then
//redis.call("expire",KYES[1],window)
//return 1
//--如果当前请求数量小于limit则返回 成功
//elseif current < limit then
//return 1
//--如果当前请求数量==limit则返回 最后一次请求
//elseif current == limit then
//return 2
//--请求数量>limit则返回 失败
//else
//return 0
//end

const (
	// to be compatible with aliyun redis, we cannot use `local key = KEYS[1]` to reuse the key
	periodScript = `local limit = tonumber(ARGV[1])
local window = tonumber(ARGV[2])
local current = redis.call("INCRBY", KEYS[1], 1)
if current == 1 then
    redis.call("expire", KEYS[1], window)
    return 1
elseif current < limit then
    return 1
elseif current == limit then
    return 2
else
    return 0
end`
	zoneDiff = 3600 * 8 // GMT+8 for our services
)

const (
	Unknown = iota
	Allowed
	HitQuota
	OverQuota

	internalOverQuota = 0
	internalAllowed   = 1
	internalHitQuota  = 2
)

var ErrUnknownCode = errors.New("unknown status code")

type (
	LimitOption func(l *PeriodLimit)

	PeriodLimit struct {
		//窗口大小，单位s
		period int
		//请求上限
		quota      int
		limitStore *redis.Client
		keyPrefix  string
		//线性限流，开启此选项后可以实现周期性的限流
		//比如quota=5时，quota实际值可能会是5.4.3.2.1呈现出周期性变化
		align bool
	}
)

func NewPeriodLimit(period, quota int, limitStore *redis.Client, keyPrefix string,
	opts ...LimitOption) *PeriodLimit {
	limiter := &PeriodLimit{
		period:     period,
		quota:      quota,
		limitStore: limitStore,
		keyPrefix:  keyPrefix,
	}

	for _, opt := range opts {
		opt(limiter)
	}

	return limiter
}

func (h *PeriodLimit) Take(key string) (int, error) {
	resp := h.limitStore.Eval(context.Background(), periodScript, []string{h.keyPrefix + key}, []string{
		strconv.Itoa(h.quota),
		strconv.Itoa(h.calcExpireSeconds()),
	})

	if resp.Err() != nil {
		return Unknown, resp.Err()
	}

	code, ok := resp.Val().(int64)
	if !ok {
		return Unknown, ErrUnknownCode
	}

	switch code {
	case internalOverQuota:
		return OverQuota, nil
	case internalAllowed:
		return Allowed, nil
	case internalHitQuota:
		return HitQuota, nil
	default:
		return Unknown, ErrUnknownCode
	}
}

func (h *PeriodLimit) calcExpireSeconds() int {
	if h.align {
		unix := time.Now().Unix() + zoneDiff
		return h.period - int(unix%int64(h.period))
	} else {
		return h.period
	}
}

func Align() LimitOption {
	return func(l *PeriodLimit) {
		l.align = true
	}
}
