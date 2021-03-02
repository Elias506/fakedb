package helpers

func IsType(elem string) bool {
	switch elem {
	case "string",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"int", "int8", "int16", "int32", "int64",
		"complex64", "complex128",
		"float32", "float64",
		"bool":
		return true
	}
	return false
}