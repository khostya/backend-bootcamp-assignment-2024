-- +goose Up
-- +goose StatementBegin
create table if not exists bootcamp.houses
(
    id                 serial primary key not null,
    address            varchar(5000)      not null unique,
    year               int                not null,
    developer          varchar(2000),

    created_at         timestamptz        not null,
    last_flat_added_at timestamptz        not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table bootcamp.houses;
-- +goose StatementEnd
