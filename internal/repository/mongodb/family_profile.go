package mongodb

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type FamilyProfileRepository struct {
    client *mongo.Client
}

func NewFamilyProfileRepository(client *mongo.Client) *FamilyProfileRepository {
    return &FamilyProfileRepository{client: client}
}

func (r *FamilyProfileRepository) InsertFamilyProfile(profile interface{}) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := r.client.Database("mydatabase").Collection("family_profiles")
    _, err := collection.InsertOne(ctx, profile)
    if err != nil {
        log.Printf("Failed to insert family profile: %v", err)
        return err
    }
    return nil
}