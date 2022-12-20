package domain

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	path "github.com/equitas-foundation/bamp-ocean/pkg/wallet/derivation-path"
	multisig "github.com/equitas-foundation/bamp-ocean/pkg/wallet/multi-sig"
	singlesig "github.com/equitas-foundation/bamp-ocean/pkg/wallet/single-sig"
	"github.com/vulpemventures/go-elements/network"
)

const (
	externalChain = 0
	internalChain = 1
)

var (
	ErrWalletMissingMnemonic         = fmt.Errorf("missing mnemonic")
	ErrWalletMissingPassword         = fmt.Errorf("missing password")
	ErrWalletMissingNetwork          = fmt.Errorf("missing network name")
	ErrWalletMissingBirthdayBlock    = fmt.Errorf("missing birthday block height")
	ErrWalletLocked                  = fmt.Errorf("wallet is locked")
	ErrWalletUnlocked                = fmt.Errorf("wallet must be locked")
	ErrWalletMaxAccountNumberReached = fmt.Errorf("reached max number of accounts")
	ErrWalletInvalidPassword         = fmt.Errorf("wrong password")
	ErrWalletInvalidNetwork          = fmt.Errorf("unknown network")
	ErrAccountNotFound               = fmt.Errorf("account not found in wallet")

	networks = map[string]*network.Network{
		"liquid":  &network.Liquid,
		"testnet": &network.Testnet,
		"regtest": &network.Regtest,
	}
)

// AddressInfo holds useful info about a derived address.
type AddressInfo struct {
	AccountKey     AccountKey
	Address        string
	BlindingKey    []byte
	DerivationPath string
	Script         string
	RedeemScript   []byte
}

// Wallet is the data structure representing a secure HD wallet, ie. protected
// by a password that encrypts/decrypts the mnemonic seed.
type Wallet struct {
	EncryptedMnemonic   []byte
	PasswordHash        []byte
	BirthdayBlockHeight uint32
	RootPath            string
	MSRootPath          string
	NetworkName         string
	AccountsByKey       map[string]*Account
	AccountKeysByIndex  map[uint32]string
	AccountKeysByName   map[string]string
	NextAccountIndex    uint32
	NextMSAccountIndex  uint32
}

// NewWallet encrypts the provided mnemonic with the passhrase and returns a new
// Wallet initialized with the encrypted mnemonic, the hash of the password,
// the given root path, network and possible a list of accounts for an already
// used one.
// The Wallet is locked by default since it is initialized without the mnemonic
// in plain text.
func NewWallet(
	mnemonic []string, password, rootPath, msRootPath, network string,
	birthdayBlock uint32, accounts []Account,
) (*Wallet, error) {
	if len(mnemonic) <= 0 {
		return nil, ErrWalletMissingMnemonic
	}
	if len(password) <= 0 {
		return nil, ErrWalletMissingPassword
	}
	if birthdayBlock == 0 {
		return nil, ErrWalletMissingBirthdayBlock
	}
	if network == "" {
		return nil, ErrWalletMissingNetwork
	}
	if _, ok := networks[network]; !ok {
		return nil, ErrWalletInvalidNetwork
	}

	if _, err := singlesig.NewWalletFromMnemonic(singlesig.NewWalletFromMnemonicArgs{
		RootPath: rootPath,
		Mnemonic: mnemonic,
	}); err != nil {
		return nil, err
	}

	strMnemonic := strings.Join(mnemonic, " ")
	encryptedMnemonic, err := MnemonicCypher.Encrypt(
		[]byte(strMnemonic), []byte(password),
	)
	if err != nil {
		return nil, err
	}

	accountsByKey := make(map[string]*Account)
	accountKeysByIndex := make(map[uint32]string)
	accountKeysByName := make(map[string]string)
	for _, a := range accounts {
		key := a.Info.Key
		accountsByKey[key.String()] = &a
		accountKeysByIndex[key.Index] = key.String()
		accountKeysByName[key.Name] = key.String()
	}

	return &Wallet{
		EncryptedMnemonic:   encryptedMnemonic,
		PasswordHash:        btcutil.Hash160([]byte(password)),
		BirthdayBlockHeight: birthdayBlock,
		RootPath:            rootPath,
		MSRootPath:          msRootPath,
		AccountsByKey:       accountsByKey,
		AccountKeysByIndex:  accountKeysByIndex,
		AccountKeysByName:   accountKeysByName,
		NetworkName:         network,
	}, nil
}

