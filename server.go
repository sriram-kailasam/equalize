package main

import (
	"fmt"
	"net/http"

	"github.com/sriram-kailasam/equalize/config"
)

func createServer(serverConfig config.Server) *http.Server {
	mux := http.NewServeMux()

	for _, location := range serverConfig.Locations {
		fmt.Printf("Registering route %s\n", location.Path)
		mux.HandleFunc(location.Path, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(int(location.Response.Status))
			fmt.Fprint(w, location.Response.Data)
		})
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", serverConfig.Port),
		Handler: mux,
	}

	return srv
}

func startServers(conf config.Config) {

	for _, serverConfig := range conf.Http.Servers {
		func() {
			srv := createServer(serverConfig)

			fmt.Printf("Starting server %s on port: %d\n", serverConfig.Name, serverConfig.Port)
			go srv.ListenAndServe()
		}()
	}
}
