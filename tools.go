package main

import "github.com/joho/godotenv"

// Load current .env file into map
func loadCurEnv() (EnvDataMap, error) {
	var env EnvDataMap
	env, err := godotenv.Read()
	if err != nil {
		return nil, err
	}
	return env, nil
}

// Writes updated env data into .env file
func writeEnv(env EnvDataMap) error {
	return godotenv.Write(env, ".env")
}