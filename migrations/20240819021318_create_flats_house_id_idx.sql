-- +goose Up
-- +goose StatementBegin
create index flats_house_id_idx on bootcamp.flats using btree (house_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index if exists flats_house_id_idx;
-- +goose StatementEnd
