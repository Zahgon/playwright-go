package playwright

const assertionsDefaultTimeout = 5000 // 5s

type playwrightAssertionsImpl struct {
	defaultTimeout *float64
}

// NewPlaywrightAssertions creates a new instance of PlaywrightAssertions
//   - timeout: default value is 5000 (ms)
func NewPlaywrightAssertions(timeout ...float64) PlaywrightAssertions {
	_ = "STUB: not implemented"
	return *new(PlaywrightAssertions)
}

func (pa *playwrightAssertionsImpl) APIResponse(response APIResponse) APIResponseAssertions {
	_ = "STUB: not implemented"
	return *new(APIResponseAssertions)
}

func (pa *playwrightAssertionsImpl) Locator(locator Locator) LocatorAssertions {
	_ = "STUB: not implemented"
	return *new(LocatorAssertions)
}

func (pa *playwrightAssertionsImpl) Page(page Page) PageAssertions {
	_ = "STUB: not implemented"
	return *new(PageAssertions)
}

type expectedTextValue struct {
	Str                 *string `json:"string,omitempty"`
	RegexSource         *string `json:"regexSource,omitempty"`
	RegexFlags          *string `json:"regexFlags,omitempty"`
	MatchSubstring      *bool   `json:"matchSubstring,omitempty"`
	IgnoreCase          *bool   `json:"ignoreCase,omitempty"`
	NormalizeWhiteSpace *bool   `json:"normalizeWhiteSpace,omitempty"`
}

type frameExpectOptions struct {
	ExpressionArg  any                 `json:"expressionArg,omitempty"`
	ExpectedText   []expectedTextValue `json:"expectedText,omitempty"`
	ExpectedNumber *float64            `json:"expectedNumber,omitempty"`
	ExpectedValue  any                 `json:"expectedValue,omitempty"`
	UseInnerText   *bool               `json:"useInnerText,omitempty"`
	IsNot          bool                `json:"isNot"`
	Timeout        *float64            `json:"timeout"`
}

type frameExpectResult struct {
	Matches  bool     `json:"matches"`
	Received any      `json:"received,omitempty"`
	TimedOut *bool    `json:"timedOut,omitempty"`
	Log      []string `json:"log,omitempty"`
}

type assertionsBase struct {
	actualLocator  Locator
	isNot          bool
	defaultTimeout *float64
}

func (b *assertionsBase) expect(
	expression string,
	options frameExpectOptions,
	expected any,
	message string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func toExpectedTextValues(
	items []any,
	matchSubstring bool,
	normalizeWhiteSpace bool,
	ignoreCase *bool,
) ([]expectedTextValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertToInterfaceList(v any) []any { _ = "STUB: not implemented"; return nil }
