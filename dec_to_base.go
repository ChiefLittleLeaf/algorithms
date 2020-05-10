package algorithms

func DecToBase(dec, base int) string {
	const charset = "0123456789ABCDEF"
	var res string
	for dec > 0 {
		rem := dec % base
		res = string(charset[rem]) + res
		dec = dec / base
	}
	return res

	// var sb strings.Builder
	//for dec > 0 {
	//		rem := dec % base
	//		sb.WriteByte(charset[rem])
	//		dec = dec / base
	//	}
	//	return res
}
