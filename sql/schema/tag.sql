CREATE TABLE IF NOT EXISTS tag (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS i_tag_name ON tag (name);

CREATE TRIGGER IF NOT EXISTS t_delete_unused_tag
AFTER DELETE ON user_tag
BEGIN
  DELETE FROM tag
  WHERE
    id = old.tag_id
    AND NOT EXISTS (
      SELECT 1
      FROM
        user_tag
      WHERE
        user_tag.tag_id = old.tag_id
    );
END;
