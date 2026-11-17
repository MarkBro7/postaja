package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/MarkBro7/postaja/vreme" //to rabis zamenjat s tistim iz githuba

	"github.com/urfave/cli/v3"
)

func main() {

	var minTemp float64
	var maxTemp float64
	var stMeritev int

	cmd := &cli.Command{
		Name:  "vreme",
		Usage: "weather station pogram",

		Flags: []cli.Flag{
			&cli.Float64Flag{Name: "minTemp", Usage: "min allowed value", Value: -50.0},
			&cli.Float64Flag{Name: "maxTemp", Usage: "max allowed value", Value: 50.0},
			&cli.IntFlag{Name: "stMeritev", Usage: "min numbers of measurrments to perform statistic calculation", Value: 3},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp = cmd.Float64("minTemp")
			maxTemp = cmd.Float64("maxTemp")
			stMeritev = cmd.Int("stMeritev")
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

	// if --help or --h was parsed we display help section and end the program
	for _, arg := range os.Args {
		if arg == "--help" || arg == "-h" {
			return
		}
	}

	var postaje map[string]vreme.Postaja
	postaje = make(map[string]vreme.Postaja)

	//dodamo 4 poljubne postaje
	vreme.DodajPostajo(postaje, "LJ01", "Fri", "Ljubljana", []float64{21.4, 22.3, 19.9, 18.5})
	vreme.DodajPostajo(postaje, "KR01", "Primskovo", "Kranj", []float64{27.5, 18.9, 37.9})
	vreme.DodajPostajo(postaje, "MB01", "Center", "Maribor", []float64{1.3, 3.8, 5.9, -3.9, 4.7})
	vreme.DodajPostajo(postaje, "CE01", "Železniška postaja", "Celje", []float64{20.3})

	//testiranje funkcije dodajMeritev()
	fmt.Println("Testiranje funckije dodajeMeritev():")
	vreme.DodajMeritev(postaje, "CE01", 19.9, minTemp, maxTemp) //dodamo ustrezno vrednost -> ni izpisa
	vreme.DodajMeritev(postaje, "CE01", 100, minTemp, maxTemp)  //dodamo neustrezno vrednost -> izpis o napaki
	vreme.DodajMeritev(postaje, "KP01", 19.9, minTemp, maxTemp) //dodamo ključu ki ni v slovarju -> izpis o napaki

	//testiranje funkcije izpisMeritev()
	fmt.Println("\nTestiranje funckije izpisMeritev():")
	vreme.IzpisMeritev(postaje) //dobimo izpis

	//testiranje funkcije izpisPorocila()
	fmt.Println("\nTestiranje funckije izpisPorocila():")
	vreme.IzpisPorocila(postaje, stMeritev) //dobimo izpis
}
