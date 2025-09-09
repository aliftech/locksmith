package util

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/sha3"
)

// Wallet represents a cryptocurrency wallet with mnemonic and derived keys
type CryptoWallet struct {
	Mnemonic   string
	Passphrase string
	keys       map[string]*bip32.Key
	mux        sync.Mutex
}

// CryptoAddress represents different cryptocurrency address types
type CryptoAddress struct {
	PrivateKeyHex string
	PublicKeyHex  string
	Address       string
}

// NewWallet creates a new cryptocurrency wallet with a mnemonic
func NewCryptoWallet(passphrase string) (*CryptoWallet, error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return nil, fmt.Errorf("failed to generate entropy: %v", err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, fmt.Errorf("failed to generate mnemonic: %v", err)
	}

	return &CryptoWallet{
		Mnemonic:   mnemonic,
		Passphrase: passphrase,
		keys:       make(map[string]*bip32.Key),
	}, nil
}

// getKey retrieves or sets a key in the wallet's key map
func (w *CryptoWallet) getCryptoKey(path string) (*bip32.Key, bool) {
	w.mux.Lock()
	defer w.mux.Unlock()
	key, ok := w.keys[path]
	return key, ok
}

func (w *CryptoWallet) setCryptoKey(path string, key *bip32.Key) {
	w.mux.Lock()
	defer w.mux.Unlock()
	w.keys[path] = key
}

// GetSeed generates the seed from the mnemonic and passphrase
func (w *CryptoWallet) GetCryptoSeed() []byte {
	return bip39.NewSeed(w.Mnemonic, w.Passphrase)
}

// GetMasterKey derives the master key from the seed
func (w *CryptoWallet) GetCryptoMasterKey() (*bip32.Key, error) {
	path := "m"
	if key, ok := w.getCryptoKey(path); ok {
		return key, nil
	}

	key, err := bip32.NewMasterKey(w.GetCryptoSeed())
	if err != nil {
		return nil, fmt.Errorf("failed to generate master key: %v", err)
	}

	w.setCryptoKey(path, key)
	return key, nil
}

// GetEthereumKey derives a key for Ethereum using BIP-44 path: m/44'/60'/0'/0/index
func (w *CryptoWallet) GetEthereumKey(index uint32) (*bip32.Key, error) {
	path := fmt.Sprintf("m/44'/60'/0'/0/%d", index)

	if key, ok := w.getCryptoKey(path); ok {
		return key, nil
	}

	masterKey, err := w.GetCryptoMasterKey()
	if err != nil {
		return nil, err
	}

	// m/44'
	purposeKey, err := masterKey.NewChildKey(0x8000002C) // 44' (hardened)
	if err != nil {
		return nil, fmt.Errorf("failed to derive purpose key: %v", err)
	}

	// m/44'/60'
	coinTypeKey, err := purposeKey.NewChildKey(0x8000003C) // 60' (ETH, hardened)
	if err != nil {
		return nil, fmt.Errorf("failed to derive coin type key: %v", err)
	}

	// m/44'/60'/0'
	accountKey, err := coinTypeKey.NewChildKey(0x80000000) // 0' (hardened)
	if err != nil {
		return nil, fmt.Errorf("failed to derive account key: %v", err)
	}

	// m/44'/60'/0'/0
	changeKey, err := accountKey.NewChildKey(0) // External chain
	if err != nil {
		return nil, fmt.Errorf("failed to derive change key: %v", err)
	}

	// m/44'/60'/0'/0/index
	key, err := changeKey.NewChildKey(index)
	if err != nil {
		return nil, fmt.Errorf("failed to derive index key: %v", err)
	}

	w.setCryptoKey(path, key)
	return key, nil
}

// GetSolanaKey derives a key for Solana using BIP-44 path: m/44'/501'/0'/0/index
func (w *CryptoWallet) GetSolanaKey(index uint32) (*bip32.Key, error) {
	path := fmt.Sprintf("m/44'/501'/0'/0/%d", index)

	if key, ok := w.getCryptoKey(path); ok {
		return key, nil
	}

	masterKey, err := w.GetCryptoMasterKey()
	if err != nil {
		return nil, err
	}

	// m/44'
	purposeKey, err := masterKey.NewChildKey(0x8000002C) // 44' (hardened)
	if err != nil {
		return nil, fmt.Errorf("failed to derive purpose key: %v", err)
	}

	// m/44'/501'
	coinTypeKey, err := purposeKey.NewChildKey(0x800001F5) // 501' (SOL, hardened)
	if err != nil {
		return nil, fmt.Errorf("failed to derive coin type key: %v", err)
	}

	// m/44'/501'/0'
	accountKey, err := coinTypeKey.NewChildKey(0x80000000) // 0' (hardened)
	if err != nil {
		return nil, fmt.Errorf("failed to derive account key: %v", err)
	}

	// m/44'/501'/0'/0
	changeKey, err := accountKey.NewChildKey(0) // External chain
	if err != nil {
		return nil, fmt.Errorf("failed to derive change key: %v", err)
	}

	// m/44'/501'/0'/0/index
	key, err := changeKey.NewChildKey(index)
	if err != nil {
		return nil, fmt.Errorf("failed to derive index key: %v", err)
	}

	w.setCryptoKey(path, key)
	return key, nil
}

