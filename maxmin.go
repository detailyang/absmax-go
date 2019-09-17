package absmax

// MinInt64 returns the smaller of a or b.
func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// MinInt32 returns the smaller of a or b.
func MinInt32(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

// MinInt16 returns the smaller of a or b.
func MinInt16(a, b int16) int16 {
	if a < b {
		return a
	}
	return b
}

// MinInt8 returns the smaller of a or b.
func MinInt8(a, b int8) int8 {
	if a < b {
		return a
	}
	return b
}

// MinFloat64 returns the smaller of a or b.
func MinFloat64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// MinFloat32 returns the smaller of a or b.
func MinFloat32(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

// MinInt returns the smaller of a or b.
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt64 returns the larger of a or b.
func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// MaxInt32 returns the larger of a or b.
func MaxInt32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

// MaxInt16 returns the larger of a or b.
func MaxInt16(a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}

// MaxInt8 returns the larger of a or b.
func MaxInt8(a, b int8) int8 {
	if a > b {
		return a
	}
	return b
}

// MaxInt returns the larger of a or b.
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MaxFloat64 returns the larger of a or b.
func MaxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// MaxFloat32 returns the larger of a or b.
func MaxFloat32(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}
