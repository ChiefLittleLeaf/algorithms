package algorithms

func ReverseString(s string) string {
	//	var sb strings.Builder
	//	for i := len(s)-1; i >= 0; i--{
	//		sb.WriteByte(s[i])
	//	}
	//	return sb.String()

	//var res string
	//for i := len(s)-1; i>=0;i-- {
	//	res = res + string(s[i])
	//}
	//return res

	var res string
	for _, r := range s {
		res = string(r) + res
	}
	return res
}
