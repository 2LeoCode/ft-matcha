CREATE TABLE IF NOT EXISTS user (
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
  suspended INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX IF NOT EXISTS i_user_mail_password ON user (mail, password);

CREATE INDEX IF NOT EXISTS i_user_birth_date ON user (birth_date);

CREATE INDEX IF NOT EXISTS i_user_last_position ON user (last_position);

CREATE INDEX IF NOT EXISTS i_user_orientation_sex ON user (orientation, sex);

CREATE INDEX IF NOT EXISTS i_user_country_city ON user (country, city);
