package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// Convert sexigesimal RA to decimal degrees
func RaToDecimal(ra string) float64 {

	parts := strings.Split(ra, ":")

	hours, _ := strconv.ParseFloat(parts[0], 64)
	minutes, _ := strconv.ParseFloat(parts[1], 64)
	seconds, _ := strconv.ParseFloat(parts[2], 64)

	raDecimal := 15 * (hours + minutes/60 + seconds/3600)
	return raDecimal
}

// Convert sexigesimal Dec to decimal degrees
func DecToDecimal(dec string) float64 {

	parts := strings.Split(dec, ":")

	degrees, _ := strconv.ParseFloat(parts[0], 64)
	arcminutes, _ := strconv.ParseFloat(parts[1], 64)
	arcseconds, _ := strconv.ParseFloat(parts[2], 64)

	decDecimal := degrees + arcminutes/60 + arcseconds/3600
	return decDecimal
}

func Fetch(url string) []byte {
	// Build the query URL
	queryURL := url

	// Make API request
	resp, err := http.Get(queryURL)
	if err != nil {
		fmt.Println("Error making API request:", err)
		return nil
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	return body
}
