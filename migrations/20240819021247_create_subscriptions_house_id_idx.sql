-- +goose Up
-- +goose StatementBegin
create index subscriptions_house_id_idx on bootcamp.subscriptions using btree (house_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index if exists subscriptions_house_id_idx;
-- +goose StatementEnd
