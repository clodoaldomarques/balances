package mysqldb

import (
	"balances/internal/app/domain/accounts"
	"context"
	"database/sql"
)

var (
	INSERT_ACCOUNT = `insert into accounts (account_id, org_id, limits, balances, created_at, updated_at, status, version) values (?, ?, ?, ?, ?, ?, ?, ?)`
	SELECT_ACCOUNT = `select account_id, org_id, limits, balances, created_at, updated_at, status, version from accounts where account_id = ?`
	UPDATE_ACCOUNT = `update accounts set limits = ?, balances = ?, updated_at = ?, status = ?, version = ? where account_id = ? and org_id = ?`
	INSERT_ENTRIES = `insert into entries (tracking_id, account_id, org_id, impacts, created_at) values (?, ?, ?, ?, ?)`
)

type Repository struct {
	ctx context.Context
	db  *sql.DB
}

func NewRepository(ctx context.Context) *Repository {
	db, err := Connect()
	if err != nil {
		panic(err)
	}

	return &Repository{ctx: ctx, db: db}
}

func (r Repository) Close() {
	r.db.Close()
}

func (r Repository) SaveNewAccount(ctx context.Context, a accounts.Account) error {
	accTable := AccountToTable(a)
	statement, err := r.db.Prepare(INSERT_ACCOUNT)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(
		accTable.AccountID,
		accTable.OrgID,
		accTable.Limits,
		accTable.Balances,
		accTable.CreatedAt,
		accTable.UpdatedAt,
		accTable.Status,
		accTable.Version,
	)
	if err != nil {
		return err
	}
	return nil
}
func (r Repository) UpdateExistingAccount(ctx context.Context, a accounts.Account) error {
	accTable := AccountToTable(a)
	statement, err := r.db.Prepare(UPDATE_ACCOUNT)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(
		accTable.Limits,
		accTable.Balances,
		accTable.UpdatedAt,
		accTable.Status,
		accTable.Version,
		accTable.AccountID,
		accTable.OrgID,
	)
	if err != nil {
		return err
	}
	return nil
}
func (r Repository) RetrieveAccountByID(ctx context.Context, accountID int64) (accounts.Account, error) {
	var acc AccountTable
	err := r.db.QueryRow(SELECT_ACCOUNT, accountID).Scan(
		&acc.AccountID,
		&acc.OrgID,
		&acc.Limits,
		&acc.Balances,
		&acc.CreatedAt,
		&acc.UpdatedAt,
		&acc.Status,
		&acc.Version,
	)
	if err != nil {
		return accounts.Account{}, err
	}
	return acc.ToEntity(), nil
}

func (r Repository) SaveEntryAndUpdateAccount(ctx context.Context, e accounts.Entry, a accounts.Account) error {
	aTable := AccountToTable(a)
	eTable, err := EntryToTable(e)
	if err != nil {
		return err
	}

	tx, err := r.db.BeginTx(r.ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stat, err := tx.Prepare(INSERT_ENTRIES)
	if err != nil {
		return err
	}
	defer stat.Close()
	_, err = stat.Exec(
		eTable.TrackingID,
		eTable.AccountID,
		eTable.OrgID,
		eTable.Impacts,
		eTable.CreatedAt,
	)
	if err != nil {
		return err
	}

	stat, err = tx.Prepare(UPDATE_ACCOUNT)
	if err != nil {
		return err
	}
	_, err = stat.Exec(
		aTable.Limits,
		aTable.Balances,
		aTable.UpdatedAt,
		aTable.Status,
		aTable.Version,
		aTable.AccountID,
		aTable.OrgID,
	)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}
