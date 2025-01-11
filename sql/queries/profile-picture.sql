-- name: ProfilePicture_Create :one
INSERT INTO profile_picture (user_id, file_name, data)
VALUES (@user_id, @file_name, @data) RETURNING *;

-- name: ProfilePicture_GetManyByUser :many
SELECT * FROM profile_picture
WHERE user_id = @user_id;

-- name: ProfilePicture_Delete :exec
DELETE FROM profile_picture
WHERE id = @id;
