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

MODULES := core gml2_1_2 gml3_1_1 gml3_2_1 gml citygml2_0

test:
	@for m in $(MODULES); do \
		echo "=== $$m ==="; \
		go -C $$m test -count=1 ./...; \
	done

cover:
	@for m in $(MODULES); do \
		echo "=== $$m ==="; \
		go -C $$m test -count=1 -coverprofile=$(GOTMPDIR)/cover_$$m.out ./...; \
		go tool cover -func=$(GOTMPDIR)/cover_$$m.out; \
	done

# ---- xsd2go-lite (code generator) ----

.PHONY: xsd2go-build xsd2go-test xsd2go-cover xsd2go-gen

xsd2go-build:
	cd $(XSD2GO_DIR) && GOWORK=off go build -o xsd2go-lite .

xsd2go-test:
	cd $(XSD2GO_DIR) && GOWORK=off go test -count=1 ./...

xsd2go-cover:
	cd $(XSD2GO_DIR) && GOWORK=off go test -count=1 -coverprofile=$(GOTMPDIR)/cover_xsd2go.out ./...
	go tool cover -func=$(GOTMPDIR)/cover_xsd2go.out

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

XAL_NS   := urn:oasis:names:tc:ciq:xsdschema:xAL:2.0
XAL_XSD  := $(XSD2GO_DIR)/schemas/citygml/xAL/xAL.xsd
GML311_NS  := http://www.opengis.net/gml
GML311_XSD := $(XSD2GO_DIR)/schemas/gml/3.1.1/base/gml.xsd
CITYGML20_NS  := http://www.opengis.net/citygml/2.0
CITYGML20_XSD := $(XSD2GO_DIR)/schemas/citygml/2.0/cityGMLBase.xsd
CITYGML20_BLDG_XSD := $(XSD2GO_DIR)/schemas/citygml/building/2.0/building.xsd

XAL2_NS  := urn:oasis:names:tc:ciq:xsdschema:xAL:2.0

.PHONY: citygml2_0-gen
citygml2_0-gen: xsd2go-build
	$(XSD2GO_BIN) \
		-n "$(XAL2_NS)" \
		-p generated \
		--with-doc \
		-o citygml2_0/generated/xal.go \
		$(XAL_XSD)
	$(XSD2GO_BIN) \
		-n "$(CITYGML20_NS)" \
		-p generated \
		--with-doc \
		--catalog "$(XLINK_NS)=$(XLINK_XSD)" \
		--catalog "$(XAL_NS)=$(XAL_XSD)" \
		--catalog "$(GML311_NS)=$(GML311_XSD)" \
		--omit-namespace "$(GML311_NS)" \
		-o citygml2_0/generated/core.go \
		$(CITYGML20_XSD)
	$(XSD2GO_BIN) \
		-n "http://www.opengis.net/citygml/building/2.0" \
		-p generated \
		--with-doc \
		--catalog "$(XLINK_NS)=$(XLINK_XSD)" \
		--catalog "$(XAL_NS)=$(XAL_XSD)" \
		--catalog "$(GML311_NS)=$(GML311_XSD)" \
		--catalog "$(CITYGML20_NS)=$(CITYGML20_XSD)" \
		--omit-namespace "$(GML311_NS)" \
		-o citygml2_0/generated/building.go \
		$(CITYGML20_BLDG_XSD)

# ---- gml-parser (CLI example) ----

GML_PARSER_DIR := cmd/gml-parser
GML_PARSER_BIN := $(GML_PARSER_DIR)/gml-parser

.PHONY: gml-parser-build gml-parser-run

gml-parser-build:
	cd $(GML_PARSER_DIR) && go build -o gml-parser .

gml-parser-run: gml-parser-build
	$(GML_PARSER_BIN) inspect -i testdata/N03/N03-20250101_13.xml --swap
