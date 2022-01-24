package openchannel

type FlowType int64

// Enums for flow types.
// FT stands for flow type.
const (
	FTCritical FlowType = iota
	FTSubCritical
	FTSuperCritical
)

func (ft FlowType) String() string {
	switch ft {
	case FTCritical:
		return "Critical"
	case FTSubCritical:
		return "Sub-Critical"
	case FTSuperCritical:
		return "Super Critical"
	}
	return "Unknown flow type"
}
