# What is it?
FastHttpRatelimit simply provides middleware for FastHttp to enforce customisable IP based ratelimits.

# How to use 
FastHttpRatelimit is very easy to use, take a look at the example below. A function called `RatelimitHandler` is provided which takes the following paramters:
- `maxRequests`: The amount of requests that are allowed in the timeframe before the ratelimit is enforced.
- `timeout`: The timeframe for which requests should be measured in.
- `next`: Your request handler.
- `onRatelimit`: The handler to be executed when a user is ratelimited.

# Motivation
In many of my projects I have needed a simple ratelimiter for FastHTTP to prevent bruteforcing and other issues.

I am releasing it as its own module to make it easier to use for myself, and so that others are able to use it.

# Example
```go
import (
	"github.com/Dot-Rar/FastHttpRatelimit"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"time"
)

func Main() {
	router := routing.New()

	handler := RatelimitHandler(
		25, // Max requests in the time period
		30 * time.Second, // The time period for which requests should be logged
		fasthttp.CompressHandler(MyHandler), // The handler to be executed upon the user not being ratelimited (let them continue)
		fasthttp.CompressHandler(ErrorHandler), // The handler to be executed upon the user being ratelimited (tell them) 
		)
	
	router.Get("/", func(ctx *routing.Context) error {
		handler(ctx.RequestCtx)
		return nil
	})
}
```