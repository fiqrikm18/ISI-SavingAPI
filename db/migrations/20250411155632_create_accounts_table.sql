-- +goose Up
-- +goose StatementBegin
create table accounts (
    id serial primary key,
    account_number varchar(20) not null,
    name varchar(255) not null,
    nik varchar(20) not null,
    phone varchar(15) not null,
    balance numeric(10,2) not null,
    created_at timestamp default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table accounts;
-- +goose StatementEnd
