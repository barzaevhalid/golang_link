package main

import (
	"fmt"
	"net/http"
	"rest_api/configs"
	"rest_api/internal/auth"
	"rest_api/internal/link"
	"rest_api/internal/stat"
	"rest_api/internal/user"
	"rest_api/pkg/db"
	"rest_api/pkg/event"
	"rest_api/pkg/middleware"
)

func App() http.Handler {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	//Repositores
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	//Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	//Handler
	link.NewLinkHanlder(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config:         conf,
		// StatRepository: statRepository,
		EventBus: eventBus,
	})
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepository,
		Config:         conf,
	})

	go statService.AddClick()

	// Middelwares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)
	return stack(router)

}
func main() {
	app := App()
	server := http.Server{
		Addr:    ":8080",
		Handler: app,
	}

	fmt.Println("Server is listening on port 8080")
	server.ListenAndServe()
}
