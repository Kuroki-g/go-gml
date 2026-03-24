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
	@echo "  xsd2go-gen [GML_VERSION=3.2.1|3.1.1|2.1.2]"
	@echo "               GML XSD → gml<version>/generated/geometry.go 生成"
	@echo "               デフォルト: GML_VERSION=3.2.1"
	@echo ""
	@echo "gml-parser (CLI example)"
	@echo "  gml-parser-build バイナリビルド"
	@echo "  gml-parser-run   inspect サブコマンド実行 (testdata/N03)"

XSD2GO_DIR := docs/go/xsd2go-lite
XSD2GO_BIN := $(XSD2GO_DIR)/xsd2go-lite
XSD2GO_TMP := $(XSD2GO_DIR)/.tmp
GML_TMP    := .tmp

# ---- xsd2go-gen version configuration ----
GML_VERSION ?= 3.2.1

ifeq ($(GML_VERSION),3.2.1)
  _GEN_NS  := http://www.opengis.net/gml/3.2
  _GEN_OUT := gml3_2_1/generated/geometry.go
  _GEN_XSD := $(XSD2GO_DIR)/schemas/gml/3.2.1/geometryAggregates.xsd
else ifeq ($(GML_VERSION),3.1.1)
  _GEN_NS  := http://www.opengis.net/gml
  _GEN_OUT := gml3_1_1/generated/geometry.go
  _GEN_XSD := $(XSD2GO_DIR)/schemas/gml/3.1.1/base/geometryAggregates.xsd
else ifeq ($(GML_VERSION),2.1.2)
  _GEN_NS  := http://www.opengis.net/gml
  _GEN_OUT := gml2_1_2/generated/geometry.go
  _GEN_XSD := $(XSD2GO_DIR)/schemas/gml/2.1.2/geometry.xsd
else
  _GEN_NS  :=
  _GEN_OUT :=
  _GEN_XSD :=
endif

# ---- go-gml (main library) ----

.PHONY: build test cover

build:
	go build ./...

$(GML_TMP):
	mkdir -p $(GML_TMP)

MODULES := core gml2_1_2 gml3_1_1 gml3_2_1 gml

test: $(GML_TMP)
	@for m in $(MODULES); do \
		echo "=== $$m ==="; \
		GOTMPDIR=$(abspath $(GML_TMP)) go -C $$m test -count=1 ./...; \
	done

cover: $(GML_TMP)
	@for m in $(MODULES); do \
		echo "=== $$m ==="; \
		GOTMPDIR=$(abspath $(GML_TMP)) go -C $$m test -count=1 -coverprofile=$(abspath $(GML_TMP))/cover_$$m.out ./...; \
		go tool cover -func=$(abspath $(GML_TMP))/cover_$$m.out; \
	done

# ---- xsd2go-lite (code generator) ----

.PHONY: xsd2go-build xsd2go-test xsd2go-cover xsd2go-gen

xsd2go-build:
	cd $(XSD2GO_DIR) && GOWORK=off go build -o xsd2go-lite .

xsd2go-test: $(XSD2GO_TMP)
	cd $(XSD2GO_DIR) && GOWORK=off GOTMPDIR=$(abspath $(XSD2GO_TMP)) go test -count=1 ./...

xsd2go-cover: $(XSD2GO_TMP)
	cd $(XSD2GO_DIR) && GOWORK=off GOTMPDIR=$(abspath $(XSD2GO_TMP)) go test -count=1 -coverprofile=.tmp/cover.out ./...
	cd $(XSD2GO_DIR) && go tool cover -func=.tmp/cover.out

XLINK_NS  := http://www.w3.org/1999/xlink
XLINK_XSD := $(XSD2GO_DIR)/schemas/xlink/xlink.xsd

xsd2go-gen: xsd2go-build
	@test -n "$(_GEN_OUT)" || (echo "Unknown GML_VERSION=$(GML_VERSION). Valid values: 3.2.1, 3.1.1, 2.1.2" >&2 && exit 1)
	$(XSD2GO_BIN) \
		-n "$(_GEN_NS)" \
		-p gml \
		--with-doc \
		--catalog "$(XLINK_NS)=$(XLINK_XSD)" \
		-o $(_GEN_OUT) \
		$(_GEN_XSD)

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
