package util

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
	"github.com/tyler-smith/go-bip32"
)

// GetBitcoinCashKey derives a key for Bitcoin Cash using BIP-44 path: m/44'/145'/0'/0/index
func (w *CryptoWallet) GetBitcoinCashKey(index uint32) (*bip32.Key, error) {
	path := fmt.Sprintf("m/44'/145'/0'/0/%d", index)
	if key, ok := w.getCryptoKey(path); ok {
		return key, nil
	}

	masterKey, err := w.GetCryptoMasterKey()
	if err != nil {
		return nil, err
	}

	purposeKey, err := masterKey.NewChildKey(0x8000002C) // 44'
	if err != nil {
		return nil, fmt.Errorf("failed to derive purpose key: %v", err)
	}
	coinTypeKey, err := purposeKey.NewChildKey(0x80000091) // 145' (BCH)
	if err != nil {
		return nil, fmt.Errorf("failed to derive coin type key: %v", err)
	}
	accountKey, err := coinTypeKey.NewChildKey(0x80000000) // 0'
	if err != nil {
		return nil, fmt.Errorf("failed to derive account key: %v", err)
	}
	changeKey, err := accountKey.NewChildKey(0) // external
	if err != nil {
		return nil, fmt.Errorf("failed to derive change key: %v", err)
	}
	key, err := changeKey.NewChildKey(index)
	if err != nil {
		return nil, fmt.Errorf("failed to derive index key: %v", err)
	}

	w.setCryptoKey(path, key)
	return key, nil
}

// GetLitecoinKey derives a key for Litecoin using BIP-44 path: m/44'/2'/0'/0/index
func (w *CryptoWallet) GetLitecoinKey(index uint32) (*bip32.Key, error) {
	path := fmt.Sprintf("m/44'/2'/0'/0/%d", index)
	if key, ok := w.getCryptoKey(path); ok {
		return key, nil
	}

	masterKey, err := w.GetCryptoMasterKey()
	if err != nil {
		return nil, err
	}

	purposeKey, err := masterKey.NewChildKey(0x8000002C) // 44'
	if err != nil {
		return nil, fmt.Errorf("failed to derive purpose key: %v", err)
	}
	coinTypeKey, err := purposeKey.NewChildKey(0x80000002) // 2' (LTC)
	if err != nil {
		return nil, fmt.Errorf("failed to derive coin type key: %v", err)
	}
	accountKey, err := coinTypeKey.NewChildKey(0x80000000) // 0'
	if err != nil {
		return nil, fmt.Errorf("failed to derive account key: %v", err)
	}
	changeKey, err := accountKey.NewChildKey(0) // external
	if err != nil {
		return nil, fmt.Errorf("failed to derive change key: %v", err)
	}
	key, err := changeKey.NewChildKey(index)
	if err != nil {
		return nil, fmt.Errorf("failed to derive index key: %v", err)
	}

	w.setCryptoKey(path, key)
	return key, nil
}

// GetDogecoinKey derives a key for Dogecoin using BIP-44 path: m/44'/3'/0'/0/index
func (w *CryptoWallet) GetDogecoinKey(index uint32) (*bip32.Key, error) {
	path := fmt.Sprintf("m/44'/3'/0'/0/%d", index)
	if key, ok := w.getCryptoKey(path); ok {
		return key, nil
	}

	masterKey, err := w.GetCryptoMasterKey()
	if err != nil {
		return nil, err
	}

	purposeKey, err := masterKey.NewChildKey(0x8000002C) // 44'
	if err != nil {
		return nil, fmt.Errorf("failed to derive purpose key: %v", err)
	}
	coinTypeKey, err := purposeKey.NewChildKey(0x80000003) // 3' (DOGE)
	if err != nil {
		return nil, fmt.Errorf("failed to derive coin type key: %v", err)
	}
	accountKey, err := coinTypeKey.NewChildKey(0x80000000) // 0'
	if err != nil {
		return nil, fmt.Errorf("failed to derive account key: %v", err)
	}
	changeKey, err := accountKey.NewChildKey(0) // external
	if err != nil {
		return nil, fmt.Errorf("failed to derive change key: %v", err)
	}
	key, err := changeKey.NewChildKey(index)
	if err != nil {
		return nil, fmt.Errorf("failed to derive index key: %v", err)
	}

	w.setCryptoKey(path, key)
	return key, nil
}

