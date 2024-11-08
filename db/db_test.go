package db

import (
    "context"
    "testing"
    "time"
    "multiplayer-mode-usage/db"
    "go.mongodb.org/mongo-driver/bson"
    "github.com/stretchr/testify/assert"
)

// MongoDB Test Configurations
const testMongoURI = "mongodb://localhost:27017"
const testDBName = "test_db"

func setupTestDB() {
    db.InitializeMongoDB(testMongoURI, testDBName)
}

func TestGetPopularModeByArea_Success(t *testing.T) {
    setupTestDB()
    testDB := db.GetDB()
    collection := testDB.Collection("mode_usage")

    // Insert test data
    collection.InsertOne(context.TODO(), bson.M{
        "area_code":    "NY",
        "mode_name":    "Battle Royale",
        "player_count": 150,
    })

    mode, count, err := db.GetPopularModeByArea("NY")
    assert.NoError(t, err)
    assert.Equal(t, "Battle Royale", mode)
    assert.Equal(t, 150, count)

    collection.Drop(context.TODO())
}

func TestGetPopularModeByArea_NoData(t *testing.T) {
    setupTestDB()
    mode, count, err := db.GetPopularModeByArea("Unknown")
    assert.Error(t, err)
    assert.Equal(t, "", mode)
    assert.Equal(t, 0, count)
}
