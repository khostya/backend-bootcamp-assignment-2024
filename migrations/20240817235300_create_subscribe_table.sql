-- +goose Up
-- +goose StatementBegin
create table if not exists bootcamp.subscriptions
(
    user_email varchar(1000) references bootcamp.users (email),
    house_id   int references bootcamp.houses (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table bootcamp.subscriptions;
-- +goose StatementEnd
