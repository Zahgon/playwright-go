package playwright

import (
	"regexp"
)

var escapedChars = map[rune]bool{
	'$':  true,
	'^':  true,
	'+':  true,
	'.':  true,
	'*':  true,
	'(':  true,
	')':  true,
	'|':  true,
	'\\': true,
	'?':  true,
	'{':  true,
	'}':  true,
	'[':  true,
	']':  true,
}

func globMustToRegex(glob string) *regexp.Regexp { _ = "STUB: not implemented"; return nil }

func resolveGlobToRegex(baseURL *string, glob string, isWebSocketUrl bool) *regexp.Regexp {
	_ = "STUB: not implemented"
	return nil
}

func resolveGlobBase(baseURL *string, match string) string { _ = "STUB: not implemented"; return "" }

// Escaped `\\?` behaves the same as `?` in our glob patterns.

// Glob symbols may be escaped in the URL and some of them such as ? affect resolution,
// so we replace them with safe components first.

// Handle special case of http*://, note that the new schema has to be
// a web schema so that slashes are properly inserted after domain.

func constructURLBasedOnBaseURL(baseURL *string, givenURL string) string {
	_ = "STUB: not implemented"
	return ""
}

// In Node.js, new URL('http://localhost') returns 'http://localhost/'.

func toWebSocketBaseURL(baseURL *string) *string { _ = "STUB: not implemented"; return nil }

// Allow http(s) baseURL to match ws(s) urls.
