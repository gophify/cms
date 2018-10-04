package routes

import (
	"github.com/gorilla/mux"
	"../ctx/authCtx"
)

func AuthRoutes(r *mux.Router) {
    r.HandleFunc("/auth/api", authCtx.Api)
    r.HandleFunc("/secret", authCtx.Secret)
    r.HandleFunc("/login", authCtx.Login)
    r.HandleFunc("/logout-api", authCtx.LogoutApi)
    r.HandleFunc("/userlog-api", authCtx.UserLog)
}