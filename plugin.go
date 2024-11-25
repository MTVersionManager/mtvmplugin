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
	// Use sets the current version. A path directory is provided for you to symlink or copy the binaries into (it is called path directory because it should be part of the system's path variable)
	// An installDir is also provided which is where the version is installed.
	Use(installDir string, pathDir string) error
	// GetCurrentVersion gets the current set version and returns an empty string if a version is not set. An installDir is provided which is where all the versions of the plugins are installed.
	// For example go 1.23.3 would be installed in installDir/1.23.3 A pathDir is also provided so you can evaluate the symlinks of the binaries in there to find the current version
	GetCurrentVersion(installDir string, pathDir string) (string, error)
}
