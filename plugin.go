package mtvmplugin

// Plugin is the plugin's functions
type Plugin interface {
	// Download starts the download
	Download(version string) error
	// Install executes after the download is finished
	Install(installDir string) error
	// GetLatestVersion returns the version number of the latest version
	GetLatestVersion() (string, error)
	// Progress returns a channel that you sent a float of the progress between 0 and 1, it's kind of like an initialization function
	Progress() chan float64
}
