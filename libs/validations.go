package libs


func MinSize(field string, size int) (result bool) {
	if len(field) < size {
		return false
	}
	return true
}

func MaxSize(field string, size int) (result bool) {
	if len(field) > size {
		return false
	}
	return true
}