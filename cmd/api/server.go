package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *application) run() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      app.routes(),
	}

	return srv.ListenAndServe()
}
