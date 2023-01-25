-- +goose Up
-- +goose StatementBegin
create table `accounts` (
    `id` bigint unsigned primary key,
    `number` bigint unsigned,
    `name` varchar(64)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `accounts`;
-- +goose StatementEnd
