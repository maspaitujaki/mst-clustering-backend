package main

import (
	"backend/database"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	initDB()

	router := mux.NewRouter().StrictSlash(true)
	// initaliseHandlers(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	log.Println("Starting the HTTP server on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))

}

func initaliseHandlers(router *mux.Router) {
	// router.HandleFunc("/penyakit/create", controllers.CreatePenyakit).Methods("POST")

}

func initDB() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Error loading .env file")
	}
	config :=
		database.Config{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Name:     os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
		}
	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	// database.MigratePenyakit(&entity.Penyakit{})
	// database.MigratePemeriksaan(&entity.Pemeriksaan{})
}

// func main() {
// 	xAtr, yAtr, nodes, _ := mst.ReadNodes("test/tc1.csv")
// 	edges, _ := mst.ReadEdges(nodes)
// 	edges = mst.QuickSortEdgesAsc(edges)

// 	tree := mst.KruskalMST(nodes, edges)
// 	result := clustering.MakeCluster(tree, 3)
// 	for _, cluster := range result {
// 		for _, node := range cluster {
// 			fmt.Printf("%s ", node.Name)
// 		}
// 		fmt.Println("Ganti Cluster")
// 	}
// 	visualisasi.MakeScatter(xAtr, yAtr, result)
// }
