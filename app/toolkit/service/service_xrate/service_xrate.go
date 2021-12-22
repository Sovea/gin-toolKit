package service_xrate

import (
	"net/http"
	"sync"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// IPRateLimiter.
type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

var (
	default_limiter = NewIPRateLimiter(1, 15)
)

// NewIPRateLimiter.
func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	i := &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}

	return i
}

func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)

	i.ips[ip] = limiter

	return limiter
}

func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]

	if !exists {
		i.mu.Unlock()
		return i.AddIP(ip)
	}

	i.mu.Unlock()

	return limiter
}

func LimitMiddleware(i *IPRateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get ip address
		ipAddr := c.ClientIP()
		limiter := i.GetLimiter(ipAddr)
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"errcode": -1,
				"errmsg":  "Too Many Requests.",
			})
			c.Abort()
			return

		} else {
			c.Next()
		}
	}
}
func DefaultLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get ip address
		ipAddr := c.ClientIP()
		limiter := default_limiter.GetLimiter(ipAddr)
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"errcode": -1,
				"errmsg":  "Too Many Requests.",
			})
			c.Abort()
			return

		} else {
			c.Next()
		}
	}
}
