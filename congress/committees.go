package congress

type Committee struct {
	Chamber           string `json:"chamber"`
	CommitteeId       string `json:"committee_id"`
	Name              string `json:"name"`
	ParentCommitteeId string `json:"parent_committee_id"`
	Subcommittee      bool   `json:"subcommittee"`
}

type CommitteeResult struct {
	Result
	Results []Committee `json:"results"`
}

func (c *Congress) GetCommittees(criteria map[string]string) (*CommitteeResult, error) {
	l := CommitteeResult{}
	err := c.GetURL(&l, congressRoot, criteria, "committees")
	if err != nil {
		return nil, err
	}
	return &l, nil
}
