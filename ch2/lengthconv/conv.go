// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package lengthconv

// FToM converts feet to meters
func FToM(f Feet) Meter { return Meter(f) * MeterInF }

// MToF converts Meters to Feet
func MToF(m Meter) Feet { return Feet(m) / FeetInM }
