// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package weightconv

// KToP converts kilos to lbs
func KToP(k Kilogram) Pound { return Pound(k) * LbInKilo }

// PToK converts pounds to kilos
func PToK(p Pound) Kilogram { return Kilogram(p) * KiloInLb }
