package translations

import "golang.org/x/text/language"

//go:generate gotext -srclang=en update -out=catalog.go -lang=en,es,de,en-US,es-419 ../server/server.go

var Languages = []language.Tag{
	language.English,
	language.Spanish,
	language.German,
	language.AmericanEnglish,
	language.LatinAmericanSpanish,
}

