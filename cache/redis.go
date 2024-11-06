package cache

import (
    "context"
    "encoding/json"
    "github.com/go-redis/redis/v8"
    "log"
    "time"
)

var ctx = context.Background()
var rdb *redis.Client


type ModeCache struct {
    Mode  string `json:"mode"`
    Count int32  `json:"count"`
}

func InitRedis(addr string) {
    rdb = redis.NewClient(&redis.Options{
        Addr: addr,
    })
    log.Println("Connected to Redis")
}

func GetPopularModeCache(areaCode string) (string, int32, error) {
    val, err := rdb.Get(ctx, areaCode).Result()
    if err != nil {
        return "", 0, err 
    }

    var modeCache ModeCache
    err = json.Unmarshal([]byte(val), &modeCache)
    if err != nil {
        return "", 0, err
    }

    return modeCache.Mode, modeCache.Count, nil
}

func SetPopularModeCache(areaCode, mode string, count int32, duration time.Duration) error {
    modeCache := ModeCache{
        Mode:  mode,
        Count: count,
    }

    data, err := json.Marshal(modeCache)
    if err != nil {
        return err
    }

    err = rdb.Set(ctx, areaCode, data, duration).Err()
    if err != nil {
        return err 
    }

    return nil
}
