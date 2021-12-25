package main

import (
	"errors"
	"strconv"
	"strings"

	"github.com/nicholassm/go-ean"
)

func validateChecksum(s []string) (bool, error) {

	c, err := ean.ChecksumEan13(s[0] + s[1] + s[2] + s[3])
	if err != nil {
		return false, errors.New("ean13 wrong")
	}
	// Prüfe Kontrollzahl
	value, _ := strconv.Atoi(s[3][1:])
	if value == c {
		return true, nil
	}
	return false, nil
}
func validateCountry(s []string, countryCode string) bool {
	if s[0] != countryCode {
		return false
	} else {
		return true
	}
}
func prepInp(ahvnr string) ([]string, error) {
	s := strings.SplitN(ahvnr, ".", 4)
	return s, nil
}

func Validate(ahvnr string) (bool, error) {
	a, _ := prepInp(ahvnr)
	statusCountry := validateCountry(a, "756")
	if statusCountry {
		statusChecksum, error := validateChecksum(a)
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
