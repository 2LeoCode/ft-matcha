CREATE TABLE IF NOT EXISTS discussion (
  id INTEGER PRIMARY KEY,
  creation_time INTEGER NOT NULL DEFAULT (unixepoch())
);

CREATE TRIGGER IF NOT EXISTS t_delete_empty_discussions
AFTER DELETE ON user_discussion
BEGIN
  DELETE FROM discussion
  WHERE
    id = old.discussion_id
    AND NOT EXISTS (
      SELECT 1
      FROM
        user_discussion
      WHERE
        user_discussion.discussion_id = old.discussion_id
    );
END;
