// Package database
package database

import (
	"github.com/google/uuid"
)

type User struct {
	ID 				uuid.UUID `json:"id"`
	FirstName string `json:"firstName"`
	LastName 	string `json:"lastName"`
	Biography string `json:"biography"`
}


var db = make(map[uuid.UUID]User, 10)

func FindAll () []User {
	users := make([]User, 0, len(db))
	for _, user := range db {
		users = append(users, user)
	}

	return users
}

func FindByID(id uuid.UUID) (User, bool){
	user, ok := db[id]
	return user, ok
}

func Insert(user User) (uuid.UUID, User) {
	id := uuid.New()
	user.ID = id
	db[id] = user
	
	return id, user
}

func Update(id uuid.UUID, user User) (User, bool) {
	if _, exists := db[id]; !exists {
		return User{}, false
	}
	user.ID = id
	db[id] = user
	return user, true
}

func Delete(id uuid.UUID) bool {
	if _, exists := db[id]; !exists {
		return false
	}

	delete(db, id)
	return true
}
