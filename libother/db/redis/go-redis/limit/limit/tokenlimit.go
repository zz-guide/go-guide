package limit

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/tal-tech/go-zero/core/logx"
	xrate "golang.org/x/time/rate"
)

//
//--每秒生成token数量即token生成速度
//local rate = tonumber(ARGV[1])

//--桶容量
//local capacity = tonumber(ARGV[2])

//--当前时间戳
//local now = tonumber(ARGV[3])

//--当前请求token数量
//local requested = tonumber(ARGV[4])

//--需要多少秒才能填满桶
//local fill_time = capacity/rate

//--向下取整,ttl为填满时间的2倍
//local ttl = math.floor(fill_time*2)

//--当前时间桶容量
//local last_tokens = tonumber(redis.call("get", KEYS[1]))

//--如果当前桶容量为0,说明是第一次进入,则默认容量为桶的最大容量
//if last_tokens == nil then
//last_tokens = capacity
//end

//--上一次刷新的时间
//local last_refreshed = tonumber(redis.call("get", KEYS[2]))

//--第一次进入则设置刷新时间为0
//if last_refreshed == nil then
//last_refreshed = 0
//end

//--距离上次请求的时间跨度
//local delta = math.max(0, now-last_refreshed)

//--距离上次请求的时间跨度,总共能生产token的数量,如果超多最大容量则丢弃多余的token
//local filled_tokens = math.min(capacity, last_tokens+(delta*rate))

//--本次请求token数量是否足够
//local allowed = filled_tokens >= requested

//--桶剩余数量
//local new_tokens = filled_tokens

//--允许本次token申请,计算剩余数量
//if allowed then
//new_tokens = filled_tokens - requested
//end

//--设置剩余token数量
//redis.call("setex", KEYS[1], ttl, new_tokens)
//--设置刷新时间
//redis.call("setex", KEYS[2], ttl, now)
//

//return allowed

const (
	// to be compatible with aliyun redis, we cannot use `local key = KEYS[1]` to reuse the key
	// KEYS[1] as tokens_key
	// KEYS[2] as timestamp_key
	script = `local rate = tonumber(ARGV[1])
local capacity = tonumber(ARGV[2])
local now = tonumber(ARGV[3])
local requested = tonumber(ARGV[4])
local fill_time = capacity/rate
local ttl = math.floor(fill_time*2)
local last_tokens = tonumber(redis.call("get", KEYS[1]))
if last_tokens == nil then
    last_tokens = capacity
end

local last_refreshed = tonumber(redis.call("get", KEYS[2]))
if last_refreshed == nil then
    last_refreshed = 0
end

local delta = math.max(0, now-last_refreshed)
local filled_tokens = math.min(capacity, last_tokens+(delta*rate))
local allowed = filled_tokens >= requested
local new_tokens = filled_tokens
if allowed then
    new_tokens = filled_tokens - requested
end

redis.call("setex", KEYS[1], ttl, new_tokens)
redis.call("setex", KEYS[2], ttl, now)

return allowed`
	tokenFormat     = "{%s}.tokens"
	timestampFormat = "{%s}.ts"
)

// A TokenLimiter controls how frequently events are allowed to happen with in one second.
type TokenLimiter struct {
	rate           int
	burst          int
	store          *redis.Client
	tokenKey       string
	timestampKey   string
	rescueLock     sync.Mutex
	redisAlive     uint32
	rescueLimiter  *xrate.Limiter
	monitorStarted bool
}

// NewTokenLimiter returns a new TokenLimiter that allows events up to rate and permits
// bursts of at most burst tokens.
func NewTokenLimiter(rate, burst int, store *redis.Client, key string) *TokenLimiter {
	tokenKey := fmt.Sprintf(tokenFormat, key)
	timestampKey := fmt.Sprintf(timestampFormat, key)

	return &TokenLimiter{
		rate:          rate,
		burst:         burst,
		store:         store,
		tokenKey:      tokenKey,
		timestampKey:  timestampKey,
		redisAlive:    1,
		rescueLimiter: xrate.NewLimiter(xrate.Every(time.Second/time.Duration(rate)), burst),
	}
}

// Allow is shorthand for AllowN(time.Now(), 1).
func (lim *TokenLimiter) Allow() bool {
	return lim.AllowN(time.Now(), 1)
}

// AllowN reports whether n events may happen at time now.
// Use this method if you intend to drop / skip events that exceed the rate rate.
// Otherwise use Reserve or Wait.
func (lim *TokenLimiter) AllowN(now time.Time, n int) bool {
	return lim.reserveN(now, n)
}

func (lim *TokenLimiter) reserveN(now time.Time, n int) bool {
	if atomic.LoadUint32(&lim.redisAlive) == 0 {
		return lim.rescueLimiter.AllowN(now, n)
	}

	resp := lim.store.Eval(
		context.Background(),
		script,
		[]string{
			lim.tokenKey,
			lim.timestampKey,
		},
		[]string{
			strconv.Itoa(lim.rate),
			strconv.Itoa(lim.burst),
			strconv.FormatInt(now.Unix(), 10),
			strconv.Itoa(n),
		})
	// redis allowed == false
	// Lua boolean false -> r Nil bulk reply
	if resp.Err() == redis.Nil {
		return false
	} else if resp.Err() != nil {
		logx.Errorf("fail to use rate limiter: %s, use in-process limiter for rescue", resp.Err())
		return lim.rescueLimiter.AllowN(now, n)
	}

	code, ok := resp.Val().(int64)
	if !ok {
		logx.Errorf("fail to eval redis script: %v, use in-process limiter for rescue", resp)
		return lim.rescueLimiter.AllowN(now, n)
	}

	// redis allowed == true
	// Lua boolean true -> r integer reply with value of 1
	return code == 1
}
