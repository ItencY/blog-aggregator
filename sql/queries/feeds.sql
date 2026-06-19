-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT feeds.name AS feed_name, feeds.url, users.name AS user_name
FROM feeds
INNER JOIN users
ON feeds.user_id = users.id;

-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT
    iff.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow iff
JOIN feeds ON iff.feed_id = feeds.id
JOIN users ON iff.user_id = users.id;

-- name: GetFeedFollowsForUser :many
SELECT
    ff.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM feed_follows ff
JOIN feeds ON ff.feed_id = feeds.id
JOIN users ON ff.user_id = users.id
WHERE ff.user_id = $1;

-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE feed_id = $1 AND user_id = $2;
