package bvg

import "time"

type Location struct {
	Type      string   `json:"type"`
	ID        string   `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Latitude  float64  `json:"latitude,omitempty"`
	Longitude float64  `json:"longitude,omitempty"`
	POI       bool     `json:"poi,omitempty"`
	Address   string   `json:"address,omitempty"`
	Distance  int      `json:"distance,omitempty"`
	Products  Products `json:"products,omitempty"`
}

type Products struct {
	Suburban bool `json:"suburban"`
	Subway   bool `json:"subway"`
	Tram     bool `json:"tram"`
	Bus      bool `json:"bus"`
	Ferry    bool `json:"ferry"`
	Express  bool `json:"express"`
	Regional bool `json:"regional"`
}

type Line struct {
	Type     string   `json:"type"`
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Mode     string   `json:"mode"`
	Product  string   `json:"product"`
	Public   bool     `json:"public"`
	Operator Operator `json:"operator"`
	Symbol   string   `json:"symbol,omitempty"`
	Nr       int      `json:"nr,omitempty"`
}

type Operator struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Stopover struct {
	Stop              Location  `json:"stop"`
	Departure         time.Time `json:"departure,omitempty"`
	DepartureDelay    int       `json:"departureDelay,omitempty"`
	Arrival           time.Time `json:"arrival,omitempty"`
	ArrivalDelay      int       `json:"arrivalDelay,omitempty"`
	DeparturePlatform string    `json:"departurePlatform,omitempty"`
	ArrivalPlatform   string    `json:"arrivalPlatform,omitempty"`
}

type Remark struct {
	Type    string `json:"type"`
	Code    string `json:"code"`
	Text    string `json:"text"`
	Summary string `json:"summary,omitempty"`
}

type Leg struct {
	Origin            Location   `json:"origin"`
	Destination       Location   `json:"destination"`
	Departure         time.Time  `json:"departure"`
	DepartureDelay    int        `json:"departureDelay,omitempty"`
	Arrival           time.Time  `json:"arrival"`
	ArrivalDelay      int        `json:"arrivalDelay,omitempty"`
	DeparturePlatform string     `json:"departurePlatform,omitempty"`
	ArrivalPlatform   string     `json:"arrivalPlatform,omitempty"`
	Line              *Line      `json:"line,omitempty"`
	Direction         string     `json:"direction,omitempty"`
	Stopovers         []Stopover `json:"stopovers,omitempty"`
}

type Journey struct {
	Legs         []Leg  `json:"legs"`
	RefreshToken string `json:"refreshToken,omitempty"`
}

type JourneysResponse struct {
	Journeys              []Journey `json:"journeys"`
	EarlierRef            string    `json:"earlierRef,omitempty"`
	LaterRef              string    `json:"laterRef,omitempty"`
	RealtimeDataUpdatedAt int64     `json:"realtimeDataUpdatedAt,omitempty"`
}

type Trip struct {
	ID          string     `json:"id"`
	Line        Line       `json:"line"`
	Direction   string     `json:"direction"`
	Origin      Location   `json:"origin"`
	Destination Location   `json:"destination"`
	Departure   time.Time  `json:"departure"`
	Arrival     time.Time  `json:"arrival"`
	Stopovers   []Stopover `json:"stopovers,omitempty"`
	Remarks     []Remark   `json:"remarks,omitempty"`
}

type Departure struct {
	TripID      string    `json:"tripId"`
	Stop        Location  `json:"stop"`
	When        time.Time `json:"when"`
	PlannedWhen time.Time `json:"plannedWhen"`
	Delay       int       `json:"delay,omitempty"`
	Platform    string    `json:"platform,omitempty"`
	Line        Line      `json:"line"`
	Direction   string    `json:"direction"`
	Remarks     []Remark  `json:"remarks,omitempty"`
}

type Arrival struct {
	TripID      string    `json:"tripId"`
	Stop        Location  `json:"stop"`
	When        time.Time `json:"when"`
	PlannedWhen time.Time `json:"plannedWhen"`
	Delay       int       `json:"delay,omitempty"`
	Platform    string    `json:"platform,omitempty"`
	Line        Line      `json:"line"`
	Direction   string    `json:"direction"`
	Remarks     []Remark  `json:"remarks,omitempty"`
}

type RadarResponse struct {
	Movements []Movement `json:"movements"`
}

type Movement struct {
	Location      Location   `json:"location"`
	Line          Line       `json:"line"`
	Direction     string     `json:"direction"`
	TripID        string     `json:"tripId"`
	NextStopovers []Stopover `json:"nextStopovers,omitempty"`
}

type DeparturesResponse struct {
	Departures            []Departure `json:"departures"`
	RealtimeDataUpdatedAt int64       `json:"realtimeDataUpdatedAt,omitempty"`
}

type ArrivalsResponse struct {
	Arrivals              []Arrival `json:"arrivals"`
	RealtimeDataUpdatedAt int64     `json:"realtimeDataUpdatedAt,omitempty"`
}

type TripsResponse struct {
	Trips                 []Trip `json:"trips"`
	RealtimeDataUpdatedAt int64  `json:"realtimeDataUpdatedAt,omitempty"`
}

type ReachableResponse struct {
	Reachable []ReachableItem `json:"reachable"`
}

type ReachableItem struct {
	Duration int        `json:"duration"`
	Stations []Location `json:"stations"`
}
