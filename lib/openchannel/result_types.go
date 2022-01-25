package openchannel

// Result type for rectangular open channel.
type ResultRectangularChannel struct {
	Discharge    float64
	BedSlope     float64
	ChannelWidth float64
	WaterDepth   float64

	// Other properties
	HydraulicRadius float64
	AverageVelocity float64
	WettedPerimeter float64
	WettedArea      float64

	// Critical flow
	HydraulicDepth          float64
	FroudeNumber            float64
	DischargeIntensity      float64
	CriticalDepth           float64
	CriticalTopWidth        float64
	CriticalWettedPerimeter float64
	CriticalWettedArea      float64
	CriticalHydraulicRadius float64
	CriticalSlope           float64

	ChannelFlowType FlowType
}

// Result type for trapezoidal open channel.
type ResultTrapezoidalChannel struct {
	Discharge    float64
	BedSlope     float64
	ChannelWidth float64
	WaterDepth   float64
	SideSlope    float64

	// Other properties
	HydraulicRadius float64
	AverageVelocity float64
	WettedPerimeter float64
	WettedArea      float64

	// Critical flow
	HydraulicDepth          float64
	FroudeNumber            float64
	DischargeIntensity      float64
	CriticalDepth           float64
	CriticalTopWidth        float64
	CriticalWettedPerimeter float64
	CriticalWettedArea      float64
	CriticalHydraulicRadius float64
	CriticalSlope           float64

	ChannelFlowType FlowType
}

// Result type for circular open channel.
type ResultCircularChannel struct {
	Discharge    float64
	BedSlope     float64
	PipeDiameter float64
	WaterDepth   float64

	// Other properties
	HydraulicRadius float64
	AverageVelocity float64
	WettedPerimeter float64
	WettedArea      float64

	// Critical flow
	HydraulicDepth          float64
	FroudeNumber            float64
	DischargeIntensity      float64
	CriticalDepth           float64
	CriticalTopWidth        float64
	CriticalWettedPerimeter float64
	CriticalWettedArea      float64
	CriticalHydraulicRadius float64
	CriticalSlope           float64

	ChannelFlowType FlowType
}

// Result for critical flow calculation
type ResultCriticalFlowCalculation struct {
	HydraulicDepth          float64
	FroudeNumber            float64
	DischargeIntensity      float64
	CriticalDepth           float64
	CriticalTopWidth        float64
	CriticalWettedPerimeter float64
	CriticalWettedArea      float64
	CriticalHydraulicRadius float64
	CriticalSlope           float64
}
