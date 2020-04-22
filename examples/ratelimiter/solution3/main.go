package main

import (
	"net/http"

	"golang.org/x/time/rate"
)

//令牌桶算法
//用户配置的平均发送速率为r，则每隔1/r秒一个令牌被加入到桶中（每秒会有r个令牌放入桶中），桶中最多可以存放b个令牌。
//用户每次请求都会拿走一个令牌，如果令牌到达时令牌桶已经满了，那么这个令牌会被丢弃；如果令牌不足，那么用户请求就会被拒绝
//另一个令牌桶算法库：https://github.com/juju/ratelimit
var limiter = rate.NewLimiter(2, 5)

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)
	// Wrap the servemux with the limit middleware.
	http.ListenAndServe(":4000", limit(mux))
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
