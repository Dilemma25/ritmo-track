package controller

import (
	"net/http"
	"ritmotrack-backend/internal/application/dto"
	userUseCase "ritmotrack-backend/internal/application/usecase/user"
	"ritmotrack-backend/internal/presentation/http/apierror"
	"ritmotrack-backend/internal/presentation/http/parser"
	"ritmotrack-backend/internal/presentation/http/responder"
	"ritmotrack-backend/internal/presentation/http/shema"
	"ritmotrack-backend/internal/presentation/http/validator"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	UserCreateUseCase  *userUseCase.CreateUserUseCase
	UserGetByIDUseCase *userUseCase.GetUserByIDUseCase
	UserGetAllUseCase  *userUseCase.GetAllUsersUseCase
	validator          *validator.Validator
	responder          *responder.Responder
}

func NewUserController(
	userCreateUseCase *userUseCase.CreateUserUseCase,
	userGetByIDUseCase *userUseCase.GetUserByIDUseCase,
	userGetAllUseCase *userUseCase.GetAllUsersUseCase,
	validator *validator.Validator,
	responder *responder.Responder,
) *UserController {
	return &UserController{
		UserCreateUseCase:  userCreateUseCase,
		UserGetByIDUseCase: userGetByIDUseCase,
		UserGetAllUseCase:  userGetAllUseCase,
		validator:          validator,
		responder:          responder,
	}
}

func (ths *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	users, _ := ths.UserGetAllUseCase.Execute(r.Context())

	ths.responder.ResponseOk(w, users)
}

func (ths *UserController) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	idInt, err := parser.ParseInt(idStr)

	if err != nil {
		ths.responder.ResponseError(w, apierror.NewErrInvalidUrlParameter(err))
		return
	}

	user, _ := ths.UserGetByIDUseCase.Execute(r.Context(), idInt)

	ths.responder.ResponseOk(w, user)
}

func (ths *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var req shema.CreateUserRequest

	if err := ths.validator.ValidateBody(r, &req); err != nil {
		ths.responder.ResponseError(w, err)

		return
	}

	userDTO := dto.CreateUserDTO{
		Login:    req.Login,
		Name:     req.Name,
		Password: req.Password,
	}

	user, err := ths.UserCreateUseCase.Execute(r.Context(), userDTO)
	if err != nil {
		ths.responder.ResponseError(w, apierror.NewErrInternal(err))
		return
	}

	ths.responder.ResponseOk(w, user)
}
