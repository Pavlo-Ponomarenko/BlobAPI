-- +migrate Up
create table blobs (
    id serial primary key not null,
    blob bytea
);
-- +migrate Down
drop table blobs;
