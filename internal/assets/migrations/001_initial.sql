-- +migrate Up
create table blobs (
    id varchar primary key not null,
    blob jsonb
);
-- +migrate Down
drop table blobs;
