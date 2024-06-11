package postrepo

import (
	"fmt"
	"mymodule/yy/db/2/internal/entities"
	"mymodule/yy/db/2/internal/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

type PostRepository interface{}

type postRepoImpl struct {
	db *sqlx.DB
}

func (r *postRepoImpl) AddPosts(posts []models.Post) ([]models.Post, error) {
	createdPosts := make([]models.Post, 0, len(posts))
	const batch = 100
	for i := 0; i < len(posts); i = i + batch {
		endOfBatch := i + batch
		if endOfBatch > len(posts) {
			endOfBatch = len(posts)
		}

		postsBatch, err := r.addPosts(posts[i:endOfBatch])
		if err != nil {
			return nil, fmt.Errorf("failed to batch insert posts: %w", err)
		}

		createdPosts = append(createdPosts, postsBatch...)
	}

	return createdPosts, nil
}

func (r *postRepoImpl) addPosts(posts []models.Post) ([]models.Post, error) {
	if len(posts) == 0 {
		return nil, nil
	}

	valuesQuery := make([]string, 0, len(posts))
	valuesArgs := make([]interface{}, 0, len(posts)*3)
	i := 1
	for _, post := range posts {
		valuesQuery = append(valuesQuery, fmt.Sprintf("$%s, $%s, $%s", i, i+1, i+2))
		i += 3
		valuesArgs = append(valuesArgs, post.UserID, post.Likes, post.CreatedAt)
	}

	query := fmt.Sprintf(`
		INSERT INTO posts (user_id, likes, created_at)
		VALUES %s
		RETURNING id, user_id, likes, created_at
	`, strings.Join(valuesQuery, ", "))

	createdEntities := make([]entities.Post, 0, len(posts))

	if err := r.db.Select(&createdEntities, query, valuesArgs...); err != nil {
		return nil, fmt.Errorf("failed to add post: %w", err)
	}

	return models.NewPosts(createdEntities), nil
}

func (r *postRepoImpl) AddLike(postID, userID uint64) (bool, error) {
	res, err := r.db.Exec(`
		INSERT INTO user_likes (user_id, post_id)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`, userID, postID)
	if err != nil {
		return false, fmt.Errorf("failed to insert user like: %w", err)
	}
	rowsAff, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("failed to get rows affected by insert: %w", err)
	}
	if rowsAff == 0 {
		return false, nil
	}

	res, err = r.db.Exec(`
		UPDATE posts SET likes = likes + 1
		WHERE posts.id = $1
	`, postID)
	if err != nil {
		return false, fmt.Errorf("failed to update likes: %w", err)
	}
	rowsAff, err = res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("faled to get affected rows: %w", err)
	}
	if rowsAff == 0 {
		return false, nil
	}

	return true, nil
}

func (r *postRepoImpl) Delete(postID uint64) (bool, error) {
	res, err := r.db.Exec(`
		DELETE FROM posts
		WHERE posts.id = $1
	`, postID)
	if err != nil {
		return false, fmt.Errorf("failed to delete post: %w", err)
	}
	rowsAff, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("falied to get rows affected by delete: %w", err)
	}
	if rowsAff == 0 {
		return false, nil
	}

	return true, nil
}

func (r *postRepoImpl) DeleteByIDs(ids []uint64) error {
	if len(ids) == 0 {
		return nil
	}

	var err error
	const batch = 100
	for i := 0; i < batch; i = i + batch {
		endOfBatch := i + batch
		if endOfBatch > len(ids) {
			endOfBatch = len(ids)
		}

		err = r.deleteByIds(ids[i:endOfBatch])
		if err != nil {
			return fmt.Errorf("failed to delete posts: %w", err)
		}
	}

	return nil
}

func (r *postRepoImpl) deleteByIds(ids []uint64) error {
	q, args, err := sqlx.In(`
			DELETE FROM posts
			WHERE posts.id
			IN (?)`,
		ids)
	if err != nil {
		return fmt.Errorf("falied to prepare query: %w", err)
	}
	q = r.db.Rebind(q)
	if _, err = r.db.Exec(q, args...); err != nil {
		return fmt.Errorf("falied to execute query: %w", err)
	}

	return nil
}

// uses transaction instead of default query
func (r *postRepoImpl) AddLikeTx(postID, userID uint64) (done bool, err error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return false, fmt.Errorf("failed to begin tx: %w", err)
	}

	defer func() {
		if err != nil {
			errRb := tx.Rollback()
			if errRb != nil {
				err = fmt.Errorf("failed to rollback tx: %w", errRb)
				return
			}
			return
		}
		err = tx.Commit()
	}()

	if _, err = tx.Exec("SET TRANSACTION ISOLATION LEVEL REPEATABLE READ;"); err != nil {
		return false, fmt.Errorf("failed to set tx iso level: %w", err)
	}

	return r.addLikeTx(tx, postID, userID)
}

func (r *postRepoImpl) addLikeTx(tx *sqlx.Tx, postID uint64, userID uint64) (bool, error) {
	res, err := tx.Exec(`
		INSERT INTO user_likes (user_id, post_id)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`, userID, postID)
	if err != nil {
		return false, fmt.Errorf("failed to execute insert tx: %w", err)
	}
	rowsAff, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("failed to get rows affected by insert tx: %w", err)
	}
	if rowsAff == 0 {
		return false, nil
	}

	res, err = tx.Exec(`
		UPDATE posts SET likes = likes + 1
		WHERE posts.id = $1
	`, postID)
	if err != nil {
		return false, fmt.Errorf("failed to exec update tx: %w", err)
	}
	rowsAff, err = res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("failed to get rows affected by update tx: %w", err)
	}
	if rowsAff == 0 {
		return false, nil
	}


	return true, nil
}
