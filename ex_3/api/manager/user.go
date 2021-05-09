package manager

import (
	"errors"
	"ex_3/api/model"
	"fmt"
	"github.com/rs/zerolog/log"
	"strconv"
)

func (m *Manager) GetUser(id string) (*model.User, error) {
	rows, err := m.Search(fmt.Sprintf("SELECT * FROM users WHERE id=%v LIMIT 1", id))
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Err(err).Msg("error while closing previous statement")
		}
	}()

	if rows.Next() {
		user := new(model.User)
		if err := rows.Scan(&user.Id, &user.Email, &user.Username, &user.Password); err != nil {
			return nil, err
		}
		return user, nil
	}

	return nil, errors.New("something went wrong :(")
}

func (m *Manager) CreateUser(user *model.User) (*model.User, error) {
	id, err := m.Query(fmt.Sprintf("INSERT INTO users (email, username, password) values ('%v', '%v', '%v')", user.Email, user.Username, user.Password))
	if err != nil {
		return nil, err
	}

	user, err = m.GetUser(strconv.FormatInt(id, 10))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *Manager) UpdateUser(user *model.User) (*model.User, error) {

	prevUser, err := m.GetUser(strconv.FormatInt(user.Id, 10))
	if err != nil {
		return nil, err
	}

	if user.Email == "" {
		user.Email = prevUser.Email
	}

	if user.Username == "" {
		user.Username = prevUser.Username
	}

	if user.Password == "" {
		user.Password = prevUser.Password
	}

	emailParam := fmt.Sprintf("email='" + user.Email + "'")
	usernameParam := fmt.Sprintf("username='" + user.Username + "'")
	passwordParam := fmt.Sprintf("password='" + user.Password + "'")

	_, err = m.Query(fmt.Sprintf("UPDATE users SET %v, %v, %v WHERE id=%v", emailParam, usernameParam, passwordParam, user.Id))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *Manager) DeleteUser(user *model.User) (*model.User, error) {
	user, err := m.GetUser(strconv.FormatInt(user.Id, 10))
	if err != nil {
		return nil, err
	}

	_, err = m.Query(fmt.Sprintf("DELETE FROM users WHERE id=%v", user.Id))
	if err != nil {
		return nil, err
	}

	return user, nil
}
