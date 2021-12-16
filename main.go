package main

import (
	router "github.com/michals92/wonderland-go/router"

	"github.com/michals92/wonderland-go/repository"
	"github.com/michals92/wonderland-go/service"
)

var (
	firestoreRepo = repository.NewFirestoreRepository()

	gridService = service.NewGridService(firestoreRepo)

	httpRouter = router.NewMuxRouter()

/*
	userController = controller.NewUserController(userService)
	pageController = controller.NewPageController(pageService, cardService)

	httpRouter = router.NewMuxRouter()*/
)

func main() {

}
