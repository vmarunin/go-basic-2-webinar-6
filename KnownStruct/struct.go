package main

type SmartyStreetsAPI []struct {
	InputIndex           int        `json:"input_index,omitempty"`
	CandidateIndex       int        `json:"candidate_index,omitempty"`
	DeliveryLine1        string     `json:"delivery_line_1,omitempty"`
	LastLine             string     `json:"last_line,omitempty"`
	DeliveryPointBarcode string     `json:"delivery_point_barcode,omitempty"`
	Components           Components `json:"components,omitempty"`
	Metadata             Metadata   `json:"metadata,omitempty"`
	Analysis             Analysis   `json:"analysis,omitempty"`
}
type Components struct {
	PrimaryNumber           string
	StreetPredirection      string
	StreetName              string
	StreetSuffix            string
	CityName                string
	StateAbbreviation       string
	Zipcode                 string
	Plus4Code               string
	DeliveryPoint           string
	DeliveryPointCheckDigit string
}
type Metadata struct {
	RecordType            string
	ZipType               string
	CountyFips            string
	CountyName            string
	CarrierRoute          string
	CongressionalDistrict string
	Rdi                   string
	ElotSequence          string
	ElotSort              string
	Latitude              float64
	Longitude             float64
	Precision             string
	TimeZone              string
	UtcOffset             int
	Dst                   bool
}

type Analysis struct {
	DpvMatchCode string
	DpvFootnotes string
	DpvCmra      string
	DpvVacant    string
	Active       string
}
