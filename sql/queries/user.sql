-- name: User_Create :exec
INSERT INTO user (
  first_name,
  last_name,
  mail,
  password,
  birth_date,
  sex,
  orientation,
  country,
  city,
  bio
) VALUES (
  @first_name,
  @last_name,
  @mail,
  @password,
  @birth_date,
  @sex,
  @orientation,
  @country,
  @city,
  @bio
);

-- name: User_GetOneById :one
SELECT
  user.*,
-- noqa: disable=all
  sqlc.embed(age_gap),
  sqlc.embed(fame_rating_gap),
  sqlc.embed(last_position),
-- noqa: enable=all
FROM
  user
  LEFT JOIN age_gap ON user.id = age_gap.user_id
  LEFT JOIN fame_rating_gap ON user.id = fame_rating_gap.user_id
  LEFT JOIN last_position ON user.id = last_position.user_id
WHERE
  user.id = @id;

-- name: User_GetOneByCredentials :one
SELECT
  user.*,
-- noqa: disable=all
  sqlc.embed(age_gap),
  sqlc.embed(fame_rating_gap),
  sqlc.embed(last_position)
-- noqa: enable=all
FROM
  user
  LEFT JOIN age_gap ON user.id = age_gap.user_id
  LEFT JOIN fame_rating_gap ON user.id = fame_rating_gap.user_id
  LEFT JOIN last_position ON user.id = last_position.user_id
WHERE
  user.mail = @mail AND user.password = @password;

-- name: User_GetManyByDiscussion :many
SELECT
  user.*
FROM
  j_user_discussion
  INNER JOIN user ON j_user_discussion.user_id = user.id
WHERE
  j_user_discussion.discussion_id = @discussion_id
ORDER BY
  user.first_name ASC,
  user.last_name ASC;

-- name: User_Update :exec
UPDATE user SET
  mail = coalesce(@mail, mail),
  password = coalesce(@password, password),
  first_name = coalesce(@first_name, first_name),
  last_name = coalesce(@last_name, last_name),
  birth_date = coalesce(@birth_date, birth_date),
  sex = coalesce(@sex, sex),
  orientation = coalesce(@orientation, orientation),
  country = coalesce(@country, country),
  city = coalesce(@city, city),
  bio = coalesce(@bio, bio),
  suspended = coalesce(@suspended, suspended)
WHERE id = @id;

-- name: User_Delete :exec
DELETE FROM user
WHERE id = @id;
