package validator

import (
	"context"
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func IsUrl(_ context.Context, rawUrl string) bool {
	return govalidator.IsURL(rawUrl)
}

func IsShortUrl(_ context.Context, pattern, shortUrl string) bool {
	return govalidator.Matches(shortUrl, pattern)
}
