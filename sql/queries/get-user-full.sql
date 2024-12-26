-- name: GetUserFull :one
SELECT
  user.*,
-- noqa: disable=all
  sqlc.embed(age_gap),
  sqlc.embed(fame_rating_gap),
  sqlc.embed(last_position)
-- noqa: enable=all
FROM
  user
  INNER JOIN age_gap ON user.id = age_gap.user_id
  INNER JOIN fame_rating_gap ON user.id = fame_rating_gap.user_id
  INNER JOIN last_position ON user.id = last_position.user_id
WHERE
  user.mail = ? AND user.password = ?;
