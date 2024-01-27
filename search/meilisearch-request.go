package search

// MeiliSearchRequest struct in Go
type MeiliSearchRequest struct {
	Q                     string   `json:"q"`
	Offset                int      `json:"offset,omitempty"`
	Limit                 int      `json:"limit,omitempty"`
	HighlightPreTag       string   `json:"highlightPreTag"`
	HighlightPostTag      string   `json:"highlightPostTag"`
	ShowMatchesPosition   bool     `json:"showMatchesPosition"`
	Sort                  []string `json:"sort"`
	AttributesToHighlight []string `json:"attributesToHighlight"`
}

// NewSearchRequest initializes a new SearchRequest in Go
func NewSearchRequest() *MeiliSearchRequest {
	return &MeiliSearchRequest{}
}

// SetAttributesToHighlight sets the AttributesToHighlight field
func (s *MeiliSearchRequest) SetAttributesToHighlight(highlight []string) *MeiliSearchRequest {
	s.AttributesToHighlight = highlight
	return s
}

// SetShowMatchesPosition sets the ShowMatchesPosition field
func (s *MeiliSearchRequest) SetShowMatchesPosition(show bool) *MeiliSearchRequest {
	s.ShowMatchesPosition = show
	return s
}

// SetQ sets the Q field
func (s *MeiliSearchRequest) SetQ(q string) *MeiliSearchRequest {
	s.Q = q
	return s
}

// SetOffset sets the Offset field
func (s *MeiliSearchRequest) SetOffset(offset int) *MeiliSearchRequest {
	s.Offset = offset
	return s
}

// SetLimit sets the Limit field
func (s *MeiliSearchRequest) SetLimit(limit int) *MeiliSearchRequest {
	s.Limit = limit
	return s
}

// SetHighlightPreTag sets the HighlightPreTag field
func (s *MeiliSearchRequest) SetHighlightPreTag(highlightPreTag string) *MeiliSearchRequest {
	s.HighlightPreTag = highlightPreTag
	return s
}

// SetHighlightPostTag sets the HighlightPostTag field
func (s *MeiliSearchRequest) SetHighlightPostTag(highlightPostTag string) *MeiliSearchRequest {
	s.HighlightPostTag = highlightPostTag
	return s
}

// SetSort sets the Sort field
func (s *MeiliSearchRequest) SetSort(sort []string) *MeiliSearchRequest {
	s.Sort = sort
	return s
}

// Build returns the built SearchRequest
func (s *MeiliSearchRequest) Build() MeiliSearchRequest {
	return *s
}
