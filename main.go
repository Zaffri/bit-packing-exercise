package main

import (
	"errors"
	"fmt"
)

func main() {
	bools := []bool{true, false, true, false, false, false, false, false}
	fmt.Printf("Input: %v \n\n", bools)
	packedBools, err := packBoolArray(bools)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Packed")
		printByte(packedBools)
	}

	unpackedBools := unpackBoolArray(packedBools)

	fmt.Println("Input:  ", bools)
	fmt.Println("Result: ", unpackedBools)
}

func packBoolArray(unpackedBools []bool) (packedBools byte, err error) {
	if len(unpackedBools) > 8 {
		return 0, errors.New("There is a max size of 8 booleans for simplicity")
	}

	for index, bool := range unpackedBools {
		if bool {
			mask := byte(1 << index)
			packedBools = packedBools | mask
		}
	}
	return packedBools, nil
}

func unpackBoolArray(packedBools byte) (unpackedBools []bool) {
	for x := range 8 {
		mask := byte(1 << x)
		unpackedBools = append(unpackedBools, packedBools&mask != 0)
	}
	return unpackedBools
}

func printByte(byteToPrint byte) {
	fmt.Printf("Decimal: %d\n", byteToPrint)
	fmt.Printf("Binary: %08b\n\n", byteToPrint)
}

func byteTest(shiftLeft int) {
	fmt.Printf("Shift left %d times", shiftLeft)
	myWeeByte := byte(0)
	mask := byte(1 << shiftLeft)
	myWeeByte = myWeeByte | mask

	printByte(myWeeByte)
}
