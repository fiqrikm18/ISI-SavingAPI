-- +goose Up
-- +goose StatementBegin
alter table mutations add column updated_at timestamp null;
alter table mutations add column deleted_at timestamp null;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table mutations drop column updated_at;
alter table mutations drop column deleted_at;
-- +goose StatementEnd
