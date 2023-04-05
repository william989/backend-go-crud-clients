package main

import (
	"fmt"
	"net/http"

	"golang-crud-rest-api/controller"
	router "golang-crud-rest-api/http"
	"golang-crud-rest-api/repository"
	"golang-crud-rest-api/service"
)

var (
	clientRepository repository.ClientRepository = repository.NewClientRepository()
	clientService    service.ClientService       = service.NewClientService(clientRepository)
	postController   controller.ClientController = controller.NewClientController(clientService)
	httpRouter       router.Router               = router.NewGinRouter()
	//httpRouter       router.Router               = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Server listo y corriendo . . .")
	})
	httpRouter.GET("/posts", postController.GetClientes)
	httpRouter.POST("/posts", postController.AddCliente)
	httpRouter.POST("/post/{id}", postController.GetClienteById)
	httpRouter.SERVE(port)
}
