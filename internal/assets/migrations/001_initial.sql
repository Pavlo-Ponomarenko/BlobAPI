-- +migrate Up
create table blobs (
    id serial primary key,
    blob bytea
);
-- +migrate Down
drop table blobs;
