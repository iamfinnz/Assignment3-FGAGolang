package main

// Import library
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

// Pendefinisian
type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

// Function utama
func main() {
	// Set up router Gin
	router := gin.Default()

	// Menampilkan index.html di folder view
	router.StaticFS("/view", http.Dir("./view"))

	// Mengarahkan routing ke /status
	router.GET("/status", getStatus)

	// Jalankan function updateFileStatus untuk mengupdate value
	go updateFileStatus()

	// Run server di port 8080
	router.Run(":8080")
}

// Function untuk mendapatkan value status
func getStatus(c *gin.Context) {
	status := readStatus()

	// Pendefinisian
	waterStatus := status.Water
	windStatus := status.Wind

	// Persiapkan response
	response := gin.H{
		"status":       status,
		"water_status": waterStatus,
		"wind_status":  windStatus,
	}

	c.JSON(http.StatusOK, response)
}

// Function untuk membaca value status dari file status.json
func readStatus() Status {
	// Baca status dari file status.json
	file, err := ioutil.ReadFile("status.json")
	// Error handling
	if err != nil {
		fmt.Printf("Error saat membaca file: %v\n", err)
	}

	var status Status
	err = json.Unmarshal(file, &status)
	// Error handling
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
	}

	// Mengembalikan nilai status
	return status
}

// Function untuk mengupdate value status yang ada di status.json
func updateFileStatus() {
	for {
		// Generate nilai random water dan wind dengan nilai maks 100
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		// Buat object status yang terdiri dari water dan wind
		status := Status{
			Water: water,
			Wind:  wind,
		}

		// Tulis status ke file status.json
		file, err := json.MarshalIndent(status, "", "  ")
		// Error handling
		if err != nil {
			fmt.Printf("Error marshalling JSON: %v\n", err)
		}

		err = ioutil.WriteFile("status.json", file, 0644)
		// Error handling
		if err != nil {
			fmt.Printf("Error saat menulikan ke file: %v\n", err)
		}

		// Atur sleep ke 15 detik
		time.Sleep(15 * time.Second)
	}
}