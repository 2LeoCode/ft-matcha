CREATE TABLE IF NOT EXISTS administrators (
  id INTEGER PRIMARY KEY,
  username TEXT NOT NULL,
  password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS discussions (
  id INTEGER PRIMARY KEY,
  creation_time INTEGER NOT NULL DEFAULT (unixepoch())
);


CREATE TABLE IF NOT EXISTS tags (
  id INTEGER PRIMARY KEY,
  name TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY,
  mail TEXT UNIQUE NOT NULL,
  password BLOB NOT NULL,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  birth_date INTEGER NOT NULL,
  sex INTEGER NOT NULL,
  orientation INTEGER NOT NULL DEFAULT 0,
  country TEXT NOT NULL,
  city TEXT NOT NULL,
  bio TEXT NOT NULL,
  fame_points INTEGER NOT NULL DEFAULT 0,
  suspended INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS interactions (
  kind TEXT NOT NULL CHECK (kind IN ('kiss', 'punch', 'block')),
  creation_time INTEGER NOT NULL DEFAULT (unixepoch()),
  giver_id INTEGER REFERENCES users (id) ON DELETE SET NULL,
  given_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS messages (
  id INTEGER PRIMARY KEY,
  creation_time INTEGER NOT NULL DEFAULT (unixepoch()),
  content TEXT NOT NULL,
  author_id INTEGER REFERENCES users (id) ON DELETE SET NULL,
  discussion_id INTEGER NOT NULL REFERENCES discussion (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS profile_pictures (
  id INTEGER PRIMARY KEY,
  user_id INTEGER NOT NULL REFERENCES users (id),
  file_name TEXT NOT NULL,
  data BLOB NOT NULL
);

CREATE TABLE IF NOT EXISTS reports (
  reason TEXT NOT NULL,
  status TEXT NOT NULL,
  creation_time INTEGER NOT NULL DEFAULT (unixepoch()),
  assigned_administrator_id INTEGER
    REFERENCES administrators (id) ON DELETE SET NULL,
  issuer_id INTEGER REFERENCES users (id) ON DELETE SET NULL,
  target_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS positions (
  user_id INTEGER UNIQUE NOT NULL REFERENCES user (id),
  latitude INTEGER NOT NULL,
  longitude INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS users_discussions (
  user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
  discussion_id INTEGER NOT NULL REFERENCES discussions (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS users_tags (
  user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
  tag_id INTEGER NOT NULL REFERENCES tags (id) ON DELETE CASCADE
);
