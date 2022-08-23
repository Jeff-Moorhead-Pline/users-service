package data

import (
	"fmt"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	Username  string    `json:"username"`
	LastLogin time.Time `json:"last_login"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Address   string    `json:"address"`
}

type UsersStore []User

// A fake datastore for user objects
var users = UsersStore{
	{
		Username:  "jmoorhead",
		LastLogin: time.Now(),
		Email:     "jmoorhead@performline.com",
		FullName:  "Jeff Moorhead",
		Address:   "1 Locust St, Madison, NJ",
	},
	{
		Username:  "kgonzalez",
		LastLogin: time.Date(2022, time.August, 21, 20, 25, 23, 0, time.Local),
		Email:     "kgonzalez@performline.com",
		FullName:  "Karen Gonzalez",
		Address:   "123 Some St, Basking Ridge, NJ",
	},
	{
		Username:  "jmacarthur",
		LastLogin: time.Date(2022, time.August, 9, 9, 14, 27, 0, time.Local),
		Email:     "jmacarthur@performline.com",
		FullName:  "Jeremy MacArthur",
		Address:   "987 Any Ave, Rockaway, NJ",
	},
}

func GetUsersDataLayer() UsersStore {

	return users
}

func FindUserDataLayer(username string) (*User, error) {

	for _, u := range users {
		if u.Username == username {
			return &u, nil
		}
	}

	return nil, echo.ErrNotFound
}

func AddUserDataLayer(u *User) error {

	for _, existing := range users {
		if strings.EqualFold(existing.Username, u.Username) {
			return fmt.Errorf("user %v already exists", u.Username)
		}
	}

	users = append(users, *u)

	return nil
}

func UpdateUserDataLayer(u *User, username string) error {

	var found bool

	for i, existing := range users {
		if strings.EqualFold(existing.Username, username) {
			found = true

			if !u.LastLogin.IsZero() {
				existing.LastLogin = u.LastLogin
			}

			if u.Email != "" {
				existing.Email = u.Email
			}

			if u.FullName != "" {
				existing.FullName = u.FullName
			}

			if u.Address != "" {
				existing.Address = u.Address
			}

			users[i] = existing
			break
		}
	}

	if !found {
		return fmt.Errorf("user not found: %v", username)
	}

	return nil
}
