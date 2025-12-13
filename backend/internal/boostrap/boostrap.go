package boostrap

import (
	"ritmotrack-backend/internal/infrastructure/config"
	httpLayer "ritmotrack-backend/internal/presentation/http"
)

func Start() {

	GetContainer().Invoke(func(config *config.Config) {
		server := httpLayer.NewServer()

		httpLayer.RegisterRouters(server.Router(), GetContainer())

		server.Start(config.Port)
	})
}
