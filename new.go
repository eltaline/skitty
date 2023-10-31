/**
* @program: kitty
*
* @description:
*
* @author: lemon
*
* @create: 2021-10-13 02:23
**/

package kitty

import (
	"github.com/eltaline/skitty/router"
	"github.com/eltaline/skitty/socket"
	"github.com/eltaline/skitty/socket/http"
	httpClient "github.com/eltaline/skitty/socket/http/client"
	httpServer "github.com/eltaline/skitty/socket/http/server"
	tcpClient "github.com/eltaline/skitty/socket/tcp/client"
	tcpServer "github.com/eltaline/skitty/socket/tcp/server"
	udpClient "github.com/eltaline/skitty/socket/udp/client"
	udpServer "github.com/eltaline/skitty/socket/udp/server"
	webSocketClient "github.com/eltaline/skitty/socket/websocket/client"
	webSocketServer "github.com/eltaline/skitty/socket/websocket/server"
)

// HTTP

func NewHttpServer[T any](addr string) *httpServer.Server[T] {
	return &httpServer.Server[T]{Addr: addr}
}

func NewHttpServerRouter[T any]() *router.Router[*http.Stream[httpServer.Conn], T] {
	return &router.Router[*http.Stream[httpServer.Conn], T]{}
}

func NewHttpServerStaticRouter() *httpServer.StaticRouter {
	return &httpServer.StaticRouter{}
}

func NewHttpClientProgress() *httpClient.Progress {
	return &httpClient.Progress{}
}

func NewHttpClient() *httpClient.Client {
	return &httpClient.Client{}
}

// WEB SOCKET

func NewWebSocketClientRouter[T any]() *router.Router[*socket.Stream[webSocketClient.Conn], T] {
	return &router.Router[*socket.Stream[webSocketClient.Conn], T]{}
}

func NewWebSocketClient[T any](addr string) *webSocketClient.Client[T] {
	return &webSocketClient.Client[T]{Addr: addr}
}

func NewWebSocketServer[T any](addr string) *webSocketServer.Server[T] {
	return &webSocketServer.Server[T]{Addr: addr}
}

func NewWebSocketServerRouter[T any]() *router.Router[*socket.Stream[webSocketServer.Conn], T] {
	return &router.Router[*socket.Stream[webSocketServer.Conn], T]{}
}

// UDP

func NewUdpServer[T any](addr string) *udpServer.Server[T] {
	return &udpServer.Server[T]{Addr: addr}
}

func NewUdpServerRouter[T any]() *router.Router[*socket.Stream[udpServer.Conn], T] {
	return &router.Router[*socket.Stream[udpServer.Conn], T]{}
}

func NewUdpClient[T any](addr string) *udpClient.Client[T] {
	return &udpClient.Client[T]{Addr: addr}
}

func NewUdpClientRouter[T any]() *router.Router[*socket.Stream[udpClient.Conn], T] {
	return &router.Router[*socket.Stream[udpClient.Conn], T]{}
}

// TCP

func NewTcpServer[T any](addr string) *tcpServer.Server[T] {
	return &tcpServer.Server[T]{Addr: addr}
}

func NewTcpServerRouter[T any]() *router.Router[*socket.Stream[tcpServer.Conn], T] {
	return &router.Router[*socket.Stream[tcpServer.Conn], T]{}
}

func NewTcpClient[T any](addr string) *tcpClient.Client[T] {
	return &tcpClient.Client[T]{Addr: addr}
}

func NewTcpClientRouter[T any]() *router.Router[*socket.Stream[tcpClient.Conn], T] {
	return &router.Router[*socket.Stream[tcpClient.Conn], T]{}
}
