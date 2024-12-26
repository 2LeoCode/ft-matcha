CREATE TABLE IF NOT EXISTS message (
  creation_time INTEGER NOT NULL DEFAULT (unixepoch()),
  content TEXT NOT NULL,
  author_id INTEGER REFERENCES user (id) ON DELETE SET NULL,
  discussion_id INTEGER NOT NULL REFERENCES discussion (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS i_message_author_id ON message (author_id);
CREATE INDEX IF NOT EXISTS i_message_discussion_id ON message (discussion_id);
