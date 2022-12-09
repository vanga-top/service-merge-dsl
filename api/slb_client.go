package api

// SLBClient Service Load Balance Client
type SLBClient interface {
	Connect(url string, port int, namespace string, opts *SLBOptions) *SLBResult
	DisConnect() *SLBResult
}

type SLBResult struct {
	Code      int
	Success   bool
	Namespace string
}

type SLBOptions struct {
	Username string
	Password string
	Token    string
}
