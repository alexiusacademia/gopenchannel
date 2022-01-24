package openchannel

type Unknown int

// Enums for unknowns.
// U stands for unknown.
const (
	UDischarge Unknown = iota
	UBedSlope
	UWaterDepth
	UChannelWidth
	UPipeDiameter
)

func (u Unknown) String() string {
	switch u {
	case UDischarge:
		return "Discharge"
	case UBedSlope:
		return "Bed Slope"
	case UWaterDepth:
		return "Water Depth"
	case UChannelWidth:
		return "Channel Width"
	case UPipeDiameter:
		return "Pipe Diameter"
	}

	return "Invalid unknown"
}
