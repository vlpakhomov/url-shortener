package encoder_test

import (
	"context"
	"testing"

	"github.com/VlPakhomov/url_shortener/internal/service/encoder"
	"github.com/VlPakhomov/url_shortener/pkg/validator"

	"github.com/stretchr/testify/assert"
)

const (
	lenShortUrlPattern = 10
	ShortUrlPattern    = "^[a-zA-Z0-9_]{10}$"
)

func TestEncode(t *testing.T) {
	t.Run("encode shortUrl according to algorithm", func(t *testing.T) {
		type testCase struct {
			token    int
			expected string
		}

		testCases := []testCase{
			{
				token:    1234567890,
				expected: "____aowuMS",
			},
			{
				token:    0,
				expected: "__________",
			},
		}
		for _, tc := range testCases {
			actual := encoder.Encode(tc.token)
			assert.Equal(t, tc.expected, actual)
		}
	})

	t.Run("check idempotent feature", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			assert.Equal(t, "____aowuMS", encoder.Encode(1234567890))
		}
	})

	t.Run("shortUrl consists of provided alphabet", func(t *testing.T) {
		for i := 0; i < 11; i++ {
			assert.True(t, validator.IsShortUrl(context.Background(), ShortUrlPattern, encoder.Encode(i)))
		}
	})

	t.Run("shortUrl has a specified length", func(t *testing.T) {
		for i := 0; i < 11; i++ {
			assert.Equal(t, 10, len(encoder.Encode(i)))
		}
	})
}
