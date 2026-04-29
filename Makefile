.PHONY: help
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "go-gml (main library)"
	@echo "  build            go build ./..."
	@echo "  test             go test ./..."
	@echo "  cover            カバレッジレポート生成"
	@echo ""
	@echo "xsd2go-lite2 (code generator)"
	@echo "  xsd2go2-build    バイナリビルド"
	@echo "  xsd2go2-test     ユニットテスト実行"
	@echo ""
	@echo "fuzz"
	@echo "  fuzz-gen         fuzz seed corpus 生成 (gml*/testdata/fuzz/FuzzReader/)"
	@echo ""
	@echo "check-coverage"
	@echo "  check-coverage   未処理 PropertyType フィールドを検出"
	@echo ""
	@echo "epsg-gen (EPSG dimension table generator)"
	@echo "  epsg-gen-build   バイナリビルド"
	@echo "  epsg-gen CODES=\"4326 6668\"  EPSG コードの次元を API から取得し core/epsg_dim_table.go を更新"
	@echo ""
	@echo "gml-parser (CLI example)"
	@echo "  gml-parser-build バイナリビルド"
	@echo "  gml-parser-run   inspect サブコマンド実行 (testdata/N03)"

XSD2GO2_DIR := cmd/xsd2go-lite2
XSD2GO2_BIN := $(XSD2GO2_DIR)/xsd2go-lite2
SCHEMA_DIR := docs/schemas

# ---- xsd2go-gen version configuration ----
GML_VERSION ?= 3.2.1

ifeq ($(GML_VERSION),3.2.1)
  _GEN_NS  := http://www.opengis.net/gml/3.2
  _GEN_OUT := gml3_2_1/generated/geometry.go
  _GEN_XSD := $(SCHEMA_DIR)/gml/3.2.1/geometryAggregates.xsd
else ifeq ($(GML_VERSION),3.1.1)
  _GEN_NS  := http://www.opengis.net/gml
  _GEN_OUT := gml3_1_1/generated/geometry.go
  _GEN_XSD := $(SCHEMA_DIR)/gml/3.1.1/base/geometryAggregates.xsd
else ifeq ($(GML_VERSION),2.1.2)
  _GEN_NS  := http://www.opengis.net/gml
  _GEN_OUT := gml2_1_2/generated/geometry.go
  _GEN_XSD := $(SCHEMA_DIR)/gml/2.1.2/geometry.xsd
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
	@echo "=== xsd2go-lite2 ==="
	GONOSUMDB='*' GOWORK=off go test -C $(XSD2GO2_DIR) -count=1 ./...

cover:
	@for m in $(MODULES); do \
		echo "=== $$m ==="; \
		go -C $$m test -count=1 -coverprofile=$(GOTMPDIR)/cover_$$m.out ./...; \
		go tool cover -func=$(GOTMPDIR)/cover_$$m.out; \
	done

# ---- check-coverage ----

CHECK_COVERAGE_DIR := cmd/check-coverage

.PHONY: check-coverage

check-coverage:
	go run ./$(CHECK_COVERAGE_DIR)/

# ---- xsd2go-lite2 (code generator) ----

.PHONY: xsd2go2-build xsd2go2-test xsd2go-gen

xsd2go2-build:
	GONOSUMDB='*' GOWORK=off go build -C $(XSD2GO2_DIR) -o xsd2go-lite2 .

xsd2go2-test:
	GONOSUMDB='*' GOWORK=off go test -C $(XSD2GO2_DIR) -count=1 ./...

XLINK_NS  := http://www.w3.org/1999/xlink
XLINK_XSD := $(SCHEMA_DIR)/xlink/xlink.xsd

xsd2go-gen: xsd2go2-build
	@test -n "$(_GEN_OUT)" || (echo "Unknown GML_VERSION=$(GML_VERSION). Valid values: 3.2.1, 3.1.1, 2.1.2" >&2 && exit 1)
	$(XSD2GO2_BIN) \
		-n "$(_GEN_NS)" \
		-p gml \
		--with-doc \
		--catalog "$(XLINK_NS)=$(XLINK_XSD)" \
		-o $(_GEN_OUT) \
		$(_GEN_XSD)

