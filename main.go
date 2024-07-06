package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/quincy0/live-ai/router"
	"github.com/quincy0/live-ai/util"
	"github.com/quincy0/qpro/qConfig"
	"github.com/quincy0/qpro/qLog"
	"github.com/quincy0/qpro/qRedis"
	"github.com/quincy0/qpro/qRoutine"
	qTrace "github.com/quincy0/qpro/qTrace"
	"github.com/quincy0/qpro/qdb"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	_ "github.com/quincy0/live-ai/store"
)

var (
	configPath string
	port       string
	mode       string
	StartCmd   = &cobra.Command{
		Use:   "server",
		Short: "Start API server",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configPath, "configPath", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&port, "port", "p", "9003", "Tcp port server listening on")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode ; eg:dev,test,prod")
}

func usage() {
	usageStr := `starting api server`
	log.Printf("%s\n", usageStr)
}

func setup() {
	//ctx := util.InitContext()
	// 1. 读取配置
	qConfig.Setup(configPath)

	// 2. 日志配置
	qLog.Init(qConfig.Settings.Log)
	// 3. 初始化数据库链接
	qdb.InitMysql()
	// 4. 初始化Redis
	qRedis.InitRedis()
}

func run() error {
	cleanup := qTrace.InitTracer()
	defer func() {
		err := cleanup(context.Background())
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	r := router.Init()

	srv := &http.Server{
		Addr:         qConfig.Settings.Application.Host + ":" + strconv.Itoa(qConfig.Settings.Application.Port),
		Handler:      r,
		ReadTimeout:  qConfig.Settings.Application.ReadTimeoutSeconds * time.Second,
		WriteTimeout: qConfig.Settings.Application.WriteTimeoutSeconds * time.Second,
	}

	qRoutine.GoSafe(func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			qLog.Fatal("listen error", zap.Any("error", err))
		}
	})
	fmt.Printf("%s Server Run http://%s:%v/ \r\n",
		util.CurrentTimeStr(),
		qConfig.Settings.Application.Host,
		qConfig.Settings.Application.Port)
	fmt.Printf("%s Swagger URL http://%s:%v/swagger/index.html \r\n",
		util.CurrentTimeStr(),
		qConfig.Settings.Application.Host,
		qConfig.Settings.Application.Port)

	fmt.Printf("%s Enter Control + C Shutdown Server \r\n", util.CurrentTimeStr())

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	fmt.Printf("%s Shutdown Server ... \r\n", util.CurrentTimeStr())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		qLog.Fatal("Server Shutdown", zap.Any("error", err))
	}
	qLog.Info("Server exiting")
	return nil
}

func main() {
	StartCmd.Execute()
}
