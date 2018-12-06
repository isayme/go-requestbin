package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	logger "github.com/isayme/go-logger"
	"github.com/isayme/go-requestbin/app"
	"github.com/isayme/go-requestbin/app/conf"
	"github.com/isayme/go-requestbin/app/manager"
	"github.com/isayme/go-requestbin/app/middleware"
	"github.com/isayme/go-requestbin/app/router"
)

var configPath = flag.String("c", "/etc/requestbin.json", "config file path")
var showVersion = flag.Bool("v", false, "show version")
var showHelp = flag.Bool("h", false, "show help")

func main() {
	flag.Parse()

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}

	if *showVersion {
		fmt.Printf("%s: %s\n", app.Name, app.Version)
		os.Exit(0)
	}

	conf.SetPath(*configPath)
	config := conf.Get()

	if config.Logger.Level != "" {
		logger.SetLevel(config.Logger.Level)
	}

	manager.Init(config)

	r := gin.New()
	r.Use(middleware.Logger)
	r.Use(gin.Recovery())

	r.Any("/:slug", router.RecordRequest)
	r.GET("/:slug/inspect", router.ListRequests)
	r.Run(fmt.Sprintf(":%d", config.HTTP.Port))
}
