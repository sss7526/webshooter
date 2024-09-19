package validator

import (
	"regexp"
)

func IsValidURL(target string) bool {
	var urlPattern = regexp.MustCompile(`^(https?:\/\/)?([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}(\/[\w\-\.~:\/?#[\]@!$&'()*+,;=%]*)?$`)

	return urlPattern.MatchString(target)
}
