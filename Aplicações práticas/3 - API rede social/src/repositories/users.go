package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.User) (uint64, error) {
	statement, error := repository.db.Prepare("insert into users (name, nick, email, password) values (?, ?, ?, ?)")
	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if error != nil {
		return 0, error
	}

	lastID, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(lastID), nil
}

func (repository Users) Find(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, error := repository.db.Query(
		"select id, name, nick, email, created_at from users where name like ? or nick like ?",
		nameOrNick, nameOrNick,
	)
	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if error = lines.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); error != nil {
			return nil, error
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository Users) FindByID(userID uint64) (models.User, error) {
	lines, error := repository.db.Query("select id, name, nick, email, created_at from users where id = ?", userID)
	if error != nil {
		return models.User{}, error
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if error = lines.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); error != nil {
			return models.User{}, error
		}
	}

	return user, nil
}

func (repository Users) Update(userID uint64, user models.User) error {
	statement, error := repository.db.Prepare("update users set name = ?, nick = ?, email = ? where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(user.Name, user.Nick, user.Email, userID); error != nil {
		return error
	}

	return nil
}

func (repository Users) Delete(userID uint64) error {
	statement, error := repository.db.Prepare("delete from users where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(userID); error != nil {
		return error
	}

	return nil
}

func (repository Users) FindByEmail(email string) (models.User, error) {
	lines, error := repository.db.Query("select id, password from users where email = ?", email)
	if error != nil {
		return models.User{}, error
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if error = lines.Scan(&user.ID, &user.Password); error != nil {
			return models.User{}, error
		}
	}

	return user, nil
}