// IsInitialized returns wheter the wallet is initialized with an encrypted
// mnemonic.
func (w *Wallet) IsInitialized() bool {
	return len(w.EncryptedMnemonic) > 0
}

// IsLocked returns whether the wallet is initialized and the plaintext
// mnemonic is set in its store.
func (w *Wallet) IsLocked() bool {
	return !w.IsInitialized() || !MnemonicStore.IsSet()
}

// GetMnemonic safely returns the plaintext mnemonic.
func (w *Wallet) GetMnemonic() ([]string, error) {
	if w.IsLocked() {
		return nil, ErrWalletLocked
	}

	return MnemonicStore.Get(), nil
}

// GetMnemonic safely returns the master blinding key.
func (w *Wallet) GetMasterBlindingKey() (string, error) {
	if w.IsLocked() {
		return "", ErrWalletLocked
	}

	mnemonic := MnemonicStore.Get()
	ww, _ := singlesig.NewWalletFromMnemonic(singlesig.NewWalletFromMnemonicArgs{
		RootPath: w.RootPath,
		Mnemonic: mnemonic,
	})
	return ww.MasterBlindingKey()
}

// Lock locks the Wallet by wiping the plaintext mnemonic from its store.
func (w *Wallet) Lock(password string) error {
	if w.IsLocked() {
		return nil
	}

	if !w.isValidPassword(password) {
		return ErrWalletInvalidPassword
	}

	MnemonicStore.Unset()
	return nil
}

// Unlock attempts to decrypt the encrypted mnemonic with the provided
// password.
func (w *Wallet) Unlock(password string) error {
	if !w.IsLocked() {
		return nil
	}

	if !w.isValidPassword(password) {
		return ErrWalletInvalidPassword
	}

	mnemonic, err := MnemonicCypher.Decrypt(w.EncryptedMnemonic, []byte(password))
	if err != nil {
		return err
	}

	MnemonicStore.Set(string(mnemonic))
	return nil
}

// ChangePassword attempts to unlock the wallet with the given currentPassword,
// then encrypts the plaintext mnemonic again with new password, stores its hash
// and, finally, locks the Wallet again.
func (w *Wallet) ChangePassword(currentPassword, newPassword string) error {
	if !w.IsLocked() {
		return ErrWalletUnlocked
	}
	if !w.isValidPassword(currentPassword) {
		return ErrWalletInvalidPassword
	}

	mnemonic, err := MnemonicCypher.Decrypt(w.EncryptedMnemonic, []byte(currentPassword))
	if err != nil {
		return err
	}

	encryptedMnemonic, err := MnemonicCypher.Encrypt(mnemonic, []byte(newPassword))
	if err != nil {
		return err
	}

	w.EncryptedMnemonic = encryptedMnemonic
	w.PasswordHash = btcutil.Hash160([]byte(newPassword))
	return nil
}

// CreateAccount creates a new account with the given name by preventing
// collisions with existing ones. If successful, returns the Account created.
func (w *Wallet) CreateAccount(name string, birthdayBlock uint32) (*Account, error) {
	if w.IsLocked() {
		return nil, ErrWalletLocked
	}
	if _, ok := w.AccountKeysByName[name]; ok {
		return nil, nil
	}
	if w.NextAccountIndex == hdkeychain.HardenedKeyStart {
		return nil, ErrWalletMaxAccountNumberReached
	}

	mnemonic := MnemonicStore.Get()

	ww, _ := singlesig.NewWalletFromMnemonic(singlesig.NewWalletFromMnemonicArgs{
		RootPath: w.RootPath,
		Mnemonic: mnemonic,
	})
	xpub, _ := ww.AccountExtendedPublicKey(singlesig.ExtendedKeyArgs{Account: w.NextAccountIndex})

	accountKey := AccountKey{name, w.NextAccountIndex}
	derivationPath, _ := path.ParseDerivationPath(w.RootPath)
	derivationPath = append(derivationPath, w.NextAccountIndex+hdkeychain.HardenedKeyStart)
	bdayBlock := w.BirthdayBlockHeight
	if birthdayBlock > bdayBlock {
		bdayBlock = birthdayBlock
	}
	accountInfo := AccountInfo{
		accountKey, []string{xpub}, derivationPath.String(),
	}
	account := &Account{
		Info:                   accountInfo,
		DerivationPathByScript: make(map[string]string),
		BirthdayBlock:          bdayBlock,
	}

	w.AccountsByKey[accountKey.String()] = account
	w.AccountKeysByIndex[accountKey.Index] = accountKey.String()
	w.AccountKeysByName[accountKey.Name] = accountKey.String()
	w.NextAccountIndex++
	return account, nil
}