// GetCardanoKey derives a key for Cardano using BIP-44 path: m/44'/1815'/0'/0/index
func (w *CryptoWallet) GetCardanoKey(index uint32) (*bip32.Key, error) {
	path := fmt.Sprintf("m/44'/1815'/0'/0/%d", index)

	if key, ok := w.getCryptoKey(path); ok {
		return key, nil
	}

	masterKey, err := w.GetCryptoMasterKey()
	if err != nil {
		return nil, err
	}

	// m/44'
	purposeKey, err := masterKey.NewChildKey(0x8000002C) // 44' (hardened)
	if err != nil {
		return nil, fmt.Errorf("failed to derive purpose key: %v", err)
	}

	// m/44'/1815'
	coinTypeKey, err := purposeKey.NewChildKey(0x80000717) // 1815' (ADA, hardened)
	if err != nil {
		return nil, fmt.Errorf("failed to derive coin type key: %v", err)
	}

	// m/44'/1815'/0'
	accountKey, err := coinTypeKey.NewChildKey(0x80000000) // 0' (hardened)
	if err != nil {
		return nil, fmt.Errorf("failed to derive account key: %v", err)
	}

	// m/44'/1815'/0'/0
	changeKey, err := accountKey.NewChildKey(0) // External chain
	if err != nil {
		return nil, fmt.Errorf("failed to derive change key: %v", err)
	}

	// m/44'/1815'/0'/0/index
	key, err := changeKey.NewChildKey(index)
	if err != nil {
		return nil, fmt.Errorf("failed to derive index key: %v", err)
	}

	w.setCryptoKey(path, key)
	return key, nil
}

