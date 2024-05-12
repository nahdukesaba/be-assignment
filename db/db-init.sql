CREATE USER postgres;
CREATE DATABASE postgres;
GRANT ALL PRIVILEGES ON DATABASE postgres TO postgres;

CREATE TABLE "user" (
    user_id integer NOT NULL,
    "username" varchar(100) NOT NULL,
    "password" varchar(255) NOT NULL,
    PRIMARY KEY(user_id)
);


CREATE TABLE account (
    account_id varchar(100) NOT NULL,
    user_id integer NOT NULL REFERENCES "user"(user_id),
    account_type varchar(100) NOT NULL,
    balance integer NOT NULL,
    "limit" integer NOT NULL,
    updated_at timestamp NOT NULL,
    PRIMARY KEY(account_id)
);

CREATE TABLE transaction (
    transaction_id integer NOT NULL,
    account_id varchar(100) NOT NULL REFERENCES account(account_id),
    amount integer NOT NULL,
    reciever varchar(100) NULL REFERENCES account(account_id),
    "status" varchar(100) NOT NULL,
    created_at timestamp NOT NULL,
    PRIMARY KEY(transaction_id)
);
