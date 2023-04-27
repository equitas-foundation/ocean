// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package queries

import (
	"context"
	"database/sql"
)

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM account WHERE namespace = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, namespace string) error {
	_, err := q.db.Exec(ctx, deleteAccount, namespace)
	return err
}

const deleteAccountScripts = `-- name: DeleteAccountScripts :exec
DELETE FROM account_script_info WHERE fk_account_name = $1
`

func (q *Queries) DeleteAccountScripts(ctx context.Context, fkAccountName string) error {
	_, err := q.db.Exec(ctx, deleteAccountScripts, fkAccountName)
	return err
}

const deleteTransactionInputAccounts = `-- name: DeleteTransactionInputAccounts :exec
DELETE FROM tx_input_account WHERE fk_tx_id=$1
`

func (q *Queries) DeleteTransactionInputAccounts(ctx context.Context, fkTxID string) error {
	_, err := q.db.Exec(ctx, deleteTransactionInputAccounts, fkTxID)
	return err
}

const deleteUtxoStatuses = `-- name: DeleteUtxoStatuses :exec
DELETE FROM utxo_status WHERE fk_utxo_id = $1
`

func (q *Queries) DeleteUtxoStatuses(ctx context.Context, fkUtxoID int32) error {
	_, err := q.db.Exec(ctx, deleteUtxoStatuses, fkUtxoID)
	return err
}

const deleteUtxosForAccountName = `-- name: DeleteUtxosForAccountName :exec
DELETE FROM utxo WHERE account_name=$1
`

