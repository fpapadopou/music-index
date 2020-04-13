package index

// Query specifies the query parameters.
type Query struct {
	Term string
	Type string
}

// Result represents an item returned from the SearchService.
type Result struct {
	Name         string
	Type         string
	Provider     string
	ProviderHRef string
	ImageURL     string
}

// Suggestion represents the output of the SuggestionsService.
type Suggestion struct {
	Name string
	Type string
}
