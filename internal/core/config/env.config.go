package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func EnvSetup() {
	// Get the directory of the current file (env_setup.go)
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	// Go up 2 levels: config/ → project-root/
	// Adjust levels based on your structure
	rootPath := filepath.Join(basepath, "../../../")
	envPath := filepath.Join(rootPath, ".env")

	// Check if file exists
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		log.Fatalf("❌ .env file not found at: %s", envPath)
	}

	// Load the .env file
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("❌ Error loading .env file:", err)
	}

	log.Println("✅ .env loaded successfully from:", envPath)
}
