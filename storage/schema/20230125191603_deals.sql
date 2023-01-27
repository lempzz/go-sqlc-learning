-- +goose Up
-- +goose StatementBegin
create table `deals` (
    `id` bigint unsigned primary key,
    `account_id` bigint unsigned not null,
    `type` tinyint unsigned not null comment "0 - BUY, 1 - SELL",
    `open_price` decimal(20, 12) unsigned not null,
    `close_price` decimal(20, 12) unsigned not null,
    `open_time` datetime not null,
    `close_time` datetime not null,
    `profit` decimal(20, 12) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `deals`;
-- +goose StatementEnd
