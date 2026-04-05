module github.com/Kuroki-g/go-gml/cmd/citygml-parser

go 1.26.0

require (
	github.com/Kuroki-g/go-gml/citygml2_0 v0.0.0
	github.com/Kuroki-g/go-gml/gml3_1_1 v0.2.1
	github.com/spf13/cobra v1.10.2
)

require (
	github.com/Kuroki-g/go-gml/core v0.2.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
)

replace (
	github.com/Kuroki-g/go-gml/citygml2_0 => ../../citygml2_0
	github.com/Kuroki-g/go-gml/core => ../../core
	github.com/Kuroki-g/go-gml/gml3_1_1 => ../../gml3_1_1
)
