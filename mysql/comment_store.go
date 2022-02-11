package mysql

import (
	"fmt"

	"github.com/gabrielcerdam-olxautos/goreddit"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CommentStore struct {
	*sqlx.DB
}

func (s *CommentStore) Comment(id uuid.UUID) (goreddit.Comment, error) {
	var t goreddit.Comment
	if err := s.Get(&t, `SELECT * FROM comments WHERE id = $1`, id); err != nil {
		return goreddit.Comment{}, fmt.Errorf("error getting comment: %w", err)
	}
	return t, nil
}

func (s *CommentStore) Comments() ([]goreddit.Comment, error) {
	var tt []goreddit.Comment
	if err := s.Select(&tt, `SELECT * FROM comments`); err != nil {
		return []goreddit.Comment{}, fmt.Errorf("error getting comments: %w", err)
	}
	return tt, nil
}

func (s *CommentStore) CreateComment(c *goreddit.Comment) error {
	if err := s.Get(c, `INSERT INTO comments VALUE ($1, $2, $3, $4) RETURNING *`,
		c.PostID,
		c.Content,
		c.ID); err != nil {
		return fmt.Errorf("error creating comment %w", err)
	}
	return nil
}

func (s *CommentStore) UpdateComment(c *goreddit.Comment) error {
	if err := s.Get(c, `UPDATE comment SET post_id = $1, content = $2, votes = $3 WHERE id = $4) RETURNING *`,
		c.PostID,
		c.Content,
		c.Votes,
		c.ID); err != nil {
		return fmt.Errorf("error updating comment %w", err)
	}
	return nil
}

func (s *CommentStore) DeleteComment(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM comments WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting comment: %w", err)
	}
	return nil
}
