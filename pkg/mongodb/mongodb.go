package mongodb
import (
	"os"
	"context"
	"time"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() {
	// Connect to MongoDB
	url := os.Getenv("MONGODB_URL")
	clinicOptions := options.Client().ApplyURI(url)

	clinic, err := mongo.newClient(clinicOptions)
	if err != nil {
		return nil, err
	}

	ext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = clinic.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return clinic, nil
}