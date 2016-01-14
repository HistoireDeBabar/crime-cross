package mocks

type MockStopAndSearchCollector struct {
}

func (m *MockStopAndSearchCollector) Collect() (data []byte, err error) {
	data = []byte("[{ \"removal_of_more_than_outer_clothing\" : \"null\", \"datetime\" : \"2015-04-01T17:30:00\", \"legislation\" : \"Misuse of Drugs Act 1971 (section 23)\", \"outcome\" : \"Suspect summonsed to court\", \"location\" : { \"latitude\" : \"51.284396\", \"longitude\" : \"-2.495092\", \"street\" : { \"id\" : 535912, \"name\" : \"On or near Longvernal\" } }, \"operation\" : null, \"self_defined_ethnicity\" : \"White - White British (W1)\", \"type\" : \"Person and Vehicle search\", \"age_range\" : \"25-34\", \"gender\" : \"Male\", \"operation_name\" : null, \"outcome_linked_to_object_of_search\" : null, \"object_of_search\" : \"Controlled drugs\", \"involved_person\" : true, \"officer_defined_ethnicity\" : \"White\" }]")
	return data, nil
}
