package api

type PlayoutStatusIO struct {
	State  string `json:"state" enums:"running,idle" jsonschema:"enum=running,enum=idle"`
	Packet uint64 `json:"packet"`
	Time   uint64 `json:"time"`
	Size   uint64 `json:"size_kb"`
}

type PlayoutStatusSwap struct {
	Address     string `json:"url"`
	Status      string `json:"status"`
	LastAddress string `json:"lasturl"`
	LastError   string `json:"lasterror"`
}

type PlayoutStatus struct {
	ID          string            `json:"id"`
	Address     string            `json:"url"`
	Stream      uint64            `json:"stream"`
	Queue       uint64            `json:"queue"`
	AQueue      uint64            `json:"aqueue"`
	Dup         uint64            `json:"dup"`
	Drop        uint64            `json:"drop"`
	Enc         uint64            `json:"enc"`
	Looping     bool              `json:"looping"`
	Duplicating bool              `json:"duplicating"`
	GOP         string            `json:"gop"`
	Debug       interface{}       `json:"debug"`
	Input       PlayoutStatusIO   `json:"input"`
	Output      PlayoutStatusIO   `json:"output"`
	Swap        PlayoutStatusSwap `json:"swap"`
}
