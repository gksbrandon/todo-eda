#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "todo" <<-EOSQL
  CREATE SCHEMA lists;

  CREATE TABLE lists.lists
  (
      id            text NOT NULL,
      user_id       text NOT NULL,
      created_at    timestamptz NOT NULL DEFAULT NOW(),
      updated_at    timestamptz NOT NULL DEFAULT NOW(),
      PRIMARY KEY (id)
  );

  CREATE INDEX list_lists_idx ON lists.lists (user_id);
  CREATE TRIGGER created_at_lists_trgr BEFORE UPDATE ON lists.lists FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
  CREATE TRIGGER updated_at_lists_trgr BEFORE UPDATE ON lists.lists FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

  CREATE TABLE lists.tasks
  (
      id            text NOT NULL,
      list_id       text NOT NULL,
      description   text NOT NULL,
      completed     bool NOT NULL DEFAULT FALSE,
      created_at    timestamptz NOT NULL DEFAULT NOW(),
      updated_at    timestamptz NOT NULL DEFAULT NOW(),
      PRIMARY KEY (id)
  );

  CREATE INDEX list_tasks_idx ON lists.tasks (list_id);
  CREATE TRIGGER created_at_tasks_trgr BEFORE UPDATE ON lists.tasks FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
  CREATE TRIGGER updated_at_tasks_trgr BEFORE UPDATE ON lists.tasks FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

  GRANT USAGE ON SCHEMA lists TO todo_user;
  GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA lists TO todo_user;
EOSQL
