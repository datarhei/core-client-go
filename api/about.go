package api

// About is some general information about the API
type About struct {
	App       string   `json:"app"`
	Auths     []string `json:"auths"`
	Name      string   `json:"name"`
	ID        string   `json:"id"`
	CreatedAt string   `json:"created_at"`
	Version   Version  `json:"version"`
}

// MinimalAbout is the minimal information about the API
type MinimalAbout struct {
	App   string   `json:"app"`
	Auths []string `json:"auths"`
}
