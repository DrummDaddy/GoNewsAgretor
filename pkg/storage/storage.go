package storage

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// База данных.
type DB struct {
	pool *pgxpool.Pool
}

// Публикация, получаемая из RSS.
type Post struct {
	ID      int    // номер записи
	Title   string // заголовок публикации
	Content string // содержимое публикации
	PubTime int64  // время публикации
	Link    string // ссылка на источник

}

func New() (*DB, error) {
	connstr := os.Getenv("newsdb")
	if connstr == "" {
		return nil, errors.New("не указано подключение к БД")
	}
	pool, err := pgxpool.Connect(context.Background(), connstr)
	if err != nil {
		return nil, err
	}
	db := DB{
		pool: pool,
	}
	return &db, nil
}

func (db *DB) StoreNews(news []Post) error {
	for _, post := range news {
		_, err := db.pool.Exec(context.Background(),
			`INSERT INTO news(title, content, pub_time, link) VALUES ($1, $2, $3, $4)`,
			post.Title,
			post.Content,
			post.PubTime,
			post.Link,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewsWithPagination возвращает новости с поддержкой фильтрации и пагинации
func (db *DB) NewsWithPagination(page, pageSize int, titleFilter string) ([]Post, error) {
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	var rows pgx.Rows
	var err error

	query := `SELECT id, title, content, pub_time, link FROM news 
              WHERE title ILIKE '%' || $1 || '%' 
              ORDER BY pub_time DESC 
              LIMIT $2 OFFSET $3`
	rows, err = db.pool.Query(context.Background(), query, titleFilter, pageSize, offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var news []Post
	for rows.Next() {
		var p Post
		err = rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.PubTime,
			&p.Link,
		)
		if err != nil {
			return nil, err
		}
		news = append(news, p)
	}
	return news, rows.Err()
}
