package controller

import (
	"errors"
	"net/http"
	"ritmotrack-backend/internal/application/apperror"
	"ritmotrack-backend/internal/application/dto"
	userUseCase "ritmotrack-backend/internal/application/usecase/user"
	"ritmotrack-backend/internal/presentation/http/apierror"
	"ritmotrack-backend/internal/presentation/http/ctxvalue"
	"ritmotrack-backend/internal/presentation/http/mapper"
	"ritmotrack-backend/internal/presentation/http/responder"
	"ritmotrack-backend/internal/presentation/http/shema"
	"ritmotrack-backend/internal/presentation/http/validator"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	CreateUserUseCase  userUseCase.CreateUserUseCase
	GetUserByIDUseCase userUseCase.GetUserByIDUseCase
	GetUserAllUseCase  userUseCase.GetAllUsersUseCase
	PatchUserUseCase   userUseCase.PatchUserUseCase
	validator          *validator.Validator
	responder          *responder.Responder
}

func NewUserController(
	createUserUseCase userUseCase.CreateUserUseCase,
	getUserByIDUseCase userUseCase.GetUserByIDUseCase,
	getAllUsersUseCase userUseCase.GetAllUsersUseCase,
	patchUserUseCase userUseCase.PatchUserUseCase,
	validator *validator.Validator,
	responder *responder.Responder,
) *UserController {
	return &UserController{
		CreateUserUseCase:  createUserUseCase,
		GetUserByIDUseCase: getUserByIDUseCase,
		GetUserAllUseCase:  getAllUsersUseCase,
		PatchUserUseCase:   patchUserUseCase,
		validator:          validator,
		responder:          responder,
	}
}

func (ths *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	users, _ := ths.GetUserAllUseCase.Execute(r.Context())

	ths.responder.ResponseOk(w, users)
}

func (ths *UserController) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	idUint, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		ths.responder.ResponseError(w, apierror.NewErrInvalidUrlParameter(err))
		return
	}

	user, _ := ths.GetUserByIDUseCase.Execute(r.Context(), uint(idUint))

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
		Password: req.Password,
	}

	user, err := ths.CreateUserUseCase.Execute(r.Context(), userDTO)
	if err != nil {

		if errors.Is(err, apperror.ErrUserAlreadyExists) {
			ths.responder.ResponseError(w, apierror.NewErrUserAlreadyExists())
			return
		}

		ths.responder.ResponseError(w, apierror.NewErrInternal(err))
		return
	}

	response := mapper.CreateUserResponseFromDTO(user)

	ths.responder.ResponseOk(w, response)
}

func (ths *UserController) Patch(w http.ResponseWriter, r *http.Request) {
	var req shema.PatchUserRequest
	//TODO проверять пустой боди и выбрасывать 400
	if err := ths.validator.ValidateBody(r, &req); err != nil {
		ths.responder.ResponseError(w, err)

		return
	}

	userId, ok := r.Context().Value(ctxvalue.UserID).(uint)

	if !ok {
		ths.responder.ResponseError(w, apierror.NewErrUserContextNotFound())
	}

	patchDTO := dto.PatchUserDTO{
		Name:     req.Name,
		Password: req.Password,
	}

	if err := ths.PatchUserUseCase.Execute(r.Context(), userId, patchDTO); err != nil {

		if errors.Is(err, apperror.ErrUserNotFound) {
			ths.responder.ResponseError(w, apierror.NewErrUserNotFound())
		}

		ths.responder.ResponseError(w, apierror.NewErrInternal(err))
		return
	}

	ths.responder.ResponseNoContent(w)
}