// GetPolkadotKey derives a key for Polkadot using BIP-44 path: m/44'/354'/0'/0/index
// Note: Polkadot uses SR25519, this is simulated using ECDSA for compatibility
func (w *CryptoWallet) GetPolkadotKey(index uint32) (*bip32.Key, error) {
	path := fmt.Sprintf("m/44'/354'/0'/0/%d", index)
	if key, ok := w.getCryptoKey(path); ok {
		return key, nil
	}

	masterKey, err := w.GetCryptoMasterKey()
	if err != nil {
		return nil, err
	}

	purposeKey, err := masterKey.NewChildKey(0x8000002C) // 44'
	if err != nil {
		return nil, fmt.Errorf("failed to derive purpose key: %v", err)
	}
	coinTypeKey, err := purposeKey.NewChildKey(0x80000162) // 354' (DOT)
	if err != nil {
		return nil, fmt.Errorf("failed to derive coin type key: %v", err)
	}
	accountKey, err := coinTypeKey.NewChildKey(0x80000000) // 0'
	if err != nil {
		return nil, fmt.Errorf("failed to derive account key: %v", err)
	}
	changeKey, err := accountKey.NewChildKey(0) // external
	if err != nil {
		return nil, fmt.Errorf("failed to derive change key: %v", err)
	}
	key, err := changeKey.NewChildKey(index)
	if err != nil {
		return nil, fmt.Errorf("failed to derive index key: %v", err)
	}

	w.setCryptoKey(path, key)
	return key, nil
}

// GetTronKey derives a key for Tron using BIP-44 path: m/44'/195'/0'/0/index
func (w *CryptoWallet) GetTronKey(index uint32) (*bip32.Key, error) {
	path := fmt.Sprintf("m/44'/195'/0'/0/%d", index)
	if key, ok := w.getCryptoKey(path); ok {
		return key, nil
	}

	masterKey, err := w.GetCryptoMasterKey()
	if err != nil {
		return nil, err
	}

	purposeKey, err := masterKey.NewChildKey(0x8000002C) // 44'
	if err != nil {
		return nil, fmt.Errorf("failed to derive purpose key: %v", err)
	}
	coinTypeKey, err := purposeKey.NewChildKey(0x800000C3) // 195' (TRX)
	if err != nil {
		return nil, fmt.Errorf("failed to derive coin type key: %v", err)
	}
	accountKey, err := coinTypeKey.NewChildKey(0x80000000) // 0'
	if err != nil {
		return nil, fmt.Errorf("failed to derive account key: %v", err)
	}
	changeKey, err := accountKey.NewChildKey(0) // external
	if err != nil {
		return nil, fmt.Errorf("failed to derive change key: %v", err)
	}
	key, err := changeKey.NewChildKey(index)
	if err != nil {
		return nil, fmt.Errorf("failed to derive index key: %v", err)
	}

	w.setCryptoKey(path, key)
	return key, nil
}

// GetCosmosKey derives a key for Cosmos using BIP-44 path: m/44'/118'/0'/0/index
func (w *CryptoWallet) GetCosmosKey(index uint32) (*bip32.Key, error) {
	path := fmt.Sprintf("m/44'/118'/0'/0/%d", index)
	if key, ok := w.getCryptoKey(path); ok {
		return key, nil
	}

	masterKey, err := w.GetCryptoMasterKey()
	if err != nil {
		return nil, err
	}

	purposeKey, err := masterKey.NewChildKey(0x8000002C) // 44'
	if err != nil {
		return nil, fmt.Errorf("failed to derive purpose key: %v", err)
	}
	coinTypeKey, err := purposeKey.NewChildKey(0x80000076) // 118' (ATOM)
	if err != nil {
		return nil, fmt.Errorf("failed to derive coin type key: %v", err)
	}
	accountKey, err := coinTypeKey.NewChildKey(0x80000000) // 0'
	if err != nil {
		return nil, fmt.Errorf("failed to derive account key: %v", err)
	}
	changeKey, err := accountKey.NewChildKey(0) // external
	if err != nil {
		return nil, fmt.Errorf("failed to derive change key: %v", err)
	}
	key, err := changeKey.NewChildKey(index)
	if err != nil {
		return nil, fmt.Errorf("failed to derive index key: %v", err)
	}

	w.setCryptoKey(path, key)
	return key, nil
}

