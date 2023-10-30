package web

import (
	"embed"
)

//go:embed swagger-ui/*
//go:embed index.html
//go:embed oauth2-redirect.html
var WebUI embed.FS
