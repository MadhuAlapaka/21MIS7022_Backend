package middleware

import (
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var ctx = context.Background()

func RateLimiter(rdb *redis.Client, limit int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr
			count, err := rdb.Incr(ctx, ip).Result()
			if err != nil {
				http.Error(w, "Rate limiter error", http.StatusInternalServerError)
				return
			}

			if count == 1 {
				rdb.Expire(ctx, ip, time.Minute)
			}

			if count > int64(limit) {
				http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
