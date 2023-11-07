CREATE TABLE "payment"."users" (
    "username" varchar PRIMARY KEY,
    "hashed_password" varchar NOT NULL,
    "full_name" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "payment"."accounts" ADD FOREIGN KEY ("owner") REFERENCES "payment"."users" ("username");

-- CREATE UNIQUE INDEX ON "payment"."accounts" ("owner", "currency");
ALTER TABLE "payment"."accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency")
