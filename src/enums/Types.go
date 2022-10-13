package enums

type Type int32

const (
	TAG      Type = 0
	CATEGORY Type = 1
)

func (t Type) String() string {
	switch t {
	case TAG:
		return "tag"
	case CATEGORY:
		return "category"
	default:
		return "unknown"
	}
}
