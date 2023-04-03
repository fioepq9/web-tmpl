package main

import (
	"fmt"
	"os"
	"os/signal"
	"webtmpl/cmd/web/handler/department"
	"webtmpl/cmd/web/util"
	"webtmpl/internal/config"
	"webtmpl/internal/logger"

	ginLogger "github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/pflag"
)

var (
	configfile = pflag.String("config", "/etc/config.toml", "the config file")
)

func main() {
	pflag.Parse()

	// read config
	c := config.Read(*configfile)

	// init global logger
	logger.Init(c)

	// serve
	go serve(c)

	// suspend
	suspend()

}

func serve(c config.Config) {

	r := gin.New()
	r.Use(
		gin.Recovery(),
		ginLogger.SetLogger(
			ginLogger.WithLogger(func(c *gin.Context, _ zerolog.Logger) zerolog.Logger {
				ctx := util.SetTraceIDWithContext(c.Request.Context())
				traceId := util.GetTraceIDFromContext(ctx)
				c.Request = c.Request.WithContext(ctx)
				c.Header("Web-Trace-ID", traceId)
				l := logger.L.With().
					Str("trace_id", traceId).
					Logger()
				c.Request = c.Request.WithContext(l.WithContext(c.Request.Context()))
				return l
			}),
		),
	)

	if c.App.Mode.Stg() || c.App.Mode.Prd() {
		gin.SetMode(gin.ReleaseMode)
	}

	department.Register(r.Group("/department"))

	err := r.Run(fmt.Sprintf("%s:%d", c.App.Host, c.App.Port))

	if err != nil {
		panic(err)
	}
}

func suspend() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch)
	for {
		sig := <-ch
		if sig == os.Interrupt {
			return
		}
	}
}
