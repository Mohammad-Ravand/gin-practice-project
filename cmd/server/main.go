package main
import (
	"fmt"
	"net/http"
	"example.com/gin-first-project/internal/config"
	"example.com/gin-first-project/internal/router"
)
func main() {
	db := config.CreateDatabaseConnection()
	// Initialize the router
	r := router.SetupRouter(db)
	// Start the server
	port := ":3001"
	fmt.Printf("Server listening on port %s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		panic(err)
	}
}
