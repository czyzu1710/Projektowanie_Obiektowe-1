package dao

import (
	"errors"
	"fmt"
	"log"
	"zad3/database"
	"zad3/model"
)

func GetUser(id string) (*model.User, error) {
	db := database.GetInstance()

	row, err := db.Query(fmt.Sprintf("SELECT * FROM UserS WHERE ID = %v LIMIT 1", id))
	if err != nil {
		log.Println("Error trying fetch User")
		return nil, err
	}

	if row.Next() {
		User := new(model.User)
		err := row.Scan(&User.Id, &User.Name, &User.Surname, &User.Birthday)
		if err != nil {
			return nil, err
		}
		return User, nil
	}

	return nil, errors.New("No such User.")
}

func GetUsers() ([]*model.User, error) {
	db := database.GetInstance()
	row, err := db.Query("SELECT * FROM UserS s")
	if err != nil {
		log.Println("Error trying fetch Users")
	}

	var Users = make([]*model.User, 0)
	for row.Next() {
		User := new(model.User)
		err := row.Scan(&User.Id, &User.Name, &User.Surname, &User.Birthday)
		if err != nil {
			return nil, err
		}
		Users = append(Users, User)
	}

	if len(Users) == 0 {
		return Users, errors.New("Error trying fetch Users or no Users are available")
	}

	return Users, nil
}

func Save(s *model.User) error {
	db := database.GetInstance()
	id, err := db.Execute(fmt.Sprintf("INSERT INTO UserS(NAME, SURNAME, BIRTHDAY) VALUES ('%v', '%v', '%v')", s.Name, s.Surname, s.Birthday))
	if err != nil {
		log.Println("Error trying insert User.")
		return err
	}
	s.Id = id
	return nil
}

func Update(s *model.User) error {
	db := database.GetInstance()
	_, err := db.Execute(fmt.Sprintf("UPDATE UserS SET NAME = '%v', SURNAME = '%v', BIRTHDAY = '%v' WHERE ID = %v", s.Name, s.Name, s.Birthday, s.Id))
	if err != nil {
		log.Println("Error trying update User.")
		return err
	}
	return nil
}

func Delete(id string) error {
	db := database.GetInstance()
	_, err := db.Execute(fmt.Sprintf("DELETE FROM UserS WHERE ID = %v", id))
	if err != nil {
		log.Println("Error trying to delete User")
		return err
	}
	return nil
}
