package status

type Status uint8

const (
	Running = iota + 1
	Stopped
)

// String - return text os enum
func (s Status) String(status uint8) string {
	return [...]string{"Running", "Waiting"}[status-1]
}

// EnumIndex - return index of enum
func (s Status) EnumIndex() int {
	return int(s)
}
