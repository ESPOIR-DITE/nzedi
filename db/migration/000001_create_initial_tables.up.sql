CREATE TABLE IF NOT EXISTS "account"
(
    "id"         varchar PRIMARY KEY,
    "date"       timestamptz,
    "password"   varchar,
    "email"      varchar     NOT NULL,
    "token"      varchar,
    "username"   varchar,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS "account_state"
(
    "id"          varchar PRIMARY KEY,
    "name"        varchar,
    "description" varchar,
    "created_at"  timestamptz NOT NULL,
    "updated_at"  timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS "user_type"
(
    "id"          varchar PRIMARY KEY,
    "name"        varchar,
    "description" varchar,
    "created_at"  timestamptz NOT NULL,
    "updated_at"  timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS "account_type"
(
    "id"            varchar PRIMARY KEY,
    "account_id"    varchar     NOT NULL,
    "account_state" varchar     NOT NULL,
    "user_type_id"  varchar     NOT NULL,
    "created_at"    timestamptz NOT NULL,
    "updated_at"    timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS "company"
(
    "id"         varchar PRIMARY KEY,
    "manager"    varchar     NOT NULL UNIQUE,
    "name"       varchar     NOT NULL,
    "url"        varchar,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS "user"
(
    "id"            varchar PRIMARY KEY,
    "account_id"    varchar     NOT NULL unique,
    "date_of_birth" timestamptz,
    "firstname"     varchar     NOT NULL,
    "lastname"      varchar     NOT NULL,
    "created_at"    timestamptz NOT NULL,
    "updated_at"    timestamptz NOT NULL
);

ALTER TABLE "account_type"
    add FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "account_type"
    add FOREIGN KEY ("account_state") REFERENCES "account_state" ("id");

ALTER TABLE "account_type"
    add FOREIGN KEY ("user_type_id") REFERENCES "user_type" ("id");

ALTER TABLE "user"
    add FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "company"
    add FOREIGN KEY ("manager") REFERENCES "account" ("id");
