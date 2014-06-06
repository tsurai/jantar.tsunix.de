package main

import (
	"github.com/tsurai/jantar"
	"github.com/tsurai/jantar-middleware/public"
	"github.com/tsurai/jantar-middleware/static"
)

func main() {
	j := jantar.New(&jantar.Config{
		Hostname: "localhost",
		Port:     3000,
	})

	jantar.Log.SetMinLevel(jantar.LogLevelDebug)

	j.AddMiddleware(&public.Public{})
	j.AddMiddleware(&static.Static{})

	j.Run()
}
