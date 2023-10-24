#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "todo" <<-EOSQL
  CREATE SCHEMA users;

  CREATE TABLE users.users
  (
      id         text NOT NULL,
      name       text NOT NULL,
      email      text NOT NULL,
      created_at timestamptz NOT NULL DEFAULT NOW(),
      updated_at timestamptz NOT NULL DEFAULT NOW(),
      PRIMARY KEY (id)
  );

  CREATE TRIGGER created_at_users_trgr BEFORE UPDATE ON users.users FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
  CREATE TRIGGER updated_at_users_trgr BEFORE UPDATE ON users.users FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

  GRANT USAGE ON SCHEMA users TO todo_user;
  GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA users TO todo_user;
EOSQL
