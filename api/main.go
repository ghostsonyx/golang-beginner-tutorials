package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type DeviceInfo struct {
	IP      string `json:"ip"`
	Country string `json:"country"`
	CC      string `json:"cc"`
}

func main() {
	router := gin.Default()
	router.GET("/device-info", getDeviceInfo)

	router.Run(":8080")
}

func getDeviceInfo(c *gin.Context) {
	response, err := http.Get("https://api.myip.com/")
	var responseDeviceInfo DeviceInfo
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		responseData, _ := io.ReadAll(response.Body)
		json.Unmarshal(responseData, &responseDeviceInfo)
	}

	c.IndentedJSON(http.StatusOK, responseDeviceInfo)
}
