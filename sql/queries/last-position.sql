-- name: LastPosition_Set :one
INSERT INTO last_position (user_id, latitude, longitude)
VALUES (@id, @latitude, @longitude)
ON CONFLICT (user_id) DO UPDATE SET
  latitude = excluded.latitude,
  longitude = excluded.longitude
RETURNING *;
