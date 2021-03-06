package looli

import (
	"fmt"
	"io"
	"os"
	"time"
)

var defaultWriter = os.Stdout

func Logger() HandlerFunc {
	return LoggerWithWriter(defaultWriter)
}

func LoggerWithWriter(out io.Writer) HandlerFunc {
	return func(c *Context) {
		start := time.Now()
		path := c.Path
		method := c.Request.Method
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		statusCode := c.statusCode
		proto := c.Request.Proto

		fmt.Fprintf(out, "[looli] %v | %3d | %11v | %s | %-4s %-8s %s\n",
			end.Format("2006/01/01 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			method,
			proto,
			path,
		)
	}
}
