//go:build unit || integration || smoke

package test

import "embed"

// By default, load all fixture files. You can adjust the below go:embed line if you have massive files that
// should be loaded from the file system or other requirements that this approach doesn't support.

//go:embed fixtures
var Fixtures embed.FS