func (q *Queries) DeleteUtxosForAccountName(ctx context.Context, accountName string) error {
	_, err := q.db.Exec(ctx, deleteUtxosForAccountName, accountName)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT namespace, index, label, xpubs, derivation_path, next_external_index, next_internal_index, fk_wallet_id FROM account WHERE namespace = $1 OR label = $1
`

func (q *Queries) GetAccount(ctx context.Context, namespace string) (Account, error) {
	row := q.db.QueryRow(ctx, getAccount, namespace)
	var i Account
	err := row.Scan(
		&i.Namespace,
		&i.Index,
		&i.Label,
		&i.Xpubs,
		&i.DerivationPath,
		&i.NextExternalIndex,
		&i.NextInternalIndex,
		&i.FkWalletID,
	)
	return i, err
}

const getAllUtxos = `-- name: GetAllUtxos :many
SELECT u.id, tx_id, vout, value, asset, value_commitment, asset_commitment, value_blinder, asset_blinder, script, redeem_script, nonce, range_proof, surjection_proof, account_name, lock_timestamp, lock_expiry_timestamp, us.id, block_height, block_time, block_hash, status, fk_utxo_id FROM utxo u left join utxo_status us on u.id = us.fk_utxo_id
`

type GetAllUtxosRow struct {
	ID                  int32
	TxID                string
	Vout                int32
	Value               int64
	Asset               string
	ValueCommitment     []byte
	AssetCommitment     []byte
	ValueBlinder        []byte
	AssetBlinder        []byte
	Script              []byte
	RedeemScript        []byte
	Nonce               []byte
	RangeProof          []byte
	SurjectionProof     []byte
	AccountName         string
	LockTimestamp       int64
	LockExpiryTimestamp int64
	ID_2                sql.NullInt32
	BlockHeight         sql.NullInt32
	BlockTime           sql.NullInt64
	BlockHash           sql.NullString
	Status              sql.NullInt32
	FkUtxoID            sql.NullInt32
}

func (q *Queries) GetAllUtxos(ctx context.Context) ([]GetAllUtxosRow, error) {
	rows, err := q.db.Query(ctx, getAllUtxos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllUtxosRow
	for rows.Next() {
		var i GetAllUtxosRow
		if err := rows.Scan(
			&i.ID,
			&i.TxID,
			&i.Vout,
			&i.Value,
			&i.Asset,
			&i.ValueCommitment,
			&i.AssetCommitment,
			&i.ValueBlinder,
			&i.AssetBlinder,
			&i.Script,
			&i.RedeemScript,
			&i.Nonce,
			&i.RangeProof,
			&i.SurjectionProof,
			&i.AccountName,
			&i.LockTimestamp,
			&i.LockExpiryTimestamp,
			&i.ID_2,
			&i.BlockHeight,
			&i.BlockTime,
			&i.BlockHash,
			&i.Status,
			&i.FkUtxoID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTransaction = `-- name: GetTransaction :many
SELECT tx_id, tx_hex, block_hash, block_height, id, account_name, fk_tx_id FROM transaction t left join tx_input_account tia on t.tx_id = tia.fk_tx_id WHERE tx_id=$1
`

type GetTransactionRow struct {
	TxID        string
	TxHex       string
	BlockHash   string
	BlockHeight int32
	ID          sql.NullInt32
	AccountName sql.NullString
	FkTxID      sql.NullString
}

func (q *Queries) GetTransaction(ctx context.Context, txID string) ([]GetTransactionRow, error) {
	rows, err := q.db.Query(ctx, getTransaction, txID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTransactionRow
	for rows.Next() {
		var i GetTransactionRow
		if err := rows.Scan(
			&i.TxID,
			&i.TxHex,
			&i.BlockHash,
			&i.BlockHeight,
			&i.ID,
			&i.AccountName,
			&i.FkTxID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUtxoForKey = `-- name: GetUtxoForKey :many
SELECT u.id, tx_id, vout, value, asset, value_commitment, asset_commitment, value_blinder, asset_blinder, script, redeem_script, nonce, range_proof, surjection_proof, account_name, lock_timestamp, lock_expiry_timestamp, us.id, block_height, block_time, block_hash, status, fk_utxo_id FROM utxo u left join utxo_status us on u.id = us.fk_utxo_id
WHERE u.tx_id = $1 AND u.vout = $2
`

type GetUtxoForKeyParams struct {
	TxID string
	Vout int32
}

type GetUtxoForKeyRow struct {
	ID                  int32
	TxID                string
	Vout                int32
	Value               int64
	Asset               string
	ValueCommitment     []byte
	AssetCommitment     []byte
	ValueBlinder        []byte
	AssetBlinder        []byte
	Script              []byte
	RedeemScript        []byte
	Nonce               []byte
	RangeProof          []byte
	SurjectionProof     []byte
	AccountName         string
	LockTimestamp       int64
	LockExpiryTimestamp int64
	ID_2                sql.NullInt32
	BlockHeight         sql.NullInt32
	BlockTime           sql.NullInt64
	BlockHash           sql.NullString
	Status              sql.NullInt32
	FkUtxoID            sql.NullInt32
}

func (q *Queries) GetUtxoForKey(ctx context.Context, arg GetUtxoForKeyParams) ([]GetUtxoForKeyRow, error) {
	rows, err := q.db.Query(ctx, getUtxoForKey, arg.TxID, arg.Vout)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUtxoForKeyRow
	for rows.Next() {
		var i GetUtxoForKeyRow
		if err := rows.Scan(
			&i.ID,
			&i.TxID,
			&i.Vout,
			&i.Value,
			&i.Asset,
			&i.ValueCommitment,
			&i.AssetCommitment,
			&i.ValueBlinder,
			&i.AssetBlinder,
			&i.Script,
			&i.RedeemScript,
			&i.Nonce,
			&i.RangeProof,
			&i.SurjectionProof,
			&i.AccountName,
			&i.LockTimestamp,
			&i.LockExpiryTimestamp,
			&i.ID_2,
			&i.BlockHeight,
			&i.BlockTime,
			&i.BlockHash,
			&i.Status,
			&i.FkUtxoID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUtxosForAccount = `-- name: GetUtxosForAccount :many
SELECT u.id, tx_id, vout, value, asset, value_commitment, asset_commitment, value_blinder, asset_blinder, script, redeem_script, nonce, range_proof, surjection_proof, account_name, lock_timestamp, lock_expiry_timestamp, us.id, block_height, block_time, block_hash, status, fk_utxo_id FROM utxo u left join utxo_status us on u.id = us.fk_utxo_id
WHERE u.account_name = $1
`

type GetUtxosForAccountRow struct {
	ID                  int32
	TxID                string
	Vout                int32
	Value               int64
	Asset               string
	ValueCommitment     []byte
	AssetCommitment     []byte
	ValueBlinder        []byte
	AssetBlinder        []byte
	Script              []byte
	RedeemScript        []byte
	Nonce               []byte
	RangeProof          []byte
	SurjectionProof     []byte
	AccountName         string
	LockTimestamp       int64
	LockExpiryTimestamp int64
	ID_2                sql.NullInt32
	BlockHeight         sql.NullInt32
	BlockTime           sql.NullInt64
	BlockHash           sql.NullString
	Status              sql.NullInt32
	FkUtxoID            sql.NullInt32
}

func (q *Queries) GetUtxosForAccount(ctx context.Context, accountName string) ([]GetUtxosForAccountRow, error) {
	rows, err := q.db.Query(ctx, getUtxosForAccount, accountName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUtxosForAccountRow
	for rows.Next() {
		var i GetUtxosForAccountRow
		if err := rows.Scan(
			&i.ID,
			&i.TxID,
			&i.Vout,
			&i.Value,
			&i.Asset,
			&i.ValueCommitment,
			&i.AssetCommitment,
			&i.ValueBlinder,
			&i.AssetBlinder,
			&i.Script,
			&i.RedeemScript,
			&i.Nonce,
			&i.RangeProof,
			&i.SurjectionProof,
			&i.AccountName,
			&i.LockTimestamp,
			&i.LockExpiryTimestamp,
			&i.ID_2,
			&i.BlockHeight,
			&i.BlockTime,
			&i.BlockHash,
			&i.Status,
			&i.FkUtxoID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUtxosForAccountName = `-- name: GetUtxosForAccountName :many
SELECT id, tx_id, vout, value, asset, value_commitment, asset_commitment, value_blinder, asset_blinder, script, redeem_script, nonce, range_proof, surjection_proof, account_name, lock_timestamp, lock_expiry_timestamp FROM utxo WHERE account_name=$1
`

func (q *Queries) GetUtxosForAccountName(ctx context.Context, accountName string) ([]Utxo, error) {
	rows, err := q.db.Query(ctx, getUtxosForAccountName, accountName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Utxo
	for rows.Next() {
		var i Utxo
		if err := rows.Scan(
			&i.ID,
			&i.TxID,
			&i.Vout,
			&i.Value,
			&i.Asset,
			&i.ValueCommitment,
			&i.AssetCommitment,
			&i.ValueBlinder,
			&i.AssetBlinder,
			&i.Script,
			&i.RedeemScript,
			&i.Nonce,
			&i.RangeProof,
			&i.SurjectionProof,
			&i.AccountName,
			&i.LockTimestamp,
			&i.LockExpiryTimestamp,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getWalletAccountsAndScripts = `-- name: GetWalletAccountsAndScripts :many
SELECT w.id as walletId,w.encrypted_mnemonic,w.password_hash,w.birthday_block_height,w.root_path,w.ms_root_path,w.network_name,w.next_account_index,w.next_ms_account_index, a.namespace,a.label,a.index,a.xpubs,a.derivation_path as account_derivation_path,a.next_external_index,a.next_internal_index,a.fk_wallet_id,asi.script,asi.derivation_path as script_derivation_path,asi.fk_account_name FROM
wallet w LEFT JOIN account a ON w.id = a.fk_wallet_id
LEFT JOIN account_script_info asi on a.namespace = asi.fk_account_name
WHERE w.id = $1
`

type GetWalletAccountsAndScriptsRow struct {
	Walletid              string
	EncryptedMnemonic     []byte
	PasswordHash          []byte
	BirthdayBlockHeight   int32
	RootPath              string
	MsRootPath            string
	NetworkName           string
	NextAccountIndex      int32
	NextMsAccountIndex    int32
	Namespace             sql.NullString
	Label                 sql.NullString
	Index                 sql.NullInt32
	Xpubs                 []string
	AccountDerivationPath sql.NullString
	NextExternalIndex     sql.NullInt32
	NextInternalIndex     sql.NullInt32
	FkWalletID            sql.NullString
	Script                sql.NullString
	ScriptDerivationPath  sql.NullString
	FkAccountName         sql.NullString
}

func (q *Queries) GetWalletAccountsAndScripts(ctx context.Context, id string) ([]GetWalletAccountsAndScriptsRow, error) {
	rows, err := q.db.Query(ctx, getWalletAccountsAndScripts, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetWalletAccountsAndScriptsRow
	for rows.Next() {
		var i GetWalletAccountsAndScriptsRow
		if err := rows.Scan(
			&i.Walletid,
			&i.EncryptedMnemonic,
			&i.PasswordHash,
			&i.BirthdayBlockHeight,
			&i.RootPath,
			&i.MsRootPath,
			&i.NetworkName,
			&i.NextAccountIndex,
			&i.NextMsAccountIndex,
			&i.Namespace,
			&i.Label,
			&i.Index,
			&i.Xpubs,
			&i.AccountDerivationPath,
			&i.NextExternalIndex,
			&i.NextInternalIndex,
			&i.FkWalletID,
			&i.Script,
			&i.ScriptDerivationPath,
			&i.FkAccountName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertAccount = `-- name: InsertAccount :one
INSERT INTO account(namespace,label,index,xpubs,derivation_path,next_external_index,next_internal_index,fk_wallet_id)
VALUES($1,$2,$3,$4,$5,$6,$7,$8) RETURNING namespace, index, label, xpubs, derivation_path, next_external_index, next_internal_index, fk_wallet_id
`

type InsertAccountParams struct {
	Namespace         string
	Label             sql.NullString
	Index             int32
	Xpubs             []string
	DerivationPath    string
	NextExternalIndex int32
	NextInternalIndex int32
	FkWalletID        string
}

func (q *Queries) InsertAccount(ctx context.Context, arg InsertAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, insertAccount,
		arg.Namespace,
		arg.Label,
		arg.Index,
		arg.Xpubs,
		arg.DerivationPath,
		arg.NextExternalIndex,
		arg.NextInternalIndex,
		arg.FkWalletID,
	)
	var i Account
	err := row.Scan(
		&i.Namespace,
		&i.Index,
		&i.Label,
		&i.Xpubs,
		&i.DerivationPath,
		&i.NextExternalIndex,
		&i.NextInternalIndex,
		&i.FkWalletID,
	)
	return i, err
}

type InsertAccountScriptsParams struct {
	Script         string
	DerivationPath string
	FkAccountName  string
}

const insertTransaction = `-- name: InsertTransaction :one
INSERT INTO transaction(tx_id,tx_hex,block_hash,block_height)
VALUES($1,$2,$3,$4) RETURNING tx_id, tx_hex, block_hash, block_height
`

type InsertTransactionParams struct {
	TxID        string
	TxHex       string
	BlockHash   string
	BlockHeight int32
}

// TRANSACTION
func (q *Queries) InsertTransaction(ctx context.Context, arg InsertTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, insertTransaction,
		arg.TxID,
		arg.TxHex,
		arg.BlockHash,
		arg.BlockHeight,
	)
	var i Transaction
	err := row.Scan(
		&i.TxID,
		&i.TxHex,
		&i.BlockHash,
		&i.BlockHeight,
	)
	return i, err
}

const insertTransactionInputAccount = `-- name: InsertTransactionInputAccount :one
INSERT INTO tx_input_account(account_name, fk_tx_id)
VALUES($1,$2) RETURNING id, account_name, fk_tx_id
`

type InsertTransactionInputAccountParams struct {
	AccountName string
	FkTxID      string
}

func (q *Queries) InsertTransactionInputAccount(ctx context.Context, arg InsertTransactionInputAccountParams) (TxInputAccount, error) {
	row := q.db.QueryRow(ctx, insertTransactionInputAccount, arg.AccountName, arg.FkTxID)
	var i TxInputAccount
	err := row.Scan(&i.ID, &i.AccountName, &i.FkTxID)
	return i, err
}

const insertUtxo = `-- name: InsertUtxo :one
INSERT INTO utxo(tx_id,vout,value,asset,value_commitment,asset_commitment,value_blinder,asset_blinder,script,redeem_script,nonce,range_proof,surjection_proof,account_name,lock_timestamp,lock_expiry_timestamp)
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16) RETURNING id, tx_id, vout, value, asset, value_commitment, asset_commitment, value_blinder, asset_blinder, script, redeem_script, nonce, range_proof, surjection_proof, account_name, lock_timestamp, lock_expiry_timestamp
`

type InsertUtxoParams struct {
	TxID                string
	Vout                int32
	Value               int64
	Asset               string
	ValueCommitment     []byte
	AssetCommitment     []byte
	ValueBlinder        []byte
	AssetBlinder        []byte
	Script              []byte
	RedeemScript        []byte
	Nonce               []byte
	RangeProof          []byte
	SurjectionProof     []byte
	AccountName         string
	LockTimestamp       int64
	LockExpiryTimestamp int64
}

// UTXO
func (q *Queries) InsertUtxo(ctx context.Context, arg InsertUtxoParams) (Utxo, error) {
	row := q.db.QueryRow(ctx, insertUtxo,
		arg.TxID,
		arg.Vout,
		arg.Value,
		arg.Asset,
		arg.ValueCommitment,
		arg.AssetCommitment,
		arg.ValueBlinder,
		arg.AssetBlinder,
		arg.Script,
		arg.RedeemScript,
		arg.Nonce,
		arg.RangeProof,
		arg.SurjectionProof,
		arg.AccountName,
		arg.LockTimestamp,
		arg.LockExpiryTimestamp,
	)
	var i Utxo
	err := row.Scan(
		&i.ID,
		&i.TxID,
		&i.Vout,
		&i.Value,
		&i.Asset,
		&i.ValueCommitment,
		&i.AssetCommitment,
		&i.ValueBlinder,
		&i.AssetBlinder,
		&i.Script,
		&i.RedeemScript,
		&i.Nonce,
		&i.RangeProof,
		&i.SurjectionProof,
		&i.AccountName,
		&i.LockTimestamp,
		&i.LockExpiryTimestamp,
	)
	return i, err
}

const insertUtxoStatus = `-- name: InsertUtxoStatus :one
INSERT INTO utxo_status(block_height,block_time,block_hash,status,fk_utxo_id)
VALUES($1,$2,$3,$4,$5) RETURNING id, block_height, block_time, block_hash, status, fk_utxo_id
`

type InsertUtxoStatusParams struct {
	BlockHeight int32
	BlockTime   int64
	BlockHash   string
	Status      int32
	FkUtxoID    int32
}

func (q *Queries) InsertUtxoStatus(ctx context.Context, arg InsertUtxoStatusParams) (UtxoStatus, error) {
	row := q.db.QueryRow(ctx, insertUtxoStatus,
		arg.BlockHeight,
		arg.BlockTime,
		arg.BlockHash,
		arg.Status,
		arg.FkUtxoID,
	)
	var i UtxoStatus
	err := row.Scan(
		&i.ID,
		&i.BlockHeight,
		&i.BlockTime,
		&i.BlockHash,
		&i.Status,
		&i.FkUtxoID,
	)
	return i, err
}

const insertWallet = `-- name: InsertWallet :one
INSERT INTO wallet(id, encrypted_mnemonic,password_hash,birthday_block_height,root_path,ms_root_path,network_name,next_account_index,next_ms_account_index)
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id, encrypted_mnemonic, password_hash, birthday_block_height, root_path, ms_root_path, network_name, next_account_index, next_ms_account_index
`

type InsertWalletParams struct {
	ID                  string
	EncryptedMnemonic   []byte
	PasswordHash        []byte
	BirthdayBlockHeight int32
	RootPath            string
	MsRootPath          string
	NetworkName         string
	NextAccountIndex    int32
	NextMsAccountIndex  int32
}

// WALLET & ACCOUNT
func (q *Queries) InsertWallet(ctx context.Context, arg InsertWalletParams) (Wallet, error) {
	row := q.db.QueryRow(ctx, insertWallet,
		arg.ID,
		arg.EncryptedMnemonic,
		arg.PasswordHash,
		arg.BirthdayBlockHeight,
		arg.RootPath,
		arg.MsRootPath,
		arg.NetworkName,
		arg.NextAccountIndex,
		arg.NextMsAccountIndex,
	)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.EncryptedMnemonic,
		&i.PasswordHash,
		&i.BirthdayBlockHeight,
		&i.RootPath,
		&i.MsRootPath,
		&i.NetworkName,
		&i.NextAccountIndex,
		&i.NextMsAccountIndex,
	)
	return i, err
}

const resetTransactions = `-- name: ResetTransactions :exec
DELETE FROM transaction
`

func (q *Queries) ResetTransactions(ctx context.Context) error {
	_, err := q.db.Exec(ctx, resetTransactions)
	return err
}

const resetUtxos = `-- name: ResetUtxos :exec
DELETE FROM utxo
`

func (q *Queries) ResetUtxos(ctx context.Context) error {
	_, err := q.db.Exec(ctx, resetUtxos)
	return err
}

const resetWallet = `-- name: ResetWallet :exec
DELETE FROM wallet
`

func (q *Queries) ResetWallet(ctx context.Context) error {
	_, err := q.db.Exec(ctx, resetWallet)
	return err
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE account SET next_external_index = $1, next_internal_index = $2, label = $3 WHERE namespace = $4 RETURNING namespace, index, label, xpubs, derivation_path, next_external_index, next_internal_index, fk_wallet_id
`

