package hash

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEqual(t *testing.T) {
	var (
		password = "lol-aza_za-lol123"
		cost     = 4
	)

	hasher := NewPasswordHasher(cost)

	hashedPassword, err := hasher.Hash(password)
	require.NoError(t, err)

	eq := hasher.Equals(hashedPassword, password)
	require.True(t, eq)
}
