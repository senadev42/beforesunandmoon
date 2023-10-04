package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// types
type Coord struct {
	ra, dec string
}

// Cobra commands
var RootCmd = &cobra.Command{
	Use:   "beforesunandmoon",
	Short: "A simple CLI tool for terminal stargazing",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify a subcommand. Use 'bsam help' for usage.")
	},
}

func newCoord(coordType string) string {
	var value string

	//no ternary operators 4u
	var randsign int = 1
	if rand.Intn(2) == 0 {
		randsign = -1
	}

	rand.Seed(time.Now().UnixNano())

	if coordType == "ra" {
		// For RA, generate hours, minutes, and seconds
		value = fmt.Sprintf("%02d:%02d:%02.2f", rand.Intn(24), rand.Intn(60), rand.Float64()*60)
	} else if coordType == "dec" {
		// For Dec, generate degrees, arcminutes, and arcseconds
		value = fmt.Sprintf("%+03d:%02d:%02.2f", rand.Intn(180)*randsign, rand.Intn(60), rand.Float64()*60)
	}

	return value
}

// end goal: build a valid query for this thing
// https://api.astrocats.space/catalog?ra=21:23:32.16&dec=-53:01:36.08&radius=400
var FeelingLuckyCmd = &cobra.Command{
	Use:   "feelinglucky",
	Short: "Picks a random spot in space and tells you what's there",
	Run: func(cmd *cobra.Command, args []string) {

		//generate some random coordinates
		ra, dec := newCoord("ra"), newCoord("dec")

		radius := 40000000

		// Build the query URL
		queryURL := fmt.Sprintf("https://api.astrocats.space/catalog?ra=%s&dec=%s&radius=%d",
			ra, dec, radius)

		print("Calling ", queryURL)

	},
}

func parseCoordinates(coord string) (float64, float64, error) {
	parts := strings.Split(coord, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid coordinates format")
	}

	lat := parts[0]
	lon := parts[1]

	// Convert coordinates from string to float64
	latFloat, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing latitude: %v", err)
	}

	lonFloat, err := strconv.ParseFloat(lon, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing longitude: %v", err)
	}

	return latFloat, lonFloat, nil
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	RootCmd.AddCommand(FeelingLuckyCmd)

	// Execute the CLI application
	Execute()
}
