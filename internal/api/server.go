package api

import (
	"context"
	"net/http"
	"github.com/appu900/TurboQueue/config"
	"github.com/appu900/TurboQueue/internal/model"
	"github.com/appu900/TurboQueue/internal/store"
	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg        *config.Config
	router     *gin.Engine
	httpServer *http.Server
	store      store.Store
	msgChans   map[string][]chan *model.Message
}

func NewServer(cfg *config.Config) *Server {
	newServerInstance := &Server{
		cfg:      cfg,
		router:   gin.Default(),
		msgChans: make(map[string][]chan *model.Message),
	}
	return newServerInstance

}

func (s *Server) Start(ctx context.Context) error {
	s.httpServer = &http.Server{
		Addr: ":" + s.cfg.Port,
	}
	return s.httpServer.ListenAndServe()
}
