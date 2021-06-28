package api

// ProgressIO represents the progress of an ffmpeg input or output
type ProgressIO struct {
	ID      string `json:"id" jsonschema:"minLength=1"`
	Address string `json:"address" jsonschema:"minLength=1"`

	// General
	Index   uint64  `json:"index"`
	Stream  uint64  `json:"stream"`
	Format  string  `json:"format"`
	Type    string  `json:"type"`
	Codec   string  `json:"codec"`
	Coder   string  `json:"coder"`
	Frame   uint64  `json:"frame"`
	FPS     float64 `json:"fps"`
	Packet  uint64  `json:"packet"`
	PPS     float64 `json:"pps"`
	Size    uint64  `json:"size_kb"`
	Bitrate float64 `json:"bitrate_kbit"`

	// Video
	Pixfmt    string  `json:"pix_fmt,omitempty"`
	Quantizer float64 `json:"q,omitempty"`
	Width     uint64  `json:"width,omitempty"`
	Height    uint64  `json:"height,omitempty"`

	// Audio
	Sampling uint64 `json:"sampling_hz,omitempty"`
	Layout   string `json:"layout,omitempty"`
	Channels uint64 `json:"channels,omitempty"`

	// avstream
	AVstream *AVstream `json:"avstream"`
}

// Progress represents the progress of an ffmpeg process
type Progress struct {
	Input     []ProgressIO `json:"inputs"`
	Output    []ProgressIO `json:"outputs"`
	Frame     uint64       `json:"frame"`
	Packet    uint64       `json:"packet"`
	FPS       float64      `json:"fps"`
	Quantizer float64      `json:"q"`
	Size      uint64       `json:"size_kb"`
	Time      float64      `json:"time"`
	Bitrate   float64      `json:"bitrate_kbit"`
	Speed     float64      `json:"speed"`
	Drop      uint64       `json:"drop"`
	Dup       uint64       `json:"dup"`
}
