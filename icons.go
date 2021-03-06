package desktop

import "fyne.io/fyne"

// AppData is an interface for accessing information about application icons
type AppData interface {
	Name() string       // Name is the name of the app usually
	Run([]string) error // Run is the command to run the app, passing any environment variables to be set

	Icon(theme string, size int) fyne.Resource // Icon returns an icon for the app in the requested theme and size
}

// ApplicationProvider describes a type that can locate icons and applications for the current system
type ApplicationProvider interface {
	AvailableApps() []AppData
	AvailableThemes() []string
	FindAppFromName(appName string) AppData
	FindAppFromWinInfo(win Window) AppData
	FindAppsMatching(pattern string) []AppData
	DefaultApps() []AppData
}
