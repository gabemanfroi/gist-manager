/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/gabemanfroi/gistManager/cmd"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	fmt.Println("Initializing...")
	if os.Getenv("APP_ENV") == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func main() {
	cmd.Execute()
}
