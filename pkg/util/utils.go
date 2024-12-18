package util

func SliceCheck(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// Not needed yet
// func SliceRemove(slice []string, item string) []string {
// 	for i, s := range slice {
// 		if s == item {
// 			return append(slice[:i], slice[i+1:]...)
// 		}
// 	}
// 	return slice
// }

// func SliceAppend(slice []string, item string) []string {
// 	if !SliceCheck(slice, item) {
// 		return append(slice, item)
// 	}
// 	return slice
// }
