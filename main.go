package main

import (
	"backend/controllers"
	"backend/database"
	"backend/entity"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	initDB()
	initStorage()

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	log.Println("Starting the HTTP server on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))

}

func initaliseHandlers(router *mux.Router) {
	// router.HandleFunc("/penyakit/create", controllers.CreatePenyakit).Methods("POST")
	router.HandleFunc("/clustering", controllers.HandleNewClustering).Methods("POST")
	router.HandleFunc("/clustering", controllers.HandleGetLogs).Methods("GET")
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
	database.MigrateLog(&entity.Log{})
	// database.MigratePemeriksaan(&entity.Pemeriksaan{})
}

func initStorage() {
	dir, _ := os.Getwd()
	resultPath := filepath.Join(dir, "results")
	os.RemoveAll(resultPath)
	uploadPath := filepath.Join(dir, "uploads")
	os.RemoveAll(uploadPath)
	os.MkdirAll(resultPath, 0700)
	os.MkdirAll(uploadPath, 0700)
	log.Println("Storage initialized")
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
