package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	u "github.com/duartqx/ddcomments/domains/entities/user"
	m "github.com/duartqx/ddcomments/domains/models"
)

type UserRepository struct {
	db *sqlx.DB
}

func GetNewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur UserRepository) GetUserModel() *u.UserEntity {
	return &u.UserEntity{}
}

func (ur UserRepository) GetDTOModel() *u.UserDTO {
	return &u.UserDTO{}
}

func (ur UserRepository) FindByID(id uuid.UUID) (m.User, error) {

	user := ur.GetUserModel()

	if err := ur.db.Get(user, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return nil, err
	}

	return user, nil
}

func (ur UserRepository) FindByEmail(email string) (m.User, error) {

	user := ur.GetUserModel()

	if err := ur.db.Get(user, "SELECT * FROM users WHERE email = $1", email); err != nil {
		return nil, err
	}

	return user, nil

}

func (ur UserRepository) Update(user m.User) error {

	if user.GetId() == uuid.Nil {
		return fmt.Errorf("Invalid user id")
	}

	if _, err := ur.db.Exec(
		"UPDATE users SET email = $1, name = $2 WHERE id = $3",
		user.GetEmail(), user.GetName(), user.GetId(),
	); err != nil {
		return err
	}

	return nil
}

func (ur UserRepository) All() (*[]m.User, error) {
	users := []m.User{}

	rows, err := ur.db.Query("SELECT * FROM users ORDER BY email ASC")
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		user := ur.GetDTOModel()

		if err := rows.Scan(&user); err != nil {
			return nil, err
		}

		var iUser m.User = user

		users = append(users, iUser)
	}

	return &users, nil
}

func (ur UserRepository) Create(user m.User) error {

	if user.GetEmail() == "" || user.GetName() == "" || user.GetPassword() == "" {
		return fmt.Errorf("Invalid User")
	}

	var id uuid.UUID

	if err := ur.db.QueryRow(
		`
			INSERT INTO users (email, name, password)
			VALUES ($1, $2, $3)
			RETURNING id
		`, user.GetEmail(), user.GetName(), user.GetPassword(),
	).Scan(&id); err != nil {
		return err
	}

	user.SetId(id)

	return nil
}

func (ur UserRepository) Delete(user m.User) error {

	if user.GetId() == uuid.Nil {
		return fmt.Errorf("Invalid User")
	}

	if _, err := ur.db.Exec("DELETE FROM users WHERE id = $1", user.GetId()); err != nil {
		return err
	}

	return nil
}

func (ur UserRepository) ExistsById(id uuid.UUID) (exists *bool) {
	ur.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM users WHERE id = $1)",
		id,
	).Scan(&exists)

	return exists
}

func (ur UserRepository) ExistsByEmail(email string) (exists *bool) {
	ur.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)",
		email,
	).Scan(&exists)

	return exists
}