type UpdateAccountParams struct {
	NextExternalIndex int32
	NextInternalIndex int32
	Label             sql.NullString
	Namespace         string
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, updateAccount,
		arg.NextExternalIndex,
		arg.NextInternalIndex,
		arg.Label,
		arg.Namespace,
	)
	var i Account
	err := row.Scan(
		&i.Namespace,
		&i.Index,
		&i.Label,
		&i.Xpubs,
		&i.DerivationPath,
		&i.NextExternalIndex,
		&i.NextInternalIndex,
		&i.FkWalletID,
	)
	return i, err
}

const updateTransaction = `-- name: UpdateTransaction :one
UPDATE transaction SET tx_hex=$1,block_hash=$2,block_height=$3 WHERE tx_id=$4 RETURNING tx_id, tx_hex, block_hash, block_height
`

type UpdateTransactionParams struct {
	TxHex       string
	BlockHash   string
	BlockHeight int32
	TxID        string
}

func (q *Queries) UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, updateTransaction,
		arg.TxHex,
		arg.BlockHash,
		arg.BlockHeight,
		arg.TxID,
	)
	var i Transaction
	err := row.Scan(
		&i.TxID,
		&i.TxHex,
		&i.BlockHash,
		&i.BlockHeight,
	)
	return i, err
}

