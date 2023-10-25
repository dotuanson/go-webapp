CREATE SCHEMA "payment";

CREATE TABLE "payment"."accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestampz NOT NULL DEFAULT (now())
);

CREATE TABLE "payment"."entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestampz NOT NULL DEFAULT (now())
);

CREATE TABLE "payment"."transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestampz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "payment"."accounts" ("owner");

CREATE INDEX ON "payment"."entries" ("account_id");

CREATE INDEX ON "payment"."transfers" ("from_account_id");

CREATE INDEX ON "payment"."transfers" ("to_account_id");

CREATE INDEX ON "payment"."transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "payment"."entries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "payment"."transfers"."amount" IS 'must be positive';

ALTER TABLE "payment"."entries" ADD FOREIGN KEY ("account_id") REFERENCES "payment"."accounts" ("id");

ALTER TABLE "payment"."transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "payment"."accounts" ("id");

ALTER TABLE "payment"."transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "payment"."accounts" ("id");
