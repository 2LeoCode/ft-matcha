CREATE TABLE IF NOT EXISTS j_user_discussion (
  user_id INTEGER NOT NULL REFERENCES user (id) ON DELETE CASCADE,
  discussion_id INTEGER NOT NULL REFERENCES discussion (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS i_user_discussion_user_id
  ON j_user_discussion (user_id);
CREATE INDEX IF NOT EXISTS i_user_discussion_discussion_id
  ON j_user_discussion (discussion_id);
