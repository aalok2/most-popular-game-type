package handler

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gorilla/mux"
    "multiplayer-mode-usage/handler"
    "multiplayer-mode-usage/service"
    "multiplayer-mode-usage/cache"
    "github.com/stretchr/testify/assert"
)

func setupTestHandler() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/popular/{area_code}", handler.GetPopularModeHandler).Methods("GET")
    return r
}

func TestGetPopularModeHandler_CacheHit(t *testing.T) {
    cache.InitRedis("localhost:6379")
    cache.SetPopularModeCache("CA", "Battle Royale", 300, 5*time.Minute)

    r := setupTestHandler()
    req, _ := http.NewRequest("GET", "/popular/CA", nil)
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), `"Battle Royale"`)
    assert.Contains(t, w.Body.String(), `"player_count":300`)
}

func TestGetPopularModeHandler_CacheMiss_DBHit(t *testing.T) {
    service.GetPopularMode("NY")  // Assuming DB has this value

    r := setupTestHandler()
    req, _ := http.NewRequest("GET", "/popular/NY", nil)
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), `"Battle Royale"`) // Assuming a DB record exists for "NY" with mode "Battle Royale"
}
