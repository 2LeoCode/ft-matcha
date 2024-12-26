CREATE TABLE IF NOT EXISTS user (
  id INTEGER PRIMARY KEY,
  mail TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  birth_date INTEGER NOT NULL,
  sex INTEGER NOT NULL CHECK (sex IN (0, 1)),
  orientation INTEGER NOT NULL CHECK (orientation IN (0, 1, 2)) DEFAULT 0,
  country INTEGER NOT NULL,
  city INTEGER NOT NULL,
  bio TEXT NOT NULL,
  suspended INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX IF NOT EXISTS i_user_mail_password ON user (mail, password);
CREATE INDEX IF NOT EXISTS i_user_age ON user (age);
CREATE INDEX IF NOT EXISTS i_user_last_position ON user (last_position);
CREATE INDEX IF NOT EXISTS i_user_orientation_sex ON user (orientation, sex);
CREATE INDEX IF NOT EXISTS i_user_country_city ON user (country, city);
