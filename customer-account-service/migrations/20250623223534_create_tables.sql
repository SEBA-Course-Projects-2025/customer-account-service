-- +goose Up
-- +goose StatementBegin
CREATE TABLE customer_accounts
(
    id               uuid primary key,
    email            varchar(255) not null unique,
    phone            varchar(255) not null unique,
    password_hash    varchar(255) not null,
    name             varchar(255) not null,
    shipping_address varchar(255) not null,
    created_at       timestamp default now(),
    updated_at       timestamp default now()
);

CREATE TABLE orders
(
    id          uuid primary key,
    customer_id uuid           not null references customer_accounts (id) on delete cascade,
    total_price numeric(12, 2) not null,
    status      varchar(40)    not null,
    created_at  timestamp default now(),
    updated_at  timestamp default now()
);

CREATE TABLE order_items
(
    id           uuid primary key not null,
    product_id   uuid             not null,
    order_id     uuid references orders (id),
    product_name varchar(255)     not null,
    quantity     int              not null,
    unit_price   numeric(12, 2)   not null,
    image_url    text             not null,
    created_at   timestamp default now(),
    updated_at   timestamp default now()
);

CREATE INDEX idx_orders_created_at ON orders (created_at);
CREATE INDEX idx_orders_total_price ON orders (total_price);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_orders_created_at;
DROP INDEX IF EXISTS idx_orders_total_price;

DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS customer_accounts;
-- +goose StatementEnd