// CreateMSAccount creates a new multisig account with the given name and cosigner xpub
// by preventing collisions with existing ones. If successful, returns the Account created.
func (w *Wallet) CreateMSAccount(
	name, cosignerXpub string, birthdayBlock uint32,
) (*Account, error) {
	if w.IsLocked() {
		return nil, ErrWalletLocked
	}
	if _, ok := w.AccountKeysByName[name]; ok {
		return nil, nil
	}
	if w.NextMSAccountIndex == hdkeychain.HardenedKeyStart {
		return nil, ErrWalletMaxAccountNumberReached
	}

	mnemonic := MnemonicStore.Get()

	derivationPath, _ := path.ParseRootDerivationPath(w.MSRootPath)
	derivationPath = append(
		derivationPath,
		w.NextMSAccountIndex+hdkeychain.HardenedKeyStart, // account'
		hdkeychain.HardenedKeyStart+2,                    // scriptType'
	)
	ww, _ := multisig.NewWalletFromMnemonic(multisig.NewWalletFromMnemonicArgs{
		RootPath: derivationPath.String(),
		Mnemonic: mnemonic,
		Xpubs:    []string{cosignerXpub},
	})

	accountKey := AccountKey{name, w.NextMSAccountIndex}
	xpub, _ := ww.AccountExtendedPublicKey()
	bdayBlock := w.BirthdayBlockHeight
	if birthdayBlock > bdayBlock {
		bdayBlock = birthdayBlock
	}
	xpubs := []string{cosignerXpub, xpub}
	sort.SliceStable(xpubs, func(i, j int) bool {
		return xpubs[i] < xpubs[j]
	})
	accountInfo := AccountInfo{
		accountKey, xpubs, derivationPath.String(),
	}
	account := &Account{
		Info:                   accountInfo,
		DerivationPathByScript: make(map[string]string),
		BirthdayBlock:          bdayBlock,
	}

	w.AccountsByKey[accountKey.String()] = account
	w.AccountKeysByIndex[accountKey.Index] = accountKey.String()
	w.AccountKeysByName[accountKey.Name] = accountKey.String()
	w.NextMSAccountIndex++
	return account, nil
}

// GetAccount safely returns an Account identified by the given name.
func (w *Wallet) GetAccount(accountName string) (*Account, error) {
	return w.getAccount(accountName)
}

// DeleteAccount safely removes an Account and all related stored info from the
// singlesig.
func (w *Wallet) DeleteAccount(accountName string) error {
	account, err := w.getAccount(accountName)
	if err != nil {
		return err
	}

	delete(w.AccountKeysByIndex, account.Info.Key.Index)
	delete(w.AccountKeysByName, account.Info.Key.Name)
	delete(w.AccountsByKey, account.Info.Key.String())
	return nil
}

// DeriveNextExternalAddressForAccount returns all useful info about the next
// new receiving address for the given account.
func (w *Wallet) DeriveNextExternalAddressForAccount(
	accountName string,
) (*AddressInfo, error) {
	return w.deriveNextAddressForAccount(accountName, externalChain)
}

// DeriveNextInternalAddressForAccount returns all useful info about the next
// new change address for the given account.
func (w *Wallet) DeriveNextInternalAddressForAccount(
	accountName string,
) (*AddressInfo, error) {
	return w.deriveNextAddressForAccount(accountName, internalChain)
}

// AllDerivedAddressesForAccount returns info about all derived receiving and
// change addresses derived so far for the given account.
func (w *Wallet) AllDerivedAddressesForAccount(
	accountName string,
) ([]AddressInfo, error) {
	return w.allDerivedAddressesForAccount(accountName, true)
}

// AllDerivedExternalAddressesForAccount returns info about all derived
// receiving addresses derived so far for the given account.
func (w *Wallet) AllDerivedExternalAddressesForAccount(
	accountName string,
) ([]AddressInfo, error) {
	return w.allDerivedAddressesForAccount(accountName, false)
}

func (w *Wallet) isValidPassword(password string) bool {
	return bytes.Equal(w.PasswordHash, btcutil.Hash160([]byte(password)))
}

func (w *Wallet) getAccount(accountName string) (*Account, error) {
	if w.IsLocked() {
		return nil, ErrWalletLocked
	}

	accountkey, ok := w.AccountKeysByName[accountName]
	if !ok {
		return nil, ErrAccountNotFound
	}
	return w.AccountsByKey[accountkey], nil
}

