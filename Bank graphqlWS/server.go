package main

import (
	"Bank_graphqlWS/graph"
	"Bank_graphqlWS/initializers"
	"Bank_graphqlWS/repository"
	"Bank_graphqlWS/service"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	// Initialize the database
	initializers.ConnectToDb()
	initializers.SyncDatabase()

	// Set up the port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Initialize the repository and service
	bankRepo := repository.NewBankRepository(initializers.DB)
	insuranceRepo := repository.NewInsuranceRepository(initializers.DB)
	purchaseRepo := repository.NewPurchaseRepository(initializers.DB)
	bankService := &service.BankServices{
		Bankrepo: bankRepo,
	}
	insuranceService := service.NewInsuranceService(insuranceRepo, purchaseRepo)

	// Create the GraphQL server with the resolver
	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			BankService:      bankService,
			InsuranceService: insuranceService,
		},
	}))

	// Add transport options for the server
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	// Set up HTTP routes for GraphQL playground and mutation queries
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	// Log and start the server
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
