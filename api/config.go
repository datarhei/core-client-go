package api

import (
	"time"
)

// ConfigData embeds config.Data
type ConfigData struct {
	CreatedAt time.Time `json:"created_at"`
	LoadedAt  time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Version   int64     `json:"version" jsonschema:"minimum=1,maximum=1"`
	ID        string    `json:"id"`
	Address   string    `json:"address"`
	Log       struct {
		Level    string `json:"level" enums:"debug,info,warn,error,silent" jsonschema:"enum=debug,enum=info,enum=warn,enum=error,enum=silent"`
		MaxLines int    `json:"max_lines"`
	} `json:"log"`
	DB struct {
		Dir string `json:"dir"`
	} `json:"db"`
	Host struct {
		Name []string `json:"name"`
		Auto bool     `json:"auto"`
	} `json:"host"`
	API struct {
		Access struct {
			Allow []string `json:"allow"`
			Block []string `json:"block"`
		} `json:"access"`
		Auth struct {
			Enable           bool   `json:"enable"`
			DisableLocalhost bool   `json:"disable_localhost"`
			Username         string `json:"username"`
			Password         string `json:"password"`
			JWT              struct {
				Secret string `json:"secret"`
			} `json:"jwt"`
		} `json:"auth"`
	} `json:"api"`
	TLS struct {
		Address  string `json:"address"`
		Enable   bool   `json:"enable"`
		Auto     bool   `json:"auto"`
		CertFile string `json:"cert_file"`
		KeyFile  string `json:"key_file"`
	} `json:"tls"`
	Storage struct {
		Disk struct {
			Dir   string `json:"dir"`
			Size  int64  `json:"max_size_mbytes"`
			Cache struct {
				Enable   bool     `json:"enable"`
				Size     uint64   `json:"max_size_mbytes"`
				TTL      int64    `json:"ttl_seconds"`
				FileSize uint64   `json:"max_file_size_mbytes"`
				Types    []string `json:"types"`
			} `json:"cache"`
		} `json:"disk"`
		Memory struct {
			Auth struct {
				Enable   bool   `json:"enable"`
				Username string `json:"username"`
				Password string `json:"password"`
			} `json:"auth"`
			Size  int64 `json:"max_size_mbytes"`
			Purge bool  `json:"purge"`
		} `json:"memory"`
		CORS struct {
			Origins []string `json:"origins"`
		} `json:"cors"`
		MimeTypes string `json:"mimetypes_file"`
	} `json:"storage"`
	RTMP struct {
		Enable    bool   `json:"enable"`
		EnableTLS bool   `json:"enable_tls"`
		Address   string `json:"address"`
		App       string `json:"app"`
		Token     string `json:"token"`
	} `json:"rtmp"`
	FFmpeg struct {
		Binary       string `json:"binary"`
		MaxProcesses int64  `json:"max_processes"`
		Access       struct {
			Input struct {
				Allow []string `json:"allow"`
				Block []string `json:"block"`
			} `json:"input"`
			Output struct {
				Allow []string `json:"allow"`
				Block []string `json:"block"`
			} `json:"output"`
		} `json:"access"`
		Log struct {
			MaxLines   int `json:"max_lines"`
			MaxHistory int `json:"max_history"`
		} `json:"log"`
	} `json:"ffmpeg"`
	Playout struct {
		Enable  bool `json:"enable"`
		MinPort int  `json:"min_port"`
		MaxPort int  `json:"max_port"`
	} `json:"playout"`
	Debug struct {
		Profiling bool `json:"profiling"`
		ForceGC   int  `json:"force_gc"`
	} `json:"debug"`
	Stats struct {
		Enable          bool     `json:"enable"`
		IPIgnoreList    []string `json:"ip_ignorelist"`
		SessionTimeout  int      `json:"session_timeout_sec"`
		Persist         bool     `json:"persist"`
		PersistInterval int      `json:"persist_interval_sec"`
		MaxBitrate      uint64   `json:"max_bitrate_mbit"`
		MaxSessions     uint64   `json:"max_sessions"`
	} `json:"stats"`
	Service struct {
		Enable bool   `json:"enable"`
		Token  string `json:"token"`
		URL    string `json:"url"`
	} `json:"service"`
	Router struct {
		BlockedPrefixes []string          `json:"blocked_prefixes"`
		Routes          map[string]string `json:"routes"`
	} `json:"router"`
}

// Config is the config returned by the API
type Config struct {
	CreatedAt time.Time `json:"created_at"`
	LoadedAt  time.Time `json:"loaded_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Config ConfigData `json:"config"`

	Overrides []string `json:"overrides"`
}

// ConfigError is used to return error messages when uploading a new config
type ConfigError map[string][]string