func (w *Wallet) deriveNextAddressForAccount(
	accountName string, chainIndex int,
) (*AddressInfo, error) {
	account, err := w.getAccount(accountName)
	if err != nil {
		return nil, err
	}

	if account.IsMultiSig() {
		return w.deriveNextAddressForMSAccount(account, chainIndex)
	}

	return w.deriveNextAddressForSSAccount(account, chainIndex)
}

func (w *Wallet) deriveNextAddressForSSAccount(
	account *Account, chainIndex int,
) (*AddressInfo, error) {
	mnemonic, _ := w.GetMnemonic()
	ww, _ := singlesig.NewWalletFromMnemonic(singlesig.NewWalletFromMnemonicArgs{
		RootPath: w.RootPath,
		Mnemonic: mnemonic,
	})

	addressIndex := account.NextExternalIndex
	if chainIndex == internalChain {
		addressIndex = account.NextInternalIndex
	}
	derivationPath := fmt.Sprintf(
		"%d'/%d/%d",
		account.Info.Key.Index, chainIndex, addressIndex,
	)
	net := networkFromName(w.NetworkName)
	addr, script, err := ww.DeriveConfidentialAddress(singlesig.DeriveConfidentialAddressArgs{
		DerivationPath: derivationPath,
		Network:        net,
	})
	if err != nil {
		return nil, err
	}

	blindingKey, _, _ := ww.DeriveBlindingKeyPair(singlesig.DeriveBlindingKeyPairArgs{
		Script: script,
	})

	account.addDerivationPath(hex.EncodeToString(script), derivationPath)
	if chainIndex == internalChain {
		account.incrementInternalIndex()
	} else {
		account.incrementExternalIndex()
	}

	return &AddressInfo{
		AccountKey:     account.Info.Key,
		Address:        addr,
		Script:         hex.EncodeToString(script),
		BlindingKey:    blindingKey.Serialize(),
		DerivationPath: derivationPath,
	}, nil
}

func (w *Wallet) deriveNextAddressForMSAccount(
	account *Account, chainIndex int,
) (*AddressInfo, error) {
	mnemonic, _ := w.GetMnemonic()
	wallet, _ := multisig.NewWalletFromMnemonic(
		multisig.NewWalletFromMnemonicArgs{
			RootPath: account.Info.DerivationPath,
			Mnemonic: mnemonic,
			Xpubs:    account.Info.Xpubs,
		},
	)

	addressIndex := account.NextExternalIndex
	if chainIndex == internalChain {
		addressIndex = account.NextInternalIndex
	}

	derivationPath := fmt.Sprintf("%d/%d", chainIndex, addressIndex)

	net := networkFromName(w.NetworkName)
	addr, script, redeemScript, err := wallet.DeriveConfidentialAddress(
		multisig.DeriveConfidentialAddressArgs{
			DerivationPath: derivationPath,
			Network:        net,
		},
	)
	if err != nil {
		return nil, err
	}

	blindingKey, _, _ := wallet.DeriveBlindingKeyPair(
		multisig.DeriveBlindingKeyPairArgs{
			Script: script,
		},
	)

	account.addDerivationPath(hex.EncodeToString(script), derivationPath)
	if chainIndex == internalChain {
		account.incrementInternalIndex()
	} else {
		account.incrementExternalIndex()
	}

	return &AddressInfo{
		AccountKey:     account.Info.Key,
		Address:        addr,
		Script:         hex.EncodeToString(script),
		BlindingKey:    blindingKey.Serialize(),
		DerivationPath: derivationPath,
		RedeemScript:   redeemScript,
	}, nil
}

func (w *Wallet) allDerivedAddressesForAccount(
	accountName string, includeInternals bool,
) ([]AddressInfo, error) {
	account, err := w.getAccount(accountName)
	if err != nil {
		return nil, err
	}

	if account.IsMultiSig() {
		return w.allDerivedAddressesForMSAccount(account, includeInternals)
	}
	return w.allDerivedAddressesForSSAccount(account, includeInternals)
}

