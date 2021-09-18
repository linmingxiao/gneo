package gneo

import (
	"net/http"
	"github.com/linmingxiao/gneo/skill/httpx"
	"github.com/linmingxiao/gneo/logx"
	"time"
)

type GServer struct {
	*APPConfig
	httpServer *http.Server
	Router *Engine
}

func CreateServer(cfg *APPConfig) *GServer  {
	server := new(GServer)
	SetMode(cfg.RunMode)
	router := New()
	router.Use(Logger())
	if cfg == nil{
		server.APPConfig = &APPConfig{}
	} else {
		server.APPConfig = cfg
		server.Router = router
	}
	return server
}

func (gs *GServer)Listen()  {
	gs.httpServer = &http.Server{
		Addr:           httpx.ResolveAddress([]string{gs.Addr}),
		Handler:        gs.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logx.Info("Server is running: ", gs.Addr)
	gs.httpServer.ListenAndServe()
}

