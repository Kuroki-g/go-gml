package gml

type AbstractCRSType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	Remarks              *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     []string               `xml:"http://www.opengis.net/gml/3.2 domainOfValidity"`
	Scope                []string               `xml:"http://www.opengis.net/gml/3.2 scope"`
}

type AbstractContinuousCoverageType struct {
	MetaDataProperty     []MetaDataPropertyType        `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType              `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType        `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                    `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                        `xml:"id,attr,omitempty"`
	BoundedBy            *BoundingShapeType            `xml:"http://www.opengis.net/gml/3.2 boundedBy,omitempty"`
	Location             *LocationPropertyType         `xml:"http://www.opengis.net/gml/3.2 location,omitempty"`
	PriorityLocation     *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml/3.2 priorityLocation,omitempty"`
	DomainSet            *DomainSetType                `xml:"http://www.opengis.net/gml/3.2 domainSet,omitempty"`
	GridDomain           *DomainSetType                `xml:"http://www.opengis.net/gml/3.2 gridDomain,omitempty"`
	MultiCurveDomain     *DomainSetType                `xml:"http://www.opengis.net/gml/3.2 multiCurveDomain,omitempty"`
	MultiPointDomain     *DomainSetType                `xml:"http://www.opengis.net/gml/3.2 multiPointDomain,omitempty"`
	MultiSolidDomain     *DomainSetType                `xml:"http://www.opengis.net/gml/3.2 multiSolidDomain,omitempty"`
	MultiSurfaceDomain   *DomainSetType                `xml:"http://www.opengis.net/gml/3.2 multiSurfaceDomain,omitempty"`
	RectifiedGridDomain  *DomainSetType                `xml:"http://www.opengis.net/gml/3.2 rectifiedGridDomain,omitempty"`
	RangeSet             *RangeSetType                 `xml:"http://www.opengis.net/gml/3.2 rangeSet,omitempty"`
	CoverageFunction     *CoverageFunctionType         `xml:"http://www.opengis.net/gml/3.2 coverageFunction,omitempty"`
}

type AbstractCoordinateOperationType struct {
	MetaDataProperty            []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description                 *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference        *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier                  *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                        []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                          string                 `xml:"id,attr,omitempty"`
	Remarks                     *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity            *string                `xml:"http://www.opengis.net/gml/3.2 domainOfValidity,omitempty"`
	Scope                       []string               `xml:"http://www.opengis.net/gml/3.2 scope"`
	OperationVersion            *string                `xml:"http://www.opengis.net/gml/3.2 operationVersion,omitempty"`
	CoordinateOperationAccuracy []string               `xml:"http://www.opengis.net/gml/3.2 coordinateOperationAccuracy"`
	SourceCRS                   *CRSPropertyType       `xml:"http://www.opengis.net/gml/3.2 sourceCRS,omitempty"`
	TargetCRS                   *CRSPropertyType       `xml:"http://www.opengis.net/gml/3.2 targetCRS,omitempty"`
}

type AbstractCoordinateSystemType struct {
	MetaDataProperty     []MetaDataPropertyType             `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType                   `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                     `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType             `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                         `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                             `xml:"id,attr,omitempty"`
	Remarks              *string                            `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	Axis                 []CoordinateSystemAxisPropertyType `xml:"http://www.opengis.net/gml/3.2 axis"`
	UsesAxis             *CoordinateSystemAxisPropertyType  `xml:"http://www.opengis.net/gml/3.2 usesAxis,omitempty"`
	AggregationType      string                             `xml:"aggregationType,attr,omitempty"`
}

type AbstractCoverageType struct {
	MetaDataProperty     []MetaDataPropertyType        `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType              `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType        `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                    `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                        `xml:"id,attr,omitempty"`
	BoundedBy            *BoundingShapeType            `xml:"http://www.opengis.net/gml/3.2 boundedBy,omitempty"`
	Location             *LocationPropertyType         `xml:"http://www.opengis.net/gml/3.2 location,omitempty"`
	PriorityLocation     *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml/3.2 priorityLocation,omitempty"`
	DomainSet            *DomainSetType                `xml:"http://www.opengis.net/gml/3.2 domainSet,omitempty"`
	GridDomain           *DomainSetType                `xml:"http://www.opengis.net/gml/3.2 gridDomain,omitempty"`
	MultiCurveDomain     *DomainSetType                `xml:"http://www.opengis.net/gml/3.2 multiCurveDomain,omitempty"`
	MultiPointDomain     *DomainSetType                `xml:"http://www.opengis.net/gml/3.2 multiPointDomain,omitempty"`
	MultiSolidDomain     *DomainSetType                `xml:"http://www.opengis.net/gml/3.2 multiSolidDomain,omitempty"`
	MultiSurfaceDomain   *DomainSetType                `xml:"http://www.opengis.net/gml/3.2 multiSurfaceDomain,omitempty"`
	RectifiedGridDomain  *DomainSetType                `xml:"http://www.opengis.net/gml/3.2 rectifiedGridDomain,omitempty"`
	RangeSet             *RangeSetType                 `xml:"http://www.opengis.net/gml/3.2 rangeSet,omitempty"`
}

type AbstractCurveSegmentType struct {
	NumDerivativesAtStart int `xml:"numDerivativesAtStart,attr,omitempty"`
	NumDerivativesAtEnd   int `xml:"numDerivativesAtEnd,attr,omitempty"`
	NumDerivativeInterior int `xml:"numDerivativeInterior,attr,omitempty"`
}

type AbstractCurveType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	SrsName              string                 `xml:"srsName,attr,omitempty"`
	SrsDimension         int                    `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                 `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                 `xml:"uomLabels,attr,omitempty"`
}

type AbstractDatumType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	Remarks              *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     *string                `xml:"http://www.opengis.net/gml/3.2 domainOfValidity,omitempty"`
	Scope                []string               `xml:"http://www.opengis.net/gml/3.2 scope"`
	AnchorDefinition     *CodeType              `xml:"http://www.opengis.net/gml/3.2 anchorDefinition,omitempty"`
	AnchorPoint          *CodeType              `xml:"http://www.opengis.net/gml/3.2 anchorPoint,omitempty"`
	RealizationEpoch     *string                `xml:"http://www.opengis.net/gml/3.2 realizationEpoch,omitempty"`
}

type AbstractFeatureCollectionType struct {
	MetaDataProperty     []MetaDataPropertyType        `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType              `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType        `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                    `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                        `xml:"id,attr,omitempty"`
	BoundedBy            *BoundingShapeType            `xml:"http://www.opengis.net/gml/3.2 boundedBy,omitempty"`
	Location             *LocationPropertyType         `xml:"http://www.opengis.net/gml/3.2 location,omitempty"`
	PriorityLocation     *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml/3.2 priorityLocation,omitempty"`
	FeatureMember        []FeaturePropertyType         `xml:"http://www.opengis.net/gml/3.2 featureMember"`
	FeatureMembers       *FeatureArrayPropertyType     `xml:"http://www.opengis.net/gml/3.2 featureMembers,omitempty"`
}

type AbstractFeatureMemberType struct {
	Owns bool `xml:"owns,attr,omitempty"`
}

type AbstractFeatureType struct {
	MetaDataProperty     []MetaDataPropertyType        `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType              `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType        `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                    `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                        `xml:"id,attr,omitempty"`
	BoundedBy            *BoundingShapeType            `xml:"http://www.opengis.net/gml/3.2 boundedBy,omitempty"`
	Location             *LocationPropertyType         `xml:"http://www.opengis.net/gml/3.2 location,omitempty"`
	PriorityLocation     *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml/3.2 priorityLocation,omitempty"`
}

type AbstractGMLType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
}

type AbstractGeneralConversionType struct {
	MetaDataProperty            []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description                 *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference        *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier                  *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                        []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                          string                 `xml:"id,attr,omitempty"`
	Remarks                     *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity            *string                `xml:"http://www.opengis.net/gml/3.2 domainOfValidity,omitempty"`
	Scope                       []string               `xml:"http://www.opengis.net/gml/3.2 scope"`
	OperationVersion            *string                `xml:"http://www.opengis.net/gml/3.2 operationVersion,omitempty"`
	CoordinateOperationAccuracy []string               `xml:"http://www.opengis.net/gml/3.2 coordinateOperationAccuracy"`
	SourceCRS                   *CRSPropertyType       `xml:"http://www.opengis.net/gml/3.2 sourceCRS,omitempty"`
	TargetCRS                   *CRSPropertyType       `xml:"http://www.opengis.net/gml/3.2 targetCRS,omitempty"`
}

type AbstractGeneralDerivedCRSType struct {
	MetaDataProperty     []MetaDataPropertyType         `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType               `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                 `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType         `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                     `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                         `xml:"id,attr,omitempty"`
	Remarks              *string                        `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     []string                       `xml:"http://www.opengis.net/gml/3.2 domainOfValidity"`
	Scope                []string                       `xml:"http://www.opengis.net/gml/3.2 scope"`
	Conversion           *GeneralConversionPropertyType `xml:"http://www.opengis.net/gml/3.2 conversion,omitempty"`
	DefinedByConversion  *GeneralConversionPropertyType `xml:"http://www.opengis.net/gml/3.2 definedByConversion,omitempty"`
}

type AbstractGeneralOperationParameterPropertyType struct {
	AbstractGeneralOperationParameter *AbstractGeneralOperationParameterType `xml:"http://www.opengis.net/gml/3.2 AbstractGeneralOperationParameter,omitempty"`
	OperationParameter                *OperationParameterType                `xml:"http://www.opengis.net/gml/3.2 OperationParameter,omitempty"`
	OperationParameterGroup           *OperationParameterGroupType           `xml:"http://www.opengis.net/gml/3.2 OperationParameterGroup,omitempty"`
	NilReason                         string                                 `xml:"nilReason,attr,omitempty"`
	RemoteSchema                      string                                 `xml:"remoteSchema,attr,omitempty"`
}

type AbstractGeneralOperationParameterType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	Remarks              *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	MinimumOccurs        *int                   `xml:"http://www.opengis.net/gml/3.2 minimumOccurs,omitempty"`
}

type AbstractGeneralParameterValuePropertyType struct {
	AbstractGeneralParameterValue *AbstractGeneralParameterValueType `xml:"http://www.opengis.net/gml/3.2 AbstractGeneralParameterValue,omitempty"`
	ParameterValue                *ParameterValueType                `xml:"http://www.opengis.net/gml/3.2 ParameterValue,omitempty"`
	ParameterValueGroup           *ParameterValueGroupType           `xml:"http://www.opengis.net/gml/3.2 ParameterValueGroup,omitempty"`
}

type AbstractGeneralParameterValueType struct {
}

type AbstractGeneralTransformationType struct {
	MetaDataProperty            []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description                 *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference        *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier                  *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                        []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                          string                 `xml:"id,attr,omitempty"`
	Remarks                     *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity            *string                `xml:"http://www.opengis.net/gml/3.2 domainOfValidity,omitempty"`
	Scope                       []string               `xml:"http://www.opengis.net/gml/3.2 scope"`
	OperationVersion            *string                `xml:"http://www.opengis.net/gml/3.2 operationVersion,omitempty"`
	CoordinateOperationAccuracy []string               `xml:"http://www.opengis.net/gml/3.2 coordinateOperationAccuracy"`
	SourceCRS                   *CRSPropertyType       `xml:"http://www.opengis.net/gml/3.2 sourceCRS,omitempty"`
	TargetCRS                   *CRSPropertyType       `xml:"http://www.opengis.net/gml/3.2 targetCRS,omitempty"`
}

type AbstractGeometricAggregateType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	SrsName              string                 `xml:"srsName,attr,omitempty"`
	SrsDimension         int                    `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                 `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                 `xml:"uomLabels,attr,omitempty"`
	AggregationType      string                 `xml:"aggregationType,attr,omitempty"`
}

type AbstractGeometricPrimitiveType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	SrsName              string                 `xml:"srsName,attr,omitempty"`
	SrsDimension         int                    `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                 `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                 `xml:"uomLabels,attr,omitempty"`
}

type AbstractGeometryType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	SrsName              string                 `xml:"srsName,attr,omitempty"`
	SrsDimension         int                    `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                 `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                 `xml:"uomLabels,attr,omitempty"`
}

type AbstractGriddedSurfaceType struct {
	AggregationType string `xml:"aggregationType,attr,omitempty"`
	Rows            string `xml:"http://www.opengis.net/gml/3.2 rows,omitempty"`
	Columns         int    `xml:"columns,attr,omitempty"`
}

type AbstractMemberType struct {
	Owns bool `xml:"owns,attr,omitempty"`
}

type AbstractMetaDataType struct {
	Id string `xml:"id,attr,omitempty"`
}

type AbstractMetadataPropertyType struct {
	Owns bool `xml:"owns,attr,omitempty"`
}

type AbstractParametricCurveSurfaceType struct {
	AggregationType string `xml:"aggregationType,attr,omitempty"`
}

type AbstractRingPropertyType struct {
	AbstractRing *AbstractRingType `xml:"http://www.opengis.net/gml/3.2 AbstractRing,omitempty"`
	LinearRing   *LinearRingType   `xml:"http://www.opengis.net/gml/3.2 LinearRing,omitempty"`
	Ring         *RingType         `xml:"http://www.opengis.net/gml/3.2 Ring,omitempty"`
}

type AbstractRingType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	SrsName              string                 `xml:"srsName,attr,omitempty"`
	SrsDimension         int                    `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                 `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                 `xml:"uomLabels,attr,omitempty"`
}

type AbstractSolidType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	SrsName              string                 `xml:"srsName,attr,omitempty"`
	SrsDimension         int                    `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                 `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                 `xml:"uomLabels,attr,omitempty"`
}

type AbstractSurfacePatchType struct {
}

type AbstractSurfaceType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	SrsName              string                 `xml:"srsName,attr,omitempty"`
	SrsDimension         int                    `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                 `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                 `xml:"uomLabels,attr,omitempty"`
}

type AbstractTimeComplexType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
}

type AbstractTimeGeometricPrimitiveType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	RelatedTime          []RelatedTimeType      `xml:"http://www.opengis.net/gml/3.2 relatedTime"`
	Frame                string                 `xml:"frame,attr,omitempty"`
}

type AbstractTimeObjectType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
}

type AbstractTimePrimitiveType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	RelatedTime          []RelatedTimeType      `xml:"http://www.opengis.net/gml/3.2 relatedTime"`
}

type AbstractTimeSliceType struct {
	MetaDataProperty     []MetaDataPropertyType     `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType           `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType             `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType     `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                 `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                     `xml:"id,attr,omitempty"`
	ValidTime            *TimePrimitivePropertyType `xml:"http://www.opengis.net/gml/3.2 validTime,omitempty"`
	DataSource           *StringOrRefType           `xml:"http://www.opengis.net/gml/3.2 dataSource,omitempty"`
}

type AbstractTimeTopologyPrimitiveType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	RelatedTime          []RelatedTimeType      `xml:"http://www.opengis.net/gml/3.2 relatedTime"`
	Complex              *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 complex,omitempty"`
}

type AbstractTopoPrimitiveType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
}

type AbstractTopologyType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
}

type AffineCSPropertyType struct {
	AffineCS     *AffineCSType `xml:"http://www.opengis.net/gml/3.2 AffineCS,omitempty"`
	NilReason    string        `xml:"nilReason,attr,omitempty"`
	RemoteSchema string        `xml:"remoteSchema,attr,omitempty"`
}

type AffineCSType struct {
	MetaDataProperty     []MetaDataPropertyType             `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType                   `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                     `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType             `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                         `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                             `xml:"id,attr,omitempty"`
	Remarks              *string                            `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	Axis                 []CoordinateSystemAxisPropertyType `xml:"http://www.opengis.net/gml/3.2 axis"`
	UsesAxis             *CoordinateSystemAxisPropertyType  `xml:"http://www.opengis.net/gml/3.2 usesAxis,omitempty"`
	AggregationType      string                             `xml:"aggregationType,attr,omitempty"`
}

type AffinePlacementType struct {
	Location     *DirectPositionType `xml:"http://www.opengis.net/gml/3.2 location,omitempty"`
	RefDirection []VectorType        `xml:"http://www.opengis.net/gml/3.2 refDirection"`
	InDimension  int                 `xml:"http://www.opengis.net/gml/3.2 inDimension,omitempty"`
	OutDimension int                 `xml:"http://www.opengis.net/gml/3.2 outDimension,omitempty"`
}

type AngleChoiceType struct {
	Angle    *AngleType    `xml:"http://www.opengis.net/gml/3.2 angle,omitempty"`
	DmsAngle *DMSAngleType `xml:"http://www.opengis.net/gml/3.2 dmsAngle,omitempty"`
}

type AngleType struct {
	Value string `xml:",chardata"`
}

type ArcByBulgeType struct {
	NumDerivativesAtStart int                     `xml:"numDerivativesAtStart,attr,omitempty"`
	NumDerivativesAtEnd   int                     `xml:"numDerivativesAtEnd,attr,omitempty"`
	NumDerivativeInterior int                     `xml:"numDerivativeInterior,attr,omitempty"`
	Bulge                 []float64               `xml:"http://www.opengis.net/gml/3.2 bulge"`
	Normal                []VectorType            `xml:"http://www.opengis.net/gml/3.2 normal"`
	PosList               *DirectPositionListType `xml:"http://www.opengis.net/gml/3.2 posList,omitempty"`
	Coordinates           *CoordinatesType        `xml:"http://www.opengis.net/gml/3.2 coordinates,omitempty"`
	Pos                   *DirectPositionType     `xml:"http://www.opengis.net/gml/3.2 pos,omitempty"`
	PointProperty         *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointProperty,omitempty"`
	PointRep              *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointRep,omitempty"`
	Interpolation         string                  `xml:"interpolation,attr,omitempty"`
	NumArc                int                     `xml:"numArc,attr,omitempty"`
}

type ArcByCenterPointType struct {
	NumDerivativesAtStart int                     `xml:"numDerivativesAtStart,attr,omitempty"`
	NumDerivativesAtEnd   int                     `xml:"numDerivativesAtEnd,attr,omitempty"`
	NumDerivativeInterior int                     `xml:"numDerivativeInterior,attr,omitempty"`
	Radius                *LengthType             `xml:"http://www.opengis.net/gml/3.2 radius,omitempty"`
	StartAngle            *AngleType              `xml:"http://www.opengis.net/gml/3.2 startAngle,omitempty"`
	EndAngle              *AngleType              `xml:"http://www.opengis.net/gml/3.2 endAngle,omitempty"`
	PosList               *DirectPositionListType `xml:"http://www.opengis.net/gml/3.2 posList,omitempty"`
	Coordinates           *CoordinatesType        `xml:"http://www.opengis.net/gml/3.2 coordinates,omitempty"`
	Pos                   *DirectPositionType     `xml:"http://www.opengis.net/gml/3.2 pos,omitempty"`
	PointProperty         *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointProperty,omitempty"`
	PointRep              *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointRep,omitempty"`
	Interpolation         string                  `xml:"interpolation,attr,omitempty"`
	NumArc                int                     `xml:"numArc,attr"`
}

type ArcStringByBulgeType struct {
	NumDerivativesAtStart int                     `xml:"numDerivativesAtStart,attr,omitempty"`
	NumDerivativesAtEnd   int                     `xml:"numDerivativesAtEnd,attr,omitempty"`
	NumDerivativeInterior int                     `xml:"numDerivativeInterior,attr,omitempty"`
	Bulge                 []float64               `xml:"http://www.opengis.net/gml/3.2 bulge"`
	Normal                []VectorType            `xml:"http://www.opengis.net/gml/3.2 normal"`
	PosList               *DirectPositionListType `xml:"http://www.opengis.net/gml/3.2 posList,omitempty"`
	Coordinates           *CoordinatesType        `xml:"http://www.opengis.net/gml/3.2 coordinates,omitempty"`
	Pos                   *DirectPositionType     `xml:"http://www.opengis.net/gml/3.2 pos,omitempty"`
	PointProperty         *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointProperty,omitempty"`
	PointRep              *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointRep,omitempty"`
	Interpolation         string                  `xml:"interpolation,attr,omitempty"`
	NumArc                int                     `xml:"numArc,attr,omitempty"`
}

type ArcStringType struct {
	NumDerivativesAtStart int                     `xml:"numDerivativesAtStart,attr,omitempty"`
	NumDerivativesAtEnd   int                     `xml:"numDerivativesAtEnd,attr,omitempty"`
	NumDerivativeInterior int                     `xml:"numDerivativeInterior,attr,omitempty"`
	PosList               *DirectPositionListType `xml:"http://www.opengis.net/gml/3.2 posList,omitempty"`
	Coordinates           *CoordinatesType        `xml:"http://www.opengis.net/gml/3.2 coordinates,omitempty"`
	Pos                   *DirectPositionType     `xml:"http://www.opengis.net/gml/3.2 pos,omitempty"`
	PointProperty         *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointProperty,omitempty"`
	PointRep              *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointRep,omitempty"`
	Interpolation         string                  `xml:"interpolation,attr,omitempty"`
	NumArc                int                     `xml:"numArc,attr,omitempty"`
}

type ArcType struct {
	NumDerivativesAtStart int                     `xml:"numDerivativesAtStart,attr,omitempty"`
	NumDerivativesAtEnd   int                     `xml:"numDerivativesAtEnd,attr,omitempty"`
	NumDerivativeInterior int                     `xml:"numDerivativeInterior,attr,omitempty"`
	PosList               *DirectPositionListType `xml:"http://www.opengis.net/gml/3.2 posList,omitempty"`
	Coordinates           *CoordinatesType        `xml:"http://www.opengis.net/gml/3.2 coordinates,omitempty"`
	Pos                   *DirectPositionType     `xml:"http://www.opengis.net/gml/3.2 pos,omitempty"`
	PointProperty         *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointProperty,omitempty"`
	PointRep              *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointRep,omitempty"`
	Interpolation         string                  `xml:"interpolation,attr,omitempty"`
	NumArc                int                     `xml:"numArc,attr,omitempty"`
}

type AreaType struct {
	Value string `xml:",chardata"`
}

type ArrayAssociationType struct {
	AbstractObject                []string                           `xml:"http://www.opengis.net/gml/3.2 AbstractObject"`
	AbstractCurveSegment          *AbstractCurveSegmentType          `xml:"http://www.opengis.net/gml/3.2 AbstractCurveSegment,omitempty"`
	AbstractGML                   *AbstractGMLType                   `xml:"http://www.opengis.net/gml/3.2 AbstractGML,omitempty"`
	AbstractGeneralParameterValue *AbstractGeneralParameterValueType `xml:"http://www.opengis.net/gml/3.2 AbstractGeneralParameterValue,omitempty"`
	AbstractMetaData              *AbstractMetaDataType              `xml:"http://www.opengis.net/gml/3.2 AbstractMetaData,omitempty"`
	AbstractValue                 *string                            `xml:"http://www.opengis.net/gml/3.2 AbstractValue,omitempty"`
	AffinePlacement               *AffinePlacementType               `xml:"http://www.opengis.net/gml/3.2 AffinePlacement,omitempty"`
	CoverageMappingRule           *MappingRuleType                   `xml:"http://www.opengis.net/gml/3.2 CoverageMappingRule,omitempty"`
	DataBlock                     *DataBlockType                     `xml:"http://www.opengis.net/gml/3.2 DataBlock,omitempty"`
	Envelope                      *EnvelopeType                      `xml:"http://www.opengis.net/gml/3.2 Envelope,omitempty"`
	File                          *FileType                          `xml:"http://www.opengis.net/gml/3.2 File,omitempty"`
	GridFunction                  *GridFunctionType                  `xml:"http://www.opengis.net/gml/3.2 GridFunction,omitempty"`
	CoverageFunction              *CoverageFunctionType              `xml:"http://www.opengis.net/gml/3.2 coverageFunction,omitempty"`
	Owns                          bool                               `xml:"owns,attr,omitempty"`
}

type ArrayType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	Members              *ArrayAssociationType  `xml:"http://www.opengis.net/gml/3.2 members,omitempty"`
}

