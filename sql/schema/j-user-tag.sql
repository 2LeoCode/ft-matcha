CREATE TABLE IF NOT EXISTS j_user_tag (
  user_id INTEGER NOT NULL REFERENCES user (id) ON DELETE CASCADE,
  tag_id INTEGER NOT NULL REFERENCES tag (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS i_user_tag_user_id ON j_user_tag (user_id);
CREATE INDEX IF NOT EXISTS i_user_tag_tag_id ON j_user_tag (tag_id);
