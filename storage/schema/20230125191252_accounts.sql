-- +goose Up
-- +goose StatementBegin
create table `accounts` (
    `id` bigint unsigned primary key,
    `number` bigint unsigned not null,
    `name` varchar(64) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `accounts`;
-- +goose StatementEnd
