IF NOT EXISTS CREATE TABLE accounts(
    account_id bigint primary key,
    org_id varchar(100) not null,
    limits json not null,
    balances json not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    status varchar(100) not null default "active",
    version bigint not null default 1
);

IF NOT EXISTS CREATE TABLE entries(
    tracking_id varchar(36) not null primary key,
    account_id bigint not null,
    org_id varchar(100) not null,
    impacts json not null,
    created_at timestamp default current_timestamp
);
