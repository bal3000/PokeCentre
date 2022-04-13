package main

import "net/http"

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	env := envelope[map[string]string]{
		"application": map[string]string{
			"status": "ok",
		},
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     "1.0.0",
		},
	}

	err := writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
