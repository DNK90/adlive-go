package token

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestJwtToken_CreateToken(t *testing.T) {
	token := &JwtToken{
		Id: "123",
		MacAddress: "abc",
	}
	err := token.Create()
	require.NoError(t, err)
	println(token)

	id, res := token.ValidateToken()
	require.True(t, res)
	require.Equal(t, "testId", id)
	println(id)
}
