package weightconv

import "fmt"

type Pound float64
type Kilogram float64

const (
	KiloInLb Kilogram = 0.4535924
	LbInKilo Pound    = 2.204623
)

func (p Pound) String() string    { return fmt.Sprintf("%g lb(s)", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%g kg(s)", k) }
