package mysql

import (
	"fmt"

	"github.com/gabrielcerdam-olxautos/goreddit"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PostStore struct {
	*sqlx.DB
}

func (s *PostStore) Post(id uuid.UUID) (goreddit.Post, error) {
	var t goreddit.Post
	if err := s.Get(&t, `SELECT * FROM posts WHERE id = $1`, id); err != nil {
		return goreddit.Post{}, fmt.Errorf("error getting post: %w", err)
	}
	return t, nil
}

func (s *PostStore) Posts() ([]goreddit.Post, error) {
	var tt []goreddit.Post
	if err := s.Select(&tt, `SELECT * FROM posts`); err != nil {
		return []goreddit.Post{}, fmt.Errorf("error getting posts %w", err)
	}
	return tt, nil
}

func (s *PostStore) CreatePost(p *goreddit.Post) error {
	if err := s.Get(p, `INSERT INTO posts VALUE ($1, $2, $3, $4, $5) RETURNING *`,
		p.ID,
		p.ThreadID,
		p.Title,
		p.Content,
		p.Votes); err != nil {
		return fmt.Errorf("error creating post %w", err)
	}
	return nil
}

func (s *PostStore) UpdatePost(p *goreddit.Post) error {
	if err := s.Get(p, `UPDATE posts SET thread_id = $1, title = $2 content = $3,votes = $4 WHERE id = $5 RETURNING *`,
		p.ThreadID,
		p.ThreadID,
		p.Content,
		p.Votes,
		p.ID); err != nil {
		return fmt.Errorf("error updating post %w", err)
	}
	return nil
}

func (s *PostStore) DeletePost(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM posts WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting post: %w", err)
	}
	return nil
}
