package openchannel

type Unknown int64

const (
	Discharge Unknown = iota
	BedSlope
	WaterDepth
	ChannelWidth
	Radius
)

func (u Unknown) String() string {
	switch u {
	case Discharge:
		return "Discharge"
	case BedSlope:
		return "Bed Slope"
	case WaterDepth:
		return "Water Depth"
	case ChannelWidth:
		return "Channel Width"
	case Radius:
		return "Radius"
	}

	return "Invalid"
}
