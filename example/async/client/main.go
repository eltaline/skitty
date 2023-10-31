/**
* @program: kitty
*
* @description:
*
* @author: lemon
*
* @create: 2022-05-23 21:58
**/

package main

import (
	"log"

	"github.com/eltaline/skitty"
	"github.com/eltaline/skitty/router"
	"github.com/eltaline/skitty/socket"
	"github.com/eltaline/skitty/socket/tcp/client"
	"github.com/eltaline/skitty/socket/tcp/server"
)

// the same as ws and udp

var tcpServer *server.Server[any]

var tcpClient *client.Client[any]

func asyncTcpServer() {

	var ready = make(chan struct{})

	tcpServer = kitty.NewTcpServer[any]("127.0.0.1:8888")

	var tcpServerRouter = kitty.NewTcpServerRouter[any]()

	tcpServerRouter.Group("/hello").Handler(func(handler *router.Handler[*socket.Stream[server.Conn], any]) {
		handler.Route("/world").Handler(func(stream *socket.Stream[server.Conn]) error {
			return stream.Emit(stream.Event(), stream.Data())
		})
	})

	tcpServer.OnSuccess = func() {
		ready <- struct{}{}
	}

	go tcpServer.SetRouter(tcpServerRouter).Start()

	<-ready
}

func asyncTcpClient() {

	var ready = make(chan struct{})
	var isRun = false

	tcpClient = kitty.NewTcpClient[any]("127.0.0.1:8888")

	var clientRouter = kitty.NewTcpClientRouter[any]()

	tcpClient.OnSuccess = func() {
		if isRun {
			return
		}
		ready <- struct{}{}
	}

	go tcpClient.SetRouter(clientRouter).Connect()

	<-ready
	isRun = true
}

func main() {
	asyncTcpServer()
	asyncTcpClient()

	var asyncClient = socket.NewAsyncClient[client.Conn, any](tcpClient)

	var stream, err = asyncClient.Emit("/hello/world", []byte("hello world"))

	log.Println(string(stream.Data()), err)

	select {}
}