func (w *Wallet) allDerivedAddressesForSSAccount(
	account *Account, includeInternals bool,
) ([]AddressInfo, error) {
	net := networkFromName(w.NetworkName)
	mnemonic, _ := w.GetMnemonic()
	ww, _ := singlesig.NewWalletFromMnemonic(singlesig.NewWalletFromMnemonicArgs{
		RootPath: w.RootPath,
		Mnemonic: mnemonic,
	})

	infoLen := account.NextExternalIndex
	if includeInternals {
		infoLen += account.NextInternalIndex
	}
	info := make([]AddressInfo, 0, infoLen)
	for i := 0; i < int(account.NextExternalIndex); i++ {
		derivationPath := fmt.Sprintf(
			"%d'/%d/%d",
			account.Info.Key.Index, externalChain, i,
		)
		addr, script, err := ww.DeriveConfidentialAddress(singlesig.DeriveConfidentialAddressArgs{
			DerivationPath: derivationPath,
			Network:        net,
		})
		if err != nil {
			return nil, err
		}
		key, _, _ := ww.DeriveBlindingKeyPair(singlesig.DeriveBlindingKeyPairArgs{
			Script: script,
		})
		info = append(info, AddressInfo{
			AccountKey:     account.Info.Key,
			Address:        addr,
			BlindingKey:    key.Serialize(),
			DerivationPath: derivationPath,
			Script:         hex.EncodeToString(script),
		})
	}
	if includeInternals {
		for i := 0; i < int(account.NextInternalIndex); i++ {
			derivationPath := fmt.Sprintf(
				"%d'/%d/%d",
				account.Info.Key.Index, internalChain, i,
			)
			addr, script, err := ww.DeriveConfidentialAddress(singlesig.DeriveConfidentialAddressArgs{
				DerivationPath: derivationPath,
				Network:        net,
			})
			if err != nil {
				return nil, err
			}
			key, _, _ := ww.DeriveBlindingKeyPair(singlesig.DeriveBlindingKeyPairArgs{
				Script: script,
			})
			info = append(info, AddressInfo{
				AccountKey:     account.Info.Key,
				Address:        addr,
				BlindingKey:    key.Serialize(),
				DerivationPath: derivationPath,
				Script:         hex.EncodeToString(script),
			})
		}
	}

	return info, nil
}

func (w *Wallet) allDerivedAddressesForMSAccount(
	account *Account, includeInternals bool,
) ([]AddressInfo, error) {
	net := networkFromName(w.NetworkName)
	mnemonic, _ := w.GetMnemonic()
	ww, _ := multisig.NewWalletFromMnemonic(multisig.NewWalletFromMnemonicArgs{
		RootPath: account.Info.DerivationPath,
		Mnemonic: mnemonic,
		Xpubs:    account.Info.Xpubs,
	})

	infoLen := account.NextExternalIndex
	if includeInternals {
		infoLen += account.NextInternalIndex
	}
	info := make([]AddressInfo, 0, infoLen)
	for i := 0; i < int(account.NextExternalIndex); i++ {
		derivationPath := fmt.Sprintf("%d/%d", externalChain, i)
		addr, script, redeemScript, err := ww.DeriveConfidentialAddress(
			multisig.DeriveConfidentialAddressArgs{
				DerivationPath: derivationPath,
				Network:        net,
			},
		)
		if err != nil {
			return nil, err
		}

		blindingKey, _, _ := ww.DeriveBlindingKeyPair(
			multisig.DeriveBlindingKeyPairArgs{
				Script: script,
			},
		)
		info = append(info, AddressInfo{
			AccountKey:     account.Info.Key,
			Address:        addr,
			BlindingKey:    blindingKey.Serialize(),
			DerivationPath: derivationPath,
			Script:         hex.EncodeToString(script),
			RedeemScript:   redeemScript,
		})
	}
	if includeInternals {
		for i := 0; i < int(account.NextInternalIndex); i++ {
			derivationPath := fmt.Sprintf("%d/%d", internalChain, i)
			addr, script, redeemScript, err := ww.DeriveConfidentialAddress(
				multisig.DeriveConfidentialAddressArgs{
					DerivationPath: derivationPath,
					Network:        net,
				},
			)
			if err != nil {
				return nil, err
			}
			blindingKey, _, _ := ww.DeriveBlindingKeyPair(multisig.DeriveBlindingKeyPairArgs{
				Script: script,
			})
			info = append(info, AddressInfo{
				AccountKey:     account.Info.Key,
				Address:        addr,
				BlindingKey:    blindingKey.Serialize(),
				DerivationPath: derivationPath,
				Script:         hex.EncodeToString(script),
				RedeemScript:   redeemScript,
			})
		}
	}

	return info, nil
}

func networkFromName(net string) *network.Network {
	return networks[net]
}
