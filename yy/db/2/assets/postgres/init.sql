CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
   name VARCHAR(255) NOT NULL,
    last_login TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS posts
(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    likes INT DEFAULT 0,
    created_at TIMESTAMP

);

CREATE TABLE IF NOT EXISTS user_likes
(
    user_id INT,
    post_id INT
);

ALTER TABLE user_likes ADD CONSTRAINT user_likes_unique
    UNIQUE (user_id, post_id);

ALTER TABLE user_likes ADD CONSTRAINT user_likes_pk
    PRIMARY KEY (user_id, post_id);
