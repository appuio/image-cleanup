package cfg

type (
	// Configuration holds a strongly-typed tree of the configuration
	Configuration struct {
		Namespace string
		Git       GitConfig      `koanf:",squash"`
		History   HistoryConfig  `koanf:",squash"`
		Orphan    OrphanConfig   `koanf:",squash"`
		Resource  ResourceConfig `koanf:",squash"`
		Log       LogConfig
		Delete    bool
	}
	// GitConfig configures git repository
	GitConfig struct {
		CommitLimit  int    `koanf:"commit-limit"`
		RepoPath     string `koanf:"repo-path"`
		Tag          bool   `koanf:"tags"`
		SortCriteria string `koanf:"sort"`
	}
	// HistoryConfig configures the history command behaviour
	HistoryConfig struct {
		Keep int
	}
	// OrphanConfig configures the orphans command behaviour
	OrphanConfig struct {
		OlderThan           string `koanf:"older-than"`
		OrphanDeletionRegex string `koanf:"deletion-pattern"`
	}
	// LogConfig configures the log
	LogConfig struct {
		LogLevel string `koanf:"level"`
		Batch    bool
		Verbose  bool
	}
	// ResourceConfig configures the resources and secrets
	ResourceConfig struct {
		Labels      []string `koanf:"label"`
		OlderThan   string   `koanf:"older-than"`
		DeleteAfter string   `koanf:"delete-after"`
	}
)

// NewDefaultConfig retrieves the hardcoded configs with sane defaults
func NewDefaultConfig() *Configuration {
	return &Configuration{
		Git: GitConfig{
			CommitLimit:  0,
			RepoPath:     ".",
			Tag:          false,
			SortCriteria: "version",
		},
		History: HistoryConfig{
			Keep: 3,
		},
		Orphan: OrphanConfig{
			OlderThan:           "1w",
			OrphanDeletionRegex: "^[a-z0-9]{40}$",
		},
		Resource: ResourceConfig{
			Labels:      []string{},
			OlderThan:   "1w",
			DeleteAfter: "24h",
		},
		Delete: false,
		Log: LogConfig{
			LogLevel: "info",
			Batch:    false,
			Verbose:  false,
		},
	}
}
