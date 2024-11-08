package db

import (
    "context"
    "errors"
    "log"
    "sync"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
    dbInstance    *mongo.Database
    clientInstance *mongo.Client
    clientInstanceErr error
    mongoOnce         sync.Once
)

type ModeUsage struct {
    AreaCode    string `bson:"area_code"`
    ModeName    string `bson:"mode_name"`
    PlayerCount int    `bson:"player_count"`
}

func InitializeMongoDB(uri, dbName string) *mongo.Database {
    mongoOnce.Do(func() {
        log.Printf("Initializing MongoDB client...")

        clientOptions := options.Client().ApplyURI(uri)
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        clientInstance, clientInstanceErr = mongo.Connect(ctx, clientOptions)
        if clientInstanceErr != nil {
            log.Fatalf("Failed to connect to MongoDB at %s: %v\n", uri, clientInstanceErr)
        }

        dbInstance = clientInstance.Database(dbName)

        err := clientInstance.Ping(ctx, nil)
        if err != nil {
            log.Fatalf("Failed to ping MongoDB: %v\n", err)
        }

        log.Printf("Successfully connected to MongoDB at %s, using database: %s\n", uri, dbName)
    })

    return dbInstance
}

func GetDB() *mongo.Database {
    return dbInstance
}

func GetPopularModeByArea(areaCode string) (string, int, error) {
    collection := GetDB().Collection("mode_usage")

    filter := bson.M{"area_code": areaCode}
    opts := options.FindOne().SetSort(bson.D{{"player_count", -1}})

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var result ModeUsage
    err := collection.FindOne(ctx, filter, opts).Decode(&result)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return "", 0, errors.New("no data found for the specified area code")
        }
        log.Println("Error fetching popular mode:", err)
        return "", 0, err
    }

    return result.ModeName, result.PlayerCount, nil
}
