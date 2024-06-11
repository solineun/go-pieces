package models

import (
	"mymodule/yy/db/2/internal/entities"
	"time"
)

type User struct {
	ID        uint64
	Name      string
	LastLogin time.Time
	Posts     []Post
}

type Post struct {
	ID        uint64
	UserID    uint64
	Likes     int64
	CreatedAt time.Time
}

func NewUsers(rawUsers []entities.UserWithPosts) []User {
	usersMap := make(map[uint64]User, 0)
	for _, rawUser := range rawUsers {
		user, ok := usersMap[rawUser.ID]
		if !ok {
			user = User{
				ID:        rawUser.ID,
				Name:      rawUser.Name,
				LastLogin: rawUser.LastLogin,
				Posts:     []Post{},
			}
		}

		if rawUser.PostID == nil {
			continue
		}

		user.Posts = append(user.Posts, Post{
			ID:        *rawUser.PostID,
			UserID:    rawUser.ID,
			Likes:     *rawUser.PostLikes,
			CreatedAt: *rawUser.PostCreatedAt,
		})

		usersMap[rawUser.ID] = user
	}

	users := make([]User, 0, len(usersMap))
	for _, user := range usersMap {
		users = append(users, user)
	}

	return users
}

func NewPosts(rawPosts []entities.Post) []Post {
	posts := make([]Post, 0, len(rawPosts))
	for _, p := range rawPosts {
		posts = append(posts, Post{
			ID: p.ID,
			UserID: p.UserID,
			Likes: p.Likes,
			CreatedAt: p.CreatedAt,
		})
	}

	return posts
}