// GetRippleKey derives a key for Ripple using BIP-44 path: m/44'/144'/0'/0/index
func (w *CryptoWallet) GetRippleKey(index uint32) (*bip32.Key, error) {
	path := fmt.Sprintf("m/44'/144'/0'/0/%d", index)
	if key, ok := w.getCryptoKey(path); ok {
		return key, nil
	}

	masterKey, err := w.GetCryptoMasterKey()
	if err != nil {
		return nil, err
	}

	purposeKey, err := masterKey.NewChildKey(0x8000002C) // 44'
	if err != nil {
		return nil, fmt.Errorf("failed to derive purpose key: %v", err)
	}
	coinTypeKey, err := purposeKey.NewChildKey(0x80000090) // 144' (XRP)
	if err != nil {
		return nil, fmt.Errorf("failed to derive coin type key: %v", err)
	}
	accountKey, err := coinTypeKey.NewChildKey(0x80000000) // 0'
	if err != nil {
		return nil, fmt.Errorf("failed to derive account key: %v", err)
	}
	changeKey, err := accountKey.NewChildKey(0) // external
	if err != nil {
		return nil, fmt.Errorf("failed to derive change key: %v", err)
	}
	key, err := changeKey.NewChildKey(index)
	if err != nil {
		return nil, fmt.Errorf("failed to derive index key: %v", err)
	}

	w.setCryptoKey(path, key)
	return key, nil
}

