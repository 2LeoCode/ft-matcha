-- name: Discussion_Create :one
INSERT INTO discussion () VALUES () RETURNING id;

-- name: Discussion_GetOneById :one
SELECT * FROM discussion WHERE id = @id;

-- name: Discussion_AddUser :exec
INSERT INTO j_user_discussion (discussion_id, user_id)
  VALUES (@id, @user_id);

-- name: Discussion_RemoveUser :exec
DELETE FROM j_user_discussion
  WHERE discussion_id = @id AND user_id = @user_id;

-- name: Discussion_GetManyByUser :many
SELECT
  discussion.*
FROM
  j_user_discussion
  INNER JOIN discussion ON j_user_discussion.discussion_id = discussion.id
WHERE
  j_user_discussion.user_id = @user_id
ORDER BY
  discussion.creation_time DESC;
