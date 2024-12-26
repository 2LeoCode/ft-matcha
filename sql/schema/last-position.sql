CREATE TABLE IF NOT EXISTS last_position (
  user_id INTEGER UNIQUE NOT NULL REFERENCES user (id),
  lattitude INTEGER NOT NULL,
  longitude INTEGER NOT NULL
);
