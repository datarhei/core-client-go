package api

type Version struct {
	Number   string `json:"number"`
	Commit   string `json:"repository_commit"`
	Branch   string `json:"repository_branch"`
	Build    string `json:"build_date"`
	Arch     string `json:"arch"`
	Compiler string `json:"compiler"`
}
