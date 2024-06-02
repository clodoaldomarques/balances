CREATE TABLE accounts(
    account_id bigint primary key,
    org_id varchar(100) not null,
    limits json not null,
    availables json not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    allow_credit tinyint(1) not null default 1,
    allow_debit tinyint(1) not null default 1,
    version bigint not null default 1
);

CREATE TABLE entries (
    tracking_id varchar(36) not null primary key,
    account_id bigint not null,
    org_id varchar(100) not null,
    impacts json not null,
    considers json not null,
    availables json not null,
    amount numeric(15,2),
    amount_consider numeric(15,2),
    operation varchar(50),
    created_at timestamp default current_timestamp,
    version bigint not null
);
