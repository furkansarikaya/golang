package utils

func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}
func FullNameNakedReturn(f, l string) (full string, lenght int) {
	full = f + " " + l
	lenght = len(full)
	return
}
