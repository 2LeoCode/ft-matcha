CREATE TABLE IF NOT EXISTS administrator (
  id INTEGER PRIMARY KEY,
  username TEXT NOT NULL,
  password TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS i_administrator_username_password
  ON administrator (username, password);
