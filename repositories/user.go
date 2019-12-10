package repositories

import (
	"fmt"
	"log"
	"strings"
	"user/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB interface {
	InsertUser(*models.User) error
	LookUpUser(*models.User) (string, bool)
}

type dataBase struct {
	*gorm.DB
}

func New() DB {
	db, err := gorm.Open("mysql", "admin:nimda@(localhost:3306)/users")

	if err != nil {
		log.Fatal(err)
	}

	return &dataBase{
		db,
	}
}

func (db *dataBase) InsertUser(user *models.User) (err error) {
	if ok := db.NewRecord(user); !ok {
		err = fmt.Errorf("create error")

		return
	}

	if errs := db.Create(user).GetErrors(); len(errs) > 0 {
		duplicate := "Duplicate: "

		for idx, elm := range errs {
			dup := strings.Split(elm.Error(), " for key ")
			key := strings.ReplaceAll(dup[1], "'", "")
			separate := ""

			if idx > 0 {
				separate = ","
			}

			duplicate = fmt.Sprintf("%s%s%s", duplicate, key, separate)
		}

		return fmt.Errorf(duplicate)
	}

	if ok := db.NewRecord(user); !ok {
		err = fmt.Errorf("create error")

		return
	}

	return nil
}

func (db *dataBase) LookUpUser(user *models.User) (id string, ok bool) {
	out := &models.User{}
	db.Find(out, user)

	if out.ID != "" {

		return out.ID, true
	}

	return "", false
}
