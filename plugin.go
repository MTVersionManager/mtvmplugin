package plugin

// Plugin is the plugin itself. It must implement the Functions interface
type Plugin struct {
	// Channel that sends the download progress, number between 0 and 1, 0 is 0% and 1 is 100%
	Progress chan float64
	// The Data variable stores any information you might need to save, it is only accessed by you. A struct would be a good option for storing multiple variables
	Data any
}

// Functions is the functions that a Plugin must have to be valid
type Functions interface {
	// Download starts the download
	Download(version string) error
	// Install executes after the download is finished
	Install(installDir string) error
	// GetLatestVersion returns the version number of the latest version
	GetLatestVersion() (string, error)
}

// New initializes the plugin
type New func() (Plugin, error)
