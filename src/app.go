package app

import (
	"fmt"
	"home-booking-api/src/routes"
	"log"
	"net/http"
	"os"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/lpernett/godotenv"
)

func App() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := os.Getenv("POSTGRES_CONNECTION_STRING")
	clerkSecret := os.Getenv("CLERK_SECRET_KEY")

	clerkClient, err := clerk.NewClient(clerkSecret)

	if err != nil {
		log.Fatal("Error initiating clerk client")
	}

	injectActiveSession := clerk.RequireSessionV2(clerkClient)

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

	fmt.Println(mux)

	//routes.RegisterElectricityRoutes(mux)

	mux.Use(injectActiveSession)

	routes.RegisterPaymentRoutes(paymentRoutes, db, clerkClient)	
	routes.RegisterBookingsRoutes(bookingRoutes, db, clerkClient)
	routes.RegisterHouseRoutes(houseRoutes, db, clerkClient)
	routes.RegisterFamilyRoutes(familyRoutes, db, clerkClient)

	log.Fatal(http.ListenAndServe(":8000", 
		handlers.CORS(handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), 
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), 
		handlers.AllowedOrigins([]string{"*"}))(mux)),
	)

}
