package main

import (
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/silenceper/dcmp/etcd"
)

var cfg *config

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	cfg = getConfig()

	etcd.InitConfig(cfg.Endpoints)

	r := gin.Default()
	registerRouter(r)
	r.Run(cfg.Listen)
}
