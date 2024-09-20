package validator

import (
	"regexp"
	"strings"
)

func IsValidURL(target string) bool {
	var urlPattern = regexp.MustCompile(`^(https?:\/\/)?([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}(\/[\w\-\.~:\/?#[\]@!$&'()*+,;=%]*)?$`)

	return urlPattern.MatchString(target)
}

func EnsureScheme(target string) string {
	if !strings.HasPrefix(target, "http://") && !strings.HasPrefix(target, "https://") && !strings.HasPrefix(target, "localhost") {
		target = "https://" + target
	} else if strings.HasPrefix(target, "localhost") {
		target = "http://" + target
	}
	return target
}
