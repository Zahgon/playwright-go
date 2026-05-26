package playwright

import (
	"io"
	"log/slog"
	"os/exec"
)

const playwrightCliVersion = "1.57.0"

var (
	logger               = slog.Default()
	playwrightCDNMirrors = []string{
		"https://playwright.azureedge.net",
		"https://playwright-akamai.azureedge.net",
		"https://playwright-verizon.azureedge.net",
	}
)

// PlaywrightDriver wraps the Playwright CLI of upstream Playwright.
//
// It's required for playwright-go to work.
type PlaywrightDriver struct {
	Version string
	options *RunOptions
}

func NewDriver(options ...*RunOptions) (*PlaywrightDriver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get default values

func getDefaultCacheDirectory() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (d *PlaywrightDriver) isUpToDateDriver() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// avoid triggering downloads and accidentally overwriting files

// Command returns an exec.Cmd for the driver.
func (d *PlaywrightDriver) Command(arg ...string) *exec.Cmd { _ = "STUB: not implemented"; return nil }

// Install downloads the driver and the browsers depending on [RunOptions].
func (d *PlaywrightDriver) Install() error { _ = "STUB: not implemented"; return nil }

// Uninstall removes the driver and the browsers.
func (d *PlaywrightDriver) Uninstall() error { _ = "STUB: not implemented"; return nil }

// DownloadDriver downloads the driver only
func (d *PlaywrightDriver) DownloadDriver() error { _ = "STUB: not implemented"; return nil }

func (d *PlaywrightDriver) log(msg string, args ...any) { _ = "STUB: not implemented"; return }

func (d *PlaywrightDriver) run() (*connection, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *PlaywrightDriver) installBrowsers() error { _ = "STUB: not implemented"; return nil }

func (d *PlaywrightDriver) uninstallBrowsers() error { _ = "STUB: not implemented"; return nil }

// RunOptions are custom options to run the driver
type RunOptions struct {
	// DriverDirectory points to the playwright driver directory.
	// It should have two subdirectories: node and package.
	// You can also specify it using the environment variable PLAYWRIGHT_DRIVER_PATH.
	//
	// Default is user cache directory + "/ms-playwright-go/x.xx.xx":
	//  - Windows: %USERPROFILE%\AppData\Local
	//  - macOS: ~/Library/Caches
	//  - Linux: ~/.cache
	DriverDirectory string
	// OnlyInstallShell only downloads the headless shell. (For chromium browsers only)
	OnlyInstallShell    bool
	SkipInstallBrowsers bool
	// if not set and SkipInstallBrowsers is false, will download all browsers (chromium, firefox, webkit)
	Browsers []string
	Verbose  bool // default true
	Stdout   io.Writer
	Stderr   io.Writer
	Logger   *slog.Logger
	// DryRun does not install browser/dependencies. It will only print information.
	DryRun bool
}

// Install does download the driver and the browsers.
//
// Use this before playwright.Run() or use playwright cli to install the driver and browsers
func Install(options ...*RunOptions) error { _ = "STUB: not implemented"; return nil }

// Run starts a Playwright instance.
//
// Requires the driver and the browsers to be installed before.
// Either use Install() or use playwright cli.
func Run(options ...*RunOptions) (*Playwright, error) { _ = "STUB: not implemented"; return nil, nil }

func transformRunOptions(options ...*RunOptions) (*RunOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if user did not set it, try to get it from env

func getNodeExecutable(driverDirectory string) string { _ = "STUB: not implemented"; return "" }

func getDriverCliJs(driverDirectory string) string { _ = "STUB: not implemented"; return "" }

func (d *PlaywrightDriver) getDriverURLs() []string { _ = "STUB: not implemented"; return nil }

// isReleaseVersion checks if the version is not a beta or alpha release
// this helps to determine the url from where to download the driver
func (d *PlaywrightDriver) isReleaseVersion() bool { _ = "STUB: not implemented"; return false }

func makeFileExecutable(path string) error { _ = "STUB: not implemented"; return nil }

func downloadDriver(driverURLs []string) (body []byte, e error) {
	_ = "STUB: not implemented"
	return nil, nil
}
