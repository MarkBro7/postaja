// Package vreme provides some basic functions to add wheater stations with measurements of temperature, update stations,
// and to calculate some statistic based on the measurements in each station

package vreme

import "fmt"

// Postaja represent weather stations with name, place and slice of measurments
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo creates a new struct of type Postaja and adds it into the map that contains all other already create stations
func DodajPostajo(postaje map[string]Postaja, kljuc string, ime string, kraj string, meritve []float64) {
	postaje[kljuc] = Postaja{ime, kraj, meritve}
}

// DodajMeritev adds a measurment to the station unless the temperature is out of then range [minTemp, maxTemp], is stations doesnt exist, print
// with warning is produced
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64, minTemp float64, maxTemp float64) {
	if (temperatura > maxTemp) || (temperatura < minTemp) {
		fmt.Println("Temperatura je izven dovoljenega območja.")
		return
	}
	p, ok := postaje[oznaka]

	if ok == false {
		fmt.Println("Ta postaja ne obstaja.")
		return
	}

	p.meritve = append(p.meritve, temperatura)
	postaje[oznaka] = p
}

func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	p, rez := postaje[oznaka]

	if (rez == false) || (len(p.meritve) < stMeritev) {
		ok = false
		return
	} else {
		ok = true
	}

	najnizja = p.meritve[0]
	najvisja = p.meritve[0]

	for _, v := range p.meritve {
		povprecna += v
		if v < najnizja {
			najnizja = v
		}
		if v > najvisja {
			najvisja = v
		}
	}

	povprecna /= float64(len(p.meritve))
	return
}

// IzpisMeritev prints the elements of the struct Postaja alongside with all the measurments this station holds for all stations in the map
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")
	for k, v := range postaje {
		fmt.Printf("%s - %s (%s): [", k, v.ime, v.kraj)
		for i, m := range v.meritve {
			fmt.Printf("%.1f", m)
			if i+1 < len(v.meritve) {
				fmt.Printf(" ")
			}
		}
		fmt.Println("]")
	}
}

// IzpisPorocila prints name of the station and some statistics (mix, max, avg) and "Vroče, Toplo, Hladno, Mrzlo" based on the avg for all stations in the map,
// if stations has less then stMeritev of measurrements warning is printed instead
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	for k, v := range postaje {
		fmt.Printf("%s: ", v.ime)
		min, max, avg, ok := statistika(postaje, k, stMeritev)

		if !ok {
			fmt.Printf("premalo meritev\n")
		} else {
			fmt.Printf("min %.1f °C, max %.1f °C, povprecje %.1f °C -> ", min, max, avg)

			switch {
			case avg >= 25.0:
				fmt.Println("Vroče")
			case (avg < 25.0) && (avg >= 15.0):
				fmt.Println("Toplo")
			case (avg < 15.0) && (avg >= 5.0):
				fmt.Println("Hladno")
			case avg < 5.0:
				fmt.Println("Mrzlo")
			}
		}
	}
}
