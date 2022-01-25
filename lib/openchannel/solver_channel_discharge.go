package openchannel

import (
	"errors"
	"math"
)

func SolveRectangularChannelDischarge(rc RectangularChannel) (ResultRectangularChannel, error) {
	wettedArea := rc.ChannelWidth * rc.WaterDepth
	wettedPerimeter := rc.ChannelWidth + (2 * rc.WaterDepth)
	hydraulicRadius := wettedArea / wettedPerimeter
	averageVelocity := (1.0 / rc.Roughness) * math.Pow(hydraulicRadius, (2.0/3.0)) * math.Sqrt(rc.BedSlope)
	discharge := averageVelocity * wettedArea

	var result ResultRectangularChannel
	result.Discharge = discharge
	result.BedSlope = rc.BedSlope
	result.ChannelWidth = rc.ChannelWidth
	result.WaterDepth = rc.WaterDepth

	result.HydraulicRadius = hydraulicRadius
	result.AverageVelocity = averageVelocity
	result.WettedPerimeter = wettedPerimeter
	result.WettedArea = wettedArea

	rcfc, err := solveForCriticalFlowRectangular(
		wettedArea,
		rc.ChannelWidth,
		averageVelocity,
		discharge,
		rc.Roughness)

	if err != nil {
		// Show error here
		return ResultRectangularChannel{}, errors.New("an error has occured")
	}

	result.HydraulicDepth = rcfc.HydraulicDepth
	result.FroudeNumber = rcfc.FroudeNumber
	result.DischargeIntensity = rcfc.DischargeIntensity
	result.CriticalDepth = rcfc.CriticalDepth
	result.CriticalTopWidth = rcfc.CriticalTopWidth
	result.CriticalWettedPerimeter = rcfc.CriticalWettedPerimeter
	result.CriticalWettedArea = rcfc.CriticalWettedArea
	result.CriticalHydraulicRadius = rcfc.CriticalHydraulicRadius
	result.CriticalSlope = rcfc.CriticalSlope

	return result, nil
}
