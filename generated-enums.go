package playwright

func getMixedState(in string) *MixedState { _ = "STUB: not implemented"; return nil }

type MixedState string

var (
	MixedStateOn    *MixedState = getMixedState("On")
	MixedStateOff               = getMixedState("Off")
	MixedStateMixed             = getMixedState("Mixed")
)

func getElementState(in string) *ElementState { _ = "STUB: not implemented"; return nil }

type ElementState string

var (
	ElementStateVisible  *ElementState = getElementState("visible")
	ElementStateHidden                 = getElementState("hidden")
	ElementStateStable                 = getElementState("stable")
	ElementStateEnabled                = getElementState("enabled")
	ElementStateDisabled               = getElementState("disabled")
	ElementStateEditable               = getElementState("editable")
)

func getAriaRole(in string) *AriaRole { _ = "STUB: not implemented"; return nil }

type AriaRole string

var (
	AriaRoleAlert            *AriaRole = getAriaRole("alert")
	AriaRoleAlertdialog                = getAriaRole("alertdialog")
	AriaRoleApplication                = getAriaRole("application")
	AriaRoleArticle                    = getAriaRole("article")
	AriaRoleBanner                     = getAriaRole("banner")
	AriaRoleBlockquote                 = getAriaRole("blockquote")
	AriaRoleButton                     = getAriaRole("button")
	AriaRoleCaption                    = getAriaRole("caption")
	AriaRoleCell                       = getAriaRole("cell")
	AriaRoleCheckbox                   = getAriaRole("checkbox")
	AriaRoleCode                       = getAriaRole("code")
	AriaRoleColumnheader               = getAriaRole("columnheader")
	AriaRoleCombobox                   = getAriaRole("combobox")
	AriaRoleComplementary              = getAriaRole("complementary")
	AriaRoleContentinfo                = getAriaRole("contentinfo")
	AriaRoleDefinition                 = getAriaRole("definition")
	AriaRoleDeletion                   = getAriaRole("deletion")
	AriaRoleDialog                     = getAriaRole("dialog")
	AriaRoleDirectory                  = getAriaRole("directory")
	AriaRoleDocument                   = getAriaRole("document")
	AriaRoleEmphasis                   = getAriaRole("emphasis")
	AriaRoleFeed                       = getAriaRole("feed")
	AriaRoleFigure                     = getAriaRole("figure")
	AriaRoleForm                       = getAriaRole("form")
	AriaRoleGeneric                    = getAriaRole("generic")
	AriaRoleGrid                       = getAriaRole("grid")
	AriaRoleGridcell                   = getAriaRole("gridcell")
	AriaRoleGroup                      = getAriaRole("group")
	AriaRoleHeading                    = getAriaRole("heading")
	AriaRoleImg                        = getAriaRole("img")
	AriaRoleInsertion                  = getAriaRole("insertion")
	AriaRoleLink                       = getAriaRole("link")
	AriaRoleList                       = getAriaRole("list")
	AriaRoleListbox                    = getAriaRole("listbox")
	AriaRoleListitem                   = getAriaRole("listitem")
	AriaRoleLog                        = getAriaRole("log")
	AriaRoleMain                       = getAriaRole("main")
	AriaRoleMarquee                    = getAriaRole("marquee")
	AriaRoleMath                       = getAriaRole("math")
	AriaRoleMeter                      = getAriaRole("meter")
	AriaRoleMenu                       = getAriaRole("menu")
	AriaRoleMenubar                    = getAriaRole("menubar")
	AriaRoleMenuitem                   = getAriaRole("menuitem")
	AriaRoleMenuitemcheckbox           = getAriaRole("menuitemcheckbox")
	AriaRoleMenuitemradio              = getAriaRole("menuitemradio")
	AriaRoleNavigation                 = getAriaRole("navigation")
	AriaRoleNone                       = getAriaRole("none")
	AriaRoleNote                       = getAriaRole("note")
	AriaRoleOption                     = getAriaRole("option")
	AriaRoleParagraph                  = getAriaRole("paragraph")
	AriaRolePresentation               = getAriaRole("presentation")
	AriaRoleProgressbar                = getAriaRole("progressbar")
	AriaRoleRadio                      = getAriaRole("radio")
	AriaRoleRadiogroup                 = getAriaRole("radiogroup")
	AriaRoleRegion                     = getAriaRole("region")
	AriaRoleRow                        = getAriaRole("row")
	AriaRoleRowgroup                   = getAriaRole("rowgroup")
	AriaRoleRowheader                  = getAriaRole("rowheader")
	AriaRoleScrollbar                  = getAriaRole("scrollbar")
	AriaRoleSearch                     = getAriaRole("search")
	AriaRoleSearchbox                  = getAriaRole("searchbox")
	AriaRoleSeparator                  = getAriaRole("separator")
	AriaRoleSlider                     = getAriaRole("slider")
	AriaRoleSpinbutton                 = getAriaRole("spinbutton")
	AriaRoleStatus                     = getAriaRole("status")
	AriaRoleStrong                     = getAriaRole("strong")
	AriaRoleSubscript                  = getAriaRole("subscript")
	AriaRoleSuperscript                = getAriaRole("superscript")
	AriaRoleSwitch                     = getAriaRole("switch")
	AriaRoleTab                        = getAriaRole("tab")
	AriaRoleTable                      = getAriaRole("table")
	AriaRoleTablist                    = getAriaRole("tablist")
	AriaRoleTabpanel                   = getAriaRole("tabpanel")
	AriaRoleTerm                       = getAriaRole("term")
	AriaRoleTextbox                    = getAriaRole("textbox")
	AriaRoleTime                       = getAriaRole("time")
	AriaRoleTimer                      = getAriaRole("timer")
	AriaRoleToolbar                    = getAriaRole("toolbar")
	AriaRoleTooltip                    = getAriaRole("tooltip")
	AriaRoleTree                       = getAriaRole("tree")
	AriaRoleTreegrid                   = getAriaRole("treegrid")
	AriaRoleTreeitem                   = getAriaRole("treeitem")
)

