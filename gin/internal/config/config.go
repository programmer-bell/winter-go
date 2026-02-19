package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)


type Config struct{
	MongoURI string
	MongoDB string
	ServerPort string
}


func Load() (Config,error){

	if err := godotenv.Load(); err != nil{
		return  Config{}, fmt.Errorf("Failed to load .env")
	}

	mongoURI, err := extractEnv("MONGODB_URI")
	if err != nil{
		return Config{}, err
	}
	mongoDB, err := extractEnv("MONGODB_DB")
	if err != nil{
		return Config{}, err
	}
	port, err := extractEnv("PORT")
	if err != nil{
		return Config{}, err
	}

	return Config{
		MongoURI: mongoURI,
		MongoDB: mongoDB,
		ServerPort: port,
	}, nil

}

func extractEnv(key string) (string, error){

	val := os.Getenv(key)

	if val == ""{
		return "",fmt.Errorf("Missing request .env")
	}

	return val, nil

}

