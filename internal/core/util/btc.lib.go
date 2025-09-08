package util

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

// Wallet represents a Bitcoin wallet with mnemonic and derived keys
type Wallet struct {
	Mnemonic   string
	Passphrase string
	keys       map[string]*bip32.Key
	mux        sync.Mutex
}

// BitcoinAddress represents different Bitcoin address types
type BitcoinAddress struct {
	PrivateKeyHex string
	PublicKeyHex  string
	P2PKHAddress  string
	P2WPKHAddress string
	P2TRAddress   string
	WIF           string
}

// NewWallet creates a new Bitcoin wallet with a mnemonic
func NewWallet(passphrase string) (*Wallet, error) {
	// Generate 128-bit entropy (12 words mnemonic)
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return nil, fmt.Errorf("failed to generate entropy: %v", err)
	}

	// Create mnemonic from entropy
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, fmt.Errorf("failed to generate mnemonic: %v", err)
	}

	return &Wallet{
		Mnemonic:   mnemonic,
		Passphrase: passphrase,
		keys:       make(map[string]*bip32.Key),
	}, nil
}

// getKey retrieves or sets a key in the wallet's key map
func (w *Wallet) getKey(path string) (*bip32.Key, bool) {
	w.mux.Lock()
	defer w.mux.Unlock()
	key, ok := w.keys[path]
	return key, ok
}

func (w *Wallet) setKey(path string, key *bip32.Key) {
	w.mux.Lock()
	defer w.mux.Unlock()
	w.keys[path] = key
}

// GetSeed generates the seed from the mnemonic and passphrase
func (w *Wallet) GetSeed() []byte {
	return bip39.NewSeed(w.Mnemonic, w.Passphrase)
}

// GetMasterKey derives the master key from the seed
func (w *Wallet) GetMasterKey() (*bip32.Key, error) {
	path := "m"
	if key, ok := w.getKey(path); ok {
		return key, nil
	}

	key, err := bip32.NewMasterKey(w.GetSeed())
	if err != nil {
		return nil, fmt.Errorf("failed to generate master key: %v", err)
	}

	w.setKey(path, key)
	return key, nil
}

// GetBIP44Key derives a key following BIP-44 path: m/44'/0'/0'/0/index
func (w *Wallet) GetBIP44Key(index uint32) (*bip32.Key, error) {
	path := fmt.Sprintf("m/44'/0'/0'/0/%d", index)

	if key, ok := w.getKey(path); ok {
		return key, nil
	}

	// Derive the path step-by-step
	masterKey, err := w.GetMasterKey()
	if err != nil {
		return nil, err
	}

	// m/44'
	purposeKey, err := masterKey.NewChildKey(0x8000002C) // 44' (hardened)
	if err != nil {
		return nil, fmt.Errorf("failed to derive purpose key: %v", err)
	}

	// m/44'/0'
	coinTypeKey, err := purposeKey.NewChildKey(0x80000000) // 0' (BTC, hardened)
	if err != nil {
		return nil, fmt.Errorf("failed to derive coin type key: %v", err)
	}

	// m/44'/0'/0'
	accountKey, err := coinTypeKey.NewChildKey(0x80000000) // 0' (hardened)
	if err != nil {
		return nil, fmt.Errorf("failed to derive account key: %v", err)
	}

	// m/44'/0'/0'/0
	changeKey, err := accountKey.NewChildKey(0) // External chain
	if err != nil {
		return nil, fmt.Errorf("failed to derive change key: %v", err)
	}

	// m/44'/0'/0'/0/index
	key, err := changeKey.NewChildKey(index)
	if err != nil {
		return nil, fmt.Errorf("failed to derive index key: %v", err)
	}

	w.setKey(path, key)
	return key, nil
}

// GetBIP84Key derives a key following BIP-84 path: m/84'/0'/0'/0/index
func (w *Wallet) GetBIP84Key(index uint32) (*bip32.Key, error) {
	path := fmt.Sprintf("m/84'/0'/0'/0/%d", index)

	if key, ok := w.getKey(path); ok {
		return key, nil
	}

	// Derive the path step-by-step
	masterKey, err := w.GetMasterKey()
	if err != nil {
		return nil, err
	}

	// m/84'
	purposeKey, err := masterKey.NewChildKey(0x80000054) // 84' (hardened)
	if err != nil {
		return nil, fmt.Errorf("failed to derive purpose key: %v", err)
	}

	// m/84'/0'
	coinTypeKey, err := purposeKey.NewChildKey(0x80000000) // 0' (BTC, hardened)
	if err != nil {
		return nil, fmt.Errorf("failed to derive coin type key: %v", err)
	}

	// m/84'/0'/0'
	accountKey, err := coinTypeKey.NewChildKey(0x80000000) // 0' (hardened)
	if err != nil {
		return nil, fmt.Errorf("failed to derive account key: %v", err)
	}

	// m/84'/0'/0'/0
	changeKey, err := accountKey.NewChildKey(0) // External chain
	if err != nil {
		return nil, fmt.Errorf("failed to derive change key: %v", err)
	}

	// m/84'/0'/0'/0/index
	key, err := changeKey.NewChildKey(index)
	if err != nil {
		return nil, fmt.Errorf("failed to derive index key: %v", err)
	}

	w.setKey(path, key)
	return key, nil
}

