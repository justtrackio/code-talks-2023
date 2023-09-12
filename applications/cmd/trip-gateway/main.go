package main

import (
	"codetalks/internal/gateway"
	"github.com/justtrackio/gosoline/pkg/application"
)

func main() {
	application.RunApiServer(gateway.Definer)
}
