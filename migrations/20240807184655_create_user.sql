-- +goose Up
-- +goose StatementBegin
create table if not exists bootcamp.users
(
    id uuid primary key default gen_random_uuid(),
    password varchar(1000) not null,
    email varchar(1000) not null unique,
    type bootcamp.user_type not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists bootcamp.users;
-- +goose StatementEnd