const updateUtxo = `-- name: UpdateUtxo :one
UPDATE utxo SET value=$1,asset=$2,value_commitment=$3,asset_commitment=$4,value_blinder=$5,asset_blinder=$6,script=$7,redeem_script=$8,nonce=$9,range_proof=$10,surjection_proof=$11,account_name=$12,lock_timestamp=$13, lock_expiry_timestamp=$14 WHERE tx_id=$15 and vout=$16 RETURNING id, tx_id, vout, value, asset, value_commitment, asset_commitment, value_blinder, asset_blinder, script, redeem_script, nonce, range_proof, surjection_proof, account_name, lock_timestamp, lock_expiry_timestamp
`

type UpdateUtxoParams struct {
	Value               int64
	Asset               string
	ValueCommitment     []byte
	AssetCommitment     []byte
	ValueBlinder        []byte
	AssetBlinder        []byte
	Script              []byte
	RedeemScript        []byte
	Nonce               []byte
	RangeProof          []byte
	SurjectionProof     []byte
	AccountName         string
	LockTimestamp       int64
	LockExpiryTimestamp int64
	TxID                string
	Vout                int32
}

func (q *Queries) UpdateUtxo(ctx context.Context, arg UpdateUtxoParams) (Utxo, error) {
	row := q.db.QueryRow(ctx, updateUtxo,
		arg.Value,
		arg.Asset,
		arg.ValueCommitment,
		arg.AssetCommitment,
		arg.ValueBlinder,
		arg.AssetBlinder,
		arg.Script,
		arg.RedeemScript,
		arg.Nonce,
		arg.RangeProof,
		arg.SurjectionProof,
		arg.AccountName,
		arg.LockTimestamp,
		arg.LockExpiryTimestamp,
		arg.TxID,
		arg.Vout,
	)
	var i Utxo
	err := row.Scan(
		&i.ID,
		&i.TxID,
		&i.Vout,
		&i.Value,
		&i.Asset,
		&i.ValueCommitment,
		&i.AssetCommitment,
		&i.ValueBlinder,
		&i.AssetBlinder,
		&i.Script,
		&i.RedeemScript,
		&i.Nonce,
		&i.RangeProof,
		&i.SurjectionProof,
		&i.AccountName,
		&i.LockTimestamp,
		&i.LockExpiryTimestamp,
	)
	return i, err
}

