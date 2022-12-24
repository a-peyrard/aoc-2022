package collection

func Intersect[M1 ~map[K]struct{}, M2 ~map[K]struct{}, K comparable](first M1, second M2) map[K]struct{} {
	var exist bool
	intersection := map[K]struct{}{}
	for k := range first {
		_, exist = second[k]
		if exist {
			intersection[k] = struct{}{}
		}
	}
	return intersection
}
