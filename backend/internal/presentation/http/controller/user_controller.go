package controller

import (
	"encoding/json"
	"fmt"
	"ritmotrack-backend/internal/application/user/dto"
	"ritmotrack-backend/internal/presentation/http/sheme"
	"ritmotrack-backend/internal/presentation/http/utils"
	"strconv"

	"github.com/go-chi/chi/v5"

	"net/http"
	"ritmotrack-backend/internal/application/user/usecase"
)

type UserController struct {
	UserCreateUseCase  usecase.UserCreateUseCase
	UserGetByIDUseCase usecase.UserGetByIDUseCase
	UserGetAllUseCase  usecase.UserGetAllUseCase
}

func (ths *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	users, _ := ths.UserGetAllUseCase.Execute(r.Context())

	utils.ResponseJSON(w, http.StatusOK, users)
}

func (ths *UserController) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	//TODO вынести парсинг id в utils
	idInt, err := strconv.ParseInt(idStr, 10, 32)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, fmt.Errorf("invalid id"))
		return
	}

	user, _ := ths.UserGetByIDUseCase.Execute(r.Context(), idInt)

	utils.ResponseJSON(w, http.StatusOK, user)
}

func (ths *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var req sheme.UserCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, fmt.Errorf("invalid json"))
		return
	}

	userDTO := dto.CreateUserDTO{
		Login:    req.Login,
		Name:     req.Name,
		Password: req.Password,
	}

	user, err := ths.UserCreateUseCase.Execute(r.Context(), userDTO)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, user)
}
