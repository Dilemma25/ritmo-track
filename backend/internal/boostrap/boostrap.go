package boostrap

import (
	"ritmotrack-backend/internal/core/config"
	"ritmotrack-backend/internal/core/di"
	httpLayer "ritmotrack-backend/internal/presentation/http"
)

func Start() {
	container := di.CreateContainer()

	container.Invoke(func(config *config.Config) {
		server := httpLayer.NewServer()

		httpLayer.RegisterRouters(server.Router(), container)

		server.Start(config.Port)
	})
}
