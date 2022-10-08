package user

import "time"

type User struct {
	// The unique identifier of the user
	Id int `json:"id"`
	// The unique or not unique username of the user
	Username string `json:"username"`
	// A generic name for the user
	Name string `json:"name"`
	// Optional last name
	Surname string `json:"surname"`
	// The email address of the user, we use this to confirm their account
	Email string `json:"email"`
	// A phone number to associate the account, maybe use it for 2-factor logins
	Phone string `json:"phone"`
	// The creation date of this user
	CreateDate time.Time `json:"createDate"`
	// Last time an edit was applied to this user, useful for update polling
	LastModified time.Time `json:"lastModified"`
	// Wether or not the user has verified their account
	Enabled bool `json:"enabled"`
	// Wether or not this user account has been deleted
	IsDeleted bool `json:"isDeleted"`
}
