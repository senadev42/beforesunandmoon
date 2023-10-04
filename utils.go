package main

import (
	"strconv"
	"strings"
)

// Convert sexigesimal RA to decimal degrees
func RaToDecimal(ra string) float64 {
	// Split the RA string by colon
	parts := strings.Split(ra, ":")
	// Parse each part into a float value
	hours, _ := strconv.ParseFloat(parts[0], 64)
	minutes, _ := strconv.ParseFloat(parts[1], 64)
	seconds, _ := strconv.ParseFloat(parts[2], 64)
	// Apply the conversion formula
	raDecimal := 15 * (hours + minutes/60 + seconds/3600)
	return raDecimal
}

// Convert sexigesimal Dec to decimal degrees
func DecToDecimal(dec string) float64 {
	// Split the Dec string by colon
	parts := strings.Split(dec, ":")
	// Parse each part into a float value
	degrees, _ := strconv.ParseFloat(parts[0], 64)
	arcminutes, _ := strconv.ParseFloat(parts[1], 64)
	arcseconds, _ := strconv.ParseFloat(parts[2], 64)
	// Apply the conversion formula
	decDecimal := degrees + arcminutes/60 + arcseconds/3600
	return decDecimal
}
