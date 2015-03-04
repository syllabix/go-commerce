package models

import (
	//"golang.org/x/crypto/bcrypt"
	"errors"
	"regexp"
)

/* Admin Users */

type Administrator struct {
	Username  string
	FirstName string
	LastName  string
	Email     email
}

type email struct {
	address string
}

var validEmail = regexp.MustCompile("^[\\w.!#$%&’*+\\-/=?\\^`{|}~]+@[a-zA-Z0-9-]+(\\.[a-zA-Z0-9-]+)*$")

func (e *email) SetAddress(email string) error {
	if validEmail.MatchString(email) {
		e.address = email
		return nil
	}
	return errors.New("Email address is not valid")
}
