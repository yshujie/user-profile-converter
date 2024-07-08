package main

import (
	"log"
	"os"
	"user-profile-converter/innternal/repository/mongodb"
	"user-profile-converter/innternal/repository/mysql"
	"user-profile-converter/internal/service"
	"user-profile-converter/pkg/mysql"
	"user-profile-converter/pkg/mongodb"

	"github.com/joho/godotenv"
	"go.mongidb.org/mongo-driver/bson"
	"context"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysqlDB,err := mysql.Connect()
	if err != nil {
		log.Fatal("Faild to connect to MySQL, err: ", err)
	}
	defer mysqlDB.Close()

	mongoClient, err := mongodb.Connect()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB, err: ", err)
	}
	defer mongoClient.Disconnect(context.Background())

	log.Println("Connected to MySQL and MongoDB")
}
