package http

import (
	"context"
	"example/test-golang/config"
	"example/test-golang/controller"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpSrv *http.Server
}

func (s Server) Start() error {
	return s.httpSrv.ListenAndServe()
}

func (s Server) Shutdown(ctx context.Context) error {
	return s.httpSrv.Shutdown(ctx)
}

func ProvideServer(cfg *config.Config, uc controller.UserController) *Server {
	r := gin.Default()
	r.GET("/user/:id", uc.GetUserById )

	port := getPortString(cfg.ServerPort)
	srv := &http.Server{Addr: port, Handler: r}

	return &Server{httpSrv: srv}
}

func getPortString(port int) string {
	return fmt.Sprintf(":%d", port)
}
