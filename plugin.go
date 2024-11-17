package plugin

type Plugin struct {
	Download func(version string) error
	// Executes when Download is finished, installDir is the directory to install in, it should only contain the binaries of the tool
	Install func(installDir string) error
	// Channel that sends the download progress, number between 0 and 1, 0 is 0% and 1 is 100%
	Progress chan float64
	// Returns the version number of the latest version available and an error if there is any
	GetLatestVersion func() (string, error)
	// The Data variable stores any information you might need to save, it is only accessed by you. A struct would be a good option for storing multiple variables
	Data any
}

// New initializes the plugin
type New func() (Plugin, error)
