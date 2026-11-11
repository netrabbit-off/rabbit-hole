package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/netrabbit-off/rabbit-hole/internal/repository"
	"github.com/netrabbit-off/rabbit-hole/pkg/helpers"
	"github.com/netrabbit-off/rabbit-hole/pkg/models"
)

type Controller struct {
	userRepo  *repository.UserRepository
	tokenRepo *repository.TokenRepository
}

func GetController(userRepo *repository.UserRepository, tokenRepo *repository.TokenRepository) *Controller {
	return &Controller{userRepo: userRepo, tokenRepo: tokenRepo}
}

func (c *Controller) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error"))
		return
	}

	id := uuid.NewString()
	c.userRepo.CreateUser(id, user.Login, user.Password)

	accessToken, err := helpers.GenToken()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error"))
		return
	}

	refreshToken, err := helpers.GenToken()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error"))
		return
	}

	if err := c.tokenRepo.CreateToken(r.Context(), accessToken, id, false); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error"))
		return
	}
	if err := c.tokenRepo.CreateToken(r.Context(), refreshToken, id, true); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error"))
		return
	}

	res, _ := json.Marshal(&models.Response{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Error:        "false",
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func (c *Controller) LoginHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error"))
		return
	}

	accessBytes := make([]byte, 32)
	if _, err := rand.Read(accessBytes); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error"))
		return
	}
	accessToken := base64.URLEncoding.EncodeToString(accessBytes)

	refreshBytes := make([]byte, 32)
	if _, err := rand.Read(refreshBytes); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error"))
		return
	}
	refreshToken := base64.URLEncoding.EncodeToString(refreshBytes)

	if err := c.tokenRepo.CreateToken(r.Context(), accessToken, user.ID, false); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error"))
		return
	}
	if err := c.tokenRepo.CreateToken(r.Context(), refreshToken, user.ID, true); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error"))
		return
	}

	res, _ := json.Marshal(&models.Response{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Error:        "false",
	})

	if c.userRepo.CheckLogin(user.Login, user.Password) {
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	}

	res, _ = json.Marshal(&models.Response{
		Error:   "invalid_credentials",
		Message: "Неправильный логин или пароль",
	})

	w.WriteHeader(http.StatusUnauthorized)
	w.Write(res)
}

func (c *Controller) CheckAuthHandler(w http.ResponseWriter, r *http.Request) {
	token := strings.TrimPrefix(r.Header["Authorization"][0], "Bearer ")
	println(token)
	userID, err := c.tokenRepo.CheckToken(r.Context(), token)
	if err != nil {
		res, _ := json.Marshal(&models.Response{
			Error: "UNAUTHORIZED",
		})
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(res))
		return
	}

	res, err := json.Marshal(&models.Response{
		Message: userID,
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func (c *Controller) UpdateTokensHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.Response{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res, _ := json.Marshal(&models.Response{
			Error: "UNAUTHORIZED",
		})
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(res))
		return
	}

	access, refresh, err := c.tokenRepo.RefreshTokens(r.Context(), req.RefreshToken)
	if err != nil {
		res, _ := json.Marshal(&models.Response{
			Error: "UNAUTHORIZED",
		})
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(res))
		return
	}

	res, _ := json.Marshal(&models.Response{
		AccessToken:  access,
		RefreshToken: refresh,
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}
