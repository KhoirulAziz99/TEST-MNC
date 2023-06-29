-- data table merchant 
SELECT * FROM merchant;
SELECT * FROM customer;
SELECT * FROM history_trx;

UPDATE customer SET balance = 0 WHERE id = 2;

INSERT INTO merchant (name, no_telephon, category, balance, created_at, updated_at)
VALUES ('ACT', '01112', 'fashion', 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO merchant (name, no_telephon, category, balance, created_at, updated_at)
VALUES ('JatiPermata', '01113', 'furniture', 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO merchant (name, no_telephon, category, balance, created_at, updated_at)
VALUES ('BeautyQ', '01114', 'Make-up', 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO merchant (name, no_telephon, category, balance, created_at, updated_at)
VALUES ('GuySensey', '01115', 'gym-tools', 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO merchant (name, no_telephon, category, balance, created_at, updated_at)
VALUES ('Hashirama Store Pat', '01116', 'biju', 1000000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

ALTER TABLE merchant
ADD CONSTRAINT constraint_name UNIQUE (name);


SELECT * FROM history_trx;

create table if not exists history_trx (
    id serial not null primary key,
    customer_email varchar(15) not null,
    merchant_name varchar(100) not null,
    created_at timestamp,
    foreign key(customer_email) references customer(email),
    foreign key(merchant_name) references merchant(name)
);

