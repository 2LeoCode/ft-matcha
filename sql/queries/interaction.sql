-- name: Interaction_Create :one
INSERT INTO interaction (kind, giver_id, given_id)
VALUES (@kind, @giver_id, @given_id);

-- name: Interaction_GetManyByGiver :many
SELECT * FROM interaction
WHERE giver_id = @giver_id AND kind IN (sqlc.slice("kinds"))
GROUP BY kind
ORDER BY creation_time ASC;

-- name: Interaction_GetManyByGiven :many
SELECT * FROM interaction
WHERE given_id = @given_id AND kind IN (sqlc.slice("kinds"))
GROUP BY kind
ORDER BY creation_time ASC;
