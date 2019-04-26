// Copyright (c) 2019, The Decred developers
// See LICENSE for details.

package internal

import "testing"

func BenchmarkMakeVoutMultilineInsertStatement_upsert3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MakeVoutMultilineInsertStatement(3, true, true)
	}
}

func TestMakeVoutMultilineInsertStatement(t *testing.T) {
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
		{"unchecked 1", args{1, false, false}, `INSERT INTO vouts (tx_hash, tx_index, tx_tree, value,
		version, pkscript, script_req_sigs, script_type, script_addresses) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id;`},
		{"unchecked 2", args{2, false, false}, `INSERT INTO vouts (tx_hash, tx_index, tx_tree, value,
		version, pkscript, script_req_sigs, script_type, script_addresses) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9), ($10, $11, $12, $13, $14, $15, $16, $17, $18) RETURNING id;`},
		{"unchecked 3", args{3, false, false}, `INSERT INTO vouts (tx_hash, tx_index, tx_tree, value,
		version, pkscript, script_req_sigs, script_type, script_addresses) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9), ($10, $11, $12, $13, $14, $15, $16, $17, $18), ($19, $20, $21, $22, $23, $24, $25, $26, $27) RETURNING id;`},
		{"upsert 1", args{1, true, true}, `INSERT INTO vouts (tx_hash, tx_index, tx_tree, value,
		version, pkscript, script_req_sigs, script_type, script_addresses) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ON CONFLICT (tx_hash, tx_index, tx_tree) DO UPDATE
	SET version = $5 RETURNING id;`},
		{"upsert 2", args{2, true, true}, `INSERT INTO vouts (tx_hash, tx_index, tx_tree, value,
		version, pkscript, script_req_sigs, script_type, script_addresses) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9), ($10, $11, $12, $13, $14, $15, $16, $17, $18) ON CONFLICT (tx_hash, tx_index, tx_tree) DO UPDATE
	SET version = $5 RETURNING id;`},
		{"checked 1", args{1, true, false}, `WITH inserting AS (INSERT INTO vouts (tx_hash, tx_index, tx_tree, value,
		version, pkscript, script_req_sigs, script_type, script_addresses) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ON CONFLICT (tx_hash, tx_index, tx_tree) DO NOTHING -- no lock on row
			RETURNING id
		)
		SELECT id FROM inserting
		UNION  ALL
		SELECT id FROM vouts
		WHERE  tx_hash = $1 AND tx_index = $2 AND tx_tree = $3 -- only executed if no INSERT
		LIMIT  1;`},
		{"checked 2", args{2, true, false}, `WITH inserting AS (INSERT INTO vouts (tx_hash, tx_index, tx_tree, value,
		version, pkscript, script_req_sigs, script_type, script_addresses) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9), ($10, $11, $12, $13, $14, $15, $16, $17, $18) ON CONFLICT (tx_hash, tx_index, tx_tree) DO NOTHING -- no lock on row
			RETURNING id
		)
		SELECT id FROM inserting
		UNION  ALL
		SELECT id FROM vouts
		WHERE  tx_hash = $1 AND tx_index = $2 AND tx_tree = $3 -- only executed if no INSERT
		LIMIT  1;`},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if got := MakeVoutMultilineInsertStatement(tt.args.N, tt.args.checked, tt.args.updateOnConflict); got != tt.want {
				t.Errorf(`MakeVoutMultilineInsertStatement() = "%v" want "%v"`, got, tt.want)
			}
		})
	}
}
func BenchmarkMakeVinMultilineInsertStatement_upsert3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MakeVinMultilineInsertStatement(3, true, true)
	}
}

func TestMakeVinMultilineInsertStatement(t *testing.T) {
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
		{"unchecked 1", args{1, false, false}, `INSERT INTO vins (tx_hash, tx_index, tx_tree, prev_tx_hash, prev_tx_index, prev_tx_tree,
		value_in, is_valid, is_mainchain, block_time, tx_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id;`},
		{"unchecked 2", args{2, false, false}, `INSERT INTO vins (tx_hash, tx_index, tx_tree, prev_tx_hash, prev_tx_index, prev_tx_tree,
		value_in, is_valid, is_mainchain, block_time, tx_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11), ($12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22) RETURNING id;`},
		{"unchecked 3", args{3, false, false}, `INSERT INTO vins (tx_hash, tx_index, tx_tree, prev_tx_hash, prev_tx_index, prev_tx_tree,
		value_in, is_valid, is_mainchain, block_time, tx_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11), ($12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22), ($23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33) RETURNING id;`},
		{"upsert 1", args{1, true, true}, `INSERT INTO vins (tx_hash, tx_index, tx_tree, prev_tx_hash, prev_tx_index, prev_tx_tree,
		value_in, is_valid, is_mainchain, block_time, tx_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) ON CONFLICT (tx_hash, tx_index, tx_tree) DO UPDATE
	SET is_valid = $8, is_mainchain = $9, block_time = $10,
		prev_tx_hash = $4, prev_tx_index = $5, prev_tx_tree = $6
	RETURNING id;`},
		{"upsert 2", args{2, true, true}, `INSERT INTO vins (tx_hash, tx_index, tx_tree, prev_tx_hash, prev_tx_index, prev_tx_tree,
		value_in, is_valid, is_mainchain, block_time, tx_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11), ($12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22) ON CONFLICT (tx_hash, tx_index, tx_tree) DO UPDATE
	SET is_valid = $8, is_mainchain = $9, block_time = $10,
		prev_tx_hash = $4, prev_tx_index = $5, prev_tx_tree = $6
	RETURNING id;`},
		{"checked 1", args{1, true, false}, `WITH inserting AS (INSERT INTO vins (tx_hash, tx_index, tx_tree, prev_tx_hash, prev_tx_index, prev_tx_tree,
		value_in, is_valid, is_mainchain, block_time, tx_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) ON CONFLICT (tx_hash, tx_index, tx_tree) DO NOTHING -- no lock on row
		RETURNING id
	)
	SELECT id FROM inserting
	UNION  ALL
	SELECT id FROM vins
	WHERE  tx_hash = $1 AND tx_index = $2 AND tx_tree = $3 -- only executed if no INSERT
	LIMIT  1;`},
		{"checked 2", args{2, true, false}, `WITH inserting AS (INSERT INTO vins (tx_hash, tx_index, tx_tree, prev_tx_hash, prev_tx_index, prev_tx_tree,
		value_in, is_valid, is_mainchain, block_time, tx_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11), ($12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22) ON CONFLICT (tx_hash, tx_index, tx_tree) DO NOTHING -- no lock on row
		RETURNING id
	)
	SELECT id FROM inserting
	UNION  ALL
	SELECT id FROM vins
	WHERE  tx_hash = $1 AND tx_index = $2 AND tx_tree = $3 -- only executed if no INSERT
	LIMIT  1;`},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if got := MakeVinMultilineInsertStatement(tt.args.N, tt.args.checked, tt.args.updateOnConflict); got != tt.want {
				t.Errorf(`MakeVinMultilineInsertStatement() = "%v" want "%v"`, got, tt.want)
			}
		})
	}
}
