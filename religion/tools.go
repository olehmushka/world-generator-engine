package religion

func PrepareCoef(in float64) float64 {
	if in < 0 {
		return 0
	}
	if in > 10 {
		return 10
	}

	return in
}
