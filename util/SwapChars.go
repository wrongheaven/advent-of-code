package util

func SwapChars(str string, i, j int) string {
	a := []byte(str)
	a[i], a[j] = a[j], a[i]
	return string(a)
}
