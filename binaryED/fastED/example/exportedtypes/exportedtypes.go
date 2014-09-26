package exportedtypes

type (
	T1 struct {
		N1 int
		N2 string
		N3 uint8
		N4 *T2
	}

	T2 struct {
		N1 int
		N2 string
		N3 uint8
		N4 []string
		N5 map[int]string
		N6 map[int]*T3
	}

	T3 struct {
		N int
	}
)
