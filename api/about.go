package api

// About is some general information about the API
type About struct {
	Name      string  `json:"name"`
	ID        string  `json:"id"`
	CreatedAt string  `json:"created_at"`
	Version   Version `json:"version"`
}

// MinimalAbout is the minimal information about the API
type MinimalAbout struct {
	Name string `json:"name"`
}
