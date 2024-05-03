package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Posts struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

func (repository Posts) Create(post models.Post) (uint64, error) {
	statement, error := repository.db.Prepare("insert into posts (title, content, author_id) values (?, ?, ?)")
	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(post.Title, post.Content, post.AuthorID)
	if error != nil {
		return 0, error
	}

	lastID, error := result.LastInsertId()

	if error != nil {
		return 0, error
	}

	return uint64(lastID), nil

}

func (repository Posts) FindById(postId uint64) (models.Post, error) {
	rows, error := repository.db.Query(`
        select p.*, u.nick from posts p
        join users u on u.id = p.author_id
        where p.id = ?
    `, postId)
	if error != nil {
		return models.Post{}, error
	}
	defer rows.Close()

	var post models.Post
	if rows.Next() {
		fmt.Println(rows.Columns())
		if error = rows.Scan(
			&post.ID,
			&post.AuthorID,
			&post.Title,
			&post.Content,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); error != nil {
			return models.Post{}, error
		}
	}

	return post, nil

}

func (repository Posts) FindAll(userId uint64) ([]models.Post, error) {
	rows, error := repository.db.Query(`
        select distinct p.*, u.nick from posts p
        join users u on u.id = p.author_id
        left join followers f on p.author_id = f.user_id
        where u.id = ? or f.follower_id = ?
        order by 1 desc
    `, userId, userId)

	if error != nil {
		return nil, error
	}

	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if error = rows.Scan(
			&post.ID,
			&post.AuthorID,
			&post.Title,
			&post.Content,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); error != nil {
			return nil, error
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (repository Posts) Update(postId uint64, post models.Post) error {
	statement, error := repository.db.Prepare("update posts set title = ?, content = ? where id = ?")
	if error != nil {
		return error
	}

	defer statement.Close()

	if _, error = statement.Exec(post.Title, post.Content, postId); error != nil {
		return error
	}

	return nil
}

func (repository Posts) Delete(postId uint64) error {
	statement, error := repository.db.Prepare("delete from posts where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(postId); error != nil {
		return error
	}

	return nil
}

func (repository Posts) Like(postId uint64) error {
	statement, error := repository.db.Prepare("update posts set likes = likes + 1 where id = ?")
	if error != nil {
		return error
	}

	defer statement.Close()

	if _, error = statement.Exec(postId); error != nil {

		return error
	}

	return nil
}

func (repository Posts) Unlike(postId uint64) error {
	statement, error := repository.db.Prepare("update posts set likes = (case when likes > 0 then likes - 1 else 0 end) where id = ?")
	if error != nil {
		return error
	}

	defer statement.Close()

	if _, error = statement.Exec(postId); error != nil {
		return error
	}

	return nil
}

func (repository Posts) FindPostByUserId(userId uint64) ([]models.Post, error) {
	rows, error := repository.db.Query(`
        select p.*, u.nick from posts p
        join users u on u.id = p.author_id
        where p.author_id = ?
        order by 1 desc
    `, userId)
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if error = rows.Scan(
			&post.ID,
			&post.AuthorID,
			&post.Title,
			&post.Content,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); error != nil {
			return nil, error
		}
		posts = append(posts, post)
	}

	return posts, nil
}
