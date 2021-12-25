package main

import (
	"errors"
	"strconv"
	"strings"

	"github.com/nicholassm/go-ean"
)

func validateChecksum(s string) (bool, error) {
	string := strings.SplitN(s, ".", 4)
	c, err := ean.ChecksumEan13(string[0] + string[1] + string[2] + string[3])
	if err != nil {
		return false, errors.New("ean13 wrong")
	}
	// Prüfe Kontrollzahl
	value, _ := strconv.Atoi(string[3][1:])
	if value == c {
		return true, nil
	}
	return false, nil
}

func validateCountry(s string, countryCode string) bool {
	string := strings.SplitN(s, ".", 4)
	if string[0] != countryCode {
		return false
	} else {
		return true
	}
}

func Validate(ahvnr string) (bool, error) {
	statusCountry := validateCountry(ahvnr, "756")
	if statusCountry {
		statusChecksum, error := validateChecksum(ahvnr)
		if statusChecksum {
			return true, nil
		}
		if error != nil {
			return false, errors.New("problem with checksum" + ahvnr)
		}
	}
	return false, errors.New("problem" + ahvnr)
}

func main() {
	println(Validate("756.9217.0769.84")) // false
	println(Validate("756.3903.6825.80")) // true
	//println(Validate("746.3903.6825.80")) // false
	//println(Validate("756.3903.6445.81")) // false

}
