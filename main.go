package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "beforesunandmoon",
	Short: "A simple CLI tool for terminal stargazing",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify a subcommand. Use 'beforesunandmoon help' for usage.")
	},
}

func newCoord(coordType string) string {
	var value string

	//rand.Seed(time.Now().UnixNano())
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	//Right Ascenation
	if coordType == "ra" {
		// For RA, generate hours, minutes, and seconds
		//example: 21:23:32.16
		hours := rand.Intn(24)         // hours range from 0 to 23
		minutes := rand.Intn(60)       // minutes range from 0 to 59
		seconds := rand.Float64() * 60 // seconds range from 0 to 59.99
		value = fmt.Sprintf("%02d:%02d:%05.2f", hours, minutes, seconds)

	}

	//Declination
	if coordType == "dec" {
		// For Dec, generate degrees, arcminutes, and arcseconds
		//example: 53:01:36.08
		degrees := rand.Intn(181) - 90    // degrees range from -90 to 90
		arcminutes := rand.Intn(60)       // arcminutes range from 0 to 59
		arcseconds := rand.Float64() * 60 // arcseconds range from 0 to 59.99
		sign := "+"                       // positive sign for north
		if degrees < 0 {
			sign = "-"         // negative sign for south
			degrees = -degrees // absolute value of degrees
		}
		value = fmt.Sprintf("%s%02d:%02d:%05.2f", sign, degrees, arcminutes, arcseconds)
	}

	return value
}

//generate some random coordinates
//ra := "08:23:07.17" //for testing
//dec := "-48:29:40.53"

// end goal: build a valid query for this thing
// https://api.astrocats.space/catalog?ra=21:23:32.16&dec=-53:01:36.08&radius=400
var ExploreCmd = &cobra.Command{
	Use:   "explore",
	Short: "Picks a random spot in space and tells you what's there",
	Run: func(cmd *cobra.Command, args []string) {

		ra := newCoord("ra")
		dec := newCoord("dec")

		radius, err := cmd.Flags().GetInt("radius")
		if err != nil {
			fmt.Println("Error fetching radius flag:", err)
			return
		}

		navigate(ra, dec, radius)

	},
}

var WhatsHereCmd = &cobra.Command{
	Use:   "whatshere",
	Short: "Picks a random spot in space and tells you what's there",
	Run: func(cmd *cobra.Command, args []string) {

		ra, err := cmd.Flags().GetString("ra")
		if err != nil {
			fmt.Println("Error fetching ra flag:", err)
			return
		}

		dec, err := cmd.Flags().GetString("dec")
		if err != nil {
			fmt.Println("Error fetching dec flag:", err)
			return
		}

		radius, err := cmd.Flags().GetInt("radius")
		if err != nil {
			fmt.Println("Error fetching radius flag:", err)
			return
		}

		navigate(ra, dec, radius)

	},
}

func navigate(ra string, dec string, radius int) {
	fmt.Printf("\nCoords\nRA:\t%s\nDEC:\t%s\nRadius:\t%d\n\n", ra, dec, radius)

	// Build the query URL
	queryURL := fmt.Sprintf("https://api.astrocats.space/catalog?ra=%s&dec=%s&radius=%d",
		ra, dec, radius)

	//get return
	body := Fetch(queryURL)

	//PARSING
	var astrodata map[string]StellarObjectData

	if err := json.Unmarshal([]byte(body), &astrodata); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	tabulator(astrodata)
}

func tabulator(astrodata map[string]StellarObjectData) {

	// Create a tab writer
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintf(w, "Name\tRA\tDec\tDiscover Date\n")

	// Iterate through the data and print each entry as a row in the table
	for name, data := range astrodata {
		// Format the data as needed for the table
		ra := "n/a"
		if len(data.Ra) > 0 {
			ra = data.Ra[0].Value
		}

		dec := "n/a"
		if len(data.Dec) > 0 {
			dec = data.Dec[0].Value
		}

		discoverDate := "n/a"
		if len(data.Discoverdate) > 0 {
			discoverDate = data.Discoverdate[0].Value
		}

		// Print the row
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", name, ra, dec, discoverDate)
	}

	// Flush the tab writer
	w.Flush()

}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Provides information about RA and DEC",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\nRA (Right Ascension) and DEC (Declination) are celestial coordinate systems used to locate objects in the sky.")
		fmt.Println("\nThe expected format for RA is in hours, minutes, and seconds (HH:MM:SS.SS), and for DEC is in degrees, minutes, and seconds (DD:MM:SS.SS). For example, a valid RA value is \"08:23:07.17\", and a valid DEC value is \"-48:29:40.53\".")
		fmt.Println("\nMore reading: https://skyandtelescope.org/astronomy-resources/right-ascension-declination-celestial-coordinates/")
	},
}

func main() {

	ExploreCmd.Flags().IntP("radius", "r", 5000, "Specify the radius value")
	RootCmd.AddCommand(ExploreCmd)

	WhatsHereCmd.Flags().IntP("radius", "r", 5000, "Specify the radius value")
	WhatsHereCmd.Flags().StringP("ra", "a", "08:23:07.17", "Specify the ra (Right Ascension) value")
	WhatsHereCmd.Flags().StringP("dec", "d", "-48:29:40.53", "Specify the dec (Declination) value")
	RootCmd.AddCommand(WhatsHereCmd)

	RootCmd.AddCommand(aboutCmd)

	// Execute the CLI application
	Execute()
}
