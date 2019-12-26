package matrix

func areUniformedVecs(vecs []Vector) (int, bool) {
	if len(vecs) == 0 {
		return 0, true
	} else if len(vecs) == 1 {
		return len(vecs[0]), true
	}

	for ind := 1; ind < len(vecs); ind++ {
		if len(vecs[ind]) != len(vecs[ind -1 ]){
			return 0, false
		}
	}

	return len(vecs[0]), true
}