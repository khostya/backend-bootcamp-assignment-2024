-- +goose Up
-- +goose StatementBegin
create type bootcamp.flat_status as enum ('created', 'approved', 'declined', 'on moderation');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop type bootcamp.flat_status;
-- +goose StatementEnd