package counter

type (
	// ICounter interface counter
	ICounter interface {
		Get() uint64
		GetLimit() uint64
		SetLimit(v uint64)
		Check() bool
	}
)
