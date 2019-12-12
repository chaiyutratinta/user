package services

import (
	"fmt"
	"log"
	"reflect"
	"user/models"
	"user/repositories"
	"user/utils"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type UserServices interface {
	CreateUser(*models.Body) (models.Token, error)
	ValidateUser(*models.Body) (models.Token, bool)
}

type user struct {
	repositories.DB
}

func New(repo repositories.DB) UserServices {

	return &user{
		repo,
	}
}

func (u *user) CreateUser(body *models.Body) (token models.Token, err error) {
	validateBody := &struct {
		UserName string `validate:"required"`
		Password string `validate:"required"`
		Email    string `validate:"required,email"`
	}{
		body.UserName,
		body.Password,
		body.Email,
	}
	validate := validator.New()

	if errors := validate.Struct(validateBody); errors != nil {
		errStr := "Bad request: "
		bodyT := reflect.TypeOf(*body)
		keyTags := map[string]string{}

		for i := 0; i < bodyT.NumField(); i++ {
			field := bodyT.Field(i)
			keyTags[field.Name] = field.Tag.Get("json")
		}

		for idx, errs := range errors.(validator.ValidationErrors) {
			separate := ""

			if idx > 0 {
				separate = ","
			}

			errStr = fmt.Sprintf("%s%s%s=%s", errStr, separate, keyTags[errs.Namespace()], errs.ActualTag())
		}

		return models.Token(""), fmt.Errorf(errStr)
	}

	userID := (uuid.New()).String()
	password := utils.Hash(body.Password)
	err = u.InsertUser(&models.User{
		ID:       userID,
		Email:    body.Email,
		Password: password,
		UserName: body.UserName,
	})

	if err != nil {
		token = models.Token("")

		return
	}

	token, err = utils.Token(userID, "all")

	return
}

func (u *user) ValidateUser(body *models.Body) (token models.Token, ok bool) {
	passSha1 := utils.Hash(body.Password)

	userID, userFound := u.LookUpUser(&models.User{
		UserName: body.UserName,
		Email:    body.Email,
		Password: passSha1,
	})

	if !userFound {

		return models.Token(""), false
	}

	token, err := utils.Token(userID, "all")

	if err != nil {
		log.Println(err)

		return models.Token(""), false
	}

	return token, true
}
