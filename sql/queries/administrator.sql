-- name: Administrator_Create :exec
INSERT INTO administrator (username, password)
  VALUES (@username, @password);

-- name: Administrator_GetOneById :one
SELECT * FROM administrator WHERE id = @id;

-- name: Administrator_GetOneByCredentials :one
SELECT * FROM administrator WHERE username = @username AND password = @password;