type AssociationRoleType struct {
	Owns         bool   `xml:"owns,attr,omitempty"`
	NilReason    string `xml:"nilReason,attr,omitempty"`
	RemoteSchema string `xml:"remoteSchema,attr,omitempty"`
}

type BSplineType struct {
	NumDerivativesAtStart int                     `xml:"numDerivativesAtStart,attr,omitempty"`
	NumDerivativesAtEnd   int                     `xml:"numDerivativesAtEnd,attr,omitempty"`
	NumDerivativeInterior int                     `xml:"numDerivativeInterior,attr,omitempty"`
	Degree                int                     `xml:"http://www.opengis.net/gml/3.2 degree,omitempty"`
	Knot                  []KnotPropertyType      `xml:"http://www.opengis.net/gml/3.2 knot"`
	PosList               *DirectPositionListType `xml:"http://www.opengis.net/gml/3.2 posList,omitempty"`
	Coordinates           *CoordinatesType        `xml:"http://www.opengis.net/gml/3.2 coordinates,omitempty"`
	Pos                   *DirectPositionType     `xml:"http://www.opengis.net/gml/3.2 pos,omitempty"`
	PointProperty         *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointProperty,omitempty"`
	PointRep              *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointRep,omitempty"`
	Interpolation         string                  `xml:"interpolation,attr,omitempty"`
	IsPolynomial          bool                    `xml:"isPolynomial,attr,omitempty"`
	KnotType              string                  `xml:"knotType,attr,omitempty"`
}

type BagType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	Member               []AssociationRoleType  `xml:"http://www.opengis.net/gml/3.2 member"`
	Members              *ArrayAssociationType  `xml:"http://www.opengis.net/gml/3.2 members,omitempty"`
}

type BaseUnitType struct {
	MetaDataProperty      []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description           *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference  *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier            *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                  []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                    string                 `xml:"id,attr,omitempty"`
	Remarks               *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	QuantityType          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 quantityType,omitempty"`
	QuantityTypeReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 quantityTypeReference,omitempty"`
	CatalogSymbol         *CodeType              `xml:"http://www.opengis.net/gml/3.2 catalogSymbol,omitempty"`
	UnitsSystem           *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 unitsSystem,omitempty"`
}

type BezierType struct {
	NumDerivativesAtStart int                     `xml:"numDerivativesAtStart,attr,omitempty"`
	NumDerivativesAtEnd   int                     `xml:"numDerivativesAtEnd,attr,omitempty"`
	NumDerivativeInterior int                     `xml:"numDerivativeInterior,attr,omitempty"`
	Degree                int                     `xml:"http://www.opengis.net/gml/3.2 degree,omitempty"`
	Knot                  []KnotPropertyType      `xml:"http://www.opengis.net/gml/3.2 knot"`
	PosList               *DirectPositionListType `xml:"http://www.opengis.net/gml/3.2 posList,omitempty"`
	Coordinates           *CoordinatesType        `xml:"http://www.opengis.net/gml/3.2 coordinates,omitempty"`
	Pos                   *DirectPositionType     `xml:"http://www.opengis.net/gml/3.2 pos,omitempty"`
	PointProperty         *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointProperty,omitempty"`
	PointRep              *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointRep,omitempty"`
	Interpolation         string                  `xml:"interpolation,attr,omitempty"`
	IsPolynomial          bool                    `xml:"isPolynomial,attr,omitempty"`
	KnotType              string                  `xml:"knotType,attr,omitempty"`
}

type BooleanPropertyType struct {
	Boolean      *string `xml:"http://www.opengis.net/gml/3.2 Boolean,omitempty"`
	NilReason    string  `xml:"nilReason,attr,omitempty"`
	RemoteSchema string  `xml:"remoteSchema,attr,omitempty"`
}

type BoundedFeatureType struct {
	MetaDataProperty     []MetaDataPropertyType        `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType              `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType        `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                    `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                        `xml:"id,attr,omitempty"`
	BoundedBy            *BoundingShapeType            `xml:"http://www.opengis.net/gml/3.2 boundedBy,omitempty"`
	Location             *LocationPropertyType         `xml:"http://www.opengis.net/gml/3.2 location,omitempty"`
	PriorityLocation     *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml/3.2 priorityLocation,omitempty"`
}

type BoundingShapeType struct {
	Envelope               *EnvelopeType               `xml:"http://www.opengis.net/gml/3.2 Envelope,omitempty"`
	EnvelopeWithTimePeriod *EnvelopeWithTimePeriodType `xml:"http://www.opengis.net/gml/3.2 EnvelopeWithTimePeriod,omitempty"`
	Null                   *string                     `xml:"http://www.opengis.net/gml/3.2 Null,omitempty"`
	NilReason              string                      `xml:"nilReason,attr,omitempty"`
}

type CRSPropertyType struct {
	AbstractCRS       *AbstractCRSType `xml:"http://www.opengis.net/gml/3.2 AbstractCRS,omitempty"`
	AbstractSingleCRS *AbstractCRSType `xml:"http://www.opengis.net/gml/3.2 AbstractSingleCRS,omitempty"`
	CompoundCRS       *CompoundCRSType `xml:"http://www.opengis.net/gml/3.2 CompoundCRS,omitempty"`
	NilReason         string           `xml:"nilReason,attr,omitempty"`
	RemoteSchema      string           `xml:"remoteSchema,attr,omitempty"`
}

type CartesianCSPropertyType struct {
	CartesianCS  *CartesianCSType `xml:"http://www.opengis.net/gml/3.2 CartesianCS,omitempty"`
	NilReason    string           `xml:"nilReason,attr,omitempty"`
	RemoteSchema string           `xml:"remoteSchema,attr,omitempty"`
}

type CartesianCSType struct {
	MetaDataProperty     []MetaDataPropertyType             `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType                   `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                     `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType             `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                         `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                             `xml:"id,attr,omitempty"`
	Remarks              *string                            `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	Axis                 []CoordinateSystemAxisPropertyType `xml:"http://www.opengis.net/gml/3.2 axis"`
	UsesAxis             *CoordinateSystemAxisPropertyType  `xml:"http://www.opengis.net/gml/3.2 usesAxis,omitempty"`
	AggregationType      string                             `xml:"aggregationType,attr,omitempty"`
}

type CategoryExtentType struct {
	Value string `xml:",chardata"`
}

type CategoryPropertyType struct {
	Category     *string `xml:"http://www.opengis.net/gml/3.2 Category,omitempty"`
	NilReason    string  `xml:"nilReason,attr,omitempty"`
	RemoteSchema string  `xml:"remoteSchema,attr,omitempty"`
}

type CircleByCenterPointType struct {
	NumDerivativesAtStart int                     `xml:"numDerivativesAtStart,attr,omitempty"`
	NumDerivativesAtEnd   int                     `xml:"numDerivativesAtEnd,attr,omitempty"`
	NumDerivativeInterior int                     `xml:"numDerivativeInterior,attr,omitempty"`
	Radius                *LengthType             `xml:"http://www.opengis.net/gml/3.2 radius,omitempty"`
	StartAngle            *AngleType              `xml:"http://www.opengis.net/gml/3.2 startAngle,omitempty"`
	EndAngle              *AngleType              `xml:"http://www.opengis.net/gml/3.2 endAngle,omitempty"`
	PosList               *DirectPositionListType `xml:"http://www.opengis.net/gml/3.2 posList,omitempty"`
	Coordinates           *CoordinatesType        `xml:"http://www.opengis.net/gml/3.2 coordinates,omitempty"`
	Pos                   *DirectPositionType     `xml:"http://www.opengis.net/gml/3.2 pos,omitempty"`
	PointProperty         *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointProperty,omitempty"`
	PointRep              *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointRep,omitempty"`
	Interpolation         string                  `xml:"interpolation,attr,omitempty"`
	NumArc                int                     `xml:"numArc,attr"`
}

type CircleType struct {
	NumDerivativesAtStart int                     `xml:"numDerivativesAtStart,attr,omitempty"`
	NumDerivativesAtEnd   int                     `xml:"numDerivativesAtEnd,attr,omitempty"`
	NumDerivativeInterior int                     `xml:"numDerivativeInterior,attr,omitempty"`
	PosList               *DirectPositionListType `xml:"http://www.opengis.net/gml/3.2 posList,omitempty"`
	Coordinates           *CoordinatesType        `xml:"http://www.opengis.net/gml/3.2 coordinates,omitempty"`
	Pos                   *DirectPositionType     `xml:"http://www.opengis.net/gml/3.2 pos,omitempty"`
	PointProperty         *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointProperty,omitempty"`
	PointRep              *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointRep,omitempty"`
	Interpolation         string                  `xml:"interpolation,attr,omitempty"`
	NumArc                int                     `xml:"numArc,attr,omitempty"`
}

type ClothoidType struct {
	NumDerivativesAtStart int     `xml:"numDerivativesAtStart,attr,omitempty"`
	NumDerivativesAtEnd   int     `xml:"numDerivativesAtEnd,attr,omitempty"`
	NumDerivativeInterior int     `xml:"numDerivativeInterior,attr,omitempty"`
	RefLocation           string  `xml:"http://www.opengis.net/gml/3.2 refLocation,omitempty"`
	ScaleFactor           float64 `xml:"http://www.opengis.net/gml/3.2 scaleFactor,omitempty"`
	StartParameter        float64 `xml:"http://www.opengis.net/gml/3.2 startParameter,omitempty"`
	EndParameter          float64 `xml:"http://www.opengis.net/gml/3.2 endParameter,omitempty"`
	Interpolation         string  `xml:"interpolation,attr,omitempty"`
}

type CodeListType struct {
	Value     string `xml:",chardata"`
	CodeSpace string `xml:"codeSpace,attr,omitempty"`
}

type CodeOrNilReasonListType struct {
	Value     string `xml:",chardata"`
	CodeSpace string `xml:"codeSpace,attr,omitempty"`
}

type CodeType struct {
	Value     string `xml:",chardata"`
	CodeSpace string `xml:"codeSpace,attr,omitempty"`
}

type CodeWithAuthorityType struct {
	Value     string `xml:",chardata"`
	CodeSpace string `xml:"codeSpace,attr"`
}

type CompositeCurveType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	SrsName              string                 `xml:"srsName,attr,omitempty"`
	SrsDimension         int                    `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                 `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                 `xml:"uomLabels,attr,omitempty"`
	CurveMember          []CurvePropertyType    `xml:"http://www.opengis.net/gml/3.2 curveMember"`
	AggregationType      string                 `xml:"aggregationType,attr,omitempty"`
}

type CompositeSolidType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	SrsName              string                 `xml:"srsName,attr,omitempty"`
	SrsDimension         int                    `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                 `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                 `xml:"uomLabels,attr,omitempty"`
	SolidMember          []SolidPropertyType    `xml:"http://www.opengis.net/gml/3.2 solidMember"`
	AggregationType      string                 `xml:"aggregationType,attr,omitempty"`
}

type CompositeSurfaceType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	SrsName              string                 `xml:"srsName,attr,omitempty"`
	SrsDimension         int                    `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                 `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                 `xml:"uomLabels,attr,omitempty"`
	SurfaceMember        []SurfacePropertyType  `xml:"http://www.opengis.net/gml/3.2 surfaceMember"`
	AggregationType      string                 `xml:"aggregationType,attr,omitempty"`
}

type CompositeValueType struct {
	MetaDataProperty     []MetaDataPropertyType  `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType        `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType          `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType  `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType              `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                  `xml:"id,attr,omitempty"`
	ValueComponent       []ValuePropertyType     `xml:"http://www.opengis.net/gml/3.2 valueComponent"`
	ValueComponents      *ValueArrayPropertyType `xml:"http://www.opengis.net/gml/3.2 valueComponents,omitempty"`
	AggregationType      string                  `xml:"aggregationType,attr,omitempty"`
}

type CompoundCRSPropertyType struct {
	CompoundCRS  *CompoundCRSType `xml:"http://www.opengis.net/gml/3.2 CompoundCRS,omitempty"`
	NilReason    string           `xml:"nilReason,attr,omitempty"`
	RemoteSchema string           `xml:"remoteSchema,attr,omitempty"`
}

type CompoundCRSType struct {
	MetaDataProperty         []MetaDataPropertyType  `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description              *StringOrRefType        `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference     *ReferenceType          `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier               *CodeWithAuthorityType  `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                     []CodeType              `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                       string                  `xml:"id,attr,omitempty"`
	Remarks                  *string                 `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity         []string                `xml:"http://www.opengis.net/gml/3.2 domainOfValidity"`
	Scope                    []string                `xml:"http://www.opengis.net/gml/3.2 scope"`
	ComponentReferenceSystem []SingleCRSPropertyType `xml:"http://www.opengis.net/gml/3.2 componentReferenceSystem"`
	IncludesSingleCRS        *SingleCRSPropertyType  `xml:"http://www.opengis.net/gml/3.2 includesSingleCRS,omitempty"`
	AggregationType          string                  `xml:"aggregationType,attr,omitempty"`
}

type ConcatenatedOperationPropertyType struct {
	ConcatenatedOperation *ConcatenatedOperationType `xml:"http://www.opengis.net/gml/3.2 ConcatenatedOperation,omitempty"`
	NilReason             string                     `xml:"nilReason,attr,omitempty"`
	RemoteSchema          string                     `xml:"remoteSchema,attr,omitempty"`
}

type ConcatenatedOperationType struct {
	MetaDataProperty            []MetaDataPropertyType            `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description                 *StringOrRefType                  `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference        *ReferenceType                    `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier                  *CodeWithAuthorityType            `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                        []CodeType                        `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                          string                            `xml:"id,attr,omitempty"`
	Remarks                     *string                           `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity            *string                           `xml:"http://www.opengis.net/gml/3.2 domainOfValidity,omitempty"`
	Scope                       []string                          `xml:"http://www.opengis.net/gml/3.2 scope"`
	OperationVersion            *string                           `xml:"http://www.opengis.net/gml/3.2 operationVersion,omitempty"`
	CoordinateOperationAccuracy []string                          `xml:"http://www.opengis.net/gml/3.2 coordinateOperationAccuracy"`
	SourceCRS                   *CRSPropertyType                  `xml:"http://www.opengis.net/gml/3.2 sourceCRS,omitempty"`
	TargetCRS                   *CRSPropertyType                  `xml:"http://www.opengis.net/gml/3.2 targetCRS,omitempty"`
	CoordOperation              []CoordinateOperationPropertyType `xml:"http://www.opengis.net/gml/3.2 coordOperation"`
	UsesOperation               *CoordinateOperationPropertyType  `xml:"http://www.opengis.net/gml/3.2 usesOperation,omitempty"`
	UsesSingleOperation         *CoordinateOperationPropertyType  `xml:"http://www.opengis.net/gml/3.2 usesSingleOperation,omitempty"`
	AggregationType             string                            `xml:"aggregationType,attr,omitempty"`
}

type ConeType struct {
	AggregationType     string `xml:"aggregationType,attr,omitempty"`
	Rows                string `xml:"http://www.opengis.net/gml/3.2 rows,omitempty"`
	Columns             int    `xml:"columns,attr,omitempty"`
	HorizontalCurveType string `xml:"horizontalCurveType,attr,omitempty"`
	VerticalCurveType   string `xml:"verticalCurveType,attr,omitempty"`
}

type ConventionalUnitType struct {
	MetaDataProperty               []MetaDataPropertyType         `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description                    *StringOrRefType               `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference           *ReferenceType                 `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier                     *CodeWithAuthorityType         `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                           []CodeType                     `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                             string                         `xml:"id,attr,omitempty"`
	Remarks                        *string                        `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	QuantityType                   *StringOrRefType               `xml:"http://www.opengis.net/gml/3.2 quantityType,omitempty"`
	QuantityTypeReference          *ReferenceType                 `xml:"http://www.opengis.net/gml/3.2 quantityTypeReference,omitempty"`
	CatalogSymbol                  *CodeType                      `xml:"http://www.opengis.net/gml/3.2 catalogSymbol,omitempty"`
	DerivationUnitTerm             []DerivationUnitTermType       `xml:"http://www.opengis.net/gml/3.2 derivationUnitTerm"`
	ConversionToPreferredUnit      *ConversionToPreferredUnitType `xml:"http://www.opengis.net/gml/3.2 conversionToPreferredUnit,omitempty"`
	RoughConversionToPreferredUnit *ConversionToPreferredUnitType `xml:"http://www.opengis.net/gml/3.2 roughConversionToPreferredUnit,omitempty"`
}

type ConversionPropertyType struct {
	Conversion   *ConversionType `xml:"http://www.opengis.net/gml/3.2 Conversion,omitempty"`
	NilReason    string          `xml:"nilReason,attr,omitempty"`
	RemoteSchema string          `xml:"remoteSchema,attr,omitempty"`
}

type ConversionToPreferredUnitType struct {
	Uom     string       `xml:"uom,attr"`
	Factor  float64      `xml:"http://www.opengis.net/gml/3.2 factor,omitempty"`
	Formula *FormulaType `xml:"http://www.opengis.net/gml/3.2 formula,omitempty"`
}

type ConversionType struct {
	MetaDataProperty            []MetaDataPropertyType                      `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description                 *StringOrRefType                            `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference        *ReferenceType                              `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier                  *CodeWithAuthorityType                      `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                        []CodeType                                  `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                          string                                      `xml:"id,attr,omitempty"`
	Remarks                     *string                                     `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity            *string                                     `xml:"http://www.opengis.net/gml/3.2 domainOfValidity,omitempty"`
	Scope                       []string                                    `xml:"http://www.opengis.net/gml/3.2 scope"`
	OperationVersion            *string                                     `xml:"http://www.opengis.net/gml/3.2 operationVersion,omitempty"`
	CoordinateOperationAccuracy []string                                    `xml:"http://www.opengis.net/gml/3.2 coordinateOperationAccuracy"`
	SourceCRS                   *CRSPropertyType                            `xml:"http://www.opengis.net/gml/3.2 sourceCRS,omitempty"`
	TargetCRS                   *CRSPropertyType                            `xml:"http://www.opengis.net/gml/3.2 targetCRS,omitempty"`
	Method                      *OperationMethodPropertyType                `xml:"http://www.opengis.net/gml/3.2 method,omitempty"`
	UsesMethod                  *OperationMethodPropertyType                `xml:"http://www.opengis.net/gml/3.2 usesMethod,omitempty"`
	ParameterValue              []AbstractGeneralParameterValuePropertyType `xml:"http://www.opengis.net/gml/3.2 parameterValue"`
	IncludesValue               *AbstractGeneralParameterValuePropertyType  `xml:"http://www.opengis.net/gml/3.2 includesValue,omitempty"`
	UsesValue                   *AbstractGeneralParameterValuePropertyType  `xml:"http://www.opengis.net/gml/3.2 usesValue,omitempty"`
}

type CoordinateOperationPropertyType struct {
	AbstractCoordinateOperation *AbstractCoordinateOperationType `xml:"http://www.opengis.net/gml/3.2 AbstractCoordinateOperation,omitempty"`
	AbstractSingleOperation     *AbstractCoordinateOperationType `xml:"http://www.opengis.net/gml/3.2 AbstractSingleOperation,omitempty"`
	ConcatenatedOperation       *ConcatenatedOperationType       `xml:"http://www.opengis.net/gml/3.2 ConcatenatedOperation,omitempty"`
	NilReason                   string                           `xml:"nilReason,attr,omitempty"`
	RemoteSchema                string                           `xml:"remoteSchema,attr,omitempty"`
}

type CoordinateSystemAxisPropertyType struct {
	CoordinateSystemAxis *CoordinateSystemAxisType `xml:"http://www.opengis.net/gml/3.2 CoordinateSystemAxis,omitempty"`
	NilReason            string                    `xml:"nilReason,attr,omitempty"`
	RemoteSchema         string                    `xml:"remoteSchema,attr,omitempty"`
}

type CoordinateSystemAxisType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	Remarks              *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	AxisAbbrev           *CodeType              `xml:"http://www.opengis.net/gml/3.2 axisAbbrev,omitempty"`
	AxisDirection        *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 axisDirection,omitempty"`
	MinimumValue         *float64               `xml:"http://www.opengis.net/gml/3.2 minimumValue,omitempty"`
	MaximumValue         *float64               `xml:"http://www.opengis.net/gml/3.2 maximumValue,omitempty"`
	RangeMeaning         *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 rangeMeaning,omitempty"`
	Uom                  string                 `xml:"uom,attr"`
}

type CoordinateSystemPropertyType struct {
	AbstractCoordinateSystem *AbstractCoordinateSystemType `xml:"http://www.opengis.net/gml/3.2 AbstractCoordinateSystem,omitempty"`
	AffineCS                 *AffineCSType                 `xml:"http://www.opengis.net/gml/3.2 AffineCS,omitempty"`
	CartesianCS              *CartesianCSType              `xml:"http://www.opengis.net/gml/3.2 CartesianCS,omitempty"`
	CylindricalCS            *CylindricalCSType            `xml:"http://www.opengis.net/gml/3.2 CylindricalCS,omitempty"`
	EllipsoidalCS            *EllipsoidalCSType            `xml:"http://www.opengis.net/gml/3.2 EllipsoidalCS,omitempty"`
	LinearCS                 *LinearCSType                 `xml:"http://www.opengis.net/gml/3.2 LinearCS,omitempty"`
	ObliqueCartesianCS       *ObliqueCartesianCSType       `xml:"http://www.opengis.net/gml/3.2 ObliqueCartesianCS,omitempty"`
	PolarCS                  *PolarCSType                  `xml:"http://www.opengis.net/gml/3.2 PolarCS,omitempty"`
	SphericalCS              *SphericalCSType              `xml:"http://www.opengis.net/gml/3.2 SphericalCS,omitempty"`
	TemporalCS               *TemporalCSType               `xml:"http://www.opengis.net/gml/3.2 TemporalCS,omitempty"`
	TimeCS                   *TimeCSType                   `xml:"http://www.opengis.net/gml/3.2 TimeCS,omitempty"`
	UserDefinedCS            *UserDefinedCSType            `xml:"http://www.opengis.net/gml/3.2 UserDefinedCS,omitempty"`
	VerticalCS               *VerticalCSType               `xml:"http://www.opengis.net/gml/3.2 VerticalCS,omitempty"`
	NilReason                string                        `xml:"nilReason,attr,omitempty"`
	RemoteSchema             string                        `xml:"remoteSchema,attr,omitempty"`
}

type CoordinatesType struct {
	Value   string `xml:",chardata"`
	Decimal string `xml:"decimal,attr,omitempty"`
	Cs      string `xml:"cs,attr,omitempty"`
	Ts      string `xml:"ts,attr,omitempty"`
}

type CountPropertyType struct {
	Count        *string `xml:"http://www.opengis.net/gml/3.2 Count,omitempty"`
	NilReason    string  `xml:"nilReason,attr,omitempty"`
	RemoteSchema string  `xml:"remoteSchema,attr,omitempty"`
}

type CoverageFunctionType struct {
	MappingRule         *StringOrRefType  `xml:"http://www.opengis.net/gml/3.2 MappingRule,omitempty"`
	CoverageMappingRule *MappingRuleType  `xml:"http://www.opengis.net/gml/3.2 CoverageMappingRule,omitempty"`
	GridFunction        *GridFunctionType `xml:"http://www.opengis.net/gml/3.2 GridFunction,omitempty"`
}

type CubicSplineType struct {
	NumDerivativesAtStart int                     `xml:"numDerivativesAtStart,attr,omitempty"`
	NumDerivativesAtEnd   int                     `xml:"numDerivativesAtEnd,attr,omitempty"`
	NumDerivativeInterior int                     `xml:"numDerivativeInterior,attr,omitempty"`
	VectorAtStart         *VectorType             `xml:"http://www.opengis.net/gml/3.2 vectorAtStart,omitempty"`
	VectorAtEnd           *VectorType             `xml:"http://www.opengis.net/gml/3.2 vectorAtEnd,omitempty"`
	PosList               *DirectPositionListType `xml:"http://www.opengis.net/gml/3.2 posList,omitempty"`
	Coordinates           *CoordinatesType        `xml:"http://www.opengis.net/gml/3.2 coordinates,omitempty"`
	Pos                   *DirectPositionType     `xml:"http://www.opengis.net/gml/3.2 pos,omitempty"`
	PointProperty         *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointProperty,omitempty"`
	PointRep              *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointRep,omitempty"`
	Interpolation         string                  `xml:"interpolation,attr,omitempty"`
	Degree                int                     `xml:"degree,attr,omitempty"`
}

