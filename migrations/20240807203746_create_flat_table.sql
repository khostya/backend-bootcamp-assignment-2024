-- +goose Up
-- +goose StatementBegin
create table if not exists bootcamp.flats
(
    id           serial primary key   not null,
    house_id     int                  not null references bootcamp.houses (id),
    price        int                  not null,
    rooms        int                  not null,
    status       bootcamp.flat_status not null,
    moderator_id uuid,
    unique (id, house_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table bootcamp.flats;
-- +goose StatementEnd
