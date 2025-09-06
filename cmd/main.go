package main

import (
	"net/http"
	"rest_api/configs"
	"rest_api/internal/auth"
	"rest_api/internal/link"
	"rest_api/internal/user"
	"rest_api/pkg/db"
	"rest_api/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositores
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)

	//Services
	authService := auth.NewAuthService(userRepository)

	//Handler
	link.NewLinkHanlder(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config:         conf,
	})
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})

	// Middelwares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	server.ListenAndServe()

}
