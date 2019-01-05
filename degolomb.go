package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// BitArray Array of bits 0/1
func BitArray(baseTen byte) (bigEndianBits []int) {
	bigEndianBits = make([]int, 8)
	for i := range bigEndianBits {
		pv := byte(1 << uint8(len(bigEndianBits)-i-1))
		if (baseTen & pv) == pv {
			bigEndianBits[i] = 1
		}
	}

	return
}

// Degolomb Extract base10 values for bits in a Exp Golomb'd byte
func Degolomb(baseTen byte, fieldWidths []int) (baseTenValues []int) {
	if len(fieldWidths) == 0 {
		return []int{int(baseTen)}
	}
	bits := BitArray(baseTen)

	basePos := 0
	for _, width := range fieldWidths {
		tally := 0
		for shft := 0; shft < width; shft++ {
			bitIdx := basePos + shft
			bitVal := int(1 << uint8(width-shft-1))
			if bits[bitIdx] == 1 {
				tally += bitVal
			}
		}
		basePos += width
		baseTenValues = append(baseTenValues, tally)
	}
	return
}
