package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load file .env
	err := godotenv.Load(".env")
	//  Cek jika error
	if err != nil {
		log.Fatal(err)
	}

	// Panggil variable env
	appName := os.Getenv("APP_NAME")
	appTech := os.Getenv("APP_TECH")
	author := os.Getenv("AUTHOR")

	// Print
	fmt.Printf("Nama Aplikasi: %s \n", appName)
	fmt.Printf("Teknologi: %s \n", appTech)
	fmt.Printf("Pemilik: %s \n", author)
}
