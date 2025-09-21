package assets

import "embed"

//go:embed data/*.txt
var Files embed.FS
