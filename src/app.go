package app

import (
	"home-booking-api/src/routes"
	"home-booking-api/src/services"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/lpernett/godotenv"
	"github.com/stripe/stripe-go/v75"
)

func App() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := os.Getenv("POSTGRES_CONNECTION_STRING")
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	db, err := sqlx.Connect("postgres", connectionString)

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	defer db.Close()

	mux := mux.NewRouter()

	paymentRoutes := mux.PathPrefix("/payment").Subrouter()
	bookingRoutes := mux.PathPrefix("/booking").Subrouter()
	houseRoutes := mux.PathPrefix("/house").Subrouter()
	familyRoutes := mux.PathPrefix("/family").Subrouter()
	userRoutes := mux.PathPrefix("/user").Subrouter()

	//routes.RegisterElectricityRoutes(mux)

	injectActiveSession, err := services.ClerkActiveSession()

	if err != nil {
		log.Fatal("Error creating clerk client")
	}

	mux.Use(injectActiveSession)

	routes.RegisterPaymentRoutes(paymentRoutes, db)
	routes.RegisterBookingsRoutes(bookingRoutes, db)
	routes.RegisterHouseRoutes(houseRoutes, db)
	routes.RegisterFamilyRoutes(familyRoutes, db)
	routes.RegisterUserRoutes(userRoutes, db)

	log.Fatal(http.ListenAndServe(":8000",
		handlers.CORS(handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(mux)),
	)

}
