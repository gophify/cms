package routes

import (
	"github.com/gorilla/mux"
	"../ctx/mediaCtx"
	"net/http"
)

func MediaRoutes(r *mux.Router) {
	r.PathPrefix("/media/").Handler(http.StripPrefix("/media/", http.FileServer(http.Dir("./media/"))))
    r.HandleFunc("/media-upload", mediaCtx.Upload)
    r.HandleFunc("/media-api", mediaCtx.Api)
}