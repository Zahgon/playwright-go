package playwright

type harRouter struct {
	localUtils     *localUtilsImpl
	harId          string
	notFoundAction HarNotFound
	urlOrPredicate any
	err            error
}

func (r *harRouter) addContextRoute(context BrowserContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *harRouter) addPageRoute(page Page) error { _ = "STUB: not implemented"; return nil }

func (r *harRouter) dispose() { _ = "STUB: not implemented"; return }

func (r *harRouter) handle(route Route) error { _ = "STUB: not implemented"; return nil }

func newHarRouter(localUtils *localUtilsImpl, file string, notFoundAction HarNotFound, urlOrPredicate any) *harRouter {
	_ = "STUB: not implemented"
	return nil
}
