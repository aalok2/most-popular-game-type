package service

import (
    "log"
    "time"
    "multiplayer-mode-usage/cache"
    "multiplayer-mode-usage/db"
)

func GetPopularMode(areaCode string) (string, int32, error) {
    mode, count, err := cache.GetPopularModeCache(areaCode)
    if err == nil {
        log.Printf("Cache hit for area code: %s. Most Played Mode: %s, Player Count: %d", areaCode, mode, count)
        return mode, count, nil
    }
    log.Printf("Cache miss for area code: %s. Fetching from database...", areaCode)

    mode, dbCount, err := db.GetPopularModeByArea(areaCode)
    if err != nil {
        log.Printf("Error fetching from database for area code: %s: %v", areaCode, err)
        return "", 0, err
    }

    cacheErr := cache.SetPopularModeCache(areaCode, mode, int32(dbCount), 5*time.Minute)
    if cacheErr != nil {
        log.Printf("Error setting cache for area code: %s: %v. However, returning data from DB.", areaCode, cacheErr)
    } else {
        log.Printf("Cache set for area code: %s, Most Played Mode: %s, Player Count: %d for 5 minutes", areaCode, mode, dbCount)
    }

    return mode, int32(dbCount), nil
}
