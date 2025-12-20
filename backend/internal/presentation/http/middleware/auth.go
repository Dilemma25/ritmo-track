package middleware

import (
	"context"
	"log"
	"net/http"
	"ritmotrack-backend/internal/application/usecase/auth"
	"ritmotrack-backend/internal/presentation/http/apierror"
	"ritmotrack-backend/internal/presentation/http/responder"
	"strings"
)

type AuthMiddleware struct {
	checkJwtUseCase auth.CheckJwtUseCase
	responder       *responder.Responder
}

func NewAuthMiddleware(
	checkJwtUseCase auth.CheckJwtUseCase,
	responder *responder.Responder,
) *AuthMiddleware {
	return &AuthMiddleware{
		checkJwtUseCase: checkJwtUseCase,
		responder:       responder,
	}
}

func (ths *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			apiErr := apierror.NewErrAuthFailed("Авторизационный токен не предоставлен")
			ths.responder.ResponseError(w, apiErr)

			return
		}

		bearerPrefix := "Bearer "

		if !strings.HasPrefix(authHeader, bearerPrefix) {
			apiErr := apierror.NewErrAuthFailed("Неверный формат токена")
			ths.responder.ResponseError(w, apiErr)

			return
		}

		token := strings.TrimPrefix(authHeader, bearerPrefix)

		tokenClaims, err := ths.checkJwtUseCase.Execute(r.Context(), token)

		if err != nil {
			log.Println(err)
			apiErr := apierror.NewErrAuthFailed("Неверный токен")
			ths.responder.ResponseError(w, apiErr)

			return
		}

		userIDKey, ok := r.Context().Value("userID").(uint)

		if !ok {

		}

		ctx := context.WithValue(r.Context(), userIDKey, tokenClaims.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
