-- +goose Up
-- +goose StatementBegin
create type bootcamp.user_type as enum ('moderator', 'client');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop type bootcamp.user_type;
-- +goose StatementEnd
