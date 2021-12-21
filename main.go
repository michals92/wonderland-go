package main

import (
	"os"

	"github.com/michals92/wonderland-go/controller"
	"github.com/michals92/wonderland-go/repository"
	"github.com/michals92/wonderland-go/router"
	"github.com/michals92/wonderland-go/service"
)

var (
	firestoreRepo  = repository.NewFirestoreRepository()
	gridService    = service.NewGridService(firestoreRepo)
	gridController = controller.NewGridController(gridService)
	httpRouter     = router.NewMuxRouter()
)

func main() {
	/*httpRouter.POST("/login", userController.Login)
	httpRouter.POST("/register", userController.Register)
	httpRouter.POST("/logout", userController.Logout)
	httpRouter.POST("/token/refresh", userController.Refresh)*/

	//httpRouter.AUTH_POST("/page", pageController.PostPage)

	httpRouter.POST("/parcels", gridController.GetParcels)
	httpRouter.POST("/parcel", gridController.AddParcel)
	httpRouter.POST("/pinArt", gridController.PinArt)
	httpRouter.GET("/unpinArt", gridController.UnpinArt)

	httpRouter.SERVE(os.Getenv("PORT"))
}
