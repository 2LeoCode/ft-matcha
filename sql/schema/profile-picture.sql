CREATE TABLE IF NOT EXISTS profile_picture (
  id INTEGER PRIMARY KEY,
  user_id INTEGER NOT NULL REFERENCES user (id),
  file_name TEXT NOT NULL,
  data BLOB NOT NULL
);

CREATE INDEX IF NOT EXISTS i_profile_picture_user_id
ON profile_picture (user_id);
