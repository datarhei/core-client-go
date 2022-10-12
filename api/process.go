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
	ID      string                   `json:"id"`
	Address string                   `json:"address" validate:"required" jsonschema:"minLength=1"`
	Options []string                 `json:"options"`
	Cleanup []ProcessConfigIOCleanup `json:"cleanup,omitempty"`
}

type ProcessConfigIOCleanup struct {
	Pattern       string `json:"pattern" validate:"required"`
	MaxFiles      uint   `json:"max_files"`
	MaxFileAge    uint   `json:"max_file_age_seconds"`
	PurgeOnDelete bool   `json:"purge_on_delete"`
}

type ProcessConfigLimits struct {
	CPU     float64 `json:"cpu_usage" jsonschema:"minimum=0,maximum=100"`
	Memory  uint64  `json:"memory_mbytes" jsonschema:"minimum=0"`
	WaitFor uint64  `json:"waitfor_seconds" jsonschema:"minimum=0"`
}

// ProcessConfig represents the configuration of an ffmpeg process
type ProcessConfig struct {
	ID             string              `json:"id"`
	Type           string              `json:"type" validate:"oneof='ffmpeg' ''" jsonschema:"enum=ffmpeg,enum="`
	Reference      string              `json:"reference"`
	Input          []ProcessConfigIO   `json:"input" validate:"required"`
	Output         []ProcessConfigIO   `json:"output" validate:"required"`
	Options        []string            `json:"options"`
	Reconnect      bool                `json:"reconnect"`
	ReconnectDelay uint64              `json:"reconnect_delay_seconds"`
	Autostart      bool                `json:"autostart"`
	StaleTimeout   uint64              `json:"stale_timeout_seconds"`
	Limits         ProcessConfigLimits `json:"limits"`
}

// ProcessReportHistoryEntry represents the logs of a run of a restream process
type ProcessReportHistoryEntry struct {
	CreatedAt int64       `json:"created_at"`
	Prelude   []string    `json:"prelude"`
	Log       [][2]string `json:"log"`
}

// ProcessReport represents the current log and the logs of previous runs of a restream process
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
	Memory    uint64    `json:"memory_bytes"`                                            // bytes
	CPU       float64   `json:"cpu_usage" swaggertype:"number" jsonschema:"type=number"` // percent
	Command   []string  `json:"command"`
}