func getColorScheme(in string) *ColorScheme { _ = "STUB: not implemented"; return nil }

type ColorScheme string

var (
	ColorSchemeLight        *ColorScheme = getColorScheme("light")
	ColorSchemeDark                      = getColorScheme("dark")
	ColorSchemeNoPreference              = getColorScheme("no-preference")
	ColorSchemeNoOverride                = getColorScheme("no-override")
)

func getForcedColors(in string) *ForcedColors { _ = "STUB: not implemented"; return nil }

type ForcedColors string

var (
	ForcedColorsActive     *ForcedColors = getForcedColors("active")
	ForcedColorsNone                     = getForcedColors("none")
	ForcedColorsNoOverride               = getForcedColors("no-override")
)

func getHarContentPolicy(in string) *HarContentPolicy { _ = "STUB: not implemented"; return nil }

type HarContentPolicy string

var (
	HarContentPolicyOmit   *HarContentPolicy = getHarContentPolicy("omit")
	HarContentPolicyEmbed                    = getHarContentPolicy("embed")
	HarContentPolicyAttach                   = getHarContentPolicy("attach")
)

func getHarMode(in string) *HarMode { _ = "STUB: not implemented"; return nil }

type HarMode string

var (
	HarModeFull    *HarMode = getHarMode("full")
	HarModeMinimal          = getHarMode("minimal")
)

func getReducedMotion(in string) *ReducedMotion { _ = "STUB: not implemented"; return nil }

type ReducedMotion string

var (
	ReducedMotionReduce       *ReducedMotion = getReducedMotion("reduce")
	ReducedMotionNoPreference                = getReducedMotion("no-preference")
	ReducedMotionNoOverride                  = getReducedMotion("no-override")
)

func getServiceWorkerPolicy(in string) *ServiceWorkerPolicy { _ = "STUB: not implemented"; return nil }

type ServiceWorkerPolicy string

var (
	ServiceWorkerPolicyAllow *ServiceWorkerPolicy = getServiceWorkerPolicy("allow")
	ServiceWorkerPolicyBlock                      = getServiceWorkerPolicy("block")
)

func getSameSiteAttribute(in string) *SameSiteAttribute { _ = "STUB: not implemented"; return nil }

type SameSiteAttribute string

var (
	SameSiteAttributeStrict *SameSiteAttribute = getSameSiteAttribute("Strict")
	SameSiteAttributeLax                       = getSameSiteAttribute("Lax")
	SameSiteAttributeNone                      = getSameSiteAttribute("None")
)

func getHarNotFound(in string) *HarNotFound { _ = "STUB: not implemented"; return nil }

type HarNotFound string

var (
	HarNotFoundAbort    *HarNotFound = getHarNotFound("abort")
	HarNotFoundFallback              = getHarNotFound("fallback")
)

func getRouteFromHarUpdateContentPolicy(in string) *RouteFromHarUpdateContentPolicy {
	_ = "STUB: not implemented"
	return nil
}

type RouteFromHarUpdateContentPolicy string

var (
	RouteFromHarUpdateContentPolicyEmbed  *RouteFromHarUpdateContentPolicy = getRouteFromHarUpdateContentPolicy("embed")
	RouteFromHarUpdateContentPolicyAttach                                  = getRouteFromHarUpdateContentPolicy("attach")
)

func getUnrouteBehavior(in string) *UnrouteBehavior { _ = "STUB: not implemented"; return nil }

