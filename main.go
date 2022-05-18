package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/alf248/gotrade/database"
	"github.com/alf248/gotrade/database/mock"
	m "github.com/alf248/gotrade/database/mock"
	"github.com/alf248/gotrade/server"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// PORT is the port that this http server will use.
	// FRONTEND is the URL to the frontend app. Needed for CORS to function.
	// MONGO is the URL to the Mongo database.
	var PORT, FRONTEND, MONGO string

	// If a 'env' file is found, then use data from that file.
	// Otherwise get data from Environment Variables.
	env := getEnvDataFromFile("env")
	if env != nil {
		PORT = env.PORT
		FRONTEND = env.FRONTEND
		MONGO = env.MONGO
		if env.UseAtlasSearch {
			database.UseAtlasSearch = true
		}
	} else {
		PORT, FRONTEND, MONGO = getEnvironmentVariables()
		database.UseAtlasSearch = true
	}

	// CONNECT TO MONGO DATABASE
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	database.Client, err = mongo.Connect(ctx, options.Client().ApplyURI(MONGO))
	if err != nil {
		panic(err)
	}
	defer database.Client.Disconnect(ctx)

	// POSSIBLY ADD MOCK DATA
	mockData := getMockDataFromFile("mockdata")
	if mockData != nil {
		if mockData.AddData {
			log.Println("ADDING MOCK DATA")
			mock.UserNames = mockData.UserNames
			m.AddMockUsers()
			m.AddMockOffers(mockData.RandomOffers)
		}
	}

	// START HTTP SERVER
	server.Start(PORT, FRONTEND)
}

func getEnvironmentVariables() (string, string, string) {

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	frontEnd := os.Getenv("FRONTEND")
	if frontEnd == "" {
		log.Fatal("$FRONTEND must be set")
	}

	mongoURL := os.Getenv("MONGO")
	if mongoURL == "" {
		log.Fatal("$MONGO must be set")
	}

	return port, frontEnd, mongoURL
}

type envData struct {
	MONGO          string `json:"MONGO"`
	FRONTEND       string `json:"FRONTEND"`
	PORT           string `json:"PORT"`
	UseAtlasSearch bool   `json:"UseAtlasSearch"`
}

func getEnvDataFromFile(filename string) *envData {
	file, err := ioutil.ReadFile(filename)
	if err == nil {
		data := envData{}
		err := json.Unmarshal([]byte(file), &data)
		if err != nil {
			panic(err)
		}

		log.Println("Found env file:")
		//log.Println("PORT", devMode.PORT)
		log.Println("MONGO", data.MONGO)
		//log.Println("FRONTEND", devMode.FRONTEND)
		log.Println("UseAtlasSearch", data.UseAtlasSearch)

		return &data
	}
	return nil
}

type mockData struct {
	AddData      bool     `json:"addData"`
	UserNames    []string `json:"userNames"`
	RandomOffers int      `json:"randomOffers"`
}

func getMockDataFromFile(filename string) *mockData {
	file, err := ioutil.ReadFile(filename)
	if err == nil {
		data := mockData{}
		err := json.Unmarshal([]byte(file), &data)
		if err != nil {
			panic(err)
		}

		return &data
	}
	return nil
}
