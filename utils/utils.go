package utils

import (
	"fmt"
	"github.com/lucasjones/reggen"
)

func GenerateId() string {
	index, err := reggen.Generate("^[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{13}$", 36)
	if err != nil {
		fmt.Println(err)
	}

	return index
}