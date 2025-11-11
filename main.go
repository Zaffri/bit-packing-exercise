package main

import (
	"errors"
	"fmt"
	"os"
	"unsafe"
)

func main() {
	bools := [8]bool{true, false, true, false, false, false, false, false}
	packedBools, err := packBoolArray(bools)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Unpacked \nSize: %d bytes \nValue: %v\n\n", unsafe.Sizeof(bools), bools)
	fmt.Printf("Packed \nSize: %d bytes \nValue: %08b (%v)\n\n", unsafe.Sizeof(packedBools), packedBools, packedBools)

	unpackedBools := unpackBoolArray(packedBools)
	fmt.Printf("Original value after unpack:\n%v\n", unpackedBools)
}

func packBoolArray(unpackedBools [8]bool) (packedBools byte, err error) {
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

func unpackBoolArray(packedBools byte) (unpackedBools [8]bool) {
	for x := range 8 {
		mask := byte(1 << x)
		unpackedBools[x] = packedBools&mask != 0
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
