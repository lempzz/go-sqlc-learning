-- +goose Up
-- +goose StatementBegin
create table `deals` (
    `id` bigint unsigned primary key,
    `account_id` bigint unsigned,
    `type` tinyint unsigned comment "0 - BUY, 1 - SELL",
    `open_price` decimal(20, 12) unsigned,
    `close_price` decimal(20, 12) unsigned,
    `open_time` datetime,
    `close_time` datetime,
    `profit` decimal(20, 12)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `deals`;
-- +goose StatementEnd