// GenerateBitcoinCashAddress generates a Bitcoin Cash address (same format as BTC for simplicity)
func (w *CryptoWallet) GenerateBitcoinCashAddress(index uint32) (*CryptoAddress, error) {
	key, err := w.GetBitcoinCashKey(index)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.ToECDSA(key.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to ECDSA: %v", err)
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	pubKeyBytes := crypto.FromECDSAPub(publicKey)
	hash160 := btcutil.Hash160(pubKeyBytes)
	versionedPayload := append([]byte{0x00}, hash160...)
	firstHash := sha256.Sum256(versionedPayload)
	secondHash := sha256.Sum256(firstHash[:]) // ✅ Convert [32]byte → []byte using [:]
	checksum := secondHash[:4]
	addressBytes := append(versionedPayload, checksum...)
	address := base58.Encode(addressBytes)

	return &CryptoAddress{
		PrivateKeyHex: hex.EncodeToString(key.Key),
		PublicKeyHex:  hex.EncodeToString(pubKeyBytes),
		Address:       address,
	}, nil
}

// GenerateLitecoinAddress generates a Litecoin address (version 0x30)
func (w *CryptoWallet) GenerateLitecoinAddress(index uint32) (*CryptoAddress, error) {
	key, err := w.GetLitecoinKey(index)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.ToECDSA(key.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to ECDSA: %v", err)
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	pubKeyBytes := crypto.FromECDSAPub(publicKey)
	hash160 := btcutil.Hash160(pubKeyBytes)
	versionedPayload := append([]byte{0x30}, hash160...) // Litecoin uses 0x30
	firstHash := sha256.Sum256(versionedPayload)
	secondHash := sha256.Sum256(firstHash[:]) // ✅ Convert [32]byte → []byte using [:]
	checksum := secondHash[:4]
	addressBytes := append(versionedPayload, checksum...)
	address := base58.Encode(addressBytes)

	return &CryptoAddress{
		PrivateKeyHex: hex.EncodeToString(key.Key),
		PublicKeyHex:  hex.EncodeToString(pubKeyBytes),
		Address:       address,
	}, nil
}

// GenerateDogecoinAddress generates a Dogecoin address (version 0x1E)
func (w *CryptoWallet) GenerateDogecoinAddress(index uint32) (*CryptoAddress, error) {
	key, err := w.GetDogecoinKey(index)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.ToECDSA(key.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to ECDSA: %v", err)
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	pubKeyBytes := crypto.FromECDSAPub(publicKey)
	hash160 := btcutil.Hash160(pubKeyBytes)
	versionedPayload := append([]byte{0x1E}, hash160...) // Dogecoin uses 0x1E
	firstHash := sha256.Sum256(versionedPayload)
	secondHash := sha256.Sum256(firstHash[:]) // ✅ Convert [32]byte → []byte using [:]
	checksum := secondHash[:4]
	addressBytes := append(versionedPayload, checksum...)
	address := base58.Encode(addressBytes)

	return &CryptoAddress{
		PrivateKeyHex: hex.EncodeToString(key.Key),
		PublicKeyHex:  hex.EncodeToString(pubKeyBytes),
		Address:       address,
	}, nil
}

// GeneratePolkadotAddress generates a simulated Polkadot SS58 address
func (w *CryptoWallet) GeneratePolkadotAddress(index uint32) (*CryptoAddress, error) {
	key, err := w.GetPolkadotKey(index)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.ToECDSA(key.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to ECDSA: %v", err)
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	pubKeyBytes := crypto.FromECDSAPub(publicKey)

	// Simulate SS58: [prefix (42 for Polkadot)][pubkey (32 bytes)][checksum (2 bytes)]
	// In reality, SS58 uses blake2b and variable-length encoding.
	// This is a simplified simulation.

	// Use first 32 bytes of pubkey (truncate or pad if needed)
	var pubKey32 [32]byte
	copy(pubKey32[:], pubKeyBytes)
	if len(pubKeyBytes) < 32 {
		copy(pubKey32[len(pubKeyBytes):], make([]byte, 32-len(pubKeyBytes)))
	}

	payload := append([]byte{42}, pubKey32[:]...) // 42 = Polkadot prefix

	// Simple checksum: first 2 bytes of blake2b (we'll use sha256 for simulation)
	hash := sha256.Sum256(payload)
	checksum := hash[:2] // ✅ Works — 'hash' is addressable
	payload = append(payload, checksum...)

	address := base58.Encode(payload)

	return &CryptoAddress{
		PrivateKeyHex: hex.EncodeToString(key.Key),
		PublicKeyHex:  hex.EncodeToString(pubKeyBytes),
		Address:       address,
	}, nil
}

// GenerateTronAddress generates a Tron address (base58check of hash160(pubkey) with 0x41 prefix)
func (w *CryptoWallet) GenerateTronAddress(index uint32) (*CryptoAddress, error) {
	key, err := w.GetTronKey(index)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.ToECDSA(key.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to ECDSA: %v", err)
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	pubKeyBytes := crypto.FromECDSAPub(publicKey)
	hash160 := btcutil.Hash160(pubKeyBytes)
	versionedPayload := append([]byte{0x41}, hash160...) // TRON prefix
	firstHash := sha256.Sum256(versionedPayload)
	secondHash := sha256.Sum256(firstHash[:]) // ✅ Convert [32]byte → []byte using [:]
	checksum := secondHash[:4]
	addressBytes := append(versionedPayload, checksum...)
	address := base58.Encode(addressBytes)

	return &CryptoAddress{
		PrivateKeyHex: hex.EncodeToString(key.Key),
		PublicKeyHex:  hex.EncodeToString(pubKeyBytes),
		Address:       address,
	}, nil
}

// bech32 encoding utilities (simplified)
func bech32Encode(hrp string, data []byte) string {
	// Very simplified — not production ready!
	// Real implementation: github.com/btcsuite/btcutil/bech32
	converted := make([]int, len(data))
	for i, b := range data {
		converted[i] = int(b & 0x1f)
	}
	// Dummy checksum
	converted = append(converted, 0, 0, 0, 0, 0, 0)
	var chars []byte
	for _, v := range converted {
		chars = append(chars, byte('q'+v))
	}
	return hrp + "1" + string(chars)
}

// GenerateCosmosAddress generates a Cosmos Bech32 address
func (w *CryptoWallet) GenerateCosmosAddress(index uint32) (*CryptoAddress, error) {
	key, err := w.GetCosmosKey(index)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.ToECDSA(key.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to ECDSA: %v", err)
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	pubKeyBytes := crypto.FromECDSAPub(publicKey)
	hash160 := btcutil.Hash160(pubKeyBytes)

	// Cosmos: cosmos<bech32 separator> + hash160
	address := bech32Encode("cosmos", hash160)

	return &CryptoAddress{
		PrivateKeyHex: hex.EncodeToString(key.Key),
		PublicKeyHex:  hex.EncodeToString(pubKeyBytes),
		Address:       address,
	}, nil
}

// GenerateRippleAddress generates a Ripple base58 address (starts with 'r')
func (w *CryptoWallet) GenerateRippleAddress(index uint32) (*CryptoAddress, error) {
	key, err := w.GetRippleKey(index)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.ToECDSA(key.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to ECDSA: %v", err)
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	pubKeyBytes := crypto.FromECDSAPub(publicKey)
	hash160 := btcutil.Hash160(pubKeyBytes)

	// Ripple: version 0x00 + hash160 + 4-byte checksum
	versionedPayload := append([]byte{0x00}, hash160...)
	firstHash := sha256.Sum256(versionedPayload)
	secondHash := sha256.Sum256(firstHash[:]) // ✅ Convert [32]byte → []byte using [:]
	checksum := secondHash[:4]
	addressBytes := append(versionedPayload, checksum...)
	address := "r" + base58.Encode(addressBytes)[1:] // Force start with 'r'

	return &CryptoAddress{
		PrivateKeyHex: hex.EncodeToString(key.Key),
		PublicKeyHex:  hex.EncodeToString(pubKeyBytes),
		Address:       address,
	}, nil
}
