package domain

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"user/models"
	"user/services"
)

type User interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type user struct {
	services.UserServices
}

func New(serv services.UserServices) (u User) {
	u = &user{
		serv,
	}

	return
}

func (u *user) Register(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	body := &models.Body{}

	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		defer req.Body.Close()

		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	token, err := u.CreateUser(body)

	if err != nil {
		switch {
		case strings.HasPrefix(err.Error(), "Duplicate: "):
			{
				errDuplicate := map[string]string{}
				errKey := strings.TrimPrefix(err.Error(), "Duplicate: ")

				for _, key := range strings.Split(errKey, ",") {
					errDuplicate[strings.Trim(key, " ")] = "duplicate."
				}

				json, _ := json.Marshal(errDuplicate)
				writer.WriteHeader(http.StatusBadRequest)
				writer.Write(json)

				return
			}
		case strings.HasPrefix(err.Error(), "Bad request: "):
			{
				required := map[string]string{}
				errKey := strings.TrimPrefix(err.Error(), "Bad request: ")

				for _, key := range strings.Split(errKey, ",") {
					required[strings.Trim(key, " ")] = "required"
				}

				json, _ := json.Marshal(required)
				writer.WriteHeader(http.StatusBadRequest)
				writer.Write(json)

				return
			}
		default:
			{
				log.Println(err)
				writer.WriteHeader(http.StatusBadRequest)

				return
			}
		}

	}

	res, _ := json.Marshal(map[string]models.Token{
		"token": token,
	})
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)

	return
}

func (u *user) Login(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	body := &models.Body{}

	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		defer req.Body.Close()

		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	token, ok := u.ValidateUser(body)

	if !ok {
		json, _ := json.Marshal(map[string]string{
			"error": "incorrect.",
		})
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(json)

		return
	}

	json, _ := json.Marshal(map[string]models.Token{
		"token": token,
	})
	writer.WriteHeader(http.StatusOK)
	writer.Write(json)

	return

}
