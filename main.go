package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "multiplayer-mode-usage/config"
    "multiplayer-mode-usage/db"
    "multiplayer-mode-usage/cache"
    "multiplayer-mode-usage/routes"
)

func main() {

    cfg := config.LoadConfig()

    fmt.Printf("Loaded Config: Redis Address - %s, Mongo URI - %s, Mongo DB Name - %s\n", cfg.RedisAddress, cfg.MongoURI, cfg.MongoDBName)

    cache.InitRedis(cfg.RedisAddress)

    db.InitializeMongoDB(cfg.MongoURI, cfg.MongoDBName)

    router := mux.NewRouter()

    routes.SetupRoutes(router)

    log.Printf("Server starting on port 8080...")

    err := http.ListenAndServe(":8080" , router)
    if err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
    log.Println("Server successfully connected and running on port 8080.")
}
