module github.com/Kuroki-g/go-gml/docs/go/gml-parser

go 1.26.0

require (
	github.com/Kuroki-g/go-gml/gml v0.0.0
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/cobra v1.10.2 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
	golang.org/x/text v0.35.0 // indirect
)

require (
	github.com/Kuroki-g/go-gml/core v0.0.0 // indirect
	github.com/Kuroki-g/go-gml/gml3_2_1 v0.0.0 // indirect
)

replace (
	github.com/Kuroki-g/go-gml/gml => ../../../gml
	github.com/Kuroki-g/go-gml/core => ../../../core
	github.com/Kuroki-g/go-gml/gml3_2_1 => ../../../gml3_2_1
)
