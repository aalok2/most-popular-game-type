package db

import (
    "context"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
    "time"
)

var db *mongo.Database

type ModeUsage struct {
    AreaCode    string `bson:"area_code"`
    ModeName    string `bson:"mode_name"`
    PlayerCount int    `bson:"player_count"`
}


func InitMongoDB(uri string, dbName string) {
    clientOptions := options.Client().ApplyURI(uri)
    
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB at %s: %v\n", uri, err)
    }
    

    db = client.Database(dbName)
    
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatalf("Failed to ping MongoDB: %v\n", err)
    }
    
    log.Printf("Successfully connected to MongoDB at %s, using database: %s\n", uri, dbName)
}

func GetDB() *mongo.Database {
    return db
}


func GetPopularModeByArea(areaCode string) (string, int, error) {
    collection := db.Collection("mode_usage")

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

