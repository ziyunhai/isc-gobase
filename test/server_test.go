package test

import (
	"github.com/isyscore/isc-gobase/cron"
	"testing"

	"github.com/isyscore/isc-gobase/logger"
	"github.com/isyscore/isc-gobase/server"
)

func TestServer(t *testing.T) {
	server.InitServer()

	server.RegisterCustomHealthCheck("/api/sample",
		func() string {
			return "OK"
		},
		func() string {
			return "OK"
		},
		func() string {
			return "OK"
		},
	)

	logger.Info("server started")

	go func() {
		for i := 0; i < 100; i++ {
			go func(idx int) {
				c_s := cron.New()
				c_s.AddFunc("*/1 * * * * ?", func() {
					logger.Debug("协程ID=：%d,我是库陈胜Debug", idx)
					logger.Info("协程ID=：%d,我是库陈胜1", idx)
					logger.Warn("协程ID=：%d,我是库陈胜2", idx)
					logger.Error("协程ID=：%d,我是库陈胜3", idx)
				})
				c_s.Start()
			}(i)
		}
	}()

	server.StartServer()
}
