package api

import (
	"net/http"
	"vercel-go/app"
)

func Handle(w http.ResponseWriter, r *http.Request) { app.Handle(w, r) }
