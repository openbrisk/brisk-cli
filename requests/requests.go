package requests

// Function exported for system/functions endpoint
type Function struct {
	Name            string `json:"name"`
	Image           string `json:"image"`
	InvocationCount int64  `json:"invocationCount"`
	Replicas        uint64 `json:"replicas"`
	EnvProcess      string `json:"envProcess"`

	// Labels are metadata for functions which may be used by the
	// back-end for making scheduling or routing decisions.
	Labels *map[string]string `json:"labels"`
}
