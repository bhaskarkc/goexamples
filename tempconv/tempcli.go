package tempconv

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func tempconvcli() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			log.Fatal(err)
		}

		f := Fahrenheit(t)
		c := Celsius(t)

		fmt.Printf("%s = %s, %s = %s\n",
			f, FToC(f),
			c, CToF(c),
		)
	}
}
