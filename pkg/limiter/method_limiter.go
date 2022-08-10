package limiter

import (
	"strings"
	"time"

	"BlogService/global"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

var buckets []BucketRule

type MethodLimiter struct {
	*Limiter
}

func NewMethodLimiter() InterfaceLimiter {
	initBuckets()
	l := &Limiter{map[string]*ratelimit.Bucket{}}
	return MethodLimiter{l}.AddBuckets(buckets...)
}

// Key 以路由的相对路径作为令牌桶映射的键
func (l MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index]
}

// GetBucket 获取接口的令牌桶
func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.limiterBuckets[key]
	return bucket, ok
}

// AddBuckets 添加接口的令牌桶
func (l MethodLimiter) AddBuckets(rules ...BucketRule) InterfaceLimiter {
	for _, rule := range rules {
		if _, ok := l.limiterBuckets[rule.Key]; !ok {
			l.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quantum)
		}
	}
	return l
}

//读取配置文件，初始化令牌桶
func initBuckets() {
	for _, lv := range global.Config.Limiter {
		buckets = append(buckets, BucketRule{
			Key:          lv.Key,
			FillInterval: lv.FillInterval * time.Second,
			Capacity:     lv.Capacity,
			Quantum:      lv.Quantum,
		})
	}
}
