package internal

const (
	CreateAtomicSwapContractTable = `CREATE TABLE IF NOT EXISTS contracts (
		tx_hash TEXT,
		vout INT4,
		p2sh_addr TEXT,
		value INT8, -- dup in contract_spends table?
		secret_hash BYTEA,
		lock_time INT8
	);`

	CreateAtomicSwapSpendTable = `CREATE TABLE IF NOT EXISTS contract_spends (
		redeem BOOL, -- redeem or refund
		tx_hash TEXT,
		vin INT4,
		value INT8,
		contract_tx TEXT, -- entry MUST exist in contracts table
		contract_vout INT4,
		secret BYTEA
	);`
)
