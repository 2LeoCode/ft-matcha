CREATE TABLE IF NOT EXISTS j_discussion_message (
  discussion_id INTEGER NOT NULL REFERENCES discussion (id) ON DELETE CASCADE,
  message_id INTEGER NOT NULL REFERENCES message (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS i_discussion_message_discussion_id
  ON j_discussion_message (discussion_id);
CREATE INDEX IF NOT EXISTS i_discussion_message_message_id
  ON j_discussion_message (discussion_id);
