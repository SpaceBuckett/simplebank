
CREATE TYPE "Currency" AS ENUM (
  'USD',
  'EUR'
);

CREATE TABLE "account" (
                           "id" bigserial PRIMARY KEY,
                           "owner" varchar NOT NULL,
                           "balance" bigint NOT NULL,
                           "currency" "Currency" NOT NULL,
                           "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
                           "id" bigserial PRIMARY KEY,
                           "account_id" bigint NOT NULL,
                           "amount" bigint NOT NULL,
                           "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
                             "id" bigserial PRIMARY KEY,
                             "from_account_id" bigint NOT NULL,
                             "to_account_id" bigint NOT NULL,
                             "amount" bigint NOT NULL,
                             "created_at" timestamptz NOT NULL DEFAULT (now())
);
CREATE TABLE "users" (
                         "username" varchar PRIMARY KEY,
                         "hashed_password" varchar NOT NULL,
                         "full_name" varchar NOT NULL,
                         "email" varchar UNIQUE NOT NULL,
                         "password_changed_at" timestamptz NOT NULL DEFAULT (now()),
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "account" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "account" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency");


CREATE INDEX ON "account" ("owner");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative and positive';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "account" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "account" ("id");

--   'USD',
--   'EUR'
-- );
--
-- CREATE TABLE "account" ("id" bigserial PRIMARY KEY, "owner" varchar NOT NULL, "balance" bigint NOT NULL, "currency" varchar NOT NULL, "created_at" timestamptz NOT NULL DEFAULT (now()));
--
-- CREATE TABLE "entereis" (
--   "id" bigserial PRIMARY KEY,
--   "account_id" bigint NOT NULL,
--   "amount" bigint NOT NULL,
--   "created_at" timestamptz NOT NULL DEFAULT (now())
-- );
--
-- CREATE TABLE "transfers" (
--   "id" bigserial PRIMARY KEY,
--   "from_account_id" bigint NOT NULL,
--   "to_account_id" bigint NOT NULL,
--   "amount" bigint NOT NULL,
--   "created_at" timestamptz NOT NULL DEFAULT (now())
-- );
--
-- CREATE INDEX ON "account" ("owner");
--
-- CREATE INDEX ON "entereis" ("account_id");
--
-- CREATE INDEX ON "transfers" ("from_account_id");
--
-- CREATE INDEX ON "transfers" ("to_account_id");
--
-- CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");
--
-- COMMENT ON COLUMN "entereis"."amount" IS 'can be negative and positive';
--
-- COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';
--
-- ALTER TABLE "entereis" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");
--
-- ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "account" ("id");
--
-- ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "account" ("id");
