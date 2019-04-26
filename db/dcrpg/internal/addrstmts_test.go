package internal

import "testing"

func BenchmarkMakeAddressRowMultilineInsertStatement_upsert3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MakeAddressRowMultilineInsertStatement(3, true, true)
	}
}

func TestMakeAddressRowMultilineInsertStatement(t *testing.T) {
	type args struct {
		N                int
		checked          bool
		updateOnConflict bool
	}
	tests := []struct {
		testName string
		args     args
		want     string
	}{
		{"unchecked 0", args{0, false, false}, ``},
		{"unchecked 1", args{1, false, false}, `INSERT INTO addresses (address, matching_tx_hash, tx_hash,
		tx_vin_vout_index, tx_vin_vout_row_id, value, block_time, is_funding, valid_mainchain, tx_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id;`},
		{"unchecked 2", args{2, false, false}, `INSERT INTO addresses (address, matching_tx_hash, tx_hash,
		tx_vin_vout_index, tx_vin_vout_row_id, value, block_time, is_funding, valid_mainchain, tx_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10), ($11, $12, $13, $14, $15, $16, $17, $18, $19, $20) RETURNING id;`},
		{"unchecked 3", args{3, false, false}, `INSERT INTO addresses (address, matching_tx_hash, tx_hash,
		tx_vin_vout_index, tx_vin_vout_row_id, value, block_time, is_funding, valid_mainchain, tx_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10), ($11, $12, $13, $14, $15, $16, $17, $18, $19, $20), ($21, $22, $23, $24, $25, $26, $27, $28, $29, $30) RETURNING id;`},
		{"upsert 1", args{1, true, true}, `INSERT INTO addresses (address, matching_tx_hash, tx_hash,
		tx_vin_vout_index, tx_vin_vout_row_id, value, block_time, is_funding, valid_mainchain, tx_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) ON CONFLICT (tx_vin_vout_row_id, address, is_funding) DO UPDATE
	SET matching_tx_hash = $2, tx_hash = $3, tx_vin_vout_index = $4,
		block_time = $7, valid_mainchain = $9 RETURNING id;`},
		{"upsert 2", args{2, true, true}, `INSERT INTO addresses (address, matching_tx_hash, tx_hash,
		tx_vin_vout_index, tx_vin_vout_row_id, value, block_time, is_funding, valid_mainchain, tx_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10), ($11, $12, $13, $14, $15, $16, $17, $18, $19, $20) ON CONFLICT (tx_vin_vout_row_id, address, is_funding) DO UPDATE
	SET matching_tx_hash = $2, tx_hash = $3, tx_vin_vout_index = $4,
		block_time = $7, valid_mainchain = $9 RETURNING id;`},
		{"checked 1", args{1, true, false}, `WITH inserting AS (INSERT INTO addresses (address, matching_tx_hash, tx_hash,
		tx_vin_vout_index, tx_vin_vout_row_id, value, block_time, is_funding, valid_mainchain, tx_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) ON CONFLICT (tx_vin_vout_row_id, address, is_funding) DO NOTHING -- no lock on row
			RETURNING id
		)
		SELECT id FROM inserting
		UNION  ALL
		SELECT id FROM addresses
		WHERE  address = $1 AND is_funding = $8 AND tx_vin_vout_row_id = $5 -- only executed if no INSERT
		LIMIT  1;`},
		{"checked 2", args{2, true, false}, `WITH inserting AS (INSERT INTO addresses (address, matching_tx_hash, tx_hash,
		tx_vin_vout_index, tx_vin_vout_row_id, value, block_time, is_funding, valid_mainchain, tx_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10), ($11, $12, $13, $14, $15, $16, $17, $18, $19, $20) ON CONFLICT (tx_vin_vout_row_id, address, is_funding) DO NOTHING -- no lock on row
			RETURNING id
		)
		SELECT id FROM inserting
		UNION  ALL
		SELECT id FROM addresses
		WHERE  address = $1 AND is_funding = $8 AND tx_vin_vout_row_id = $5 -- only executed if no INSERT
		LIMIT  1;`},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if got := MakeAddressRowMultilineInsertStatement(tt.args.N, tt.args.checked, tt.args.updateOnConflict); got != tt.want {
				t.Errorf(`MakeAddressRowMultilineInsertStatement() = "%v" want "%v"`, got, tt.want)
			}
		})
	}
}
