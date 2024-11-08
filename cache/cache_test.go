package cache

import (
    "testing"
    "time"
    "multiplayer-mode-usage/cache"
    "github.com/stretchr/testify/assert"
)

// Redis Test Configurations
const testRedisAddr = "localhost:6379"

func setupTestCache() {
    cache.InitRedis(testRedisAddr)
}

func TestSetAndGetPopularModeCache_Success(t *testing.T) {
    setupTestCache()

    // Set data in the cache
    err := cache.SetPopularModeCache("NY", "Battle Royale", 150, 1*time.Minute)
    assert.NoError(t, err)

    // Retrieve data from cache
    mode, count, err := cache.GetPopularModeCache("NY")
    assert.NoError(t, err)
    assert.Equal(t, "Battle Royale", mode)
    assert.Equal(t, int32(150), count)
}

func TestGetPopularModeCache_CacheMiss(t *testing.T) {
    setupTestCache()

    // Retrieve non-existent key
    mode, count, err := cache.GetPopularModeCache("Unknown")
    assert.Error(t, err)
    assert.Equal(t, "", mode)
    assert.Equal(t, int32(0), count)
}
