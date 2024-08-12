-- +goose Up
-- +goose StatementBegin
create schema if not exists bootcamp;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop schema if exists bootcamp;
-- +goose StatementEnd