type CurveArrayPropertyType struct {
	AbstractCurve   *AbstractCurveType   `xml:"http://www.opengis.net/gml/3.2 AbstractCurve,omitempty"`
	AbstractRing    *AbstractRingType    `xml:"http://www.opengis.net/gml/3.2 AbstractRing,omitempty"`
	CompositeCurve  *CompositeCurveType  `xml:"http://www.opengis.net/gml/3.2 CompositeCurve,omitempty"`
	Curve           *CurveType           `xml:"http://www.opengis.net/gml/3.2 Curve,omitempty"`
	LineString      *LineStringType      `xml:"http://www.opengis.net/gml/3.2 LineString,omitempty"`
	OrientableCurve *OrientableCurveType `xml:"http://www.opengis.net/gml/3.2 OrientableCurve,omitempty"`
	Owns            bool                 `xml:"owns,attr,omitempty"`
}

type CurvePropertyType struct {
	AbstractCurve   *AbstractCurveType   `xml:"http://www.opengis.net/gml/3.2 AbstractCurve,omitempty"`
	AbstractRing    *AbstractRingType    `xml:"http://www.opengis.net/gml/3.2 AbstractRing,omitempty"`
	CompositeCurve  *CompositeCurveType  `xml:"http://www.opengis.net/gml/3.2 CompositeCurve,omitempty"`
	Curve           *CurveType           `xml:"http://www.opengis.net/gml/3.2 Curve,omitempty"`
	LineString      *LineStringType      `xml:"http://www.opengis.net/gml/3.2 LineString,omitempty"`
	OrientableCurve *OrientableCurveType `xml:"http://www.opengis.net/gml/3.2 OrientableCurve,omitempty"`
	NilReason       string               `xml:"nilReason,attr,omitempty"`
	RemoteSchema    string               `xml:"remoteSchema,attr,omitempty"`
	Owns            bool                 `xml:"owns,attr,omitempty"`
}

type CurveSegmentArrayPropertyType struct {
	AbstractCurveSegment *AbstractCurveSegmentType `xml:"http://www.opengis.net/gml/3.2 AbstractCurveSegment,omitempty"`
	ArcByCenterPoint     *ArcByCenterPointType     `xml:"http://www.opengis.net/gml/3.2 ArcByCenterPoint,omitempty"`
	ArcString            *ArcStringType            `xml:"http://www.opengis.net/gml/3.2 ArcString,omitempty"`
	ArcStringByBulge     *ArcStringByBulgeType     `xml:"http://www.opengis.net/gml/3.2 ArcStringByBulge,omitempty"`
	BSpline              *BSplineType              `xml:"http://www.opengis.net/gml/3.2 BSpline,omitempty"`
	Clothoid             *ClothoidType             `xml:"http://www.opengis.net/gml/3.2 Clothoid,omitempty"`
	CubicSpline          *CubicSplineType          `xml:"http://www.opengis.net/gml/3.2 CubicSpline,omitempty"`
	GeodesicString       *GeodesicStringType       `xml:"http://www.opengis.net/gml/3.2 GeodesicString,omitempty"`
	LineStringSegment    *LineStringSegmentType    `xml:"http://www.opengis.net/gml/3.2 LineStringSegment,omitempty"`
	OffsetCurve          *OffsetCurveType          `xml:"http://www.opengis.net/gml/3.2 OffsetCurve,omitempty"`
}

type CurveType struct {
	MetaDataProperty     []MetaDataPropertyType         `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType               `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                 `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType         `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                     `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                         `xml:"id,attr,omitempty"`
	SrsName              string                         `xml:"srsName,attr,omitempty"`
	SrsDimension         int                            `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                         `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                         `xml:"uomLabels,attr,omitempty"`
	Segments             *CurveSegmentArrayPropertyType `xml:"http://www.opengis.net/gml/3.2 segments,omitempty"`
}

type CylinderType struct {
	AggregationType     string `xml:"aggregationType,attr,omitempty"`
	Rows                string `xml:"http://www.opengis.net/gml/3.2 rows,omitempty"`
	Columns             int    `xml:"columns,attr,omitempty"`
	HorizontalCurveType string `xml:"horizontalCurveType,attr,omitempty"`
	VerticalCurveType   string `xml:"verticalCurveType,attr,omitempty"`
}

type CylindricalCSPropertyType struct {
	CylindricalCS *CylindricalCSType `xml:"http://www.opengis.net/gml/3.2 CylindricalCS,omitempty"`
	NilReason     string             `xml:"nilReason,attr,omitempty"`
	RemoteSchema  string             `xml:"remoteSchema,attr,omitempty"`
}

type CylindricalCSType struct {
	MetaDataProperty     []MetaDataPropertyType             `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType                   `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                     `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType             `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                         `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                             `xml:"id,attr,omitempty"`
	Remarks              *string                            `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	Axis                 []CoordinateSystemAxisPropertyType `xml:"http://www.opengis.net/gml/3.2 axis"`
	UsesAxis             *CoordinateSystemAxisPropertyType  `xml:"http://www.opengis.net/gml/3.2 usesAxis,omitempty"`
	AggregationType      string                             `xml:"aggregationType,attr,omitempty"`
}

type DMSAngleType struct {
	Degrees        *DegreesType `xml:"http://www.opengis.net/gml/3.2 degrees,omitempty"`
	DecimalMinutes *float64     `xml:"http://www.opengis.net/gml/3.2 decimalMinutes,omitempty"`
	Minutes        *int         `xml:"http://www.opengis.net/gml/3.2 minutes,omitempty"`
	Seconds        *float64     `xml:"http://www.opengis.net/gml/3.2 seconds,omitempty"`
}

type DataBlockType struct {
	RangeParameters            *AssociationRoleType `xml:"http://www.opengis.net/gml/3.2 rangeParameters,omitempty"`
	TupleList                  *CoordinatesType     `xml:"http://www.opengis.net/gml/3.2 tupleList,omitempty"`
	DoubleOrNilReasonTupleList *string              `xml:"http://www.opengis.net/gml/3.2 doubleOrNilReasonTupleList,omitempty"`
}

type DatumPropertyType struct {
	AbstractDatum    *AbstractDatumType    `xml:"http://www.opengis.net/gml/3.2 AbstractDatum,omitempty"`
	EngineeringDatum *EngineeringDatumType `xml:"http://www.opengis.net/gml/3.2 EngineeringDatum,omitempty"`
	GeodeticDatum    *GeodeticDatumType    `xml:"http://www.opengis.net/gml/3.2 GeodeticDatum,omitempty"`
	ImageDatum       *ImageDatumType       `xml:"http://www.opengis.net/gml/3.2 ImageDatum,omitempty"`
	TemporalDatum    *TemporalDatumType    `xml:"http://www.opengis.net/gml/3.2 TemporalDatum,omitempty"`
	VerticalDatum    *VerticalDatumType    `xml:"http://www.opengis.net/gml/3.2 VerticalDatum,omitempty"`
	NilReason        string                `xml:"nilReason,attr,omitempty"`
	RemoteSchema     string                `xml:"remoteSchema,attr,omitempty"`
}

type DefinitionBaseType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
}

type DefinitionProxyType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	Remarks              *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DefinitionRef        *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 definitionRef,omitempty"`
}

type DefinitionType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	Remarks              *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
}

type DegreesType struct {
	Value     int    `xml:",chardata"`
	Direction string `xml:"direction,attr,omitempty"`
}

type DerivationUnitTermType struct {
	Uom      string `xml:"uom,attr"`
	Exponent int    `xml:"exponent,attr,omitempty"`
}

type DerivedCRSPropertyType struct {
	DerivedCRS   *DerivedCRSType `xml:"http://www.opengis.net/gml/3.2 DerivedCRS,omitempty"`
	NilReason    string          `xml:"nilReason,attr,omitempty"`
	RemoteSchema string          `xml:"remoteSchema,attr,omitempty"`
}

type DerivedCRSType struct {
	MetaDataProperty     []MetaDataPropertyType         `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType               `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                 `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType         `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                     `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                         `xml:"id,attr,omitempty"`
	Remarks              *string                        `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     []string                       `xml:"http://www.opengis.net/gml/3.2 domainOfValidity"`
	Scope                []string                       `xml:"http://www.opengis.net/gml/3.2 scope"`
	Conversion           *GeneralConversionPropertyType `xml:"http://www.opengis.net/gml/3.2 conversion,omitempty"`
	DefinedByConversion  *GeneralConversionPropertyType `xml:"http://www.opengis.net/gml/3.2 definedByConversion,omitempty"`
	BaseCRS              *SingleCRSPropertyType         `xml:"http://www.opengis.net/gml/3.2 baseCRS,omitempty"`
	DerivedCRSType       *CodeWithAuthorityType         `xml:"http://www.opengis.net/gml/3.2 derivedCRSType,omitempty"`
	CoordinateSystem     *CoordinateSystemPropertyType  `xml:"http://www.opengis.net/gml/3.2 coordinateSystem,omitempty"`
	UsesCS               *CoordinateSystemPropertyType  `xml:"http://www.opengis.net/gml/3.2 usesCS,omitempty"`
}

type DerivedUnitType struct {
	MetaDataProperty      []MetaDataPropertyType   `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description           *StringOrRefType         `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference  *ReferenceType           `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier            *CodeWithAuthorityType   `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                  []CodeType               `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                    string                   `xml:"id,attr,omitempty"`
	Remarks               *string                  `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	QuantityType          *StringOrRefType         `xml:"http://www.opengis.net/gml/3.2 quantityType,omitempty"`
	QuantityTypeReference *ReferenceType           `xml:"http://www.opengis.net/gml/3.2 quantityTypeReference,omitempty"`
	CatalogSymbol         *CodeType                `xml:"http://www.opengis.net/gml/3.2 catalogSymbol,omitempty"`
	DerivationUnitTerm    []DerivationUnitTermType `xml:"http://www.opengis.net/gml/3.2 derivationUnitTerm"`
}

type DictionaryEntryType struct {
	Owns                              bool                                   `xml:"owns,attr,omitempty"`
	Definition                        *DefinitionType                        `xml:"http://www.opengis.net/gml/3.2 Definition,omitempty"`
	AbstractCRS                       *AbstractCRSType                       `xml:"http://www.opengis.net/gml/3.2 AbstractCRS,omitempty"`
	AbstractCoordinateOperation       *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml/3.2 AbstractCoordinateOperation,omitempty"`
	AbstractCoordinateSystem          *AbstractCoordinateSystemType          `xml:"http://www.opengis.net/gml/3.2 AbstractCoordinateSystem,omitempty"`
	AbstractDatum                     *AbstractDatumType                     `xml:"http://www.opengis.net/gml/3.2 AbstractDatum,omitempty"`
	AbstractGeneralOperationParameter *AbstractGeneralOperationParameterType `xml:"http://www.opengis.net/gml/3.2 AbstractGeneralOperationParameter,omitempty"`
	CoordinateSystemAxis              *CoordinateSystemAxisType              `xml:"http://www.opengis.net/gml/3.2 CoordinateSystemAxis,omitempty"`
	DefinitionCollection              *DictionaryType                        `xml:"http://www.opengis.net/gml/3.2 DefinitionCollection,omitempty"`
	DefinitionProxy                   *DefinitionProxyType                   `xml:"http://www.opengis.net/gml/3.2 DefinitionProxy,omitempty"`
	Dictionary                        *DictionaryType                        `xml:"http://www.opengis.net/gml/3.2 Dictionary,omitempty"`
	Ellipsoid                         *EllipsoidType                         `xml:"http://www.opengis.net/gml/3.2 Ellipsoid,omitempty"`
	OperationMethod                   *OperationMethodType                   `xml:"http://www.opengis.net/gml/3.2 OperationMethod,omitempty"`
	PrimeMeridian                     *PrimeMeridianType                     `xml:"http://www.opengis.net/gml/3.2 PrimeMeridian,omitempty"`
	TimeReferenceSystem               *TimeReferenceSystemType               `xml:"http://www.opengis.net/gml/3.2 TimeReferenceSystem,omitempty"`
	UnitDefinition                    *UnitDefinitionType                    `xml:"http://www.opengis.net/gml/3.2 UnitDefinition,omitempty"`
	NilReason                         string                                 `xml:"nilReason,attr,omitempty"`
	RemoteSchema                      string                                 `xml:"remoteSchema,attr,omitempty"`
}

type DictionaryType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	Remarks              *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DictionaryEntry      *DictionaryEntryType   `xml:"http://www.opengis.net/gml/3.2 dictionaryEntry,omitempty"`
	DefinitionMember     *DictionaryEntryType   `xml:"http://www.opengis.net/gml/3.2 definitionMember,omitempty"`
	IndirectEntry        *IndirectEntryType     `xml:"http://www.opengis.net/gml/3.2 indirectEntry,omitempty"`
	AggregationType      string                 `xml:"aggregationType,attr,omitempty"`
}

type DirectPositionListType struct {
	Value        string `xml:",chardata"`
	Count        int    `xml:"count,attr,omitempty"`
	SrsName      string `xml:"srsName,attr,omitempty"`
	SrsDimension int    `xml:"srsDimension,attr,omitempty"`
	AxisLabels   string `xml:"axisLabels,attr,omitempty"`
	UomLabels    string `xml:"uomLabels,attr,omitempty"`
}

type DirectPositionType struct {
	Value        string `xml:",chardata"`
	SrsName      string `xml:"srsName,attr,omitempty"`
	SrsDimension int    `xml:"srsDimension,attr,omitempty"`
	AxisLabels   string `xml:"axisLabels,attr,omitempty"`
	UomLabels    string `xml:"uomLabels,attr,omitempty"`
}

type DirectedEdgePropertyType struct {
	Edge         *EdgeType `xml:"http://www.opengis.net/gml/3.2 Edge,omitempty"`
	Orientation  string    `xml:"orientation,attr,omitempty"`
	NilReason    string    `xml:"nilReason,attr,omitempty"`
	RemoteSchema string    `xml:"remoteSchema,attr,omitempty"`
	Owns         bool      `xml:"owns,attr,omitempty"`
}

type DirectedFacePropertyType struct {
	Face         *FaceType `xml:"http://www.opengis.net/gml/3.2 Face,omitempty"`
	Orientation  string    `xml:"orientation,attr,omitempty"`
	NilReason    string    `xml:"nilReason,attr,omitempty"`
	RemoteSchema string    `xml:"remoteSchema,attr,omitempty"`
	Owns         bool      `xml:"owns,attr,omitempty"`
}

type DirectedNodePropertyType struct {
	Node         *NodeType `xml:"http://www.opengis.net/gml/3.2 Node,omitempty"`
	Orientation  string    `xml:"orientation,attr,omitempty"`
	NilReason    string    `xml:"nilReason,attr,omitempty"`
	RemoteSchema string    `xml:"remoteSchema,attr,omitempty"`
	Owns         bool      `xml:"owns,attr,omitempty"`
}

type DirectedObservationAtDistanceType struct {
	MetaDataProperty     []MetaDataPropertyType        `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType              `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType        `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                    `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                        `xml:"id,attr,omitempty"`
	BoundedBy            *BoundingShapeType            `xml:"http://www.opengis.net/gml/3.2 boundedBy,omitempty"`
	Location             *LocationPropertyType         `xml:"http://www.opengis.net/gml/3.2 location,omitempty"`
	PriorityLocation     *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml/3.2 priorityLocation,omitempty"`
	ValidTime            *TimePrimitivePropertyType    `xml:"http://www.opengis.net/gml/3.2 validTime,omitempty"`
	Using                *ProcedurePropertyType        `xml:"http://www.opengis.net/gml/3.2 using,omitempty"`
	Target               *TargetPropertyType           `xml:"http://www.opengis.net/gml/3.2 target,omitempty"`
	Subject              *TargetPropertyType           `xml:"http://www.opengis.net/gml/3.2 subject,omitempty"`
	ResultOf             *ResultType                   `xml:"http://www.opengis.net/gml/3.2 resultOf,omitempty"`
	Direction            *DirectionPropertyType        `xml:"http://www.opengis.net/gml/3.2 direction,omitempty"`
	Distance             *MeasureType                  `xml:"http://www.opengis.net/gml/3.2 distance,omitempty"`
}

type DirectedObservationType struct {
	MetaDataProperty     []MetaDataPropertyType        `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType              `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType        `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                    `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                        `xml:"id,attr,omitempty"`
	BoundedBy            *BoundingShapeType            `xml:"http://www.opengis.net/gml/3.2 boundedBy,omitempty"`
	Location             *LocationPropertyType         `xml:"http://www.opengis.net/gml/3.2 location,omitempty"`
	PriorityLocation     *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml/3.2 priorityLocation,omitempty"`
	ValidTime            *TimePrimitivePropertyType    `xml:"http://www.opengis.net/gml/3.2 validTime,omitempty"`
	Using                *ProcedurePropertyType        `xml:"http://www.opengis.net/gml/3.2 using,omitempty"`
	Target               *TargetPropertyType           `xml:"http://www.opengis.net/gml/3.2 target,omitempty"`
	Subject              *TargetPropertyType           `xml:"http://www.opengis.net/gml/3.2 subject,omitempty"`
	ResultOf             *ResultType                   `xml:"http://www.opengis.net/gml/3.2 resultOf,omitempty"`
	Direction            *DirectionPropertyType        `xml:"http://www.opengis.net/gml/3.2 direction,omitempty"`
}

type DirectedTopoSolidPropertyType struct {
	TopoSolid    *TopoSolidType `xml:"http://www.opengis.net/gml/3.2 TopoSolid,omitempty"`
	Orientation  string         `xml:"orientation,attr,omitempty"`
	NilReason    string         `xml:"nilReason,attr,omitempty"`
	RemoteSchema string         `xml:"remoteSchema,attr,omitempty"`
	Owns         bool           `xml:"owns,attr,omitempty"`
}

type DirectionDescriptionType struct {
	CompassPoint string         `xml:"http://www.opengis.net/gml/3.2 compassPoint,omitempty"`
	Keyword      *CodeType      `xml:"http://www.opengis.net/gml/3.2 keyword,omitempty"`
	Description  string         `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	Reference    *ReferenceType `xml:"http://www.opengis.net/gml/3.2 reference,omitempty"`
}

type DirectionPropertyType struct {
	DirectionVector      *DirectionVectorType      `xml:"http://www.opengis.net/gml/3.2 DirectionVector,omitempty"`
	DirectionDescription *DirectionDescriptionType `xml:"http://www.opengis.net/gml/3.2 DirectionDescription,omitempty"`
	CompassPoint         string                    `xml:"http://www.opengis.net/gml/3.2 CompassPoint,omitempty"`
	DirectionKeyword     *CodeType                 `xml:"http://www.opengis.net/gml/3.2 DirectionKeyword,omitempty"`
	DirectionString      *StringOrRefType          `xml:"http://www.opengis.net/gml/3.2 DirectionString,omitempty"`
	Owns                 bool                      `xml:"owns,attr,omitempty"`
	NilReason            string                    `xml:"nilReason,attr,omitempty"`
	RemoteSchema         string                    `xml:"remoteSchema,attr,omitempty"`
}

type DirectionVectorType struct {
	Vector          *VectorType `xml:"http://www.opengis.net/gml/3.2 vector,omitempty"`
	HorizontalAngle *AngleType  `xml:"http://www.opengis.net/gml/3.2 horizontalAngle,omitempty"`
	VerticalAngle   *AngleType  `xml:"http://www.opengis.net/gml/3.2 verticalAngle,omitempty"`
}

type DiscreteCoverageType struct {
	MetaDataProperty     []MetaDataPropertyType        `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType              `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType        `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                    `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                        `xml:"id,attr,omitempty"`
	BoundedBy            *BoundingShapeType            `xml:"http://www.opengis.net/gml/3.2 boundedBy,omitempty"`
	Location             *LocationPropertyType         `xml:"http://www.opengis.net/gml/3.2 location,omitempty"`
	PriorityLocation     *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml/3.2 priorityLocation,omitempty"`
	DomainSet            *DomainSetType                `xml:"http://www.opengis.net/gml/3.2 domainSet,omitempty"`
	GridDomain           *DomainSetType                `xml:"http://www.opengis.net/gml/3.2 gridDomain,omitempty"`
	MultiCurveDomain     *DomainSetType                `xml:"http://www.opengis.net/gml/3.2 multiCurveDomain,omitempty"`
	MultiPointDomain     *DomainSetType                `xml:"http://www.opengis.net/gml/3.2 multiPointDomain,omitempty"`
	MultiSolidDomain     *DomainSetType                `xml:"http://www.opengis.net/gml/3.2 multiSolidDomain,omitempty"`
	MultiSurfaceDomain   *DomainSetType                `xml:"http://www.opengis.net/gml/3.2 multiSurfaceDomain,omitempty"`
	RectifiedGridDomain  *DomainSetType                `xml:"http://www.opengis.net/gml/3.2 rectifiedGridDomain,omitempty"`
	RangeSet             *RangeSetType                 `xml:"http://www.opengis.net/gml/3.2 rangeSet,omitempty"`
	CoverageFunction     *CoverageFunctionType         `xml:"http://www.opengis.net/gml/3.2 coverageFunction,omitempty"`
}

type DomainSetType struct {
	AbstractGeometry           *AbstractGeometryType           `xml:"http://www.opengis.net/gml/3.2 AbstractGeometry,omitempty"`
	AbstractGeometricAggregate *AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml/3.2 AbstractGeometricAggregate,omitempty"`
	AbstractGeometricPrimitive *AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml/3.2 AbstractGeometricPrimitive,omitempty"`
	AbstractImplicitGeometry   *AbstractGeometryType           `xml:"http://www.opengis.net/gml/3.2 AbstractImplicitGeometry,omitempty"`
	GeometricComplex           *GeometricComplexType           `xml:"http://www.opengis.net/gml/3.2 GeometricComplex,omitempty"`
	AbstractTimeObject         *AbstractTimeObjectType         `xml:"http://www.opengis.net/gml/3.2 AbstractTimeObject,omitempty"`
	AbstractTimeComplex        *AbstractTimeComplexType        `xml:"http://www.opengis.net/gml/3.2 AbstractTimeComplex,omitempty"`
	AbstractTimePrimitive      *AbstractTimePrimitiveType      `xml:"http://www.opengis.net/gml/3.2 AbstractTimePrimitive,omitempty"`
	Owns                       bool                            `xml:"owns,attr,omitempty"`
	NilReason                  string                          `xml:"nilReason,attr,omitempty"`
	RemoteSchema               string                          `xml:"remoteSchema,attr,omitempty"`
}

type DynamicFeatureCollectionType struct {
	MetaDataProperty     []MetaDataPropertyType        `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType              `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType        `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                    `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                        `xml:"id,attr,omitempty"`
	BoundedBy            *BoundingShapeType            `xml:"http://www.opengis.net/gml/3.2 boundedBy,omitempty"`
	Location             *LocationPropertyType         `xml:"http://www.opengis.net/gml/3.2 location,omitempty"`
	PriorityLocation     *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml/3.2 priorityLocation,omitempty"`
	DynamicMembers       *DynamicFeatureMemberType     `xml:"http://www.opengis.net/gml/3.2 dynamicMembers,omitempty"`
}

type DynamicFeatureMemberType struct {
	Owns                     bool                          `xml:"owns,attr,omitempty"`
	DynamicFeature           []DynamicFeatureType          `xml:"http://www.opengis.net/gml/3.2 DynamicFeature"`
	DynamicFeatureCollection *DynamicFeatureCollectionType `xml:"http://www.opengis.net/gml/3.2 DynamicFeatureCollection,omitempty"`
	NilReason                string                        `xml:"nilReason,attr,omitempty"`
	RemoteSchema             string                        `xml:"remoteSchema,attr,omitempty"`
}

