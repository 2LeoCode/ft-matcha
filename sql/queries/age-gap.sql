-- name: AgeGap_Set :one
INSERT INTO age_gap (user_id, min, max)
VALUES (@user_id, @min, @max)
ON CONFLICT (user_id) DO UPDATE SET
  min = excluded.min,
  max = excluded.max
RETURNING *;
