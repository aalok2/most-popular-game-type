package routes

import (
    "net/http"
    "github.com/gorilla/mux"
    "multiplayer-mode-usage/handler"
    "multiplayer-mode-usage/middleware"
)

func SetupRoutes(router *mux.Router) {
    router.Handle("/api/popular-mode/{area_code}",
        middleware.ValidateAreaCodeMiddleware(http.HandlerFunc(handler.GetPopularModeHandler)),
    ).Methods("GET")
}
