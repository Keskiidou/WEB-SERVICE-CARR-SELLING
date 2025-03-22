package server

import (
	"insurance-soap-go/handlers"
	"insurance-soap-go/initializers"
	"insurance-soap-go/repositories"
	"insurance-soap-go/services"
	"net/http"
	"os"
)

func StartServer() {
	// Initialize the repository, service, and handler
	purchaseRepo := repositories.NewPurchaseRepository(initializers.DB)
	insuranceService := services.NewInsuranceService(purchaseRepo)
	insuranceHandler := handlers.NewInsuranceHandler(insuranceService)

	// Set up the SOAP endpoint
	http.Handle("/ws", insuranceHandler)

	// Get the PORT from the environment variables
	port := os.Getenv("PORT")

	// Start the HTTP server
	println("SOAP server is running on http://localhost:" + port + "/ws")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic("Failed to start the server: " + err.Error())
	}
}