type DynamicFeatureType struct {
	MetaDataProperty     []MetaDataPropertyType        `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType              `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType        `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                    `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                        `xml:"id,attr,omitempty"`
	BoundedBy            *BoundingShapeType            `xml:"http://www.opengis.net/gml/3.2 boundedBy,omitempty"`
	Location             *LocationPropertyType         `xml:"http://www.opengis.net/gml/3.2 location,omitempty"`
	PriorityLocation     *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml/3.2 priorityLocation,omitempty"`
}

type EdgeType struct {
	MetaDataProperty     []MetaDataPropertyType     `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType           `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType             `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType     `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                 `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                     `xml:"id,attr,omitempty"`
	Container            *TopoSolidPropertyType     `xml:"http://www.opengis.net/gml/3.2 container,omitempty"`
	DirectedNode         []DirectedNodePropertyType `xml:"http://www.opengis.net/gml/3.2 directedNode"`
	DirectedFace         []DirectedFacePropertyType `xml:"http://www.opengis.net/gml/3.2 directedFace"`
	CurveProperty        *CurvePropertyType         `xml:"http://www.opengis.net/gml/3.2 curveProperty,omitempty"`
	AggregationType      string                     `xml:"aggregationType,attr,omitempty"`
}

type EllipsoidPropertyType struct {
	Ellipsoid    *EllipsoidType `xml:"http://www.opengis.net/gml/3.2 Ellipsoid,omitempty"`
	NilReason    string         `xml:"nilReason,attr,omitempty"`
	RemoteSchema string         `xml:"remoteSchema,attr,omitempty"`
}

type EllipsoidType struct {
	MetaDataProperty        []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description             *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference    *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier              *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                    []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                      string                 `xml:"id,attr,omitempty"`
	Remarks                 *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	SemiMajorAxis           *MeasureType           `xml:"http://www.opengis.net/gml/3.2 semiMajorAxis,omitempty"`
	SecondDefiningParameter *string                `xml:"http://www.opengis.net/gml/3.2 secondDefiningParameter,omitempty"`
}

type EllipsoidalCSPropertyType struct {
	EllipsoidalCS *EllipsoidalCSType `xml:"http://www.opengis.net/gml/3.2 EllipsoidalCS,omitempty"`
	NilReason     string             `xml:"nilReason,attr,omitempty"`
	RemoteSchema  string             `xml:"remoteSchema,attr,omitempty"`
}

type EllipsoidalCSType struct {
	MetaDataProperty     []MetaDataPropertyType             `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType                   `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                     `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType             `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                         `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                             `xml:"id,attr,omitempty"`
	Remarks              *string                            `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	Axis                 []CoordinateSystemAxisPropertyType `xml:"http://www.opengis.net/gml/3.2 axis"`
	UsesAxis             *CoordinateSystemAxisPropertyType  `xml:"http://www.opengis.net/gml/3.2 usesAxis,omitempty"`
	AggregationType      string                             `xml:"aggregationType,attr,omitempty"`
}

type EngineeringCRSPropertyType struct {
	EngineeringCRS *EngineeringCRSType `xml:"http://www.opengis.net/gml/3.2 EngineeringCRS,omitempty"`
	NilReason      string              `xml:"nilReason,attr,omitempty"`
	RemoteSchema   string              `xml:"remoteSchema,attr,omitempty"`
}

type EngineeringCRSType struct {
	MetaDataProperty     []MetaDataPropertyType        `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType              `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType        `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                    `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                        `xml:"id,attr,omitempty"`
	Remarks              *string                       `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     []string                      `xml:"http://www.opengis.net/gml/3.2 domainOfValidity"`
	Scope                []string                      `xml:"http://www.opengis.net/gml/3.2 scope"`
	EngineeringDatum     *EngineeringDatumPropertyType `xml:"http://www.opengis.net/gml/3.2 engineeringDatum,omitempty"`
	UsesEngineeringDatum *EngineeringDatumPropertyType `xml:"http://www.opengis.net/gml/3.2 usesEngineeringDatum,omitempty"`
	AffineCS             *AffineCSPropertyType         `xml:"http://www.opengis.net/gml/3.2 affineCS,omitempty"`
	UsesAffineCS         *AffineCSPropertyType         `xml:"http://www.opengis.net/gml/3.2 usesAffineCS,omitempty"`
	CartesianCS          *CartesianCSPropertyType      `xml:"http://www.opengis.net/gml/3.2 cartesianCS,omitempty"`
	UsesCartesianCS      *CartesianCSPropertyType      `xml:"http://www.opengis.net/gml/3.2 usesCartesianCS,omitempty"`
	CylindricalCS        *CylindricalCSPropertyType    `xml:"http://www.opengis.net/gml/3.2 cylindricalCS,omitempty"`
	LinearCS             *LinearCSPropertyType         `xml:"http://www.opengis.net/gml/3.2 linearCS,omitempty"`
	PolarCS              *PolarCSPropertyType          `xml:"http://www.opengis.net/gml/3.2 polarCS,omitempty"`
	SphericalCS          *SphericalCSPropertyType      `xml:"http://www.opengis.net/gml/3.2 sphericalCS,omitempty"`
	UsesSphericalCS      *SphericalCSPropertyType      `xml:"http://www.opengis.net/gml/3.2 usesSphericalCS,omitempty"`
	UserDefinedCS        *UserDefinedCSPropertyType    `xml:"http://www.opengis.net/gml/3.2 userDefinedCS,omitempty"`
	CoordinateSystem     *CoordinateSystemPropertyType `xml:"http://www.opengis.net/gml/3.2 coordinateSystem,omitempty"`
	UsesCS               *CoordinateSystemPropertyType `xml:"http://www.opengis.net/gml/3.2 usesCS,omitempty"`
}

type EngineeringDatumPropertyType struct {
	EngineeringDatum *EngineeringDatumType `xml:"http://www.opengis.net/gml/3.2 EngineeringDatum,omitempty"`
	NilReason        string                `xml:"nilReason,attr,omitempty"`
	RemoteSchema     string                `xml:"remoteSchema,attr,omitempty"`
}

type EngineeringDatumType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	Remarks              *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     *string                `xml:"http://www.opengis.net/gml/3.2 domainOfValidity,omitempty"`
	Scope                []string               `xml:"http://www.opengis.net/gml/3.2 scope"`
	AnchorDefinition     *CodeType              `xml:"http://www.opengis.net/gml/3.2 anchorDefinition,omitempty"`
	AnchorPoint          *CodeType              `xml:"http://www.opengis.net/gml/3.2 anchorPoint,omitempty"`
	RealizationEpoch     *string                `xml:"http://www.opengis.net/gml/3.2 realizationEpoch,omitempty"`
}

type EnvelopeType struct {
	Pos          []DirectPositionType `xml:"http://www.opengis.net/gml/3.2 pos"`
	Coordinates  *CoordinatesType     `xml:"http://www.opengis.net/gml/3.2 coordinates,omitempty"`
	LowerCorner  *DirectPositionType  `xml:"http://www.opengis.net/gml/3.2 lowerCorner,omitempty"`
	UpperCorner  *DirectPositionType  `xml:"http://www.opengis.net/gml/3.2 upperCorner,omitempty"`
	SrsName      string               `xml:"srsName,attr,omitempty"`
	SrsDimension int                  `xml:"srsDimension,attr,omitempty"`
	AxisLabels   string               `xml:"axisLabels,attr,omitempty"`
	UomLabels    string               `xml:"uomLabels,attr,omitempty"`
}

type EnvelopeWithTimePeriodType struct {
	Pos           []DirectPositionType `xml:"http://www.opengis.net/gml/3.2 pos"`
	Coordinates   *CoordinatesType     `xml:"http://www.opengis.net/gml/3.2 coordinates,omitempty"`
	LowerCorner   *DirectPositionType  `xml:"http://www.opengis.net/gml/3.2 lowerCorner,omitempty"`
	UpperCorner   *DirectPositionType  `xml:"http://www.opengis.net/gml/3.2 upperCorner,omitempty"`
	SrsName       string               `xml:"srsName,attr,omitempty"`
	SrsDimension  int                  `xml:"srsDimension,attr,omitempty"`
	AxisLabels    string               `xml:"axisLabels,attr,omitempty"`
	UomLabels     string               `xml:"uomLabels,attr,omitempty"`
	BeginPosition *TimePositionType    `xml:"http://www.opengis.net/gml/3.2 beginPosition,omitempty"`
	EndPosition   *TimePositionType    `xml:"http://www.opengis.net/gml/3.2 endPosition,omitempty"`
	Frame         string               `xml:"frame,attr,omitempty"`
}

type FaceOrTopoSolidPropertyType struct {
	Face         *FaceType      `xml:"http://www.opengis.net/gml/3.2 Face,omitempty"`
	TopoSolid    *TopoSolidType `xml:"http://www.opengis.net/gml/3.2 TopoSolid,omitempty"`
	NilReason    string         `xml:"nilReason,attr,omitempty"`
	RemoteSchema string         `xml:"remoteSchema,attr,omitempty"`
	Owns         bool           `xml:"owns,attr,omitempty"`
}

type FaceType struct {
	MetaDataProperty     []MetaDataPropertyType          `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType                `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                  `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType          `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                      `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                          `xml:"id,attr,omitempty"`
	Isolated             []NodePropertyType              `xml:"http://www.opengis.net/gml/3.2 isolated"`
	DirectedEdge         []DirectedEdgePropertyType      `xml:"http://www.opengis.net/gml/3.2 directedEdge"`
	DirectedTopoSolid    []DirectedTopoSolidPropertyType `xml:"http://www.opengis.net/gml/3.2 directedTopoSolid"`
	SurfaceProperty      *SurfacePropertyType            `xml:"http://www.opengis.net/gml/3.2 surfaceProperty,omitempty"`
	Universal            bool                            `xml:"universal,attr,omitempty"`
	AggregationType      string                          `xml:"aggregationType,attr,omitempty"`
}

type FeatureArrayPropertyType struct {
	AbstractFeature            *AbstractFeatureType            `xml:"http://www.opengis.net/gml/3.2 AbstractFeature,omitempty"`
	AbstractContinuousCoverage *AbstractContinuousCoverageType `xml:"http://www.opengis.net/gml/3.2 AbstractContinuousCoverage,omitempty"`
	AbstractCoverage           *AbstractCoverageType           `xml:"http://www.opengis.net/gml/3.2 AbstractCoverage,omitempty"`
	AbstractFeatureCollection  *AbstractFeatureCollectionType  `xml:"http://www.opengis.net/gml/3.2 AbstractFeatureCollection,omitempty"`
	DynamicFeature             *DynamicFeatureType             `xml:"http://www.opengis.net/gml/3.2 DynamicFeature,omitempty"`
	FeatureCollection          *FeatureCollectionType          `xml:"http://www.opengis.net/gml/3.2 FeatureCollection,omitempty"`
	Observation                *ObservationType                `xml:"http://www.opengis.net/gml/3.2 Observation,omitempty"`
}

type FeatureCollectionType struct {
	MetaDataProperty     []MetaDataPropertyType        `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType              `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType        `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                    `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                        `xml:"id,attr,omitempty"`
	BoundedBy            *BoundingShapeType            `xml:"http://www.opengis.net/gml/3.2 boundedBy,omitempty"`
	Location             *LocationPropertyType         `xml:"http://www.opengis.net/gml/3.2 location,omitempty"`
	PriorityLocation     *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml/3.2 priorityLocation,omitempty"`
	FeatureMember        []FeaturePropertyType         `xml:"http://www.opengis.net/gml/3.2 featureMember"`
	FeatureMembers       *FeatureArrayPropertyType     `xml:"http://www.opengis.net/gml/3.2 featureMembers,omitempty"`
}

type FeaturePropertyType struct {
	AbstractFeature            *AbstractFeatureType            `xml:"http://www.opengis.net/gml/3.2 AbstractFeature,omitempty"`
	AbstractContinuousCoverage *AbstractContinuousCoverageType `xml:"http://www.opengis.net/gml/3.2 AbstractContinuousCoverage,omitempty"`
	AbstractCoverage           *AbstractCoverageType           `xml:"http://www.opengis.net/gml/3.2 AbstractCoverage,omitempty"`
	AbstractFeatureCollection  *AbstractFeatureCollectionType  `xml:"http://www.opengis.net/gml/3.2 AbstractFeatureCollection,omitempty"`
	DynamicFeature             *DynamicFeatureType             `xml:"http://www.opengis.net/gml/3.2 DynamicFeature,omitempty"`
	FeatureCollection          *FeatureCollectionType          `xml:"http://www.opengis.net/gml/3.2 FeatureCollection,omitempty"`
	Observation                *ObservationType                `xml:"http://www.opengis.net/gml/3.2 Observation,omitempty"`
	Owns                       bool                            `xml:"owns,attr,omitempty"`
	NilReason                  string                          `xml:"nilReason,attr,omitempty"`
	RemoteSchema               string                          `xml:"remoteSchema,attr,omitempty"`
}

type FileType struct {
	RangeParameters *AssociationRoleType `xml:"http://www.opengis.net/gml/3.2 rangeParameters,omitempty"`
	FileStructure   *CodeType            `xml:"http://www.opengis.net/gml/3.2 fileStructure,omitempty"`
	MimeType        string               `xml:"http://www.opengis.net/gml/3.2 mimeType,omitempty"`
	Compression     string               `xml:"http://www.opengis.net/gml/3.2 compression,omitempty"`
	FileName        string               `xml:"http://www.opengis.net/gml/3.2 fileName,omitempty"`
	FileReference   string               `xml:"http://www.opengis.net/gml/3.2 fileReference,omitempty"`
}

type FormulaType struct {
	A float64 `xml:"http://www.opengis.net/gml/3.2 a,omitempty"`
	B float64 `xml:"http://www.opengis.net/gml/3.2 b,omitempty"`
	C float64 `xml:"http://www.opengis.net/gml/3.2 c,omitempty"`
	D float64 `xml:"http://www.opengis.net/gml/3.2 d,omitempty"`
}

type GeneralConversionPropertyType struct {
	AbstractGeneralConversion *AbstractGeneralConversionType `xml:"http://www.opengis.net/gml/3.2 AbstractGeneralConversion,omitempty"`
	Conversion                *ConversionType                `xml:"http://www.opengis.net/gml/3.2 Conversion,omitempty"`
	NilReason                 string                         `xml:"nilReason,attr,omitempty"`
	RemoteSchema              string                         `xml:"remoteSchema,attr,omitempty"`
}

type GeneralTransformationPropertyType struct {
	AbstractGeneralTransformation *AbstractGeneralTransformationType `xml:"http://www.opengis.net/gml/3.2 AbstractGeneralTransformation,omitempty"`
	Transformation                *TransformationType                `xml:"http://www.opengis.net/gml/3.2 Transformation,omitempty"`
	NilReason                     string                             `xml:"nilReason,attr,omitempty"`
	RemoteSchema                  string                             `xml:"remoteSchema,attr,omitempty"`
}

type GenericMetaDataType struct {
	Id string `xml:"id,attr,omitempty"`
}

type GeocentricCRSPropertyType struct {
	GeocentricCRS *GeocentricCRSType `xml:"http://www.opengis.net/gml/3.2 GeocentricCRS,omitempty"`
	NilReason     string             `xml:"nilReason,attr,omitempty"`
	RemoteSchema  string             `xml:"remoteSchema,attr,omitempty"`
}

type GeocentricCRSType struct {
	MetaDataProperty     []MetaDataPropertyType     `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType           `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType             `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType     `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                 `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                     `xml:"id,attr,omitempty"`
	Remarks              *string                    `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     []string                   `xml:"http://www.opengis.net/gml/3.2 domainOfValidity"`
	Scope                []string                   `xml:"http://www.opengis.net/gml/3.2 scope"`
	UsesGeodeticDatum    *GeodeticDatumPropertyType `xml:"http://www.opengis.net/gml/3.2 usesGeodeticDatum,omitempty"`
	UsesCartesianCS      *CartesianCSPropertyType   `xml:"http://www.opengis.net/gml/3.2 usesCartesianCS,omitempty"`
	UsesSphericalCS      *SphericalCSPropertyType   `xml:"http://www.opengis.net/gml/3.2 usesSphericalCS,omitempty"`
}

type GeodesicStringType struct {
	NumDerivativesAtStart int                     `xml:"numDerivativesAtStart,attr,omitempty"`
	NumDerivativesAtEnd   int                     `xml:"numDerivativesAtEnd,attr,omitempty"`
	NumDerivativeInterior int                     `xml:"numDerivativeInterior,attr,omitempty"`
	PosList               *DirectPositionListType `xml:"http://www.opengis.net/gml/3.2 posList,omitempty"`
	Pos                   *DirectPositionType     `xml:"http://www.opengis.net/gml/3.2 pos,omitempty"`
	PointProperty         *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointProperty,omitempty"`
	Interpolation         string                  `xml:"interpolation,attr,omitempty"`
}

type GeodesicType struct {
	NumDerivativesAtStart int                     `xml:"numDerivativesAtStart,attr,omitempty"`
	NumDerivativesAtEnd   int                     `xml:"numDerivativesAtEnd,attr,omitempty"`
	NumDerivativeInterior int                     `xml:"numDerivativeInterior,attr,omitempty"`
	PosList               *DirectPositionListType `xml:"http://www.opengis.net/gml/3.2 posList,omitempty"`
	Pos                   *DirectPositionType     `xml:"http://www.opengis.net/gml/3.2 pos,omitempty"`
	PointProperty         *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointProperty,omitempty"`
	Interpolation         string                  `xml:"interpolation,attr,omitempty"`
}

type GeodeticCRSPropertyType struct {
	GeodeticCRS  *GeodeticCRSType `xml:"http://www.opengis.net/gml/3.2 GeodeticCRS,omitempty"`
	NilReason    string           `xml:"nilReason,attr,omitempty"`
	RemoteSchema string           `xml:"remoteSchema,attr,omitempty"`
}

type GeodeticCRSType struct {
	MetaDataProperty     []MetaDataPropertyType     `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType           `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType             `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType     `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                 `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                     `xml:"id,attr,omitempty"`
	Remarks              *string                    `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     []string                   `xml:"http://www.opengis.net/gml/3.2 domainOfValidity"`
	Scope                []string                   `xml:"http://www.opengis.net/gml/3.2 scope"`
	GeodeticDatum        *GeodeticDatumPropertyType `xml:"http://www.opengis.net/gml/3.2 geodeticDatum,omitempty"`
	UsesGeodeticDatum    *GeodeticDatumPropertyType `xml:"http://www.opengis.net/gml/3.2 usesGeodeticDatum,omitempty"`
	EllipsoidalCS        *EllipsoidalCSPropertyType `xml:"http://www.opengis.net/gml/3.2 ellipsoidalCS,omitempty"`
	UsesEllipsoidalCS    *EllipsoidalCSPropertyType `xml:"http://www.opengis.net/gml/3.2 usesEllipsoidalCS,omitempty"`
	CartesianCS          *CartesianCSPropertyType   `xml:"http://www.opengis.net/gml/3.2 cartesianCS,omitempty"`
	UsesCartesianCS      *CartesianCSPropertyType   `xml:"http://www.opengis.net/gml/3.2 usesCartesianCS,omitempty"`
	SphericalCS          *SphericalCSPropertyType   `xml:"http://www.opengis.net/gml/3.2 sphericalCS,omitempty"`
	UsesSphericalCS      *SphericalCSPropertyType   `xml:"http://www.opengis.net/gml/3.2 usesSphericalCS,omitempty"`
}

type GeodeticDatumPropertyType struct {
	GeodeticDatum *GeodeticDatumType `xml:"http://www.opengis.net/gml/3.2 GeodeticDatum,omitempty"`
	NilReason     string             `xml:"nilReason,attr,omitempty"`
	RemoteSchema  string             `xml:"remoteSchema,attr,omitempty"`
}

type GeodeticDatumType struct {
	MetaDataProperty     []MetaDataPropertyType     `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType           `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType             `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType     `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                 `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                     `xml:"id,attr,omitempty"`
	Remarks              *string                    `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     *string                    `xml:"http://www.opengis.net/gml/3.2 domainOfValidity,omitempty"`
	Scope                []string                   `xml:"http://www.opengis.net/gml/3.2 scope"`
	AnchorDefinition     *CodeType                  `xml:"http://www.opengis.net/gml/3.2 anchorDefinition,omitempty"`
	AnchorPoint          *CodeType                  `xml:"http://www.opengis.net/gml/3.2 anchorPoint,omitempty"`
	RealizationEpoch     *string                    `xml:"http://www.opengis.net/gml/3.2 realizationEpoch,omitempty"`
	PrimeMeridian        *PrimeMeridianPropertyType `xml:"http://www.opengis.net/gml/3.2 primeMeridian,omitempty"`
	UsesPrimeMeridian    *PrimeMeridianPropertyType `xml:"http://www.opengis.net/gml/3.2 usesPrimeMeridian,omitempty"`
	Ellipsoid            *EllipsoidPropertyType     `xml:"http://www.opengis.net/gml/3.2 ellipsoid,omitempty"`
	UsesEllipsoid        *EllipsoidPropertyType     `xml:"http://www.opengis.net/gml/3.2 usesEllipsoid,omitempty"`
}

type GeographicCRSPropertyType struct {
	GeographicCRS *GeographicCRSType `xml:"http://www.opengis.net/gml/3.2 GeographicCRS,omitempty"`
	NilReason     string             `xml:"nilReason,attr,omitempty"`
	RemoteSchema  string             `xml:"remoteSchema,attr,omitempty"`
}

type GeographicCRSType struct {
	MetaDataProperty     []MetaDataPropertyType     `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType           `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType             `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType     `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                 `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                     `xml:"id,attr,omitempty"`
	Remarks              *string                    `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     []string                   `xml:"http://www.opengis.net/gml/3.2 domainOfValidity"`
	Scope                []string                   `xml:"http://www.opengis.net/gml/3.2 scope"`
	UsesEllipsoidalCS    *EllipsoidalCSPropertyType `xml:"http://www.opengis.net/gml/3.2 usesEllipsoidalCS,omitempty"`
	UsesGeodeticDatum    *GeodeticDatumPropertyType `xml:"http://www.opengis.net/gml/3.2 usesGeodeticDatum,omitempty"`
}

type GeometricComplexPropertyType struct {
	GeometricComplex *GeometricComplexType `xml:"http://www.opengis.net/gml/3.2 GeometricComplex,omitempty"`
	CompositeCurve   *CompositeCurveType   `xml:"http://www.opengis.net/gml/3.2 CompositeCurve,omitempty"`
	CompositeSurface *CompositeSurfaceType `xml:"http://www.opengis.net/gml/3.2 CompositeSurface,omitempty"`
	CompositeSolid   *CompositeSolidType   `xml:"http://www.opengis.net/gml/3.2 CompositeSolid,omitempty"`
	Owns             bool                  `xml:"owns,attr,omitempty"`
	NilReason        string                `xml:"nilReason,attr,omitempty"`
	RemoteSchema     string                `xml:"remoteSchema,attr,omitempty"`
}

type GeometricComplexType struct {
	MetaDataProperty     []MetaDataPropertyType           `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType                 `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                   `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType           `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                       `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                           `xml:"id,attr,omitempty"`
	SrsName              string                           `xml:"srsName,attr,omitempty"`
	SrsDimension         int                              `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                           `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                           `xml:"uomLabels,attr,omitempty"`
	Element              []GeometricPrimitivePropertyType `xml:"http://www.opengis.net/gml/3.2 element"`
	AggregationType      string                           `xml:"aggregationType,attr,omitempty"`
}

type GeometricPrimitivePropertyType struct {
	AbstractGeometricPrimitive *AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml/3.2 AbstractGeometricPrimitive,omitempty"`
	AbstractCurve              *AbstractCurveType              `xml:"http://www.opengis.net/gml/3.2 AbstractCurve,omitempty"`
	AbstractSolid              *AbstractSolidType              `xml:"http://www.opengis.net/gml/3.2 AbstractSolid,omitempty"`
	AbstractSurface            *AbstractSurfaceType            `xml:"http://www.opengis.net/gml/3.2 AbstractSurface,omitempty"`
	Point                      *PointType                      `xml:"http://www.opengis.net/gml/3.2 Point,omitempty"`
	Owns                       bool                            `xml:"owns,attr,omitempty"`
	NilReason                  string                          `xml:"nilReason,attr,omitempty"`
	RemoteSchema               string                          `xml:"remoteSchema,attr,omitempty"`
}