XAL_NS   := urn:oasis:names:tc:ciq:xsdschema:xAL:2.0
XAL_XSD  := $(SCHEMA_DIR)/citygml/xAL/xAL.xsd
GML311_NS  := http://www.opengis.net/gml
GML311_XSD := $(SCHEMA_DIR)/gml/3.1.1/base/gml.xsd
CITYGML20_NS  := http://www.opengis.net/citygml/2.0
CITYGML20_XSD := $(SCHEMA_DIR)/citygml/2.0/cityGMLBase.xsd
CITYGML20_BLDG_XSD := $(SCHEMA_DIR)/citygml/building/2.0/building.xsd

XAL2_NS  := urn:oasis:names:tc:ciq:xsdschema:xAL:2.0

.PHONY: citygml2_0-gen
citygml2_0-gen: xsd2go2-build
	$(XSD2GO2_BIN) \
		-n "$(XAL2_NS)" \
		-p generated \
		--with-doc \
		-o citygml2_0/generated/xal.go \
		$(XAL_XSD)
	$(XSD2GO2_BIN) \
		-n "$(CITYGML20_NS)" \
		-p generated \
		--with-doc \
		--catalog "$(XLINK_NS)=$(XLINK_XSD)" \
		--catalog "$(XAL_NS)=$(XAL_XSD)" \
		--catalog "$(GML311_NS)=$(GML311_XSD)" \
		--map-namespace "$(GML311_NS)=github.com/Kuroki-g/go-gml/gml3_1_1/generated" \
		-o citygml2_0/generated/core.go \
		$(CITYGML20_XSD)
	$(XSD2GO2_BIN) \
		-n "http://www.opengis.net/citygml/building/2.0" \
		-p generated \
		--with-doc \
		--catalog "$(XLINK_NS)=$(XLINK_XSD)" \
		--catalog "$(XAL_NS)=$(XAL_XSD)" \
		--catalog "$(GML311_NS)=$(GML311_XSD)" \
		--catalog "$(CITYGML20_NS)=$(CITYGML20_XSD)" \
		--map-namespace "$(GML311_NS)=github.com/Kuroki-g/go-gml/gml3_1_1/generated" \
		-o citygml2_0/generated/building.go \
		$(CITYGML20_BLDG_XSD)

# ---- fuzz seed corpus ----

.PHONY: fuzz-gen

fuzz-gen:
	uv run python3 scripts/gen_fuzz_seeds.py

# ---- gml-parser (CLI example) ----

GML_PARSER_DIR := cmd/gml-parser
GML_PARSER_BIN := $(GML_PARSER_DIR)/gml-parser

.PHONY: gml-parser-build gml-parser-run

gml-parser-build:
	cd $(GML_PARSER_DIR) && go build -o gml-parser .

gml-parser-run: gml-parser-build
	$(GML_PARSER_BIN) inspect -i testdata/N03/N03-20250101_13.xml --swap

# ---- citygml-parser (CLI) ----

CITYGML_PARSER_DIR := cmd/citygml-parser
CITYGML_PARSER_BIN := $(CITYGML_PARSER_DIR)/.tmp/citygml-parser

.PHONY: citygml-parser-build citygml-parser-run

citygml-parser-build:
	go -C $(CITYGML_PARSER_DIR) build -o .tmp/citygml-parser .

citygml-parser-run: citygml-parser-build
	$(CITYGML_PARSER_BIN) inspect -i testdata/13105_bunkyo-ku_pref_2023_citygml_2_op/udx/bldg/53394548_bldg_6697_op.gml

# ---- epsg-gen (EPSG dimension table generator) ----

EPSG_GEN_DIR := cmd/epsg-gen
EPSG_GEN_BIN := $(EPSG_GEN_DIR)/.tmp/epsg-gen

.PHONY: epsg-gen-build epsg-gen

epsg-gen-build:
	go -C $(EPSG_GEN_DIR) build -o .tmp/epsg-gen .

epsg-gen: epsg-gen-build
	$(EPSG_GEN_BIN) $(CODES)
