package bvg

import (
	"fmt"
	"strings"
)

type endpoint string
type endpointTemplate string

const (
	endpointTemplateStops              endpointTemplate = "/stops"
	endpointTemplateStopsReachableFrom endpointTemplate = "/stops/reachable-from"
	endpointTemplateStopsId            endpointTemplate = "/stops/:id"
	endpointTemplateStopsIdDepartures  endpointTemplate = "/stops/:id/departures"
	endpointTemplateStopsIdArrivals    endpointTemplate = "/stops/:id/arrivals"
	endpointTemplateJourneys           endpointTemplate = "/journeys"
	endpointTemplateJourneysRef        endpointTemplate = "/journeys/:ref"
	endpointTemplateTrips              endpointTemplate = "/trips"
	endpointTemplateTripsId            endpointTemplate = "/trips/:id"
	endpointTemplateLocationsNearby    endpointTemplate = "/locations/nearby"
	endpointTemplateLocations          endpointTemplate = "/locations"
	endpointTemplateRadar              endpointTemplate = "/radar"
)

func (e endpointTemplate) template(params ...map[string]string) endpoint {
	if len(params) > 0 {
		for key, value := range params[0] {
			e = endpointTemplate(strings.ReplaceAll(string(e), fmt.Sprintf(":%s", key), value))
		}
	}
	return endpoint(e)
}

func stopsReachableFrom() endpoint {
	return endpointTemplateStopsReachableFrom.template()
}

func stopsId(id string) endpoint {
	return endpointTemplateStopsId.template(map[string]string{"id": id})
}

func stopsIdDepartures(id string) endpoint {
	return endpointTemplateStopsIdDepartures.template(map[string]string{"id": id})
}

func stopsIdArrivals(id string) endpoint {
	return endpointTemplateStopsIdArrivals.template(map[string]string{"id": id})
}

func journeys() endpoint {
	return endpointTemplateJourneys.template()
}

func tripsId(id string) endpoint {
	return endpointTemplateTripsId.template(map[string]string{"id": id})
}

func trips() endpoint {
	return endpointTemplateTrips.template()
}

func locationsNearby() endpoint {
	return endpointTemplateLocationsNearby.template()
}

func locations() endpoint {
	return endpointTemplateLocations.template()
}

func radar() endpoint {
	return endpointTemplateRadar.template()
}

func journeysRef(ref string) endpoint {
	return endpointTemplateJourneysRef.template(map[string]string{"ref": ref})
}

func stops() endpoint {
	return endpointTemplateStops.template()
}