type GeometryArrayPropertyType struct {
	AbstractGeometry           *AbstractGeometryType           `xml:"http://www.opengis.net/gml/3.2 AbstractGeometry,omitempty"`
	AbstractGeometricAggregate *AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml/3.2 AbstractGeometricAggregate,omitempty"`
	AbstractGeometricPrimitive *AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml/3.2 AbstractGeometricPrimitive,omitempty"`
	AbstractImplicitGeometry   *AbstractGeometryType           `xml:"http://www.opengis.net/gml/3.2 AbstractImplicitGeometry,omitempty"`
	GeometricComplex           *GeometricComplexType           `xml:"http://www.opengis.net/gml/3.2 GeometricComplex,omitempty"`
	Owns                       bool                            `xml:"owns,attr,omitempty"`
}

type GeometryPropertyType struct {
	AbstractGeometry           *AbstractGeometryType           `xml:"http://www.opengis.net/gml/3.2 AbstractGeometry,omitempty"`
	AbstractGeometricAggregate *AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml/3.2 AbstractGeometricAggregate,omitempty"`
	AbstractGeometricPrimitive *AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml/3.2 AbstractGeometricPrimitive,omitempty"`
	AbstractImplicitGeometry   *AbstractGeometryType           `xml:"http://www.opengis.net/gml/3.2 AbstractImplicitGeometry,omitempty"`
	GeometricComplex           *GeometricComplexType           `xml:"http://www.opengis.net/gml/3.2 GeometricComplex,omitempty"`
	NilReason                  string                          `xml:"nilReason,attr,omitempty"`
	RemoteSchema               string                          `xml:"remoteSchema,attr,omitempty"`
	Owns                       bool                            `xml:"owns,attr,omitempty"`
}

type GridEnvelopeType struct {
	Low  string `xml:"http://www.opengis.net/gml/3.2 low,omitempty"`
	High string `xml:"http://www.opengis.net/gml/3.2 high,omitempty"`
}

type GridFunctionType struct {
	SequenceRule *SequenceRuleType `xml:"http://www.opengis.net/gml/3.2 sequenceRule,omitempty"`
	StartPoint   string            `xml:"http://www.opengis.net/gml/3.2 startPoint,omitempty"`
}

type GridLengthType struct {
	Value string `xml:",chardata"`
}

type GridLimitsType struct {
	GridEnvelope *GridEnvelopeType `xml:"http://www.opengis.net/gml/3.2 GridEnvelope,omitempty"`
}

type GridType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	SrsName              string                 `xml:"srsName,attr,omitempty"`
	SrsDimension         int                    `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                 `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                 `xml:"uomLabels,attr,omitempty"`
	Limits               *GridLimitsType        `xml:"http://www.opengis.net/gml/3.2 limits,omitempty"`
	AxisName             []string               `xml:"http://www.opengis.net/gml/3.2 axisName"`
	Dimension            int                    `xml:"dimension,attr"`
}

type HistoryPropertyType struct {
	AbstractTimeSlice  []AbstractTimeSliceType `xml:"http://www.opengis.net/gml/3.2 AbstractTimeSlice"`
	MovingObjectStatus *MovingObjectStatusType `xml:"http://www.opengis.net/gml/3.2 MovingObjectStatus,omitempty"`
	Owns               bool                    `xml:"owns,attr,omitempty"`
}

type IdentifiedObjectType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	Remarks              *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
}

type ImageCRSPropertyType struct {
	ImageCRS     *ImageCRSType `xml:"http://www.opengis.net/gml/3.2 ImageCRS,omitempty"`
	NilReason    string        `xml:"nilReason,attr,omitempty"`
	RemoteSchema string        `xml:"remoteSchema,attr,omitempty"`
}

type ImageCRSType struct {
	MetaDataProperty       []MetaDataPropertyType          `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description            *StringOrRefType                `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference   *ReferenceType                  `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier             *CodeWithAuthorityType          `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                   []CodeType                      `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                     string                          `xml:"id,attr,omitempty"`
	Remarks                *string                         `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity       []string                        `xml:"http://www.opengis.net/gml/3.2 domainOfValidity"`
	Scope                  []string                        `xml:"http://www.opengis.net/gml/3.2 scope"`
	ImageDatum             *ImageDatumPropertyType         `xml:"http://www.opengis.net/gml/3.2 imageDatum,omitempty"`
	UsesImageDatum         *ImageDatumPropertyType         `xml:"http://www.opengis.net/gml/3.2 usesImageDatum,omitempty"`
	CartesianCS            *CartesianCSPropertyType        `xml:"http://www.opengis.net/gml/3.2 cartesianCS,omitempty"`
	UsesCartesianCS        *CartesianCSPropertyType        `xml:"http://www.opengis.net/gml/3.2 usesCartesianCS,omitempty"`
	AffineCS               *AffineCSPropertyType           `xml:"http://www.opengis.net/gml/3.2 affineCS,omitempty"`
	UsesAffineCS           *AffineCSPropertyType           `xml:"http://www.opengis.net/gml/3.2 usesAffineCS,omitempty"`
	UsesObliqueCartesianCS *ObliqueCartesianCSPropertyType `xml:"http://www.opengis.net/gml/3.2 usesObliqueCartesianCS,omitempty"`
}

type ImageDatumPropertyType struct {
	ImageDatum   *ImageDatumType `xml:"http://www.opengis.net/gml/3.2 ImageDatum,omitempty"`
	NilReason    string          `xml:"nilReason,attr,omitempty"`
	RemoteSchema string          `xml:"remoteSchema,attr,omitempty"`
}

type ImageDatumType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	Remarks              *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     *string                `xml:"http://www.opengis.net/gml/3.2 domainOfValidity,omitempty"`
	Scope                []string               `xml:"http://www.opengis.net/gml/3.2 scope"`
	AnchorDefinition     *CodeType              `xml:"http://www.opengis.net/gml/3.2 anchorDefinition,omitempty"`
	AnchorPoint          *CodeType              `xml:"http://www.opengis.net/gml/3.2 anchorPoint,omitempty"`
	RealizationEpoch     *string                `xml:"http://www.opengis.net/gml/3.2 realizationEpoch,omitempty"`
	PixelInCell          *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 pixelInCell,omitempty"`
}

type IndirectEntryType struct {
	DefinitionProxy *DefinitionProxyType `xml:"http://www.opengis.net/gml/3.2 DefinitionProxy,omitempty"`
}

type InlinePropertyType struct {
	Owns bool `xml:"owns,attr,omitempty"`
}

type KnotPropertyType struct {
	Knot *KnotType `xml:"http://www.opengis.net/gml/3.2 Knot,omitempty"`
}

type KnotType struct {
	Value        float64 `xml:"http://www.opengis.net/gml/3.2 value,omitempty"`
	Multiplicity int     `xml:"http://www.opengis.net/gml/3.2 multiplicity,omitempty"`
	Weight       float64 `xml:"http://www.opengis.net/gml/3.2 weight,omitempty"`
}

type LengthType struct {
	Value string `xml:",chardata"`
}

type LineStringSegmentArrayPropertyType struct {
	LineStringSegment *LineStringSegmentType `xml:"http://www.opengis.net/gml/3.2 LineStringSegment,omitempty"`
}

type LineStringSegmentType struct {
	NumDerivativesAtStart int                     `xml:"numDerivativesAtStart,attr,omitempty"`
	NumDerivativesAtEnd   int                     `xml:"numDerivativesAtEnd,attr,omitempty"`
	NumDerivativeInterior int                     `xml:"numDerivativeInterior,attr,omitempty"`
	PosList               *DirectPositionListType `xml:"http://www.opengis.net/gml/3.2 posList,omitempty"`
	Coordinates           *CoordinatesType        `xml:"http://www.opengis.net/gml/3.2 coordinates,omitempty"`
	Pos                   *DirectPositionType     `xml:"http://www.opengis.net/gml/3.2 pos,omitempty"`
	PointProperty         *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointProperty,omitempty"`
	PointRep              *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointRep,omitempty"`
	Interpolation         string                  `xml:"interpolation,attr,omitempty"`
}

type LineStringType struct {
	MetaDataProperty     []MetaDataPropertyType  `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType        `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType          `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType  `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType              `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                  `xml:"id,attr,omitempty"`
	SrsName              string                  `xml:"srsName,attr,omitempty"`
	SrsDimension         int                     `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                  `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                  `xml:"uomLabels,attr,omitempty"`
	PosList              *DirectPositionListType `xml:"http://www.opengis.net/gml/3.2 posList,omitempty"`
	Coordinates          *CoordinatesType        `xml:"http://www.opengis.net/gml/3.2 coordinates,omitempty"`
	Pos                  *DirectPositionType     `xml:"http://www.opengis.net/gml/3.2 pos,omitempty"`
	PointProperty        *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointProperty,omitempty"`
	PointRep             *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointRep,omitempty"`
}

type LinearCSPropertyType struct {
	LinearCS     *LinearCSType `xml:"http://www.opengis.net/gml/3.2 LinearCS,omitempty"`
	NilReason    string        `xml:"nilReason,attr,omitempty"`
	RemoteSchema string        `xml:"remoteSchema,attr,omitempty"`
}

type LinearCSType struct {
	MetaDataProperty     []MetaDataPropertyType             `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType                   `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                     `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType             `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                         `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                             `xml:"id,attr,omitempty"`
	Remarks              *string                            `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	Axis                 []CoordinateSystemAxisPropertyType `xml:"http://www.opengis.net/gml/3.2 axis"`
	UsesAxis             *CoordinateSystemAxisPropertyType  `xml:"http://www.opengis.net/gml/3.2 usesAxis,omitempty"`
	AggregationType      string                             `xml:"aggregationType,attr,omitempty"`
}

type LinearRingPropertyType struct {
	LinearRing *LinearRingType `xml:"http://www.opengis.net/gml/3.2 LinearRing,omitempty"`
}

type LinearRingType struct {
	MetaDataProperty     []MetaDataPropertyType  `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType        `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType          `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType  `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType              `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                  `xml:"id,attr,omitempty"`
	SrsName              string                  `xml:"srsName,attr,omitempty"`
	SrsDimension         int                     `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                  `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                  `xml:"uomLabels,attr,omitempty"`
	PosList              *DirectPositionListType `xml:"http://www.opengis.net/gml/3.2 posList,omitempty"`
	Coordinates          *CoordinatesType        `xml:"http://www.opengis.net/gml/3.2 coordinates,omitempty"`
	Pos                  *DirectPositionType     `xml:"http://www.opengis.net/gml/3.2 pos,omitempty"`
	PointProperty        *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointProperty,omitempty"`
	PointRep             *PointPropertyType      `xml:"http://www.opengis.net/gml/3.2 pointRep,omitempty"`
}

type LocationPropertyType struct {
	AbstractGeometry           *AbstractGeometryType           `xml:"http://www.opengis.net/gml/3.2 AbstractGeometry,omitempty"`
	AbstractGeometricAggregate *AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml/3.2 AbstractGeometricAggregate,omitempty"`
	AbstractGeometricPrimitive *AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml/3.2 AbstractGeometricPrimitive,omitempty"`
	AbstractImplicitGeometry   *AbstractGeometryType           `xml:"http://www.opengis.net/gml/3.2 AbstractImplicitGeometry,omitempty"`
	GeometricComplex           *GeometricComplexType           `xml:"http://www.opengis.net/gml/3.2 GeometricComplex,omitempty"`
	LocationKeyWord            *CodeType                       `xml:"http://www.opengis.net/gml/3.2 LocationKeyWord,omitempty"`
	LocationString             *StringOrRefType                `xml:"http://www.opengis.net/gml/3.2 LocationString,omitempty"`
	Null                       *string                         `xml:"http://www.opengis.net/gml/3.2 Null,omitempty"`
	NilReason                  string                          `xml:"nilReason,attr,omitempty"`
	RemoteSchema               string                          `xml:"remoteSchema,attr,omitempty"`
}

type MappingRuleType struct {
	RuleDefinition string         `xml:"http://www.opengis.net/gml/3.2 ruleDefinition,omitempty"`
	RuleReference  *ReferenceType `xml:"http://www.opengis.net/gml/3.2 ruleReference,omitempty"`
}

type MeasureListType struct {
	Value string `xml:",chardata"`
	Uom   string `xml:"uom,attr"`
}

type MeasureOrNilReasonListType struct {
	Value string `xml:",chardata"`
	Uom   string `xml:"uom,attr"`
}

type MeasureType struct {
	Value float64 `xml:",chardata"`
	Uom   string  `xml:"uom,attr"`
}

type MetaDataPropertyType struct {
	AbstractMetaData *AbstractMetaDataType `xml:"http://www.opengis.net/gml/3.2 AbstractMetaData,omitempty"`
	GenericMetaData  *GenericMetaDataType  `xml:"http://www.opengis.net/gml/3.2 GenericMetaData,omitempty"`
	About            string                `xml:"about,attr,omitempty"`
	NilReason        string                `xml:"nilReason,attr,omitempty"`
	RemoteSchema     string                `xml:"remoteSchema,attr,omitempty"`
}

type MovingObjectStatusType struct {
	MetaDataProperty     []MetaDataPropertyType        `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType              `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType        `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                    `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                        `xml:"id,attr,omitempty"`
	ValidTime            *TimePrimitivePropertyType    `xml:"http://www.opengis.net/gml/3.2 validTime,omitempty"`
	DataSource           *StringOrRefType              `xml:"http://www.opengis.net/gml/3.2 dataSource,omitempty"`
	Speed                *MeasureType                  `xml:"http://www.opengis.net/gml/3.2 speed,omitempty"`
	Bearing              *DirectionPropertyType        `xml:"http://www.opengis.net/gml/3.2 bearing,omitempty"`
	Acceleration         *MeasureType                  `xml:"http://www.opengis.net/gml/3.2 acceleration,omitempty"`
	Elevation            *MeasureType                  `xml:"http://www.opengis.net/gml/3.2 elevation,omitempty"`
	Status               *StringOrRefType              `xml:"http://www.opengis.net/gml/3.2 status,omitempty"`
	StatusReference      *ReferenceType                `xml:"http://www.opengis.net/gml/3.2 statusReference,omitempty"`
	Position             *GeometryPropertyType         `xml:"http://www.opengis.net/gml/3.2 position,omitempty"`
	Pos                  *DirectPositionType           `xml:"http://www.opengis.net/gml/3.2 pos,omitempty"`
	LocationName         *CodeType                     `xml:"http://www.opengis.net/gml/3.2 locationName,omitempty"`
	LocationReference    *ReferenceType                `xml:"http://www.opengis.net/gml/3.2 locationReference,omitempty"`
	Location             *LocationPropertyType         `xml:"http://www.opengis.net/gml/3.2 location,omitempty"`
	PriorityLocation     *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml/3.2 priorityLocation,omitempty"`
}

type MultiCurvePropertyType struct {
	MultiCurve   *MultiCurveType `xml:"http://www.opengis.net/gml/3.2 MultiCurve,omitempty"`
	NilReason    string          `xml:"nilReason,attr,omitempty"`
	RemoteSchema string          `xml:"remoteSchema,attr,omitempty"`
	Owns         bool            `xml:"owns,attr,omitempty"`
}

type MultiCurveType struct {
	MetaDataProperty     []MetaDataPropertyType  `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType        `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType          `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType  `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType              `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                  `xml:"id,attr,omitempty"`
	SrsName              string                  `xml:"srsName,attr,omitempty"`
	SrsDimension         int                     `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                  `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                  `xml:"uomLabels,attr,omitempty"`
	AggregationType      string                  `xml:"aggregationType,attr,omitempty"`
	CurveMember          []CurvePropertyType     `xml:"http://www.opengis.net/gml/3.2 curveMember"`
	CurveMembers         *CurveArrayPropertyType `xml:"http://www.opengis.net/gml/3.2 curveMembers,omitempty"`
}

type MultiGeometryPropertyType struct {
	AbstractGeometricAggregate *AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml/3.2 AbstractGeometricAggregate,omitempty"`
	MultiCurve                 *MultiCurveType                 `xml:"http://www.opengis.net/gml/3.2 MultiCurve,omitempty"`
	MultiGeometry              *MultiGeometryType              `xml:"http://www.opengis.net/gml/3.2 MultiGeometry,omitempty"`
	MultiPoint                 *MultiPointType                 `xml:"http://www.opengis.net/gml/3.2 MultiPoint,omitempty"`
	MultiSolid                 *MultiSolidType                 `xml:"http://www.opengis.net/gml/3.2 MultiSolid,omitempty"`
	MultiSurface               *MultiSurfaceType               `xml:"http://www.opengis.net/gml/3.2 MultiSurface,omitempty"`
	NilReason                  string                          `xml:"nilReason,attr,omitempty"`
	RemoteSchema               string                          `xml:"remoteSchema,attr,omitempty"`
	Owns                       bool                            `xml:"owns,attr,omitempty"`
}

type MultiGeometryType struct {
	MetaDataProperty     []MetaDataPropertyType     `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType           `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType             `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType     `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                 `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                     `xml:"id,attr,omitempty"`
	SrsName              string                     `xml:"srsName,attr,omitempty"`
	SrsDimension         int                        `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                     `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                     `xml:"uomLabels,attr,omitempty"`
	AggregationType      string                     `xml:"aggregationType,attr,omitempty"`
	GeometryMember       []GeometryPropertyType     `xml:"http://www.opengis.net/gml/3.2 geometryMember"`
	GeometryMembers      *GeometryArrayPropertyType `xml:"http://www.opengis.net/gml/3.2 geometryMembers,omitempty"`
}

type MultiPointPropertyType struct {
	MultiPoint   *MultiPointType `xml:"http://www.opengis.net/gml/3.2 MultiPoint,omitempty"`
	NilReason    string          `xml:"nilReason,attr,omitempty"`
	RemoteSchema string          `xml:"remoteSchema,attr,omitempty"`
	Owns         bool            `xml:"owns,attr,omitempty"`
}

type MultiPointType struct {
	MetaDataProperty     []MetaDataPropertyType  `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType        `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType          `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType  `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType              `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                  `xml:"id,attr,omitempty"`
	SrsName              string                  `xml:"srsName,attr,omitempty"`
	SrsDimension         int                     `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                  `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                  `xml:"uomLabels,attr,omitempty"`
	AggregationType      string                  `xml:"aggregationType,attr,omitempty"`
	PointMember          []PointPropertyType     `xml:"http://www.opengis.net/gml/3.2 pointMember"`
	PointMembers         *PointArrayPropertyType `xml:"http://www.opengis.net/gml/3.2 pointMembers,omitempty"`
}

type MultiSolidPropertyType struct {
	MultiSolid   *MultiSolidType `xml:"http://www.opengis.net/gml/3.2 MultiSolid,omitempty"`
	NilReason    string          `xml:"nilReason,attr,omitempty"`
	RemoteSchema string          `xml:"remoteSchema,attr,omitempty"`
	Owns         bool            `xml:"owns,attr,omitempty"`
}

type MultiSolidType struct {
	MetaDataProperty     []MetaDataPropertyType  `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType        `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType          `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType  `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType              `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                  `xml:"id,attr,omitempty"`
	SrsName              string                  `xml:"srsName,attr,omitempty"`
	SrsDimension         int                     `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                  `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                  `xml:"uomLabels,attr,omitempty"`
	AggregationType      string                  `xml:"aggregationType,attr,omitempty"`
	SolidMember          []SolidPropertyType     `xml:"http://www.opengis.net/gml/3.2 solidMember"`
	SolidMembers         *SolidArrayPropertyType `xml:"http://www.opengis.net/gml/3.2 solidMembers,omitempty"`
}

type MultiSurfacePropertyType struct {
	MultiSurface *MultiSurfaceType `xml:"http://www.opengis.net/gml/3.2 MultiSurface,omitempty"`
	NilReason    string            `xml:"nilReason,attr,omitempty"`
	RemoteSchema string            `xml:"remoteSchema,attr,omitempty"`
	Owns         bool              `xml:"owns,attr,omitempty"`
}

type MultiSurfaceType struct {
	MetaDataProperty     []MetaDataPropertyType    `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType          `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType            `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType    `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                    `xml:"id,attr,omitempty"`
	SrsName              string                    `xml:"srsName,attr,omitempty"`
	SrsDimension         int                       `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                    `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                    `xml:"uomLabels,attr,omitempty"`
	AggregationType      string                    `xml:"aggregationType,attr,omitempty"`
	SurfaceMember        []SurfacePropertyType     `xml:"http://www.opengis.net/gml/3.2 surfaceMember"`
	SurfaceMembers       *SurfaceArrayPropertyType `xml:"http://www.opengis.net/gml/3.2 surfaceMembers,omitempty"`
}

type NodeOrEdgePropertyType struct {
	Node         *NodeType `xml:"http://www.opengis.net/gml/3.2 Node,omitempty"`
	Edge         *EdgeType `xml:"http://www.opengis.net/gml/3.2 Edge,omitempty"`
	NilReason    string    `xml:"nilReason,attr,omitempty"`
	RemoteSchema string    `xml:"remoteSchema,attr,omitempty"`
	Owns         bool      `xml:"owns,attr,omitempty"`
}

type NodePropertyType struct {
	Node         *NodeType `xml:"http://www.opengis.net/gml/3.2 Node,omitempty"`
	NilReason    string    `xml:"nilReason,attr,omitempty"`
	RemoteSchema string    `xml:"remoteSchema,attr,omitempty"`
	Owns         bool      `xml:"owns,attr,omitempty"`
}

type NodeType struct {
	MetaDataProperty     []MetaDataPropertyType       `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType             `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType               `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType       `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                   `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                       `xml:"id,attr,omitempty"`
	Container            *FaceOrTopoSolidPropertyType `xml:"http://www.opengis.net/gml/3.2 container,omitempty"`
	DirectedEdge         []DirectedEdgePropertyType   `xml:"http://www.opengis.net/gml/3.2 directedEdge"`
	PointProperty        *PointPropertyType           `xml:"http://www.opengis.net/gml/3.2 pointProperty,omitempty"`
	AggregationType      string                       `xml:"aggregationType,attr,omitempty"`
}

type ObliqueCartesianCSPropertyType struct {
	ObliqueCartesianCS *ObliqueCartesianCSType `xml:"http://www.opengis.net/gml/3.2 ObliqueCartesianCS,omitempty"`
	NilReason          string                  `xml:"nilReason,attr,omitempty"`
	RemoteSchema       string                  `xml:"remoteSchema,attr,omitempty"`
}

type ObliqueCartesianCSType struct {
	MetaDataProperty     []MetaDataPropertyType             `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType                   `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                     `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType             `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                         `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                             `xml:"id,attr,omitempty"`
	Remarks              *string                            `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	Axis                 []CoordinateSystemAxisPropertyType `xml:"http://www.opengis.net/gml/3.2 axis"`
	UsesAxis             *CoordinateSystemAxisPropertyType  `xml:"http://www.opengis.net/gml/3.2 usesAxis,omitempty"`
	AggregationType      string                             `xml:"aggregationType,attr,omitempty"`
}

