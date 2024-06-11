package userrepo

import (
	"fmt"
	"mymodule/yy/db/2/internal/entities"
	"mymodule/yy/db/2/internal/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetAll() ([]entities.User, error)
}

type userRepoImpl struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepoImpl {
	return &userRepoImpl{
		db: db,
	}
}

// default way to select and scan rows
func (r *userRepoImpl) GetAll() ([]*entities.User, error) {
	result := make([]*entities.User, 0)

	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("falied to get users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		user := new(entities.User)
		if err := rows.Scan(&user.ID, &user.Name, &user.LastLogin); err != nil {
			return nil, fmt.Errorf("failed to scan result: %w", err)
		}
		result = append(result, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// uses sqlx.DB.Select()
func (r *userRepoImpl) GetAll2() ([]entities.User, error) {
	users := make([]entities.User, 0)
	if err := r.db.Select(&users, "SELECT * FROM users"); err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}

	return users, nil
}

func (r *userRepoImpl) GetAllUsersWithPosts() ([]models.User, error) {
	rawUsers := make([]entities.UserWithPosts, 0)

	err := r.db.Select(&rawUsers, `
		SELECT users.id AS user_id , users.name AS user_name,
	users.last_login AS user_last_login, posts.id AS post_id, 
	posts.likes AS post_likes, posts.created_at AS post_created_at
		FROM users
			LEFT JOIN posts ON users.id = posts.user_id
		ORDER BY users.last_login DESC, posts.created_at DESC
	`)

	if err != nil {
		return nil, fmt.Errorf("failed to get users with posts: %w", err)
	}

	return models.NewUsers(rawUsers), nil
}

// uses default logic for IN operand
func (r *userRepoImpl) GetUsersByIDs(ids []uint64) ([]models.User, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	args := make([]interface{}, 0, len(ids))
	inParams := make([]string, 0, len(ids))

	for i, id := range ids {
		args = append(args, id)
		inParams = append(inParams, fmt.Sprintf("$%d", i+1))
	}

	query := fmt.Sprintf(`
		SELECT users.id AS user_id , users.name AS user_name,
	users.last_login AS user_last_login, posts.id AS post_id, 
	posts.likes AS post_likes, posts.created_at AS post_created_at
		FROM users
			LEFT JOIN posts ON users.id = posts.user_id
		WHERE users.id IN (%s)
		ORDER BY users.last_login DESC, posts.created_at DESC
	`, strings.Join(inParams, ","))

	rawUsers := make([]entities.UserWithPosts, 0)

	if err := r.db.Select(&rawUsers, query, args...); err != nil {
		return nil, fmt.Errorf("falied to get users by ids:%w", err)
	}

	return models.NewUsers(rawUsers), nil
}

// uses sqlx.In
func (r *userRepoImpl) GetUsersByIDs2(ids []uint64) ([]models.User, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	query, args, err := sqlx.In(`
	SELECT users.id AS user_id , users.name AS user_name,
users.last_login AS user_last_login, posts.id AS post_id, 
posts.likes AS post_likes, posts.created_at AS post_created_at
	FROM users
		LEFT JOIN posts ON users.id = posts.user_id
	WHERE users.id IN (?)
	ORDER BY users.last_login DESC, posts.created_at DESC
`)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}

	query = r.db.Rebind(query)
	rawUsers := make([]entities.UserWithPosts, 0)
	if err := r.db.Select(&rawUsers, query, args...); err != nil {
		return nil, fmt.Errorf("falied to get users by ids:%w", err)
	}

	return models.NewUsers(rawUsers), nil
}

func (r *userRepoImpl) AddUser(user models.User) (*models.User, error) {
	var userID uint64
	err := r.db.Get(
		&userID,
		`
			INSERT INTO users (name, last_login)
			VALUES($1, $2)
			RETURNING id	
		`,
		user.Name,
		user.LastLogin,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to add user")
	}

	user.ID = userID
	if len(user.Posts) == 0 {
		return &user, nil
	}

	valuesQuery := make([]string, 0, len(user.Posts))
	valuesArgs := make([]interface{}, 0, len(user.Posts)*3)
	i := 1
	for _, post := range user.Posts {
		valuesQuery = append(valuesQuery, fmt.Sprintf("$%s, $%s, $%s", i, i+1, i+2))
		i += 3
		valuesArgs = append(valuesArgs, user.ID, post.Likes, post.CreatedAt)
	}

	query := fmt.Sprintf(`
		INSERT INTO posts (user_id, likes, created_at)
		VALUES %s
		RETURNING id, user_id, likes, created_at
	`, strings.Join(valuesQuery, ", "))

	posts := make([]entities.Post, 0, len(user.Posts))
	err = r.db.Select(&posts, query, valuesArgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to add post: %w", err)
	}

	user.Posts = models.NewPosts(posts)
	return &user, nil
}