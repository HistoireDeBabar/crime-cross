package main

type StopAndSearch struct {
	Legislation             string
	DateTime                string
	SelfDefinedEthnicity    string `json:"self_defined_ethnicity"`
	AgeRange                string `json:"age_range"`
	Type                    string
	Gender                  string
	OperationName           string `json:operation_name"`
	ObjectOfSearch          string `json:"out_of_search"`
	InvolvedPerson          bool   `json:"involved_person"`
	OfficerDefinedEthnicity string `json:"officer_defined_ethnicity"`
	Location                Location
}

type Location struct {
	Latitude  string
	Longitude string
	Street    Street
}

type Street struct {
	Id   int
	Name string
}
