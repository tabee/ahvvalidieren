//package main provide main functions to validate swiss ahv number.
package main

import (
	"errors"
	"strconv"
	"strings"

	"github.com/nicholassm/go-ean"
)

type ahvnr struct {
	number    string
	numberS   []string
	validated bool
}

func newAhvnr(input string) *ahvnr {
	a := ahvnr{number: input}
	a.numberS = strings.SplitN(input, ".", 4) // prepar inpu for checks.
	//a.validated = true
	return &a
}

// check country number for switzerland.
func validateCountry(s []string) bool {
	countryCode := "756" // country code 756 for switzerland
	if s[0] != countryCode {
		return false
	} else {
		return true
	}
}

// check ean13 checksum.
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

// run all checks.
func Validate(input string) (bool, error) {
	ahvnr := newAhvnr(input)
	statusCountry := validateCountry(ahvnr.numberS)
	if statusCountry {
		statusChecksum, _ := validateChecksum(ahvnr.numberS)
		if statusChecksum {
			ahvnr.validated = true
			return ahvnr.validated, nil
		}
	}
	return ahvnr.validated, errors.New("problem mit" + input)
}

func main() {
	arr := [2]string{"56.9217.0769.84", "756.3903.6825.80"}
	for j := 0; j < len(arr); j++ {
		v, err := Validate(arr[j])
		if err != nil {
			errors.New("Fehler")
		}
		println(v)
	}
}
