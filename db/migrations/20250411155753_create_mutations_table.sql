-- +goose Up
-- +goose StatementBegin
create table mutations (
    id serial primary key,
    account_id integer not null references accounts(id),
    type varchar(10) not null,
    amount numeric(10,2) not null,
    created_at timestamp default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table mutations;
-- +goose StatementEnd
