package FastHttpRatelimit

import (
	"github.com/valyala/fasthttp"
	"time"
)

var(
	ratelimits = make(map[string]int)
)

func RatelimitHandler(maxRequests int, timeout time.Duration, next fasthttp.RequestHandler, onRatelimit fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		ip := ctx.RemoteIP().String()

		if _, contains := ratelimits[ip]; contains {
			ratelimits[ip]++
		} else {
			ratelimits[ip] = 1
		}

		go func() {
			time.Sleep(timeout)

			ratelimits[ip]--
			if ratelimits[ip] <= 0 {
				delete(ratelimits, ip)
			}
		}()

		if ratelimits[ip] > maxRequests {
			onRatelimit(ctx)
		} else {
			next(ctx)
		}
	})
}
