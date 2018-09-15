package main

import (
	"go-simple-blog/service"
)

func main () {
	serv := service.NewServer()
	serv.Start()
}