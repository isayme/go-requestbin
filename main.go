package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gdig "github.com/isayme/go-gdig"
	logger "github.com/isayme/go-logger"
	"github.com/isayme/go-requestbin/app/conf"
	"github.com/isayme/go-requestbin/app/middleware"
	"github.com/isayme/go-requestbin/app/router"
	"github.com/isayme/go-requestbin/app/util"
)

var showVersion = flag.Bool("v", false, "show version")
var showHelp = flag.Bool("h", false, "show help")

func main() {
	flag.Parse()

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}

	if *showVersion {
		fmt.Printf("%s: %s\n", util.Name, util.Version)
		os.Exit(0)
	}

	config := conf.Get()

	if config.Logger.Level != "" {
		logger.SetLevel(config.Logger.Level)
	}

	r := gin.New()
	r.Use(middleware.Logger)
	r.Use(gin.Recovery())

	err := gdig.Invoke(func(request *router.Request) error {
		r.Any("/:slug", func(c *gin.Context) {
			if slug := c.Param("slug"); slug == "version" {
				c.JSON(http.StatusOK, map[string]interface{}{
					"name":    util.Name,
					"version": util.Version,
				})
			} else {
				request.RecordRequest(c)
			}
		})
		r.GET("/:slug/inspect", request.ListRequests)
		return nil
	})
	if err != nil {
		panic(err)
	}

	addr := fmt.Sprintf(":%d", config.HTTP.Port)
	logger.Infof("listen %s ...", addr)
	r.Run(addr)
}
