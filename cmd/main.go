package main

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/link"
	"go/adv-demo/internal/stat"
	"go/adv-demo/internal/user"
	"go/adv-demo/pkg/db"
	"go/adv-demo/pkg/event"
	"go/adv-demo/pkg/middleware"
	"net/http"
)

func App() http.Handler {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	//Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	//Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.StatService(stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	//Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	//while dont have service
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		EventBus:       eventBus,
		Config:         conf,
	})
	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepository,
		Config:         conf,
	})

	go statService.AddClick()

	//Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	return stack(router)
}

func main() {
	app := App()
	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	fmt.Println("Server is listenning on port 8081")
	server.ListenAndServe()
}
