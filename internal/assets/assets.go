package assets

import "embed"

//go:embed css/* js/* images/*
var Files embed.FS
