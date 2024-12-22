package mysqldb

import (
	"balances/internal/app/domain/accounts"
	"balances/internal/worker/domain/daily"
	"balances/pkg/logger"
	"context"
	"database/sql"
	"time"
)

var (
	INSERT_BALANCE            = `insert into daily_balances (date, account_id, org_id, balances, version) values (?, ?, ?, ?, ?)`
	UPDATE_BALANCE            = `update daily_balances set balances = ?, version = ? where date = ? and  account_id = ? and org_id = ?`
	SELECT_LAST_BALANCE       = `select db.date, db.account_id, db.org_id, db.balances, db.version from daily_balances db WHERE db.account_id = ? and db.org_id = ? order by db.date DESC limit 1`
	SELECT_BALANCES_BY_PERIOD = `select db.date, db.account_id, db.org_id, db.balances, db.version from daily_balances db where db.account_id = ? and db.org_id = ? and db.date BETWEEN ? and ? order by db.date`
)

type Repository struct {
	ctx context.Context
	db  *sql.DB
}

func NewRepository(ctx context.Context) *Repository {
	db, err := Connect()
	if err != nil {
		logger.Error(ctx, "error on connect to database", logger.Fields{"error": err.Error()})
		return &Repository{ctx: ctx}
	}

	return &Repository{ctx: ctx, db: db}
}

func (r Repository) Close() {
	r.db.Close()
}

func (r Repository) SaveNewBalance(ctx context.Context, b daily.Balance) error {
	return nil
}
func (r Repository) UpdateExistingBalance(ctx context.Context, b daily.Balance) error {
	return nil
}
func (r Repository) RetrieveLastBalance(ctx context.Context, accountID int64, orgID string) (daily.Balance, error) {
	return daily.Balance{}, nil

}
func (r Repository) RetrieveBalanceByPeriod(ctx context.Context, accountID int64, orgID string, initialDate, finalDate time.Time) ([]daily.Balance, error) {
	return nil, nil
}

func (r Repository) SaveNewAccount(ctx context.Context, a accounts.Account) error {
	acc := buildAccountTable(a)
	statement, err := r.db.Prepare(INSERT_ACCOUNT)
	if err != nil {
		logger.Error(ctx, "error on save new account", logger.Fields{"account": a, "error": err.Error(), "sql": INSERT_ACCOUNT})
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(
		acc.AccountID,
		acc.OrgID,
		acc.Limits,
		acc.Balances,
		acc.CreatedAt,
		acc.UpdatedAt,
		acc.Status,
		acc.Version,
	)
	if err != nil {
		return err
	}
	return nil
}
func (r Repository) UpdateExistingAccount(ctx context.Context, a accounts.Account) error {
	acc := buildAccountTable(a)
	statement, err := r.db.Prepare(UPDATE_ACCOUNT)
	if err != nil {
		logger.Error(ctx, "error on update existing account", logger.Fields{"account": a, "error": err.Error(), "sql": UPDATE_ACCOUNT})
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(
		acc.Limits,
		acc.Balances,
		acc.UpdatedAt,
		acc.Status,
		acc.Version,
		acc.AccountID,
		acc.OrgID,
	)
	if err != nil {
		return err
	}
	return nil
}
func (r Repository) RetrieveAccountByID(ctx context.Context, accountID int64, orgID string) (accounts.Account, error) {
	var acc Account
	err := r.db.QueryRow(SELECT_ACCOUNT, accountID, orgID).Scan(
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
		logger.Error(ctx, "error on retrieve existing account", logger.Fields{"account": accountID, "error": err.Error(), "sql": SELECT_ACCOUNT})
		return accounts.Account{}, err
	}
	return acc.toEntity(), nil
}

func (r Repository) SaveEntryAndUpdateAccount(ctx context.Context, e accounts.Entry, a accounts.Account) error {
	tx, err := r.db.BeginTx(r.ctx, nil)
	if err != nil {
		logger.Error(ctx, "error on start new transaction", logger.Fields{"account": a, "entry": e, "error": err.Error()})
		return err
	}
	defer tx.Rollback()

	stat, err := tx.Prepare(INSERT_ENTRIES)
	if err != nil {
		logger.Error(ctx, "error on save new entry", logger.Fields{"account": a, "entry": e, "error": err.Error(), "sql": INSERT_ENTRIES})
		return err
	}
	defer stat.Close()

	en, err := buildEntriesTable(e)
	if err != nil {
		logger.Error(ctx, "error on parse to table", logger.Fields{"entry": e, "error": err})
		return err
	}
	_, err = stat.Exec(
		en.TrackingID,
		en.AccountID,
		en.OrgID,
		en.Impacts,
		en.CreatedAt,
	)
	if err != nil {
		return err
	}

	acc := buildAccountTable(a)
	stat, err = tx.Prepare(UPDATE_ACCOUNT)
	if err != nil {
		logger.Error(ctx, "error on prepare statement", logger.Fields{"account": a, "error": err.Error(), "sql": UPDATE_ACCOUNT})
		return err
	}
	_, err = stat.Exec(
		acc.Limits,
		acc.Balances,
		acc.UpdatedAt,
		acc.Status,
		acc.Version,
		acc.AccountID,
		acc.OrgID,
	)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}
