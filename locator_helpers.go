package playwright

import (
	"regexp"
)

func convertRegexp(reg *regexp.Regexp) (pattern, flags string) {
	_ = "STUB: not implemented"
	return "", ""
}

func escapeForAttributeSelector(text any, exact bool) string { _ = "STUB: not implemented"; return "" }

func escapeForTextSelector(text any, exact bool) string { _ = "STUB: not implemented"; return "" }

func escapeRegexForSelector(re *regexp.Regexp) string { _ = "STUB: not implemented"; return "" }

func escapeText(s string) string { _ = "STUB: not implemented"; return "" }

func getByAltTextSelector(text any, exact bool) string { _ = "STUB: not implemented"; return "" }

func getByAttributeTextSelector(attrName string, text any, exact bool) string {
	_ = "STUB: not implemented"
	return ""
}

func getByLabelSelector(text any, exact bool) string { _ = "STUB: not implemented"; return "" }

func getByPlaceholderSelector(text any, exact bool) string { _ = "STUB: not implemented"; return "" }

func getByRoleSelector(role AriaRole, options ...LocatorGetByRoleOptions) string {
	_ = "STUB: not implemented"
	return ""
}

func getByTextSelector(text any, exact bool) string { _ = "STUB: not implemented"; return "" }

func getByTestIdSelector(testIdAttributeName string, testId any) string {
	_ = "STUB: not implemented"
	return ""
}

func getByTitleSelector(text any, exact bool) string { _ = "STUB: not implemented"; return "" }

func getTestIdAttributeName() string { _ = "STUB: not implemented"; return "" }

func setTestIdAttributeName(name string) { _ = "STUB: not implemented"; return }
