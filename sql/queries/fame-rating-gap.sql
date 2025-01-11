-- name: FameRatingGap_Set :one
INSERT INTO fame_rating_gap (user_id, min, max)
VALUES (@id, @min, @max)
ON CONFLICT (user_id) DO UPDATE SET
  min = excluded.min,
  max = excluded.max
RETURNING *;