type ObservationType struct {
	MetaDataProperty     []MetaDataPropertyType        `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType              `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType        `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                    `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                        `xml:"id,attr,omitempty"`
	BoundedBy            *BoundingShapeType            `xml:"http://www.opengis.net/gml/3.2 boundedBy,omitempty"`
	Location             *LocationPropertyType         `xml:"http://www.opengis.net/gml/3.2 location,omitempty"`
	PriorityLocation     *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml/3.2 priorityLocation,omitempty"`
	ValidTime            *TimePrimitivePropertyType    `xml:"http://www.opengis.net/gml/3.2 validTime,omitempty"`
	Using                *ProcedurePropertyType        `xml:"http://www.opengis.net/gml/3.2 using,omitempty"`
	Target               *TargetPropertyType           `xml:"http://www.opengis.net/gml/3.2 target,omitempty"`
	Subject              *TargetPropertyType           `xml:"http://www.opengis.net/gml/3.2 subject,omitempty"`
	ResultOf             *ResultType                   `xml:"http://www.opengis.net/gml/3.2 resultOf,omitempty"`
}

type OffsetCurveType struct {
	NumDerivativesAtStart int                `xml:"numDerivativesAtStart,attr,omitempty"`
	NumDerivativesAtEnd   int                `xml:"numDerivativesAtEnd,attr,omitempty"`
	NumDerivativeInterior int                `xml:"numDerivativeInterior,attr,omitempty"`
	OffsetBase            *CurvePropertyType `xml:"http://www.opengis.net/gml/3.2 offsetBase,omitempty"`
	Distance              *LengthType        `xml:"http://www.opengis.net/gml/3.2 distance,omitempty"`
	RefDirection          *VectorType        `xml:"http://www.opengis.net/gml/3.2 refDirection,omitempty"`
}

type OperationMethodPropertyType struct {
	OperationMethod *OperationMethodType `xml:"http://www.opengis.net/gml/3.2 OperationMethod,omitempty"`
	NilReason       string               `xml:"nilReason,attr,omitempty"`
	RemoteSchema    string               `xml:"remoteSchema,attr,omitempty"`
}

type OperationMethodType struct {
	MetaDataProperty          []MetaDataPropertyType                          `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description               *StringOrRefType                                `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference      *ReferenceType                                  `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier                *CodeWithAuthorityType                          `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                      []CodeType                                      `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                        string                                          `xml:"id,attr,omitempty"`
	Remarks                   *string                                         `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	SourceDimensions          *int                                            `xml:"http://www.opengis.net/gml/3.2 sourceDimensions,omitempty"`
	TargetDimensions          *int                                            `xml:"http://www.opengis.net/gml/3.2 targetDimensions,omitempty"`
	Parameter                 []AbstractGeneralOperationParameterPropertyType `xml:"http://www.opengis.net/gml/3.2 parameter"`
	GeneralOperationParameter *AbstractGeneralOperationParameterPropertyType  `xml:"http://www.opengis.net/gml/3.2 generalOperationParameter,omitempty"`
	IncludesParameter         *AbstractGeneralOperationParameterPropertyType  `xml:"http://www.opengis.net/gml/3.2 includesParameter,omitempty"`
	FormulaCitation           *string                                         `xml:"http://www.opengis.net/gml/3.2 formulaCitation,omitempty"`
	Formula                   *CodeType                                       `xml:"http://www.opengis.net/gml/3.2 formula,omitempty"`
	MethodFormula             *CodeType                                       `xml:"http://www.opengis.net/gml/3.2 methodFormula,omitempty"`
}

type OperationParameterGroupPropertyType struct {
	OperationParameterGroup *OperationParameterGroupType `xml:"http://www.opengis.net/gml/3.2 OperationParameterGroup,omitempty"`
	NilReason               string                       `xml:"nilReason,attr,omitempty"`
	RemoteSchema            string                       `xml:"remoteSchema,attr,omitempty"`
}

type OperationParameterGroupType struct {
	MetaDataProperty          []MetaDataPropertyType                          `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description               *StringOrRefType                                `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference      *ReferenceType                                  `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier                *CodeWithAuthorityType                          `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                      []CodeType                                      `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                        string                                          `xml:"id,attr,omitempty"`
	Remarks                   *string                                         `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	MinimumOccurs             *int                                            `xml:"http://www.opengis.net/gml/3.2 minimumOccurs,omitempty"`
	MaximumOccurs             *int                                            `xml:"http://www.opengis.net/gml/3.2 maximumOccurs,omitempty"`
	Parameter                 []AbstractGeneralOperationParameterPropertyType `xml:"http://www.opengis.net/gml/3.2 parameter"`
	GeneralOperationParameter *AbstractGeneralOperationParameterPropertyType  `xml:"http://www.opengis.net/gml/3.2 generalOperationParameter,omitempty"`
	IncludesParameter         *AbstractGeneralOperationParameterPropertyType  `xml:"http://www.opengis.net/gml/3.2 includesParameter,omitempty"`
}

type OperationParameterPropertyType struct {
	OperationParameter *OperationParameterType `xml:"http://www.opengis.net/gml/3.2 OperationParameter,omitempty"`
	NilReason          string                  `xml:"nilReason,attr,omitempty"`
	RemoteSchema       string                  `xml:"remoteSchema,attr,omitempty"`
}

type OperationParameterType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	Remarks              *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	MinimumOccurs        *int                   `xml:"http://www.opengis.net/gml/3.2 minimumOccurs,omitempty"`
}

type OperationPropertyType struct {
	AbstractOperation             *AbstractCoordinateOperationType   `xml:"http://www.opengis.net/gml/3.2 AbstractOperation,omitempty"`
	AbstractGeneralConversion     *AbstractGeneralConversionType     `xml:"http://www.opengis.net/gml/3.2 AbstractGeneralConversion,omitempty"`
	AbstractGeneralTransformation *AbstractGeneralTransformationType `xml:"http://www.opengis.net/gml/3.2 AbstractGeneralTransformation,omitempty"`
	NilReason                     string                             `xml:"nilReason,attr,omitempty"`
	RemoteSchema                  string                             `xml:"remoteSchema,attr,omitempty"`
}

type OrientableCurveType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	SrsName              string                 `xml:"srsName,attr,omitempty"`
	SrsDimension         int                    `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                 `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                 `xml:"uomLabels,attr,omitempty"`
	BaseCurve            *CurvePropertyType     `xml:"http://www.opengis.net/gml/3.2 baseCurve,omitempty"`
	Orientation          string                 `xml:"orientation,attr,omitempty"`
}

type OrientableSurfaceType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	SrsName              string                 `xml:"srsName,attr,omitempty"`
	SrsDimension         int                    `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                 `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                 `xml:"uomLabels,attr,omitempty"`
	BaseSurface          *SurfacePropertyType   `xml:"http://www.opengis.net/gml/3.2 baseSurface,omitempty"`
	Orientation          string                 `xml:"orientation,attr,omitempty"`
}

type ParameterValueGroupType struct {
	ParameterValue []AbstractGeneralParameterValuePropertyType `xml:"http://www.opengis.net/gml/3.2 parameterValue"`
	IncludesValue  *AbstractGeneralParameterValuePropertyType  `xml:"http://www.opengis.net/gml/3.2 includesValue,omitempty"`
	UsesValue      *AbstractGeneralParameterValuePropertyType  `xml:"http://www.opengis.net/gml/3.2 usesValue,omitempty"`
	Group          *OperationParameterGroupPropertyType        `xml:"http://www.opengis.net/gml/3.2 group,omitempty"`
	ValuesOfGroup  *OperationParameterGroupPropertyType        `xml:"http://www.opengis.net/gml/3.2 valuesOfGroup,omitempty"`
}

type ParameterValueType struct {
	OperationParameter *OperationParameterPropertyType `xml:"http://www.opengis.net/gml/3.2 operationParameter,omitempty"`
	ValueOfParameter   *OperationParameterPropertyType `xml:"http://www.opengis.net/gml/3.2 valueOfParameter,omitempty"`
	Value              *MeasureType                    `xml:"http://www.opengis.net/gml/3.2 value,omitempty"`
	DmsAngleValue      *DMSAngleType                   `xml:"http://www.opengis.net/gml/3.2 dmsAngleValue,omitempty"`
	StringValue        *string                         `xml:"http://www.opengis.net/gml/3.2 stringValue,omitempty"`
	IntegerValue       *int                            `xml:"http://www.opengis.net/gml/3.2 integerValue,omitempty"`
	BooleanValue       *bool                           `xml:"http://www.opengis.net/gml/3.2 booleanValue,omitempty"`
	ValueList          *MeasureListType                `xml:"http://www.opengis.net/gml/3.2 valueList,omitempty"`
	IntegerValueList   *string                         `xml:"http://www.opengis.net/gml/3.2 integerValueList,omitempty"`
	ValueFile          *string                         `xml:"http://www.opengis.net/gml/3.2 valueFile,omitempty"`
}

type PassThroughOperationPropertyType struct {
	PassThroughOperation *PassThroughOperationType `xml:"http://www.opengis.net/gml/3.2 PassThroughOperation,omitempty"`
	NilReason            string                    `xml:"nilReason,attr,omitempty"`
	RemoteSchema         string                    `xml:"remoteSchema,attr,omitempty"`
}

type PassThroughOperationType struct {
	MetaDataProperty            []MetaDataPropertyType           `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description                 *StringOrRefType                 `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference        *ReferenceType                   `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier                  *CodeWithAuthorityType           `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                        []CodeType                       `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                          string                           `xml:"id,attr,omitempty"`
	Remarks                     *string                          `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity            *string                          `xml:"http://www.opengis.net/gml/3.2 domainOfValidity,omitempty"`
	Scope                       []string                         `xml:"http://www.opengis.net/gml/3.2 scope"`
	OperationVersion            *string                          `xml:"http://www.opengis.net/gml/3.2 operationVersion,omitempty"`
	CoordinateOperationAccuracy []string                         `xml:"http://www.opengis.net/gml/3.2 coordinateOperationAccuracy"`
	SourceCRS                   *CRSPropertyType                 `xml:"http://www.opengis.net/gml/3.2 sourceCRS,omitempty"`
	TargetCRS                   *CRSPropertyType                 `xml:"http://www.opengis.net/gml/3.2 targetCRS,omitempty"`
	ModifiedCoordinate          []int                            `xml:"http://www.opengis.net/gml/3.2 modifiedCoordinate"`
	CoordOperation              *CoordinateOperationPropertyType `xml:"http://www.opengis.net/gml/3.2 coordOperation,omitempty"`
	UsesOperation               *CoordinateOperationPropertyType `xml:"http://www.opengis.net/gml/3.2 usesOperation,omitempty"`
	UsesSingleOperation         *CoordinateOperationPropertyType `xml:"http://www.opengis.net/gml/3.2 usesSingleOperation,omitempty"`
	AggregationType             string                           `xml:"aggregationType,attr,omitempty"`
}

type PointArrayPropertyType struct {
	Point *PointType `xml:"http://www.opengis.net/gml/3.2 Point,omitempty"`
	Owns  bool       `xml:"owns,attr,omitempty"`
}

type PointPropertyType struct {
	Point        *PointType `xml:"http://www.opengis.net/gml/3.2 Point,omitempty"`
	NilReason    string     `xml:"nilReason,attr,omitempty"`
	RemoteSchema string     `xml:"remoteSchema,attr,omitempty"`
	Owns         bool       `xml:"owns,attr,omitempty"`
}

type PointType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	SrsName              string                 `xml:"srsName,attr,omitempty"`
	SrsDimension         int                    `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                 `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                 `xml:"uomLabels,attr,omitempty"`
	Pos                  *DirectPositionType    `xml:"http://www.opengis.net/gml/3.2 pos,omitempty"`
	Coordinates          *CoordinatesType       `xml:"http://www.opengis.net/gml/3.2 coordinates,omitempty"`
}

type PolarCSPropertyType struct {
	PolarCS      *PolarCSType `xml:"http://www.opengis.net/gml/3.2 PolarCS,omitempty"`
	NilReason    string       `xml:"nilReason,attr,omitempty"`
	RemoteSchema string       `xml:"remoteSchema,attr,omitempty"`
}

type PolarCSType struct {
	MetaDataProperty     []MetaDataPropertyType             `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType                   `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                     `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType             `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                         `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                             `xml:"id,attr,omitempty"`
	Remarks              *string                            `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	Axis                 []CoordinateSystemAxisPropertyType `xml:"http://www.opengis.net/gml/3.2 axis"`
	UsesAxis             *CoordinateSystemAxisPropertyType  `xml:"http://www.opengis.net/gml/3.2 usesAxis,omitempty"`
	AggregationType      string                             `xml:"aggregationType,attr,omitempty"`
}

type PolygonPatchType struct {
	Exterior      *AbstractRingPropertyType  `xml:"http://www.opengis.net/gml/3.2 exterior,omitempty"`
	Interior      []AbstractRingPropertyType `xml:"http://www.opengis.net/gml/3.2 interior"`
	Interpolation string                     `xml:"interpolation,attr,omitempty"`
}

type PolygonType struct {
	MetaDataProperty     []MetaDataPropertyType     `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType           `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType             `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType     `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                 `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                     `xml:"id,attr,omitempty"`
	SrsName              string                     `xml:"srsName,attr,omitempty"`
	SrsDimension         int                        `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                     `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                     `xml:"uomLabels,attr,omitempty"`
	Exterior             *AbstractRingPropertyType  `xml:"http://www.opengis.net/gml/3.2 exterior,omitempty"`
	Interior             []AbstractRingPropertyType `xml:"http://www.opengis.net/gml/3.2 interior"`
}

type PrimeMeridianPropertyType struct {
	PrimeMeridian *PrimeMeridianType `xml:"http://www.opengis.net/gml/3.2 PrimeMeridian,omitempty"`
	NilReason     string             `xml:"nilReason,attr,omitempty"`
	RemoteSchema  string             `xml:"remoteSchema,attr,omitempty"`
}

type PrimeMeridianType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	Remarks              *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	GreenwichLongitude   *AngleType             `xml:"http://www.opengis.net/gml/3.2 greenwichLongitude,omitempty"`
}

type PriorityLocationPropertyType struct {
	AbstractGeometry           *AbstractGeometryType           `xml:"http://www.opengis.net/gml/3.2 AbstractGeometry,omitempty"`
	AbstractGeometricAggregate *AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml/3.2 AbstractGeometricAggregate,omitempty"`
	AbstractGeometricPrimitive *AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml/3.2 AbstractGeometricPrimitive,omitempty"`
	AbstractImplicitGeometry   *AbstractGeometryType           `xml:"http://www.opengis.net/gml/3.2 AbstractImplicitGeometry,omitempty"`
	GeometricComplex           *GeometricComplexType           `xml:"http://www.opengis.net/gml/3.2 GeometricComplex,omitempty"`
	LocationKeyWord            *CodeType                       `xml:"http://www.opengis.net/gml/3.2 LocationKeyWord,omitempty"`
	LocationString             *StringOrRefType                `xml:"http://www.opengis.net/gml/3.2 LocationString,omitempty"`
	Null                       *string                         `xml:"http://www.opengis.net/gml/3.2 Null,omitempty"`
	NilReason                  string                          `xml:"nilReason,attr,omitempty"`
	RemoteSchema               string                          `xml:"remoteSchema,attr,omitempty"`
	Priority                   string                          `xml:"priority,attr,omitempty"`
}

type ProcedurePropertyType struct {
	AbstractFeature            *AbstractFeatureType            `xml:"http://www.opengis.net/gml/3.2 AbstractFeature,omitempty"`
	AbstractContinuousCoverage *AbstractContinuousCoverageType `xml:"http://www.opengis.net/gml/3.2 AbstractContinuousCoverage,omitempty"`
	AbstractCoverage           *AbstractCoverageType           `xml:"http://www.opengis.net/gml/3.2 AbstractCoverage,omitempty"`
	AbstractFeatureCollection  *AbstractFeatureCollectionType  `xml:"http://www.opengis.net/gml/3.2 AbstractFeatureCollection,omitempty"`
	DynamicFeature             *DynamicFeatureType             `xml:"http://www.opengis.net/gml/3.2 DynamicFeature,omitempty"`
	FeatureCollection          *FeatureCollectionType          `xml:"http://www.opengis.net/gml/3.2 FeatureCollection,omitempty"`
	Observation                *ObservationType                `xml:"http://www.opengis.net/gml/3.2 Observation,omitempty"`
	Owns                       bool                            `xml:"owns,attr,omitempty"`
	NilReason                  string                          `xml:"nilReason,attr,omitempty"`
	RemoteSchema               string                          `xml:"remoteSchema,attr,omitempty"`
}

type ProjectedCRSPropertyType struct {
	ProjectedCRS *ProjectedCRSType `xml:"http://www.opengis.net/gml/3.2 ProjectedCRS,omitempty"`
	NilReason    string            `xml:"nilReason,attr,omitempty"`
	RemoteSchema string            `xml:"remoteSchema,attr,omitempty"`
}

type ProjectedCRSType struct {
	MetaDataProperty     []MetaDataPropertyType         `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType               `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                 `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType         `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                     `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                         `xml:"id,attr,omitempty"`
	Remarks              *string                        `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     []string                       `xml:"http://www.opengis.net/gml/3.2 domainOfValidity"`
	Scope                []string                       `xml:"http://www.opengis.net/gml/3.2 scope"`
	Conversion           *GeneralConversionPropertyType `xml:"http://www.opengis.net/gml/3.2 conversion,omitempty"`
	DefinedByConversion  *GeneralConversionPropertyType `xml:"http://www.opengis.net/gml/3.2 definedByConversion,omitempty"`
	CartesianCS          *CartesianCSPropertyType       `xml:"http://www.opengis.net/gml/3.2 cartesianCS,omitempty"`
	UsesCartesianCS      *CartesianCSPropertyType       `xml:"http://www.opengis.net/gml/3.2 usesCartesianCS,omitempty"`
	BaseGeodeticCRS      *GeodeticCRSPropertyType       `xml:"http://www.opengis.net/gml/3.2 baseGeodeticCRS,omitempty"`
	BaseGeographicCRS    *GeographicCRSPropertyType     `xml:"http://www.opengis.net/gml/3.2 baseGeographicCRS,omitempty"`
}

type QuantityExtentType struct {
	Value string `xml:",chardata"`
}

type QuantityPropertyType struct {
	Quantity     *string `xml:"http://www.opengis.net/gml/3.2 Quantity,omitempty"`
	NilReason    string  `xml:"nilReason,attr,omitempty"`
	RemoteSchema string  `xml:"remoteSchema,attr,omitempty"`
}

type RangeSetType struct {
	ValueArray              []ValueArrayType            `xml:"http://www.opengis.net/gml/3.2 ValueArray"`
	AbstractScalarValueList []string                    `xml:"http://www.opengis.net/gml/3.2 AbstractScalarValueList"`
	BooleanList             *string                     `xml:"http://www.opengis.net/gml/3.2 BooleanList,omitempty"`
	CategoryList            *CodeOrNilReasonListType    `xml:"http://www.opengis.net/gml/3.2 CategoryList,omitempty"`
	CountList               *string                     `xml:"http://www.opengis.net/gml/3.2 CountList,omitempty"`
	QuantityList            *MeasureOrNilReasonListType `xml:"http://www.opengis.net/gml/3.2 QuantityList,omitempty"`
	DataBlock               *DataBlockType              `xml:"http://www.opengis.net/gml/3.2 DataBlock,omitempty"`
	File                    *FileType                   `xml:"http://www.opengis.net/gml/3.2 File,omitempty"`
}

type RectangleType struct {
	Exterior      *AbstractRingPropertyType `xml:"http://www.opengis.net/gml/3.2 exterior,omitempty"`
	Interpolation string                    `xml:"interpolation,attr,omitempty"`
}

type RectifiedGridType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	SrsName              string                 `xml:"srsName,attr,omitempty"`
	SrsDimension         int                    `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                 `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                 `xml:"uomLabels,attr,omitempty"`
	Limits               *GridLimitsType        `xml:"http://www.opengis.net/gml/3.2 limits,omitempty"`
	AxisName             []string               `xml:"http://www.opengis.net/gml/3.2 axisName"`
	Dimension            int                    `xml:"dimension,attr"`
	Origin               *PointPropertyType     `xml:"http://www.opengis.net/gml/3.2 origin,omitempty"`
	OffsetVector         []VectorType           `xml:"http://www.opengis.net/gml/3.2 offsetVector"`
}

type ReferenceType struct {
	Owns         bool   `xml:"owns,attr,omitempty"`
	NilReason    string `xml:"nilReason,attr,omitempty"`
	RemoteSchema string `xml:"remoteSchema,attr,omitempty"`
}

type RelatedTimeType struct {
	AbstractTimePrimitive          *AbstractTimePrimitiveType          `xml:"http://www.opengis.net/gml/3.2 AbstractTimePrimitive,omitempty"`
	AbstractTimeGeometricPrimitive *AbstractTimeGeometricPrimitiveType `xml:"http://www.opengis.net/gml/3.2 AbstractTimeGeometricPrimitive,omitempty"`
	AbstractTimeTopologyPrimitive  *AbstractTimeTopologyPrimitiveType  `xml:"http://www.opengis.net/gml/3.2 AbstractTimeTopologyPrimitive,omitempty"`
	NilReason                      string                              `xml:"nilReason,attr,omitempty"`
	RemoteSchema                   string                              `xml:"remoteSchema,attr,omitempty"`
	Owns                           bool                                `xml:"owns,attr,omitempty"`
	RelativePosition               string                              `xml:"relativePosition,attr,omitempty"`
}

type ResultType struct {
	Owns         bool   `xml:"owns,attr,omitempty"`
	NilReason    string `xml:"nilReason,attr,omitempty"`
	RemoteSchema string `xml:"remoteSchema,attr,omitempty"`
}

type RingPropertyType struct {
	Ring *RingType `xml:"http://www.opengis.net/gml/3.2 Ring,omitempty"`
}

type RingType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	SrsName              string                 `xml:"srsName,attr,omitempty"`
	SrsDimension         int                    `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                 `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                 `xml:"uomLabels,attr,omitempty"`
	CurveMember          []CurvePropertyType    `xml:"http://www.opengis.net/gml/3.2 curveMember"`
	AggregationType      string                 `xml:"aggregationType,attr,omitempty"`
}

type ScaleType struct {
	Value string `xml:",chardata"`
}

type SequenceRuleType struct {
	Value     string `xml:",chardata"`
	Order     string `xml:"order,attr,omitempty"`
	AxisOrder string `xml:"axisOrder,attr,omitempty"`
}

type ShellPropertyType struct {
	Shell *ShellType `xml:"http://www.opengis.net/gml/3.2 Shell,omitempty"`
}

type ShellType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	SrsName              string                 `xml:"srsName,attr,omitempty"`
	SrsDimension         int                    `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                 `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                 `xml:"uomLabels,attr,omitempty"`
	SurfaceMember        []SurfacePropertyType  `xml:"http://www.opengis.net/gml/3.2 surfaceMember"`
	AggregationType      string                 `xml:"aggregationType,attr,omitempty"`
}

type SingleCRSPropertyType struct {
	AbstractSingleCRS         *AbstractCRSType               `xml:"http://www.opengis.net/gml/3.2 AbstractSingleCRS,omitempty"`
	AbstractGeneralDerivedCRS *AbstractGeneralDerivedCRSType `xml:"http://www.opengis.net/gml/3.2 AbstractGeneralDerivedCRS,omitempty"`
	EngineeringCRS            *EngineeringCRSType            `xml:"http://www.opengis.net/gml/3.2 EngineeringCRS,omitempty"`
	GeocentricCRS             *GeocentricCRSType             `xml:"http://www.opengis.net/gml/3.2 GeocentricCRS,omitempty"`
	GeodeticCRS               *GeodeticCRSType               `xml:"http://www.opengis.net/gml/3.2 GeodeticCRS,omitempty"`
	GeographicCRS             *GeographicCRSType             `xml:"http://www.opengis.net/gml/3.2 GeographicCRS,omitempty"`
	ImageCRS                  *ImageCRSType                  `xml:"http://www.opengis.net/gml/3.2 ImageCRS,omitempty"`
	TemporalCRS               *TemporalCRSType               `xml:"http://www.opengis.net/gml/3.2 TemporalCRS,omitempty"`
	VerticalCRS               *VerticalCRSType               `xml:"http://www.opengis.net/gml/3.2 VerticalCRS,omitempty"`
	NilReason                 string                         `xml:"nilReason,attr,omitempty"`
	RemoteSchema              string                         `xml:"remoteSchema,attr,omitempty"`
}

