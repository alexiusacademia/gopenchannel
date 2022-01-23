package openchannel

type FlowType int64

const (
	Critical FlowType = iota
	SubCritical
	SuperCritical
)

func (ft FlowType) String() string {
	switch ft {
	case Critical:
		return "Critical"
	case SubCritical:
		return "Sub-Critical"
	case SuperCritical:
		return "Super Critical"
	}
	return "unknown flow type"
}
