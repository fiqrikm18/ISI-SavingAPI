-- +goose Up
-- +goose StatementBegin
alter table accounts add column updated_at timestamp null;
alter table accounts add column deleted_at timestamp null;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table accounts drop column updated_at;
alter table accounts drop column deleted_at;
-- +goose StatementEnd
