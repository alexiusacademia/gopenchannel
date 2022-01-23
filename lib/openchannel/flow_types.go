package openchannel

type FlowType int64

const (
	Critical FlowType = iota
	SubCritical
	SuperCritical
)

/*
func (ft FlowType) String() string {
	switch ft {
	case Critical:
		return "critical"
	case SubCritical:
		return "sub-critical"
	case SuperCritical:
		return "super critical"
	}
	return "unknown flow type"
}
*/
