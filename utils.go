package thingfulx

import (
	"regexp"
)

var (
	// Collection of regexps used for our quick nasty categorise function.
	experimentRegexp = regexp.MustCompile("arduino|sunspot|mbed|beaglebone|linux|computer|raspberry|experiment")

	energyRegexp = regexp.MustCompile("thermostat|energy|electricity|carbon|gas|power|meter|currentcost|cc128|solar|efergy|watt|current cost")

	healthRegexp = regexp.MustCompile("weight|heart|fitbit|medical|exercise|health")

	homeRegexp = regexp.MustCompile("room|house|home|conservatory|indoor|maison|curtain|kitchen|dining|garage|garden")

	floraRegexp = regexp.MustCompile("animal|shark|turtle|tree|bird|cat|dog|pet|plant|botanic|biologic|soil|agriculture|flora|fauna")

	environmentRegexp = regexp.MustCompile("airquality|air quality|environment|weather|pollution|outdoor|temperature|pm10|iceberg|earthquake|soil|atm|flood")

	transportRegexp = regexp.MustCompile("ship|vehicle|buoy|vessel|truck|transport|traffic|car|automobile|transport")
)

// Categorise is an extract of Thingful's original quick'n'nasty categorisation
// logic.
func Categorise(str string) Category {
	switch {
	case experimentRegexp.MatchString(str):
		return Experiment
	case energyRegexp.MatchString(str):
		return Energy
	case healthRegexp.MatchString(str):
		return Health
	case homeRegexp.MatchString(str):
		return Home
	case floraRegexp.MatchString(str):
		return Flora
	case environmentRegexp.MatchString(str):
		return Environment
	case transportRegexp.MatchString(str):
		return Transport
	default:
		return Miscellaneous
	}
}
