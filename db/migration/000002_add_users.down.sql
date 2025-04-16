ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS "fk_owner_currency";
ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";
DROP TABLE IF EXISTS "users";
