package main

import (
	"context"
	"flag"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/iyidan/gindemo/api"
	"github.com/iyidan/gindemo/conf"
	"github.com/iyidan/gindemo/log"
	"github.com/iyidan/gindemo/middleware"
	"github.com/iyidan/gindemo/models"
)

func main() {
	// register signal handler
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		os.Kill,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	// startup
	flagParse()
	conf.Startup()
	log.Startup()
	models.Startup()
	defer models.Close()

	// init server
	_, shutdownCh, closedCh := initServer()

	// blocked until received a signal or server closed
	select {
	case s := <-sc:
		log.Errorf("main: received signal: %v", s)
		close(shutdownCh)
		<-closedCh
	case err := <-closedCh:
		log.Errorf("received error from server closedCh: %s", err)
	}
}

func flagParse() {
	var echoVersion bool
	flag.BoolVar(&echoVersion, "v", false, "-v get current version")
	flag.Parse()
	// echo version
	if echoVersion {
		os.Stdout.WriteString(VERSION + "\n")
		os.Exit(0)
	}
}

func initServer() (*http.Server, chan<- struct{}, <-chan error) {
	if conf.Bool("pprof") {
		// open pprof
		go func() {
			log.Fatal(http.ListenAndServe(conf.String("pprofAddr"), nil))
		}()
	}
	// production model
	if conf.IsOnProd() {
		gin.SetMode(gin.ReleaseMode)
	}

	// register middleware and apis
	router := gin.New()
	router.Use(
		gin.Recovery(),
		middleware.Accesslog(),
		middleware.MaxBodyLimit(conf.Int64("maxBodySize")))
	api.Register(router)

	// start server
	shutdownCh := make(chan struct{})
	closedCh := make(chan error, 1)
	listenErrch := make(chan error, 1)
	s := &http.Server{
		Addr:           conf.String("addr"),
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}
	// dev connState debug
	if conf.IsOnDev() {
		s.ConnState = func(conn net.Conn, state http.ConnState) {
			log.Infof("%s -> %s", conn.RemoteAddr().String(), state.String())
		}
	}
	go func() {
		err := s.ListenAndServe()
		listenErrch <- err
		close(listenErrch)
	}()

	// graceful shutdown
	go func() {
		select {
		case <-shutdownCh:
			timeout := time.Second * 60
			ctx, cancelf := context.WithTimeout(context.Background(), timeout)
			err := s.Shutdown(ctx)
			cancelf()
			if err != nil {
				log.Errorf("server shutdown error: %s", err)
			} else {
				log.Info("server shutdown ok")
			}
			err = <-listenErrch
			closedCh <- err
			close(closedCh)
		case err := <-listenErrch:
			closedCh <- err
			close(closedCh)
		}
	}()

	return s, shutdownCh, closedCh
}
