package service

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type (
	ServerContext struct {
		HttpServer *http.Server
		Log        *logrus.Logger
	}
)

func NewServer() *ServerContext {
	return &ServerContext{}
}

func (s *ServerContext) Start() {
	s.loadLog()
	s.loadHttp()
	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, syscall.SIGINT, syscall.SIGTERM)
	for range signChan {
		s.Stop()
		os.Exit(1)
	}
}

func (s *ServerContext) loadLog() {
	s.Log = logrus.New()
	l , err := logrus.ParseLevel("debug")
	if err != nil {
		s.Log.Fatal(err.Error())
		return
	}
	s.Log.Level = l
	s.Log.Formatter = &logrus.JSONFormatter{
		//TimestampFormat:"20060102150405",
	}
}


func (s *ServerContext) loadHttp() {
	r := gin.Default()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.GET("/", func(context *gin.Context) {
		s.Log.Info(context.Params, "======")
		context.Writer.WriteString("Hello World.")
	})

	s.HttpServer = &http.Server{
		Handler:        r,
		MaxHeaderBytes: 10 << 20,
		WriteTimeout:   time.Second * 5,
		ReadTimeout:    time.Second * 5,
		Addr:           ":8081",
	}

	err := s.HttpServer.ListenAndServe()

	if err != nil {
		s.Log.Warn(err.Error())
	}
}

func (s *ServerContext) Stop() {
	//关闭http服务器
	s.HttpServer.Shutdown(nil)
}
