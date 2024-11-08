package api

import (
	db "github.com/Sandhya-Pratama/simple-bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Store struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Store {

	server := &Store{
		store: store,
	}

	router := gin.Default()

	//add router
	router.POST("/accounts", server.createAccount)

	server.router = router
	return server
}

// start run http server
func (server *Store) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}

}
