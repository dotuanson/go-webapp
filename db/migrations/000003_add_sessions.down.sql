ALTER TABLE IF EXISTS "payment"."sessions" DROP CONSTRAINT IF EXISTS "sessions_username_fkey";

DROP TABLE IF EXISTS "payment"."sessions";