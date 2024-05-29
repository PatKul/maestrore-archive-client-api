package domain

import (
	"fmt"
	"maestrore/core"
	"maestrore/domain/location"
	"maestrore/domain/sales"
	"net/http"
)

/**
 * APIService
 */
type APIService struct {
	config    *core.Config
	encryptor *core.Encryptor
	database  *core.MySqlDatabase
}

func NewAPIService(config *core.Config) *APIService {
	service := &APIService{config: config}

	return service
}

/**
 * Initialize the API service, connect to the database
 * @return error
 */
func (srv *APIService) Init() error {
	srv.encryptor = core.NewEncryptor()
	srv.database = core.NewMySqlDatabase(srv.config, srv.encryptor)

	error := srv.database.Connect()
	if error != nil {
		return fmt.Errorf("failed to connect to database: %s", error.Error())
	}

	fmt.Printf("Connected to database - Host: %s,  User: %s !!\n",
		srv.config.DatabaseHost,
		srv.config.DatabaseUser)

	return nil
}

func (srv *APIService) Run() {
	defer srv.database.Close()

	router := http.NewServeMux()

	// route handling
	defaultRouteHandler := NewDefaultRouteHandler(router, srv.encryptor)
	defaultRouteHandler.RegisterRoute()

	locationRouteHandler := location.NewRouteHandler(router, srv.database.GetConnection())
	locationRouteHandler.RegisterRoute()

	salesRouteHandler := sales.NewRouteHandler(router, srv.database.GetConnection())
	salesRouteHandler.RegisterRoute()

	handlers := CorsMiddleware(router)

	address := fmt.Sprintf(":%s", srv.config.Port)

	fmt.Printf("Server running on port %s\n\n", srv.config.Port)
	error := http.ListenAndServe(address, handlers)
	if error != nil {
		panic(error)
	}
}
