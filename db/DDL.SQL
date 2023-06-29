create table if not exists customer (
    id serial not null primary key,
    name varchar(100) not null,
    email varchar(100) unique not null,
    password varchar(100) unique not null,
    balance integer,
    created_at timestamp,
    updated_at timestamp,
    is_deleted bool default false
);

create table if not exists merchant (
    id serial not null primary key,
    name varchar(100) unique not null,
    no_telephon varchar(15) unique not null,
    category varchar(100) not null,
    balance integer,
    created_at timestamp,
    updated_at timestamp,
    active_status bool default true
);

create table if not exists history_log (
    id serial not null primary key,
    customer_id integer not null,
    created_at timestamp,
    foreign key(customer_id) references customer(id)
);

create table if not exists history_trx (
    id serial not null primary key,
    customer_email integer not null,
    merchant_name integer not null,
    created_at timestamp,
    foreign key(customer_email) references customer(email),
    foreign key(merchant_name) references merchant(name)
);