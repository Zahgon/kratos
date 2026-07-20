package http

type redirect struct {
	URL  string
	Code int
}

func (r *redirect) Redirect() (string, int) { _ = "STUB: not implemented"; return "", 0 }

func (r *redirect) Error() string { _ = "STUB: not implemented"; return "" }

func NewRedirect(url string, code int) Redirector {
	_ = "STUB: not implemented"
	return *new(Redirector)
}
