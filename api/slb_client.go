package api

// SLBClient Service Load Balance Client
type SLBClient interface {
	Connect(host string, port int, namespace string, token string) *SLBResult
	DisConnect() *SLBResult
}

type SLBResult struct {
	Code      int
	Success   bool
	Namespace string
}
