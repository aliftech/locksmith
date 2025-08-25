package main

import (
	"fmt"
	"log"

	"github.com/aliftech/locksmith/src/app/migrations"
	"github.com/aliftech/locksmith/src/config"
	"github.com/aliftech/locksmith/src/libraries"
)

func init() {
	config.Connect()
	migrations.MainMigration()
}

func main() {
	PUBLIC_KEY, PRIVATE_KEY, err := libraries.GenerateBTCeKey()
	if err != nil {
		log.Fatal("ERROR! Invalid private key")
	}

	fmt.Printf("Private key: %s \n", PRIVATE_KEY)
	fmt.Printf("Public key: %s", PUBLIC_KEY)
}
