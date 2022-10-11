package api

type AVstreamIO struct {
	State  string `json:"state" enums:"running,idle" jsonschema:"enum=running,enum=idle"`
	Packet uint64 `json:"packet"`
	Time   uint64 `json:"time"`
	Size   uint64 `json:"size_kb"`
}

type AVstream struct {
	Input       AVstreamIO `json:"input"`
	Output      AVstreamIO `json:"output"`
	Aqueue      uint64     `json:"aqueue"`
	Queue       uint64     `json:"queue"`
	Dup         uint64     `json:"dup"`
	Drop        uint64     `json:"drop"`
	Enc         uint64     `json:"enc"`
	Looping     bool       `json:"looping"`
	Duplicating bool       `json:"duplicating"`
	GOP         string     `json:"gop"`
}
