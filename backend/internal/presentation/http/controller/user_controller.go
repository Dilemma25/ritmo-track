package controller

import (
	"fmt"
	"net/http"
	"ritmotrack-backend/internal/application/user/dto"
	"ritmotrack-backend/internal/application/user/usecase"
	"ritmotrack-backend/internal/presentation/http/parser"
	"ritmotrack-backend/internal/presentation/http/request"
	"ritmotrack-backend/internal/presentation/http/validator"
	"ritmotrack-backend/internal/presentation/http/writer"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	UserCreateUseCase  *usecase.UserCreateUseCase
	UserGetByIDUseCase *usecase.UserGetByIDUseCase
	UserGetAllUseCase  *usecase.UserGetAllUseCase
	validator          *validator.Validator
}

func NewUserController(
	userCreateUseCase *usecase.UserCreateUseCase,
	userGetByIDUseCase *usecase.UserGetByIDUseCase,
	userGetAllUseCase *usecase.UserGetAllUseCase,
) *UserController {
	return &UserController{
		UserCreateUseCase:  userCreateUseCase,
		UserGetByIDUseCase: userGetByIDUseCase,
		UserGetAllUseCase:  userGetAllUseCase,
		validator:          validator.New(),
	}
}

func (ths *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	users, _ := ths.UserGetAllUseCase.Execute(r.Context())

	writer.WriteResponseJSON(w, http.StatusOK, users)
}

func (ths *UserController) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	idInt, err := parser.ParseInt(idStr)

	if err != nil {
		writer.WriteErrorJSON(w, http.StatusBadRequest, fmt.Errorf("invalid id"))
		return
	}

	user, _ := ths.UserGetByIDUseCase.Execute(r.Context(), idInt)

	writer.WriteResponseJSON(w, http.StatusOK, user)
}

func (ths *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var req request.UserCreateRequest

	if err := parser.ParseRequestBody(r, &req); err != nil {
		writer.WriteErrorJSON(w, http.StatusBadRequest, err)

		return
	}

	if errs := ths.validator.ValidateStruct(&req); len(errs) > 0 {
		for _, err := range errs {
			writer.WriteErrorJSON(w, http.StatusBadRequest, fmt.Errorf(err))
		}
		return
	}

	userDTO := dto.CreateUserDTO{
		Login:    req.Login,
		Name:     req.Name,
		Password: req.Password,
	}

	user, err := ths.UserCreateUseCase.Execute(r.Context(), userDTO)
	if err != nil {
		writer.WriteErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	writer.WriteResponseJSON(w, http.StatusOK, user)
}
