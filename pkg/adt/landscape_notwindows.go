//go:build !windows

package adt

// findLandscapeFilesFromRegistry is a stub on non-Windows platforms.
// On Linux and macOS, SAP landscape files are found via environment variables
// and well-known default paths, not the Windows registry.
func findLandscapeFilesFromRegistry() []string {
	return nil
}
