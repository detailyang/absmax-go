package absmax

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestABSInt64(t *testing.T) {
	for i := -1234567; i <= 1234567; i++ {
		require.Equal(t, int64(math.Abs(float64(i))), AbsInt64(int64(i)))
	}
}

func TestABSInt32(t *testing.T) {
	for i := -123456; i <= 123456; i++ {
		require.Equal(t, int64(math.Abs(float64(i))), int64(AbsInt32(int32(i))))
	}
}

func TestABSInt16(t *testing.T) {
	for i := -12345; i <= 12345; i++ {
		require.Equal(t, int64(math.Abs(float64(i))), int64(AbsInt16(int16(i))))
	}
}

func TestABSInt8(t *testing.T) {
	for i := -123; i <= 123; i++ {
		require.Equal(t, int64(math.Abs(float64(i))), int64(AbsInt8(int8(i))))
	}
}
func TestABSFloat32(t *testing.T) {
	for i := -123456; i <= 123456; i++ {
		require.Equal(t, int64(math.Abs(float64(i))), int64(AbsFloat32(float32(i))))
	}
}

func TestABSFloat64(t *testing.T) {
	for i := -123456; i <= 123456; i++ {
		require.Equal(t, int64(math.Abs(float64(i))), int64(AbsFloat64(float64(i))))
	}
}
