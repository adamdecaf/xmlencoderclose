package embedpkg

import "embed"

// Importing embed used to crash xmlencoderclose on Go 1.27:
// internal error: package "embed" without types was imported from ...
var FS embed.FS
