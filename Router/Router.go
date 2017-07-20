package Router

import (
	"../Handlers"
	"github.com/julienschmidt/httprouter"
)

func Routes() (routes *httprouter.Router) {
	routes = httprouter.New()
	routes.GET("/", Handlers.HomeHandler)
	routes.GET("/name/:name", Handlers.NameHandler)
	routes.GET("/booksjson", Handlers.BookJSONHandler)
	routes.GET("/bookshtml", Handlers.BookHTMLHandler)
	routes.GET("/users", Handlers.UsersHandler)
	return
}
