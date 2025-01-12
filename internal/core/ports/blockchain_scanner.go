package ports

import (
	"github.com/equitas-foundation/bamp-ocean/internal/core/domain"
)

// BlockchainScanner is the abstraction for any kind of service representing an
// Elements node. It gives info about txs and utxos related to one or more HD
// accounts in a aync way (via channels), and lets broadcast transactions over
// the Liquid network.
type BlockchainScanner interface {
	// Start starts the service.
	Start()
	// Stop stops the service.
	Stop()

	// WatchForAccount instructs the scanner to start notifying about txs/utxos
	// related to the given list of addresses belonging to the given HD account.
	WatchForAccount(
		accountName string, startingBlockHeight uint32,
		addresses []domain.AddressInfo,
	)
	// WatchForUtxos instructs the scanner to notify when the given utxos are
	// either spent or confirmed.
	WatchForUtxos(
		accountName string, utxos []domain.UtxoInfo,
	)
	// RestoreAccount makes the scanner discover and retuen all the used
	// addresses for a certain account represented by its account index, xpub
	// and master blinding key.
	RestoreAccount(
		accountIndex uint32, accountName string,
		xpubs []string, masterBlindingKey []byte,
		startingBlockHeight, addressesThreshold uint32,
	) ([]domain.AddressInfo, []domain.AddressInfo, error)
	// StopWatchForAccount instructs the scanner to stop notifying about
	// txs/utxos related to any address belonging to the given HD account.
	StopWatchForAccount(accountName string)

	// GetUtxoChannel returns the channel where notification about utxos realated
	// to the given HD account are sent.
	GetUtxoChannel(accountName string) chan []*domain.Utxo
	// GetTxChannel returns the channel where notification about txs realated to
	// the given HD account are sent.
	GetTxChannel(accountName string) chan *domain.Transaction

	// GetLatestBlock returns the header of the latest block of the blockchain.
	GetLatestBlock() ([]byte, uint32, error)
	// GetBlockHash returns the hash of the block identified by its height.
	GetBlockHash(height uint32) ([]byte, error)
	// GetUtxos is a sync function to get info about the utxos represented by
	// given outpoints (UtxoKeys).
	GetUtxos(utxos []domain.Utxo) ([]domain.Utxo, error)
	// GetUtxos is a sync function to get all utxos for the given list of addresses.
	GetUtxosForAddresses(addresses []domain.AddressInfo) ([]*domain.Utxo, error)
	// BroadcastTransaction sends the given raw tx (in hex string) over the
	// network in order to be included in a later block of the Liquid blockchain.
	BroadcastTransaction(txHex string) (string, error)
}
