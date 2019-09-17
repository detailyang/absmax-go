package absmax

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMin(t *testing.T) {
	require.Equal(t, -1, int(MinInt(-1, 1)))
	require.Equal(t, -1, int(MinInt32(-1, 1)))
	require.Equal(t, -1, int(MinInt64(-1, 1)))
	require.Equal(t, -1, int(MinInt16(-1, 1)))
	require.Equal(t, -1, int(MinInt8(-1, 1)))
	require.Equal(t, -1, int(MinFloat32(-1, 1)))
	require.Equal(t, -1, int(MinFloat64(-1, 1)))
}

func TestMax(t *testing.T) {
	require.Equal(t, 1, int(MaxInt(-1, 1)))
	require.Equal(t, 1, int(MaxInt32(-1, 1)))
	require.Equal(t, 1, int(MaxInt64(-1, 1)))
	require.Equal(t, 1, int(MaxInt16(-1, 1)))
	require.Equal(t, 1, int(MaxInt8(-1, 1)))
	require.Equal(t, 1, int(MaxFloat32(-1, 1)))
	require.Equal(t, 1, int(MaxFloat64(-1, 1)))
}