// GenerateEthereumAddress generates an Ethereum key pair and address
func (w *CryptoWallet) GenerateEthereumAddress(index uint32) (*CryptoAddress, error) {
	key, err := w.GetEthereumKey(index)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.ToECDSA(key.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to ECDSA: %v", err)
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	publicKeyBytes := crypto.FromECDSAPub(publicKey)
	address := crypto.PubkeyToAddress(*publicKey).Hex()

	return &CryptoAddress{
		PrivateKeyHex: hex.EncodeToString(key.Key),
		PublicKeyHex:  hex.EncodeToString(publicKeyBytes),
		Address:       address,
	}, nil
}

// GenerateCardanoAddress generates a Cardano key pair and address (simplified base address)
func (w *CryptoWallet) GenerateCardanoAddress(index uint32) (*CryptoAddress, error) {
	key, err := w.GetCardanoKey(index)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.ToECDSA(key.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to ECDSA: %v", err)
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	publicKeyBytes := crypto.FromECDSAPub(publicKey)

	// Simplified Cardano address generation (base address)
	hash28 := func(data []byte) []byte {
		hasher := sha3.New256()
		hasher.Write(data)
		return hasher.Sum(nil)[:28]
	}

	paymentKeyHash := hash28(publicKeyBytes)
	stakeKeyHash := hash28(publicKeyBytes) // Simplified, typically separate key
	header := byte(0x01)                   // Base address header
	addrBytes := append([]byte{header}, append(paymentKeyHash, stakeKeyHash...)...)

	// Cardano address is base58 encoded with a CRC32 checksum
	address := base58.Encode(addrBytes)

	return &CryptoAddress{
		PrivateKeyHex: hex.EncodeToString(key.Key),
		PublicKeyHex:  hex.EncodeToString(publicKeyBytes),
		Address:       address,
	}, nil
}

func GenerateRandomCryptoAddress(cryptoType string) (*CryptoAddress, error) {
	privateKeyBytes := make([]byte, 32)
	_, err := rand.Read(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	switch strings.ToLower(cryptoType) {
	case "ethereum", "avax", "bsc", "polygon":
		privateKey, err := crypto.ToECDSA(privateKeyBytes)
		if err != nil {
			return nil, fmt.Errorf("failed to convert to ECDSA: %v", err)
		}
		publicKey := privateKey.Public().(*ecdsa.PublicKey)
		publicKeyBytes := crypto.FromECDSAPub(publicKey)
		address := crypto.PubkeyToAddress(*publicKey).Hex()

		return &CryptoAddress{
			PrivateKeyHex: hex.EncodeToString(privateKeyBytes),
			PublicKeyHex:  hex.EncodeToString(publicKeyBytes),
			Address:       address,
		}, nil

	case "bitcoin", "bitcoincash", "litecoin", "dogecoin", "tron", "ripple":
		privateKey, err := crypto.ToECDSA(privateKeyBytes)
		if err != nil {
			return nil, fmt.Errorf("failed to convert to ECDSA: %v", err)
		}
		publicKey := privateKey.Public().(*ecdsa.PublicKey)
		publicKeyBytes := crypto.FromECDSAPub(publicKey)
		hash160 := btcutil.Hash160(publicKeyBytes)

		var version byte
		switch cryptoType {
		case "bitcoin":
			version = 0x00
		case "bitcoincash":
			version = 0x00
		case "litecoin":
			version = 0x30
		case "dogecoin":
			version = 0x1E
		case "tron":
			version = 0x41
		case "ripple":
			version = 0x00
		}

		versionedPayload := append([]byte{version}, hash160...)
		firstHash := sha256.Sum256(versionedPayload)
		secondHash := sha256.Sum256(firstHash[:]) // ✅ Convert [32]byte → []byte using [:]
		checksum := secondHash[:4]
		addressBytes := append(versionedPayload, checksum...)
		address := base58.Encode(addressBytes)

		if cryptoType == "ripple" {
			address = "r" + address[1:]
		}

		return &CryptoAddress{
			PrivateKeyHex: hex.EncodeToString(privateKeyBytes),
			PublicKeyHex:  hex.EncodeToString(publicKeyBytes),
			Address:       address,
		}, nil

	case "polkadot":
		privateKey, err := crypto.ToECDSA(privateKeyBytes)
		if err != nil {
			return nil, fmt.Errorf("failed to convert to ECDSA: %v", err)
		}
		publicKey := privateKey.Public().(*ecdsa.PublicKey)
		publicKeyBytes := crypto.FromECDSAPub(publicKey)

		var pubKey32 [32]byte
		copy(pubKey32[:], publicKeyBytes)
		payload := append([]byte{42}, pubKey32[:]...)
		hash := sha256.Sum256(payload)
		checksum := hash[:2]
		payload = append(payload, checksum...)
		address := base58.Encode(payload)

		return &CryptoAddress{
			PrivateKeyHex: hex.EncodeToString(privateKeyBytes),
			PublicKeyHex:  hex.EncodeToString(publicKeyBytes),
			Address:       address,
		}, nil

	case "cosmos":
		privateKey, err := crypto.ToECDSA(privateKeyBytes)
		if err != nil {
			return nil, fmt.Errorf("failed to convert to ECDSA: %v", err)
		}
		publicKey := privateKey.Public().(*ecdsa.PublicKey)
		publicKeyBytes := crypto.FromECDSAPub(publicKey)
		hash160 := btcutil.Hash160(publicKeyBytes)
		address := bech32Encode("cosmos", hash160)

		return &CryptoAddress{
			PrivateKeyHex: hex.EncodeToString(privateKeyBytes),
			PublicKeyHex:  hex.EncodeToString(publicKeyBytes),
			Address:       address,
		}, nil

	case "cardano":
		// Existing code
		privateKey, err := crypto.ToECDSA(privateKeyBytes)
		if err != nil {
			return nil, fmt.Errorf("failed to convert to ECDSA: %v", err)
		}
		publicKey := privateKey.Public().(*ecdsa.PublicKey)
		publicKeyBytes := crypto.FromECDSAPub(publicKey)

		hash28 := func(data []byte) []byte {
			hasher := sha3.New256()
			hasher.Write(data)
			return hasher.Sum(nil)[:28]
		}

		paymentKeyHash := hash28(publicKeyBytes)
		stakeKeyHash := hash28(publicKeyBytes)
		header := byte(0x01)
		addrBytes := append([]byte{header}, append(paymentKeyHash, stakeKeyHash...)...)
		address := base58.Encode(addrBytes)

		return &CryptoAddress{
			PrivateKeyHex: hex.EncodeToString(privateKeyBytes),
			PublicKeyHex:  hex.EncodeToString(publicKeyBytes),
			Address:       address,
		}, nil

	default:
		return nil, fmt.Errorf("unsupported cryptocurrency: %s", cryptoType)
	}
}
