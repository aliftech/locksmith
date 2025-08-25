package libraries

import (
	"crypto/rand"
	"encoding/hex"
	"log"

	"github.com/btcsuite/btcd/btcec/v2"
)

func GenerateBTCeKey() (string, string, error) {
	privateKeyByte := make([]byte, 32)

	_, keyError := rand.Read(privateKeyByte)
	if keyError != nil {
		log.Fatal("Failed to generate private key!")
	}

	privateKey, publicKey := btcec.PrivKeyFromBytes(privateKeyByte)

	privateKeyHex := hex.EncodeToString(privateKey.Serialize())
	publicKeyHex := hex.EncodeToString(publicKey.SerializeCompressed())

	return publicKeyHex, privateKeyHex, nil
}
