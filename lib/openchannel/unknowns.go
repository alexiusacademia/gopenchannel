package openchannel

type Unknown int64

const (
	Discharge Unknown = iota
	BedSlope
	WaterDepth
	ChannelWidth
	Diameter
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
	case Diameter:
		return "Pipe Diameter"
	}

	return "Invalid"
}
