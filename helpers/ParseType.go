package helpers

import (
	"strconv"
	"fmt"
)

//ParseType checking on correct type in s (string)
func ParseType(s string, sType string) (interface{}, error) {
	var i interface{} = nil
	var err error = nil
	switch sType {
	case "int":
		i, err = strconv.ParseInt(s, 0, 0)
	case "int8":
		i, err = strconv.ParseInt(s, 0, 8)
	case "int16":
		i, err = strconv.ParseInt(s, 0, 16)
	case "int32":
		i, err = strconv.ParseInt(s, 0, 32)
	case "int64":
		i, err = strconv.ParseInt(s, 0, 64)
	case "uint":
		i, err = strconv.ParseUint(s, 0, 0)
	case "uint8":
		i, err = strconv.ParseUint(s, 0, 8)
	case "uint16":
		i, err = strconv.ParseUint(s, 0, 16)
	case "uint32":
		i, err = strconv.ParseUint(s, 0, 32)
	case "uint64":
		i, err = strconv.ParseUint(s, 0, 64)
	case "string":
		if len(s) >= 2 {
			if s[0] == '"' && s[len(s)-1] == '"' {
				i, err = s, nil
			} else {
				i, err = nil, fmt.Errorf(`invalid syntax - %s; use " before and after string value`, s)
			}
		} else {
			i, err = nil, fmt.Errorf(`invalid syntax - %s; use " before and after string value`, s)
		}
	case "bool":
		i, err = strconv.ParseBool(s)
	case "complex64":
		i, err = strconv.ParseComplex(s, 64)
	case "complex128":
		i, err = strconv.ParseComplex(s, 128)
	case "float32":
		i, err = strconv.ParseFloat(s, 32)
	case "float64":
		i, err = strconv.ParseFloat(s, 64)
	}
	return i, err
}
