package status

type Status uint8

const (
	Close = iota + 1
	HalfOpen
	Open
)

// String - return text os enum
func String(status uint8) string {
	return [...]string{"Close", "Half-Open", "Open"}[status-1]
}

// EnumIndex - return index of enum
func (s Status) EnumIndex() int {
	return int(s)
}
