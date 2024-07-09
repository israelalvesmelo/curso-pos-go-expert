package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/israelalvesmelo/curso-pos-go-expert/APIs/configs"
	_ "github.com/israelalvesmelo/curso-pos-go-expert/APIs/docs"
	"github.com/israelalvesmelo/curso-pos-go-expert/APIs/internal/entity"
	"github.com/israelalvesmelo/curso-pos-go-expert/APIs/internal/infra/database"
	"github.com/israelalvesmelo/curso-pos-go-expert/APIs/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Israel
// @contact.url    https://github.com/israelalvesmelo

// @license.name   My License
// @license.url    https://github.com/israelalvesmelo

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Use(middleware.WithValue("jwt", configs.TokenAuth))
	router.Use(middleware.WithValue("JwtExperesIn", configs.JwtExperesIn))

	router.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth)) // That middleware will search for a JWT token in a http request and safe it in the context
		r.Use(jwtauth.Authenticator)               // That middleware will validate the JWT token with the secret
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	router.Post("/users", userHandler.CreateUser)
	router.Post("/users/generate_token", userHandler.GetJWT)

	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", router)
}

// LogRequest is a middleware that logs the request method and path. It is an example of a custom middleware.
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
