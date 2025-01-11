-- name: Message_Create :one
INSERT INTO message (author_id, discussion_id, content)
VALUES (@author_id, @discussion_id, @content)
RETURNING *;

-- name: Message_GetManyByDiscussion :many
SELECT * FROM message
WHERE discussion_id = @discussion_id
ORDER BY creation_time ASC;

-- name: Message_Delete :exec
DELETE FROM message
WHERE id = @id;
