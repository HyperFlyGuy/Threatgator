-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

-- name: GetPostsForUser :many
SELECT posts.*, feed_follows.user_id AS user_id, feeds.name as feed_name
FROM posts
INNER JOIN feed_follows ON posts.feed_id  = feed_follows.feed_id
INNER JOIN feeds ON posts.feed_id = feeds.id
WHERE feed_follows.user_id = $1
ORDER BY published_at DESC
LIMIT $2;