const updateWallet = `-- name: UpdateWallet :one
UPDATE wallet SET encrypted_mnemonic = $2, password_hash = $3, birthday_block_height = $4, root_path = $5, ms_root_path = $6, network_name = $7, next_account_index = $8, next_ms_account_index = $9 WHERE id = $1 RETURNING id, encrypted_mnemonic, password_hash, birthday_block_height, root_path, ms_root_path, network_name, next_account_index, next_ms_account_index
`

type UpdateWalletParams struct {
	ID                  string
	EncryptedMnemonic   []byte
	PasswordHash        []byte
	BirthdayBlockHeight int32
	RootPath            string
	MsRootPath          string
	NetworkName         string
	NextAccountIndex    int32
	NextMsAccountIndex  int32
}

func (q *Queries) UpdateWallet(ctx context.Context, arg UpdateWalletParams) (Wallet, error) {
	row := q.db.QueryRow(ctx, updateWallet,
		arg.ID,
		arg.EncryptedMnemonic,
		arg.PasswordHash,
		arg.BirthdayBlockHeight,
		arg.RootPath,
		arg.MsRootPath,
		arg.NetworkName,
		arg.NextAccountIndex,
		arg.NextMsAccountIndex,
	)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.EncryptedMnemonic,
		&i.PasswordHash,
		&i.BirthdayBlockHeight,
		&i.RootPath,
		&i.MsRootPath,
		&i.NetworkName,
		&i.NextAccountIndex,
		&i.NextMsAccountIndex,
	)
	return i, err
}