type SingleOperationPropertyType struct {
	AbstractSingleOperation *AbstractCoordinateOperationType `xml:"http://www.opengis.net/gml/3.2 AbstractSingleOperation,omitempty"`
	AbstractOperation       *AbstractCoordinateOperationType `xml:"http://www.opengis.net/gml/3.2 AbstractOperation,omitempty"`
	PassThroughOperation    *PassThroughOperationType        `xml:"http://www.opengis.net/gml/3.2 PassThroughOperation,omitempty"`
	NilReason               string                           `xml:"nilReason,attr,omitempty"`
	RemoteSchema            string                           `xml:"remoteSchema,attr,omitempty"`
}

type SolidArrayPropertyType struct {
	AbstractSolid  *AbstractSolidType  `xml:"http://www.opengis.net/gml/3.2 AbstractSolid,omitempty"`
	CompositeSolid *CompositeSolidType `xml:"http://www.opengis.net/gml/3.2 CompositeSolid,omitempty"`
	Solid          *SolidType          `xml:"http://www.opengis.net/gml/3.2 Solid,omitempty"`
	Owns           bool                `xml:"owns,attr,omitempty"`
}

type SolidPropertyType struct {
	AbstractSolid  *AbstractSolidType  `xml:"http://www.opengis.net/gml/3.2 AbstractSolid,omitempty"`
	CompositeSolid *CompositeSolidType `xml:"http://www.opengis.net/gml/3.2 CompositeSolid,omitempty"`
	Solid          *SolidType          `xml:"http://www.opengis.net/gml/3.2 Solid,omitempty"`
	NilReason      string              `xml:"nilReason,attr,omitempty"`
	RemoteSchema   string              `xml:"remoteSchema,attr,omitempty"`
	Owns           bool                `xml:"owns,attr,omitempty"`
}

type SolidType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	SrsName              string                 `xml:"srsName,attr,omitempty"`
	SrsDimension         int                    `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                 `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                 `xml:"uomLabels,attr,omitempty"`
	Exterior             *ShellPropertyType     `xml:"http://www.opengis.net/gml/3.2 exterior,omitempty"`
	Interior             []ShellPropertyType    `xml:"http://www.opengis.net/gml/3.2 interior"`
}

type SpeedType struct {
	Value string `xml:",chardata"`
}

type SphereType struct {
	AggregationType     string `xml:"aggregationType,attr,omitempty"`
	Rows                string `xml:"http://www.opengis.net/gml/3.2 rows,omitempty"`
	Columns             int    `xml:"columns,attr,omitempty"`
	HorizontalCurveType string `xml:"horizontalCurveType,attr,omitempty"`
	VerticalCurveType   string `xml:"verticalCurveType,attr,omitempty"`
}

type SphericalCSPropertyType struct {
	SphericalCS  *SphericalCSType `xml:"http://www.opengis.net/gml/3.2 SphericalCS,omitempty"`
	NilReason    string           `xml:"nilReason,attr,omitempty"`
	RemoteSchema string           `xml:"remoteSchema,attr,omitempty"`
}

type SphericalCSType struct {
	MetaDataProperty     []MetaDataPropertyType             `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType                   `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                     `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType             `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                         `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                             `xml:"id,attr,omitempty"`
	Remarks              *string                            `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	Axis                 []CoordinateSystemAxisPropertyType `xml:"http://www.opengis.net/gml/3.2 axis"`
	UsesAxis             *CoordinateSystemAxisPropertyType  `xml:"http://www.opengis.net/gml/3.2 usesAxis,omitempty"`
	AggregationType      string                             `xml:"aggregationType,attr,omitempty"`
}

type StringOrRefType struct {
	Value        string `xml:",chardata"`
	NilReason    string `xml:"nilReason,attr,omitempty"`
	RemoteSchema string `xml:"remoteSchema,attr,omitempty"`
}

type SurfaceArrayPropertyType struct {
	AbstractSurface   *AbstractSurfaceType   `xml:"http://www.opengis.net/gml/3.2 AbstractSurface,omitempty"`
	CompositeSurface  *CompositeSurfaceType  `xml:"http://www.opengis.net/gml/3.2 CompositeSurface,omitempty"`
	OrientableSurface *OrientableSurfaceType `xml:"http://www.opengis.net/gml/3.2 OrientableSurface,omitempty"`
	Polygon           *PolygonType           `xml:"http://www.opengis.net/gml/3.2 Polygon,omitempty"`
	Shell             *ShellType             `xml:"http://www.opengis.net/gml/3.2 Shell,omitempty"`
	Surface           *SurfaceType           `xml:"http://www.opengis.net/gml/3.2 Surface,omitempty"`
	Owns              bool                   `xml:"owns,attr,omitempty"`
}

type SurfacePatchArrayPropertyType struct {
	AbstractSurfacePatch           *AbstractSurfacePatchType           `xml:"http://www.opengis.net/gml/3.2 AbstractSurfacePatch,omitempty"`
	AbstractParametricCurveSurface *AbstractParametricCurveSurfaceType `xml:"http://www.opengis.net/gml/3.2 AbstractParametricCurveSurface,omitempty"`
	PolygonPatch                   *PolygonPatchType                   `xml:"http://www.opengis.net/gml/3.2 PolygonPatch,omitempty"`
	Rectangle                      *RectangleType                      `xml:"http://www.opengis.net/gml/3.2 Rectangle,omitempty"`
	Triangle                       *TriangleType                       `xml:"http://www.opengis.net/gml/3.2 Triangle,omitempty"`
}

type SurfacePropertyType struct {
	AbstractSurface   *AbstractSurfaceType   `xml:"http://www.opengis.net/gml/3.2 AbstractSurface,omitempty"`
	CompositeSurface  *CompositeSurfaceType  `xml:"http://www.opengis.net/gml/3.2 CompositeSurface,omitempty"`
	OrientableSurface *OrientableSurfaceType `xml:"http://www.opengis.net/gml/3.2 OrientableSurface,omitempty"`
	Polygon           *PolygonType           `xml:"http://www.opengis.net/gml/3.2 Polygon,omitempty"`
	Shell             *ShellType             `xml:"http://www.opengis.net/gml/3.2 Shell,omitempty"`
	Surface           *SurfaceType           `xml:"http://www.opengis.net/gml/3.2 Surface,omitempty"`
	NilReason         string                 `xml:"nilReason,attr,omitempty"`
	RemoteSchema      string                 `xml:"remoteSchema,attr,omitempty"`
	Owns              bool                   `xml:"owns,attr,omitempty"`
}

type SurfaceType struct {
	MetaDataProperty     []MetaDataPropertyType         `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType               `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                 `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType         `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                     `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                         `xml:"id,attr,omitempty"`
	SrsName              string                         `xml:"srsName,attr,omitempty"`
	SrsDimension         int                            `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                         `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                         `xml:"uomLabels,attr,omitempty"`
	Patches              *SurfacePatchArrayPropertyType `xml:"http://www.opengis.net/gml/3.2 patches,omitempty"`
	PolygonPatches       *SurfacePatchArrayPropertyType `xml:"http://www.opengis.net/gml/3.2 polygonPatches,omitempty"`
	TrianglePatches      *SurfacePatchArrayPropertyType `xml:"http://www.opengis.net/gml/3.2 trianglePatches,omitempty"`
}

type TargetPropertyType struct {
	AbstractFeature            *AbstractFeatureType            `xml:"http://www.opengis.net/gml/3.2 AbstractFeature,omitempty"`
	AbstractContinuousCoverage *AbstractContinuousCoverageType `xml:"http://www.opengis.net/gml/3.2 AbstractContinuousCoverage,omitempty"`
	AbstractCoverage           *AbstractCoverageType           `xml:"http://www.opengis.net/gml/3.2 AbstractCoverage,omitempty"`
	AbstractFeatureCollection  *AbstractFeatureCollectionType  `xml:"http://www.opengis.net/gml/3.2 AbstractFeatureCollection,omitempty"`
	DynamicFeature             *DynamicFeatureType             `xml:"http://www.opengis.net/gml/3.2 DynamicFeature,omitempty"`
	FeatureCollection          *FeatureCollectionType          `xml:"http://www.opengis.net/gml/3.2 FeatureCollection,omitempty"`
	Observation                *ObservationType                `xml:"http://www.opengis.net/gml/3.2 Observation,omitempty"`
	AbstractGeometry           *AbstractGeometryType           `xml:"http://www.opengis.net/gml/3.2 AbstractGeometry,omitempty"`
	AbstractGeometricAggregate *AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml/3.2 AbstractGeometricAggregate,omitempty"`
	AbstractGeometricPrimitive *AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml/3.2 AbstractGeometricPrimitive,omitempty"`
	AbstractImplicitGeometry   *AbstractGeometryType           `xml:"http://www.opengis.net/gml/3.2 AbstractImplicitGeometry,omitempty"`
	GeometricComplex           *GeometricComplexType           `xml:"http://www.opengis.net/gml/3.2 GeometricComplex,omitempty"`
	Owns                       bool                            `xml:"owns,attr,omitempty"`
	NilReason                  string                          `xml:"nilReason,attr,omitempty"`
	RemoteSchema               string                          `xml:"remoteSchema,attr,omitempty"`
}

type TemporalCRSPropertyType struct {
	TemporalCRS  *TemporalCRSType `xml:"http://www.opengis.net/gml/3.2 TemporalCRS,omitempty"`
	NilReason    string           `xml:"nilReason,attr,omitempty"`
	RemoteSchema string           `xml:"remoteSchema,attr,omitempty"`
}

type TemporalCRSType struct {
	MetaDataProperty     []MetaDataPropertyType     `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType           `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType             `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType     `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                 `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                     `xml:"id,attr,omitempty"`
	Remarks              *string                    `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     []string                   `xml:"http://www.opengis.net/gml/3.2 domainOfValidity"`
	Scope                []string                   `xml:"http://www.opengis.net/gml/3.2 scope"`
	TemporalDatum        *TemporalDatumPropertyType `xml:"http://www.opengis.net/gml/3.2 temporalDatum,omitempty"`
	UsesTemporalDatum    *TemporalDatumPropertyType `xml:"http://www.opengis.net/gml/3.2 usesTemporalDatum,omitempty"`
	TimeCS               *TimeCSPropertyType        `xml:"http://www.opengis.net/gml/3.2 timeCS,omitempty"`
	UsesTimeCS           *TimeCSPropertyType        `xml:"http://www.opengis.net/gml/3.2 usesTimeCS,omitempty"`
	UsesTemporalCS       *TemporalCSPropertyType    `xml:"http://www.opengis.net/gml/3.2 usesTemporalCS,omitempty"`
}

type TemporalCSPropertyType struct {
	TemporalCS   *TemporalCSType `xml:"http://www.opengis.net/gml/3.2 TemporalCS,omitempty"`
	NilReason    string          `xml:"nilReason,attr,omitempty"`
	RemoteSchema string          `xml:"remoteSchema,attr,omitempty"`
}

type TemporalCSType struct {
	MetaDataProperty     []MetaDataPropertyType             `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType                   `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                     `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType             `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                         `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                             `xml:"id,attr,omitempty"`
	Remarks              *string                            `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	Axis                 []CoordinateSystemAxisPropertyType `xml:"http://www.opengis.net/gml/3.2 axis"`
	UsesAxis             *CoordinateSystemAxisPropertyType  `xml:"http://www.opengis.net/gml/3.2 usesAxis,omitempty"`
	AggregationType      string                             `xml:"aggregationType,attr,omitempty"`
}

type TemporalDatumBaseType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	Remarks              *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     *string                `xml:"http://www.opengis.net/gml/3.2 domainOfValidity,omitempty"`
	Scope                []string               `xml:"http://www.opengis.net/gml/3.2 scope"`
	AnchorDefinition     *CodeType              `xml:"http://www.opengis.net/gml/3.2 anchorDefinition,omitempty"`
	AnchorPoint          *CodeType              `xml:"http://www.opengis.net/gml/3.2 anchorPoint,omitempty"`
	RealizationEpoch     *string                `xml:"http://www.opengis.net/gml/3.2 realizationEpoch,omitempty"`
}

type TemporalDatumPropertyType struct {
	TemporalDatum *TemporalDatumType `xml:"http://www.opengis.net/gml/3.2 TemporalDatum,omitempty"`
	NilReason     string             `xml:"nilReason,attr,omitempty"`
	RemoteSchema  string             `xml:"remoteSchema,attr,omitempty"`
}

type TemporalDatumType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	Remarks              *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     *string                `xml:"http://www.opengis.net/gml/3.2 domainOfValidity,omitempty"`
	Scope                []string               `xml:"http://www.opengis.net/gml/3.2 scope"`
	AnchorDefinition     *CodeType              `xml:"http://www.opengis.net/gml/3.2 anchorDefinition,omitempty"`
	AnchorPoint          *CodeType              `xml:"http://www.opengis.net/gml/3.2 anchorPoint,omitempty"`
	RealizationEpoch     *string                `xml:"http://www.opengis.net/gml/3.2 realizationEpoch,omitempty"`
	Origin               *string                `xml:"http://www.opengis.net/gml/3.2 origin,omitempty"`
}

type TimeCSPropertyType struct {
	TimeCS       *TimeCSType `xml:"http://www.opengis.net/gml/3.2 TimeCS,omitempty"`
	NilReason    string      `xml:"nilReason,attr,omitempty"`
	RemoteSchema string      `xml:"remoteSchema,attr,omitempty"`
}

type TimeCSType struct {
	MetaDataProperty     []MetaDataPropertyType             `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType                   `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                     `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType             `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                         `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                             `xml:"id,attr,omitempty"`
	Remarks              *string                            `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	Axis                 []CoordinateSystemAxisPropertyType `xml:"http://www.opengis.net/gml/3.2 axis"`
	UsesAxis             *CoordinateSystemAxisPropertyType  `xml:"http://www.opengis.net/gml/3.2 usesAxis,omitempty"`
	AggregationType      string                             `xml:"aggregationType,attr,omitempty"`
}

type TimeCalendarEraPropertyType struct {
	TimeCalendarEra *TimeCalendarEraType `xml:"http://www.opengis.net/gml/3.2 TimeCalendarEra,omitempty"`
	Owns            bool                 `xml:"owns,attr,omitempty"`
	NilReason       string               `xml:"nilReason,attr,omitempty"`
	RemoteSchema    string               `xml:"remoteSchema,attr,omitempty"`
}

type TimeCalendarEraType struct {
	MetaDataProperty     []MetaDataPropertyType  `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType        `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType          `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType  `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType              `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                  `xml:"id,attr,omitempty"`
	Remarks              *string                 `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	ReferenceEvent       *StringOrRefType        `xml:"http://www.opengis.net/gml/3.2 referenceEvent,omitempty"`
	ReferenceDate        string                  `xml:"http://www.opengis.net/gml/3.2 referenceDate,omitempty"`
	JulianReference      float64                 `xml:"http://www.opengis.net/gml/3.2 julianReference,omitempty"`
	EpochOfUse           *TimePeriodPropertyType `xml:"http://www.opengis.net/gml/3.2 epochOfUse,omitempty"`
}

type TimeCalendarPropertyType struct {
	TimeCalendar *TimeCalendarType `xml:"http://www.opengis.net/gml/3.2 TimeCalendar,omitempty"`
	Owns         bool              `xml:"owns,attr,omitempty"`
	NilReason    string            `xml:"nilReason,attr,omitempty"`
	RemoteSchema string            `xml:"remoteSchema,attr,omitempty"`
}

type TimeCalendarType struct {
	MetaDataProperty     []MetaDataPropertyType        `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType              `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType        `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                    `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                        `xml:"id,attr,omitempty"`
	Remarks              *string                       `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     string                        `xml:"http://www.opengis.net/gml/3.2 domainOfValidity,omitempty"`
	ReferenceFrame       []TimeCalendarEraPropertyType `xml:"http://www.opengis.net/gml/3.2 referenceFrame"`
}

type TimeClockPropertyType struct {
	TimeClock    *TimeClockType `xml:"http://www.opengis.net/gml/3.2 TimeClock,omitempty"`
	Owns         bool           `xml:"owns,attr,omitempty"`
	NilReason    string         `xml:"nilReason,attr,omitempty"`
	RemoteSchema string         `xml:"remoteSchema,attr,omitempty"`
}

type TimeClockType struct {
	MetaDataProperty     []MetaDataPropertyType     `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType           `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType             `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType     `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                 `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                     `xml:"id,attr,omitempty"`
	Remarks              *string                    `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     string                     `xml:"http://www.opengis.net/gml/3.2 domainOfValidity,omitempty"`
	ReferenceEvent       *StringOrRefType           `xml:"http://www.opengis.net/gml/3.2 referenceEvent,omitempty"`
	ReferenceTime        string                     `xml:"http://www.opengis.net/gml/3.2 referenceTime,omitempty"`
	UtcReference         string                     `xml:"http://www.opengis.net/gml/3.2 utcReference,omitempty"`
	DateBasis            []TimeCalendarPropertyType `xml:"http://www.opengis.net/gml/3.2 dateBasis"`
}

type TimeCoordinateSystemType struct {
	MetaDataProperty     []MetaDataPropertyType   `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType         `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType           `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType   `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType               `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                   `xml:"id,attr,omitempty"`
	Remarks              *string                  `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     string                   `xml:"http://www.opengis.net/gml/3.2 domainOfValidity,omitempty"`
	Interval             *TimeIntervalLengthType  `xml:"http://www.opengis.net/gml/3.2 interval,omitempty"`
	OriginPosition       *TimePositionType        `xml:"http://www.opengis.net/gml/3.2 originPosition,omitempty"`
	Origin               *TimeInstantPropertyType `xml:"http://www.opengis.net/gml/3.2 origin,omitempty"`
}

type TimeEdgePropertyType struct {
	TimeEdge     *TimeEdgeType `xml:"http://www.opengis.net/gml/3.2 TimeEdge,omitempty"`
	NilReason    string        `xml:"nilReason,attr,omitempty"`
	RemoteSchema string        `xml:"remoteSchema,attr,omitempty"`
	Owns         bool          `xml:"owns,attr,omitempty"`
}

type TimeEdgeType struct {
	MetaDataProperty     []MetaDataPropertyType  `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType        `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType          `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType  `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType              `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                  `xml:"id,attr,omitempty"`
	RelatedTime          []RelatedTimeType       `xml:"http://www.opengis.net/gml/3.2 relatedTime"`
	Complex              *ReferenceType          `xml:"http://www.opengis.net/gml/3.2 complex,omitempty"`
	Start                *TimeNodePropertyType   `xml:"http://www.opengis.net/gml/3.2 start,omitempty"`
	End                  *TimeNodePropertyType   `xml:"http://www.opengis.net/gml/3.2 end,omitempty"`
	Extent               *TimePeriodPropertyType `xml:"http://www.opengis.net/gml/3.2 extent,omitempty"`
}

type TimeInstantPropertyType struct {
	TimeInstant  *TimeInstantType `xml:"http://www.opengis.net/gml/3.2 TimeInstant,omitempty"`
	NilReason    string           `xml:"nilReason,attr,omitempty"`
	RemoteSchema string           `xml:"remoteSchema,attr,omitempty"`
	Owns         bool             `xml:"owns,attr,omitempty"`
}

type TimeInstantType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	RelatedTime          []RelatedTimeType      `xml:"http://www.opengis.net/gml/3.2 relatedTime"`
	Frame                string                 `xml:"frame,attr,omitempty"`
	TimePosition         *TimePositionType      `xml:"http://www.opengis.net/gml/3.2 timePosition,omitempty"`
}

type TimeIntervalLengthType struct {
	Value  float64 `xml:",chardata"`
	Unit   string  `xml:"unit,attr"`
	Radix  int     `xml:"radix,attr,omitempty"`
	Factor int     `xml:"factor,attr,omitempty"`
}

type TimeNodePropertyType struct {
	TimeNode     *TimeNodeType `xml:"http://www.opengis.net/gml/3.2 TimeNode,omitempty"`
	NilReason    string        `xml:"nilReason,attr,omitempty"`
	RemoteSchema string        `xml:"remoteSchema,attr,omitempty"`
	Owns         bool          `xml:"owns,attr,omitempty"`
}

type TimeNodeType struct {
	MetaDataProperty     []MetaDataPropertyType   `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType         `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType           `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType   `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType               `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                   `xml:"id,attr,omitempty"`
	RelatedTime          []RelatedTimeType        `xml:"http://www.opengis.net/gml/3.2 relatedTime"`
	Complex              *ReferenceType           `xml:"http://www.opengis.net/gml/3.2 complex,omitempty"`
	PreviousEdge         []TimeEdgePropertyType   `xml:"http://www.opengis.net/gml/3.2 previousEdge"`
	NextEdge             []TimeEdgePropertyType   `xml:"http://www.opengis.net/gml/3.2 nextEdge"`
	Position             *TimeInstantPropertyType `xml:"http://www.opengis.net/gml/3.2 position,omitempty"`
}

type TimeOrdinalEraPropertyType struct {
	TimeOrdinalEra *TimeOrdinalEraType `xml:"http://www.opengis.net/gml/3.2 TimeOrdinalEra,omitempty"`
	Owns           bool                `xml:"owns,attr,omitempty"`
	NilReason      string              `xml:"nilReason,attr,omitempty"`
	RemoteSchema   string              `xml:"remoteSchema,attr,omitempty"`
}

type TimeOrdinalEraType struct {
	MetaDataProperty     []MetaDataPropertyType       `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType             `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType               `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType       `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                   `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                       `xml:"id,attr,omitempty"`
	Remarks              *string                      `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	RelatedTime          []RelatedTimeType            `xml:"http://www.opengis.net/gml/3.2 relatedTime"`
	Start                *TimeNodePropertyType        `xml:"http://www.opengis.net/gml/3.2 start,omitempty"`
	End                  *TimeNodePropertyType        `xml:"http://www.opengis.net/gml/3.2 end,omitempty"`
	Extent               *TimePeriodPropertyType      `xml:"http://www.opengis.net/gml/3.2 extent,omitempty"`
	Member               []TimeOrdinalEraPropertyType `xml:"http://www.opengis.net/gml/3.2 member"`
	Group                *ReferenceType               `xml:"http://www.opengis.net/gml/3.2 group,omitempty"`
}

type TimeOrdinalReferenceSystemType struct {
	MetaDataProperty     []MetaDataPropertyType       `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType             `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType               `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType       `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                   `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                       `xml:"id,attr,omitempty"`
	Remarks              *string                      `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     string                       `xml:"http://www.opengis.net/gml/3.2 domainOfValidity,omitempty"`
	Component            []TimeOrdinalEraPropertyType `xml:"http://www.opengis.net/gml/3.2 component"`
}

type TimePeriodPropertyType struct {
	TimePeriod   *TimePeriodType `xml:"http://www.opengis.net/gml/3.2 TimePeriod,omitempty"`
	NilReason    string          `xml:"nilReason,attr,omitempty"`
	RemoteSchema string          `xml:"remoteSchema,attr,omitempty"`
	Owns         bool            `xml:"owns,attr,omitempty"`
}

type TimePeriodType struct {
	MetaDataProperty     []MetaDataPropertyType   `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType         `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType           `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType   `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType               `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                   `xml:"id,attr,omitempty"`
	RelatedTime          []RelatedTimeType        `xml:"http://www.opengis.net/gml/3.2 relatedTime"`
	Frame                string                   `xml:"frame,attr,omitempty"`
	BeginPosition        *TimePositionType        `xml:"http://www.opengis.net/gml/3.2 beginPosition,omitempty"`
	Begin                *TimeInstantPropertyType `xml:"http://www.opengis.net/gml/3.2 begin,omitempty"`
	EndPosition          *TimePositionType        `xml:"http://www.opengis.net/gml/3.2 endPosition,omitempty"`
	End                  *TimeInstantPropertyType `xml:"http://www.opengis.net/gml/3.2 end,omitempty"`
	Duration             *string                  `xml:"http://www.opengis.net/gml/3.2 duration,omitempty"`
	TimeInterval         *TimeIntervalLengthType  `xml:"http://www.opengis.net/gml/3.2 timeInterval,omitempty"`
}

