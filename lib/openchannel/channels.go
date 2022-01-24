package openchannel

import "github.com/alexiusacademia/gopenchannel/lib/math"

type ChannelType int

const (
	ChannelTypeRectangular ChannelType = iota
	ChannelTypeTrapezoidal
	ChannelTypeCircular
	ChannelTypeIrregular
)

func (ct ChannelType) String() string {
	switch ct {
	case ChannelTypeRectangular:
		return "Rectangular Channel"
	case ChannelTypeTrapezoidal:
		return "Trapezoidal Channel"
	case ChannelTypeCircular:
		return "Circular Channel"
	case ChannelTypeIrregular:
		return "Irregular Channel"
	}
	return "Unknown Channel"
}

type RectangularChannel struct {
	Unknown      Unknown
	Discharge    float64
	BedSlope     float64
	ChannelWidth float64
	WaterDepth   float64
	Roughness    float64
}

type TrapezoidalChannel struct {
	Unknown      Unknown
	Discharge    float64
	BedSlope     float64
	ChannelWidth float64
	WaterDepth   float64
	SideSlope    float64
	Roughness    float64
}

type CircularChannel struct {
	Unknown    Unknown
	Diameter   float64
	WaterDepth float64
	Roughness  float64
}

type IrregularChannel struct {
	Discharge      float64
	BedSlope       float64
	Nodes          []math.Node
	WaterElevation float64
	Roughness      float64
}
