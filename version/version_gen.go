package version

import (
	"fmt"
	"runtime"
)

var (
	Name        = "ketch-forwarder"
	Description = "ketch-forwarder"
	ReleaseName = ""
	Version     = "0.0.0"
	Prerelease  = ""
)

// String returns the complete version string, including prerelease
func String() string {
	s := fmt.Sprintf("%s %s %s", runtime.GOOS, runtime.GOARCH, runtime.Version())
	if Prerelease != "" {
		return fmt.Sprintf("%s-%s %s", Version, Prerelease, s)
	}
	return fmt.Sprintf("%s %s", Version, s)
}
