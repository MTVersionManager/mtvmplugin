package mtvmplugin

// Plugin is the plugin's functions
type Plugin interface {
	// Download starts the download, please send the download progress through the progress channel, send 0 for 0% and 1 for 100%
	Download(version string, progress chan float64) error
	// Install executes after the download is finished
	Install(installDir string) error
	// GetLatestVersion returns the version number of the latest version
	GetLatestVersion() (string, error)
	// Sort returns a sorted slice of versions sorted in ascending order (oldest first) and an error.
	Sort(versions []string) ([]string, error)
}
