package static

import "embed"

//go:embed *
var staticContent embed.FS

func GetStaticContent() embed.FS {
	return staticContent
}