type UnrouteBehavior string

var (
	UnrouteBehaviorWait         *UnrouteBehavior = getUnrouteBehavior("wait")
	UnrouteBehaviorIgnoreErrors                  = getUnrouteBehavior("ignoreErrors")
	UnrouteBehaviorDefault                       = getUnrouteBehavior("default")
)

func getMouseButton(in string) *MouseButton { _ = "STUB: not implemented"; return nil }

type MouseButton string

var (
	MouseButtonLeft   *MouseButton = getMouseButton("left")
	MouseButtonRight               = getMouseButton("right")
	MouseButtonMiddle              = getMouseButton("middle")
)

func getKeyboardModifier(in string) *KeyboardModifier { _ = "STUB: not implemented"; return nil }

type KeyboardModifier string

var (
	KeyboardModifierAlt           *KeyboardModifier = getKeyboardModifier("Alt")
	KeyboardModifierControl                         = getKeyboardModifier("Control")
	KeyboardModifierControlOrMeta                   = getKeyboardModifier("ControlOrMeta")
	KeyboardModifierMeta                            = getKeyboardModifier("Meta")
	KeyboardModifierShift                           = getKeyboardModifier("Shift")
)

func getScreenshotAnimations(in string) *ScreenshotAnimations {
	_ = "STUB: not implemented"
	return nil
}

type ScreenshotAnimations string

var (
	ScreenshotAnimationsDisabled *ScreenshotAnimations = getScreenshotAnimations("disabled")
	ScreenshotAnimationsAllow                          = getScreenshotAnimations("allow")
)

func getScreenshotCaret(in string) *ScreenshotCaret { _ = "STUB: not implemented"; return nil }

type ScreenshotCaret string

var (
	ScreenshotCaretHide    *ScreenshotCaret = getScreenshotCaret("hide")
	ScreenshotCaretInitial                  = getScreenshotCaret("initial")
)

func getScreenshotScale(in string) *ScreenshotScale { _ = "STUB: not implemented"; return nil }

type ScreenshotScale string

var (
	ScreenshotScaleCss    *ScreenshotScale = getScreenshotScale("css")
	ScreenshotScaleDevice                  = getScreenshotScale("device")
)

func getScreenshotType(in string) *ScreenshotType { _ = "STUB: not implemented"; return nil }

type ScreenshotType string

var (
	ScreenshotTypePng  *ScreenshotType = getScreenshotType("png")
	ScreenshotTypeJpeg                 = getScreenshotType("jpeg")
)

func getWaitForSelectorState(in string) *WaitForSelectorState {
	_ = "STUB: not implemented"
	return nil
}

type WaitForSelectorState string

var (
	WaitForSelectorStateAttached *WaitForSelectorState = getWaitForSelectorState("attached")
	WaitForSelectorStateDetached                       = getWaitForSelectorState("detached")
	WaitForSelectorStateVisible                        = getWaitForSelectorState("visible")
	WaitForSelectorStateHidden                         = getWaitForSelectorState("hidden")
)

func getWaitUntilState(in string) *WaitUntilState { _ = "STUB: not implemented"; return nil }

type WaitUntilState string

var (
	WaitUntilStateLoad             *WaitUntilState = getWaitUntilState("load")
	WaitUntilStateDomcontentloaded                 = getWaitUntilState("domcontentloaded")
	WaitUntilStateNetworkidle                      = getWaitUntilState("networkidle")
	WaitUntilStateCommit                           = getWaitUntilState("commit")
)

func getLoadState(in string) *LoadState { _ = "STUB: not implemented"; return nil }

type LoadState string

var (
	LoadStateLoad             *LoadState = getLoadState("load")
	LoadStateDomcontentloaded            = getLoadState("domcontentloaded")
	LoadStateNetworkidle                 = getLoadState("networkidle")
)

func getContrast(in string) *Contrast { _ = "STUB: not implemented"; return nil }

type Contrast string

var (
	ContrastNoPreference *Contrast = getContrast("no-preference")
	ContrastMore                   = getContrast("more")
	ContrastNoOverride             = getContrast("no-override")
)

func getMedia(in string) *Media { _ = "STUB: not implemented"; return nil }

type Media string

var (
	MediaScreen     *Media = getMedia("screen")
	MediaPrint             = getMedia("print")
	MediaNoOverride        = getMedia("no-override")
)

func getHttpCredentialsSend(in string) *HttpCredentialsSend { _ = "STUB: not implemented"; return nil }

type HttpCredentialsSend string

var (
	HttpCredentialsSendUnauthorized *HttpCredentialsSend = getHttpCredentialsSend("unauthorized")
	HttpCredentialsSendAlways                            = getHttpCredentialsSend("always")
)
