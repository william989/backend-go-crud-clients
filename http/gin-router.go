package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ginRouter struct{}

var (
	ginDispatcher = gin.Default()
)

func NewGinRouter() Router {
	return &ginRouter{}
}

func (*ginRouter) GET(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	ginDispatcher.GET(uri, func(c *gin.Context) {
		f(c.Writer, c.Request)
	})
}

func (*ginRouter) POST(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	ginDispatcher.POST(uri, func(c *gin.Context) {
		f(c.Writer, c.Request)
	})
}

func (*ginRouter) SERVE(port string) {
	fmt.Printf("Gin HTTP server running on port %v", port)
	ginDispatcher.Run(port)
}
