CREATE TRIGGER IF NOT EXISTS tr_delete_empty_discussions
AFTER DELETE ON users_discussions
BEGIN
  DELETE FROM discussions
  WHERE
    id = old.discussion_id
    AND NOT EXISTS (
      SELECT 1
      FROM
        users_discussions
      WHERE
        users_discussions.discussion_id = old.discussion_id
    );
END;

CREATE TRIGGER IF NOT EXISTS tr_delete_unused_tags
AFTER DELETE ON users_tags
BEGIN
  DELETE FROM tags
  WHERE
    id = old.tag_id
    AND NOT EXISTS (
      SELECT 1
      FROM
        users_tags
      WHERE
        users_tags.tag_id = old.tag_id
    );
END;
