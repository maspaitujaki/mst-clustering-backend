package controllers

import (
	"backend/clustering"
	"backend/database"
	"backend/entity"
	"backend/mst"
	"backend/visualisasi"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func HandleNewClustering(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleNewClustering")
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		log.Println("Error parsing multipart form: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	uploadedFiles, handler, err := r.FormFile("dataFile")
	if err != nil {
		log.Println("Error getting file: ", err)
		http.Error(w, "Bad input", http.StatusBadRequest)
		return
	}
	defer uploadedFiles.Close()
	dir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	filename := handler.Filename

	fileLocation := filepath.Join(dir, "uploads", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("Error opening file: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFiles); err != nil {
		log.Println("Error copying file: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("CSV file uploaded successfully: " + fileLocation)
	log.Println("Starting clustering...")

	xAtr, yAtr, nodes, err := mst.ReadNodes(fileLocation)
	if err != nil {
		log.Println("Error reading nodes: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	node := r.FormValue("clusterCount")
	if node == "" {
		log.Println("Error getting cluster count")
		http.Error(w, "Error getting cluster count", http.StatusBadRequest)
		return
	}

	nNode, _ := strconv.Atoi(node)
	if nNode > len(nodes) {
		log.Println("Error: cluster count must not greater than node count")
		http.Error(w, "Cluster count must not greater than node count", http.StatusBadRequest)
		return
	}

	edges, _ := mst.ReadEdges(nodes)
	edges = mst.QuickSortEdgesAsc(edges)

	tree := mst.KruskalMST(nodes, edges)
	result := clustering.MakeCluster(tree, nNode)
	visualisasi.MakeScatter(xAtr, yAtr, result, filepath.Join(dir, "results", strings.Split(filename, ".")[0]+".png"))

	log.Println("Clustering finished, results on:" + filepath.Join(dir, "results", strings.Split(filename, ".")[0]+".png"))

	var newLog entity.Log
	newLog.NamaFileAsli = filename
	newLog.Tanggal = time.Now()
	newLog.N_Cluster = nNode

	// Penambahan ke database
	if result := database.Connector.Create(&newLog); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	img, err := os.Open(filepath.Join(dir, "results", strings.Split(filename, ".")[0]+".png"))
	if err != nil {
		log.Fatal(err) // perhaps handle this nicer
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/png") // <-- set the content-type header
	w.WriteHeader(http.StatusOK)
	io.Copy(w, img)
}

func HandleGetLogs(w http.ResponseWriter, r *http.Request) {
	var logs []entity.Log
	database.Connector.Find(&logs)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(logs)

}
