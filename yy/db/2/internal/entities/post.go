package entities

import "time"

type Post struct {
	ID        uint64    `db:"id"`
	UserID    uint64    `db:"user_id"`
	Likes     int64     `db:"likes"`
	CreatedAt time.Time `db:"created_at"`
}