type TimePositionType struct {
	Value                 string `xml:",chardata"`
	Frame                 string `xml:"frame,attr,omitempty"`
	CalendarEraName       string `xml:"calendarEraName,attr,omitempty"`
	IndeterminatePosition string `xml:"indeterminatePosition,attr,omitempty"`
}

type TimePrimitivePropertyType struct {
	AbstractTimePrimitive          *AbstractTimePrimitiveType          `xml:"http://www.opengis.net/gml/3.2 AbstractTimePrimitive,omitempty"`
	AbstractTimeGeometricPrimitive *AbstractTimeGeometricPrimitiveType `xml:"http://www.opengis.net/gml/3.2 AbstractTimeGeometricPrimitive,omitempty"`
	AbstractTimeTopologyPrimitive  *AbstractTimeTopologyPrimitiveType  `xml:"http://www.opengis.net/gml/3.2 AbstractTimeTopologyPrimitive,omitempty"`
	NilReason                      string                              `xml:"nilReason,attr,omitempty"`
	RemoteSchema                   string                              `xml:"remoteSchema,attr,omitempty"`
	Owns                           bool                                `xml:"owns,attr,omitempty"`
}

type TimeReferenceSystemType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	Remarks              *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     string                 `xml:"http://www.opengis.net/gml/3.2 domainOfValidity,omitempty"`
}

type TimeTopologyComplexPropertyType struct {
	TimeTopologyComplex *TimeTopologyComplexType `xml:"http://www.opengis.net/gml/3.2 TimeTopologyComplex,omitempty"`
	NilReason           string                   `xml:"nilReason,attr,omitempty"`
	RemoteSchema        string                   `xml:"remoteSchema,attr,omitempty"`
	Owns                bool                     `xml:"owns,attr,omitempty"`
}

type TimeTopologyComplexType struct {
	MetaDataProperty     []MetaDataPropertyType              `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType                    `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                      `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType              `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                          `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                              `xml:"id,attr,omitempty"`
	Primitive            []TimeTopologyPrimitivePropertyType `xml:"http://www.opengis.net/gml/3.2 primitive"`
}

type TimeTopologyPrimitivePropertyType struct {
	AbstractTimeTopologyPrimitive *AbstractTimeTopologyPrimitiveType `xml:"http://www.opengis.net/gml/3.2 AbstractTimeTopologyPrimitive,omitempty"`
	TimeEdge                      *TimeEdgeType                      `xml:"http://www.opengis.net/gml/3.2 TimeEdge,omitempty"`
	TimeNode                      *TimeNodeType                      `xml:"http://www.opengis.net/gml/3.2 TimeNode,omitempty"`
	NilReason                     string                             `xml:"nilReason,attr,omitempty"`
	RemoteSchema                  string                             `xml:"remoteSchema,attr,omitempty"`
	Owns                          bool                               `xml:"owns,attr,omitempty"`
}

type TimeType struct {
	Value string `xml:",chardata"`
}

type TinType struct {
	MetaDataProperty     []MetaDataPropertyType               `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType                     `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                       `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType               `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                           `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                               `xml:"id,attr,omitempty"`
	SrsName              string                               `xml:"srsName,attr,omitempty"`
	SrsDimension         int                                  `xml:"srsDimension,attr,omitempty"`
	AxisLabels           string                               `xml:"axisLabels,attr,omitempty"`
	UomLabels            string                               `xml:"uomLabels,attr,omitempty"`
	Patches              *SurfacePatchArrayPropertyType       `xml:"http://www.opengis.net/gml/3.2 patches,omitempty"`
	PolygonPatches       *SurfacePatchArrayPropertyType       `xml:"http://www.opengis.net/gml/3.2 polygonPatches,omitempty"`
	TrianglePatches      *SurfacePatchArrayPropertyType       `xml:"http://www.opengis.net/gml/3.2 trianglePatches,omitempty"`
	StopLines            []LineStringSegmentArrayPropertyType `xml:"http://www.opengis.net/gml/3.2 stopLines"`
	BreakLines           []LineStringSegmentArrayPropertyType `xml:"http://www.opengis.net/gml/3.2 breakLines"`
	MaxLength            *LengthType                          `xml:"http://www.opengis.net/gml/3.2 maxLength,omitempty"`
	ControlPoint         string                               `xml:"http://www.opengis.net/gml/3.2 controlPoint,omitempty"`
}

type TopoComplexPropertyType struct {
	TopoComplex  *TopoComplexType `xml:"http://www.opengis.net/gml/3.2 TopoComplex,omitempty"`
	NilReason    string           `xml:"nilReason,attr,omitempty"`
	RemoteSchema string           `xml:"remoteSchema,attr,omitempty"`
}

type TopoComplexType struct {
	MetaDataProperty     []MetaDataPropertyType             `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType                   `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                     `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType             `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                         `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                             `xml:"id,attr,omitempty"`
	MaximalComplex       *TopoComplexPropertyType           `xml:"http://www.opengis.net/gml/3.2 maximalComplex,omitempty"`
	SuperComplex         []TopoComplexPropertyType          `xml:"http://www.opengis.net/gml/3.2 superComplex"`
	SubComplex           []TopoComplexPropertyType          `xml:"http://www.opengis.net/gml/3.2 subComplex"`
	TopoPrimitiveMember  []TopoPrimitiveMemberType          `xml:"http://www.opengis.net/gml/3.2 topoPrimitiveMember"`
	TopoPrimitiveMembers *TopoPrimitiveArrayAssociationType `xml:"http://www.opengis.net/gml/3.2 topoPrimitiveMembers,omitempty"`
	IsMaximal            bool                               `xml:"isMaximal,attr,omitempty"`
	AggregationType      string                             `xml:"aggregationType,attr,omitempty"`
}

type TopoCurvePropertyType struct {
	TopoCurve *TopoCurveType `xml:"http://www.opengis.net/gml/3.2 TopoCurve,omitempty"`
	Owns      bool           `xml:"owns,attr,omitempty"`
}

type TopoCurveType struct {
	MetaDataProperty     []MetaDataPropertyType     `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType           `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType             `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType     `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                 `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                     `xml:"id,attr,omitempty"`
	DirectedEdge         []DirectedEdgePropertyType `xml:"http://www.opengis.net/gml/3.2 directedEdge"`
	AggregationType      string                     `xml:"aggregationType,attr,omitempty"`
}

type TopoPointPropertyType struct {
	TopoPoint *TopoPointType `xml:"http://www.opengis.net/gml/3.2 TopoPoint,omitempty"`
	Owns      bool           `xml:"owns,attr,omitempty"`
}

type TopoPointType struct {
	MetaDataProperty     []MetaDataPropertyType    `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType          `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType            `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType    `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                    `xml:"id,attr,omitempty"`
	DirectedNode         *DirectedNodePropertyType `xml:"http://www.opengis.net/gml/3.2 directedNode,omitempty"`
}

type TopoPrimitiveArrayAssociationType struct {
	AbstractTopoPrimitive *AbstractTopoPrimitiveType `xml:"http://www.opengis.net/gml/3.2 AbstractTopoPrimitive,omitempty"`
	Edge                  *EdgeType                  `xml:"http://www.opengis.net/gml/3.2 Edge,omitempty"`
	Face                  *FaceType                  `xml:"http://www.opengis.net/gml/3.2 Face,omitempty"`
	Node                  *NodeType                  `xml:"http://www.opengis.net/gml/3.2 Node,omitempty"`
	TopoSolid             *TopoSolidType             `xml:"http://www.opengis.net/gml/3.2 TopoSolid,omitempty"`
	Owns                  bool                       `xml:"owns,attr,omitempty"`
}

type TopoPrimitiveMemberType struct {
	AbstractTopoPrimitive *AbstractTopoPrimitiveType `xml:"http://www.opengis.net/gml/3.2 AbstractTopoPrimitive,omitempty"`
	Edge                  *EdgeType                  `xml:"http://www.opengis.net/gml/3.2 Edge,omitempty"`
	Face                  *FaceType                  `xml:"http://www.opengis.net/gml/3.2 Face,omitempty"`
	Node                  *NodeType                  `xml:"http://www.opengis.net/gml/3.2 Node,omitempty"`
	TopoSolid             *TopoSolidType             `xml:"http://www.opengis.net/gml/3.2 TopoSolid,omitempty"`
	NilReason             string                     `xml:"nilReason,attr,omitempty"`
	RemoteSchema          string                     `xml:"remoteSchema,attr,omitempty"`
	Owns                  bool                       `xml:"owns,attr,omitempty"`
}

type TopoSolidPropertyType struct {
	TopoSolid    *TopoSolidType `xml:"http://www.opengis.net/gml/3.2 TopoSolid,omitempty"`
	NilReason    string         `xml:"nilReason,attr,omitempty"`
	RemoteSchema string         `xml:"remoteSchema,attr,omitempty"`
	Owns         bool           `xml:"owns,attr,omitempty"`
}

type TopoSolidType struct {
	MetaDataProperty     []MetaDataPropertyType     `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType           `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType             `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType     `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                 `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                     `xml:"id,attr,omitempty"`
	Isolated             []NodeOrEdgePropertyType   `xml:"http://www.opengis.net/gml/3.2 isolated"`
	DirectedFace         []DirectedFacePropertyType `xml:"http://www.opengis.net/gml/3.2 directedFace"`
	SolidProperty        *SolidPropertyType         `xml:"http://www.opengis.net/gml/3.2 solidProperty,omitempty"`
	Universal            bool                       `xml:"universal,attr,omitempty"`
	AggregationType      string                     `xml:"aggregationType,attr,omitempty"`
}

type TopoSurfacePropertyType struct {
	TopoSurface *TopoSurfaceType `xml:"http://www.opengis.net/gml/3.2 TopoSurface,omitempty"`
	Owns        bool             `xml:"owns,attr,omitempty"`
}

type TopoSurfaceType struct {
	MetaDataProperty     []MetaDataPropertyType     `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType           `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType             `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType     `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                 `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                     `xml:"id,attr,omitempty"`
	DirectedFace         []DirectedFacePropertyType `xml:"http://www.opengis.net/gml/3.2 directedFace"`
	AggregationType      string                     `xml:"aggregationType,attr,omitempty"`
}

type TopoVolumePropertyType struct {
	TopoVolume *TopoVolumeType `xml:"http://www.opengis.net/gml/3.2 TopoVolume,omitempty"`
	Owns       bool            `xml:"owns,attr,omitempty"`
}

type TopoVolumeType struct {
	MetaDataProperty     []MetaDataPropertyType          `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType                `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                  `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType          `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                      `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                          `xml:"id,attr,omitempty"`
	DirectedTopoSolid    []DirectedTopoSolidPropertyType `xml:"http://www.opengis.net/gml/3.2 directedTopoSolid"`
	AggregationType      string                          `xml:"aggregationType,attr,omitempty"`
}

type TransformationPropertyType struct {
	Transformation *TransformationType `xml:"http://www.opengis.net/gml/3.2 Transformation,omitempty"`
	NilReason      string              `xml:"nilReason,attr,omitempty"`
	RemoteSchema   string              `xml:"remoteSchema,attr,omitempty"`
}

type TransformationType struct {
	MetaDataProperty            []MetaDataPropertyType                      `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description                 *StringOrRefType                            `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference        *ReferenceType                              `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier                  *CodeWithAuthorityType                      `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                        []CodeType                                  `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                          string                                      `xml:"id,attr,omitempty"`
	Remarks                     *string                                     `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity            *string                                     `xml:"http://www.opengis.net/gml/3.2 domainOfValidity,omitempty"`
	Scope                       []string                                    `xml:"http://www.opengis.net/gml/3.2 scope"`
	OperationVersion            *string                                     `xml:"http://www.opengis.net/gml/3.2 operationVersion,omitempty"`
	CoordinateOperationAccuracy []string                                    `xml:"http://www.opengis.net/gml/3.2 coordinateOperationAccuracy"`
	SourceCRS                   *CRSPropertyType                            `xml:"http://www.opengis.net/gml/3.2 sourceCRS,omitempty"`
	TargetCRS                   *CRSPropertyType                            `xml:"http://www.opengis.net/gml/3.2 targetCRS,omitempty"`
	Method                      *OperationMethodPropertyType                `xml:"http://www.opengis.net/gml/3.2 method,omitempty"`
	UsesMethod                  *OperationMethodPropertyType                `xml:"http://www.opengis.net/gml/3.2 usesMethod,omitempty"`
	ParameterValue              []AbstractGeneralParameterValuePropertyType `xml:"http://www.opengis.net/gml/3.2 parameterValue"`
	IncludesValue               *AbstractGeneralParameterValuePropertyType  `xml:"http://www.opengis.net/gml/3.2 includesValue,omitempty"`
	UsesValue                   *AbstractGeneralParameterValuePropertyType  `xml:"http://www.opengis.net/gml/3.2 usesValue,omitempty"`
}

type TriangleType struct {
	Exterior      *AbstractRingPropertyType `xml:"http://www.opengis.net/gml/3.2 exterior,omitempty"`
	Interpolation string                    `xml:"interpolation,attr,omitempty"`
}

type UnitDefinitionType struct {
	MetaDataProperty      []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description           *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference  *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier            *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                  []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                    string                 `xml:"id,attr,omitempty"`
	Remarks               *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	QuantityType          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 quantityType,omitempty"`
	QuantityTypeReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 quantityTypeReference,omitempty"`
	CatalogSymbol         *CodeType              `xml:"http://www.opengis.net/gml/3.2 catalogSymbol,omitempty"`
}

type UnitOfMeasureType struct {
	Uom string `xml:"uom,attr"`
}

type UserDefinedCSPropertyType struct {
	UserDefinedCS *UserDefinedCSType `xml:"http://www.opengis.net/gml/3.2 UserDefinedCS,omitempty"`
	NilReason     string             `xml:"nilReason,attr,omitempty"`
	RemoteSchema  string             `xml:"remoteSchema,attr,omitempty"`
}

type UserDefinedCSType struct {
	MetaDataProperty     []MetaDataPropertyType             `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType                   `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                     `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType             `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                         `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                             `xml:"id,attr,omitempty"`
	Remarks              *string                            `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	Axis                 []CoordinateSystemAxisPropertyType `xml:"http://www.opengis.net/gml/3.2 axis"`
	UsesAxis             *CoordinateSystemAxisPropertyType  `xml:"http://www.opengis.net/gml/3.2 usesAxis,omitempty"`
	AggregationType      string                             `xml:"aggregationType,attr,omitempty"`
}

type ValueArrayPropertyType struct {
	AbstractValue              *string                         `xml:"http://www.opengis.net/gml/3.2 AbstractValue,omitempty"`
	AbstractScalarValue        *string                         `xml:"http://www.opengis.net/gml/3.2 AbstractScalarValue,omitempty"`
	AbstractScalarValueList    *string                         `xml:"http://www.opengis.net/gml/3.2 AbstractScalarValueList,omitempty"`
	CategoryExtent             *CategoryExtentType             `xml:"http://www.opengis.net/gml/3.2 CategoryExtent,omitempty"`
	CompositeValue             *CompositeValueType             `xml:"http://www.opengis.net/gml/3.2 CompositeValue,omitempty"`
	CountExtent                *string                         `xml:"http://www.opengis.net/gml/3.2 CountExtent,omitempty"`
	QuantityExtent             *QuantityExtentType             `xml:"http://www.opengis.net/gml/3.2 QuantityExtent,omitempty"`
	AbstractGeometry           *AbstractGeometryType           `xml:"http://www.opengis.net/gml/3.2 AbstractGeometry,omitempty"`
	AbstractGeometricAggregate *AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml/3.2 AbstractGeometricAggregate,omitempty"`
	AbstractGeometricPrimitive *AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml/3.2 AbstractGeometricPrimitive,omitempty"`
	AbstractImplicitGeometry   *AbstractGeometryType           `xml:"http://www.opengis.net/gml/3.2 AbstractImplicitGeometry,omitempty"`
	GeometricComplex           *GeometricComplexType           `xml:"http://www.opengis.net/gml/3.2 GeometricComplex,omitempty"`
	AbstractTimeObject         *AbstractTimeObjectType         `xml:"http://www.opengis.net/gml/3.2 AbstractTimeObject,omitempty"`
	AbstractTimeComplex        *AbstractTimeComplexType        `xml:"http://www.opengis.net/gml/3.2 AbstractTimeComplex,omitempty"`
	AbstractTimePrimitive      *AbstractTimePrimitiveType      `xml:"http://www.opengis.net/gml/3.2 AbstractTimePrimitive,omitempty"`
	Null                       *string                         `xml:"http://www.opengis.net/gml/3.2 Null,omitempty"`
	Owns                       bool                            `xml:"owns,attr,omitempty"`
}

type ValueArrayType struct {
	MetaDataProperty     []MetaDataPropertyType  `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType        `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType          `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType  `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType              `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                  `xml:"id,attr,omitempty"`
	ValueComponent       []ValuePropertyType     `xml:"http://www.opengis.net/gml/3.2 valueComponent"`
	ValueComponents      *ValueArrayPropertyType `xml:"http://www.opengis.net/gml/3.2 valueComponents,omitempty"`
	AggregationType      string                  `xml:"aggregationType,attr,omitempty"`
	CodeSpace            string                  `xml:"codeSpace,attr,omitempty"`
	Uom                  string                  `xml:"uom,attr,omitempty"`
}

type ValuePropertyType struct {
	AbstractValue              *string                         `xml:"http://www.opengis.net/gml/3.2 AbstractValue,omitempty"`
	AbstractScalarValue        *string                         `xml:"http://www.opengis.net/gml/3.2 AbstractScalarValue,omitempty"`
	AbstractScalarValueList    *string                         `xml:"http://www.opengis.net/gml/3.2 AbstractScalarValueList,omitempty"`
	CategoryExtent             *CategoryExtentType             `xml:"http://www.opengis.net/gml/3.2 CategoryExtent,omitempty"`
	CompositeValue             *CompositeValueType             `xml:"http://www.opengis.net/gml/3.2 CompositeValue,omitempty"`
	CountExtent                *string                         `xml:"http://www.opengis.net/gml/3.2 CountExtent,omitempty"`
	QuantityExtent             *QuantityExtentType             `xml:"http://www.opengis.net/gml/3.2 QuantityExtent,omitempty"`
	AbstractGeometry           *AbstractGeometryType           `xml:"http://www.opengis.net/gml/3.2 AbstractGeometry,omitempty"`
	AbstractGeometricAggregate *AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml/3.2 AbstractGeometricAggregate,omitempty"`
	AbstractGeometricPrimitive *AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml/3.2 AbstractGeometricPrimitive,omitempty"`
	AbstractImplicitGeometry   *AbstractGeometryType           `xml:"http://www.opengis.net/gml/3.2 AbstractImplicitGeometry,omitempty"`
	GeometricComplex           *GeometricComplexType           `xml:"http://www.opengis.net/gml/3.2 GeometricComplex,omitempty"`
	AbstractTimeObject         *AbstractTimeObjectType         `xml:"http://www.opengis.net/gml/3.2 AbstractTimeObject,omitempty"`
	AbstractTimeComplex        *AbstractTimeComplexType        `xml:"http://www.opengis.net/gml/3.2 AbstractTimeComplex,omitempty"`
	AbstractTimePrimitive      *AbstractTimePrimitiveType      `xml:"http://www.opengis.net/gml/3.2 AbstractTimePrimitive,omitempty"`
	Null                       *string                         `xml:"http://www.opengis.net/gml/3.2 Null,omitempty"`
	NilReason                  string                          `xml:"nilReason,attr,omitempty"`
	RemoteSchema               string                          `xml:"remoteSchema,attr,omitempty"`
	Owns                       bool                            `xml:"owns,attr,omitempty"`
}

type VectorType struct {
	Value string `xml:",chardata"`
}

type VerticalCRSPropertyType struct {
	VerticalCRS  *VerticalCRSType `xml:"http://www.opengis.net/gml/3.2 VerticalCRS,omitempty"`
	NilReason    string           `xml:"nilReason,attr,omitempty"`
	RemoteSchema string           `xml:"remoteSchema,attr,omitempty"`
}

type VerticalCRSType struct {
	MetaDataProperty     []MetaDataPropertyType     `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType           `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType             `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType     `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                 `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                     `xml:"id,attr,omitempty"`
	Remarks              *string                    `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     []string                   `xml:"http://www.opengis.net/gml/3.2 domainOfValidity"`
	Scope                []string                   `xml:"http://www.opengis.net/gml/3.2 scope"`
	VerticalCS           *VerticalCSPropertyType    `xml:"http://www.opengis.net/gml/3.2 verticalCS,omitempty"`
	UsesVerticalCS       *VerticalCSPropertyType    `xml:"http://www.opengis.net/gml/3.2 usesVerticalCS,omitempty"`
	VerticalDatum        *VerticalDatumPropertyType `xml:"http://www.opengis.net/gml/3.2 verticalDatum,omitempty"`
	UsesVerticalDatum    *VerticalDatumPropertyType `xml:"http://www.opengis.net/gml/3.2 usesVerticalDatum,omitempty"`
}

type VerticalCSPropertyType struct {
	VerticalCS   *VerticalCSType `xml:"http://www.opengis.net/gml/3.2 VerticalCS,omitempty"`
	NilReason    string          `xml:"nilReason,attr,omitempty"`
	RemoteSchema string          `xml:"remoteSchema,attr,omitempty"`
}

type VerticalCSType struct {
	MetaDataProperty     []MetaDataPropertyType             `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType                   `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType                     `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType             `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType                         `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                             `xml:"id,attr,omitempty"`
	Remarks              *string                            `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	Axis                 []CoordinateSystemAxisPropertyType `xml:"http://www.opengis.net/gml/3.2 axis"`
	UsesAxis             *CoordinateSystemAxisPropertyType  `xml:"http://www.opengis.net/gml/3.2 usesAxis,omitempty"`
	AggregationType      string                             `xml:"aggregationType,attr,omitempty"`
}

type VerticalDatumPropertyType struct {
	VerticalDatum *VerticalDatumType `xml:"http://www.opengis.net/gml/3.2 VerticalDatum,omitempty"`
	NilReason     string             `xml:"nilReason,attr,omitempty"`
	RemoteSchema  string             `xml:"remoteSchema,attr,omitempty"`
}

type VerticalDatumType struct {
	MetaDataProperty     []MetaDataPropertyType `xml:"http://www.opengis.net/gml/3.2 metaDataProperty"`
	Description          *StringOrRefType       `xml:"http://www.opengis.net/gml/3.2 description,omitempty"`
	DescriptionReference *ReferenceType         `xml:"http://www.opengis.net/gml/3.2 descriptionReference,omitempty"`
	Identifier           *CodeWithAuthorityType `xml:"http://www.opengis.net/gml/3.2 identifier,omitempty"`
	Name                 []CodeType             `xml:"http://www.opengis.net/gml/3.2 name"`
	Id                   string                 `xml:"id,attr,omitempty"`
	Remarks              *string                `xml:"http://www.opengis.net/gml/3.2 remarks,omitempty"`
	DomainOfValidity     *string                `xml:"http://www.opengis.net/gml/3.2 domainOfValidity,omitempty"`
	Scope                []string               `xml:"http://www.opengis.net/gml/3.2 scope"`
	AnchorDefinition     *CodeType              `xml:"http://www.opengis.net/gml/3.2 anchorDefinition,omitempty"`
	AnchorPoint          *CodeType              `xml:"http://www.opengis.net/gml/3.2 anchorPoint,omitempty"`
	RealizationEpoch     *string                `xml:"http://www.opengis.net/gml/3.2 realizationEpoch,omitempty"`
}

type VolumeType struct {
	Value string `xml:",chardata"`
}
