-- name: CreatePost :one
INSERT INTO posts (created_at, updated_at, title, url, description, published_at, feed_id) VALUES (
    $1, $1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetPostsForUser :many
WITH user_feeds AS (
    SELECT feed_id FROM feed_follows WHERE user_id = $1
)
SELECT posts.id, posts.created_at, posts.updated_at, posts.title, posts.url, posts.description, posts.published_at, posts.feed_id FROM posts, user_feeds
WHERE posts.feed_id = user_feeds.feed_id ORDER BY posts.published_at DESC LIMIT $2;