// GetBIP86Key derives a key following BIP-86 path: m/86'/0'/0'/0/index
func (w *Wallet) GetBIP86Key(index uint32) (*bip32.Key, error) {
	path := fmt.Sprintf("m/86'/0'/0'/0/%d", index)

	if key, ok := w.getKey(path); ok {
		return key, nil
	}

	// Derive the path step-by-step
	masterKey, err := w.GetMasterKey()
	if err != nil {
		return nil, err
	}

	// m/86'
	purposeKey, err := masterKey.NewChildKey(0x80000056) // 86' (hardened)
	if err != nil {
		return nil, fmt.Errorf("failed to derive purpose key: %v", err)
	}

	// m/86'/0'
	coinTypeKey, err := purposeKey.NewChildKey(0x80000000) // 0' (BTC, hardened)
	if err != nil {
		return nil, fmt.Errorf("failed to derive coin type key: %v", err)
	}

	// m/86'/0'/0'
	accountKey, err := coinTypeKey.NewChildKey(0x80000000) // 0' (hardened)
	if err != nil {
		return nil, fmt.Errorf("failed to derive account key: %v", err)
	}

	// m/86'/0'/0'/0
	changeKey, err := accountKey.NewChildKey(0) // External chain
	if err != nil {
		return nil, fmt.Errorf("failed to derive change key: %v", err)
	}

	// m/86'/0'/0'/0/index
	key, err := changeKey.NewChildKey(index)
	if err != nil {
		return nil, fmt.Errorf("failed to derive index key: %v", err)
	}

	w.setKey(path, key)
	return key, nil
}

// GenerateBTCKey generates a Bitcoin key pair and addresses
func GenerateBTCKey(compress bool) (*BitcoinAddress, error) {
	privateKeyByte := make([]byte, 32)

	_, err := rand.Read(privateKeyByte)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	privateKey, publicKey := btcec.PrivKeyFromBytes(privateKeyByte)

	privateKeyHex := hex.EncodeToString(privateKey.Serialize())
	publicKeyHex := hex.EncodeToString(publicKey.SerializeCompressed())

	// Generate WIF
	wif, err := btcutil.NewWIF(privateKey, &chaincfg.MainNetParams, compress)
	if err != nil {
		return nil, fmt.Errorf("failed to generate WIF: %v", err)
	}

	// Generate P2PKH address
	serializedPubKey := wif.SerializePubKey()
	addressPubKey, err := btcutil.NewAddressPubKey(serializedPubKey, &chaincfg.MainNetParams)
	if err != nil {
		return nil, fmt.Errorf("failed to generate P2PKH address: %v", err)
	}

	// Generate P2WPKH (SegWit, bech32) address
	witnessProg := btcutil.Hash160(serializedPubKey)
	addressWitnessPubKeyHash, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, &chaincfg.MainNetParams)
	if err != nil {
		return nil, fmt.Errorf("failed to generate P2WPKH address: %v", err)
	}

	// Generate P2TR (Taproot) address
	tapKey := txscript.ComputeTaprootKeyNoScript(publicKey)
	addressTaproot, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(tapKey), &chaincfg.MainNetParams)
	if err != nil {
		return nil, fmt.Errorf("failed to generate P2TR address: %v", err)
	}

	return &BitcoinAddress{
		PrivateKeyHex: privateKeyHex,
		PublicKeyHex:  publicKeyHex,
		P2PKHAddress:  addressPubKey.EncodeAddress(),
		P2WPKHAddress: addressWitnessPubKeyHash.EncodeAddress(),
		P2TRAddress:   addressTaproot.EncodeAddress(),
		WIF:           wif.String(),
	}, nil
}

// GenerateBTCWalletKey generates a Bitcoin wallet key pair and addresses for a given derivation path
func (w *Wallet) GenerateBTCWalletKey(purpose uint32, index uint32) (*BitcoinAddress, error) {
	var key *bip32.Key
	var err error

	switch purpose {
	case 0x8000002C: // BIP-44
		key, err = w.GetBIP44Key(index)
	case 0x80000054: // BIP-84
		key, err = w.GetBIP84Key(index)
	case 0x80000056: // BIP-86
		key, err = w.GetBIP86Key(index)
	default:
		return nil, fmt.Errorf("unsupported purpose: %d", purpose)
	}

	if err != nil {
		return nil, err
	}

	privateKey, publicKey := btcec.PrivKeyFromBytes(key.Key)

	privateKeyHex := hex.EncodeToString(privateKey.Serialize())
	publicKeyHex := hex.EncodeToString(publicKey.SerializeCompressed())

	// Generate WIF
	wif, err := btcutil.NewWIF(privateKey, &chaincfg.MainNetParams, true)
	if err != nil {
		return nil, fmt.Errorf("failed to generate WIF: %v", err)
	}

	// Generate P2PKH address
	serializedPubKey := wif.SerializePubKey()
	addressPubKey, err := btcutil.NewAddressPubKey(serializedPubKey, &chaincfg.MainNetParams)
	if err != nil {
		return nil, fmt.Errorf("failed to generate P2PKH address: %v", err)
	}

	// Generate P2WPKH (SegWit, bech32) address
	witnessProg := btcutil.Hash160(serializedPubKey)
	addressWitnessPubKeyHash, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, &chaincfg.MainNetParams)
	if err != nil {
		return nil, fmt.Errorf("failed to generate P2WPKH address: %v", err)
	}

	// Generate P2TR (Taproot) address
	tapKey := txscript.ComputeTaprootKeyNoScript(publicKey)
	addressTaproot, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(tapKey), &chaincfg.MainNetParams)
	if err != nil {
		return nil, fmt.Errorf("failed to generate P2TR address: %v", err)
	}

	return &BitcoinAddress{
		PrivateKeyHex: privateKeyHex,
		PublicKeyHex:  publicKeyHex,
		P2PKHAddress:  addressPubKey.EncodeAddress(),
		P2WPKHAddress: addressWitnessPubKeyHash.EncodeAddress(),
		P2TRAddress:   addressTaproot.EncodeAddress(),
		WIF:           wif.String(),
	}, nil
}
