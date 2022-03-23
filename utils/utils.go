package utils

//RemoveIndex - remove de acordo determinanda posição do array
func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

//RemoveFirstPostion Remove a primeira posição do slice
func RemoveFirstPostion(s []string) []string {
	return RemoveIndex(s, 0)
}
