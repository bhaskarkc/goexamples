package tempconv

import (
	"fmt"
	"time"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (c Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", c)
}

// Celsius to Fahrenheit conversion.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// Fahrenheit to Celsius conversion.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	const day = 24 * time.Hour
	fmt.Println(day.Seconds())

	fmt.Printf("Absolute Zero! %v\n", AbsoluteZeroC)
	fmt.Printf("Freezing! %v\n", FreezingC)
	fmt.Printf("Boiling! %v\n", BoilingC)

	fmt.Printf("AbsoluteZero Fahrenheit! %v\n", CToF(AbsoluteZeroC))
	fmt.Printf("Freezing Fahrenheit! %v\n", CToF(FreezingC))
	fmt.Printf("Boiling Fahrenheit! %v\n", CToF(BoilingC))
}
