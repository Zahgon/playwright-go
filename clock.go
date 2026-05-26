package playwright

type clockImpl struct {
	browserCtx *browserContextImpl
}

func newClock(bCtx *browserContextImpl) Clock { _ = "STUB: not implemented"; return *new(Clock) }

func (c *clockImpl) FastForward(ticks any) error { _ = "STUB: not implemented"; return nil }

func (c *clockImpl) Install(options ...ClockInstallOptions) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *clockImpl) PauseAt(time any) error { _ = "STUB: not implemented"; return nil }

func (c *clockImpl) Resume() error { _ = "STUB: not implemented"; return nil }

func (c *clockImpl) RunFor(ticks any) error { _ = "STUB: not implemented"; return nil }

func (c *clockImpl) SetFixedTime(time any) error { _ = "STUB: not implemented"; return nil }

func (c *clockImpl) SetSystemTime(time any) error { _ = "STUB: not implemented"; return nil }

func parseTime(t any) (map[string]any, error) { _ = "STUB: not implemented"; return nil, nil }

func parseTicks(ticks any) (map[string]any, error) { _ = "STUB: not implemented"; return nil, nil }
