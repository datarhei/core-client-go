package api

// Process represents all information on a process
type Process struct {
	ID        string         `json:"id" jsonschema:"minLength=1"`
	Type      string         `json:"type" jsonschema:"enum=ffmpeg"`
	Reference string         `json:"reference"`
	CreatedAt int64          `json:"created_at" jsonschema:"minimum=0"`
	Config    *ProcessConfig `json:"config,omitempty"`
	State     *ProcessState  `json:"state,omitempty"`
	Report    *ProcessReport `json:"report,omitempty"`
	Metadata  Metadata       `json:"metadata,omitempty"`
}

// ProcessConfigIO represents an input or output of an ffmpeg process config
type ProcessConfigIO struct {
	ID      string   `json:"id" validate:"required" jsonschema:"minLength=1"`
	Address string   `json:"address" validate:"required" jsonschema:"minLength=1"`
	Options []string `json:"options"`
}

// ProcessConfig represents the configuration of an ffmpeg process
type ProcessConfig struct {
	ID             string            `json:"id" validate:"required" jsonschema:"minLength=1"`
	Type           string            `json:"type" validate:"required" enums:"ffmpeg" jsonschema:"enum=ffmpeg"`
	Reference      string            `json:"reference"`
	Input          []ProcessConfigIO `json:"input" validate:"required"`
	Output         []ProcessConfigIO `json:"output" validate:"required"`
	Options        []string          `json:"options"`
	Reconnect      bool              `json:"reconnect"`
	ReconnectDelay int64             `json:"reconnect_delay_seconds" jsonschema:"minimum=1"`
	Autostart      bool              `json:"autostart"`
	StaleTimeout   int64             `json:"stale_timeout_seconds" jsonschema:"minimum=0"`
}

// ProcessReportHistoryEntry represents the logs of a run of an ffmpeg process
type ProcessReportHistoryEntry struct {
	CreatedAt int64       `json:"created_at"`
	Prelude   []string    `json:"prelude"`
	Log       [][2]string `json:"log"`
}

// ProcessReport represents the current log and the logs of previous runs of an ffmpeg process
type ProcessReport struct {
	ProcessReportHistoryEntry
	History []ProcessReportHistoryEntry `json:"history"`
}

// ProcessState represents the current state of an ffmpeg process
type ProcessState struct {
	Order     string    `json:"order" jsonschema:"enum=start,enum=stop"`
	State     string    `json:"exec" jsonschema:"enum=finished,enum=starting,enum=running,enum=finishing,enum=killed,enum=failed"`
	Runtime   int64     `json:"runtime_seconds" jsonschema:"minimum=0"`
	Reconnect int64     `json:"reconnect_seconds"`
	LastLog   string    `json:"last_logline"`
	Progress  *Progress `json:"progress"`
}
