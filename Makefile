.PHONY: help
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "go-gml (main library)"
	@echo "  build            go build ./..."
	@echo "  test             go test ./..."
	@echo "  cover            カバレッジレポート生成"
	@echo ""
	@echo "xsd2go-lite (code generator)"
	@echo "  xsd2go-build     バイナリビルド"
	@echo "  xsd2go-test      ユニットテスト実行"
	@echo "  xsd2go-cover     カバレッジレポート生成"
	@echo "  xsd2go-gen       GML XSD → pkg/gml/v3/geometry.go 生成"
	@echo ""
	@echo "gml-parser (CLI example)"
	@echo "  gml-parser-build バイナリビルド"
	@echo "  gml-parser-run   inspect サブコマンド実行 (testdata/N03)"

XSD2GO_DIR := docs/go/xsd2go-lite
XSD2GO_BIN := $(XSD2GO_DIR)/xsd2go-lite
XSD2GO_TMP := $(XSD2GO_DIR)/.tmp
GML_NS     := http://www.opengis.net/gml/3.2
GML_TMP    := .tmp

# ---- go-gml (main library) ----

.PHONY: build test cover

build:
	go build ./...

$(GML_TMP):
	mkdir -p $(GML_TMP)

test: $(GML_TMP)
	GOTMPDIR=$(GML_TMP) go test -count=1 ./...

cover: $(GML_TMP)
	GOTMPDIR=$(GML_TMP) go test -count=1 -coverprofile=$(GML_TMP)/cover.out ./...
	go tool cover -func=$(GML_TMP)/cover.out

# ---- xsd2go-lite (code generator) ----

.PHONY: xsd2go-build xsd2go-test xsd2go-cover xsd2go-gen

xsd2go-build:
	cd $(XSD2GO_DIR) && GOWORK=off go build -o xsd2go-lite .

xsd2go-test: $(XSD2GO_TMP)
	cd $(XSD2GO_DIR) && GOWORK=off GOTMPDIR=.tmp go test -count=1 ./...

xsd2go-cover: $(XSD2GO_TMP)
	cd $(XSD2GO_DIR) && GOWORK=off GOTMPDIR=.tmp go test -count=1 -coverprofile=.tmp/cover.out ./...
	cd $(XSD2GO_DIR) && go tool cover -func=.tmp/cover.out

GML2_NS := http://www.opengis.net/gml

xsd2go-gen: xsd2go-build
	$(XSD2GO_BIN) \
		-n "$(GML_NS)" \
		-p gml \
		--with-doc \
		-o pkg/gml/v3/geometry.go \
		$(XSD2GO_DIR)/schemas/gml/3.2.1/geometryAggregates.xsd

xsd2go-gen-v2: xsd2go-build
	$(XSD2GO_BIN) \
		-n "$(GML2_NS)" \
		-p gml \
		--with-doc \
		-o pkg/gml/v2/geometry.go \
		$(XSD2GO_DIR)/schemas/gml/2.1.2/geometry.xsd

$(XSD2GO_TMP):
	mkdir -p $(XSD2GO_TMP)

# ---- gml-parser (CLI example) ----

GML_PARSER_DIR := docs/go/gml-parser
GML_PARSER_BIN := $(GML_PARSER_DIR)/gml-parser

.PHONY: gml-parser-build gml-parser-run

gml-parser-build:
	cd $(GML_PARSER_DIR) && go build -o gml-parser .

gml-parser-run: gml-parser-build
	$(GML_PARSER_BIN) inspect -i testdata/N03/N03-20250101_13.xml --swap
