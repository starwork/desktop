package desktop

// DeskSettings describes the configuration options available for Fyne desktop
type DeskSettings interface {
	Background() string
	IconTheme() string
	LauncherIcons() []string
	LauncherIconSize() int
	LauncherDisableTaskbar() bool
	LauncherDisableZoom() bool
	LauncherZoomScale() float64
	AddChangeListener(listener chan DeskSettings)
}
