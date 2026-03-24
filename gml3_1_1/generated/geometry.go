package gml

type AbsoluteExternalPositionalAccuracyType struct {
	// A description of the position accuracy parameter(s) provided.
	MeasureDescription *CodeType `xml:"http://www.opengis.net/gml measureDescription,omitempty"`
	// A quantitative result defined by the evaluation procedure used, and identified by the measureDescription.
	Result *MeasureType `xml:"http://www.opengis.net/gml result,omitempty"`
}

type AbstractContinuousCoverageType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType         `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType         `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType         `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType         `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType         `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType         `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType         `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType         `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType         `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType         `xml:"http://www.opengis.net/gml srsName"`
	Id                      string             `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	BoundedBy               *BoundingShapeType `xml:"http://www.opengis.net/gml boundedBy,omitempty"`
	// deprecated in GML version 3.1
	Location            *LocationPropertyType         `xml:"http://www.opengis.net/gml location,omitempty"`
	PriorityLocation    *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml priorityLocation,omitempty"`
	DomainSet           *DomainSetType                `xml:"http://www.opengis.net/gml domainSet,omitempty"`
	GridDomain          *GridDomainType               `xml:"http://www.opengis.net/gml gridDomain,omitempty"`
	MultiCurveDomain    *MultiCurveDomainType         `xml:"http://www.opengis.net/gml multiCurveDomain,omitempty"`
	MultiPointDomain    *MultiPointDomainType         `xml:"http://www.opengis.net/gml multiPointDomain,omitempty"`
	MultiSolidDomain    *MultiSolidDomainType         `xml:"http://www.opengis.net/gml multiSolidDomain,omitempty"`
	MultiSurfaceDomain  *MultiSurfaceDomainType       `xml:"http://www.opengis.net/gml multiSurfaceDomain,omitempty"`
	RectifiedGridDomain *RectifiedGridDomainType      `xml:"http://www.opengis.net/gml rectifiedGridDomain,omitempty"`
	RangeSet            *RangeSetType                 `xml:"http://www.opengis.net/gml rangeSet,omitempty"`
	Dimension           *int                          `xml:"dimension,attr,omitempty"`
	CoverageFunction    *CoverageFunctionType         `xml:"http://www.opengis.net/gml coverageFunction,omitempty"`
}

type AbstractCoordinateOperationBaseType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
}

type AbstractCoordinateOperationType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this coordinate operation. The first coordinateOperationID, if any, is normally the primary identification code, and any others are aliases.
	CoordinateOperationID []IdentifierType `xml:"http://www.opengis.net/gml coordinateOperationID"`
	// Comments on or information about this coordinate operation, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Version of the coordinate transformation (i.e., instantiation due to the stochastic nature of the parameters). Mandatory when describing a transformation, and should not be supplied for a conversion.
	OperationVersion *string `xml:"http://www.opengis.net/gml operationVersion,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
	// Unordered set of estimates of the impact of this coordinate operation on point position accuracy. Gives position error estimates for target coordinates of this coordinate operation, assuming no errors in source coordinates.
	AbstractPositionalAccuracy         []AbstractPositionalAccuracyType         `xml:"http://www.opengis.net/gml _positionalAccuracy"`
	AbsoluteExternalPositionalAccuracy []AbsoluteExternalPositionalAccuracyType `xml:"http://www.opengis.net/gml absoluteExternalPositionalAccuracy"`
	CovarianceMatrix                   []CovarianceMatrixType                   `xml:"http://www.opengis.net/gml covarianceMatrix"`
	RelativeInternalPositionalAccuracy []RelativeInternalPositionalAccuracyType `xml:"http://www.opengis.net/gml relativeInternalPositionalAccuracy"`
	// Association to the source CRS (coordinate reference system) of this coordinate operation.
	SourceCRS *CRSRefType `xml:"http://www.opengis.net/gml sourceCRS,omitempty"`
	// Association to the target CRS (coordinate reference system) of this coordinate operation. For constraints on multiplicity of "sourceCRS" and "targetCRS", see UML model of Coordinate Operation package in OGC Abstract Specification topic 2.
	TargetCRS *CRSRefType `xml:"http://www.opengis.net/gml targetCRS,omitempty"`
}

type AbstractCoordinateSystemBaseType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
}

type AbstractCoordinateSystemType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this coordinate system. The first csID, if any, is normally the primary identification code, and any others are aliases.
	CsID []IdentifierType `xml:"http://www.opengis.net/gml csID"`
	// Comments on or information about this coordinate system, including data source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Ordered sequence of associations to the coordinate system axes included in this coordinate system.
	UsesAxis []CoordinateSystemAxisRefType `xml:"http://www.opengis.net/gml usesAxis"`
}

type AbstractCoverageType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType         `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType         `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType         `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType         `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType         `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType         `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType         `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType         `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType         `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType         `xml:"http://www.opengis.net/gml srsName"`
	Id                      string             `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	BoundedBy               *BoundingShapeType `xml:"http://www.opengis.net/gml boundedBy,omitempty"`
	// deprecated in GML version 3.1
	Location            *LocationPropertyType         `xml:"http://www.opengis.net/gml location,omitempty"`
	PriorityLocation    *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml priorityLocation,omitempty"`
	DomainSet           *DomainSetType                `xml:"http://www.opengis.net/gml domainSet,omitempty"`
	GridDomain          *GridDomainType               `xml:"http://www.opengis.net/gml gridDomain,omitempty"`
	MultiCurveDomain    *MultiCurveDomainType         `xml:"http://www.opengis.net/gml multiCurveDomain,omitempty"`
	MultiPointDomain    *MultiPointDomainType         `xml:"http://www.opengis.net/gml multiPointDomain,omitempty"`
	MultiSolidDomain    *MultiSolidDomainType         `xml:"http://www.opengis.net/gml multiSolidDomain,omitempty"`
	MultiSurfaceDomain  *MultiSurfaceDomainType       `xml:"http://www.opengis.net/gml multiSurfaceDomain,omitempty"`
	RectifiedGridDomain *RectifiedGridDomainType      `xml:"http://www.opengis.net/gml rectifiedGridDomain,omitempty"`
	RangeSet            *RangeSetType                 `xml:"http://www.opengis.net/gml rangeSet,omitempty"`
	Dimension           *int                          `xml:"dimension,attr,omitempty"`
}

type AbstractCurveSegmentType struct {
	// The attribute "numDerivativesAtStart" specifies the type of continuity between this curve segment and its predecessor. If this is the first curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtStart *int `xml:"numDerivativesAtStart,attr,omitempty"`
	// The attribute "numDerivativesAtEnd" specifies the type of continuity between this curve segment and its successor. If this is the last curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtEnd *int `xml:"numDerivativesAtEnd,attr,omitempty"`
	// The attribute "numDerivativesInterior" specifies the type of continuity that is guaranteed interior to the curve. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativeInterior *int `xml:"numDerivativeInterior,attr,omitempty"`
}

type AbstractCurveType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
}

type AbstractDatumBaseType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
}

type AbstractDatumType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this datum. The first datumID, if any, is normally the primary identification code, and any others are aliases.
	DatumID []IdentifierType `xml:"http://www.opengis.net/gml datumID"`
	// Comments on this reference system, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Description, possibly including coordinates, of the point or points used to anchor the datum to the Earth. Also known as the "origin", especially for engineering and image datums. The codeSpace attribute can be used to reference a source of more detailed on this point or surface, or on a set of such descriptions.
	// - For a geodetic datum, this point is also known as the fundamental point, which is traditionally the point where the relationship between geoid and ellipsoid is defined. In some cases, the "fundamental point" may consist of a number of points. In those cases, the parameters defining the geoid/ellipsoid relationship have been averaged for these points, and the averages adopted as the datum definition.
	// - For an engineering datum, the anchor point may be a physical point, or it may be a point with defined coordinates in another CRS. When appropriate, the coordinates of this anchor point can be referenced in another document, such as referencing a GML feature that references or includes a point position.
	// - For an image datum, the anchor point is usually either the centre of the image or the corner of the image.
	// - For a temporal datum, this attribute is not defined. Instead of the anchor point, a temporal datum carries a separate time origin of type DateTime.
	AnchorPoint *CodeType `xml:"http://www.opengis.net/gml anchorPoint,omitempty"`
	// The time after which this datum definition is valid. This time may be precise (e.g. 1997.0 for IRTF97) or merely a year (e.g. 1983 for NAD83). In the latter case, the epoch usually refers to the year in which a major recalculation of the geodetic control network, underlying the datum, was executed or initiated. An old datum can remain valid after a new datum is defined. Alternatively, a datum may be superseded by a later datum, in which case the realization epoch for the new datum defines the upper limit for the validity of the superseded datum.
	RealizationEpoch *string `xml:"http://www.opengis.net/gml realizationEpoch,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
}

type AbstractDiscreteCoverageType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType         `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType         `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType         `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType         `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType         `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType         `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType         `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType         `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType         `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType         `xml:"http://www.opengis.net/gml srsName"`
	Id                      string             `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	BoundedBy               *BoundingShapeType `xml:"http://www.opengis.net/gml boundedBy,omitempty"`
	// deprecated in GML version 3.1
	Location            *LocationPropertyType         `xml:"http://www.opengis.net/gml location,omitempty"`
	PriorityLocation    *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml priorityLocation,omitempty"`
	DomainSet           *DomainSetType                `xml:"http://www.opengis.net/gml domainSet,omitempty"`
	GridDomain          *GridDomainType               `xml:"http://www.opengis.net/gml gridDomain,omitempty"`
	MultiCurveDomain    *MultiCurveDomainType         `xml:"http://www.opengis.net/gml multiCurveDomain,omitempty"`
	MultiPointDomain    *MultiPointDomainType         `xml:"http://www.opengis.net/gml multiPointDomain,omitempty"`
	MultiSolidDomain    *MultiSolidDomainType         `xml:"http://www.opengis.net/gml multiSolidDomain,omitempty"`
	MultiSurfaceDomain  *MultiSurfaceDomainType       `xml:"http://www.opengis.net/gml multiSurfaceDomain,omitempty"`
	RectifiedGridDomain *RectifiedGridDomainType      `xml:"http://www.opengis.net/gml rectifiedGridDomain,omitempty"`
	RangeSet            *RangeSetType                 `xml:"http://www.opengis.net/gml rangeSet,omitempty"`
	Dimension           *int                          `xml:"dimension,attr,omitempty"`
	CoverageFunction    *CoverageFunctionType         `xml:"http://www.opengis.net/gml coverageFunction,omitempty"`
}

type AbstractFeatureCollectionType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType         `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType         `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType         `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType         `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType         `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType         `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType         `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType         `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType         `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType         `xml:"http://www.opengis.net/gml srsName"`
	Id                      string             `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	BoundedBy               *BoundingShapeType `xml:"http://www.opengis.net/gml boundedBy,omitempty"`
	// deprecated in GML version 3.1
	Location         *LocationPropertyType         `xml:"http://www.opengis.net/gml location,omitempty"`
	PriorityLocation *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml priorityLocation,omitempty"`
	FeatureMember    []FeaturePropertyType         `xml:"http://www.opengis.net/gml featureMember"`
	FeatureMembers   *FeatureArrayPropertyType     `xml:"http://www.opengis.net/gml featureMembers,omitempty"`
}

type AbstractFeatureType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType         `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType         `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType         `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType         `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType         `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType         `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType         `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType         `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType         `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType         `xml:"http://www.opengis.net/gml srsName"`
	Id                      string             `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	BoundedBy               *BoundingShapeType `xml:"http://www.opengis.net/gml boundedBy,omitempty"`
	// deprecated in GML version 3.1
	Location         *LocationPropertyType         `xml:"http://www.opengis.net/gml location,omitempty"`
	PriorityLocation *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml priorityLocation,omitempty"`
}

type AbstractGMLType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
}

type AbstractGeneralConversionType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this coordinate operation. The first coordinateOperationID, if any, is normally the primary identification code, and any others are aliases.
	CoordinateOperationID []IdentifierType `xml:"http://www.opengis.net/gml coordinateOperationID"`
	// Comments on or information about this coordinate operation, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Version of the coordinate transformation (i.e., instantiation due to the stochastic nature of the parameters). Mandatory when describing a transformation, and should not be supplied for a conversion.
	OperationVersion *string `xml:"http://www.opengis.net/gml operationVersion,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
	// Unordered set of estimates of the impact of this coordinate operation on point position accuracy. Gives position error estimates for target coordinates of this coordinate operation, assuming no errors in source coordinates.
	AbstractPositionalAccuracy         []AbstractPositionalAccuracyType         `xml:"http://www.opengis.net/gml _positionalAccuracy"`
	AbsoluteExternalPositionalAccuracy []AbsoluteExternalPositionalAccuracyType `xml:"http://www.opengis.net/gml absoluteExternalPositionalAccuracy"`
	CovarianceMatrix                   []CovarianceMatrixType                   `xml:"http://www.opengis.net/gml covarianceMatrix"`
	RelativeInternalPositionalAccuracy []RelativeInternalPositionalAccuracyType `xml:"http://www.opengis.net/gml relativeInternalPositionalAccuracy"`
	// Association to the source CRS (coordinate reference system) of this coordinate operation.
	SourceCRS *CRSRefType `xml:"http://www.opengis.net/gml sourceCRS,omitempty"`
	// Association to the target CRS (coordinate reference system) of this coordinate operation. For constraints on multiplicity of "sourceCRS" and "targetCRS", see UML model of Coordinate Operation package in OGC Abstract Specification topic 2.
	TargetCRS *CRSRefType `xml:"http://www.opengis.net/gml targetCRS,omitempty"`
}

type AbstractGeneralDerivedCRSType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alterative identifications of this reference system. The first srsID, if any, is normally the primary identification code, and any others are aliases.
	SrsID []IdentifierType `xml:"http://www.opengis.net/gml srsID"`
	// Comments on or information about this reference system, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
	// Association to the coordinate reference system used by this derived CRS.
	BaseCRS *CoordinateReferenceSystemRefType `xml:"http://www.opengis.net/gml baseCRS,omitempty"`
	// Association to the coordinate conversion used to define this derived CRS.
	DefinedByConversion *GeneralConversionRefType `xml:"http://www.opengis.net/gml definedByConversion,omitempty"`
}

type AbstractGeneralOperationParameterRefType struct {
	AbstractGeneralOperationParameter *AbstractGeneralOperationParameterType `xml:"http://www.opengis.net/gml _GeneralOperationParameter,omitempty"`
	OperationParameter                *OperationParameterType                `xml:"http://www.opengis.net/gml OperationParameter,omitempty"`
	OperationParameterGroup           *OperationParameterGroupType           `xml:"http://www.opengis.net/gml OperationParameterGroup,omitempty"`
	RemoteSchema                      string                                 `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                         string                                 `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                              string                                 `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                              string                                 `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                           string                                 `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                             string                                 `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                              string                                 `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                           string                                 `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type AbstractGeneralOperationParameterType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// The minimum number of times that values for this parameter group or parameter are required. If this attribute is omitted, the minimum number is one.
	MinimumOccurs *int `xml:"http://www.opengis.net/gml minimumOccurs,omitempty"`
}

type AbstractGeneralParameterValueType struct {
}

type AbstractGeneralTransformationType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this coordinate operation. The first coordinateOperationID, if any, is normally the primary identification code, and any others are aliases.
	CoordinateOperationID []IdentifierType `xml:"http://www.opengis.net/gml coordinateOperationID"`
	// Comments on or information about this coordinate operation, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Version of the coordinate transformation (i.e., instantiation due to the stochastic nature of the parameters). Mandatory when describing a transformation, and should not be supplied for a conversion.
	OperationVersion *string `xml:"http://www.opengis.net/gml operationVersion,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
	// Unordered set of estimates of the impact of this coordinate operation on point position accuracy. Gives position error estimates for target coordinates of this coordinate operation, assuming no errors in source coordinates.
	AbstractPositionalAccuracy         []AbstractPositionalAccuracyType         `xml:"http://www.opengis.net/gml _positionalAccuracy"`
	AbsoluteExternalPositionalAccuracy []AbsoluteExternalPositionalAccuracyType `xml:"http://www.opengis.net/gml absoluteExternalPositionalAccuracy"`
	CovarianceMatrix                   []CovarianceMatrixType                   `xml:"http://www.opengis.net/gml covarianceMatrix"`
	RelativeInternalPositionalAccuracy []RelativeInternalPositionalAccuracyType `xml:"http://www.opengis.net/gml relativeInternalPositionalAccuracy"`
	// Association to the source CRS (coordinate reference system) of this coordinate operation.
	SourceCRS *CRSRefType `xml:"http://www.opengis.net/gml sourceCRS,omitempty"`
	// Association to the target CRS (coordinate reference system) of this coordinate operation. For constraints on multiplicity of "sourceCRS" and "targetCRS", see UML model of Coordinate Operation package in OGC Abstract Specification topic 2.
	TargetCRS *CRSRefType `xml:"http://www.opengis.net/gml targetCRS,omitempty"`
}

type AbstractGeometricAggregateType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
}

type AbstractGeometricPrimitiveType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
}

type AbstractGeometryType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
}

type AbstractGriddedSurfaceType struct {
	// The attribute rows gives the number
	// of rows in the parameter grid.
	Rows int `xml:"http://www.opengis.net/gml rows,omitempty"`
	// The attribute columns gives the number
	// of columns in the parameter grid.
	Columns int      `xml:"http://www.opengis.net/gml columns,omitempty"`
	Row     []string `xml:"http://www.opengis.net/gml row"`
}

type AbstractMetaDataType struct {
	Id string `xml:"http://www.opengis.net/gml id,attr,omitempty"`
}

type AbstractParametricCurveSurfaceType struct {
}

type AbstractPositionalAccuracyType struct {
	// A description of the position accuracy parameter(s) provided.
	MeasureDescription *CodeType `xml:"http://www.opengis.net/gml measureDescription,omitempty"`
}

type AbstractReferenceSystemBaseType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
}

type AbstractReferenceSystemType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alterative identifications of this reference system. The first srsID, if any, is normally the primary identification code, and any others are aliases.
	SrsID []IdentifierType `xml:"http://www.opengis.net/gml srsID"`
	// Comments on or information about this reference system, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
}

type AbstractRingPropertyType struct {
	// The "_Ring" element is the abstract head of the substituition group for all closed boundaries of a surface patch.
	AbstractRing *AbstractRingType `xml:"http://www.opengis.net/gml _Ring,omitempty"`
	LinearRing   *LinearRingType   `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	Ring         *RingType         `xml:"http://www.opengis.net/gml Ring,omitempty"`
}

type AbstractRingType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
}

type AbstractSolidType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
}

type AbstractStyleType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
}

type AbstractSurfacePatchType struct {
}

type AbstractSurfaceType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
}

type AbstractTimeComplexType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
}

type AbstractTimeGeometricPrimitiveType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType        `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType        `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType        `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType        `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType        `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType        `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType        `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType        `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType        `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType        `xml:"http://www.opengis.net/gml srsName"`
	Id                      string            `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	RelatedTime             []RelatedTimeType `xml:"http://www.opengis.net/gml relatedTime"`
	Frame                   *string           `xml:"frame,attr,omitempty"`
}

type AbstractTimeObjectType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
}

type AbstractTimePrimitiveType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType        `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType        `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType        `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType        `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType        `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType        `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType        `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType        `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType        `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType        `xml:"http://www.opengis.net/gml srsName"`
	Id                      string            `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	RelatedTime             []RelatedTimeType `xml:"http://www.opengis.net/gml relatedTime"`
}

type AbstractTimeReferenceSystemType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	DomainOfValidity        string     `xml:"http://www.opengis.net/gml domainOfValidity,omitempty"`
}

type AbstractTimeSliceType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType                 `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType                 `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType                 `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType                 `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType                 `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType                 `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType                 `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType                 `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType                 `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType                 `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	ValidTime               *TimePrimitivePropertyType `xml:"http://www.opengis.net/gml validTime,omitempty"`
	DataSource              *StringOrRefType           `xml:"http://www.opengis.net/gml dataSource,omitempty"`
}

type AbstractTimeTopologyPrimitiveType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType        `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType        `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType        `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType        `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType        `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType        `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType        `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType        `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType        `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType        `xml:"http://www.opengis.net/gml srsName"`
	Id                      string            `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	RelatedTime             []RelatedTimeType `xml:"http://www.opengis.net/gml relatedTime"`
	Complex                 *ReferenceType    `xml:"http://www.opengis.net/gml complex,omitempty"`
}

type AbstractTopoPrimitiveType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType             `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType             `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType             `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType             `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType             `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType             `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType             `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType             `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType             `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType             `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                 `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	Isolated                []IsolatedPropertyType `xml:"http://www.opengis.net/gml isolated"`
	Container               *ContainerPropertyType `xml:"http://www.opengis.net/gml container,omitempty"`
}

type AbstractTopologyType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
}

type AffinePlacementType struct {
	// The location property gives
	// the target of the parameter space origin. This is the vector
	// (x0, y0, z0) in the formulae above.
	Location *DirectPositionType `xml:"http://www.opengis.net/gml location,omitempty"`
	// The attribute refDirection gives the
	// target directions for the co-ordinate basis vectors of the
	// parameter space. These are the columns of the matrix in the
	// formulae given above. The number of directions given shall be
	// inDimension. The dimension of the directions shall be
	// outDimension.
	RefDirection []VectorType `xml:"http://www.opengis.net/gml refDirection"`
	// Dimension of the constructive parameter
	// space.
	InDimension int `xml:"http://www.opengis.net/gml inDimension,omitempty"`
	// Dimension of the co-ordinate space.
	OutDimension int `xml:"http://www.opengis.net/gml outDimension,omitempty"`
}

type AngleChoiceType struct {
	Angle    *MeasureType  `xml:"http://www.opengis.net/gml angle,omitempty"`
	DmsAngle *DMSAngleType `xml:"http://www.opengis.net/gml dmsAngle,omitempty"`
}

type AngleType struct {
	Value string `xml:",chardata"`
}

type ArcByBulgeType struct {
	// The attribute "numDerivativesAtStart" specifies the type of continuity between this curve segment and its predecessor. If this is the first curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtStart *int `xml:"numDerivativesAtStart,attr,omitempty"`
	// The attribute "numDerivativesAtEnd" specifies the type of continuity between this curve segment and its successor. If this is the last curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtEnd *int `xml:"numDerivativesAtEnd,attr,omitempty"`
	// The attribute "numDerivativesInterior" specifies the type of continuity that is guaranteed interior to the curve. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativeInterior *int `xml:"numDerivativeInterior,attr,omitempty"`
	// The bulge controls the offset of each arc's midpoint. The "bulge" is the real number multiplier for the normal that determines the offset direction of the midpoint of each arc. The length of the bulge sequence is exactly 1 less than the length of the control point array, since a bulge is needed for each pair of adjacent points in the control point array. The bulge is not given by a distance, since it is simply a multiplier for the normal.
	// The midpoint of the resulting arc is given by: midPoint = ((startPoint + endPoint)/2.0) + bulge*normal
	Bulge []float64 `xml:"http://www.opengis.net/gml bulge"`
	// The attribute "normal" is a vector normal (perpendicular) to the chord of the arc, the line joining the first and last
	// point of the arc. In a 2D coordinate system, there are only two possible directions for the normal, and it is often given as a signed real, indicating its length, with a positive sign indicating a left turn angle from the chord line, and a negative sign indicating a right turn from the chord. In 3D, the normal determines the plane of the arc, along with the start and endPoint of the arc.
	// The normal is usually a unit vector, but this is not absolutely necessary. If the normal is a zero vector, the geometric object becomes equivalent to the straight line between the two end points. The length of the normal sequence is exactly the same as for the bulge sequence, 1 less than the control point sequence length.
	Normal  []VectorType            `xml:"http://www.opengis.net/gml normal"`
	PosList *DirectPositionListType `xml:"http://www.opengis.net/gml posList,omitempty"`
	// Deprecated with GML version 3.1.0. Use "posList" instead.
	Coordinates *CoordinatesType     `xml:"http://www.opengis.net/gml coordinates,omitempty"`
	Pos         []DirectPositionType `xml:"http://www.opengis.net/gml pos"`
	// This property element either references a point via the XLink-attributes or contains the point element. pointProperty
	// is the predefined property which can be used by GML Application Schemas whenever a GML Feature has a property with a value that
	// is substitutable for Point.
	PointProperty []PointPropertyType `xml:"http://www.opengis.net/gml pointProperty"`
	// Deprecated with GML version 3.1.0. Use "pointProperty" instead. Included for backwards compatibility with GML 3.0.0.
	PointRep []PointPropertyType `xml:"http://www.opengis.net/gml pointRep"`
	// The attribute "interpolation" specifies the curve interpolation mechanism used for this segment. This mechanism
	// uses the control points and control parameters to determine the position of this curve segment. For an ArcStringByBulge the interpolation is fixed as "circularArc2PointWithBulge".
	Interpolation *string `xml:"interpolation,attr,omitempty"`
	// The number of arcs in the arc string can be explicitly stated in this attribute. The number of control points in the arc string must be numArc + 1.
	NumArc *int `xml:"numArc,attr,omitempty"`
}

type ArcByCenterPointType struct {
	// The attribute "numDerivativesAtStart" specifies the type of continuity between this curve segment and its predecessor. If this is the first curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtStart *int `xml:"numDerivativesAtStart,attr,omitempty"`
	// The attribute "numDerivativesAtEnd" specifies the type of continuity between this curve segment and its successor. If this is the last curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtEnd *int `xml:"numDerivativesAtEnd,attr,omitempty"`
	// The attribute "numDerivativesInterior" specifies the type of continuity that is guaranteed interior to the curve. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativeInterior *int `xml:"numDerivativeInterior,attr,omitempty"`
	// The radius of the arc.
	Radius *LengthType `xml:"http://www.opengis.net/gml radius,omitempty"`
	// The bearing of the arc at the start.
	StartAngle *AngleType `xml:"http://www.opengis.net/gml startAngle,omitempty"`
	// The bearing of the arc at the end.
	EndAngle *AngleType              `xml:"http://www.opengis.net/gml endAngle,omitempty"`
	PosList  *DirectPositionListType `xml:"http://www.opengis.net/gml posList,omitempty"`
	// Deprecated with GML version 3.1.0. Use "posList" instead.
	Coordinates *CoordinatesType    `xml:"http://www.opengis.net/gml coordinates,omitempty"`
	Pos         *DirectPositionType `xml:"http://www.opengis.net/gml pos,omitempty"`
	// This property element either references a point via the XLink-attributes or contains the point element. pointProperty
	// is the predefined property which can be used by GML Application Schemas whenever a GML Feature has a property with a value that
	// is substitutable for Point.
	PointProperty *PointPropertyType `xml:"http://www.opengis.net/gml pointProperty,omitempty"`
	// Deprecated with GML version 3.1.0. Use "pointProperty" instead. Included for backwards compatibility with GML 3.0.0.
	PointRep *PointPropertyType `xml:"http://www.opengis.net/gml pointRep,omitempty"`
	// The attribute "interpolation" specifies the curve interpolation mechanism used for this segment. This mechanism
	// uses the control points and control parameters to determine the position of this curve segment. For an ArcByCenterPoint the interpolation is fixed as "circularArcCenterPointWithRadius".
	Interpolation *string `xml:"interpolation,attr,omitempty"`
	// Since this type describes always a single arc, the attribute is fixed to "1".
	NumArc int `xml:"numArc,attr"`
}

type ArcStringByBulgeType struct {
	// The attribute "numDerivativesAtStart" specifies the type of continuity between this curve segment and its predecessor. If this is the first curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtStart *int `xml:"numDerivativesAtStart,attr,omitempty"`
	// The attribute "numDerivativesAtEnd" specifies the type of continuity between this curve segment and its successor. If this is the last curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtEnd *int `xml:"numDerivativesAtEnd,attr,omitempty"`
	// The attribute "numDerivativesInterior" specifies the type of continuity that is guaranteed interior to the curve. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativeInterior *int `xml:"numDerivativeInterior,attr,omitempty"`
	// The bulge controls the offset of each arc's midpoint. The "bulge" is the real number multiplier for the normal that determines the offset direction of the midpoint of each arc. The length of the bulge sequence is exactly 1 less than the length of the control point array, since a bulge is needed for each pair of adjacent points in the control point array. The bulge is not given by a distance, since it is simply a multiplier for the normal.
	// The midpoint of the resulting arc is given by: midPoint = ((startPoint + endPoint)/2.0) + bulge*normal
	Bulge []float64 `xml:"http://www.opengis.net/gml bulge"`
	// The attribute "normal" is a vector normal (perpendicular) to the chord of the arc, the line joining the first and last
	// point of the arc. In a 2D coordinate system, there are only two possible directions for the normal, and it is often given as a signed real, indicating its length, with a positive sign indicating a left turn angle from the chord line, and a negative sign indicating a right turn from the chord. In 3D, the normal determines the plane of the arc, along with the start and endPoint of the arc.
	// The normal is usually a unit vector, but this is not absolutely necessary. If the normal is a zero vector, the geometric object becomes equivalent to the straight line between the two end points. The length of the normal sequence is exactly the same as for the bulge sequence, 1 less than the control point sequence length.
	Normal  []VectorType            `xml:"http://www.opengis.net/gml normal"`
	PosList *DirectPositionListType `xml:"http://www.opengis.net/gml posList,omitempty"`
	// Deprecated with GML version 3.1.0. Use "posList" instead.
	Coordinates *CoordinatesType     `xml:"http://www.opengis.net/gml coordinates,omitempty"`
	Pos         []DirectPositionType `xml:"http://www.opengis.net/gml pos"`
	// This property element either references a point via the XLink-attributes or contains the point element. pointProperty
	// is the predefined property which can be used by GML Application Schemas whenever a GML Feature has a property with a value that
	// is substitutable for Point.
	PointProperty []PointPropertyType `xml:"http://www.opengis.net/gml pointProperty"`
	// Deprecated with GML version 3.1.0. Use "pointProperty" instead. Included for backwards compatibility with GML 3.0.0.
	PointRep []PointPropertyType `xml:"http://www.opengis.net/gml pointRep"`
	// The attribute "interpolation" specifies the curve interpolation mechanism used for this segment. This mechanism
	// uses the control points and control parameters to determine the position of this curve segment. For an ArcStringByBulge the interpolation is fixed as "circularArc2PointWithBulge".
	Interpolation *string `xml:"interpolation,attr,omitempty"`
	// The number of arcs in the arc string can be explicitly stated in this attribute. The number of control points in the arc string must be numArc + 1.
	NumArc *int `xml:"numArc,attr,omitempty"`
}

type ArcStringType struct {
	// The attribute "numDerivativesAtStart" specifies the type of continuity between this curve segment and its predecessor. If this is the first curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtStart *int `xml:"numDerivativesAtStart,attr,omitempty"`
	// The attribute "numDerivativesAtEnd" specifies the type of continuity between this curve segment and its successor. If this is the last curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtEnd *int `xml:"numDerivativesAtEnd,attr,omitempty"`
	// The attribute "numDerivativesInterior" specifies the type of continuity that is guaranteed interior to the curve. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativeInterior *int                    `xml:"numDerivativeInterior,attr,omitempty"`
	PosList               *DirectPositionListType `xml:"http://www.opengis.net/gml posList,omitempty"`
	// Deprecated with GML version 3.1.0. Use "posList" instead.
	Coordinates *CoordinatesType     `xml:"http://www.opengis.net/gml coordinates,omitempty"`
	Pos         []DirectPositionType `xml:"http://www.opengis.net/gml pos"`
	// This property element either references a point via the XLink-attributes or contains the point element. pointProperty
	// is the predefined property which can be used by GML Application Schemas whenever a GML Feature has a property with a value that
	// is substitutable for Point.
	PointProperty []PointPropertyType `xml:"http://www.opengis.net/gml pointProperty"`
	// Deprecated with GML version 3.1.0. Use "pointProperty" instead. Included for backwards compatibility with GML 3.0.0.
	PointRep []PointPropertyType `xml:"http://www.opengis.net/gml pointRep"`
	// The attribute "interpolation" specifies the curve interpolation mechanism used for this segment. This mechanism
	// uses the control points and control parameters to determine the position of this curve segment. For an ArcString the interpolation is fixed as "circularArc3Points".
	Interpolation *string `xml:"interpolation,attr,omitempty"`
	// The number of arcs in the arc string can be explicitly stated in this attribute. The number of control points in the arc string must be 2 * numArc + 1.
	NumArc *int `xml:"numArc,attr,omitempty"`
}

type ArcType struct {
	// The attribute "numDerivativesAtStart" specifies the type of continuity between this curve segment and its predecessor. If this is the first curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtStart *int `xml:"numDerivativesAtStart,attr,omitempty"`
	// The attribute "numDerivativesAtEnd" specifies the type of continuity between this curve segment and its successor. If this is the last curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtEnd *int `xml:"numDerivativesAtEnd,attr,omitempty"`
	// The attribute "numDerivativesInterior" specifies the type of continuity that is guaranteed interior to the curve. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativeInterior *int                    `xml:"numDerivativeInterior,attr,omitempty"`
	PosList               *DirectPositionListType `xml:"http://www.opengis.net/gml posList,omitempty"`
	// Deprecated with GML version 3.1.0. Use "posList" instead.
	Coordinates *CoordinatesType     `xml:"http://www.opengis.net/gml coordinates,omitempty"`
	Pos         []DirectPositionType `xml:"http://www.opengis.net/gml pos"`
	// This property element either references a point via the XLink-attributes or contains the point element. pointProperty
	// is the predefined property which can be used by GML Application Schemas whenever a GML Feature has a property with a value that
	// is substitutable for Point.
	PointProperty []PointPropertyType `xml:"http://www.opengis.net/gml pointProperty"`
	// Deprecated with GML version 3.1.0. Use "pointProperty" instead. Included for backwards compatibility with GML 3.0.0.
	PointRep []PointPropertyType `xml:"http://www.opengis.net/gml pointRep"`
	// The attribute "interpolation" specifies the curve interpolation mechanism used for this segment. This mechanism
	// uses the control points and control parameters to determine the position of this curve segment. For an ArcString the interpolation is fixed as "circularArc3Points".
	Interpolation *string `xml:"interpolation,attr,omitempty"`
	// The number of arcs in the arc string can be explicitly stated in this attribute. The number of control points in the arc string must be 2 * numArc + 1.
	NumArc *int `xml:"numArc,attr,omitempty"`
}

type AreaType struct {
	Value string `xml:",chardata"`
}

type ArrayAssociationType struct {
	// This abstract element is the head of a substitutionGroup hierararchy which may contain either simpleContent or complexContent elements.  It is used to assert the model position of "class" elements declared in other GML schemas.
	AbstractObject                    []string                                `xml:"http://www.opengis.net/gml _Object"`
	Array                             []ArrayType                             `xml:"http://www.opengis.net/gml Array"`
	Bag                               []BagType                               `xml:"http://www.opengis.net/gml Bag"`
	BaseUnit                          []BaseUnitType                          `xml:"http://www.opengis.net/gml BaseUnit"`
	CartesianCS                       []CartesianCSType                       `xml:"http://www.opengis.net/gml CartesianCS"`
	CompositeCurve                    []CompositeCurveType                    `xml:"http://www.opengis.net/gml CompositeCurve"`
	CompositeSolid                    []CompositeSolidType                    `xml:"http://www.opengis.net/gml CompositeSolid"`
	CompositeSurface                  []CompositeSurfaceType                  `xml:"http://www.opengis.net/gml CompositeSurface"`
	CompoundCRS                       []CompoundCRSType                       `xml:"http://www.opengis.net/gml CompoundCRS"`
	ConcatenatedOperation             []ConcatenatedOperationType             `xml:"http://www.opengis.net/gml ConcatenatedOperation"`
	ConventionalUnit                  []ConventionalUnitType                  `xml:"http://www.opengis.net/gml ConventionalUnit"`
	Conversion                        []ConversionType                        `xml:"http://www.opengis.net/gml Conversion"`
	CoordinateSystemAxis              []CoordinateSystemAxisType              `xml:"http://www.opengis.net/gml CoordinateSystemAxis"`
	Curve                             []CurveType                             `xml:"http://www.opengis.net/gml Curve"`
	CylindricalCS                     []CylindricalCSType                     `xml:"http://www.opengis.net/gml CylindricalCS"`
	Definition                        []DefinitionType                        `xml:"http://www.opengis.net/gml Definition"`
	DefinitionCollection              []DictionaryType                        `xml:"http://www.opengis.net/gml DefinitionCollection"`
	DefinitionProxy                   []DefinitionProxyType                   `xml:"http://www.opengis.net/gml DefinitionProxy"`
	DerivedCRS                        []DerivedCRSType                        `xml:"http://www.opengis.net/gml DerivedCRS"`
	DerivedUnit                       []DerivedUnitType                       `xml:"http://www.opengis.net/gml DerivedUnit"`
	Dictionary                        []DictionaryType                        `xml:"http://www.opengis.net/gml Dictionary"`
	DirectedObservation               []DirectedObservationType               `xml:"http://www.opengis.net/gml DirectedObservation"`
	DirectedObservationAtDistance     []DirectedObservationAtDistanceType     `xml:"http://www.opengis.net/gml DirectedObservationAtDistance"`
	Edge                              []EdgeType                              `xml:"http://www.opengis.net/gml Edge"`
	Ellipsoid                         []EllipsoidType                         `xml:"http://www.opengis.net/gml Ellipsoid"`
	EllipsoidalCS                     []EllipsoidalCSType                     `xml:"http://www.opengis.net/gml EllipsoidalCS"`
	EngineeringCRS                    []EngineeringCRSType                    `xml:"http://www.opengis.net/gml EngineeringCRS"`
	EngineeringDatum                  []EngineeringDatumType                  `xml:"http://www.opengis.net/gml EngineeringDatum"`
	Face                              []FaceType                              `xml:"http://www.opengis.net/gml Face"`
	FeatureCollection                 []FeatureCollectionType                 `xml:"http://www.opengis.net/gml FeatureCollection"`
	FeatureStyle                      []FeatureStyleType                      `xml:"http://www.opengis.net/gml FeatureStyle"`
	GenericMetaData                   []GenericMetaDataType                   `xml:"http://www.opengis.net/gml GenericMetaData"`
	GeocentricCRS                     []GeocentricCRSType                     `xml:"http://www.opengis.net/gml GeocentricCRS"`
	GeodeticDatum                     []GeodeticDatumType                     `xml:"http://www.opengis.net/gml GeodeticDatum"`
	GeographicCRS                     []GeographicCRSType                     `xml:"http://www.opengis.net/gml GeographicCRS"`
	GeometricComplex                  []GeometricComplexType                  `xml:"http://www.opengis.net/gml GeometricComplex"`
	GeometryStyle                     []GeometryStyleType                     `xml:"http://www.opengis.net/gml GeometryStyle"`
	GraphStyle                        []GraphStyleType                        `xml:"http://www.opengis.net/gml GraphStyle"`
	Grid                              []GridType                              `xml:"http://www.opengis.net/gml Grid"`
	GridCoverage                      []GridCoverageType                      `xml:"http://www.opengis.net/gml GridCoverage"`
	ImageCRS                          []ImageCRSType                          `xml:"http://www.opengis.net/gml ImageCRS"`
	ImageDatum                        []ImageDatumType                        `xml:"http://www.opengis.net/gml ImageDatum"`
	LabelStyle                        []LabelStyleType                        `xml:"http://www.opengis.net/gml LabelStyle"`
	LineString                        []LineStringType                        `xml:"http://www.opengis.net/gml LineString"`
	LinearCS                          []LinearCSType                          `xml:"http://www.opengis.net/gml LinearCS"`
	LinearRing                        []LinearRingType                        `xml:"http://www.opengis.net/gml LinearRing"`
	MovingObjectStatus                []MovingObjectStatusType                `xml:"http://www.opengis.net/gml MovingObjectStatus"`
	MultiCurve                        []MultiCurveType                        `xml:"http://www.opengis.net/gml MultiCurve"`
	MultiCurveCoverage                []MultiCurveCoverageType                `xml:"http://www.opengis.net/gml MultiCurveCoverage"`
	MultiGeometry                     []MultiGeometryType                     `xml:"http://www.opengis.net/gml MultiGeometry"`
	MultiLineString                   []MultiLineStringType                   `xml:"http://www.opengis.net/gml MultiLineString"`
	MultiPoint                        []MultiPointType                        `xml:"http://www.opengis.net/gml MultiPoint"`
	MultiPointCoverage                []MultiPointCoverageType                `xml:"http://www.opengis.net/gml MultiPointCoverage"`
	MultiPolygon                      []MultiPolygonType                      `xml:"http://www.opengis.net/gml MultiPolygon"`
	MultiSolid                        []MultiSolidType                        `xml:"http://www.opengis.net/gml MultiSolid"`
	MultiSolidCoverage                []MultiSolidCoverageType                `xml:"http://www.opengis.net/gml MultiSolidCoverage"`
	MultiSurface                      []MultiSurfaceType                      `xml:"http://www.opengis.net/gml MultiSurface"`
	MultiSurfaceCoverage              []MultiSurfaceCoverageType              `xml:"http://www.opengis.net/gml MultiSurfaceCoverage"`
	Node                              []NodeType                              `xml:"http://www.opengis.net/gml Node"`
	ObliqueCartesianCS                []ObliqueCartesianCSType                `xml:"http://www.opengis.net/gml ObliqueCartesianCS"`
	Observation                       []ObservationType                       `xml:"http://www.opengis.net/gml Observation"`
	OperationMethod                   []OperationMethodType                   `xml:"http://www.opengis.net/gml OperationMethod"`
	OperationParameter                []OperationParameterType                `xml:"http://www.opengis.net/gml OperationParameter"`
	OperationParameterGroup           []OperationParameterGroupType           `xml:"http://www.opengis.net/gml OperationParameterGroup"`
	OrientableCurve                   []OrientableCurveType                   `xml:"http://www.opengis.net/gml OrientableCurve"`
	OrientableSurface                 []OrientableSurfaceType                 `xml:"http://www.opengis.net/gml OrientableSurface"`
	PassThroughOperation              []PassThroughOperationType              `xml:"http://www.opengis.net/gml PassThroughOperation"`
	Point                             []PointType                             `xml:"http://www.opengis.net/gml Point"`
	PolarCS                           []PolarCSType                           `xml:"http://www.opengis.net/gml PolarCS"`
	Polygon                           []PolygonType                           `xml:"http://www.opengis.net/gml Polygon"`
	PolyhedralSurface                 []PolyhedralSurfaceType                 `xml:"http://www.opengis.net/gml PolyhedralSurface"`
	PrimeMeridian                     []PrimeMeridianType                     `xml:"http://www.opengis.net/gml PrimeMeridian"`
	ProjectedCRS                      []ProjectedCRSType                      `xml:"http://www.opengis.net/gml ProjectedCRS"`
	RectifiedGrid                     []RectifiedGridType                     `xml:"http://www.opengis.net/gml RectifiedGrid"`
	RectifiedGridCoverage             []RectifiedGridCoverageType             `xml:"http://www.opengis.net/gml RectifiedGridCoverage"`
	Ring                              []RingType                              `xml:"http://www.opengis.net/gml Ring"`
	Solid                             []SolidType                             `xml:"http://www.opengis.net/gml Solid"`
	SphericalCS                       []SphericalCSType                       `xml:"http://www.opengis.net/gml SphericalCS"`
	Style                             []StyleType                             `xml:"http://www.opengis.net/gml Style"`
	Surface                           []SurfaceType                           `xml:"http://www.opengis.net/gml Surface"`
	TemporalCRS                       []TemporalCRSType                       `xml:"http://www.opengis.net/gml TemporalCRS"`
	TemporalCS                        []TemporalCSType                        `xml:"http://www.opengis.net/gml TemporalCS"`
	TemporalDatum                     []TemporalDatumType                     `xml:"http://www.opengis.net/gml TemporalDatum"`
	TimeCalendar                      []TimeCalendarType                      `xml:"http://www.opengis.net/gml TimeCalendar"`
	TimeCalendarEra                   []TimeCalendarEraType                   `xml:"http://www.opengis.net/gml TimeCalendarEra"`
	TimeClock                         []TimeClockType                         `xml:"http://www.opengis.net/gml TimeClock"`
	TimeCoordinateSystem              []TimeCoordinateSystemType              `xml:"http://www.opengis.net/gml TimeCoordinateSystem"`
	TimeEdge                          []TimeEdgeType                          `xml:"http://www.opengis.net/gml TimeEdge"`
	TimeInstant                       []TimeInstantType                       `xml:"http://www.opengis.net/gml TimeInstant"`
	TimeNode                          []TimeNodeType                          `xml:"http://www.opengis.net/gml TimeNode"`
	TimeOrdinalReferenceSystem        []TimeOrdinalReferenceSystemType        `xml:"http://www.opengis.net/gml TimeOrdinalReferenceSystem"`
	TimePeriod                        []TimePeriodType                        `xml:"http://www.opengis.net/gml TimePeriod"`
	TimeTopologyComplex               []TimeTopologyComplexType               `xml:"http://www.opengis.net/gml TimeTopologyComplex"`
	Tin                               []TinType                               `xml:"http://www.opengis.net/gml Tin"`
	TopoComplex                       []TopoComplexType                       `xml:"http://www.opengis.net/gml TopoComplex"`
	TopoSolid                         []TopoSolidType                         `xml:"http://www.opengis.net/gml TopoSolid"`
	TopologyStyle                     []TopologyStyleType                     `xml:"http://www.opengis.net/gml TopologyStyle"`
	Transformation                    []TransformationType                    `xml:"http://www.opengis.net/gml Transformation"`
	TriangulatedSurface               []TriangulatedSurfaceType               `xml:"http://www.opengis.net/gml TriangulatedSurface"`
	UnitDefinition                    []UnitDefinitionType                    `xml:"http://www.opengis.net/gml UnitDefinition"`
	UserDefinedCS                     []UserDefinedCSType                     `xml:"http://www.opengis.net/gml UserDefinedCS"`
	VerticalCRS                       []VerticalCRSType                       `xml:"http://www.opengis.net/gml VerticalCRS"`
	VerticalCS                        []VerticalCSType                        `xml:"http://www.opengis.net/gml VerticalCS"`
	VerticalDatum                     []VerticalDatumType                     `xml:"http://www.opengis.net/gml VerticalDatum"`
	AbstractCRS                       []AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _CRS"`
	AbstractContinuousCoverage        []AbstractContinuousCoverageType        `xml:"http://www.opengis.net/gml _ContinuousCoverage"`
	AbstractCoordinateOperation       []AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _CoordinateOperation"`
	AbstractCoordinateReferenceSystem []AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _CoordinateReferenceSystem"`
	AbstractCoordinateSystem          []AbstractCoordinateSystemType          `xml:"http://www.opengis.net/gml _CoordinateSystem"`
	AbstractCoverage                  []AbstractCoverageType                  `xml:"http://www.opengis.net/gml _Coverage"`
	AbstractCurve                     []AbstractCurveType                     `xml:"http://www.opengis.net/gml _Curve"`
	AbstractDatum                     []AbstractDatumType                     `xml:"http://www.opengis.net/gml _Datum"`
	AbstractDiscreteCoverage          []AbstractDiscreteCoverageType          `xml:"http://www.opengis.net/gml _DiscreteCoverage"`
	AbstractFeature                   []AbstractFeatureType                   `xml:"http://www.opengis.net/gml _Feature"`
	AbstractFeatureCollection         []AbstractFeatureCollectionType         `xml:"http://www.opengis.net/gml _FeatureCollection"`
	AbstractGML                       []AbstractGMLType                       `xml:"http://www.opengis.net/gml _GML"`
	AbstractGeneralConversion         []AbstractGeneralConversionType         `xml:"http://www.opengis.net/gml _GeneralConversion"`
	AbstractGeneralDerivedCRS         []AbstractGeneralDerivedCRSType         `xml:"http://www.opengis.net/gml _GeneralDerivedCRS"`
	AbstractGeneralOperationParameter []AbstractGeneralOperationParameterType `xml:"http://www.opengis.net/gml _GeneralOperationParameter"`
	AbstractGeneralTransformation     []AbstractGeneralTransformationType     `xml:"http://www.opengis.net/gml _GeneralTransformation"`
	AbstractGeometricAggregate        []AbstractGeometricAggregateType        `xml:"http://www.opengis.net/gml _GeometricAggregate"`
	AbstractGeometricPrimitive        []AbstractGeometricPrimitiveType        `xml:"http://www.opengis.net/gml _GeometricPrimitive"`
	AbstractGeometry                  []AbstractGeometryType                  `xml:"http://www.opengis.net/gml _Geometry"`
	AbstractImplicitGeometry          []AbstractGeometryType                  `xml:"http://www.opengis.net/gml _ImplicitGeometry"`
	AbstractMetaData                  []AbstractMetaDataType                  `xml:"http://www.opengis.net/gml _MetaData"`
	AbstractOperation                 []AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _Operation"`
	AbstractReferenceSystem           []AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _ReferenceSystem"`
	AbstractRing                      []AbstractRingType                      `xml:"http://www.opengis.net/gml _Ring"`
	AbstractSingleOperation           []AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _SingleOperation"`
	AbstractSolid                     []AbstractSolidType                     `xml:"http://www.opengis.net/gml _Solid"`
	AbstractStyle                     []AbstractStyleType                     `xml:"http://www.opengis.net/gml _Style"`
	AbstractSurface                   []AbstractSurfaceType                   `xml:"http://www.opengis.net/gml _Surface"`
	AbstractTimeComplex               []AbstractTimeComplexType               `xml:"http://www.opengis.net/gml _TimeComplex"`
	AbstractTimeGeometricPrimitive    []AbstractTimeGeometricPrimitiveType    `xml:"http://www.opengis.net/gml _TimeGeometricPrimitive"`
	AbstractTimeObject                []AbstractTimeObjectType                `xml:"http://www.opengis.net/gml _TimeObject"`
	AbstractTimePrimitive             []AbstractTimePrimitiveType             `xml:"http://www.opengis.net/gml _TimePrimitive"`
	AbstractTimeReferenceSystem       []AbstractTimeReferenceSystemType       `xml:"http://www.opengis.net/gml _TimeReferenceSystem"`
	AbstractTimeSlice                 []AbstractTimeSliceType                 `xml:"http://www.opengis.net/gml _TimeSlice"`
	AbstractTimeTopologyPrimitive     []AbstractTimeTopologyPrimitiveType     `xml:"http://www.opengis.net/gml _TimeTopologyPrimitive"`
	AbstractTopoPrimitive             []AbstractTopoPrimitiveType             `xml:"http://www.opengis.net/gml _TopoPrimitive"`
	AbstractTopology                  []AbstractTopologyType                  `xml:"http://www.opengis.net/gml _Topology"`
}

type ArrayType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType            `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType            `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType            `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType            `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType            `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType            `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType            `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType            `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType            `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType            `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	Members                 *ArrayAssociationType `xml:"http://www.opengis.net/gml members,omitempty"`
}

type AssociationType struct {
	// This abstract element is the head of a substitutionGroup hierararchy which may contain either simpleContent or complexContent elements.  It is used to assert the model position of "class" elements declared in other GML schemas.
	AbstractObject                    *string                                `xml:"http://www.opengis.net/gml _Object,omitempty"`
	Array                             *ArrayType                             `xml:"http://www.opengis.net/gml Array,omitempty"`
	Bag                               *BagType                               `xml:"http://www.opengis.net/gml Bag,omitempty"`
	BaseUnit                          *BaseUnitType                          `xml:"http://www.opengis.net/gml BaseUnit,omitempty"`
	CartesianCS                       *CartesianCSType                       `xml:"http://www.opengis.net/gml CartesianCS,omitempty"`
	CompositeCurve                    *CompositeCurveType                    `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	CompositeSolid                    *CompositeSolidType                    `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	CompositeSurface                  *CompositeSurfaceType                  `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	CompoundCRS                       *CompoundCRSType                       `xml:"http://www.opengis.net/gml CompoundCRS,omitempty"`
	ConcatenatedOperation             *ConcatenatedOperationType             `xml:"http://www.opengis.net/gml ConcatenatedOperation,omitempty"`
	ConventionalUnit                  *ConventionalUnitType                  `xml:"http://www.opengis.net/gml ConventionalUnit,omitempty"`
	Conversion                        *ConversionType                        `xml:"http://www.opengis.net/gml Conversion,omitempty"`
	CoordinateSystemAxis              *CoordinateSystemAxisType              `xml:"http://www.opengis.net/gml CoordinateSystemAxis,omitempty"`
	Curve                             *CurveType                             `xml:"http://www.opengis.net/gml Curve,omitempty"`
	CylindricalCS                     *CylindricalCSType                     `xml:"http://www.opengis.net/gml CylindricalCS,omitempty"`
	Definition                        *DefinitionType                        `xml:"http://www.opengis.net/gml Definition,omitempty"`
	DefinitionCollection              *DictionaryType                        `xml:"http://www.opengis.net/gml DefinitionCollection,omitempty"`
	DefinitionProxy                   *DefinitionProxyType                   `xml:"http://www.opengis.net/gml DefinitionProxy,omitempty"`
	DerivedCRS                        *DerivedCRSType                        `xml:"http://www.opengis.net/gml DerivedCRS,omitempty"`
	DerivedUnit                       *DerivedUnitType                       `xml:"http://www.opengis.net/gml DerivedUnit,omitempty"`
	Dictionary                        *DictionaryType                        `xml:"http://www.opengis.net/gml Dictionary,omitempty"`
	DirectedObservation               *DirectedObservationType               `xml:"http://www.opengis.net/gml DirectedObservation,omitempty"`
	DirectedObservationAtDistance     *DirectedObservationAtDistanceType     `xml:"http://www.opengis.net/gml DirectedObservationAtDistance,omitempty"`
	Edge                              *EdgeType                              `xml:"http://www.opengis.net/gml Edge,omitempty"`
	Ellipsoid                         *EllipsoidType                         `xml:"http://www.opengis.net/gml Ellipsoid,omitempty"`
	EllipsoidalCS                     *EllipsoidalCSType                     `xml:"http://www.opengis.net/gml EllipsoidalCS,omitempty"`
	EngineeringCRS                    *EngineeringCRSType                    `xml:"http://www.opengis.net/gml EngineeringCRS,omitempty"`
	EngineeringDatum                  *EngineeringDatumType                  `xml:"http://www.opengis.net/gml EngineeringDatum,omitempty"`
	Face                              *FaceType                              `xml:"http://www.opengis.net/gml Face,omitempty"`
	FeatureCollection                 *FeatureCollectionType                 `xml:"http://www.opengis.net/gml FeatureCollection,omitempty"`
	FeatureStyle                      *FeatureStyleType                      `xml:"http://www.opengis.net/gml FeatureStyle,omitempty"`
	GenericMetaData                   *GenericMetaDataType                   `xml:"http://www.opengis.net/gml GenericMetaData,omitempty"`
	GeocentricCRS                     *GeocentricCRSType                     `xml:"http://www.opengis.net/gml GeocentricCRS,omitempty"`
	GeodeticDatum                     *GeodeticDatumType                     `xml:"http://www.opengis.net/gml GeodeticDatum,omitempty"`
	GeographicCRS                     *GeographicCRSType                     `xml:"http://www.opengis.net/gml GeographicCRS,omitempty"`
	GeometricComplex                  *GeometricComplexType                  `xml:"http://www.opengis.net/gml GeometricComplex,omitempty"`
	GeometryStyle                     *GeometryStyleType                     `xml:"http://www.opengis.net/gml GeometryStyle,omitempty"`
	GraphStyle                        *GraphStyleType                        `xml:"http://www.opengis.net/gml GraphStyle,omitempty"`
	Grid                              *GridType                              `xml:"http://www.opengis.net/gml Grid,omitempty"`
	GridCoverage                      *GridCoverageType                      `xml:"http://www.opengis.net/gml GridCoverage,omitempty"`
	ImageCRS                          *ImageCRSType                          `xml:"http://www.opengis.net/gml ImageCRS,omitempty"`
	ImageDatum                        *ImageDatumType                        `xml:"http://www.opengis.net/gml ImageDatum,omitempty"`
	LabelStyle                        *LabelStyleType                        `xml:"http://www.opengis.net/gml LabelStyle,omitempty"`
	LineString                        *LineStringType                        `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearCS                          *LinearCSType                          `xml:"http://www.opengis.net/gml LinearCS,omitempty"`
	LinearRing                        *LinearRingType                        `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MovingObjectStatus                *MovingObjectStatusType                `xml:"http://www.opengis.net/gml MovingObjectStatus,omitempty"`
	MultiCurve                        *MultiCurveType                        `xml:"http://www.opengis.net/gml MultiCurve,omitempty"`
	MultiCurveCoverage                *MultiCurveCoverageType                `xml:"http://www.opengis.net/gml MultiCurveCoverage,omitempty"`
	MultiGeometry                     *MultiGeometryType                     `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString                   *MultiLineStringType                   `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                        *MultiPointType                        `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPointCoverage                *MultiPointCoverageType                `xml:"http://www.opengis.net/gml MultiPointCoverage,omitempty"`
	MultiPolygon                      *MultiPolygonType                      `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	MultiSolid                        *MultiSolidType                        `xml:"http://www.opengis.net/gml MultiSolid,omitempty"`
	MultiSolidCoverage                *MultiSolidCoverageType                `xml:"http://www.opengis.net/gml MultiSolidCoverage,omitempty"`
	MultiSurface                      *MultiSurfaceType                      `xml:"http://www.opengis.net/gml MultiSurface,omitempty"`
	MultiSurfaceCoverage              *MultiSurfaceCoverageType              `xml:"http://www.opengis.net/gml MultiSurfaceCoverage,omitempty"`
	Node                              *NodeType                              `xml:"http://www.opengis.net/gml Node,omitempty"`
	ObliqueCartesianCS                *ObliqueCartesianCSType                `xml:"http://www.opengis.net/gml ObliqueCartesianCS,omitempty"`
	Observation                       *ObservationType                       `xml:"http://www.opengis.net/gml Observation,omitempty"`
	OperationMethod                   *OperationMethodType                   `xml:"http://www.opengis.net/gml OperationMethod,omitempty"`
	OperationParameter                *OperationParameterType                `xml:"http://www.opengis.net/gml OperationParameter,omitempty"`
	OperationParameterGroup           *OperationParameterGroupType           `xml:"http://www.opengis.net/gml OperationParameterGroup,omitempty"`
	OrientableCurve                   *OrientableCurveType                   `xml:"http://www.opengis.net/gml OrientableCurve,omitempty"`
	OrientableSurface                 *OrientableSurfaceType                 `xml:"http://www.opengis.net/gml OrientableSurface,omitempty"`
	PassThroughOperation              *PassThroughOperationType              `xml:"http://www.opengis.net/gml PassThroughOperation,omitempty"`
	Point                             *PointType                             `xml:"http://www.opengis.net/gml Point,omitempty"`
	PolarCS                           *PolarCSType                           `xml:"http://www.opengis.net/gml PolarCS,omitempty"`
	Polygon                           *PolygonType                           `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	PolyhedralSurface                 *PolyhedralSurfaceType                 `xml:"http://www.opengis.net/gml PolyhedralSurface,omitempty"`
	PrimeMeridian                     *PrimeMeridianType                     `xml:"http://www.opengis.net/gml PrimeMeridian,omitempty"`
	ProjectedCRS                      *ProjectedCRSType                      `xml:"http://www.opengis.net/gml ProjectedCRS,omitempty"`
	RectifiedGrid                     *RectifiedGridType                     `xml:"http://www.opengis.net/gml RectifiedGrid,omitempty"`
	RectifiedGridCoverage             *RectifiedGridCoverageType             `xml:"http://www.opengis.net/gml RectifiedGridCoverage,omitempty"`
	Ring                              *RingType                              `xml:"http://www.opengis.net/gml Ring,omitempty"`
	Solid                             *SolidType                             `xml:"http://www.opengis.net/gml Solid,omitempty"`
	SphericalCS                       *SphericalCSType                       `xml:"http://www.opengis.net/gml SphericalCS,omitempty"`
	Style                             *StyleType                             `xml:"http://www.opengis.net/gml Style,omitempty"`
	Surface                           *SurfaceType                           `xml:"http://www.opengis.net/gml Surface,omitempty"`
	TemporalCRS                       *TemporalCRSType                       `xml:"http://www.opengis.net/gml TemporalCRS,omitempty"`
	TemporalCS                        *TemporalCSType                        `xml:"http://www.opengis.net/gml TemporalCS,omitempty"`
	TemporalDatum                     *TemporalDatumType                     `xml:"http://www.opengis.net/gml TemporalDatum,omitempty"`
	TimeCalendar                      *TimeCalendarType                      `xml:"http://www.opengis.net/gml TimeCalendar,omitempty"`
	TimeCalendarEra                   *TimeCalendarEraType                   `xml:"http://www.opengis.net/gml TimeCalendarEra,omitempty"`
	TimeClock                         *TimeClockType                         `xml:"http://www.opengis.net/gml TimeClock,omitempty"`
	TimeCoordinateSystem              *TimeCoordinateSystemType              `xml:"http://www.opengis.net/gml TimeCoordinateSystem,omitempty"`
	TimeEdge                          *TimeEdgeType                          `xml:"http://www.opengis.net/gml TimeEdge,omitempty"`
	TimeInstant                       *TimeInstantType                       `xml:"http://www.opengis.net/gml TimeInstant,omitempty"`
	TimeNode                          *TimeNodeType                          `xml:"http://www.opengis.net/gml TimeNode,omitempty"`
	TimeOrdinalReferenceSystem        *TimeOrdinalReferenceSystemType        `xml:"http://www.opengis.net/gml TimeOrdinalReferenceSystem,omitempty"`
	TimePeriod                        *TimePeriodType                        `xml:"http://www.opengis.net/gml TimePeriod,omitempty"`
	TimeTopologyComplex               *TimeTopologyComplexType               `xml:"http://www.opengis.net/gml TimeTopologyComplex,omitempty"`
	Tin                               *TinType                               `xml:"http://www.opengis.net/gml Tin,omitempty"`
	TopoComplex                       *TopoComplexType                       `xml:"http://www.opengis.net/gml TopoComplex,omitempty"`
	TopoSolid                         *TopoSolidType                         `xml:"http://www.opengis.net/gml TopoSolid,omitempty"`
	TopologyStyle                     *TopologyStyleType                     `xml:"http://www.opengis.net/gml TopologyStyle,omitempty"`
	Transformation                    *TransformationType                    `xml:"http://www.opengis.net/gml Transformation,omitempty"`
	TriangulatedSurface               *TriangulatedSurfaceType               `xml:"http://www.opengis.net/gml TriangulatedSurface,omitempty"`
	UnitDefinition                    *UnitDefinitionType                    `xml:"http://www.opengis.net/gml UnitDefinition,omitempty"`
	UserDefinedCS                     *UserDefinedCSType                     `xml:"http://www.opengis.net/gml UserDefinedCS,omitempty"`
	VerticalCRS                       *VerticalCRSType                       `xml:"http://www.opengis.net/gml VerticalCRS,omitempty"`
	VerticalCS                        *VerticalCSType                        `xml:"http://www.opengis.net/gml VerticalCS,omitempty"`
	VerticalDatum                     *VerticalDatumType                     `xml:"http://www.opengis.net/gml VerticalDatum,omitempty"`
	AbstractCRS                       *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _CRS,omitempty"`
	AbstractContinuousCoverage        *AbstractContinuousCoverageType        `xml:"http://www.opengis.net/gml _ContinuousCoverage,omitempty"`
	AbstractCoordinateOperation       *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _CoordinateOperation,omitempty"`
	AbstractCoordinateReferenceSystem *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _CoordinateReferenceSystem,omitempty"`
	AbstractCoordinateSystem          *AbstractCoordinateSystemType          `xml:"http://www.opengis.net/gml _CoordinateSystem,omitempty"`
	AbstractCoverage                  *AbstractCoverageType                  `xml:"http://www.opengis.net/gml _Coverage,omitempty"`
	AbstractCurve                     *AbstractCurveType                     `xml:"http://www.opengis.net/gml _Curve,omitempty"`
	AbstractDatum                     *AbstractDatumType                     `xml:"http://www.opengis.net/gml _Datum,omitempty"`
	AbstractDiscreteCoverage          *AbstractDiscreteCoverageType          `xml:"http://www.opengis.net/gml _DiscreteCoverage,omitempty"`
	AbstractFeature                   *AbstractFeatureType                   `xml:"http://www.opengis.net/gml _Feature,omitempty"`
	AbstractFeatureCollection         *AbstractFeatureCollectionType         `xml:"http://www.opengis.net/gml _FeatureCollection,omitempty"`
	AbstractGML                       *AbstractGMLType                       `xml:"http://www.opengis.net/gml _GML,omitempty"`
	AbstractGeneralConversion         *AbstractGeneralConversionType         `xml:"http://www.opengis.net/gml _GeneralConversion,omitempty"`
	AbstractGeneralDerivedCRS         *AbstractGeneralDerivedCRSType         `xml:"http://www.opengis.net/gml _GeneralDerivedCRS,omitempty"`
	AbstractGeneralOperationParameter *AbstractGeneralOperationParameterType `xml:"http://www.opengis.net/gml _GeneralOperationParameter,omitempty"`
	AbstractGeneralTransformation     *AbstractGeneralTransformationType     `xml:"http://www.opengis.net/gml _GeneralTransformation,omitempty"`
	AbstractGeometricAggregate        *AbstractGeometricAggregateType        `xml:"http://www.opengis.net/gml _GeometricAggregate,omitempty"`
	AbstractGeometricPrimitive        *AbstractGeometricPrimitiveType        `xml:"http://www.opengis.net/gml _GeometricPrimitive,omitempty"`
	AbstractGeometry                  *AbstractGeometryType                  `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	AbstractImplicitGeometry          *AbstractGeometryType                  `xml:"http://www.opengis.net/gml _ImplicitGeometry,omitempty"`
	AbstractMetaData                  *AbstractMetaDataType                  `xml:"http://www.opengis.net/gml _MetaData,omitempty"`
	AbstractOperation                 *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _Operation,omitempty"`
	AbstractReferenceSystem           *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _ReferenceSystem,omitempty"`
	AbstractRing                      *AbstractRingType                      `xml:"http://www.opengis.net/gml _Ring,omitempty"`
	AbstractSingleOperation           *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _SingleOperation,omitempty"`
	AbstractSolid                     *AbstractSolidType                     `xml:"http://www.opengis.net/gml _Solid,omitempty"`
	AbstractStyle                     *AbstractStyleType                     `xml:"http://www.opengis.net/gml _Style,omitempty"`
	AbstractSurface                   *AbstractSurfaceType                   `xml:"http://www.opengis.net/gml _Surface,omitempty"`
	AbstractTimeComplex               *AbstractTimeComplexType               `xml:"http://www.opengis.net/gml _TimeComplex,omitempty"`
	AbstractTimeGeometricPrimitive    *AbstractTimeGeometricPrimitiveType    `xml:"http://www.opengis.net/gml _TimeGeometricPrimitive,omitempty"`
	AbstractTimeObject                *AbstractTimeObjectType                `xml:"http://www.opengis.net/gml _TimeObject,omitempty"`
	AbstractTimePrimitive             *AbstractTimePrimitiveType             `xml:"http://www.opengis.net/gml _TimePrimitive,omitempty"`
	AbstractTimeReferenceSystem       *AbstractTimeReferenceSystemType       `xml:"http://www.opengis.net/gml _TimeReferenceSystem,omitempty"`
	AbstractTimeSlice                 *AbstractTimeSliceType                 `xml:"http://www.opengis.net/gml _TimeSlice,omitempty"`
	AbstractTimeTopologyPrimitive     *AbstractTimeTopologyPrimitiveType     `xml:"http://www.opengis.net/gml _TimeTopologyPrimitive,omitempty"`
	AbstractTopoPrimitive             *AbstractTopoPrimitiveType             `xml:"http://www.opengis.net/gml _TopoPrimitive,omitempty"`
	AbstractTopology                  *AbstractTopologyType                  `xml:"http://www.opengis.net/gml _Topology,omitempty"`
	RemoteSchema                      string                                 `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                         string                                 `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                              string                                 `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                              string                                 `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                           string                                 `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                             string                                 `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                              string                                 `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                           string                                 `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type BSplineType struct {
	// The attribute "numDerivativesAtStart" specifies the type of continuity between this curve segment and its predecessor. If this is the first curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtStart *int `xml:"numDerivativesAtStart,attr,omitempty"`
	// The attribute "numDerivativesAtEnd" specifies the type of continuity between this curve segment and its successor. If this is the last curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtEnd *int `xml:"numDerivativesAtEnd,attr,omitempty"`
	// The attribute "numDerivativesInterior" specifies the type of continuity that is guaranteed interior to the curve. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativeInterior *int `xml:"numDerivativeInterior,attr,omitempty"`
	// The attribute "degree" shall be the degree of the polynomial used for interpolation in this spline.
	Degree int `xml:"http://www.opengis.net/gml degree,omitempty"`
	// The property "knot" shall be the sequence of distinct knots used to define the spline basis functions.
	Knot    []KnotPropertyType      `xml:"http://www.opengis.net/gml knot"`
	PosList *DirectPositionListType `xml:"http://www.opengis.net/gml posList,omitempty"`
	// Deprecated with GML version 3.1.0. Use "posList" instead.
	Coordinates *CoordinatesType     `xml:"http://www.opengis.net/gml coordinates,omitempty"`
	Pos         []DirectPositionType `xml:"http://www.opengis.net/gml pos"`
	// This property element either references a point via the XLink-attributes or contains the point element. pointProperty
	// is the predefined property which can be used by GML Application Schemas whenever a GML Feature has a property with a value that
	// is substitutable for Point.
	PointProperty []PointPropertyType `xml:"http://www.opengis.net/gml pointProperty"`
	// Deprecated with GML version 3.1.0. Use "pointProperty" instead. Included for backwards compatibility with GML 3.0.0.
	PointRep []PointPropertyType `xml:"http://www.opengis.net/gml pointRep"`
	// The attribute "interpolation" specifies the curve interpolation mechanism used for this segment. This mechanism
	// uses the control points and control parameters to determine the position of this curve segment. For a BSpline the interpolation can be either "polynomialSpline" or "rationalSpline", default is "polynomialSpline".
	Interpolation *string `xml:"interpolation,attr,omitempty"`
	// The attribute isPolynomial is set to true if this is a polynomial spline.
	IsPolynomial *bool `xml:"isPolynomial,attr,omitempty"`
	// The attribute "knotType" gives the type of knot distribution used in defining this spline. This is for information only
	// and is set according to the different construction-functions.
	KnotType *string `xml:"knotType,attr,omitempty"`
}

type BagType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType            `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType            `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType            `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType            `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType            `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType            `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType            `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType            `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType            `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType            `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	Member                  []AssociationType     `xml:"http://www.opengis.net/gml member"`
	Members                 *ArrayAssociationType `xml:"http://www.opengis.net/gml members,omitempty"`
}

type BaseStyleDescriptorType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType           `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType           `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType           `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType           `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType           `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType           `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType           `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType           `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType           `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType           `xml:"http://www.opengis.net/gml srsName"`
	Id                      string               `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	SpatialResolution       *ScaleType           `xml:"http://www.opengis.net/gml spatialResolution,omitempty"`
	StyleVariation          []StyleVariationType `xml:"http://www.opengis.net/gml styleVariation"`
	Animate                 []string             `xml:"http://www.w3.org/2001/SMIL20/ animate"`
	AnimateMotion           []string             `xml:"http://www.w3.org/2001/SMIL20/ animateMotion"`
	AnimateColor            []string             `xml:"http://www.w3.org/2001/SMIL20/ animateColor"`
	Set                     []string             `xml:"http://www.w3.org/2001/SMIL20/ set"`
}

type BaseUnitType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Informal description of the phenomenon or type of quantity that is measured or observed. For example, "length", "angle", "time", "pressure", or "temperature". When the quantity is the result of an observation or measurement, this term is known as Observable Type or Measurand.
	QuantityType *StringOrRefType `xml:"http://www.opengis.net/gml quantityType,omitempty"`
	// For global understanding of a unit of measure, it is often possible to reference an item in a catalog of units, using a symbol in that catalog. The "codeSpace" attribute in "CodeType" identifies a namespace for the catalog symbol value, and might reference the catalog. The "string" value in "CodeType" contains the value of a symbol that is unique within this catalog namespace. This symbol often appears explicitly in the catalog, but it could be a combination of symbols using a specified algebra of units. For example, the symbol "cm" might indicate that it is the "m" symbol combined with the "c" prefix.
	CatalogSymbol *CodeType      `xml:"http://www.opengis.net/gml catalogSymbol,omitempty"`
	UnitsSystem   *ReferenceType `xml:"http://www.opengis.net/gml unitsSystem,omitempty"`
}

type BezierType struct {
	// The attribute "numDerivativesAtStart" specifies the type of continuity between this curve segment and its predecessor. If this is the first curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtStart *int `xml:"numDerivativesAtStart,attr,omitempty"`
	// The attribute "numDerivativesAtEnd" specifies the type of continuity between this curve segment and its successor. If this is the last curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtEnd *int `xml:"numDerivativesAtEnd,attr,omitempty"`
	// The attribute "numDerivativesInterior" specifies the type of continuity that is guaranteed interior to the curve. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativeInterior *int `xml:"numDerivativeInterior,attr,omitempty"`
	// The attribute "degree" shall be the degree of the polynomial used for interpolation in this spline.
	Degree int `xml:"http://www.opengis.net/gml degree,omitempty"`
	// The property "knot" shall be the sequence of distinct knots used to define the spline basis functions.
	Knot    []KnotPropertyType      `xml:"http://www.opengis.net/gml knot"`
	PosList *DirectPositionListType `xml:"http://www.opengis.net/gml posList,omitempty"`
	// Deprecated with GML version 3.1.0. Use "posList" instead.
	Coordinates *CoordinatesType     `xml:"http://www.opengis.net/gml coordinates,omitempty"`
	Pos         []DirectPositionType `xml:"http://www.opengis.net/gml pos"`
	// This property element either references a point via the XLink-attributes or contains the point element. pointProperty
	// is the predefined property which can be used by GML Application Schemas whenever a GML Feature has a property with a value that
	// is substitutable for Point.
	PointProperty []PointPropertyType `xml:"http://www.opengis.net/gml pointProperty"`
	// Deprecated with GML version 3.1.0. Use "pointProperty" instead. Included for backwards compatibility with GML 3.0.0.
	PointRep []PointPropertyType `xml:"http://www.opengis.net/gml pointRep"`
	// The attribute "interpolation" specifies the curve interpolation mechanism used for this segment. This mechanism
	// uses the control points and control parameters to determine the position of this curve segment. For a BSpline the interpolation can be either "polynomialSpline" or "rationalSpline", default is "polynomialSpline".
	Interpolation *string `xml:"interpolation,attr,omitempty"`
	// The attribute isPolynomial is set to true if this is a polynomial spline.
	IsPolynomial *bool `xml:"isPolynomial,attr,omitempty"`
	// The attribute "knotType" gives the type of knot distribution used in defining this spline. This is for information only
	// and is set according to the different construction-functions.
	KnotType *string `xml:"knotType,attr,omitempty"`
}

type BooleanPropertyType struct {
	// This abstract element is the head of a substitutionGroup hierararchy which may contain either simpleContent or complexContent elements.  It is used to assert the model position of "class" elements declared in other GML schemas.
	AbstractObject                    *string                                `xml:"http://www.opengis.net/gml _Object,omitempty"`
	Array                             *ArrayType                             `xml:"http://www.opengis.net/gml Array,omitempty"`
	Bag                               *BagType                               `xml:"http://www.opengis.net/gml Bag,omitempty"`
	BaseUnit                          *BaseUnitType                          `xml:"http://www.opengis.net/gml BaseUnit,omitempty"`
	CartesianCS                       *CartesianCSType                       `xml:"http://www.opengis.net/gml CartesianCS,omitempty"`
	CompositeCurve                    *CompositeCurveType                    `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	CompositeSolid                    *CompositeSolidType                    `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	CompositeSurface                  *CompositeSurfaceType                  `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	CompoundCRS                       *CompoundCRSType                       `xml:"http://www.opengis.net/gml CompoundCRS,omitempty"`
	ConcatenatedOperation             *ConcatenatedOperationType             `xml:"http://www.opengis.net/gml ConcatenatedOperation,omitempty"`
	ConventionalUnit                  *ConventionalUnitType                  `xml:"http://www.opengis.net/gml ConventionalUnit,omitempty"`
	Conversion                        *ConversionType                        `xml:"http://www.opengis.net/gml Conversion,omitempty"`
	CoordinateSystemAxis              *CoordinateSystemAxisType              `xml:"http://www.opengis.net/gml CoordinateSystemAxis,omitempty"`
	Curve                             *CurveType                             `xml:"http://www.opengis.net/gml Curve,omitempty"`
	CylindricalCS                     *CylindricalCSType                     `xml:"http://www.opengis.net/gml CylindricalCS,omitempty"`
	Definition                        *DefinitionType                        `xml:"http://www.opengis.net/gml Definition,omitempty"`
	DefinitionCollection              *DictionaryType                        `xml:"http://www.opengis.net/gml DefinitionCollection,omitempty"`
	DefinitionProxy                   *DefinitionProxyType                   `xml:"http://www.opengis.net/gml DefinitionProxy,omitempty"`
	DerivedCRS                        *DerivedCRSType                        `xml:"http://www.opengis.net/gml DerivedCRS,omitempty"`
	DerivedUnit                       *DerivedUnitType                       `xml:"http://www.opengis.net/gml DerivedUnit,omitempty"`
	Dictionary                        *DictionaryType                        `xml:"http://www.opengis.net/gml Dictionary,omitempty"`
	DirectedObservation               *DirectedObservationType               `xml:"http://www.opengis.net/gml DirectedObservation,omitempty"`
	DirectedObservationAtDistance     *DirectedObservationAtDistanceType     `xml:"http://www.opengis.net/gml DirectedObservationAtDistance,omitempty"`
	Edge                              *EdgeType                              `xml:"http://www.opengis.net/gml Edge,omitempty"`
	Ellipsoid                         *EllipsoidType                         `xml:"http://www.opengis.net/gml Ellipsoid,omitempty"`
	EllipsoidalCS                     *EllipsoidalCSType                     `xml:"http://www.opengis.net/gml EllipsoidalCS,omitempty"`
	EngineeringCRS                    *EngineeringCRSType                    `xml:"http://www.opengis.net/gml EngineeringCRS,omitempty"`
	EngineeringDatum                  *EngineeringDatumType                  `xml:"http://www.opengis.net/gml EngineeringDatum,omitempty"`
	Face                              *FaceType                              `xml:"http://www.opengis.net/gml Face,omitempty"`
	FeatureCollection                 *FeatureCollectionType                 `xml:"http://www.opengis.net/gml FeatureCollection,omitempty"`
	FeatureStyle                      *FeatureStyleType                      `xml:"http://www.opengis.net/gml FeatureStyle,omitempty"`
	GenericMetaData                   *GenericMetaDataType                   `xml:"http://www.opengis.net/gml GenericMetaData,omitempty"`
	GeocentricCRS                     *GeocentricCRSType                     `xml:"http://www.opengis.net/gml GeocentricCRS,omitempty"`
	GeodeticDatum                     *GeodeticDatumType                     `xml:"http://www.opengis.net/gml GeodeticDatum,omitempty"`
	GeographicCRS                     *GeographicCRSType                     `xml:"http://www.opengis.net/gml GeographicCRS,omitempty"`
	GeometricComplex                  *GeometricComplexType                  `xml:"http://www.opengis.net/gml GeometricComplex,omitempty"`
	GeometryStyle                     *GeometryStyleType                     `xml:"http://www.opengis.net/gml GeometryStyle,omitempty"`
	GraphStyle                        *GraphStyleType                        `xml:"http://www.opengis.net/gml GraphStyle,omitempty"`
	Grid                              *GridType                              `xml:"http://www.opengis.net/gml Grid,omitempty"`
	GridCoverage                      *GridCoverageType                      `xml:"http://www.opengis.net/gml GridCoverage,omitempty"`
	ImageCRS                          *ImageCRSType                          `xml:"http://www.opengis.net/gml ImageCRS,omitempty"`
	ImageDatum                        *ImageDatumType                        `xml:"http://www.opengis.net/gml ImageDatum,omitempty"`
	LabelStyle                        *LabelStyleType                        `xml:"http://www.opengis.net/gml LabelStyle,omitempty"`
	LineString                        *LineStringType                        `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearCS                          *LinearCSType                          `xml:"http://www.opengis.net/gml LinearCS,omitempty"`
	LinearRing                        *LinearRingType                        `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MovingObjectStatus                *MovingObjectStatusType                `xml:"http://www.opengis.net/gml MovingObjectStatus,omitempty"`
	MultiCurve                        *MultiCurveType                        `xml:"http://www.opengis.net/gml MultiCurve,omitempty"`
	MultiCurveCoverage                *MultiCurveCoverageType                `xml:"http://www.opengis.net/gml MultiCurveCoverage,omitempty"`
	MultiGeometry                     *MultiGeometryType                     `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString                   *MultiLineStringType                   `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                        *MultiPointType                        `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPointCoverage                *MultiPointCoverageType                `xml:"http://www.opengis.net/gml MultiPointCoverage,omitempty"`
	MultiPolygon                      *MultiPolygonType                      `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	MultiSolid                        *MultiSolidType                        `xml:"http://www.opengis.net/gml MultiSolid,omitempty"`
	MultiSolidCoverage                *MultiSolidCoverageType                `xml:"http://www.opengis.net/gml MultiSolidCoverage,omitempty"`
	MultiSurface                      *MultiSurfaceType                      `xml:"http://www.opengis.net/gml MultiSurface,omitempty"`
	MultiSurfaceCoverage              *MultiSurfaceCoverageType              `xml:"http://www.opengis.net/gml MultiSurfaceCoverage,omitempty"`
	Node                              *NodeType                              `xml:"http://www.opengis.net/gml Node,omitempty"`
	ObliqueCartesianCS                *ObliqueCartesianCSType                `xml:"http://www.opengis.net/gml ObliqueCartesianCS,omitempty"`
	Observation                       *ObservationType                       `xml:"http://www.opengis.net/gml Observation,omitempty"`
	OperationMethod                   *OperationMethodType                   `xml:"http://www.opengis.net/gml OperationMethod,omitempty"`
	OperationParameter                *OperationParameterType                `xml:"http://www.opengis.net/gml OperationParameter,omitempty"`
	OperationParameterGroup           *OperationParameterGroupType           `xml:"http://www.opengis.net/gml OperationParameterGroup,omitempty"`
	OrientableCurve                   *OrientableCurveType                   `xml:"http://www.opengis.net/gml OrientableCurve,omitempty"`
	OrientableSurface                 *OrientableSurfaceType                 `xml:"http://www.opengis.net/gml OrientableSurface,omitempty"`
	PassThroughOperation              *PassThroughOperationType              `xml:"http://www.opengis.net/gml PassThroughOperation,omitempty"`
	Point                             *PointType                             `xml:"http://www.opengis.net/gml Point,omitempty"`
	PolarCS                           *PolarCSType                           `xml:"http://www.opengis.net/gml PolarCS,omitempty"`
	Polygon                           *PolygonType                           `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	PolyhedralSurface                 *PolyhedralSurfaceType                 `xml:"http://www.opengis.net/gml PolyhedralSurface,omitempty"`
	PrimeMeridian                     *PrimeMeridianType                     `xml:"http://www.opengis.net/gml PrimeMeridian,omitempty"`
	ProjectedCRS                      *ProjectedCRSType                      `xml:"http://www.opengis.net/gml ProjectedCRS,omitempty"`
	RectifiedGrid                     *RectifiedGridType                     `xml:"http://www.opengis.net/gml RectifiedGrid,omitempty"`
	RectifiedGridCoverage             *RectifiedGridCoverageType             `xml:"http://www.opengis.net/gml RectifiedGridCoverage,omitempty"`
	Ring                              *RingType                              `xml:"http://www.opengis.net/gml Ring,omitempty"`
	Solid                             *SolidType                             `xml:"http://www.opengis.net/gml Solid,omitempty"`
	SphericalCS                       *SphericalCSType                       `xml:"http://www.opengis.net/gml SphericalCS,omitempty"`
	Style                             *StyleType                             `xml:"http://www.opengis.net/gml Style,omitempty"`
	Surface                           *SurfaceType                           `xml:"http://www.opengis.net/gml Surface,omitempty"`
	TemporalCRS                       *TemporalCRSType                       `xml:"http://www.opengis.net/gml TemporalCRS,omitempty"`
	TemporalCS                        *TemporalCSType                        `xml:"http://www.opengis.net/gml TemporalCS,omitempty"`
	TemporalDatum                     *TemporalDatumType                     `xml:"http://www.opengis.net/gml TemporalDatum,omitempty"`
	TimeCalendar                      *TimeCalendarType                      `xml:"http://www.opengis.net/gml TimeCalendar,omitempty"`
	TimeCalendarEra                   *TimeCalendarEraType                   `xml:"http://www.opengis.net/gml TimeCalendarEra,omitempty"`
	TimeClock                         *TimeClockType                         `xml:"http://www.opengis.net/gml TimeClock,omitempty"`
	TimeCoordinateSystem              *TimeCoordinateSystemType              `xml:"http://www.opengis.net/gml TimeCoordinateSystem,omitempty"`
	TimeEdge                          *TimeEdgeType                          `xml:"http://www.opengis.net/gml TimeEdge,omitempty"`
	TimeInstant                       *TimeInstantType                       `xml:"http://www.opengis.net/gml TimeInstant,omitempty"`
	TimeNode                          *TimeNodeType                          `xml:"http://www.opengis.net/gml TimeNode,omitempty"`
	TimeOrdinalReferenceSystem        *TimeOrdinalReferenceSystemType        `xml:"http://www.opengis.net/gml TimeOrdinalReferenceSystem,omitempty"`
	TimePeriod                        *TimePeriodType                        `xml:"http://www.opengis.net/gml TimePeriod,omitempty"`
	TimeTopologyComplex               *TimeTopologyComplexType               `xml:"http://www.opengis.net/gml TimeTopologyComplex,omitempty"`
	Tin                               *TinType                               `xml:"http://www.opengis.net/gml Tin,omitempty"`
	TopoComplex                       *TopoComplexType                       `xml:"http://www.opengis.net/gml TopoComplex,omitempty"`
	TopoSolid                         *TopoSolidType                         `xml:"http://www.opengis.net/gml TopoSolid,omitempty"`
	TopologyStyle                     *TopologyStyleType                     `xml:"http://www.opengis.net/gml TopologyStyle,omitempty"`
	Transformation                    *TransformationType                    `xml:"http://www.opengis.net/gml Transformation,omitempty"`
	TriangulatedSurface               *TriangulatedSurfaceType               `xml:"http://www.opengis.net/gml TriangulatedSurface,omitempty"`
	UnitDefinition                    *UnitDefinitionType                    `xml:"http://www.opengis.net/gml UnitDefinition,omitempty"`
	UserDefinedCS                     *UserDefinedCSType                     `xml:"http://www.opengis.net/gml UserDefinedCS,omitempty"`
	VerticalCRS                       *VerticalCRSType                       `xml:"http://www.opengis.net/gml VerticalCRS,omitempty"`
	VerticalCS                        *VerticalCSType                        `xml:"http://www.opengis.net/gml VerticalCS,omitempty"`
	VerticalDatum                     *VerticalDatumType                     `xml:"http://www.opengis.net/gml VerticalDatum,omitempty"`
	AbstractCRS                       *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _CRS,omitempty"`
	AbstractContinuousCoverage        *AbstractContinuousCoverageType        `xml:"http://www.opengis.net/gml _ContinuousCoverage,omitempty"`
	AbstractCoordinateOperation       *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _CoordinateOperation,omitempty"`
	AbstractCoordinateReferenceSystem *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _CoordinateReferenceSystem,omitempty"`
	AbstractCoordinateSystem          *AbstractCoordinateSystemType          `xml:"http://www.opengis.net/gml _CoordinateSystem,omitempty"`
	AbstractCoverage                  *AbstractCoverageType                  `xml:"http://www.opengis.net/gml _Coverage,omitempty"`
	AbstractCurve                     *AbstractCurveType                     `xml:"http://www.opengis.net/gml _Curve,omitempty"`
	AbstractDatum                     *AbstractDatumType                     `xml:"http://www.opengis.net/gml _Datum,omitempty"`
	AbstractDiscreteCoverage          *AbstractDiscreteCoverageType          `xml:"http://www.opengis.net/gml _DiscreteCoverage,omitempty"`
	AbstractFeature                   *AbstractFeatureType                   `xml:"http://www.opengis.net/gml _Feature,omitempty"`
	AbstractFeatureCollection         *AbstractFeatureCollectionType         `xml:"http://www.opengis.net/gml _FeatureCollection,omitempty"`
	AbstractGML                       *AbstractGMLType                       `xml:"http://www.opengis.net/gml _GML,omitempty"`
	AbstractGeneralConversion         *AbstractGeneralConversionType         `xml:"http://www.opengis.net/gml _GeneralConversion,omitempty"`
	AbstractGeneralDerivedCRS         *AbstractGeneralDerivedCRSType         `xml:"http://www.opengis.net/gml _GeneralDerivedCRS,omitempty"`
	AbstractGeneralOperationParameter *AbstractGeneralOperationParameterType `xml:"http://www.opengis.net/gml _GeneralOperationParameter,omitempty"`
	AbstractGeneralTransformation     *AbstractGeneralTransformationType     `xml:"http://www.opengis.net/gml _GeneralTransformation,omitempty"`
	AbstractGeometricAggregate        *AbstractGeometricAggregateType        `xml:"http://www.opengis.net/gml _GeometricAggregate,omitempty"`
	AbstractGeometricPrimitive        *AbstractGeometricPrimitiveType        `xml:"http://www.opengis.net/gml _GeometricPrimitive,omitempty"`
	AbstractGeometry                  *AbstractGeometryType                  `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	AbstractImplicitGeometry          *AbstractGeometryType                  `xml:"http://www.opengis.net/gml _ImplicitGeometry,omitempty"`
	AbstractMetaData                  *AbstractMetaDataType                  `xml:"http://www.opengis.net/gml _MetaData,omitempty"`
	AbstractOperation                 *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _Operation,omitempty"`
	AbstractReferenceSystem           *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _ReferenceSystem,omitempty"`
	AbstractRing                      *AbstractRingType                      `xml:"http://www.opengis.net/gml _Ring,omitempty"`
	AbstractSingleOperation           *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _SingleOperation,omitempty"`
	AbstractSolid                     *AbstractSolidType                     `xml:"http://www.opengis.net/gml _Solid,omitempty"`
	AbstractStyle                     *AbstractStyleType                     `xml:"http://www.opengis.net/gml _Style,omitempty"`
	AbstractSurface                   *AbstractSurfaceType                   `xml:"http://www.opengis.net/gml _Surface,omitempty"`
	AbstractTimeComplex               *AbstractTimeComplexType               `xml:"http://www.opengis.net/gml _TimeComplex,omitempty"`
	AbstractTimeGeometricPrimitive    *AbstractTimeGeometricPrimitiveType    `xml:"http://www.opengis.net/gml _TimeGeometricPrimitive,omitempty"`
	AbstractTimeObject                *AbstractTimeObjectType                `xml:"http://www.opengis.net/gml _TimeObject,omitempty"`
	AbstractTimePrimitive             *AbstractTimePrimitiveType             `xml:"http://www.opengis.net/gml _TimePrimitive,omitempty"`
	AbstractTimeReferenceSystem       *AbstractTimeReferenceSystemType       `xml:"http://www.opengis.net/gml _TimeReferenceSystem,omitempty"`
	AbstractTimeSlice                 *AbstractTimeSliceType                 `xml:"http://www.opengis.net/gml _TimeSlice,omitempty"`
	AbstractTimeTopologyPrimitive     *AbstractTimeTopologyPrimitiveType     `xml:"http://www.opengis.net/gml _TimeTopologyPrimitive,omitempty"`
	AbstractTopoPrimitive             *AbstractTopoPrimitiveType             `xml:"http://www.opengis.net/gml _TopoPrimitive,omitempty"`
	AbstractTopology                  *AbstractTopologyType                  `xml:"http://www.opengis.net/gml _Topology,omitempty"`
	Null                              *string                                `xml:"http://www.opengis.net/gml Null,omitempty"`
	// Aggregate value built using the Composite pattern.
	CompositeValue *CompositeValueType `xml:"http://www.opengis.net/gml CompositeValue,omitempty"`
	ValueArray     *ValueArrayType     `xml:"http://www.opengis.net/gml ValueArray,omitempty"`
	// A value from two-valued logic, using the XML Schema boolean type.  An instance may take the values {true, false, 1, 0}.
	Boolean *bool `xml:"http://www.opengis.net/gml Boolean,omitempty"`
	// A term representing a classification.  It has an optional XML attribute codeSpace, whose value is a URI which identifies a dictionary, codelist or authority for the term.
	Category *CodeType `xml:"http://www.opengis.net/gml Category,omitempty"`
	// A numeric value with a scale.  The content of the element is an amount using the XML Schema type double which permits decimal or scientific notation.  An XML attribute uom (unit of measure) is required, whose value is a URI which identifies the definition of the scale or units by which the numeric value must be multiplied.
	Quantity *MeasureType `xml:"http://www.opengis.net/gml Quantity,omitempty"`
	// An integer representing a frequency of occurrence.
	Count *int `xml:"http://www.opengis.net/gml Count,omitempty"`
	// XML List based on XML Schema boolean type.  An element of this type contains a space-separated list of boolean values {0,1,true,false}
	BooleanList *string `xml:"http://www.opengis.net/gml BooleanList,omitempty"`
	// A space-separated list of terms or nulls.  A single XML attribute codeSpace may be provided, which authorises all the terms in the list.
	CategoryList *CodeOrNullListType `xml:"http://www.opengis.net/gml CategoryList,omitempty"`
	// A space separated list of amounts or nulls.  The amounts use the XML Schema type double.  A single XML attribute uom (unit of measure) is required, whose value is a URI which identifies the definition of the scale or units by which all the amounts in the list must be multiplied.
	QuantityList *MeasureOrNullListType `xml:"http://www.opengis.net/gml QuantityList,omitempty"`
	// A space-separated list of integers or nulls.
	CountList *string `xml:"http://www.opengis.net/gml CountList,omitempty"`
	// Utility element to store a 2-point range of ordinal values. If one member is a null, then this is a single ended interval.
	CategoryExtent *CategoryExtentType `xml:"http://www.opengis.net/gml CategoryExtent,omitempty"`
	// Utility element to store a 2-point range of numeric values. If one member is a null, then this is a single ended interval.
	QuantityExtent *QuantityExtentType `xml:"http://www.opengis.net/gml QuantityExtent,omitempty"`
	// Utility element to store a 2-point range of frequency values. If one member is a null, then this is a single ended interval.
	CountExtent  *string `xml:"http://www.opengis.net/gml CountExtent,omitempty"`
	RemoteSchema string  `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string  `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string  `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string  `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string  `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string  `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string  `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string  `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type BoundedFeatureType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType         `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType         `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType         `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType         `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType         `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType         `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType         `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType         `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType         `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType         `xml:"http://www.opengis.net/gml srsName"`
	Id                      string             `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	BoundedBy               *BoundingShapeType `xml:"http://www.opengis.net/gml boundedBy,omitempty"`
	// deprecated in GML version 3.1
	Location         *LocationPropertyType         `xml:"http://www.opengis.net/gml location,omitempty"`
	PriorityLocation *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml priorityLocation,omitempty"`
}

type BoundingShapeType struct {
	Envelope               *EnvelopeType               `xml:"http://www.opengis.net/gml Envelope,omitempty"`
	EnvelopeWithTimePeriod *EnvelopeWithTimePeriodType `xml:"http://www.opengis.net/gml EnvelopeWithTimePeriod,omitempty"`
	Null                   *string                     `xml:"http://www.opengis.net/gml Null,omitempty"`
}

type CRSRefType struct {
	// Abstract coordinate reference system, usually defined by a coordinate system and a datum. This abstract complexType shall not be used, extended, or restricted, in an Application Schema, to define a concrete subtype with a meaning equivalent to a concrete subtype specified in this document.
	AbstractCRS                       *AbstractReferenceSystemType   `xml:"http://www.opengis.net/gml _CRS,omitempty"`
	CompoundCRS                       *CompoundCRSType               `xml:"http://www.opengis.net/gml CompoundCRS,omitempty"`
	DerivedCRS                        *DerivedCRSType                `xml:"http://www.opengis.net/gml DerivedCRS,omitempty"`
	EngineeringCRS                    *EngineeringCRSType            `xml:"http://www.opengis.net/gml EngineeringCRS,omitempty"`
	GeocentricCRS                     *GeocentricCRSType             `xml:"http://www.opengis.net/gml GeocentricCRS,omitempty"`
	GeographicCRS                     *GeographicCRSType             `xml:"http://www.opengis.net/gml GeographicCRS,omitempty"`
	ImageCRS                          *ImageCRSType                  `xml:"http://www.opengis.net/gml ImageCRS,omitempty"`
	ProjectedCRS                      *ProjectedCRSType              `xml:"http://www.opengis.net/gml ProjectedCRS,omitempty"`
	TemporalCRS                       *TemporalCRSType               `xml:"http://www.opengis.net/gml TemporalCRS,omitempty"`
	VerticalCRS                       *VerticalCRSType               `xml:"http://www.opengis.net/gml VerticalCRS,omitempty"`
	AbstractCoordinateReferenceSystem *AbstractReferenceSystemType   `xml:"http://www.opengis.net/gml _CoordinateReferenceSystem,omitempty"`
	AbstractGeneralDerivedCRS         *AbstractGeneralDerivedCRSType `xml:"http://www.opengis.net/gml _GeneralDerivedCRS,omitempty"`
	RemoteSchema                      string                         `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                         string                         `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                              string                         `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                              string                         `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                           string                         `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                             string                         `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                              string                         `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                           string                         `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type CartesianCSRefType struct {
	CartesianCS  *CartesianCSType `xml:"http://www.opengis.net/gml CartesianCS,omitempty"`
	RemoteSchema string           `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string           `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string           `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string           `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string           `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string           `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string           `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string           `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type CartesianCSType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this coordinate system. The first csID, if any, is normally the primary identification code, and any others are aliases.
	CsID []IdentifierType `xml:"http://www.opengis.net/gml csID"`
	// Comments on or information about this coordinate system, including data source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Ordered sequence of associations to the coordinate system axes included in this coordinate system.
	UsesAxis []CoordinateSystemAxisRefType `xml:"http://www.opengis.net/gml usesAxis"`
}

type CategoryExtentType struct {
	Value string `xml:",chardata"`
}

type CategoryPropertyType struct {
	// This abstract element is the head of a substitutionGroup hierararchy which may contain either simpleContent or complexContent elements.  It is used to assert the model position of "class" elements declared in other GML schemas.
	AbstractObject                    *string                                `xml:"http://www.opengis.net/gml _Object,omitempty"`
	Array                             *ArrayType                             `xml:"http://www.opengis.net/gml Array,omitempty"`
	Bag                               *BagType                               `xml:"http://www.opengis.net/gml Bag,omitempty"`
	BaseUnit                          *BaseUnitType                          `xml:"http://www.opengis.net/gml BaseUnit,omitempty"`
	CartesianCS                       *CartesianCSType                       `xml:"http://www.opengis.net/gml CartesianCS,omitempty"`
	CompositeCurve                    *CompositeCurveType                    `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	CompositeSolid                    *CompositeSolidType                    `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	CompositeSurface                  *CompositeSurfaceType                  `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	CompoundCRS                       *CompoundCRSType                       `xml:"http://www.opengis.net/gml CompoundCRS,omitempty"`
	ConcatenatedOperation             *ConcatenatedOperationType             `xml:"http://www.opengis.net/gml ConcatenatedOperation,omitempty"`
	ConventionalUnit                  *ConventionalUnitType                  `xml:"http://www.opengis.net/gml ConventionalUnit,omitempty"`
	Conversion                        *ConversionType                        `xml:"http://www.opengis.net/gml Conversion,omitempty"`
	CoordinateSystemAxis              *CoordinateSystemAxisType              `xml:"http://www.opengis.net/gml CoordinateSystemAxis,omitempty"`
	Curve                             *CurveType                             `xml:"http://www.opengis.net/gml Curve,omitempty"`
	CylindricalCS                     *CylindricalCSType                     `xml:"http://www.opengis.net/gml CylindricalCS,omitempty"`
	Definition                        *DefinitionType                        `xml:"http://www.opengis.net/gml Definition,omitempty"`
	DefinitionCollection              *DictionaryType                        `xml:"http://www.opengis.net/gml DefinitionCollection,omitempty"`
	DefinitionProxy                   *DefinitionProxyType                   `xml:"http://www.opengis.net/gml DefinitionProxy,omitempty"`
	DerivedCRS                        *DerivedCRSType                        `xml:"http://www.opengis.net/gml DerivedCRS,omitempty"`
	DerivedUnit                       *DerivedUnitType                       `xml:"http://www.opengis.net/gml DerivedUnit,omitempty"`
	Dictionary                        *DictionaryType                        `xml:"http://www.opengis.net/gml Dictionary,omitempty"`
	DirectedObservation               *DirectedObservationType               `xml:"http://www.opengis.net/gml DirectedObservation,omitempty"`
	DirectedObservationAtDistance     *DirectedObservationAtDistanceType     `xml:"http://www.opengis.net/gml DirectedObservationAtDistance,omitempty"`
	Edge                              *EdgeType                              `xml:"http://www.opengis.net/gml Edge,omitempty"`
	Ellipsoid                         *EllipsoidType                         `xml:"http://www.opengis.net/gml Ellipsoid,omitempty"`
	EllipsoidalCS                     *EllipsoidalCSType                     `xml:"http://www.opengis.net/gml EllipsoidalCS,omitempty"`
	EngineeringCRS                    *EngineeringCRSType                    `xml:"http://www.opengis.net/gml EngineeringCRS,omitempty"`
	EngineeringDatum                  *EngineeringDatumType                  `xml:"http://www.opengis.net/gml EngineeringDatum,omitempty"`
	Face                              *FaceType                              `xml:"http://www.opengis.net/gml Face,omitempty"`
	FeatureCollection                 *FeatureCollectionType                 `xml:"http://www.opengis.net/gml FeatureCollection,omitempty"`
	FeatureStyle                      *FeatureStyleType                      `xml:"http://www.opengis.net/gml FeatureStyle,omitempty"`
	GenericMetaData                   *GenericMetaDataType                   `xml:"http://www.opengis.net/gml GenericMetaData,omitempty"`
	GeocentricCRS                     *GeocentricCRSType                     `xml:"http://www.opengis.net/gml GeocentricCRS,omitempty"`
	GeodeticDatum                     *GeodeticDatumType                     `xml:"http://www.opengis.net/gml GeodeticDatum,omitempty"`
	GeographicCRS                     *GeographicCRSType                     `xml:"http://www.opengis.net/gml GeographicCRS,omitempty"`
	GeometricComplex                  *GeometricComplexType                  `xml:"http://www.opengis.net/gml GeometricComplex,omitempty"`
	GeometryStyle                     *GeometryStyleType                     `xml:"http://www.opengis.net/gml GeometryStyle,omitempty"`
	GraphStyle                        *GraphStyleType                        `xml:"http://www.opengis.net/gml GraphStyle,omitempty"`
	Grid                              *GridType                              `xml:"http://www.opengis.net/gml Grid,omitempty"`
	GridCoverage                      *GridCoverageType                      `xml:"http://www.opengis.net/gml GridCoverage,omitempty"`
	ImageCRS                          *ImageCRSType                          `xml:"http://www.opengis.net/gml ImageCRS,omitempty"`
	ImageDatum                        *ImageDatumType                        `xml:"http://www.opengis.net/gml ImageDatum,omitempty"`
	LabelStyle                        *LabelStyleType                        `xml:"http://www.opengis.net/gml LabelStyle,omitempty"`
	LineString                        *LineStringType                        `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearCS                          *LinearCSType                          `xml:"http://www.opengis.net/gml LinearCS,omitempty"`
	LinearRing                        *LinearRingType                        `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MovingObjectStatus                *MovingObjectStatusType                `xml:"http://www.opengis.net/gml MovingObjectStatus,omitempty"`
	MultiCurve                        *MultiCurveType                        `xml:"http://www.opengis.net/gml MultiCurve,omitempty"`
	MultiCurveCoverage                *MultiCurveCoverageType                `xml:"http://www.opengis.net/gml MultiCurveCoverage,omitempty"`
	MultiGeometry                     *MultiGeometryType                     `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString                   *MultiLineStringType                   `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                        *MultiPointType                        `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPointCoverage                *MultiPointCoverageType                `xml:"http://www.opengis.net/gml MultiPointCoverage,omitempty"`
	MultiPolygon                      *MultiPolygonType                      `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	MultiSolid                        *MultiSolidType                        `xml:"http://www.opengis.net/gml MultiSolid,omitempty"`
	MultiSolidCoverage                *MultiSolidCoverageType                `xml:"http://www.opengis.net/gml MultiSolidCoverage,omitempty"`
	MultiSurface                      *MultiSurfaceType                      `xml:"http://www.opengis.net/gml MultiSurface,omitempty"`
	MultiSurfaceCoverage              *MultiSurfaceCoverageType              `xml:"http://www.opengis.net/gml MultiSurfaceCoverage,omitempty"`
	Node                              *NodeType                              `xml:"http://www.opengis.net/gml Node,omitempty"`
	ObliqueCartesianCS                *ObliqueCartesianCSType                `xml:"http://www.opengis.net/gml ObliqueCartesianCS,omitempty"`
	Observation                       *ObservationType                       `xml:"http://www.opengis.net/gml Observation,omitempty"`
	OperationMethod                   *OperationMethodType                   `xml:"http://www.opengis.net/gml OperationMethod,omitempty"`
	OperationParameter                *OperationParameterType                `xml:"http://www.opengis.net/gml OperationParameter,omitempty"`
	OperationParameterGroup           *OperationParameterGroupType           `xml:"http://www.opengis.net/gml OperationParameterGroup,omitempty"`
	OrientableCurve                   *OrientableCurveType                   `xml:"http://www.opengis.net/gml OrientableCurve,omitempty"`
	OrientableSurface                 *OrientableSurfaceType                 `xml:"http://www.opengis.net/gml OrientableSurface,omitempty"`
	PassThroughOperation              *PassThroughOperationType              `xml:"http://www.opengis.net/gml PassThroughOperation,omitempty"`
	Point                             *PointType                             `xml:"http://www.opengis.net/gml Point,omitempty"`
	PolarCS                           *PolarCSType                           `xml:"http://www.opengis.net/gml PolarCS,omitempty"`
	Polygon                           *PolygonType                           `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	PolyhedralSurface                 *PolyhedralSurfaceType                 `xml:"http://www.opengis.net/gml PolyhedralSurface,omitempty"`
	PrimeMeridian                     *PrimeMeridianType                     `xml:"http://www.opengis.net/gml PrimeMeridian,omitempty"`
	ProjectedCRS                      *ProjectedCRSType                      `xml:"http://www.opengis.net/gml ProjectedCRS,omitempty"`
	RectifiedGrid                     *RectifiedGridType                     `xml:"http://www.opengis.net/gml RectifiedGrid,omitempty"`
	RectifiedGridCoverage             *RectifiedGridCoverageType             `xml:"http://www.opengis.net/gml RectifiedGridCoverage,omitempty"`
	Ring                              *RingType                              `xml:"http://www.opengis.net/gml Ring,omitempty"`
	Solid                             *SolidType                             `xml:"http://www.opengis.net/gml Solid,omitempty"`
	SphericalCS                       *SphericalCSType                       `xml:"http://www.opengis.net/gml SphericalCS,omitempty"`
	Style                             *StyleType                             `xml:"http://www.opengis.net/gml Style,omitempty"`
	Surface                           *SurfaceType                           `xml:"http://www.opengis.net/gml Surface,omitempty"`
	TemporalCRS                       *TemporalCRSType                       `xml:"http://www.opengis.net/gml TemporalCRS,omitempty"`
	TemporalCS                        *TemporalCSType                        `xml:"http://www.opengis.net/gml TemporalCS,omitempty"`
	TemporalDatum                     *TemporalDatumType                     `xml:"http://www.opengis.net/gml TemporalDatum,omitempty"`
	TimeCalendar                      *TimeCalendarType                      `xml:"http://www.opengis.net/gml TimeCalendar,omitempty"`
	TimeCalendarEra                   *TimeCalendarEraType                   `xml:"http://www.opengis.net/gml TimeCalendarEra,omitempty"`
	TimeClock                         *TimeClockType                         `xml:"http://www.opengis.net/gml TimeClock,omitempty"`
	TimeCoordinateSystem              *TimeCoordinateSystemType              `xml:"http://www.opengis.net/gml TimeCoordinateSystem,omitempty"`
	TimeEdge                          *TimeEdgeType                          `xml:"http://www.opengis.net/gml TimeEdge,omitempty"`
	TimeInstant                       *TimeInstantType                       `xml:"http://www.opengis.net/gml TimeInstant,omitempty"`
	TimeNode                          *TimeNodeType                          `xml:"http://www.opengis.net/gml TimeNode,omitempty"`
	TimeOrdinalReferenceSystem        *TimeOrdinalReferenceSystemType        `xml:"http://www.opengis.net/gml TimeOrdinalReferenceSystem,omitempty"`
	TimePeriod                        *TimePeriodType                        `xml:"http://www.opengis.net/gml TimePeriod,omitempty"`
	TimeTopologyComplex               *TimeTopologyComplexType               `xml:"http://www.opengis.net/gml TimeTopologyComplex,omitempty"`
	Tin                               *TinType                               `xml:"http://www.opengis.net/gml Tin,omitempty"`
	TopoComplex                       *TopoComplexType                       `xml:"http://www.opengis.net/gml TopoComplex,omitempty"`
	TopoSolid                         *TopoSolidType                         `xml:"http://www.opengis.net/gml TopoSolid,omitempty"`
	TopologyStyle                     *TopologyStyleType                     `xml:"http://www.opengis.net/gml TopologyStyle,omitempty"`
	Transformation                    *TransformationType                    `xml:"http://www.opengis.net/gml Transformation,omitempty"`
	TriangulatedSurface               *TriangulatedSurfaceType               `xml:"http://www.opengis.net/gml TriangulatedSurface,omitempty"`
	UnitDefinition                    *UnitDefinitionType                    `xml:"http://www.opengis.net/gml UnitDefinition,omitempty"`
	UserDefinedCS                     *UserDefinedCSType                     `xml:"http://www.opengis.net/gml UserDefinedCS,omitempty"`
	VerticalCRS                       *VerticalCRSType                       `xml:"http://www.opengis.net/gml VerticalCRS,omitempty"`
	VerticalCS                        *VerticalCSType                        `xml:"http://www.opengis.net/gml VerticalCS,omitempty"`
	VerticalDatum                     *VerticalDatumType                     `xml:"http://www.opengis.net/gml VerticalDatum,omitempty"`
	AbstractCRS                       *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _CRS,omitempty"`
	AbstractContinuousCoverage        *AbstractContinuousCoverageType        `xml:"http://www.opengis.net/gml _ContinuousCoverage,omitempty"`
	AbstractCoordinateOperation       *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _CoordinateOperation,omitempty"`
	AbstractCoordinateReferenceSystem *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _CoordinateReferenceSystem,omitempty"`
	AbstractCoordinateSystem          *AbstractCoordinateSystemType          `xml:"http://www.opengis.net/gml _CoordinateSystem,omitempty"`
	AbstractCoverage                  *AbstractCoverageType                  `xml:"http://www.opengis.net/gml _Coverage,omitempty"`
	AbstractCurve                     *AbstractCurveType                     `xml:"http://www.opengis.net/gml _Curve,omitempty"`
	AbstractDatum                     *AbstractDatumType                     `xml:"http://www.opengis.net/gml _Datum,omitempty"`
	AbstractDiscreteCoverage          *AbstractDiscreteCoverageType          `xml:"http://www.opengis.net/gml _DiscreteCoverage,omitempty"`
	AbstractFeature                   *AbstractFeatureType                   `xml:"http://www.opengis.net/gml _Feature,omitempty"`
	AbstractFeatureCollection         *AbstractFeatureCollectionType         `xml:"http://www.opengis.net/gml _FeatureCollection,omitempty"`
	AbstractGML                       *AbstractGMLType                       `xml:"http://www.opengis.net/gml _GML,omitempty"`
	AbstractGeneralConversion         *AbstractGeneralConversionType         `xml:"http://www.opengis.net/gml _GeneralConversion,omitempty"`
	AbstractGeneralDerivedCRS         *AbstractGeneralDerivedCRSType         `xml:"http://www.opengis.net/gml _GeneralDerivedCRS,omitempty"`
	AbstractGeneralOperationParameter *AbstractGeneralOperationParameterType `xml:"http://www.opengis.net/gml _GeneralOperationParameter,omitempty"`
	AbstractGeneralTransformation     *AbstractGeneralTransformationType     `xml:"http://www.opengis.net/gml _GeneralTransformation,omitempty"`
	AbstractGeometricAggregate        *AbstractGeometricAggregateType        `xml:"http://www.opengis.net/gml _GeometricAggregate,omitempty"`
	AbstractGeometricPrimitive        *AbstractGeometricPrimitiveType        `xml:"http://www.opengis.net/gml _GeometricPrimitive,omitempty"`
	AbstractGeometry                  *AbstractGeometryType                  `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	AbstractImplicitGeometry          *AbstractGeometryType                  `xml:"http://www.opengis.net/gml _ImplicitGeometry,omitempty"`
	AbstractMetaData                  *AbstractMetaDataType                  `xml:"http://www.opengis.net/gml _MetaData,omitempty"`
	AbstractOperation                 *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _Operation,omitempty"`
	AbstractReferenceSystem           *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _ReferenceSystem,omitempty"`
	AbstractRing                      *AbstractRingType                      `xml:"http://www.opengis.net/gml _Ring,omitempty"`
	AbstractSingleOperation           *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _SingleOperation,omitempty"`
	AbstractSolid                     *AbstractSolidType                     `xml:"http://www.opengis.net/gml _Solid,omitempty"`
	AbstractStyle                     *AbstractStyleType                     `xml:"http://www.opengis.net/gml _Style,omitempty"`
	AbstractSurface                   *AbstractSurfaceType                   `xml:"http://www.opengis.net/gml _Surface,omitempty"`
	AbstractTimeComplex               *AbstractTimeComplexType               `xml:"http://www.opengis.net/gml _TimeComplex,omitempty"`
	AbstractTimeGeometricPrimitive    *AbstractTimeGeometricPrimitiveType    `xml:"http://www.opengis.net/gml _TimeGeometricPrimitive,omitempty"`
	AbstractTimeObject                *AbstractTimeObjectType                `xml:"http://www.opengis.net/gml _TimeObject,omitempty"`
	AbstractTimePrimitive             *AbstractTimePrimitiveType             `xml:"http://www.opengis.net/gml _TimePrimitive,omitempty"`
	AbstractTimeReferenceSystem       *AbstractTimeReferenceSystemType       `xml:"http://www.opengis.net/gml _TimeReferenceSystem,omitempty"`
	AbstractTimeSlice                 *AbstractTimeSliceType                 `xml:"http://www.opengis.net/gml _TimeSlice,omitempty"`
	AbstractTimeTopologyPrimitive     *AbstractTimeTopologyPrimitiveType     `xml:"http://www.opengis.net/gml _TimeTopologyPrimitive,omitempty"`
	AbstractTopoPrimitive             *AbstractTopoPrimitiveType             `xml:"http://www.opengis.net/gml _TopoPrimitive,omitempty"`
	AbstractTopology                  *AbstractTopologyType                  `xml:"http://www.opengis.net/gml _Topology,omitempty"`
	Null                              *string                                `xml:"http://www.opengis.net/gml Null,omitempty"`
	// Aggregate value built using the Composite pattern.
	CompositeValue *CompositeValueType `xml:"http://www.opengis.net/gml CompositeValue,omitempty"`
	ValueArray     *ValueArrayType     `xml:"http://www.opengis.net/gml ValueArray,omitempty"`
	// A value from two-valued logic, using the XML Schema boolean type.  An instance may take the values {true, false, 1, 0}.
	Boolean *bool `xml:"http://www.opengis.net/gml Boolean,omitempty"`
	// A term representing a classification.  It has an optional XML attribute codeSpace, whose value is a URI which identifies a dictionary, codelist or authority for the term.
	Category *CodeType `xml:"http://www.opengis.net/gml Category,omitempty"`
	// A numeric value with a scale.  The content of the element is an amount using the XML Schema type double which permits decimal or scientific notation.  An XML attribute uom (unit of measure) is required, whose value is a URI which identifies the definition of the scale or units by which the numeric value must be multiplied.
	Quantity *MeasureType `xml:"http://www.opengis.net/gml Quantity,omitempty"`
	// An integer representing a frequency of occurrence.
	Count *int `xml:"http://www.opengis.net/gml Count,omitempty"`
	// XML List based on XML Schema boolean type.  An element of this type contains a space-separated list of boolean values {0,1,true,false}
	BooleanList *string `xml:"http://www.opengis.net/gml BooleanList,omitempty"`
	// A space-separated list of terms or nulls.  A single XML attribute codeSpace may be provided, which authorises all the terms in the list.
	CategoryList *CodeOrNullListType `xml:"http://www.opengis.net/gml CategoryList,omitempty"`
	// A space separated list of amounts or nulls.  The amounts use the XML Schema type double.  A single XML attribute uom (unit of measure) is required, whose value is a URI which identifies the definition of the scale or units by which all the amounts in the list must be multiplied.
	QuantityList *MeasureOrNullListType `xml:"http://www.opengis.net/gml QuantityList,omitempty"`
	// A space-separated list of integers or nulls.
	CountList *string `xml:"http://www.opengis.net/gml CountList,omitempty"`
	// Utility element to store a 2-point range of ordinal values. If one member is a null, then this is a single ended interval.
	CategoryExtent *CategoryExtentType `xml:"http://www.opengis.net/gml CategoryExtent,omitempty"`
	// Utility element to store a 2-point range of numeric values. If one member is a null, then this is a single ended interval.
	QuantityExtent *QuantityExtentType `xml:"http://www.opengis.net/gml QuantityExtent,omitempty"`
	// Utility element to store a 2-point range of frequency values. If one member is a null, then this is a single ended interval.
	CountExtent  *string `xml:"http://www.opengis.net/gml CountExtent,omitempty"`
	RemoteSchema string  `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string  `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string  `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string  `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string  `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string  `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string  `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string  `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type CircleByCenterPointType struct {
	// The attribute "numDerivativesAtStart" specifies the type of continuity between this curve segment and its predecessor. If this is the first curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtStart *int `xml:"numDerivativesAtStart,attr,omitempty"`
	// The attribute "numDerivativesAtEnd" specifies the type of continuity between this curve segment and its successor. If this is the last curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtEnd *int `xml:"numDerivativesAtEnd,attr,omitempty"`
	// The attribute "numDerivativesInterior" specifies the type of continuity that is guaranteed interior to the curve. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativeInterior *int `xml:"numDerivativeInterior,attr,omitempty"`
	// The radius of the arc.
	Radius *LengthType `xml:"http://www.opengis.net/gml radius,omitempty"`
	// The bearing of the arc at the start.
	StartAngle *AngleType `xml:"http://www.opengis.net/gml startAngle,omitempty"`
	// The bearing of the arc at the end.
	EndAngle *AngleType              `xml:"http://www.opengis.net/gml endAngle,omitempty"`
	PosList  *DirectPositionListType `xml:"http://www.opengis.net/gml posList,omitempty"`
	// Deprecated with GML version 3.1.0. Use "posList" instead.
	Coordinates *CoordinatesType    `xml:"http://www.opengis.net/gml coordinates,omitempty"`
	Pos         *DirectPositionType `xml:"http://www.opengis.net/gml pos,omitempty"`
	// This property element either references a point via the XLink-attributes or contains the point element. pointProperty
	// is the predefined property which can be used by GML Application Schemas whenever a GML Feature has a property with a value that
	// is substitutable for Point.
	PointProperty *PointPropertyType `xml:"http://www.opengis.net/gml pointProperty,omitempty"`
	// Deprecated with GML version 3.1.0. Use "pointProperty" instead. Included for backwards compatibility with GML 3.0.0.
	PointRep *PointPropertyType `xml:"http://www.opengis.net/gml pointRep,omitempty"`
	// The attribute "interpolation" specifies the curve interpolation mechanism used for this segment. This mechanism
	// uses the control points and control parameters to determine the position of this curve segment. For an ArcByCenterPoint the interpolation is fixed as "circularArcCenterPointWithRadius".
	Interpolation *string `xml:"interpolation,attr,omitempty"`
	// Since this type describes always a single arc, the attribute is fixed to "1".
	NumArc int `xml:"numArc,attr"`
}

type CircleType struct {
	// The attribute "numDerivativesAtStart" specifies the type of continuity between this curve segment and its predecessor. If this is the first curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtStart *int `xml:"numDerivativesAtStart,attr,omitempty"`
	// The attribute "numDerivativesAtEnd" specifies the type of continuity between this curve segment and its successor. If this is the last curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtEnd *int `xml:"numDerivativesAtEnd,attr,omitempty"`
	// The attribute "numDerivativesInterior" specifies the type of continuity that is guaranteed interior to the curve. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativeInterior *int                    `xml:"numDerivativeInterior,attr,omitempty"`
	PosList               *DirectPositionListType `xml:"http://www.opengis.net/gml posList,omitempty"`
	// Deprecated with GML version 3.1.0. Use "posList" instead.
	Coordinates *CoordinatesType     `xml:"http://www.opengis.net/gml coordinates,omitempty"`
	Pos         []DirectPositionType `xml:"http://www.opengis.net/gml pos"`
	// This property element either references a point via the XLink-attributes or contains the point element. pointProperty
	// is the predefined property which can be used by GML Application Schemas whenever a GML Feature has a property with a value that
	// is substitutable for Point.
	PointProperty []PointPropertyType `xml:"http://www.opengis.net/gml pointProperty"`
	// Deprecated with GML version 3.1.0. Use "pointProperty" instead. Included for backwards compatibility with GML 3.0.0.
	PointRep []PointPropertyType `xml:"http://www.opengis.net/gml pointRep"`
	// The attribute "interpolation" specifies the curve interpolation mechanism used for this segment. This mechanism
	// uses the control points and control parameters to determine the position of this curve segment. For an ArcString the interpolation is fixed as "circularArc3Points".
	Interpolation *string `xml:"interpolation,attr,omitempty"`
	// The number of arcs in the arc string can be explicitly stated in this attribute. The number of control points in the arc string must be 2 * numArc + 1.
	NumArc *int `xml:"numArc,attr,omitempty"`
}

type ClothoidType struct {
	// The attribute "numDerivativesAtStart" specifies the type of continuity between this curve segment and its predecessor. If this is the first curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtStart *int `xml:"numDerivativesAtStart,attr,omitempty"`
	// The attribute "numDerivativesAtEnd" specifies the type of continuity between this curve segment and its successor. If this is the last curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtEnd *int `xml:"numDerivativesAtEnd,attr,omitempty"`
	// The attribute "numDerivativesInterior" specifies the type of continuity that is guaranteed interior to the curve. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativeInterior *int   `xml:"numDerivativeInterior,attr,omitempty"`
	RefLocation           string `xml:"http://www.opengis.net/gml refLocation,omitempty"`
	// The element gives the value for the
	// constant in the Fresnel's integrals.
	ScaleFactor float64 `xml:"http://www.opengis.net/gml scaleFactor,omitempty"`
	// The startParameter is the arc length
	// distance from the inflection point that will be the start
	// point for this curve segment. This shall be lower limit
	// used in the Fresnel integral and is the value of the
	// constructive parameter of this curve segment at its start
	// point. The startParameter can either be positive or
	// negative.
	// NOTE! If 0.0 (zero), lies between the startParameter and
	// the endParameter of the clothoid, then the curve goes
	// through the clothoid's inflection point, and the direction
	// of its radius of curvature, given by the second
	// derivative vector, changes sides with respect to the
	// tangent vector. The term length distance for the
	StartParameter float64 `xml:"http://www.opengis.net/gml startParameter,omitempty"`
	// The endParameter is the arc length
	// distance from the inflection point that will be the end
	// point for this curve segment. This shall be upper limit
	// used in the Fresnel integral and is the value of the
	// constructive parameter of this curve segment at its
	// start point. The startParameter can either be positive
	// or negative.
	EndParameter float64 `xml:"http://www.opengis.net/gml endParameter,omitempty"`
}

type CodeListType struct {
	Value     string  `xml:",chardata"`
	CodeSpace *string `xml:"codeSpace,attr,omitempty"`
}

type CodeOrNullListType struct {
	Value     string  `xml:",chardata"`
	CodeSpace *string `xml:"codeSpace,attr,omitempty"`
}

type CodeType struct {
	Value     string  `xml:",chardata"`
	CodeSpace *string `xml:"codeSpace,attr,omitempty"`
}

type CompositeCurvePropertyType struct {
	CompositeCurve *CompositeCurveType `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	RemoteSchema   string              `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField      string              `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href           string              `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role           string              `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole        string              `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title          string              `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show           string              `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate        string              `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type CompositeCurveType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
	// This element references or contains one curve in the composite curve. The curves are contiguous, the collection of curves is ordered.
	// NOTE: This definition allows for a nested structure, i.e. a CompositeCurve may use, for example, another CompositeCurve as a curve member.
	CurveMember []CurvePropertyType `xml:"http://www.opengis.net/gml curveMember"`
}

type CompositeSolidPropertyType struct {
	CompositeSolid *CompositeSolidType `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	RemoteSchema   string              `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField      string              `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href           string              `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role           string              `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole        string              `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title          string              `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show           string              `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate        string              `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type CompositeSolidType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
	// This element references or contains one solid in the composite solid. The solids are contiguous.
	// NOTE: This definition allows for a nested structure, i.e. a CompositeSolid may use, for example, another CompositeSolid as a member.
	SolidMember []SolidPropertyType `xml:"http://www.opengis.net/gml solidMember"`
}

type CompositeSurfacePropertyType struct {
	CompositeSurface *CompositeSurfaceType `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	RemoteSchema     string                `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField        string                `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href             string                `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role             string                `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole          string                `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title            string                `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show             string                `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate          string                `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type CompositeSurfaceType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
	// This element references or contains one surface in the composite surface. The surfaces are contiguous.
	// NOTE: This definition allows for a nested structure, i.e. a CompositeSurface may use, for example, another CompositeSurface as a member.
	SurfaceMember []SurfacePropertyType `xml:"http://www.opengis.net/gml surfaceMember"`
}

type CompositeValueType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Element which refers to, or contains, a Value.  This version is used in CompositeValues.
	ValueComponent []ValuePropertyType `xml:"http://www.opengis.net/gml valueComponent"`
	// Element which refers to, or contains, a set of homogeneously typed Values.
	ValueComponents *ValueArrayPropertyType `xml:"http://www.opengis.net/gml valueComponents,omitempty"`
}

type CompoundCRSRefType struct {
	CompoundCRS  *CompoundCRSType `xml:"http://www.opengis.net/gml CompoundCRS,omitempty"`
	RemoteSchema string           `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string           `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string           `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string           `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string           `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string           `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string           `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string           `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type CompoundCRSType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alterative identifications of this reference system. The first srsID, if any, is normally the primary identification code, and any others are aliases.
	SrsID []IdentifierType `xml:"http://www.opengis.net/gml srsID"`
	// Comments on or information about this reference system, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
	// Ordered sequence of associations to all the component coordinate reference systems included in this compound coordinate reference system.
	IncludesCRS []CoordinateReferenceSystemRefType `xml:"http://www.opengis.net/gml includesCRS"`
}

type ConcatenatedOperationRefType struct {
	ConcatenatedOperation *ConcatenatedOperationType `xml:"http://www.opengis.net/gml ConcatenatedOperation,omitempty"`
	RemoteSchema          string                     `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField             string                     `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                  string                     `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                  string                     `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole               string                     `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                 string                     `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                  string                     `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate               string                     `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type ConcatenatedOperationType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this coordinate operation. The first coordinateOperationID, if any, is normally the primary identification code, and any others are aliases.
	CoordinateOperationID []IdentifierType `xml:"http://www.opengis.net/gml coordinateOperationID"`
	// Comments on or information about this coordinate operation, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Version of the coordinate transformation (i.e., instantiation due to the stochastic nature of the parameters). Mandatory when describing a transformation, and should not be supplied for a conversion.
	OperationVersion *string `xml:"http://www.opengis.net/gml operationVersion,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
	// Unordered set of estimates of the impact of this coordinate operation on point position accuracy. Gives position error estimates for target coordinates of this coordinate operation, assuming no errors in source coordinates.
	AbstractPositionalAccuracy         []AbstractPositionalAccuracyType         `xml:"http://www.opengis.net/gml _positionalAccuracy"`
	AbsoluteExternalPositionalAccuracy []AbsoluteExternalPositionalAccuracyType `xml:"http://www.opengis.net/gml absoluteExternalPositionalAccuracy"`
	CovarianceMatrix                   []CovarianceMatrixType                   `xml:"http://www.opengis.net/gml covarianceMatrix"`
	RelativeInternalPositionalAccuracy []RelativeInternalPositionalAccuracyType `xml:"http://www.opengis.net/gml relativeInternalPositionalAccuracy"`
	// Association to the source CRS (coordinate reference system) of this coordinate operation.
	SourceCRS *CRSRefType `xml:"http://www.opengis.net/gml sourceCRS,omitempty"`
	// Association to the target CRS (coordinate reference system) of this coordinate operation. For constraints on multiplicity of "sourceCRS" and "targetCRS", see UML model of Coordinate Operation package in OGC Abstract Specification topic 2.
	TargetCRS *CRSRefType `xml:"http://www.opengis.net/gml targetCRS,omitempty"`
	// Ordered sequence of associations to the two or more single operations used by this concatenated operation.
	UsesSingleOperation []SingleOperationRefType `xml:"http://www.opengis.net/gml usesSingleOperation"`
}

type ConeType struct {
	// The attribute rows gives the number
	// of rows in the parameter grid.
	Rows int `xml:"http://www.opengis.net/gml rows,omitempty"`
	// The attribute columns gives the number
	// of columns in the parameter grid.
	Columns             int      `xml:"http://www.opengis.net/gml columns,omitempty"`
	Row                 []string `xml:"http://www.opengis.net/gml row"`
	HorizontalCurveType *string  `xml:"horizontalCurveType,attr,omitempty"`
	VerticalCurveType   *string  `xml:"verticalCurveType,attr,omitempty"`
}

type ContainerPropertyType struct {
	Face         *FaceType      `xml:"http://www.opengis.net/gml Face,omitempty"`
	TopoSolid    *TopoSolidType `xml:"http://www.opengis.net/gml TopoSolid,omitempty"`
	RemoteSchema string         `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string         `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string         `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string         `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string         `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string         `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string         `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string         `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type ConventionalUnitType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Informal description of the phenomenon or type of quantity that is measured or observed. For example, "length", "angle", "time", "pressure", or "temperature". When the quantity is the result of an observation or measurement, this term is known as Observable Type or Measurand.
	QuantityType *StringOrRefType `xml:"http://www.opengis.net/gml quantityType,omitempty"`
	// For global understanding of a unit of measure, it is often possible to reference an item in a catalog of units, using a symbol in that catalog. The "codeSpace" attribute in "CodeType" identifies a namespace for the catalog symbol value, and might reference the catalog. The "string" value in "CodeType" contains the value of a symbol that is unique within this catalog namespace. This symbol often appears explicitly in the catalog, but it could be a combination of symbols using a specified algebra of units. For example, the symbol "cm" might indicate that it is the "m" symbol combined with the "c" prefix.
	CatalogSymbol      *CodeType                `xml:"http://www.opengis.net/gml catalogSymbol,omitempty"`
	DerivationUnitTerm []DerivationUnitTermType `xml:"http://www.opengis.net/gml derivationUnitTerm"`
	// This element is included when this unit has an accurate conversion to the preferred unit for this quantity type.
	ConversionToPreferredUnit *ConversionToPreferredUnitType `xml:"http://www.opengis.net/gml conversionToPreferredUnit,omitempty"`
	// This element is included when the correct definition of this unit is unknown, but this unit has a rough or inaccurate conversion to the preferred unit for this quantity type.
	RoughConversionToPreferredUnit *ConversionToPreferredUnitType `xml:"http://www.opengis.net/gml roughConversionToPreferredUnit,omitempty"`
}

type ConversionRefType struct {
	Conversion   *ConversionType `xml:"http://www.opengis.net/gml Conversion,omitempty"`
	RemoteSchema string          `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string          `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string          `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string          `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string          `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string          `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string          `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string          `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type ConversionToPreferredUnitType struct {
	// Reference to a unit of measure definition, usually within the same XML document but possibly outside the XML document which contains this reference. For a reference within the same XML document, the "#" symbol should be used, followed by a text abbreviation of the unit name. However, the "#" symbol may be optional, and still may be interpreted as a reference.
	Uom string `xml:"uom,attr"`
	// Specification of the scale factor by which a value using this unit of measure can be multiplied to obtain the corresponding value using the preferred unit of measure.
	Factor float64 `xml:"http://www.opengis.net/gml factor,omitempty"`
	// Specification of the formula by which a value using this unit of measure can be converted to obtain the corresponding value using the preferred unit of measure.
	Formula *FormulaType `xml:"http://www.opengis.net/gml formula,omitempty"`
}

type ConversionType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this coordinate operation. The first coordinateOperationID, if any, is normally the primary identification code, and any others are aliases.
	CoordinateOperationID []IdentifierType `xml:"http://www.opengis.net/gml coordinateOperationID"`
	// Comments on or information about this coordinate operation, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Version of the coordinate transformation (i.e., instantiation due to the stochastic nature of the parameters). Mandatory when describing a transformation, and should not be supplied for a conversion.
	OperationVersion *string `xml:"http://www.opengis.net/gml operationVersion,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
	// Unordered set of estimates of the impact of this coordinate operation on point position accuracy. Gives position error estimates for target coordinates of this coordinate operation, assuming no errors in source coordinates.
	AbstractPositionalAccuracy         []AbstractPositionalAccuracyType         `xml:"http://www.opengis.net/gml _positionalAccuracy"`
	AbsoluteExternalPositionalAccuracy []AbsoluteExternalPositionalAccuracyType `xml:"http://www.opengis.net/gml absoluteExternalPositionalAccuracy"`
	CovarianceMatrix                   []CovarianceMatrixType                   `xml:"http://www.opengis.net/gml covarianceMatrix"`
	RelativeInternalPositionalAccuracy []RelativeInternalPositionalAccuracyType `xml:"http://www.opengis.net/gml relativeInternalPositionalAccuracy"`
	// Association to the source CRS (coordinate reference system) of this coordinate operation.
	SourceCRS *CRSRefType `xml:"http://www.opengis.net/gml sourceCRS,omitempty"`
	// Association to the target CRS (coordinate reference system) of this coordinate operation. For constraints on multiplicity of "sourceCRS" and "targetCRS", see UML model of Coordinate Operation package in OGC Abstract Specification topic 2.
	TargetCRS *CRSRefType `xml:"http://www.opengis.net/gml targetCRS,omitempty"`
	// Association to the operation method used by this coordinate operation.
	UsesMethod *OperationMethodRefType `xml:"http://www.opengis.net/gml usesMethod,omitempty"`
	// Unordered list of composition associations to the set of parameter values used by this conversion operation.
	UsesValue []ParameterValueType `xml:"http://www.opengis.net/gml usesValue"`
}

type CoordType struct {
	X float64 `xml:"http://www.opengis.net/gml X,omitempty"`
	Y float64 `xml:"http://www.opengis.net/gml Y,omitempty"`
	Z float64 `xml:"http://www.opengis.net/gml Z,omitempty"`
}

type CoordinateOperationRefType struct {
	AbstractCoordinateOperation   *AbstractCoordinateOperationType   `xml:"http://www.opengis.net/gml _CoordinateOperation,omitempty"`
	ConcatenatedOperation         *ConcatenatedOperationType         `xml:"http://www.opengis.net/gml ConcatenatedOperation,omitempty"`
	Conversion                    *ConversionType                    `xml:"http://www.opengis.net/gml Conversion,omitempty"`
	PassThroughOperation          *PassThroughOperationType          `xml:"http://www.opengis.net/gml PassThroughOperation,omitempty"`
	Transformation                *TransformationType                `xml:"http://www.opengis.net/gml Transformation,omitempty"`
	AbstractGeneralConversion     *AbstractGeneralConversionType     `xml:"http://www.opengis.net/gml _GeneralConversion,omitempty"`
	AbstractGeneralTransformation *AbstractGeneralTransformationType `xml:"http://www.opengis.net/gml _GeneralTransformation,omitempty"`
	AbstractOperation             *AbstractCoordinateOperationType   `xml:"http://www.opengis.net/gml _Operation,omitempty"`
	AbstractSingleOperation       *AbstractCoordinateOperationType   `xml:"http://www.opengis.net/gml _SingleOperation,omitempty"`
	RemoteSchema                  string                             `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                     string                             `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                          string                             `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                          string                             `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                       string                             `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                         string                             `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                          string                             `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                       string                             `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type CoordinateReferenceSystemRefType struct {
	// A coordinate reference system consists of an ordered sequence of coordinate system axes that are related to the earth through a datum. A coordinate reference system is defined by one datum and by one coordinate system. Most coordinate reference system do not move relative to the earth, except for engineering coordinate reference systems defined on moving platforms such as cars, ships, aircraft, and spacecraft. For further information, see OGC Abstract Specification Topic 2.
	//
	// Coordinate reference systems are commonly divided into sub-types. The common classification criterion for sub-typing of coordinate reference systems is the way in which they deal with earth curvature. This has a direct effect on the portion of the earth's surface that can be covered by that type of CRS with an acceptable degree of error. The exception to the rule is the subtype "Temporal" which has been added by analogy.
	AbstractCoordinateReferenceSystem *AbstractReferenceSystemType   `xml:"http://www.opengis.net/gml _CoordinateReferenceSystem,omitempty"`
	DerivedCRS                        *DerivedCRSType                `xml:"http://www.opengis.net/gml DerivedCRS,omitempty"`
	EngineeringCRS                    *EngineeringCRSType            `xml:"http://www.opengis.net/gml EngineeringCRS,omitempty"`
	GeocentricCRS                     *GeocentricCRSType             `xml:"http://www.opengis.net/gml GeocentricCRS,omitempty"`
	GeographicCRS                     *GeographicCRSType             `xml:"http://www.opengis.net/gml GeographicCRS,omitempty"`
	ImageCRS                          *ImageCRSType                  `xml:"http://www.opengis.net/gml ImageCRS,omitempty"`
	ProjectedCRS                      *ProjectedCRSType              `xml:"http://www.opengis.net/gml ProjectedCRS,omitempty"`
	TemporalCRS                       *TemporalCRSType               `xml:"http://www.opengis.net/gml TemporalCRS,omitempty"`
	VerticalCRS                       *VerticalCRSType               `xml:"http://www.opengis.net/gml VerticalCRS,omitempty"`
	AbstractGeneralDerivedCRS         *AbstractGeneralDerivedCRSType `xml:"http://www.opengis.net/gml _GeneralDerivedCRS,omitempty"`
	RemoteSchema                      string                         `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                         string                         `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                              string                         `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                              string                         `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                           string                         `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                             string                         `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                              string                         `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                           string                         `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type CoordinateSystemAxisBaseType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
}

type CoordinateSystemAxisRefType struct {
	CoordinateSystemAxis *CoordinateSystemAxisType `xml:"http://www.opengis.net/gml CoordinateSystemAxis,omitempty"`
	RemoteSchema         string                    `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField            string                    `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                 string                    `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                 string                    `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole              string                    `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                string                    `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                 string                    `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate              string                    `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type CoordinateSystemAxisType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this coordinate system axis. The first axisID, if any, is normally the primary identification code, and any others are aliases.
	AxisID []IdentifierType `xml:"http://www.opengis.net/gml axisID"`
	// Comments on or information about this coordinate system axis, including data source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// The abbreviation used for this coordinate system axis. This abbreviation can be used to identify the ordinates in a coordinate tuple. Examples are X and Y. The codeSpace attribute can reference a source of more information on a set of standardized abbreviations, or on this abbreviation.
	AxisAbbrev *CodeType `xml:"http://www.opengis.net/gml axisAbbrev,omitempty"`
	// Direction of this coordinate system axis (or in the case of Cartesian projected coordinates, the direction of this coordinate system axis at the origin). Examples: north or south, east or west, up or down. Within any set of coordinate system axes, only one of each pair of terms can be used. For earth-fixed CRSs, this direction is often approximate and intended to provide a human interpretable meaning to the axis. When a geodetic datum is used, the precise directions of the axes may therefore vary slightly from this approximate direction. Note that an EngineeringCRS can include specific descriptions of the directions of its coordinate system axes. For example, the path of a linear CRS axis can be referenced in another document, such as referencing a GML feature that references or includes a curve geometry. The codeSpace attribute can reference a source of more information on a set of standardized directions, or on this direction.
	AxisDirection *CodeType `xml:"http://www.opengis.net/gml axisDirection,omitempty"`
	Uom           string    `xml:"http://www.opengis.net/gml uom,attr,omitempty"`
}

type CoordinateSystemRefType struct {
	AbstractCoordinateSystem *AbstractCoordinateSystemType `xml:"http://www.opengis.net/gml _CoordinateSystem,omitempty"`
	CartesianCS              *CartesianCSType              `xml:"http://www.opengis.net/gml CartesianCS,omitempty"`
	CylindricalCS            *CylindricalCSType            `xml:"http://www.opengis.net/gml CylindricalCS,omitempty"`
	EllipsoidalCS            *EllipsoidalCSType            `xml:"http://www.opengis.net/gml EllipsoidalCS,omitempty"`
	LinearCS                 *LinearCSType                 `xml:"http://www.opengis.net/gml LinearCS,omitempty"`
	ObliqueCartesianCS       *ObliqueCartesianCSType       `xml:"http://www.opengis.net/gml ObliqueCartesianCS,omitempty"`
	PolarCS                  *PolarCSType                  `xml:"http://www.opengis.net/gml PolarCS,omitempty"`
	SphericalCS              *SphericalCSType              `xml:"http://www.opengis.net/gml SphericalCS,omitempty"`
	TemporalCS               *TemporalCSType               `xml:"http://www.opengis.net/gml TemporalCS,omitempty"`
	UserDefinedCS            *UserDefinedCSType            `xml:"http://www.opengis.net/gml UserDefinedCS,omitempty"`
	VerticalCS               *VerticalCSType               `xml:"http://www.opengis.net/gml VerticalCS,omitempty"`
	RemoteSchema             string                        `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                string                        `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                     string                        `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                     string                        `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                  string                        `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                    string                        `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                     string                        `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                  string                        `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type CoordinatesType struct {
	Value   string  `xml:",chardata"`
	Decimal *string `xml:"decimal,attr,omitempty"`
	Cs      *string `xml:"cs,attr,omitempty"`
	Ts      *string `xml:"ts,attr,omitempty"`
}

type CountPropertyType struct {
	// This abstract element is the head of a substitutionGroup hierararchy which may contain either simpleContent or complexContent elements.  It is used to assert the model position of "class" elements declared in other GML schemas.
	AbstractObject                    *string                                `xml:"http://www.opengis.net/gml _Object,omitempty"`
	Array                             *ArrayType                             `xml:"http://www.opengis.net/gml Array,omitempty"`
	Bag                               *BagType                               `xml:"http://www.opengis.net/gml Bag,omitempty"`
	BaseUnit                          *BaseUnitType                          `xml:"http://www.opengis.net/gml BaseUnit,omitempty"`
	CartesianCS                       *CartesianCSType                       `xml:"http://www.opengis.net/gml CartesianCS,omitempty"`
	CompositeCurve                    *CompositeCurveType                    `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	CompositeSolid                    *CompositeSolidType                    `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	CompositeSurface                  *CompositeSurfaceType                  `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	CompoundCRS                       *CompoundCRSType                       `xml:"http://www.opengis.net/gml CompoundCRS,omitempty"`
	ConcatenatedOperation             *ConcatenatedOperationType             `xml:"http://www.opengis.net/gml ConcatenatedOperation,omitempty"`
	ConventionalUnit                  *ConventionalUnitType                  `xml:"http://www.opengis.net/gml ConventionalUnit,omitempty"`
	Conversion                        *ConversionType                        `xml:"http://www.opengis.net/gml Conversion,omitempty"`
	CoordinateSystemAxis              *CoordinateSystemAxisType              `xml:"http://www.opengis.net/gml CoordinateSystemAxis,omitempty"`
	Curve                             *CurveType                             `xml:"http://www.opengis.net/gml Curve,omitempty"`
	CylindricalCS                     *CylindricalCSType                     `xml:"http://www.opengis.net/gml CylindricalCS,omitempty"`
	Definition                        *DefinitionType                        `xml:"http://www.opengis.net/gml Definition,omitempty"`
	DefinitionCollection              *DictionaryType                        `xml:"http://www.opengis.net/gml DefinitionCollection,omitempty"`
	DefinitionProxy                   *DefinitionProxyType                   `xml:"http://www.opengis.net/gml DefinitionProxy,omitempty"`
	DerivedCRS                        *DerivedCRSType                        `xml:"http://www.opengis.net/gml DerivedCRS,omitempty"`
	DerivedUnit                       *DerivedUnitType                       `xml:"http://www.opengis.net/gml DerivedUnit,omitempty"`
	Dictionary                        *DictionaryType                        `xml:"http://www.opengis.net/gml Dictionary,omitempty"`
	DirectedObservation               *DirectedObservationType               `xml:"http://www.opengis.net/gml DirectedObservation,omitempty"`
	DirectedObservationAtDistance     *DirectedObservationAtDistanceType     `xml:"http://www.opengis.net/gml DirectedObservationAtDistance,omitempty"`
	Edge                              *EdgeType                              `xml:"http://www.opengis.net/gml Edge,omitempty"`
	Ellipsoid                         *EllipsoidType                         `xml:"http://www.opengis.net/gml Ellipsoid,omitempty"`
	EllipsoidalCS                     *EllipsoidalCSType                     `xml:"http://www.opengis.net/gml EllipsoidalCS,omitempty"`
	EngineeringCRS                    *EngineeringCRSType                    `xml:"http://www.opengis.net/gml EngineeringCRS,omitempty"`
	EngineeringDatum                  *EngineeringDatumType                  `xml:"http://www.opengis.net/gml EngineeringDatum,omitempty"`
	Face                              *FaceType                              `xml:"http://www.opengis.net/gml Face,omitempty"`
	FeatureCollection                 *FeatureCollectionType                 `xml:"http://www.opengis.net/gml FeatureCollection,omitempty"`
	FeatureStyle                      *FeatureStyleType                      `xml:"http://www.opengis.net/gml FeatureStyle,omitempty"`
	GenericMetaData                   *GenericMetaDataType                   `xml:"http://www.opengis.net/gml GenericMetaData,omitempty"`
	GeocentricCRS                     *GeocentricCRSType                     `xml:"http://www.opengis.net/gml GeocentricCRS,omitempty"`
	GeodeticDatum                     *GeodeticDatumType                     `xml:"http://www.opengis.net/gml GeodeticDatum,omitempty"`
	GeographicCRS                     *GeographicCRSType                     `xml:"http://www.opengis.net/gml GeographicCRS,omitempty"`
	GeometricComplex                  *GeometricComplexType                  `xml:"http://www.opengis.net/gml GeometricComplex,omitempty"`
	GeometryStyle                     *GeometryStyleType                     `xml:"http://www.opengis.net/gml GeometryStyle,omitempty"`
	GraphStyle                        *GraphStyleType                        `xml:"http://www.opengis.net/gml GraphStyle,omitempty"`
	Grid                              *GridType                              `xml:"http://www.opengis.net/gml Grid,omitempty"`
	GridCoverage                      *GridCoverageType                      `xml:"http://www.opengis.net/gml GridCoverage,omitempty"`
	ImageCRS                          *ImageCRSType                          `xml:"http://www.opengis.net/gml ImageCRS,omitempty"`
	ImageDatum                        *ImageDatumType                        `xml:"http://www.opengis.net/gml ImageDatum,omitempty"`
	LabelStyle                        *LabelStyleType                        `xml:"http://www.opengis.net/gml LabelStyle,omitempty"`
	LineString                        *LineStringType                        `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearCS                          *LinearCSType                          `xml:"http://www.opengis.net/gml LinearCS,omitempty"`
	LinearRing                        *LinearRingType                        `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MovingObjectStatus                *MovingObjectStatusType                `xml:"http://www.opengis.net/gml MovingObjectStatus,omitempty"`
	MultiCurve                        *MultiCurveType                        `xml:"http://www.opengis.net/gml MultiCurve,omitempty"`
	MultiCurveCoverage                *MultiCurveCoverageType                `xml:"http://www.opengis.net/gml MultiCurveCoverage,omitempty"`
	MultiGeometry                     *MultiGeometryType                     `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString                   *MultiLineStringType                   `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                        *MultiPointType                        `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPointCoverage                *MultiPointCoverageType                `xml:"http://www.opengis.net/gml MultiPointCoverage,omitempty"`
	MultiPolygon                      *MultiPolygonType                      `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	MultiSolid                        *MultiSolidType                        `xml:"http://www.opengis.net/gml MultiSolid,omitempty"`
	MultiSolidCoverage                *MultiSolidCoverageType                `xml:"http://www.opengis.net/gml MultiSolidCoverage,omitempty"`
	MultiSurface                      *MultiSurfaceType                      `xml:"http://www.opengis.net/gml MultiSurface,omitempty"`
	MultiSurfaceCoverage              *MultiSurfaceCoverageType              `xml:"http://www.opengis.net/gml MultiSurfaceCoverage,omitempty"`
	Node                              *NodeType                              `xml:"http://www.opengis.net/gml Node,omitempty"`
	ObliqueCartesianCS                *ObliqueCartesianCSType                `xml:"http://www.opengis.net/gml ObliqueCartesianCS,omitempty"`
	Observation                       *ObservationType                       `xml:"http://www.opengis.net/gml Observation,omitempty"`
	OperationMethod                   *OperationMethodType                   `xml:"http://www.opengis.net/gml OperationMethod,omitempty"`
	OperationParameter                *OperationParameterType                `xml:"http://www.opengis.net/gml OperationParameter,omitempty"`
	OperationParameterGroup           *OperationParameterGroupType           `xml:"http://www.opengis.net/gml OperationParameterGroup,omitempty"`
	OrientableCurve                   *OrientableCurveType                   `xml:"http://www.opengis.net/gml OrientableCurve,omitempty"`
	OrientableSurface                 *OrientableSurfaceType                 `xml:"http://www.opengis.net/gml OrientableSurface,omitempty"`
	PassThroughOperation              *PassThroughOperationType              `xml:"http://www.opengis.net/gml PassThroughOperation,omitempty"`
	Point                             *PointType                             `xml:"http://www.opengis.net/gml Point,omitempty"`
	PolarCS                           *PolarCSType                           `xml:"http://www.opengis.net/gml PolarCS,omitempty"`
	Polygon                           *PolygonType                           `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	PolyhedralSurface                 *PolyhedralSurfaceType                 `xml:"http://www.opengis.net/gml PolyhedralSurface,omitempty"`
	PrimeMeridian                     *PrimeMeridianType                     `xml:"http://www.opengis.net/gml PrimeMeridian,omitempty"`
	ProjectedCRS                      *ProjectedCRSType                      `xml:"http://www.opengis.net/gml ProjectedCRS,omitempty"`
	RectifiedGrid                     *RectifiedGridType                     `xml:"http://www.opengis.net/gml RectifiedGrid,omitempty"`
	RectifiedGridCoverage             *RectifiedGridCoverageType             `xml:"http://www.opengis.net/gml RectifiedGridCoverage,omitempty"`
	Ring                              *RingType                              `xml:"http://www.opengis.net/gml Ring,omitempty"`
	Solid                             *SolidType                             `xml:"http://www.opengis.net/gml Solid,omitempty"`
	SphericalCS                       *SphericalCSType                       `xml:"http://www.opengis.net/gml SphericalCS,omitempty"`
	Style                             *StyleType                             `xml:"http://www.opengis.net/gml Style,omitempty"`
	Surface                           *SurfaceType                           `xml:"http://www.opengis.net/gml Surface,omitempty"`
	TemporalCRS                       *TemporalCRSType                       `xml:"http://www.opengis.net/gml TemporalCRS,omitempty"`
	TemporalCS                        *TemporalCSType                        `xml:"http://www.opengis.net/gml TemporalCS,omitempty"`
	TemporalDatum                     *TemporalDatumType                     `xml:"http://www.opengis.net/gml TemporalDatum,omitempty"`
	TimeCalendar                      *TimeCalendarType                      `xml:"http://www.opengis.net/gml TimeCalendar,omitempty"`
	TimeCalendarEra                   *TimeCalendarEraType                   `xml:"http://www.opengis.net/gml TimeCalendarEra,omitempty"`
	TimeClock                         *TimeClockType                         `xml:"http://www.opengis.net/gml TimeClock,omitempty"`
	TimeCoordinateSystem              *TimeCoordinateSystemType              `xml:"http://www.opengis.net/gml TimeCoordinateSystem,omitempty"`
	TimeEdge                          *TimeEdgeType                          `xml:"http://www.opengis.net/gml TimeEdge,omitempty"`
	TimeInstant                       *TimeInstantType                       `xml:"http://www.opengis.net/gml TimeInstant,omitempty"`
	TimeNode                          *TimeNodeType                          `xml:"http://www.opengis.net/gml TimeNode,omitempty"`
	TimeOrdinalReferenceSystem        *TimeOrdinalReferenceSystemType        `xml:"http://www.opengis.net/gml TimeOrdinalReferenceSystem,omitempty"`
	TimePeriod                        *TimePeriodType                        `xml:"http://www.opengis.net/gml TimePeriod,omitempty"`
	TimeTopologyComplex               *TimeTopologyComplexType               `xml:"http://www.opengis.net/gml TimeTopologyComplex,omitempty"`
	Tin                               *TinType                               `xml:"http://www.opengis.net/gml Tin,omitempty"`
	TopoComplex                       *TopoComplexType                       `xml:"http://www.opengis.net/gml TopoComplex,omitempty"`
	TopoSolid                         *TopoSolidType                         `xml:"http://www.opengis.net/gml TopoSolid,omitempty"`
	TopologyStyle                     *TopologyStyleType                     `xml:"http://www.opengis.net/gml TopologyStyle,omitempty"`
	Transformation                    *TransformationType                    `xml:"http://www.opengis.net/gml Transformation,omitempty"`
	TriangulatedSurface               *TriangulatedSurfaceType               `xml:"http://www.opengis.net/gml TriangulatedSurface,omitempty"`
	UnitDefinition                    *UnitDefinitionType                    `xml:"http://www.opengis.net/gml UnitDefinition,omitempty"`
	UserDefinedCS                     *UserDefinedCSType                     `xml:"http://www.opengis.net/gml UserDefinedCS,omitempty"`
	VerticalCRS                       *VerticalCRSType                       `xml:"http://www.opengis.net/gml VerticalCRS,omitempty"`
	VerticalCS                        *VerticalCSType                        `xml:"http://www.opengis.net/gml VerticalCS,omitempty"`
	VerticalDatum                     *VerticalDatumType                     `xml:"http://www.opengis.net/gml VerticalDatum,omitempty"`
	AbstractCRS                       *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _CRS,omitempty"`
	AbstractContinuousCoverage        *AbstractContinuousCoverageType        `xml:"http://www.opengis.net/gml _ContinuousCoverage,omitempty"`
	AbstractCoordinateOperation       *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _CoordinateOperation,omitempty"`
	AbstractCoordinateReferenceSystem *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _CoordinateReferenceSystem,omitempty"`
	AbstractCoordinateSystem          *AbstractCoordinateSystemType          `xml:"http://www.opengis.net/gml _CoordinateSystem,omitempty"`
	AbstractCoverage                  *AbstractCoverageType                  `xml:"http://www.opengis.net/gml _Coverage,omitempty"`
	AbstractCurve                     *AbstractCurveType                     `xml:"http://www.opengis.net/gml _Curve,omitempty"`
	AbstractDatum                     *AbstractDatumType                     `xml:"http://www.opengis.net/gml _Datum,omitempty"`
	AbstractDiscreteCoverage          *AbstractDiscreteCoverageType          `xml:"http://www.opengis.net/gml _DiscreteCoverage,omitempty"`
	AbstractFeature                   *AbstractFeatureType                   `xml:"http://www.opengis.net/gml _Feature,omitempty"`
	AbstractFeatureCollection         *AbstractFeatureCollectionType         `xml:"http://www.opengis.net/gml _FeatureCollection,omitempty"`
	AbstractGML                       *AbstractGMLType                       `xml:"http://www.opengis.net/gml _GML,omitempty"`
	AbstractGeneralConversion         *AbstractGeneralConversionType         `xml:"http://www.opengis.net/gml _GeneralConversion,omitempty"`
	AbstractGeneralDerivedCRS         *AbstractGeneralDerivedCRSType         `xml:"http://www.opengis.net/gml _GeneralDerivedCRS,omitempty"`
	AbstractGeneralOperationParameter *AbstractGeneralOperationParameterType `xml:"http://www.opengis.net/gml _GeneralOperationParameter,omitempty"`
	AbstractGeneralTransformation     *AbstractGeneralTransformationType     `xml:"http://www.opengis.net/gml _GeneralTransformation,omitempty"`
	AbstractGeometricAggregate        *AbstractGeometricAggregateType        `xml:"http://www.opengis.net/gml _GeometricAggregate,omitempty"`
	AbstractGeometricPrimitive        *AbstractGeometricPrimitiveType        `xml:"http://www.opengis.net/gml _GeometricPrimitive,omitempty"`
	AbstractGeometry                  *AbstractGeometryType                  `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	AbstractImplicitGeometry          *AbstractGeometryType                  `xml:"http://www.opengis.net/gml _ImplicitGeometry,omitempty"`
	AbstractMetaData                  *AbstractMetaDataType                  `xml:"http://www.opengis.net/gml _MetaData,omitempty"`
	AbstractOperation                 *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _Operation,omitempty"`
	AbstractReferenceSystem           *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _ReferenceSystem,omitempty"`
	AbstractRing                      *AbstractRingType                      `xml:"http://www.opengis.net/gml _Ring,omitempty"`
	AbstractSingleOperation           *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _SingleOperation,omitempty"`
	AbstractSolid                     *AbstractSolidType                     `xml:"http://www.opengis.net/gml _Solid,omitempty"`
	AbstractStyle                     *AbstractStyleType                     `xml:"http://www.opengis.net/gml _Style,omitempty"`
	AbstractSurface                   *AbstractSurfaceType                   `xml:"http://www.opengis.net/gml _Surface,omitempty"`
	AbstractTimeComplex               *AbstractTimeComplexType               `xml:"http://www.opengis.net/gml _TimeComplex,omitempty"`
	AbstractTimeGeometricPrimitive    *AbstractTimeGeometricPrimitiveType    `xml:"http://www.opengis.net/gml _TimeGeometricPrimitive,omitempty"`
	AbstractTimeObject                *AbstractTimeObjectType                `xml:"http://www.opengis.net/gml _TimeObject,omitempty"`
	AbstractTimePrimitive             *AbstractTimePrimitiveType             `xml:"http://www.opengis.net/gml _TimePrimitive,omitempty"`
	AbstractTimeReferenceSystem       *AbstractTimeReferenceSystemType       `xml:"http://www.opengis.net/gml _TimeReferenceSystem,omitempty"`
	AbstractTimeSlice                 *AbstractTimeSliceType                 `xml:"http://www.opengis.net/gml _TimeSlice,omitempty"`
	AbstractTimeTopologyPrimitive     *AbstractTimeTopologyPrimitiveType     `xml:"http://www.opengis.net/gml _TimeTopologyPrimitive,omitempty"`
	AbstractTopoPrimitive             *AbstractTopoPrimitiveType             `xml:"http://www.opengis.net/gml _TopoPrimitive,omitempty"`
	AbstractTopology                  *AbstractTopologyType                  `xml:"http://www.opengis.net/gml _Topology,omitempty"`
	Null                              *string                                `xml:"http://www.opengis.net/gml Null,omitempty"`
	// Aggregate value built using the Composite pattern.
	CompositeValue *CompositeValueType `xml:"http://www.opengis.net/gml CompositeValue,omitempty"`
	ValueArray     *ValueArrayType     `xml:"http://www.opengis.net/gml ValueArray,omitempty"`
	// A value from two-valued logic, using the XML Schema boolean type.  An instance may take the values {true, false, 1, 0}.
	Boolean *bool `xml:"http://www.opengis.net/gml Boolean,omitempty"`
	// A term representing a classification.  It has an optional XML attribute codeSpace, whose value is a URI which identifies a dictionary, codelist or authority for the term.
	Category *CodeType `xml:"http://www.opengis.net/gml Category,omitempty"`
	// A numeric value with a scale.  The content of the element is an amount using the XML Schema type double which permits decimal or scientific notation.  An XML attribute uom (unit of measure) is required, whose value is a URI which identifies the definition of the scale or units by which the numeric value must be multiplied.
	Quantity *MeasureType `xml:"http://www.opengis.net/gml Quantity,omitempty"`
	// An integer representing a frequency of occurrence.
	Count *int `xml:"http://www.opengis.net/gml Count,omitempty"`
	// XML List based on XML Schema boolean type.  An element of this type contains a space-separated list of boolean values {0,1,true,false}
	BooleanList *string `xml:"http://www.opengis.net/gml BooleanList,omitempty"`
	// A space-separated list of terms or nulls.  A single XML attribute codeSpace may be provided, which authorises all the terms in the list.
	CategoryList *CodeOrNullListType `xml:"http://www.opengis.net/gml CategoryList,omitempty"`
	// A space separated list of amounts or nulls.  The amounts use the XML Schema type double.  A single XML attribute uom (unit of measure) is required, whose value is a URI which identifies the definition of the scale or units by which all the amounts in the list must be multiplied.
	QuantityList *MeasureOrNullListType `xml:"http://www.opengis.net/gml QuantityList,omitempty"`
	// A space-separated list of integers or nulls.
	CountList *string `xml:"http://www.opengis.net/gml CountList,omitempty"`
	// Utility element to store a 2-point range of ordinal values. If one member is a null, then this is a single ended interval.
	CategoryExtent *CategoryExtentType `xml:"http://www.opengis.net/gml CategoryExtent,omitempty"`
	// Utility element to store a 2-point range of numeric values. If one member is a null, then this is a single ended interval.
	QuantityExtent *QuantityExtentType `xml:"http://www.opengis.net/gml QuantityExtent,omitempty"`
	// Utility element to store a 2-point range of frequency values. If one member is a null, then this is a single ended interval.
	CountExtent  *string `xml:"http://www.opengis.net/gml CountExtent,omitempty"`
	RemoteSchema string  `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string  `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string  `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string  `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string  `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string  `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string  `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string  `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type CovarianceElementType struct {
	// Row number of this covariance element value.
	RowIndex *int `xml:"http://www.opengis.net/gml rowIndex,omitempty"`
	// Column number of this covariance element value.
	ColumnIndex *int `xml:"http://www.opengis.net/gml columnIndex,omitempty"`
	// Value of covariance matrix element.
	Covariance *float64 `xml:"http://www.opengis.net/gml covariance,omitempty"`
}

type CovarianceMatrixType struct {
	// A description of the position accuracy parameter(s) provided.
	MeasureDescription *CodeType `xml:"http://www.opengis.net/gml measureDescription,omitempty"`
	// Ordered sequence of units of measure, corresponding to the row and column index numbers of the covariance matrix, starting with row and column 1 and ending with row/column N. Each unit of measure is for the ordinate reflected in the relevant row and column of the covariance matrix.
	UnitOfMeasure []UnitOfMeasureType `xml:"http://www.opengis.net/gml unitOfMeasure"`
	// Unordered set of elements in this covariance matrix. Because the covariance matrix is symmetrical, only the elements in the upper or lower diagonal part (including the main diagonal) of the matrix need to be specified. Any zero valued covariance elements can be omitted.
	IncludesElement []CovarianceElementType `xml:"http://www.opengis.net/gml includesElement"`
}

type CoverageFunctionType struct {
	// Description of a rule for associating members from the domainSet with members of the rangeSet.
	MappingRule  *StringOrRefType  `xml:"http://www.opengis.net/gml MappingRule,omitempty"`
	GridFunction *GridFunctionType `xml:"http://www.opengis.net/gml GridFunction,omitempty"`
	IndexMap     *IndexMapType     `xml:"http://www.opengis.net/gml IndexMap,omitempty"`
}

type CubicSplineType struct {
	// The attribute "numDerivativesAtStart" specifies the type of continuity between this curve segment and its predecessor. If this is the first curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtStart *int `xml:"numDerivativesAtStart,attr,omitempty"`
	// The attribute "numDerivativesAtEnd" specifies the type of continuity between this curve segment and its successor. If this is the last curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtEnd *int `xml:"numDerivativesAtEnd,attr,omitempty"`
	// The attribute "numDerivativesInterior" specifies the type of continuity that is guaranteed interior to the curve. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativeInterior *int `xml:"numDerivativeInterior,attr,omitempty"`
	// "vectorAtStart" is the unit tangent vector at the start point of the spline.
	VectorAtStart *VectorType `xml:"http://www.opengis.net/gml vectorAtStart,omitempty"`
	// "vectorAtEnd" is the unit tangent vector at the end point of the spline.
	VectorAtEnd *VectorType             `xml:"http://www.opengis.net/gml vectorAtEnd,omitempty"`
	PosList     *DirectPositionListType `xml:"http://www.opengis.net/gml posList,omitempty"`
	// Deprecated with GML version 3.1.0. Use "posList" instead.
	Coordinates *CoordinatesType     `xml:"http://www.opengis.net/gml coordinates,omitempty"`
	Pos         []DirectPositionType `xml:"http://www.opengis.net/gml pos"`
	// This property element either references a point via the XLink-attributes or contains the point element. pointProperty
	// is the predefined property which can be used by GML Application Schemas whenever a GML Feature has a property with a value that
	// is substitutable for Point.
	PointProperty []PointPropertyType `xml:"http://www.opengis.net/gml pointProperty"`
	// Deprecated with GML version 3.1.0. Use "pointProperty" instead. Included for backwards compatibility with GML 3.0.0.
	PointRep []PointPropertyType `xml:"http://www.opengis.net/gml pointRep"`
	// The attribute "interpolation" specifies the curve interpolation mechanism used for this segment. This mechanism
	// uses the control points and control parameters to determine the position of this curve segment. For a CubicSpline the interpolation is fixed as "cubicSpline".
	Interpolation *string `xml:"interpolation,attr,omitempty"`
	// The degree for a cubic spline is "3".
	Degree *int `xml:"degree,attr,omitempty"`
}

type CurveArrayPropertyType struct {
	// The "_Curve" element is the abstract head of the substituition group for all (continuous) curve elements.
	AbstractCurve   []AbstractCurveType   `xml:"http://www.opengis.net/gml _Curve"`
	CompositeCurve  []CompositeCurveType  `xml:"http://www.opengis.net/gml CompositeCurve"`
	Curve           []CurveType           `xml:"http://www.opengis.net/gml Curve"`
	LineString      []LineStringType      `xml:"http://www.opengis.net/gml LineString"`
	OrientableCurve []OrientableCurveType `xml:"http://www.opengis.net/gml OrientableCurve"`
}

type CurvePropertyType struct {
	// The "_Curve" element is the abstract head of the substituition group for all (continuous) curve elements.
	AbstractCurve   *AbstractCurveType   `xml:"http://www.opengis.net/gml _Curve,omitempty"`
	CompositeCurve  *CompositeCurveType  `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	Curve           *CurveType           `xml:"http://www.opengis.net/gml Curve,omitempty"`
	LineString      *LineStringType      `xml:"http://www.opengis.net/gml LineString,omitempty"`
	OrientableCurve *OrientableCurveType `xml:"http://www.opengis.net/gml OrientableCurve,omitempty"`
	RemoteSchema    string               `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField       string               `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href            string               `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role            string               `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole         string               `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title           string               `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show            string               `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate         string               `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type CurveSegmentArrayPropertyType struct {
	// The "_CurveSegment" element is the abstract head of the substituition group for all curve segment elements, i.e. continuous segments of the same interpolation mechanism.
	AbstractCurveSegment []AbstractCurveSegmentType `xml:"http://www.opengis.net/gml _CurveSegment"`
	Arc                  []ArcType                  `xml:"http://www.opengis.net/gml Arc"`
	ArcByBulge           []ArcByBulgeType           `xml:"http://www.opengis.net/gml ArcByBulge"`
	ArcByCenterPoint     []ArcByCenterPointType     `xml:"http://www.opengis.net/gml ArcByCenterPoint"`
	ArcString            []ArcStringType            `xml:"http://www.opengis.net/gml ArcString"`
	ArcStringByBulge     []ArcStringByBulgeType     `xml:"http://www.opengis.net/gml ArcStringByBulge"`
	BSpline              []BSplineType              `xml:"http://www.opengis.net/gml BSpline"`
	Bezier               []BezierType               `xml:"http://www.opengis.net/gml Bezier"`
	Circle               []CircleType               `xml:"http://www.opengis.net/gml Circle"`
	CircleByCenterPoint  []CircleByCenterPointType  `xml:"http://www.opengis.net/gml CircleByCenterPoint"`
	Clothoid             []ClothoidType             `xml:"http://www.opengis.net/gml Clothoid"`
	CubicSpline          []CubicSplineType          `xml:"http://www.opengis.net/gml CubicSpline"`
	Geodesic             []GeodesicType             `xml:"http://www.opengis.net/gml Geodesic"`
	GeodesicString       []GeodesicStringType       `xml:"http://www.opengis.net/gml GeodesicString"`
	LineStringSegment    []LineStringSegmentType    `xml:"http://www.opengis.net/gml LineStringSegment"`
	OffsetCurve          []OffsetCurveType          `xml:"http://www.opengis.net/gml OffsetCurve"`
}

type CurveType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
	// This element encapsulates the segments of the curve.
	Segments *CurveSegmentArrayPropertyType `xml:"http://www.opengis.net/gml segments,omitempty"`
}

type CylinderType struct {
	// The attribute rows gives the number
	// of rows in the parameter grid.
	Rows int `xml:"http://www.opengis.net/gml rows,omitempty"`
	// The attribute columns gives the number
	// of columns in the parameter grid.
	Columns             int      `xml:"http://www.opengis.net/gml columns,omitempty"`
	Row                 []string `xml:"http://www.opengis.net/gml row"`
	HorizontalCurveType *string  `xml:"horizontalCurveType,attr,omitempty"`
	VerticalCurveType   *string  `xml:"verticalCurveType,attr,omitempty"`
}

type CylindricalCSRefType struct {
	CylindricalCS *CylindricalCSType `xml:"http://www.opengis.net/gml CylindricalCS,omitempty"`
	RemoteSchema  string             `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField     string             `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href          string             `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role          string             `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole       string             `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title         string             `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show          string             `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate       string             `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type CylindricalCSType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this coordinate system. The first csID, if any, is normally the primary identification code, and any others are aliases.
	CsID []IdentifierType `xml:"http://www.opengis.net/gml csID"`
	// Comments on or information about this coordinate system, including data source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Ordered sequence of associations to the coordinate system axes included in this coordinate system.
	UsesAxis []CoordinateSystemAxisRefType `xml:"http://www.opengis.net/gml usesAxis"`
}

type DMSAngleType struct {
	Degrees        *DegreesType `xml:"http://www.opengis.net/gml degrees,omitempty"`
	DecimalMinutes *float64     `xml:"http://www.opengis.net/gml decimalMinutes,omitempty"`
	Minutes        *int         `xml:"http://www.opengis.net/gml minutes,omitempty"`
	Seconds        *float64     `xml:"http://www.opengis.net/gml seconds,omitempty"`
}

type DataBlockType struct {
	RangeParameters       *RangeParametersType `xml:"http://www.opengis.net/gml rangeParameters,omitempty"`
	TupleList             *CoordinatesType     `xml:"http://www.opengis.net/gml tupleList,omitempty"`
	DoubleOrNullTupleList *string              `xml:"http://www.opengis.net/gml doubleOrNullTupleList,omitempty"`
}

type DatumRefType struct {
	AbstractDatum    *AbstractDatumType    `xml:"http://www.opengis.net/gml _Datum,omitempty"`
	EngineeringDatum *EngineeringDatumType `xml:"http://www.opengis.net/gml EngineeringDatum,omitempty"`
	GeodeticDatum    *GeodeticDatumType    `xml:"http://www.opengis.net/gml GeodeticDatum,omitempty"`
	ImageDatum       *ImageDatumType       `xml:"http://www.opengis.net/gml ImageDatum,omitempty"`
	TemporalDatum    *TemporalDatumType    `xml:"http://www.opengis.net/gml TemporalDatum,omitempty"`
	VerticalDatum    *VerticalDatumType    `xml:"http://www.opengis.net/gml VerticalDatum,omitempty"`
	RemoteSchema     string                `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField        string                `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href             string                `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role             string                `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole          string                `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title            string                `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show             string                `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate          string                `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type DefaultStylePropertyType struct {
	// The value of the top-level property. It is an abstract element. Used as the head element of the substitution group for extensibility purposes.
	AbstractStyle *AbstractStyleType `xml:"http://www.opengis.net/gml _Style,omitempty"`
	Style         *StyleType         `xml:"http://www.opengis.net/gml Style,omitempty"`
	About         *string            `xml:"about,attr,omitempty"`
	RemoteSchema  string             `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField     string             `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href          string             `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role          string             `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole       string             `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title         string             `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show          string             `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate       string             `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type DefinitionProxyType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// A reference to a remote entry in this dictionary, used when this dictionary entry is identified to allow external references to this specific entry. The remote entry referenced can be in a dictionary in the same or different XML document.
	DefinitionRef *ReferenceType `xml:"http://www.opengis.net/gml definitionRef,omitempty"`
}

type DefinitionType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
}

type DegreesType struct {
	Value     int     `xml:",chardata"`
	Direction *string `xml:"direction,attr,omitempty"`
}

type DerivationUnitTermType struct {
	// Reference to a unit of measure definition, usually within the same XML document but possibly outside the XML document which contains this reference. For a reference within the same XML document, the "#" symbol should be used, followed by a text abbreviation of the unit name. However, the "#" symbol may be optional, and still may be interpreted as a reference.
	Uom      string `xml:"uom,attr"`
	Exponent *int   `xml:"exponent,attr,omitempty"`
}

type DerivedCRSRefType struct {
	DerivedCRS   *DerivedCRSType `xml:"http://www.opengis.net/gml DerivedCRS,omitempty"`
	RemoteSchema string          `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string          `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string          `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string          `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string          `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string          `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string          `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string          `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type DerivedCRSType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alterative identifications of this reference system. The first srsID, if any, is normally the primary identification code, and any others are aliases.
	SrsID []IdentifierType `xml:"http://www.opengis.net/gml srsID"`
	// Comments on or information about this reference system, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
	// Association to the coordinate reference system used by this derived CRS.
	BaseCRS *CoordinateReferenceSystemRefType `xml:"http://www.opengis.net/gml baseCRS,omitempty"`
	// Association to the coordinate conversion used to define this derived CRS.
	DefinedByConversion *GeneralConversionRefType `xml:"http://www.opengis.net/gml definedByConversion,omitempty"`
	DerivedCRSType      *DerivedCRSTypeType       `xml:"http://www.opengis.net/gml derivedCRSType,omitempty"`
	// Association to the coordinate system used by this CRS.
	UsesCS *CoordinateSystemRefType `xml:"http://www.opengis.net/gml usesCS,omitempty"`
}

type DerivedCRSTypeType struct {
	Value string `xml:",chardata"`
	// Reference to a source of information specifying the values and meanings of all the allowed string values for this DerivedCRSTypeType.
	CodeSpace string `xml:"codeSpace,attr"`
}

type DerivedUnitType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Informal description of the phenomenon or type of quantity that is measured or observed. For example, "length", "angle", "time", "pressure", or "temperature". When the quantity is the result of an observation or measurement, this term is known as Observable Type or Measurand.
	QuantityType *StringOrRefType `xml:"http://www.opengis.net/gml quantityType,omitempty"`
	// For global understanding of a unit of measure, it is often possible to reference an item in a catalog of units, using a symbol in that catalog. The "codeSpace" attribute in "CodeType" identifies a namespace for the catalog symbol value, and might reference the catalog. The "string" value in "CodeType" contains the value of a symbol that is unique within this catalog namespace. This symbol often appears explicitly in the catalog, but it could be a combination of symbols using a specified algebra of units. For example, the symbol "cm" might indicate that it is the "m" symbol combined with the "c" prefix.
	CatalogSymbol      *CodeType                `xml:"http://www.opengis.net/gml catalogSymbol,omitempty"`
	DerivationUnitTerm []DerivationUnitTermType `xml:"http://www.opengis.net/gml derivationUnitTerm"`
}

type DictionaryEntryType struct {
	// This element in a dictionary entry contains the actual definition.
	Definition                        *DefinitionType                        `xml:"http://www.opengis.net/gml Definition,omitempty"`
	BaseUnit                          *BaseUnitType                          `xml:"http://www.opengis.net/gml BaseUnit,omitempty"`
	CartesianCS                       *CartesianCSType                       `xml:"http://www.opengis.net/gml CartesianCS,omitempty"`
	CompoundCRS                       *CompoundCRSType                       `xml:"http://www.opengis.net/gml CompoundCRS,omitempty"`
	ConcatenatedOperation             *ConcatenatedOperationType             `xml:"http://www.opengis.net/gml ConcatenatedOperation,omitempty"`
	ConventionalUnit                  *ConventionalUnitType                  `xml:"http://www.opengis.net/gml ConventionalUnit,omitempty"`
	Conversion                        *ConversionType                        `xml:"http://www.opengis.net/gml Conversion,omitempty"`
	CoordinateSystemAxis              *CoordinateSystemAxisType              `xml:"http://www.opengis.net/gml CoordinateSystemAxis,omitempty"`
	CylindricalCS                     *CylindricalCSType                     `xml:"http://www.opengis.net/gml CylindricalCS,omitempty"`
	DefinitionCollection              *DictionaryType                        `xml:"http://www.opengis.net/gml DefinitionCollection,omitempty"`
	DefinitionProxy                   *DefinitionProxyType                   `xml:"http://www.opengis.net/gml DefinitionProxy,omitempty"`
	DerivedCRS                        *DerivedCRSType                        `xml:"http://www.opengis.net/gml DerivedCRS,omitempty"`
	DerivedUnit                       *DerivedUnitType                       `xml:"http://www.opengis.net/gml DerivedUnit,omitempty"`
	Dictionary                        *DictionaryType                        `xml:"http://www.opengis.net/gml Dictionary,omitempty"`
	Ellipsoid                         *EllipsoidType                         `xml:"http://www.opengis.net/gml Ellipsoid,omitempty"`
	EllipsoidalCS                     *EllipsoidalCSType                     `xml:"http://www.opengis.net/gml EllipsoidalCS,omitempty"`
	EngineeringCRS                    *EngineeringCRSType                    `xml:"http://www.opengis.net/gml EngineeringCRS,omitempty"`
	EngineeringDatum                  *EngineeringDatumType                  `xml:"http://www.opengis.net/gml EngineeringDatum,omitempty"`
	GeocentricCRS                     *GeocentricCRSType                     `xml:"http://www.opengis.net/gml GeocentricCRS,omitempty"`
	GeodeticDatum                     *GeodeticDatumType                     `xml:"http://www.opengis.net/gml GeodeticDatum,omitempty"`
	GeographicCRS                     *GeographicCRSType                     `xml:"http://www.opengis.net/gml GeographicCRS,omitempty"`
	ImageCRS                          *ImageCRSType                          `xml:"http://www.opengis.net/gml ImageCRS,omitempty"`
	ImageDatum                        *ImageDatumType                        `xml:"http://www.opengis.net/gml ImageDatum,omitempty"`
	LinearCS                          *LinearCSType                          `xml:"http://www.opengis.net/gml LinearCS,omitempty"`
	ObliqueCartesianCS                *ObliqueCartesianCSType                `xml:"http://www.opengis.net/gml ObliqueCartesianCS,omitempty"`
	OperationMethod                   *OperationMethodType                   `xml:"http://www.opengis.net/gml OperationMethod,omitempty"`
	OperationParameter                *OperationParameterType                `xml:"http://www.opengis.net/gml OperationParameter,omitempty"`
	OperationParameterGroup           *OperationParameterGroupType           `xml:"http://www.opengis.net/gml OperationParameterGroup,omitempty"`
	PassThroughOperation              *PassThroughOperationType              `xml:"http://www.opengis.net/gml PassThroughOperation,omitempty"`
	PolarCS                           *PolarCSType                           `xml:"http://www.opengis.net/gml PolarCS,omitempty"`
	PrimeMeridian                     *PrimeMeridianType                     `xml:"http://www.opengis.net/gml PrimeMeridian,omitempty"`
	ProjectedCRS                      *ProjectedCRSType                      `xml:"http://www.opengis.net/gml ProjectedCRS,omitempty"`
	SphericalCS                       *SphericalCSType                       `xml:"http://www.opengis.net/gml SphericalCS,omitempty"`
	TemporalCRS                       *TemporalCRSType                       `xml:"http://www.opengis.net/gml TemporalCRS,omitempty"`
	TemporalCS                        *TemporalCSType                        `xml:"http://www.opengis.net/gml TemporalCS,omitempty"`
	TemporalDatum                     *TemporalDatumType                     `xml:"http://www.opengis.net/gml TemporalDatum,omitempty"`
	TimeCalendar                      *TimeCalendarType                      `xml:"http://www.opengis.net/gml TimeCalendar,omitempty"`
	TimeCalendarEra                   *TimeCalendarEraType                   `xml:"http://www.opengis.net/gml TimeCalendarEra,omitempty"`
	TimeClock                         *TimeClockType                         `xml:"http://www.opengis.net/gml TimeClock,omitempty"`
	TimeCoordinateSystem              *TimeCoordinateSystemType              `xml:"http://www.opengis.net/gml TimeCoordinateSystem,omitempty"`
	TimeOrdinalReferenceSystem        *TimeOrdinalReferenceSystemType        `xml:"http://www.opengis.net/gml TimeOrdinalReferenceSystem,omitempty"`
	Transformation                    *TransformationType                    `xml:"http://www.opengis.net/gml Transformation,omitempty"`
	UnitDefinition                    *UnitDefinitionType                    `xml:"http://www.opengis.net/gml UnitDefinition,omitempty"`
	UserDefinedCS                     *UserDefinedCSType                     `xml:"http://www.opengis.net/gml UserDefinedCS,omitempty"`
	VerticalCRS                       *VerticalCRSType                       `xml:"http://www.opengis.net/gml VerticalCRS,omitempty"`
	VerticalCS                        *VerticalCSType                        `xml:"http://www.opengis.net/gml VerticalCS,omitempty"`
	VerticalDatum                     *VerticalDatumType                     `xml:"http://www.opengis.net/gml VerticalDatum,omitempty"`
	AbstractCRS                       *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _CRS,omitempty"`
	AbstractCoordinateOperation       *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _CoordinateOperation,omitempty"`
	AbstractCoordinateReferenceSystem *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _CoordinateReferenceSystem,omitempty"`
	AbstractCoordinateSystem          *AbstractCoordinateSystemType          `xml:"http://www.opengis.net/gml _CoordinateSystem,omitempty"`
	AbstractDatum                     *AbstractDatumType                     `xml:"http://www.opengis.net/gml _Datum,omitempty"`
	AbstractGeneralConversion         *AbstractGeneralConversionType         `xml:"http://www.opengis.net/gml _GeneralConversion,omitempty"`
	AbstractGeneralDerivedCRS         *AbstractGeneralDerivedCRSType         `xml:"http://www.opengis.net/gml _GeneralDerivedCRS,omitempty"`
	AbstractGeneralOperationParameter *AbstractGeneralOperationParameterType `xml:"http://www.opengis.net/gml _GeneralOperationParameter,omitempty"`
	AbstractGeneralTransformation     *AbstractGeneralTransformationType     `xml:"http://www.opengis.net/gml _GeneralTransformation,omitempty"`
	AbstractOperation                 *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _Operation,omitempty"`
	AbstractReferenceSystem           *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _ReferenceSystem,omitempty"`
	AbstractSingleOperation           *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _SingleOperation,omitempty"`
	AbstractTimeReferenceSystem       *AbstractTimeReferenceSystemType       `xml:"http://www.opengis.net/gml _TimeReferenceSystem,omitempty"`
	RemoteSchema                      string                                 `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                         string                                 `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                              string                                 `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                              string                                 `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                           string                                 `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                             string                                 `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                              string                                 `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                           string                                 `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type DictionaryType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// An entry in this dictionary. The content of an entry can itself be a lower level dictionary or definition collection. This element follows the standard GML property model, so the value may be provided directly or by reference. Note that if the value is provided by reference, this definition does not carry a handle (gml:id) in this context, so does not allow external references to this specific entry in this context. When used in this way the referenced definition will usually be in a dictionary in the same XML document.
	DictionaryEntry  *DictionaryEntryType `xml:"http://www.opengis.net/gml dictionaryEntry,omitempty"`
	DefinitionMember *DictionaryEntryType `xml:"http://www.opengis.net/gml definitionMember,omitempty"`
	// An identified reference to a remote entry in this dictionary, to be used when this entry should be identified to allow external references to this specific entry.
	IndirectEntry *IndirectEntryType `xml:"http://www.opengis.net/gml indirectEntry,omitempty"`
}

type DirectPositionListType struct {
	Value string `xml:",chardata"`
	// "count" allows to specify the number of direct positions in the list. If the attribute count is present then
	// the attribute srsDimension shall be present, too.
	Count *int `xml:"count,attr,omitempty"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
}

type DirectPositionType struct {
	Value string `xml:",chardata"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
}

type DirectedEdgePropertyType struct {
	Edge         *EdgeType `xml:"http://www.opengis.net/gml Edge,omitempty"`
	Orientation  *string   `xml:"orientation,attr,omitempty"`
	RemoteSchema string    `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string    `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string    `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string    `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string    `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string    `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string    `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string    `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type DirectedFacePropertyType struct {
	Face         *FaceType `xml:"http://www.opengis.net/gml Face,omitempty"`
	Orientation  *string   `xml:"orientation,attr,omitempty"`
	RemoteSchema string    `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string    `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string    `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string    `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string    `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string    `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string    `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string    `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type DirectedNodePropertyType struct {
	Node         *NodeType `xml:"http://www.opengis.net/gml Node,omitempty"`
	Orientation  *string   `xml:"orientation,attr,omitempty"`
	RemoteSchema string    `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string    `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string    `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string    `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string    `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string    `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string    `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string    `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type DirectedObservationAtDistanceType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType         `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType         `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType         `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType         `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType         `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType         `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType         `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType         `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType         `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType         `xml:"http://www.opengis.net/gml srsName"`
	Id                      string             `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	BoundedBy               *BoundingShapeType `xml:"http://www.opengis.net/gml boundedBy,omitempty"`
	// deprecated in GML version 3.1
	Location         *LocationPropertyType         `xml:"http://www.opengis.net/gml location,omitempty"`
	PriorityLocation *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml priorityLocation,omitempty"`
	ValidTime        *TimePrimitivePropertyType    `xml:"http://www.opengis.net/gml validTime,omitempty"`
	// This element contains or points to a description of a sensor, instrument or procedure used for the observation
	Using *FeaturePropertyType `xml:"http://www.opengis.net/gml using,omitempty"`
	// This element contains or points to the specimen, region or station which is the object of the observation
	Target  *TargetPropertyType `xml:"http://www.opengis.net/gml target,omitempty"`
	Subject *TargetPropertyType `xml:"http://www.opengis.net/gml subject,omitempty"`
	// The result of the observation: an image, external object, etc
	ResultOf  *AssociationType       `xml:"http://www.opengis.net/gml resultOf,omitempty"`
	Direction *DirectionPropertyType `xml:"http://www.opengis.net/gml direction,omitempty"`
	Distance  *MeasureType           `xml:"http://www.opengis.net/gml distance,omitempty"`
}

type DirectedObservationType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType         `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType         `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType         `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType         `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType         `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType         `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType         `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType         `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType         `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType         `xml:"http://www.opengis.net/gml srsName"`
	Id                      string             `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	BoundedBy               *BoundingShapeType `xml:"http://www.opengis.net/gml boundedBy,omitempty"`
	// deprecated in GML version 3.1
	Location         *LocationPropertyType         `xml:"http://www.opengis.net/gml location,omitempty"`
	PriorityLocation *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml priorityLocation,omitempty"`
	ValidTime        *TimePrimitivePropertyType    `xml:"http://www.opengis.net/gml validTime,omitempty"`
	// This element contains or points to a description of a sensor, instrument or procedure used for the observation
	Using *FeaturePropertyType `xml:"http://www.opengis.net/gml using,omitempty"`
	// This element contains or points to the specimen, region or station which is the object of the observation
	Target  *TargetPropertyType `xml:"http://www.opengis.net/gml target,omitempty"`
	Subject *TargetPropertyType `xml:"http://www.opengis.net/gml subject,omitempty"`
	// The result of the observation: an image, external object, etc
	ResultOf  *AssociationType       `xml:"http://www.opengis.net/gml resultOf,omitempty"`
	Direction *DirectionPropertyType `xml:"http://www.opengis.net/gml direction,omitempty"`
}

type DirectedTopoSolidPropertyType struct {
	TopoSolid    *TopoSolidType `xml:"http://www.opengis.net/gml TopoSolid,omitempty"`
	Orientation  *string        `xml:"orientation,attr,omitempty"`
	RemoteSchema string         `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string         `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string         `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string         `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string         `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string         `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string         `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string         `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type DirectionPropertyType struct {
	DirectionVector  *DirectionVectorType `xml:"http://www.opengis.net/gml DirectionVector,omitempty"`
	CompassPoint     *string              `xml:"http://www.opengis.net/gml CompassPoint,omitempty"`
	DirectionKeyword *CodeType            `xml:"http://www.opengis.net/gml DirectionKeyword,omitempty"`
	DirectionString  *StringOrRefType     `xml:"http://www.opengis.net/gml DirectionString,omitempty"`
	RemoteSchema     string               `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField        string               `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href             string               `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role             string               `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole          string               `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title            string               `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show             string               `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate          string               `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type DirectionVectorType struct {
	Vector          *VectorType `xml:"http://www.opengis.net/gml vector,omitempty"`
	HorizontalAngle *AngleType  `xml:"http://www.opengis.net/gml horizontalAngle,omitempty"`
	VerticalAngle   *AngleType  `xml:"http://www.opengis.net/gml verticalAngle,omitempty"`
}

type DomainSetType struct {
	// The "_Geometry" element is the abstract head of the substituition group for all geometry elements of GML 3. This
	// includes pre-defined and user-defined geometry elements. Any geometry element must be a direct or indirect extension/restriction
	// of AbstractGeometryType and must be directly or indirectly in the substitution group of "_Geometry".
	AbstractGeometry           *AbstractGeometryType           `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	CompositeCurve             *CompositeCurveType             `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	CompositeSolid             *CompositeSolidType             `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	CompositeSurface           *CompositeSurfaceType           `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	Curve                      *CurveType                      `xml:"http://www.opengis.net/gml Curve,omitempty"`
	GeometricComplex           *GeometricComplexType           `xml:"http://www.opengis.net/gml GeometricComplex,omitempty"`
	Grid                       *GridType                       `xml:"http://www.opengis.net/gml Grid,omitempty"`
	LineString                 *LineStringType                 `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearRing                 *LinearRingType                 `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MultiCurve                 *MultiCurveType                 `xml:"http://www.opengis.net/gml MultiCurve,omitempty"`
	MultiGeometry              *MultiGeometryType              `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString            *MultiLineStringType            `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                 *MultiPointType                 `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPolygon               *MultiPolygonType               `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	MultiSolid                 *MultiSolidType                 `xml:"http://www.opengis.net/gml MultiSolid,omitempty"`
	MultiSurface               *MultiSurfaceType               `xml:"http://www.opengis.net/gml MultiSurface,omitempty"`
	OrientableCurve            *OrientableCurveType            `xml:"http://www.opengis.net/gml OrientableCurve,omitempty"`
	OrientableSurface          *OrientableSurfaceType          `xml:"http://www.opengis.net/gml OrientableSurface,omitempty"`
	Point                      *PointType                      `xml:"http://www.opengis.net/gml Point,omitempty"`
	Polygon                    *PolygonType                    `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	PolyhedralSurface          *PolyhedralSurfaceType          `xml:"http://www.opengis.net/gml PolyhedralSurface,omitempty"`
	RectifiedGrid              *RectifiedGridType              `xml:"http://www.opengis.net/gml RectifiedGrid,omitempty"`
	Ring                       *RingType                       `xml:"http://www.opengis.net/gml Ring,omitempty"`
	Solid                      *SolidType                      `xml:"http://www.opengis.net/gml Solid,omitempty"`
	Surface                    *SurfaceType                    `xml:"http://www.opengis.net/gml Surface,omitempty"`
	Tin                        *TinType                        `xml:"http://www.opengis.net/gml Tin,omitempty"`
	TriangulatedSurface        *TriangulatedSurfaceType        `xml:"http://www.opengis.net/gml TriangulatedSurface,omitempty"`
	AbstractCurve              *AbstractCurveType              `xml:"http://www.opengis.net/gml _Curve,omitempty"`
	AbstractGeometricAggregate *AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml _GeometricAggregate,omitempty"`
	AbstractGeometricPrimitive *AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml _GeometricPrimitive,omitempty"`
	AbstractImplicitGeometry   *AbstractGeometryType           `xml:"http://www.opengis.net/gml _ImplicitGeometry,omitempty"`
	AbstractRing               *AbstractRingType               `xml:"http://www.opengis.net/gml _Ring,omitempty"`
	AbstractSolid              *AbstractSolidType              `xml:"http://www.opengis.net/gml _Solid,omitempty"`
	AbstractSurface            *AbstractSurfaceType            `xml:"http://www.opengis.net/gml _Surface,omitempty"`
	// This abstract element acts as the head of the substitution group for temporal primitives and complexes.
	AbstractTimeObject             *AbstractTimeObjectType             `xml:"http://www.opengis.net/gml _TimeObject,omitempty"`
	TimeEdge                       *TimeEdgeType                       `xml:"http://www.opengis.net/gml TimeEdge,omitempty"`
	TimeInstant                    *TimeInstantType                    `xml:"http://www.opengis.net/gml TimeInstant,omitempty"`
	TimeNode                       *TimeNodeType                       `xml:"http://www.opengis.net/gml TimeNode,omitempty"`
	TimePeriod                     *TimePeriodType                     `xml:"http://www.opengis.net/gml TimePeriod,omitempty"`
	TimeTopologyComplex            *TimeTopologyComplexType            `xml:"http://www.opengis.net/gml TimeTopologyComplex,omitempty"`
	AbstractTimeComplex            *AbstractTimeComplexType            `xml:"http://www.opengis.net/gml _TimeComplex,omitempty"`
	AbstractTimeGeometricPrimitive *AbstractTimeGeometricPrimitiveType `xml:"http://www.opengis.net/gml _TimeGeometricPrimitive,omitempty"`
	AbstractTimePrimitive          *AbstractTimePrimitiveType          `xml:"http://www.opengis.net/gml _TimePrimitive,omitempty"`
	AbstractTimeTopologyPrimitive  *AbstractTimeTopologyPrimitiveType  `xml:"http://www.opengis.net/gml _TimeTopologyPrimitive,omitempty"`
	RemoteSchema                   string                              `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                      string                              `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                           string                              `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                           string                              `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                        string                              `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                          string                              `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                           string                              `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                        string                              `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type DynamicFeatureCollectionType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType         `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType         `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType         `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType         `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType         `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType         `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType         `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType         `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType         `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType         `xml:"http://www.opengis.net/gml srsName"`
	Id                      string             `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	BoundedBy               *BoundingShapeType `xml:"http://www.opengis.net/gml boundedBy,omitempty"`
	// deprecated in GML version 3.1
	Location         *LocationPropertyType         `xml:"http://www.opengis.net/gml location,omitempty"`
	PriorityLocation *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml priorityLocation,omitempty"`
	FeatureMember    []FeaturePropertyType         `xml:"http://www.opengis.net/gml featureMember"`
	FeatureMembers   *FeatureArrayPropertyType     `xml:"http://www.opengis.net/gml featureMembers,omitempty"`
}

type DynamicFeatureType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType         `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType         `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType         `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType         `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType         `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType         `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType         `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType         `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType         `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType         `xml:"http://www.opengis.net/gml srsName"`
	Id                      string             `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	BoundedBy               *BoundingShapeType `xml:"http://www.opengis.net/gml boundedBy,omitempty"`
	// deprecated in GML version 3.1
	Location         *LocationPropertyType         `xml:"http://www.opengis.net/gml location,omitempty"`
	PriorityLocation *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml priorityLocation,omitempty"`
}

type EdgeType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType                 `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType                 `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType                 `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType                 `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType                 `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType                 `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType                 `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType                 `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType                 `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType                 `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	Isolated                []IsolatedPropertyType     `xml:"http://www.opengis.net/gml isolated"`
	Container               *ContainerPropertyType     `xml:"http://www.opengis.net/gml container,omitempty"`
	DirectedNode            []DirectedNodePropertyType `xml:"http://www.opengis.net/gml directedNode"`
	DirectedFace            []DirectedFacePropertyType `xml:"http://www.opengis.net/gml directedFace"`
	// This property element either references a curve via the XLink-attributes or contains the curve element. curveProperty is the
	// predefined property which can be used by GML Application Schemas whenever a GML Feature has a property with a value that is
	// substitutable for _Curve.
	CurveProperty *CurvePropertyType `xml:"http://www.opengis.net/gml curveProperty,omitempty"`
}

type EllipsoidBaseType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
}

type EllipsoidRefType struct {
	Ellipsoid    *EllipsoidType `xml:"http://www.opengis.net/gml Ellipsoid,omitempty"`
	RemoteSchema string         `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string         `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string         `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string         `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string         `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string         `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string         `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string         `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type EllipsoidType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this ellipsoid. The first ellipsoidID, if any, is normally the primary identification code, and any others are aliases.
	EllipsoidID []IdentifierType `xml:"http://www.opengis.net/gml ellipsoidID"`
	// Comments on or information about this ellipsoid, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Length of the semi-major axis of the ellipsoid, with its units. Uses the MeasureType with the restriction that the unit of measure referenced by uom must be suitable for a length, such as metres or feet.
	SemiMajorAxis           *MeasureType                 `xml:"http://www.opengis.net/gml semiMajorAxis,omitempty"`
	SecondDefiningParameter *SecondDefiningParameterType `xml:"http://www.opengis.net/gml secondDefiningParameter,omitempty"`
}

type EllipsoidalCSRefType struct {
	EllipsoidalCS *EllipsoidalCSType `xml:"http://www.opengis.net/gml EllipsoidalCS,omitempty"`
	RemoteSchema  string             `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField     string             `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href          string             `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role          string             `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole       string             `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title         string             `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show          string             `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate       string             `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type EllipsoidalCSType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this coordinate system. The first csID, if any, is normally the primary identification code, and any others are aliases.
	CsID []IdentifierType `xml:"http://www.opengis.net/gml csID"`
	// Comments on or information about this coordinate system, including data source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Ordered sequence of associations to the coordinate system axes included in this coordinate system.
	UsesAxis []CoordinateSystemAxisRefType `xml:"http://www.opengis.net/gml usesAxis"`
}

type EngineeringCRSRefType struct {
	EngineeringCRS *EngineeringCRSType `xml:"http://www.opengis.net/gml EngineeringCRS,omitempty"`
	RemoteSchema   string              `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField      string              `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href           string              `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role           string              `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole        string              `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title          string              `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show           string              `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate        string              `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type EngineeringCRSType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alterative identifications of this reference system. The first srsID, if any, is normally the primary identification code, and any others are aliases.
	SrsID []IdentifierType `xml:"http://www.opengis.net/gml srsID"`
	// Comments on or information about this reference system, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
	// Association to the coordinate system used by this CRS.
	UsesCS *CoordinateSystemRefType `xml:"http://www.opengis.net/gml usesCS,omitempty"`
	// Association to the engineering datum used by this CRS.
	UsesEngineeringDatum *EngineeringDatumRefType `xml:"http://www.opengis.net/gml usesEngineeringDatum,omitempty"`
}

type EngineeringDatumRefType struct {
	EngineeringDatum *EngineeringDatumType `xml:"http://www.opengis.net/gml EngineeringDatum,omitempty"`
	RemoteSchema     string                `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField        string                `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href             string                `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role             string                `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole          string                `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title            string                `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show             string                `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate          string                `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type EngineeringDatumType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this datum. The first datumID, if any, is normally the primary identification code, and any others are aliases.
	DatumID []IdentifierType `xml:"http://www.opengis.net/gml datumID"`
	// Comments on this reference system, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Description, possibly including coordinates, of the point or points used to anchor the datum to the Earth. Also known as the "origin", especially for engineering and image datums. The codeSpace attribute can be used to reference a source of more detailed on this point or surface, or on a set of such descriptions.
	// - For a geodetic datum, this point is also known as the fundamental point, which is traditionally the point where the relationship between geoid and ellipsoid is defined. In some cases, the "fundamental point" may consist of a number of points. In those cases, the parameters defining the geoid/ellipsoid relationship have been averaged for these points, and the averages adopted as the datum definition.
	// - For an engineering datum, the anchor point may be a physical point, or it may be a point with defined coordinates in another CRS. When appropriate, the coordinates of this anchor point can be referenced in another document, such as referencing a GML feature that references or includes a point position.
	// - For an image datum, the anchor point is usually either the centre of the image or the corner of the image.
	// - For a temporal datum, this attribute is not defined. Instead of the anchor point, a temporal datum carries a separate time origin of type DateTime.
	AnchorPoint *CodeType `xml:"http://www.opengis.net/gml anchorPoint,omitempty"`
	// The time after which this datum definition is valid. This time may be precise (e.g. 1997.0 for IRTF97) or merely a year (e.g. 1983 for NAD83). In the latter case, the epoch usually refers to the year in which a major recalculation of the geodetic control network, underlying the datum, was executed or initiated. An old datum can remain valid after a new datum is defined. Alternatively, a datum may be superseded by a later datum, in which case the realization epoch for the new datum defines the upper limit for the validity of the superseded datum.
	RealizationEpoch *string `xml:"http://www.opengis.net/gml realizationEpoch,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
}

type EnvelopeType struct {
	// deprecated with GML version 3.0
	Coord []CoordType `xml:"http://www.opengis.net/gml coord"`
	// Deprecated with GML version 3.1. Use the explicit properties "lowerCorner" and "upperCorner" instead.
	Pos []DirectPositionType `xml:"http://www.opengis.net/gml pos"`
	// Deprecated with GML version 3.1.0. Use the explicit properties "lowerCorner" and "upperCorner" instead.
	Coordinates *CoordinatesType    `xml:"http://www.opengis.net/gml coordinates,omitempty"`
	LowerCorner *DirectPositionType `xml:"http://www.opengis.net/gml lowerCorner,omitempty"`
	UpperCorner *DirectPositionType `xml:"http://www.opengis.net/gml upperCorner,omitempty"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
}

type EnvelopeWithTimePeriodType struct {
	// deprecated with GML version 3.0
	Coord []CoordType `xml:"http://www.opengis.net/gml coord"`
	// Deprecated with GML version 3.1. Use the explicit properties "lowerCorner" and "upperCorner" instead.
	Pos []DirectPositionType `xml:"http://www.opengis.net/gml pos"`
	// Deprecated with GML version 3.1.0. Use the explicit properties "lowerCorner" and "upperCorner" instead.
	Coordinates *CoordinatesType    `xml:"http://www.opengis.net/gml coordinates,omitempty"`
	LowerCorner *DirectPositionType `xml:"http://www.opengis.net/gml lowerCorner,omitempty"`
	UpperCorner *DirectPositionType `xml:"http://www.opengis.net/gml upperCorner,omitempty"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
	// Direct representation of a temporal position
	TimePosition []TimePositionType `xml:"http://www.opengis.net/gml timePosition"`
	Frame        *string            `xml:"frame,attr,omitempty"`
}

type ExtentType struct {
	// Description of spatial and/or temporal extent of this object.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Unordered list of vertical intervals whose union describes the spatial domain of this object.
	VerticalExtent []EnvelopeType `xml:"http://www.opengis.net/gml verticalExtent"`
	// Unordered list of time periods whose union describes the spatial domain of this object.
	TemporalExtent []TimePeriodType `xml:"http://www.opengis.net/gml temporalExtent"`
	// Unordered list of bounding boxes (or envelopes) whose union describes the spatial domain of this object.
	BoundingBox []EnvelopeType `xml:"http://www.opengis.net/gml boundingBox"`
	// Unordered list of bounding polygons whose union describes the spatial domain of this object.
	BoundingPolygon []PolygonType `xml:"http://www.opengis.net/gml boundingPolygon"`
}

type FaceType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType                      `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType                      `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType                      `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType                      `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType                      `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType                      `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType                      `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType                      `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType                      `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType                      `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                          `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	Isolated                []IsolatedPropertyType          `xml:"http://www.opengis.net/gml isolated"`
	Container               *ContainerPropertyType          `xml:"http://www.opengis.net/gml container,omitempty"`
	DirectedEdge            []DirectedEdgePropertyType      `xml:"http://www.opengis.net/gml directedEdge"`
	DirectedTopoSolid       []DirectedTopoSolidPropertyType `xml:"http://www.opengis.net/gml directedTopoSolid"`
	// This property element either references a surface via the XLink-attributes or contains the surface element. surfaceProperty is the predefined property which can be used by GML Application Schemas whenever a GML Feature has a property with a value that is substitutable for _Surface.
	SurfaceProperty *SurfacePropertyType `xml:"http://www.opengis.net/gml surfaceProperty,omitempty"`
}

type FeatureArrayPropertyType struct {
	AbstractFeature               []AbstractFeatureType               `xml:"http://www.opengis.net/gml _Feature"`
	DirectedObservation           []DirectedObservationType           `xml:"http://www.opengis.net/gml DirectedObservation"`
	DirectedObservationAtDistance []DirectedObservationAtDistanceType `xml:"http://www.opengis.net/gml DirectedObservationAtDistance"`
	FeatureCollection             []FeatureCollectionType             `xml:"http://www.opengis.net/gml FeatureCollection"`
	GridCoverage                  []GridCoverageType                  `xml:"http://www.opengis.net/gml GridCoverage"`
	MultiCurveCoverage            []MultiCurveCoverageType            `xml:"http://www.opengis.net/gml MultiCurveCoverage"`
	MultiPointCoverage            []MultiPointCoverageType            `xml:"http://www.opengis.net/gml MultiPointCoverage"`
	MultiSolidCoverage            []MultiSolidCoverageType            `xml:"http://www.opengis.net/gml MultiSolidCoverage"`
	MultiSurfaceCoverage          []MultiSurfaceCoverageType          `xml:"http://www.opengis.net/gml MultiSurfaceCoverage"`
	Observation                   []ObservationType                   `xml:"http://www.opengis.net/gml Observation"`
	RectifiedGridCoverage         []RectifiedGridCoverageType         `xml:"http://www.opengis.net/gml RectifiedGridCoverage"`
	AbstractContinuousCoverage    []AbstractContinuousCoverageType    `xml:"http://www.opengis.net/gml _ContinuousCoverage"`
	AbstractCoverage              []AbstractCoverageType              `xml:"http://www.opengis.net/gml _Coverage"`
	AbstractDiscreteCoverage      []AbstractDiscreteCoverageType      `xml:"http://www.opengis.net/gml _DiscreteCoverage"`
	AbstractFeatureCollection     []AbstractFeatureCollectionType     `xml:"http://www.opengis.net/gml _FeatureCollection"`
}

type FeatureCollectionType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType         `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType         `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType         `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType         `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType         `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType         `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType         `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType         `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType         `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType         `xml:"http://www.opengis.net/gml srsName"`
	Id                      string             `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	BoundedBy               *BoundingShapeType `xml:"http://www.opengis.net/gml boundedBy,omitempty"`
	// deprecated in GML version 3.1
	Location         *LocationPropertyType         `xml:"http://www.opengis.net/gml location,omitempty"`
	PriorityLocation *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml priorityLocation,omitempty"`
	FeatureMember    []FeaturePropertyType         `xml:"http://www.opengis.net/gml featureMember"`
	FeatureMembers   *FeatureArrayPropertyType     `xml:"http://www.opengis.net/gml featureMembers,omitempty"`
}

type FeaturePropertyType struct {
	AbstractFeature               *AbstractFeatureType               `xml:"http://www.opengis.net/gml _Feature,omitempty"`
	DirectedObservation           *DirectedObservationType           `xml:"http://www.opengis.net/gml DirectedObservation,omitempty"`
	DirectedObservationAtDistance *DirectedObservationAtDistanceType `xml:"http://www.opengis.net/gml DirectedObservationAtDistance,omitempty"`
	FeatureCollection             *FeatureCollectionType             `xml:"http://www.opengis.net/gml FeatureCollection,omitempty"`
	GridCoverage                  *GridCoverageType                  `xml:"http://www.opengis.net/gml GridCoverage,omitempty"`
	MultiCurveCoverage            *MultiCurveCoverageType            `xml:"http://www.opengis.net/gml MultiCurveCoverage,omitempty"`
	MultiPointCoverage            *MultiPointCoverageType            `xml:"http://www.opengis.net/gml MultiPointCoverage,omitempty"`
	MultiSolidCoverage            *MultiSolidCoverageType            `xml:"http://www.opengis.net/gml MultiSolidCoverage,omitempty"`
	MultiSurfaceCoverage          *MultiSurfaceCoverageType          `xml:"http://www.opengis.net/gml MultiSurfaceCoverage,omitempty"`
	Observation                   *ObservationType                   `xml:"http://www.opengis.net/gml Observation,omitempty"`
	RectifiedGridCoverage         *RectifiedGridCoverageType         `xml:"http://www.opengis.net/gml RectifiedGridCoverage,omitempty"`
	AbstractContinuousCoverage    *AbstractContinuousCoverageType    `xml:"http://www.opengis.net/gml _ContinuousCoverage,omitempty"`
	AbstractCoverage              *AbstractCoverageType              `xml:"http://www.opengis.net/gml _Coverage,omitempty"`
	AbstractDiscreteCoverage      *AbstractDiscreteCoverageType      `xml:"http://www.opengis.net/gml _DiscreteCoverage,omitempty"`
	AbstractFeatureCollection     *AbstractFeatureCollectionType     `xml:"http://www.opengis.net/gml _FeatureCollection,omitempty"`
	RemoteSchema                  string                             `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                     string                             `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                          string                             `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                          string                             `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                       string                             `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                         string                             `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                          string                             `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                       string                             `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type FeatureStylePropertyType struct {
	// The style descriptor for features.
	FeatureStyle *FeatureStyleType `xml:"http://www.opengis.net/gml FeatureStyle,omitempty"`
	About        *string           `xml:"about,attr,omitempty"`
	RemoteSchema string            `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string            `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string            `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string            `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string            `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string            `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string            `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string            `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type FeatureStyleType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType                  `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType                  `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType                  `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType                  `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType                  `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType                  `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType                  `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType                  `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType                  `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType                  `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                      `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	FeatureConstraint       string                      `xml:"http://www.opengis.net/gml featureConstraint,omitempty"`
	GeometryStyle           []GeometryStylePropertyType `xml:"http://www.opengis.net/gml geometryStyle"`
	TopologyStyle           []TopologyStylePropertyType `xml:"http://www.opengis.net/gml topologyStyle"`
	LabelStyle              *LabelStylePropertyType     `xml:"http://www.opengis.net/gml labelStyle,omitempty"`
	FeatureType             *string                     `xml:"featureType,attr,omitempty"`
	BaseType                *string                     `xml:"baseType,attr,omitempty"`
	QueryGrammar            *string                     `xml:"queryGrammar,attr,omitempty"`
}

type FileType struct {
	RangeParameters *RangeParametersType `xml:"http://www.opengis.net/gml rangeParameters,omitempty"`
	FileName        string               `xml:"http://www.opengis.net/gml fileName,omitempty"`
	FileStructure   string               `xml:"http://www.opengis.net/gml fileStructure,omitempty"`
	MimeType        string               `xml:"http://www.opengis.net/gml mimeType,omitempty"`
	Compression     string               `xml:"http://www.opengis.net/gml compression,omitempty"`
}

type FormulaType struct {
	A float64 `xml:"http://www.opengis.net/gml a,omitempty"`
	B float64 `xml:"http://www.opengis.net/gml b,omitempty"`
	C float64 `xml:"http://www.opengis.net/gml c,omitempty"`
	D float64 `xml:"http://www.opengis.net/gml d,omitempty"`
}

type GeneralConversionRefType struct {
	AbstractGeneralConversion *AbstractGeneralConversionType `xml:"http://www.opengis.net/gml _GeneralConversion,omitempty"`
	Conversion                *ConversionType                `xml:"http://www.opengis.net/gml Conversion,omitempty"`
	RemoteSchema              string                         `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                 string                         `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                      string                         `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                      string                         `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                   string                         `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                     string                         `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                      string                         `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                   string                         `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type GeneralTransformationRefType struct {
	AbstractGeneralTransformation *AbstractGeneralTransformationType `xml:"http://www.opengis.net/gml _GeneralTransformation,omitempty"`
	Transformation                *TransformationType                `xml:"http://www.opengis.net/gml Transformation,omitempty"`
	RemoteSchema                  string                             `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                     string                             `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                          string                             `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                          string                             `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                       string                             `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                         string                             `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                          string                             `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                       string                             `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type GenericMetaDataType struct {
	Id string `xml:"http://www.opengis.net/gml id,attr,omitempty"`
}

type GeocentricCRSRefType struct {
	GeocentricCRS *GeocentricCRSType `xml:"http://www.opengis.net/gml GeocentricCRS,omitempty"`
	RemoteSchema  string             `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField     string             `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href          string             `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role          string             `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole       string             `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title         string             `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show          string             `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate       string             `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type GeocentricCRSType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alterative identifications of this reference system. The first srsID, if any, is normally the primary identification code, and any others are aliases.
	SrsID []IdentifierType `xml:"http://www.opengis.net/gml srsID"`
	// Comments on or information about this reference system, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
	// Association to the geodetic datum used by this CRS.
	UsesGeodeticDatum *GeodeticDatumRefType `xml:"http://www.opengis.net/gml usesGeodeticDatum,omitempty"`
	// Association to the Cartesian coordinate system used by this CRS.
	UsesCartesianCS *CartesianCSRefType `xml:"http://www.opengis.net/gml usesCartesianCS,omitempty"`
	// Association to the spherical coordinate system used by this CRS.
	UsesSphericalCS *SphericalCSRefType `xml:"http://www.opengis.net/gml usesSphericalCS,omitempty"`
}

type GeodesicStringType struct {
	// The attribute "numDerivativesAtStart" specifies the type of continuity between this curve segment and its predecessor. If this is the first curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtStart *int `xml:"numDerivativesAtStart,attr,omitempty"`
	// The attribute "numDerivativesAtEnd" specifies the type of continuity between this curve segment and its successor. If this is the last curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtEnd *int `xml:"numDerivativesAtEnd,attr,omitempty"`
	// The attribute "numDerivativesInterior" specifies the type of continuity that is guaranteed interior to the curve. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativeInterior *int                    `xml:"numDerivativeInterior,attr,omitempty"`
	PosList               *DirectPositionListType `xml:"http://www.opengis.net/gml posList,omitempty"`
	Pos                   *DirectPositionType     `xml:"http://www.opengis.net/gml pos,omitempty"`
	// This property element either references a point via the XLink-attributes or contains the point element. pointProperty
	// is the predefined property which can be used by GML Application Schemas whenever a GML Feature has a property with a value that
	// is substitutable for Point.
	PointProperty *PointPropertyType `xml:"http://www.opengis.net/gml pointProperty,omitempty"`
	// The attribute "interpolation" specifies the
	// curve interpolation mechanism used for this segment. This
	// mechanism uses the control points and control parameters to
	// determine the position of this curve segment. For an
	// GeodesicString the interpolation is fixed as "geodesic".
	Interpolation *string `xml:"interpolation,attr,omitempty"`
}

type GeodesicType struct {
	// The attribute "numDerivativesAtStart" specifies the type of continuity between this curve segment and its predecessor. If this is the first curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtStart *int `xml:"numDerivativesAtStart,attr,omitempty"`
	// The attribute "numDerivativesAtEnd" specifies the type of continuity between this curve segment and its successor. If this is the last curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtEnd *int `xml:"numDerivativesAtEnd,attr,omitempty"`
	// The attribute "numDerivativesInterior" specifies the type of continuity that is guaranteed interior to the curve. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativeInterior *int                    `xml:"numDerivativeInterior,attr,omitempty"`
	PosList               *DirectPositionListType `xml:"http://www.opengis.net/gml posList,omitempty"`
	Pos                   *DirectPositionType     `xml:"http://www.opengis.net/gml pos,omitempty"`
	// This property element either references a point via the XLink-attributes or contains the point element. pointProperty
	// is the predefined property which can be used by GML Application Schemas whenever a GML Feature has a property with a value that
	// is substitutable for Point.
	PointProperty *PointPropertyType `xml:"http://www.opengis.net/gml pointProperty,omitempty"`
	// The attribute "interpolation" specifies the
	// curve interpolation mechanism used for this segment. This
	// mechanism uses the control points and control parameters to
	// determine the position of this curve segment. For an
	// GeodesicString the interpolation is fixed as "geodesic".
	Interpolation *string `xml:"interpolation,attr,omitempty"`
}

type GeodeticDatumRefType struct {
	GeodeticDatum *GeodeticDatumType `xml:"http://www.opengis.net/gml GeodeticDatum,omitempty"`
	RemoteSchema  string             `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField     string             `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href          string             `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role          string             `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole       string             `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title         string             `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show          string             `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate       string             `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type GeodeticDatumType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this datum. The first datumID, if any, is normally the primary identification code, and any others are aliases.
	DatumID []IdentifierType `xml:"http://www.opengis.net/gml datumID"`
	// Comments on this reference system, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Description, possibly including coordinates, of the point or points used to anchor the datum to the Earth. Also known as the "origin", especially for engineering and image datums. The codeSpace attribute can be used to reference a source of more detailed on this point or surface, or on a set of such descriptions.
	// - For a geodetic datum, this point is also known as the fundamental point, which is traditionally the point where the relationship between geoid and ellipsoid is defined. In some cases, the "fundamental point" may consist of a number of points. In those cases, the parameters defining the geoid/ellipsoid relationship have been averaged for these points, and the averages adopted as the datum definition.
	// - For an engineering datum, the anchor point may be a physical point, or it may be a point with defined coordinates in another CRS. When appropriate, the coordinates of this anchor point can be referenced in another document, such as referencing a GML feature that references or includes a point position.
	// - For an image datum, the anchor point is usually either the centre of the image or the corner of the image.
	// - For a temporal datum, this attribute is not defined. Instead of the anchor point, a temporal datum carries a separate time origin of type DateTime.
	AnchorPoint *CodeType `xml:"http://www.opengis.net/gml anchorPoint,omitempty"`
	// The time after which this datum definition is valid. This time may be precise (e.g. 1997.0 for IRTF97) or merely a year (e.g. 1983 for NAD83). In the latter case, the epoch usually refers to the year in which a major recalculation of the geodetic control network, underlying the datum, was executed or initiated. An old datum can remain valid after a new datum is defined. Alternatively, a datum may be superseded by a later datum, in which case the realization epoch for the new datum defines the upper limit for the validity of the superseded datum.
	RealizationEpoch *string `xml:"http://www.opengis.net/gml realizationEpoch,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
	// Association to the prime meridian used by this geodetic datum.
	UsesPrimeMeridian *PrimeMeridianRefType `xml:"http://www.opengis.net/gml usesPrimeMeridian,omitempty"`
	// Association to the ellipsoid used by this geodetic datum.
	UsesEllipsoid *EllipsoidRefType `xml:"http://www.opengis.net/gml usesEllipsoid,omitempty"`
}

type GeographicCRSRefType struct {
	GeographicCRS *GeographicCRSType `xml:"http://www.opengis.net/gml GeographicCRS,omitempty"`
	RemoteSchema  string             `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField     string             `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href          string             `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role          string             `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole       string             `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title         string             `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show          string             `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate       string             `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type GeographicCRSType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alterative identifications of this reference system. The first srsID, if any, is normally the primary identification code, and any others are aliases.
	SrsID []IdentifierType `xml:"http://www.opengis.net/gml srsID"`
	// Comments on or information about this reference system, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
	// Association to the ellipsoidal coordinate system used by this CRS.
	UsesEllipsoidalCS *EllipsoidalCSRefType `xml:"http://www.opengis.net/gml usesEllipsoidalCS,omitempty"`
	// Association to the geodetic datum used by this CRS.
	UsesGeodeticDatum *GeodeticDatumRefType `xml:"http://www.opengis.net/gml usesGeodeticDatum,omitempty"`
}

type GeometricComplexPropertyType struct {
	GeometricComplex *GeometricComplexType `xml:"http://www.opengis.net/gml GeometricComplex,omitempty"`
	CompositeCurve   *CompositeCurveType   `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	CompositeSurface *CompositeSurfaceType `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	CompositeSolid   *CompositeSolidType   `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	RemoteSchema     string                `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField        string                `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href             string                `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role             string                `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole          string                `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title            string                `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show             string                `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate          string                `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type GeometricComplexType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string                          `xml:"uomLabels,attr,omitempty"`
	Element   []GeometricPrimitivePropertyType `xml:"http://www.opengis.net/gml element"`
}

type GeometricPrimitivePropertyType struct {
	// The "_GeometricPrimitive" element is the abstract head of the substituition group for all (pre- and user-defined)
	// geometric primitives.
	AbstractGeometricPrimitive *AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml _GeometricPrimitive,omitempty"`
	CompositeCurve             *CompositeCurveType             `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	CompositeSolid             *CompositeSolidType             `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	CompositeSurface           *CompositeSurfaceType           `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	Curve                      *CurveType                      `xml:"http://www.opengis.net/gml Curve,omitempty"`
	LineString                 *LineStringType                 `xml:"http://www.opengis.net/gml LineString,omitempty"`
	OrientableCurve            *OrientableCurveType            `xml:"http://www.opengis.net/gml OrientableCurve,omitempty"`
	OrientableSurface          *OrientableSurfaceType          `xml:"http://www.opengis.net/gml OrientableSurface,omitempty"`
	Point                      *PointType                      `xml:"http://www.opengis.net/gml Point,omitempty"`
	Polygon                    *PolygonType                    `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	PolyhedralSurface          *PolyhedralSurfaceType          `xml:"http://www.opengis.net/gml PolyhedralSurface,omitempty"`
	Solid                      *SolidType                      `xml:"http://www.opengis.net/gml Solid,omitempty"`
	Surface                    *SurfaceType                    `xml:"http://www.opengis.net/gml Surface,omitempty"`
	Tin                        *TinType                        `xml:"http://www.opengis.net/gml Tin,omitempty"`
	TriangulatedSurface        *TriangulatedSurfaceType        `xml:"http://www.opengis.net/gml TriangulatedSurface,omitempty"`
	AbstractCurve              *AbstractCurveType              `xml:"http://www.opengis.net/gml _Curve,omitempty"`
	AbstractSolid              *AbstractSolidType              `xml:"http://www.opengis.net/gml _Solid,omitempty"`
	AbstractSurface            *AbstractSurfaceType            `xml:"http://www.opengis.net/gml _Surface,omitempty"`
	RemoteSchema               string                          `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                  string                          `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                       string                          `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                       string                          `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                    string                          `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                      string                          `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                       string                          `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                    string                          `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type GeometryArrayPropertyType struct {
	// The "_Geometry" element is the abstract head of the substituition group for all geometry elements of GML 3. This
	// includes pre-defined and user-defined geometry elements. Any geometry element must be a direct or indirect extension/restriction
	// of AbstractGeometryType and must be directly or indirectly in the substitution group of "_Geometry".
	AbstractGeometry           []AbstractGeometryType           `xml:"http://www.opengis.net/gml _Geometry"`
	CompositeCurve             []CompositeCurveType             `xml:"http://www.opengis.net/gml CompositeCurve"`
	CompositeSolid             []CompositeSolidType             `xml:"http://www.opengis.net/gml CompositeSolid"`
	CompositeSurface           []CompositeSurfaceType           `xml:"http://www.opengis.net/gml CompositeSurface"`
	Curve                      []CurveType                      `xml:"http://www.opengis.net/gml Curve"`
	GeometricComplex           []GeometricComplexType           `xml:"http://www.opengis.net/gml GeometricComplex"`
	Grid                       []GridType                       `xml:"http://www.opengis.net/gml Grid"`
	LineString                 []LineStringType                 `xml:"http://www.opengis.net/gml LineString"`
	LinearRing                 []LinearRingType                 `xml:"http://www.opengis.net/gml LinearRing"`
	MultiCurve                 []MultiCurveType                 `xml:"http://www.opengis.net/gml MultiCurve"`
	MultiGeometry              []MultiGeometryType              `xml:"http://www.opengis.net/gml MultiGeometry"`
	MultiLineString            []MultiLineStringType            `xml:"http://www.opengis.net/gml MultiLineString"`
	MultiPoint                 []MultiPointType                 `xml:"http://www.opengis.net/gml MultiPoint"`
	MultiPolygon               []MultiPolygonType               `xml:"http://www.opengis.net/gml MultiPolygon"`
	MultiSolid                 []MultiSolidType                 `xml:"http://www.opengis.net/gml MultiSolid"`
	MultiSurface               []MultiSurfaceType               `xml:"http://www.opengis.net/gml MultiSurface"`
	OrientableCurve            []OrientableCurveType            `xml:"http://www.opengis.net/gml OrientableCurve"`
	OrientableSurface          []OrientableSurfaceType          `xml:"http://www.opengis.net/gml OrientableSurface"`
	Point                      []PointType                      `xml:"http://www.opengis.net/gml Point"`
	Polygon                    []PolygonType                    `xml:"http://www.opengis.net/gml Polygon"`
	PolyhedralSurface          []PolyhedralSurfaceType          `xml:"http://www.opengis.net/gml PolyhedralSurface"`
	RectifiedGrid              []RectifiedGridType              `xml:"http://www.opengis.net/gml RectifiedGrid"`
	Ring                       []RingType                       `xml:"http://www.opengis.net/gml Ring"`
	Solid                      []SolidType                      `xml:"http://www.opengis.net/gml Solid"`
	Surface                    []SurfaceType                    `xml:"http://www.opengis.net/gml Surface"`
	Tin                        []TinType                        `xml:"http://www.opengis.net/gml Tin"`
	TriangulatedSurface        []TriangulatedSurfaceType        `xml:"http://www.opengis.net/gml TriangulatedSurface"`
	AbstractCurve              []AbstractCurveType              `xml:"http://www.opengis.net/gml _Curve"`
	AbstractGeometricAggregate []AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml _GeometricAggregate"`
	AbstractGeometricPrimitive []AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml _GeometricPrimitive"`
	AbstractImplicitGeometry   []AbstractGeometryType           `xml:"http://www.opengis.net/gml _ImplicitGeometry"`
	AbstractRing               []AbstractRingType               `xml:"http://www.opengis.net/gml _Ring"`
	AbstractSolid              []AbstractSolidType              `xml:"http://www.opengis.net/gml _Solid"`
	AbstractSurface            []AbstractSurfaceType            `xml:"http://www.opengis.net/gml _Surface"`
}

type GeometryPropertyType struct {
	// The "_Geometry" element is the abstract head of the substituition group for all geometry elements of GML 3. This
	// includes pre-defined and user-defined geometry elements. Any geometry element must be a direct or indirect extension/restriction
	// of AbstractGeometryType and must be directly or indirectly in the substitution group of "_Geometry".
	AbstractGeometry           *AbstractGeometryType           `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	CompositeCurve             *CompositeCurveType             `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	CompositeSolid             *CompositeSolidType             `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	CompositeSurface           *CompositeSurfaceType           `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	Curve                      *CurveType                      `xml:"http://www.opengis.net/gml Curve,omitempty"`
	GeometricComplex           *GeometricComplexType           `xml:"http://www.opengis.net/gml GeometricComplex,omitempty"`
	Grid                       *GridType                       `xml:"http://www.opengis.net/gml Grid,omitempty"`
	LineString                 *LineStringType                 `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearRing                 *LinearRingType                 `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MultiCurve                 *MultiCurveType                 `xml:"http://www.opengis.net/gml MultiCurve,omitempty"`
	MultiGeometry              *MultiGeometryType              `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString            *MultiLineStringType            `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                 *MultiPointType                 `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPolygon               *MultiPolygonType               `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	MultiSolid                 *MultiSolidType                 `xml:"http://www.opengis.net/gml MultiSolid,omitempty"`
	MultiSurface               *MultiSurfaceType               `xml:"http://www.opengis.net/gml MultiSurface,omitempty"`
	OrientableCurve            *OrientableCurveType            `xml:"http://www.opengis.net/gml OrientableCurve,omitempty"`
	OrientableSurface          *OrientableSurfaceType          `xml:"http://www.opengis.net/gml OrientableSurface,omitempty"`
	Point                      *PointType                      `xml:"http://www.opengis.net/gml Point,omitempty"`
	Polygon                    *PolygonType                    `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	PolyhedralSurface          *PolyhedralSurfaceType          `xml:"http://www.opengis.net/gml PolyhedralSurface,omitempty"`
	RectifiedGrid              *RectifiedGridType              `xml:"http://www.opengis.net/gml RectifiedGrid,omitempty"`
	Ring                       *RingType                       `xml:"http://www.opengis.net/gml Ring,omitempty"`
	Solid                      *SolidType                      `xml:"http://www.opengis.net/gml Solid,omitempty"`
	Surface                    *SurfaceType                    `xml:"http://www.opengis.net/gml Surface,omitempty"`
	Tin                        *TinType                        `xml:"http://www.opengis.net/gml Tin,omitempty"`
	TriangulatedSurface        *TriangulatedSurfaceType        `xml:"http://www.opengis.net/gml TriangulatedSurface,omitempty"`
	AbstractCurve              *AbstractCurveType              `xml:"http://www.opengis.net/gml _Curve,omitempty"`
	AbstractGeometricAggregate *AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml _GeometricAggregate,omitempty"`
	AbstractGeometricPrimitive *AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml _GeometricPrimitive,omitempty"`
	AbstractImplicitGeometry   *AbstractGeometryType           `xml:"http://www.opengis.net/gml _ImplicitGeometry,omitempty"`
	AbstractRing               *AbstractRingType               `xml:"http://www.opengis.net/gml _Ring,omitempty"`
	AbstractSolid              *AbstractSolidType              `xml:"http://www.opengis.net/gml _Solid,omitempty"`
	AbstractSurface            *AbstractSurfaceType            `xml:"http://www.opengis.net/gml _Surface,omitempty"`
	RemoteSchema               string                          `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                  string                          `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                       string                          `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                       string                          `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                    string                          `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                      string                          `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                       string                          `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                    string                          `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type GeometryStylePropertyType struct {
	// The style descriptor for geometries of a feature.
	GeometryStyle *GeometryStyleType `xml:"http://www.opengis.net/gml GeometryStyle,omitempty"`
	About         *string            `xml:"about,attr,omitempty"`
	RemoteSchema  string             `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField     string             `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href          string             `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role          string             `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole       string             `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title         string             `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show          string             `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate       string             `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type GeometryStyleType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType              `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType              `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType              `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType              `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType              `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType              `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType              `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType              `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType              `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType              `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	SpatialResolution       *ScaleType              `xml:"http://www.opengis.net/gml spatialResolution,omitempty"`
	StyleVariation          []StyleVariationType    `xml:"http://www.opengis.net/gml styleVariation"`
	Animate                 []string                `xml:"http://www.w3.org/2001/SMIL20/ animate"`
	AnimateMotion           []string                `xml:"http://www.w3.org/2001/SMIL20/ animateMotion"`
	AnimateColor            []string                `xml:"http://www.w3.org/2001/SMIL20/ animateColor"`
	Set                     []string                `xml:"http://www.w3.org/2001/SMIL20/ set"`
	LabelStyle              *LabelStylePropertyType `xml:"http://www.opengis.net/gml labelStyle,omitempty"`
	// The symbol property. Extends the gml:AssociationType to allow for remote referencing of symbols.
	Symbol *SymbolType `xml:"http://www.opengis.net/gml symbol,omitempty"`
	// Deprecated in GML version 3.1.0. Use symbol with inline content instead.
	Style            string  `xml:"http://www.opengis.net/gml style,omitempty"`
	GeometryProperty *string `xml:"geometryProperty,attr,omitempty"`
	GeometryType     *string `xml:"geometryType,attr,omitempty"`
}

type GraphStylePropertyType struct {
	// The style descriptor for a graph consisting of a number of features. Describes graph-specific style attributes.
	GraphStyle   *GraphStyleType `xml:"http://www.opengis.net/gml GraphStyle,omitempty"`
	About        *string         `xml:"about,attr,omitempty"`
	RemoteSchema string          `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string          `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string          `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string          `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string          `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string          `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string          `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string          `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type GraphStyleType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType           `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType           `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType           `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType           `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType           `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType           `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType           `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType           `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType           `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType           `xml:"http://www.opengis.net/gml srsName"`
	Id                      string               `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	SpatialResolution       *ScaleType           `xml:"http://www.opengis.net/gml spatialResolution,omitempty"`
	StyleVariation          []StyleVariationType `xml:"http://www.opengis.net/gml styleVariation"`
	Animate                 []string             `xml:"http://www.w3.org/2001/SMIL20/ animate"`
	AnimateMotion           []string             `xml:"http://www.w3.org/2001/SMIL20/ animateMotion"`
	AnimateColor            []string             `xml:"http://www.w3.org/2001/SMIL20/ animateColor"`
	Set                     []string             `xml:"http://www.w3.org/2001/SMIL20/ set"`
	Planar                  bool                 `xml:"http://www.opengis.net/gml planar,omitempty"`
	Directed                bool                 `xml:"http://www.opengis.net/gml directed,omitempty"`
	Grid                    bool                 `xml:"http://www.opengis.net/gml grid,omitempty"`
	MinDistance             float64              `xml:"http://www.opengis.net/gml minDistance,omitempty"`
	MinAngle                float64              `xml:"http://www.opengis.net/gml minAngle,omitempty"`
	GraphType               string               `xml:"http://www.opengis.net/gml graphType,omitempty"`
	DrawingType             string               `xml:"http://www.opengis.net/gml drawingType,omitempty"`
	LineType                string               `xml:"http://www.opengis.net/gml lineType,omitempty"`
	AestheticCriteria       []string             `xml:"http://www.opengis.net/gml aestheticCriteria"`
}

type GridCoverageType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType         `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType         `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType         `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType         `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType         `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType         `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType         `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType         `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType         `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType         `xml:"http://www.opengis.net/gml srsName"`
	Id                      string             `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	BoundedBy               *BoundingShapeType `xml:"http://www.opengis.net/gml boundedBy,omitempty"`
	// deprecated in GML version 3.1
	Location            *LocationPropertyType         `xml:"http://www.opengis.net/gml location,omitempty"`
	PriorityLocation    *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml priorityLocation,omitempty"`
	DomainSet           *DomainSetType                `xml:"http://www.opengis.net/gml domainSet,omitempty"`
	GridDomain          *GridDomainType               `xml:"http://www.opengis.net/gml gridDomain,omitempty"`
	MultiCurveDomain    *MultiCurveDomainType         `xml:"http://www.opengis.net/gml multiCurveDomain,omitempty"`
	MultiPointDomain    *MultiPointDomainType         `xml:"http://www.opengis.net/gml multiPointDomain,omitempty"`
	MultiSolidDomain    *MultiSolidDomainType         `xml:"http://www.opengis.net/gml multiSolidDomain,omitempty"`
	MultiSurfaceDomain  *MultiSurfaceDomainType       `xml:"http://www.opengis.net/gml multiSurfaceDomain,omitempty"`
	RectifiedGridDomain *RectifiedGridDomainType      `xml:"http://www.opengis.net/gml rectifiedGridDomain,omitempty"`
	RangeSet            *RangeSetType                 `xml:"http://www.opengis.net/gml rangeSet,omitempty"`
	Dimension           *int                          `xml:"dimension,attr,omitempty"`
	CoverageFunction    *CoverageFunctionType         `xml:"http://www.opengis.net/gml coverageFunction,omitempty"`
}

type GridDomainType struct {
	// The "_Geometry" element is the abstract head of the substituition group for all geometry elements of GML 3. This
	// includes pre-defined and user-defined geometry elements. Any geometry element must be a direct or indirect extension/restriction
	// of AbstractGeometryType and must be directly or indirectly in the substitution group of "_Geometry".
	AbstractGeometry           *AbstractGeometryType           `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	CompositeCurve             *CompositeCurveType             `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	CompositeSolid             *CompositeSolidType             `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	CompositeSurface           *CompositeSurfaceType           `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	Curve                      *CurveType                      `xml:"http://www.opengis.net/gml Curve,omitempty"`
	GeometricComplex           *GeometricComplexType           `xml:"http://www.opengis.net/gml GeometricComplex,omitempty"`
	Grid                       *GridType                       `xml:"http://www.opengis.net/gml Grid,omitempty"`
	LineString                 *LineStringType                 `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearRing                 *LinearRingType                 `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MultiCurve                 *MultiCurveType                 `xml:"http://www.opengis.net/gml MultiCurve,omitempty"`
	MultiGeometry              *MultiGeometryType              `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString            *MultiLineStringType            `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                 *MultiPointType                 `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPolygon               *MultiPolygonType               `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	MultiSolid                 *MultiSolidType                 `xml:"http://www.opengis.net/gml MultiSolid,omitempty"`
	MultiSurface               *MultiSurfaceType               `xml:"http://www.opengis.net/gml MultiSurface,omitempty"`
	OrientableCurve            *OrientableCurveType            `xml:"http://www.opengis.net/gml OrientableCurve,omitempty"`
	OrientableSurface          *OrientableSurfaceType          `xml:"http://www.opengis.net/gml OrientableSurface,omitempty"`
	Point                      *PointType                      `xml:"http://www.opengis.net/gml Point,omitempty"`
	Polygon                    *PolygonType                    `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	PolyhedralSurface          *PolyhedralSurfaceType          `xml:"http://www.opengis.net/gml PolyhedralSurface,omitempty"`
	RectifiedGrid              *RectifiedGridType              `xml:"http://www.opengis.net/gml RectifiedGrid,omitempty"`
	Ring                       *RingType                       `xml:"http://www.opengis.net/gml Ring,omitempty"`
	Solid                      *SolidType                      `xml:"http://www.opengis.net/gml Solid,omitempty"`
	Surface                    *SurfaceType                    `xml:"http://www.opengis.net/gml Surface,omitempty"`
	Tin                        *TinType                        `xml:"http://www.opengis.net/gml Tin,omitempty"`
	TriangulatedSurface        *TriangulatedSurfaceType        `xml:"http://www.opengis.net/gml TriangulatedSurface,omitempty"`
	AbstractCurve              *AbstractCurveType              `xml:"http://www.opengis.net/gml _Curve,omitempty"`
	AbstractGeometricAggregate *AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml _GeometricAggregate,omitempty"`
	AbstractGeometricPrimitive *AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml _GeometricPrimitive,omitempty"`
	AbstractImplicitGeometry   *AbstractGeometryType           `xml:"http://www.opengis.net/gml _ImplicitGeometry,omitempty"`
	AbstractRing               *AbstractRingType               `xml:"http://www.opengis.net/gml _Ring,omitempty"`
	AbstractSolid              *AbstractSolidType              `xml:"http://www.opengis.net/gml _Solid,omitempty"`
	AbstractSurface            *AbstractSurfaceType            `xml:"http://www.opengis.net/gml _Surface,omitempty"`
	// This abstract element acts as the head of the substitution group for temporal primitives and complexes.
	AbstractTimeObject             *AbstractTimeObjectType             `xml:"http://www.opengis.net/gml _TimeObject,omitempty"`
	TimeEdge                       *TimeEdgeType                       `xml:"http://www.opengis.net/gml TimeEdge,omitempty"`
	TimeInstant                    *TimeInstantType                    `xml:"http://www.opengis.net/gml TimeInstant,omitempty"`
	TimeNode                       *TimeNodeType                       `xml:"http://www.opengis.net/gml TimeNode,omitempty"`
	TimePeriod                     *TimePeriodType                     `xml:"http://www.opengis.net/gml TimePeriod,omitempty"`
	TimeTopologyComplex            *TimeTopologyComplexType            `xml:"http://www.opengis.net/gml TimeTopologyComplex,omitempty"`
	AbstractTimeComplex            *AbstractTimeComplexType            `xml:"http://www.opengis.net/gml _TimeComplex,omitempty"`
	AbstractTimeGeometricPrimitive *AbstractTimeGeometricPrimitiveType `xml:"http://www.opengis.net/gml _TimeGeometricPrimitive,omitempty"`
	AbstractTimePrimitive          *AbstractTimePrimitiveType          `xml:"http://www.opengis.net/gml _TimePrimitive,omitempty"`
	AbstractTimeTopologyPrimitive  *AbstractTimeTopologyPrimitiveType  `xml:"http://www.opengis.net/gml _TimeTopologyPrimitive,omitempty"`
	RemoteSchema                   string                              `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                      string                              `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                           string                              `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                           string                              `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                        string                              `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                          string                              `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                           string                              `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                        string                              `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type GridEnvelopeType struct {
	Low  string `xml:"http://www.opengis.net/gml low,omitempty"`
	High string `xml:"http://www.opengis.net/gml high,omitempty"`
}

type GridFunctionType struct {
	// If absent, the implied value is "Linear".
	SequenceRule *SequenceRuleType `xml:"http://www.opengis.net/gml sequenceRule,omitempty"`
	// Index position of the first grid post, which must lie somwhere in the GridEnvelope.  If absent, the startPoint is equal to the value of gridEnvelope::low from the grid definition.
	StartPoint string `xml:"http://www.opengis.net/gml startPoint,omitempty"`
}

type GridLengthType struct {
	Value string `xml:",chardata"`
}

type GridLimitsType struct {
	GridEnvelope *GridEnvelopeType `xml:"http://www.opengis.net/gml GridEnvelope,omitempty"`
}

type GridType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string         `xml:"uomLabels,attr,omitempty"`
	Limits    *GridLimitsType `xml:"http://www.opengis.net/gml limits,omitempty"`
	AxisName  []string        `xml:"http://www.opengis.net/gml axisName"`
	Dimension int             `xml:"dimension,attr"`
}

type HistoryPropertyType struct {
	AbstractTimeSlice  []AbstractTimeSliceType  `xml:"http://www.opengis.net/gml _TimeSlice"`
	MovingObjectStatus []MovingObjectStatusType `xml:"http://www.opengis.net/gml MovingObjectStatus"`
}

type IdentifierType struct {
	// The code or name for this Identifier, often from a controlled list or pattern defined by a code space. The optional codeSpace attribute is normally included to identify or reference a code space within which one or more codes are defined. This code space is often defined by some authority organization, where one organization may define multiple code spaces. The range and format of each Code Space identifier is defined by that code space authority. Information about that code space authority can be included as metaDataProperty elements which are optionally allowed in all CRS objects.
	Name                    *CodeType `xml:"http://www.opengis.net/gml name,omitempty"`
	CoordinateOperationName *CodeType `xml:"http://www.opengis.net/gml coordinateOperationName,omitempty"`
	CsName                  *CodeType `xml:"http://www.opengis.net/gml csName,omitempty"`
	DatumName               *CodeType `xml:"http://www.opengis.net/gml datumName,omitempty"`
	EllipsoidName           *CodeType `xml:"http://www.opengis.net/gml ellipsoidName,omitempty"`
	GroupName               *CodeType `xml:"http://www.opengis.net/gml groupName,omitempty"`
	MeridianName            *CodeType `xml:"http://www.opengis.net/gml meridianName,omitempty"`
	MethodName              *CodeType `xml:"http://www.opengis.net/gml methodName,omitempty"`
	ParameterName           *CodeType `xml:"http://www.opengis.net/gml parameterName,omitempty"`
	SrsName                 *CodeType `xml:"http://www.opengis.net/gml srsName,omitempty"`
	// Identifier of the version of the associated codeSpace or code, as specified by the codeSpace or code authority. This version is included only when the "code" or "codeSpace" uses versions. When appropriate, the version is identified by the effective date, coded using ISO 8601 date format.
	Version *string `xml:"http://www.opengis.net/gml version,omitempty"`
	// Remarks about this code or alias.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
}

type ImageCRSRefType struct {
	ImageCRS     *ImageCRSType `xml:"http://www.opengis.net/gml ImageCRS,omitempty"`
	RemoteSchema string        `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string        `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string        `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string        `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string        `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string        `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string        `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string        `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type ImageCRSType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alterative identifications of this reference system. The first srsID, if any, is normally the primary identification code, and any others are aliases.
	SrsID []IdentifierType `xml:"http://www.opengis.net/gml srsID"`
	// Comments on or information about this reference system, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
	// Association to the image datum used by this CRS.
	UsesImageDatum *ImageDatumRefType `xml:"http://www.opengis.net/gml usesImageDatum,omitempty"`
	// Association to the Cartesian coordinate system used by this CRS.
	UsesCartesianCS *CartesianCSRefType `xml:"http://www.opengis.net/gml usesCartesianCS,omitempty"`
	// Association to the oblique Cartesian coordinate system used by this CRS.
	UsesObliqueCartesianCS *ObliqueCartesianCSRefType `xml:"http://www.opengis.net/gml usesObliqueCartesianCS,omitempty"`
}

type ImageDatumRefType struct {
	ImageDatum   *ImageDatumType `xml:"http://www.opengis.net/gml ImageDatum,omitempty"`
	RemoteSchema string          `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string          `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string          `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string          `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string          `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string          `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string          `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string          `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type ImageDatumType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this datum. The first datumID, if any, is normally the primary identification code, and any others are aliases.
	DatumID []IdentifierType `xml:"http://www.opengis.net/gml datumID"`
	// Comments on this reference system, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Description, possibly including coordinates, of the point or points used to anchor the datum to the Earth. Also known as the "origin", especially for engineering and image datums. The codeSpace attribute can be used to reference a source of more detailed on this point or surface, or on a set of such descriptions.
	// - For a geodetic datum, this point is also known as the fundamental point, which is traditionally the point where the relationship between geoid and ellipsoid is defined. In some cases, the "fundamental point" may consist of a number of points. In those cases, the parameters defining the geoid/ellipsoid relationship have been averaged for these points, and the averages adopted as the datum definition.
	// - For an engineering datum, the anchor point may be a physical point, or it may be a point with defined coordinates in another CRS. When appropriate, the coordinates of this anchor point can be referenced in another document, such as referencing a GML feature that references or includes a point position.
	// - For an image datum, the anchor point is usually either the centre of the image or the corner of the image.
	// - For a temporal datum, this attribute is not defined. Instead of the anchor point, a temporal datum carries a separate time origin of type DateTime.
	AnchorPoint *CodeType `xml:"http://www.opengis.net/gml anchorPoint,omitempty"`
	// The time after which this datum definition is valid. This time may be precise (e.g. 1997.0 for IRTF97) or merely a year (e.g. 1983 for NAD83). In the latter case, the epoch usually refers to the year in which a major recalculation of the geodetic control network, underlying the datum, was executed or initiated. An old datum can remain valid after a new datum is defined. Alternatively, a datum may be superseded by a later datum, in which case the realization epoch for the new datum defines the upper limit for the validity of the superseded datum.
	RealizationEpoch *string `xml:"http://www.opengis.net/gml realizationEpoch,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope       *string          `xml:"http://www.opengis.net/gml scope,omitempty"`
	PixelInCell *PixelInCellType `xml:"http://www.opengis.net/gml pixelInCell,omitempty"`
}

type IndexMapType struct {
	// If absent, the implied value is "Linear".
	SequenceRule *SequenceRuleType `xml:"http://www.opengis.net/gml sequenceRule,omitempty"`
	// Index position of the first grid post, which must lie somwhere in the GridEnvelope.  If absent, the startPoint is equal to the value of gridEnvelope::low from the grid definition.
	StartPoint  string `xml:"http://www.opengis.net/gml startPoint,omitempty"`
	LookUpTable string `xml:"http://www.opengis.net/gml lookUpTable,omitempty"`
}

type IndirectEntryType struct {
	DefinitionProxy *DefinitionProxyType `xml:"http://www.opengis.net/gml DefinitionProxy,omitempty"`
}

type IsolatedPropertyType struct {
	Node         *NodeType `xml:"http://www.opengis.net/gml Node,omitempty"`
	Edge         *EdgeType `xml:"http://www.opengis.net/gml Edge,omitempty"`
	RemoteSchema string    `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string    `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string    `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string    `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string    `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string    `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string    `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string    `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type KnotPropertyType struct {
	Knot *KnotType `xml:"http://www.opengis.net/gml Knot,omitempty"`
}

type KnotType struct {
	// The property "value" is the value of the parameter at the knot of the spline. The sequence of knots shall be a non-decreasing sequence. That is, each knot's value in the sequence shall be equal to or greater than the previous knot's value. The use of equal consecutive knots is normally handled using the multiplicity.
	Value float64 `xml:"http://www.opengis.net/gml value,omitempty"`
	// The property "multiplicity" is the multiplicity of this knot used in the definition of the spline (with the same weight).
	Multiplicity int `xml:"http://www.opengis.net/gml multiplicity,omitempty"`
	// The property "weight" is the value of the averaging weight used for this knot of the spline.
	Weight float64 `xml:"http://www.opengis.net/gml weight,omitempty"`
}

type LabelStylePropertyType struct {
	// The style descriptor for labels of a feature, geometry or topology.
	LabelStyle   *LabelStyleType `xml:"http://www.opengis.net/gml LabelStyle,omitempty"`
	About        *string         `xml:"about,attr,omitempty"`
	RemoteSchema string          `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string          `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string          `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string          `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string          `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string          `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string          `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string          `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type LabelStyleType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType           `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType           `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType           `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType           `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType           `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType           `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType           `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType           `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType           `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType           `xml:"http://www.opengis.net/gml srsName"`
	Id                      string               `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	SpatialResolution       *ScaleType           `xml:"http://www.opengis.net/gml spatialResolution,omitempty"`
	StyleVariation          []StyleVariationType `xml:"http://www.opengis.net/gml styleVariation"`
	Animate                 []string             `xml:"http://www.w3.org/2001/SMIL20/ animate"`
	AnimateMotion           []string             `xml:"http://www.w3.org/2001/SMIL20/ animateMotion"`
	AnimateColor            []string             `xml:"http://www.w3.org/2001/SMIL20/ animateColor"`
	Set                     []string             `xml:"http://www.w3.org/2001/SMIL20/ set"`
	Style                   string               `xml:"http://www.opengis.net/gml style,omitempty"`
	Label                   *LabelType           `xml:"http://www.opengis.net/gml label,omitempty"`
}

type LabelType struct {
	LabelExpression []string `xml:"http://www.opengis.net/gml LabelExpression"`
	Transform       string   `xml:"http://www.opengis.net/gml transform,attr,omitempty"`
}

type LengthType struct {
	Value string `xml:",chardata"`
}

type LineStringPropertyType struct {
	LineString   *LineStringType `xml:"http://www.opengis.net/gml LineString,omitempty"`
	RemoteSchema string          `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string          `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string          `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string          `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string          `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string          `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string          `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string          `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type LineStringSegmentArrayPropertyType struct {
	LineStringSegment []LineStringSegmentType `xml:"http://www.opengis.net/gml LineStringSegment"`
}

type LineStringSegmentType struct {
	// The attribute "numDerivativesAtStart" specifies the type of continuity between this curve segment and its predecessor. If this is the first curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtStart *int `xml:"numDerivativesAtStart,attr,omitempty"`
	// The attribute "numDerivativesAtEnd" specifies the type of continuity between this curve segment and its successor. If this is the last curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtEnd *int `xml:"numDerivativesAtEnd,attr,omitempty"`
	// The attribute "numDerivativesInterior" specifies the type of continuity that is guaranteed interior to the curve. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativeInterior *int                    `xml:"numDerivativeInterior,attr,omitempty"`
	PosList               *DirectPositionListType `xml:"http://www.opengis.net/gml posList,omitempty"`
	// Deprecated with GML version 3.1.0. Use "posList" instead.
	Coordinates *CoordinatesType     `xml:"http://www.opengis.net/gml coordinates,omitempty"`
	Pos         []DirectPositionType `xml:"http://www.opengis.net/gml pos"`
	// This property element either references a point via the XLink-attributes or contains the point element. pointProperty
	// is the predefined property which can be used by GML Application Schemas whenever a GML Feature has a property with a value that
	// is substitutable for Point.
	PointProperty []PointPropertyType `xml:"http://www.opengis.net/gml pointProperty"`
	// Deprecated with GML version 3.1.0. Use "pointProperty" instead. Included for backwards compatibility with GML 3.0.0.
	PointRep []PointPropertyType `xml:"http://www.opengis.net/gml pointRep"`
	// The attribute "interpolation" specifies the curve interpolation mechanism used for this segment. This mechanism
	// uses the control points and control parameters to determine the position of this curve segment. For a LineStringSegment the interpolation is fixed as "linear".
	Interpolation *string `xml:"interpolation,attr,omitempty"`
}

type LineStringType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string                 `xml:"uomLabels,attr,omitempty"`
	PosList   *DirectPositionListType `xml:"http://www.opengis.net/gml posList,omitempty"`
	// Deprecated with GML version 3.1.0. Use "posList" instead.
	Coordinates *CoordinatesType     `xml:"http://www.opengis.net/gml coordinates,omitempty"`
	Pos         []DirectPositionType `xml:"http://www.opengis.net/gml pos"`
	// This property element either references a point via the XLink-attributes or contains the point element. pointProperty
	// is the predefined property which can be used by GML Application Schemas whenever a GML Feature has a property with a value that
	// is substitutable for Point.
	PointProperty []PointPropertyType `xml:"http://www.opengis.net/gml pointProperty"`
	// Deprecated with GML version 3.1.0. Use "pointProperty" instead. Included for backwards compatibility
	// with GML 3.0.0.
	PointRep []PointPropertyType `xml:"http://www.opengis.net/gml pointRep"`
	// Deprecated with GML version 3.0. Use "pos" instead. The "coord" element is included for backwards
	// compatibility with GML 2.
	Coord []CoordType `xml:"http://www.opengis.net/gml coord"`
}

type LinearCSRefType struct {
	LinearCS     *LinearCSType `xml:"http://www.opengis.net/gml LinearCS,omitempty"`
	RemoteSchema string        `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string        `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string        `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string        `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string        `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string        `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string        `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string        `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type LinearCSType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this coordinate system. The first csID, if any, is normally the primary identification code, and any others are aliases.
	CsID []IdentifierType `xml:"http://www.opengis.net/gml csID"`
	// Comments on or information about this coordinate system, including data source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Ordered sequence of associations to the coordinate system axes included in this coordinate system.
	UsesAxis []CoordinateSystemAxisRefType `xml:"http://www.opengis.net/gml usesAxis"`
}

type LinearRingPropertyType struct {
	LinearRing *LinearRingType `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
}

type LinearRingType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string                 `xml:"uomLabels,attr,omitempty"`
	PosList   *DirectPositionListType `xml:"http://www.opengis.net/gml posList,omitempty"`
	// Deprecated with GML version 3.1.0. Use "posList" instead.
	Coordinates *CoordinatesType `xml:"http://www.opengis.net/gml coordinates,omitempty"`
	// Deprecated with GML version 3.0 and included for backwards compatibility with GML 2. Use "pos" elements instead.
	Coord []CoordType          `xml:"http://www.opengis.net/gml coord"`
	Pos   []DirectPositionType `xml:"http://www.opengis.net/gml pos"`
	// This property element either references a point via the XLink-attributes or contains the point element. pointProperty
	// is the predefined property which can be used by GML Application Schemas whenever a GML Feature has a property with a value that
	// is substitutable for Point.
	PointProperty []PointPropertyType `xml:"http://www.opengis.net/gml pointProperty"`
	// Deprecated with GML version 3.1.0. Use "pointProperty" instead. Included for backwards compatibility with GML 3.0.0.
	PointRep []PointPropertyType `xml:"http://www.opengis.net/gml pointRep"`
}

type LocationPropertyType struct {
	// The "_Geometry" element is the abstract head of the substituition group for all geometry elements of GML 3. This
	// includes pre-defined and user-defined geometry elements. Any geometry element must be a direct or indirect extension/restriction
	// of AbstractGeometryType and must be directly or indirectly in the substitution group of "_Geometry".
	AbstractGeometry           *AbstractGeometryType           `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	CompositeCurve             *CompositeCurveType             `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	CompositeSolid             *CompositeSolidType             `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	CompositeSurface           *CompositeSurfaceType           `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	Curve                      *CurveType                      `xml:"http://www.opengis.net/gml Curve,omitempty"`
	GeometricComplex           *GeometricComplexType           `xml:"http://www.opengis.net/gml GeometricComplex,omitempty"`
	Grid                       *GridType                       `xml:"http://www.opengis.net/gml Grid,omitempty"`
	LineString                 *LineStringType                 `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearRing                 *LinearRingType                 `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MultiCurve                 *MultiCurveType                 `xml:"http://www.opengis.net/gml MultiCurve,omitempty"`
	MultiGeometry              *MultiGeometryType              `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString            *MultiLineStringType            `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                 *MultiPointType                 `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPolygon               *MultiPolygonType               `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	MultiSolid                 *MultiSolidType                 `xml:"http://www.opengis.net/gml MultiSolid,omitempty"`
	MultiSurface               *MultiSurfaceType               `xml:"http://www.opengis.net/gml MultiSurface,omitempty"`
	OrientableCurve            *OrientableCurveType            `xml:"http://www.opengis.net/gml OrientableCurve,omitempty"`
	OrientableSurface          *OrientableSurfaceType          `xml:"http://www.opengis.net/gml OrientableSurface,omitempty"`
	Point                      *PointType                      `xml:"http://www.opengis.net/gml Point,omitempty"`
	Polygon                    *PolygonType                    `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	PolyhedralSurface          *PolyhedralSurfaceType          `xml:"http://www.opengis.net/gml PolyhedralSurface,omitempty"`
	RectifiedGrid              *RectifiedGridType              `xml:"http://www.opengis.net/gml RectifiedGrid,omitempty"`
	Ring                       *RingType                       `xml:"http://www.opengis.net/gml Ring,omitempty"`
	Solid                      *SolidType                      `xml:"http://www.opengis.net/gml Solid,omitempty"`
	Surface                    *SurfaceType                    `xml:"http://www.opengis.net/gml Surface,omitempty"`
	Tin                        *TinType                        `xml:"http://www.opengis.net/gml Tin,omitempty"`
	TriangulatedSurface        *TriangulatedSurfaceType        `xml:"http://www.opengis.net/gml TriangulatedSurface,omitempty"`
	AbstractCurve              *AbstractCurveType              `xml:"http://www.opengis.net/gml _Curve,omitempty"`
	AbstractGeometricAggregate *AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml _GeometricAggregate,omitempty"`
	AbstractGeometricPrimitive *AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml _GeometricPrimitive,omitempty"`
	AbstractImplicitGeometry   *AbstractGeometryType           `xml:"http://www.opengis.net/gml _ImplicitGeometry,omitempty"`
	AbstractRing               *AbstractRingType               `xml:"http://www.opengis.net/gml _Ring,omitempty"`
	AbstractSolid              *AbstractSolidType              `xml:"http://www.opengis.net/gml _Solid,omitempty"`
	AbstractSurface            *AbstractSurfaceType            `xml:"http://www.opengis.net/gml _Surface,omitempty"`
	LocationKeyWord            *CodeType                       `xml:"http://www.opengis.net/gml LocationKeyWord,omitempty"`
	LocationString             *StringOrRefType                `xml:"http://www.opengis.net/gml LocationString,omitempty"`
	Null                       *string                         `xml:"http://www.opengis.net/gml Null,omitempty"`
	RemoteSchema               string                          `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                  string                          `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                       string                          `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                       string                          `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                    string                          `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                      string                          `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                       string                          `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                    string                          `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type MeasureListType struct {
	Value string `xml:",chardata"`
	Uom   string `xml:"uom,attr"`
}

type MeasureOrNullListType struct {
	Value string `xml:",chardata"`
	Uom   string `xml:"uom,attr"`
}

type MeasureType struct {
	Value float64 `xml:",chardata"`
	Uom   string  `xml:"uom,attr"`
}

type MetaDataPropertyType struct {
	About        *string `xml:"about,attr,omitempty"`
	RemoteSchema string  `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string  `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string  `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string  `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string  `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string  `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string  `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string  `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type MovingObjectStatusType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType                 `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType                 `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType                 `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType                 `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType                 `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType                 `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType                 `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType                 `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType                 `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType                 `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	ValidTime               *TimePrimitivePropertyType `xml:"http://www.opengis.net/gml validTime,omitempty"`
	DataSource              *StringOrRefType           `xml:"http://www.opengis.net/gml dataSource,omitempty"`
	// Deprecated in GML 3.1.0
	Location         *LocationPropertyType         `xml:"http://www.opengis.net/gml location,omitempty"`
	PriorityLocation *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml priorityLocation,omitempty"`
	Speed            *MeasureType                  `xml:"http://www.opengis.net/gml speed,omitempty"`
	Bearing          *DirectionPropertyType        `xml:"http://www.opengis.net/gml bearing,omitempty"`
	Acceleration     *MeasureType                  `xml:"http://www.opengis.net/gml acceleration,omitempty"`
	Elevation        *MeasureType                  `xml:"http://www.opengis.net/gml elevation,omitempty"`
	Status           *StringOrRefType              `xml:"http://www.opengis.net/gml status,omitempty"`
}

type MultiCurveCoverageType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType         `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType         `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType         `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType         `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType         `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType         `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType         `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType         `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType         `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType         `xml:"http://www.opengis.net/gml srsName"`
	Id                      string             `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	BoundedBy               *BoundingShapeType `xml:"http://www.opengis.net/gml boundedBy,omitempty"`
	// deprecated in GML version 3.1
	Location            *LocationPropertyType         `xml:"http://www.opengis.net/gml location,omitempty"`
	PriorityLocation    *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml priorityLocation,omitempty"`
	DomainSet           *DomainSetType                `xml:"http://www.opengis.net/gml domainSet,omitempty"`
	GridDomain          *GridDomainType               `xml:"http://www.opengis.net/gml gridDomain,omitempty"`
	MultiCurveDomain    *MultiCurveDomainType         `xml:"http://www.opengis.net/gml multiCurveDomain,omitempty"`
	MultiPointDomain    *MultiPointDomainType         `xml:"http://www.opengis.net/gml multiPointDomain,omitempty"`
	MultiSolidDomain    *MultiSolidDomainType         `xml:"http://www.opengis.net/gml multiSolidDomain,omitempty"`
	MultiSurfaceDomain  *MultiSurfaceDomainType       `xml:"http://www.opengis.net/gml multiSurfaceDomain,omitempty"`
	RectifiedGridDomain *RectifiedGridDomainType      `xml:"http://www.opengis.net/gml rectifiedGridDomain,omitempty"`
	RangeSet            *RangeSetType                 `xml:"http://www.opengis.net/gml rangeSet,omitempty"`
	Dimension           *int                          `xml:"dimension,attr,omitempty"`
	CoverageFunction    *CoverageFunctionType         `xml:"http://www.opengis.net/gml coverageFunction,omitempty"`
}

type MultiCurveDomainType struct {
	// The "_Geometry" element is the abstract head of the substituition group for all geometry elements of GML 3. This
	// includes pre-defined and user-defined geometry elements. Any geometry element must be a direct or indirect extension/restriction
	// of AbstractGeometryType and must be directly or indirectly in the substitution group of "_Geometry".
	AbstractGeometry           *AbstractGeometryType           `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	CompositeCurve             *CompositeCurveType             `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	CompositeSolid             *CompositeSolidType             `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	CompositeSurface           *CompositeSurfaceType           `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	Curve                      *CurveType                      `xml:"http://www.opengis.net/gml Curve,omitempty"`
	GeometricComplex           *GeometricComplexType           `xml:"http://www.opengis.net/gml GeometricComplex,omitempty"`
	Grid                       *GridType                       `xml:"http://www.opengis.net/gml Grid,omitempty"`
	LineString                 *LineStringType                 `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearRing                 *LinearRingType                 `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MultiCurve                 *MultiCurveType                 `xml:"http://www.opengis.net/gml MultiCurve,omitempty"`
	MultiGeometry              *MultiGeometryType              `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString            *MultiLineStringType            `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                 *MultiPointType                 `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPolygon               *MultiPolygonType               `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	MultiSolid                 *MultiSolidType                 `xml:"http://www.opengis.net/gml MultiSolid,omitempty"`
	MultiSurface               *MultiSurfaceType               `xml:"http://www.opengis.net/gml MultiSurface,omitempty"`
	OrientableCurve            *OrientableCurveType            `xml:"http://www.opengis.net/gml OrientableCurve,omitempty"`
	OrientableSurface          *OrientableSurfaceType          `xml:"http://www.opengis.net/gml OrientableSurface,omitempty"`
	Point                      *PointType                      `xml:"http://www.opengis.net/gml Point,omitempty"`
	Polygon                    *PolygonType                    `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	PolyhedralSurface          *PolyhedralSurfaceType          `xml:"http://www.opengis.net/gml PolyhedralSurface,omitempty"`
	RectifiedGrid              *RectifiedGridType              `xml:"http://www.opengis.net/gml RectifiedGrid,omitempty"`
	Ring                       *RingType                       `xml:"http://www.opengis.net/gml Ring,omitempty"`
	Solid                      *SolidType                      `xml:"http://www.opengis.net/gml Solid,omitempty"`
	Surface                    *SurfaceType                    `xml:"http://www.opengis.net/gml Surface,omitempty"`
	Tin                        *TinType                        `xml:"http://www.opengis.net/gml Tin,omitempty"`
	TriangulatedSurface        *TriangulatedSurfaceType        `xml:"http://www.opengis.net/gml TriangulatedSurface,omitempty"`
	AbstractCurve              *AbstractCurveType              `xml:"http://www.opengis.net/gml _Curve,omitempty"`
	AbstractGeometricAggregate *AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml _GeometricAggregate,omitempty"`
	AbstractGeometricPrimitive *AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml _GeometricPrimitive,omitempty"`
	AbstractImplicitGeometry   *AbstractGeometryType           `xml:"http://www.opengis.net/gml _ImplicitGeometry,omitempty"`
	AbstractRing               *AbstractRingType               `xml:"http://www.opengis.net/gml _Ring,omitempty"`
	AbstractSolid              *AbstractSolidType              `xml:"http://www.opengis.net/gml _Solid,omitempty"`
	AbstractSurface            *AbstractSurfaceType            `xml:"http://www.opengis.net/gml _Surface,omitempty"`
	// This abstract element acts as the head of the substitution group for temporal primitives and complexes.
	AbstractTimeObject             *AbstractTimeObjectType             `xml:"http://www.opengis.net/gml _TimeObject,omitempty"`
	TimeEdge                       *TimeEdgeType                       `xml:"http://www.opengis.net/gml TimeEdge,omitempty"`
	TimeInstant                    *TimeInstantType                    `xml:"http://www.opengis.net/gml TimeInstant,omitempty"`
	TimeNode                       *TimeNodeType                       `xml:"http://www.opengis.net/gml TimeNode,omitempty"`
	TimePeriod                     *TimePeriodType                     `xml:"http://www.opengis.net/gml TimePeriod,omitempty"`
	TimeTopologyComplex            *TimeTopologyComplexType            `xml:"http://www.opengis.net/gml TimeTopologyComplex,omitempty"`
	AbstractTimeComplex            *AbstractTimeComplexType            `xml:"http://www.opengis.net/gml _TimeComplex,omitempty"`
	AbstractTimeGeometricPrimitive *AbstractTimeGeometricPrimitiveType `xml:"http://www.opengis.net/gml _TimeGeometricPrimitive,omitempty"`
	AbstractTimePrimitive          *AbstractTimePrimitiveType          `xml:"http://www.opengis.net/gml _TimePrimitive,omitempty"`
	AbstractTimeTopologyPrimitive  *AbstractTimeTopologyPrimitiveType  `xml:"http://www.opengis.net/gml _TimeTopologyPrimitive,omitempty"`
	RemoteSchema                   string                              `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                      string                              `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                           string                              `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                           string                              `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                        string                              `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                          string                              `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                           string                              `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                        string                              `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type MultiCurvePropertyType struct {
	MultiCurve   *MultiCurveType `xml:"http://www.opengis.net/gml MultiCurve,omitempty"`
	RemoteSchema string          `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string          `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string          `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string          `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string          `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string          `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string          `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string          `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type MultiCurveType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
	// This property element either references a curve via the XLink-attributes or contains the curve element. A curve element is any element which is substitutable for "_Curve".
	CurveMember []CurvePropertyType `xml:"http://www.opengis.net/gml curveMember"`
	// This property element contains a list of curves. The order of the elements is significant and shall be preserved when processing the array.
	CurveMembers *CurveArrayPropertyType `xml:"http://www.opengis.net/gml curveMembers,omitempty"`
}

type MultiGeometryPropertyType struct {
	// The "_GeometricAggregate" element is the abstract head of the substituition group for all geometric aggremates.
	AbstractGeometricAggregate *AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml _GeometricAggregate,omitempty"`
	MultiCurve                 *MultiCurveType                 `xml:"http://www.opengis.net/gml MultiCurve,omitempty"`
	MultiGeometry              *MultiGeometryType              `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString            *MultiLineStringType            `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                 *MultiPointType                 `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPolygon               *MultiPolygonType               `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	MultiSolid                 *MultiSolidType                 `xml:"http://www.opengis.net/gml MultiSolid,omitempty"`
	MultiSurface               *MultiSurfaceType               `xml:"http://www.opengis.net/gml MultiSurface,omitempty"`
	RemoteSchema               string                          `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                  string                          `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                       string                          `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                       string                          `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                    string                          `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                      string                          `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                       string                          `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                    string                          `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type MultiGeometryType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
	// This property element either references a geometry element via the XLink-attributes or contains the geometry element.
	GeometryMember []GeometryPropertyType `xml:"http://www.opengis.net/gml geometryMember"`
	// This property element contains a list of geometry elements. The order of the elements is significant and shall be preserved when processing the array.
	GeometryMembers *GeometryArrayPropertyType `xml:"http://www.opengis.net/gml geometryMembers,omitempty"`
}

type MultiLineStringPropertyType struct {
	// Deprecated with GML 3.0 and included for backwards compatibility with GML 2. Use the "MultiCurve" element instead.
	MultiLineString *MultiLineStringType `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	RemoteSchema    string               `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField       string               `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href            string               `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role            string               `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole         string               `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title           string               `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show            string               `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate         string               `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type MultiLineStringType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
	// Deprecated with GML 3.0 and included only for backwards compatibility with GML 2.0. Use "curveMember" instead.
	// This property element either references a line string via the XLink-attributes or contains the line string element.
	LineStringMember []LineStringPropertyType `xml:"http://www.opengis.net/gml lineStringMember"`
}

type MultiPointCoverageType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType         `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType         `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType         `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType         `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType         `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType         `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType         `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType         `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType         `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType         `xml:"http://www.opengis.net/gml srsName"`
	Id                      string             `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	BoundedBy               *BoundingShapeType `xml:"http://www.opengis.net/gml boundedBy,omitempty"`
	// deprecated in GML version 3.1
	Location            *LocationPropertyType         `xml:"http://www.opengis.net/gml location,omitempty"`
	PriorityLocation    *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml priorityLocation,omitempty"`
	DomainSet           *DomainSetType                `xml:"http://www.opengis.net/gml domainSet,omitempty"`
	GridDomain          *GridDomainType               `xml:"http://www.opengis.net/gml gridDomain,omitempty"`
	MultiCurveDomain    *MultiCurveDomainType         `xml:"http://www.opengis.net/gml multiCurveDomain,omitempty"`
	MultiPointDomain    *MultiPointDomainType         `xml:"http://www.opengis.net/gml multiPointDomain,omitempty"`
	MultiSolidDomain    *MultiSolidDomainType         `xml:"http://www.opengis.net/gml multiSolidDomain,omitempty"`
	MultiSurfaceDomain  *MultiSurfaceDomainType       `xml:"http://www.opengis.net/gml multiSurfaceDomain,omitempty"`
	RectifiedGridDomain *RectifiedGridDomainType      `xml:"http://www.opengis.net/gml rectifiedGridDomain,omitempty"`
	RangeSet            *RangeSetType                 `xml:"http://www.opengis.net/gml rangeSet,omitempty"`
	Dimension           *int                          `xml:"dimension,attr,omitempty"`
	CoverageFunction    *CoverageFunctionType         `xml:"http://www.opengis.net/gml coverageFunction,omitempty"`
}

type MultiPointDomainType struct {
	// The "_Geometry" element is the abstract head of the substituition group for all geometry elements of GML 3. This
	// includes pre-defined and user-defined geometry elements. Any geometry element must be a direct or indirect extension/restriction
	// of AbstractGeometryType and must be directly or indirectly in the substitution group of "_Geometry".
	AbstractGeometry           *AbstractGeometryType           `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	CompositeCurve             *CompositeCurveType             `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	CompositeSolid             *CompositeSolidType             `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	CompositeSurface           *CompositeSurfaceType           `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	Curve                      *CurveType                      `xml:"http://www.opengis.net/gml Curve,omitempty"`
	GeometricComplex           *GeometricComplexType           `xml:"http://www.opengis.net/gml GeometricComplex,omitempty"`
	Grid                       *GridType                       `xml:"http://www.opengis.net/gml Grid,omitempty"`
	LineString                 *LineStringType                 `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearRing                 *LinearRingType                 `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MultiCurve                 *MultiCurveType                 `xml:"http://www.opengis.net/gml MultiCurve,omitempty"`
	MultiGeometry              *MultiGeometryType              `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString            *MultiLineStringType            `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                 *MultiPointType                 `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPolygon               *MultiPolygonType               `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	MultiSolid                 *MultiSolidType                 `xml:"http://www.opengis.net/gml MultiSolid,omitempty"`
	MultiSurface               *MultiSurfaceType               `xml:"http://www.opengis.net/gml MultiSurface,omitempty"`
	OrientableCurve            *OrientableCurveType            `xml:"http://www.opengis.net/gml OrientableCurve,omitempty"`
	OrientableSurface          *OrientableSurfaceType          `xml:"http://www.opengis.net/gml OrientableSurface,omitempty"`
	Point                      *PointType                      `xml:"http://www.opengis.net/gml Point,omitempty"`
	Polygon                    *PolygonType                    `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	PolyhedralSurface          *PolyhedralSurfaceType          `xml:"http://www.opengis.net/gml PolyhedralSurface,omitempty"`
	RectifiedGrid              *RectifiedGridType              `xml:"http://www.opengis.net/gml RectifiedGrid,omitempty"`
	Ring                       *RingType                       `xml:"http://www.opengis.net/gml Ring,omitempty"`
	Solid                      *SolidType                      `xml:"http://www.opengis.net/gml Solid,omitempty"`
	Surface                    *SurfaceType                    `xml:"http://www.opengis.net/gml Surface,omitempty"`
	Tin                        *TinType                        `xml:"http://www.opengis.net/gml Tin,omitempty"`
	TriangulatedSurface        *TriangulatedSurfaceType        `xml:"http://www.opengis.net/gml TriangulatedSurface,omitempty"`
	AbstractCurve              *AbstractCurveType              `xml:"http://www.opengis.net/gml _Curve,omitempty"`
	AbstractGeometricAggregate *AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml _GeometricAggregate,omitempty"`
	AbstractGeometricPrimitive *AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml _GeometricPrimitive,omitempty"`
	AbstractImplicitGeometry   *AbstractGeometryType           `xml:"http://www.opengis.net/gml _ImplicitGeometry,omitempty"`
	AbstractRing               *AbstractRingType               `xml:"http://www.opengis.net/gml _Ring,omitempty"`
	AbstractSolid              *AbstractSolidType              `xml:"http://www.opengis.net/gml _Solid,omitempty"`
	AbstractSurface            *AbstractSurfaceType            `xml:"http://www.opengis.net/gml _Surface,omitempty"`
	// This abstract element acts as the head of the substitution group for temporal primitives and complexes.
	AbstractTimeObject             *AbstractTimeObjectType             `xml:"http://www.opengis.net/gml _TimeObject,omitempty"`
	TimeEdge                       *TimeEdgeType                       `xml:"http://www.opengis.net/gml TimeEdge,omitempty"`
	TimeInstant                    *TimeInstantType                    `xml:"http://www.opengis.net/gml TimeInstant,omitempty"`
	TimeNode                       *TimeNodeType                       `xml:"http://www.opengis.net/gml TimeNode,omitempty"`
	TimePeriod                     *TimePeriodType                     `xml:"http://www.opengis.net/gml TimePeriod,omitempty"`
	TimeTopologyComplex            *TimeTopologyComplexType            `xml:"http://www.opengis.net/gml TimeTopologyComplex,omitempty"`
	AbstractTimeComplex            *AbstractTimeComplexType            `xml:"http://www.opengis.net/gml _TimeComplex,omitempty"`
	AbstractTimeGeometricPrimitive *AbstractTimeGeometricPrimitiveType `xml:"http://www.opengis.net/gml _TimeGeometricPrimitive,omitempty"`
	AbstractTimePrimitive          *AbstractTimePrimitiveType          `xml:"http://www.opengis.net/gml _TimePrimitive,omitempty"`
	AbstractTimeTopologyPrimitive  *AbstractTimeTopologyPrimitiveType  `xml:"http://www.opengis.net/gml _TimeTopologyPrimitive,omitempty"`
	RemoteSchema                   string                              `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                      string                              `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                           string                              `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                           string                              `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                        string                              `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                          string                              `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                           string                              `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                        string                              `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type MultiPointPropertyType struct {
	MultiPoint   *MultiPointType `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	RemoteSchema string          `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string          `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string          `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string          `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string          `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string          `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string          `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string          `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type MultiPointType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
	// This property element either references a Point via the XLink-attributes or contains the Point element.
	PointMember []PointPropertyType `xml:"http://www.opengis.net/gml pointMember"`
	// This property element contains a list of points. The order of the elements is significant and shall be preserved when processing the array.
	PointMembers *PointArrayPropertyType `xml:"http://www.opengis.net/gml pointMembers,omitempty"`
}

type MultiPolygonPropertyType struct {
	// Deprecated with GML 3.0 and included for backwards compatibility with GML 2. Use the "MultiSurface" element instead.
	MultiPolygon *MultiPolygonType `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	RemoteSchema string            `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string            `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string            `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string            `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string            `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string            `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string            `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string            `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type MultiPolygonType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
	// Deprecated with GML 3.0 and included only for backwards compatibility with GML 2.0. Use "surfaceMember" instead.
	// This property element either references a polygon via the XLink-attributes or contains the polygon element.
	PolygonMember []PolygonPropertyType `xml:"http://www.opengis.net/gml polygonMember"`
}

type MultiSolidCoverageType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType         `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType         `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType         `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType         `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType         `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType         `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType         `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType         `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType         `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType         `xml:"http://www.opengis.net/gml srsName"`
	Id                      string             `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	BoundedBy               *BoundingShapeType `xml:"http://www.opengis.net/gml boundedBy,omitempty"`
	// deprecated in GML version 3.1
	Location            *LocationPropertyType         `xml:"http://www.opengis.net/gml location,omitempty"`
	PriorityLocation    *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml priorityLocation,omitempty"`
	DomainSet           *DomainSetType                `xml:"http://www.opengis.net/gml domainSet,omitempty"`
	GridDomain          *GridDomainType               `xml:"http://www.opengis.net/gml gridDomain,omitempty"`
	MultiCurveDomain    *MultiCurveDomainType         `xml:"http://www.opengis.net/gml multiCurveDomain,omitempty"`
	MultiPointDomain    *MultiPointDomainType         `xml:"http://www.opengis.net/gml multiPointDomain,omitempty"`
	MultiSolidDomain    *MultiSolidDomainType         `xml:"http://www.opengis.net/gml multiSolidDomain,omitempty"`
	MultiSurfaceDomain  *MultiSurfaceDomainType       `xml:"http://www.opengis.net/gml multiSurfaceDomain,omitempty"`
	RectifiedGridDomain *RectifiedGridDomainType      `xml:"http://www.opengis.net/gml rectifiedGridDomain,omitempty"`
	RangeSet            *RangeSetType                 `xml:"http://www.opengis.net/gml rangeSet,omitempty"`
	Dimension           *int                          `xml:"dimension,attr,omitempty"`
	CoverageFunction    *CoverageFunctionType         `xml:"http://www.opengis.net/gml coverageFunction,omitempty"`
}

type MultiSolidDomainType struct {
	// The "_Geometry" element is the abstract head of the substituition group for all geometry elements of GML 3. This
	// includes pre-defined and user-defined geometry elements. Any geometry element must be a direct or indirect extension/restriction
	// of AbstractGeometryType and must be directly or indirectly in the substitution group of "_Geometry".
	AbstractGeometry           *AbstractGeometryType           `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	CompositeCurve             *CompositeCurveType             `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	CompositeSolid             *CompositeSolidType             `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	CompositeSurface           *CompositeSurfaceType           `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	Curve                      *CurveType                      `xml:"http://www.opengis.net/gml Curve,omitempty"`
	GeometricComplex           *GeometricComplexType           `xml:"http://www.opengis.net/gml GeometricComplex,omitempty"`
	Grid                       *GridType                       `xml:"http://www.opengis.net/gml Grid,omitempty"`
	LineString                 *LineStringType                 `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearRing                 *LinearRingType                 `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MultiCurve                 *MultiCurveType                 `xml:"http://www.opengis.net/gml MultiCurve,omitempty"`
	MultiGeometry              *MultiGeometryType              `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString            *MultiLineStringType            `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                 *MultiPointType                 `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPolygon               *MultiPolygonType               `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	MultiSolid                 *MultiSolidType                 `xml:"http://www.opengis.net/gml MultiSolid,omitempty"`
	MultiSurface               *MultiSurfaceType               `xml:"http://www.opengis.net/gml MultiSurface,omitempty"`
	OrientableCurve            *OrientableCurveType            `xml:"http://www.opengis.net/gml OrientableCurve,omitempty"`
	OrientableSurface          *OrientableSurfaceType          `xml:"http://www.opengis.net/gml OrientableSurface,omitempty"`
	Point                      *PointType                      `xml:"http://www.opengis.net/gml Point,omitempty"`
	Polygon                    *PolygonType                    `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	PolyhedralSurface          *PolyhedralSurfaceType          `xml:"http://www.opengis.net/gml PolyhedralSurface,omitempty"`
	RectifiedGrid              *RectifiedGridType              `xml:"http://www.opengis.net/gml RectifiedGrid,omitempty"`
	Ring                       *RingType                       `xml:"http://www.opengis.net/gml Ring,omitempty"`
	Solid                      *SolidType                      `xml:"http://www.opengis.net/gml Solid,omitempty"`
	Surface                    *SurfaceType                    `xml:"http://www.opengis.net/gml Surface,omitempty"`
	Tin                        *TinType                        `xml:"http://www.opengis.net/gml Tin,omitempty"`
	TriangulatedSurface        *TriangulatedSurfaceType        `xml:"http://www.opengis.net/gml TriangulatedSurface,omitempty"`
	AbstractCurve              *AbstractCurveType              `xml:"http://www.opengis.net/gml _Curve,omitempty"`
	AbstractGeometricAggregate *AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml _GeometricAggregate,omitempty"`
	AbstractGeometricPrimitive *AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml _GeometricPrimitive,omitempty"`
	AbstractImplicitGeometry   *AbstractGeometryType           `xml:"http://www.opengis.net/gml _ImplicitGeometry,omitempty"`
	AbstractRing               *AbstractRingType               `xml:"http://www.opengis.net/gml _Ring,omitempty"`
	AbstractSolid              *AbstractSolidType              `xml:"http://www.opengis.net/gml _Solid,omitempty"`
	AbstractSurface            *AbstractSurfaceType            `xml:"http://www.opengis.net/gml _Surface,omitempty"`
	// This abstract element acts as the head of the substitution group for temporal primitives and complexes.
	AbstractTimeObject             *AbstractTimeObjectType             `xml:"http://www.opengis.net/gml _TimeObject,omitempty"`
	TimeEdge                       *TimeEdgeType                       `xml:"http://www.opengis.net/gml TimeEdge,omitempty"`
	TimeInstant                    *TimeInstantType                    `xml:"http://www.opengis.net/gml TimeInstant,omitempty"`
	TimeNode                       *TimeNodeType                       `xml:"http://www.opengis.net/gml TimeNode,omitempty"`
	TimePeriod                     *TimePeriodType                     `xml:"http://www.opengis.net/gml TimePeriod,omitempty"`
	TimeTopologyComplex            *TimeTopologyComplexType            `xml:"http://www.opengis.net/gml TimeTopologyComplex,omitempty"`
	AbstractTimeComplex            *AbstractTimeComplexType            `xml:"http://www.opengis.net/gml _TimeComplex,omitempty"`
	AbstractTimeGeometricPrimitive *AbstractTimeGeometricPrimitiveType `xml:"http://www.opengis.net/gml _TimeGeometricPrimitive,omitempty"`
	AbstractTimePrimitive          *AbstractTimePrimitiveType          `xml:"http://www.opengis.net/gml _TimePrimitive,omitempty"`
	AbstractTimeTopologyPrimitive  *AbstractTimeTopologyPrimitiveType  `xml:"http://www.opengis.net/gml _TimeTopologyPrimitive,omitempty"`
	RemoteSchema                   string                              `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                      string                              `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                           string                              `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                           string                              `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                        string                              `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                          string                              `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                           string                              `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                        string                              `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type MultiSolidPropertyType struct {
	MultiSolid   *MultiSolidType `xml:"http://www.opengis.net/gml MultiSolid,omitempty"`
	RemoteSchema string          `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string          `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string          `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string          `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string          `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string          `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string          `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string          `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type MultiSolidType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
	// This property element either references a solid via the XLink-attributes or contains the solid element. A solid element is any element which is substitutable for "_Solid".
	SolidMember []SolidPropertyType `xml:"http://www.opengis.net/gml solidMember"`
	// This property element contains a list of solids. The order of the elements is significant and shall be preserved when processing the array.
	SolidMembers *SolidArrayPropertyType `xml:"http://www.opengis.net/gml solidMembers,omitempty"`
}

type MultiSurfaceCoverageType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType         `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType         `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType         `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType         `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType         `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType         `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType         `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType         `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType         `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType         `xml:"http://www.opengis.net/gml srsName"`
	Id                      string             `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	BoundedBy               *BoundingShapeType `xml:"http://www.opengis.net/gml boundedBy,omitempty"`
	// deprecated in GML version 3.1
	Location            *LocationPropertyType         `xml:"http://www.opengis.net/gml location,omitempty"`
	PriorityLocation    *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml priorityLocation,omitempty"`
	DomainSet           *DomainSetType                `xml:"http://www.opengis.net/gml domainSet,omitempty"`
	GridDomain          *GridDomainType               `xml:"http://www.opengis.net/gml gridDomain,omitempty"`
	MultiCurveDomain    *MultiCurveDomainType         `xml:"http://www.opengis.net/gml multiCurveDomain,omitempty"`
	MultiPointDomain    *MultiPointDomainType         `xml:"http://www.opengis.net/gml multiPointDomain,omitempty"`
	MultiSolidDomain    *MultiSolidDomainType         `xml:"http://www.opengis.net/gml multiSolidDomain,omitempty"`
	MultiSurfaceDomain  *MultiSurfaceDomainType       `xml:"http://www.opengis.net/gml multiSurfaceDomain,omitempty"`
	RectifiedGridDomain *RectifiedGridDomainType      `xml:"http://www.opengis.net/gml rectifiedGridDomain,omitempty"`
	RangeSet            *RangeSetType                 `xml:"http://www.opengis.net/gml rangeSet,omitempty"`
	Dimension           *int                          `xml:"dimension,attr,omitempty"`
	CoverageFunction    *CoverageFunctionType         `xml:"http://www.opengis.net/gml coverageFunction,omitempty"`
}

type MultiSurfaceDomainType struct {
	// The "_Geometry" element is the abstract head of the substituition group for all geometry elements of GML 3. This
	// includes pre-defined and user-defined geometry elements. Any geometry element must be a direct or indirect extension/restriction
	// of AbstractGeometryType and must be directly or indirectly in the substitution group of "_Geometry".
	AbstractGeometry           *AbstractGeometryType           `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	CompositeCurve             *CompositeCurveType             `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	CompositeSolid             *CompositeSolidType             `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	CompositeSurface           *CompositeSurfaceType           `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	Curve                      *CurveType                      `xml:"http://www.opengis.net/gml Curve,omitempty"`
	GeometricComplex           *GeometricComplexType           `xml:"http://www.opengis.net/gml GeometricComplex,omitempty"`
	Grid                       *GridType                       `xml:"http://www.opengis.net/gml Grid,omitempty"`
	LineString                 *LineStringType                 `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearRing                 *LinearRingType                 `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MultiCurve                 *MultiCurveType                 `xml:"http://www.opengis.net/gml MultiCurve,omitempty"`
	MultiGeometry              *MultiGeometryType              `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString            *MultiLineStringType            `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                 *MultiPointType                 `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPolygon               *MultiPolygonType               `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	MultiSolid                 *MultiSolidType                 `xml:"http://www.opengis.net/gml MultiSolid,omitempty"`
	MultiSurface               *MultiSurfaceType               `xml:"http://www.opengis.net/gml MultiSurface,omitempty"`
	OrientableCurve            *OrientableCurveType            `xml:"http://www.opengis.net/gml OrientableCurve,omitempty"`
	OrientableSurface          *OrientableSurfaceType          `xml:"http://www.opengis.net/gml OrientableSurface,omitempty"`
	Point                      *PointType                      `xml:"http://www.opengis.net/gml Point,omitempty"`
	Polygon                    *PolygonType                    `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	PolyhedralSurface          *PolyhedralSurfaceType          `xml:"http://www.opengis.net/gml PolyhedralSurface,omitempty"`
	RectifiedGrid              *RectifiedGridType              `xml:"http://www.opengis.net/gml RectifiedGrid,omitempty"`
	Ring                       *RingType                       `xml:"http://www.opengis.net/gml Ring,omitempty"`
	Solid                      *SolidType                      `xml:"http://www.opengis.net/gml Solid,omitempty"`
	Surface                    *SurfaceType                    `xml:"http://www.opengis.net/gml Surface,omitempty"`
	Tin                        *TinType                        `xml:"http://www.opengis.net/gml Tin,omitempty"`
	TriangulatedSurface        *TriangulatedSurfaceType        `xml:"http://www.opengis.net/gml TriangulatedSurface,omitempty"`
	AbstractCurve              *AbstractCurveType              `xml:"http://www.opengis.net/gml _Curve,omitempty"`
	AbstractGeometricAggregate *AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml _GeometricAggregate,omitempty"`
	AbstractGeometricPrimitive *AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml _GeometricPrimitive,omitempty"`
	AbstractImplicitGeometry   *AbstractGeometryType           `xml:"http://www.opengis.net/gml _ImplicitGeometry,omitempty"`
	AbstractRing               *AbstractRingType               `xml:"http://www.opengis.net/gml _Ring,omitempty"`
	AbstractSolid              *AbstractSolidType              `xml:"http://www.opengis.net/gml _Solid,omitempty"`
	AbstractSurface            *AbstractSurfaceType            `xml:"http://www.opengis.net/gml _Surface,omitempty"`
	// This abstract element acts as the head of the substitution group for temporal primitives and complexes.
	AbstractTimeObject             *AbstractTimeObjectType             `xml:"http://www.opengis.net/gml _TimeObject,omitempty"`
	TimeEdge                       *TimeEdgeType                       `xml:"http://www.opengis.net/gml TimeEdge,omitempty"`
	TimeInstant                    *TimeInstantType                    `xml:"http://www.opengis.net/gml TimeInstant,omitempty"`
	TimeNode                       *TimeNodeType                       `xml:"http://www.opengis.net/gml TimeNode,omitempty"`
	TimePeriod                     *TimePeriodType                     `xml:"http://www.opengis.net/gml TimePeriod,omitempty"`
	TimeTopologyComplex            *TimeTopologyComplexType            `xml:"http://www.opengis.net/gml TimeTopologyComplex,omitempty"`
	AbstractTimeComplex            *AbstractTimeComplexType            `xml:"http://www.opengis.net/gml _TimeComplex,omitempty"`
	AbstractTimeGeometricPrimitive *AbstractTimeGeometricPrimitiveType `xml:"http://www.opengis.net/gml _TimeGeometricPrimitive,omitempty"`
	AbstractTimePrimitive          *AbstractTimePrimitiveType          `xml:"http://www.opengis.net/gml _TimePrimitive,omitempty"`
	AbstractTimeTopologyPrimitive  *AbstractTimeTopologyPrimitiveType  `xml:"http://www.opengis.net/gml _TimeTopologyPrimitive,omitempty"`
	RemoteSchema                   string                              `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                      string                              `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                           string                              `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                           string                              `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                        string                              `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                          string                              `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                           string                              `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                        string                              `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type MultiSurfacePropertyType struct {
	MultiSurface *MultiSurfaceType `xml:"http://www.opengis.net/gml MultiSurface,omitempty"`
	RemoteSchema string            `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string            `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string            `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string            `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string            `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string            `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string            `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string            `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type MultiSurfaceType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
	// This property element either references a surface via the XLink-attributes or contains the surface element. A surface element is any element which is substitutable for "_Surface".
	SurfaceMember []SurfacePropertyType `xml:"http://www.opengis.net/gml surfaceMember"`
	// This property element contains a list of surfaces. The order of the elements is significant and shall be preserved when processing the array.
	SurfaceMembers *SurfaceArrayPropertyType `xml:"http://www.opengis.net/gml surfaceMembers,omitempty"`
}

type NodeType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType                 `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType                 `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType                 `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType                 `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType                 `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType                 `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType                 `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType                 `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType                 `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType                 `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	Isolated                []IsolatedPropertyType     `xml:"http://www.opengis.net/gml isolated"`
	Container               *ContainerPropertyType     `xml:"http://www.opengis.net/gml container,omitempty"`
	DirectedEdge            []DirectedEdgePropertyType `xml:"http://www.opengis.net/gml directedEdge"`
	// This property element either references a point via the XLink-attributes or contains the point element. pointProperty
	// is the predefined property which can be used by GML Application Schemas whenever a GML Feature has a property with a value that
	// is substitutable for Point.
	PointProperty *PointPropertyType `xml:"http://www.opengis.net/gml pointProperty,omitempty"`
}

type ObliqueCartesianCSRefType struct {
	ObliqueCartesianCS *ObliqueCartesianCSType `xml:"http://www.opengis.net/gml ObliqueCartesianCS,omitempty"`
	RemoteSchema       string                  `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField          string                  `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href               string                  `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role               string                  `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole            string                  `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title              string                  `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show               string                  `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate            string                  `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type ObliqueCartesianCSType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this coordinate system. The first csID, if any, is normally the primary identification code, and any others are aliases.
	CsID []IdentifierType `xml:"http://www.opengis.net/gml csID"`
	// Comments on or information about this coordinate system, including data source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Ordered sequence of associations to the coordinate system axes included in this coordinate system.
	UsesAxis []CoordinateSystemAxisRefType `xml:"http://www.opengis.net/gml usesAxis"`
}

type ObservationType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType         `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType         `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType         `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType         `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType         `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType         `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType         `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType         `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType         `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType         `xml:"http://www.opengis.net/gml srsName"`
	Id                      string             `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	BoundedBy               *BoundingShapeType `xml:"http://www.opengis.net/gml boundedBy,omitempty"`
	// deprecated in GML version 3.1
	Location         *LocationPropertyType         `xml:"http://www.opengis.net/gml location,omitempty"`
	PriorityLocation *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml priorityLocation,omitempty"`
	ValidTime        *TimePrimitivePropertyType    `xml:"http://www.opengis.net/gml validTime,omitempty"`
	// This element contains or points to a description of a sensor, instrument or procedure used for the observation
	Using *FeaturePropertyType `xml:"http://www.opengis.net/gml using,omitempty"`
	// This element contains or points to the specimen, region or station which is the object of the observation
	Target  *TargetPropertyType `xml:"http://www.opengis.net/gml target,omitempty"`
	Subject *TargetPropertyType `xml:"http://www.opengis.net/gml subject,omitempty"`
	// The result of the observation: an image, external object, etc
	ResultOf *AssociationType `xml:"http://www.opengis.net/gml resultOf,omitempty"`
}

type OffsetCurveType struct {
	// The attribute "numDerivativesAtStart" specifies the type of continuity between this curve segment and its predecessor. If this is the first curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtStart *int `xml:"numDerivativesAtStart,attr,omitempty"`
	// The attribute "numDerivativesAtEnd" specifies the type of continuity between this curve segment and its successor. If this is the last curve segment in the curve, one of these values, as appropriate, is ignored. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativesAtEnd *int `xml:"numDerivativesAtEnd,attr,omitempty"`
	// The attribute "numDerivativesInterior" specifies the type of continuity that is guaranteed interior to the curve. The default value of "0" means simple continuity, which is a mandatory minimum level of continuity. This level is referred to as "C 0 " in mathematical texts. A value of 1 means that the function and its first derivative are continuous at the appropriate end point: "C 1 " continuity. A value of "n" for any integer means the function and its first n derivatives are continuous: "C n " continuity.
	// NOTE: Use of these values is only appropriate when the basic curve definition is an underdetermined system. For example, line string segments cannot support continuity above C 0 , since there is no spare control parameter to adjust the incoming angle at the end points of the segment. Spline functions on the other hand often have extra degrees of freedom on end segments that allow them to adjust the values of the derivatives to support C 1 or higher continuity.
	NumDerivativeInterior *int `xml:"numDerivativeInterior,attr,omitempty"`
	// offsetBase is a reference to thecurve from which this
	// curve is define	as an offset.
	OffsetBase *CurvePropertyType `xml:"http://www.opengis.net/gml offsetBase,omitempty"`
	// distance is the distance at which the
	// offset curve is generated from the basis curve. In 2D systems, positive distances
	// are to be to the left of the basis curve, and the negative distances are to be to the
	// right of the basis curve.
	Distance *LengthType `xml:"http://www.opengis.net/gml distance,omitempty"`
	// refDistance is used to define the vector
	// direction of the offset curve from the basis curve. It can
	// be omitted in the 2D case, where the distance can be
	// positive or negative. In that case, distance defines left
	// side (positive distance) or right side (negative distance)
	// with respect to the tangent to the basis curve.
	//
	// In 3D the basis curve shall have a well defined tangent
	// direction for every point. The offset curve at any point
	// in 3D, the basis curve shall have a well-defined tangent
	// direction for every point. The offset curve at any point
	// (parameter) on the basis curve c is in the direction
	// -   -   -         -
	// s = v x t  where  v = c.refDirection()
	// and
	// -
	// t = c.tangent()
	// -
	// For the offset direction to be well-defined, v shall not
	// on any point of the curve be in the same, or opposite,
	// direction as
	// -
	// t.
	//
	// The default value of the refDirection shall be the local
	// co-ordinate axis vector for elevation, which indicates up for
	// the curve in a geographic sense.
	//
	// NOTE! If the refDirection is the positive tangent to the
	// local elevation axis ("points upward"), then the offset
	// vector points to the left of the curve when viewed from
	// above.
	RefDirection *VectorType `xml:"http://www.opengis.net/gml refDirection,omitempty"`
}

type OperationMethodBaseType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
}

type OperationMethodRefType struct {
	OperationMethod *OperationMethodType `xml:"http://www.opengis.net/gml OperationMethod,omitempty"`
	RemoteSchema    string               `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField       string               `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href            string               `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role            string               `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole         string               `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title           string               `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show            string               `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate         string               `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type OperationMethodType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this operation method. The first methodID, if any, is normally the primary identification code, and any others are aliases.
	MethodID []IdentifierType `xml:"http://www.opengis.net/gml methodID"`
	// Comments on or information about this operation method, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Formula(s) used by this operation method. The value may be a reference to a publication. Note that the operation method may not be analytic, in which case this element references or contains the procedure, not an analytic formula.
	MethodFormula *CodeType `xml:"http://www.opengis.net/gml methodFormula,omitempty"`
	// Number of dimensions in the source CRS of this operation method.
	SourceDimensions *int `xml:"http://www.opengis.net/gml sourceDimensions,omitempty"`
	// Number of dimensions in the target CRS of this operation method.
	TargetDimensions *int `xml:"http://www.opengis.net/gml targetDimensions,omitempty"`
	// Unordered list of associations to the set of operation parameters and parameter groups used by this operation method.
	UsesParameter []AbstractGeneralOperationParameterRefType `xml:"http://www.opengis.net/gml usesParameter"`
}

type OperationParameterBaseType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// The minimum number of times that values for this parameter group or parameter are required. If this attribute is omitted, the minimum number is one.
	MinimumOccurs *int `xml:"http://www.opengis.net/gml minimumOccurs,omitempty"`
}

type OperationParameterGroupBaseType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// The minimum number of times that values for this parameter group or parameter are required. If this attribute is omitted, the minimum number is one.
	MinimumOccurs *int `xml:"http://www.opengis.net/gml minimumOccurs,omitempty"`
}

type OperationParameterGroupRefType struct {
	OperationParameterGroup *OperationParameterGroupType `xml:"http://www.opengis.net/gml OperationParameterGroup,omitempty"`
	RemoteSchema            string                       `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField               string                       `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                    string                       `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                    string                       `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                 string                       `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                   string                       `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                    string                       `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                 string                       `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type OperationParameterGroupType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// The minimum number of times that values for this parameter group or parameter are required. If this attribute is omitted, the minimum number is one.
	MinimumOccurs *int `xml:"http://www.opengis.net/gml minimumOccurs,omitempty"`
	// Set of alternative identifications of this operation parameter group. The first groupID, if any, is normally the primary identification code, and any others are aliases.
	GroupID []IdentifierType `xml:"http://www.opengis.net/gml groupID"`
	// Comments on or information about this operation parameter group, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// The maximum number of times that values for this parameter group can be included. If this attribute is omitted, the maximum number is one.
	MaximumOccurs *int `xml:"http://www.opengis.net/gml maximumOccurs,omitempty"`
	// Unordered list of associations to the set of operation parameters that are members of this group.
	IncludesParameter []AbstractGeneralOperationParameterRefType `xml:"http://www.opengis.net/gml includesParameter"`
}

type OperationParameterRefType struct {
	OperationParameter *OperationParameterType `xml:"http://www.opengis.net/gml OperationParameter,omitempty"`
	RemoteSchema       string                  `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField          string                  `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href               string                  `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role               string                  `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole            string                  `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title              string                  `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show               string                  `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate            string                  `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type OperationParameterType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// The minimum number of times that values for this parameter group or parameter are required. If this attribute is omitted, the minimum number is one.
	MinimumOccurs *int `xml:"http://www.opengis.net/gml minimumOccurs,omitempty"`
	// Set of alternative identifications of this operation parameter. The first parameterID, if any, is normally the primary identification code, and any others are aliases.
	ParameterID []IdentifierType `xml:"http://www.opengis.net/gml parameterID"`
	// Comments on or information about this operation parameter, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
}

type OperationRefType struct {
	// A parameterized mathematical operation on coordinates that transforms or converts coordinates to another coordinate reference system. This coordinate operation uses an operation method, usually with associated parameter values. However, operation methods and parameter values are directly associated with concrete subtypes, not with this abstract type.
	//
	// This abstract complexType shall not be directly used, extended, or restricted in a compliant Application Schema.
	AbstractOperation             *AbstractCoordinateOperationType   `xml:"http://www.opengis.net/gml _Operation,omitempty"`
	Conversion                    *ConversionType                    `xml:"http://www.opengis.net/gml Conversion,omitempty"`
	Transformation                *TransformationType                `xml:"http://www.opengis.net/gml Transformation,omitempty"`
	AbstractGeneralConversion     *AbstractGeneralConversionType     `xml:"http://www.opengis.net/gml _GeneralConversion,omitempty"`
	AbstractGeneralTransformation *AbstractGeneralTransformationType `xml:"http://www.opengis.net/gml _GeneralTransformation,omitempty"`
	RemoteSchema                  string                             `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                     string                             `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                          string                             `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                          string                             `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                       string                             `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                         string                             `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                          string                             `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                       string                             `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type OrientableCurveType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
	// References or contains the base curve (positive orientation).
	// NOTE: This definition allows for a nested structure, i.e. an OrientableCurve may use another OrientableCurve as its base curve.
	BaseCurve *CurvePropertyType `xml:"http://www.opengis.net/gml baseCurve,omitempty"`
	// If the orientation is "+", then the OrientableCurve is identical to the baseCurve. If the orientation is "-", then the OrientableCurve is related to another _Curve with a parameterization that reverses the sense of the curve traversal. "+" is the default value.
	Orientation *string `xml:"orientation,attr,omitempty"`
}

type OrientableSurfaceType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
	// References or contains the base surface (positive orientation).
	BaseSurface *SurfacePropertyType `xml:"http://www.opengis.net/gml baseSurface,omitempty"`
	// If the orientation is "+", then the OrientableSurface is identical to the baseSurface. If the orientation is "-", then the OrientableSurface is a reference to a Surface with an up-normal that reverses the direction for this OrientableSurface, the sense of "the top of the surface". "+" is the default value.
	Orientation *string `xml:"orientation,attr,omitempty"`
}

type ParameterValueGroupType struct {
	// Unordered set of composition associations to the parameter values and groups of values included in this group.
	IncludesValue []AbstractGeneralParameterValueType `xml:"http://www.opengis.net/gml includesValue"`
	// Association to the operation parameter group for which this element provides parameter values.
	ValuesOfGroup *OperationParameterGroupRefType `xml:"http://www.opengis.net/gml valuesOfGroup,omitempty"`
}

type ParameterValueType struct {
	// Association to the operation parameter that this is a value of.
	ValueOfParameter *OperationParameterRefType `xml:"http://www.opengis.net/gml valueOfParameter,omitempty"`
	// Numeric value of an operation parameter, with its associated unit of measure.
	Value *MeasureType `xml:"http://www.opengis.net/gml value,omitempty"`
	// Value of an angle operation parameter, in either degree-minute-second format or single value format.
	DmsAngleValue *DMSAngleType `xml:"http://www.opengis.net/gml dmsAngleValue,omitempty"`
	// String value of an operation parameter. A string value does not have an associated unit of measure.
	StringValue *string `xml:"http://www.opengis.net/gml stringValue,omitempty"`
	// Positive integer value of an operation parameter, usually used for a count. An integer value does not have an associated unit of measure.
	IntegerValue *int `xml:"http://www.opengis.net/gml integerValue,omitempty"`
	// Boolean value of an operation parameter. A Boolean value does not have an associated unit of measure.
	BooleanValue *bool `xml:"http://www.opengis.net/gml booleanValue,omitempty"`
	// Ordered sequence of two or more numeric values of an operation parameter list, where each value has the same associated unit of measure. An element of this type contains a space-separated sequence of double values.
	ValueList *MeasureListType `xml:"http://www.opengis.net/gml valueList,omitempty"`
	// Ordered sequence of two or more integer values of an operation parameter list, usually used for counts. These integer values do not have an associated unit of measure. An element of this type contains a space-separated sequence of integer values.
	IntegerValueList *string `xml:"http://www.opengis.net/gml integerValueList,omitempty"`
	// Reference to a file or a part of a file containing one or more parameter values, each numeric value with its associated unit of measure. When referencing a part of a file, that file must contain multiple identified parts, such as an XML encoded document. Furthermore, the referenced file or part of a file can reference another part of the same or different files, as allowed in XML documents.
	ValueFile *string `xml:"http://www.opengis.net/gml valueFile,omitempty"`
}

type PassThroughOperationRefType struct {
	PassThroughOperation *PassThroughOperationType `xml:"http://www.opengis.net/gml PassThroughOperation,omitempty"`
	RemoteSchema         string                    `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField            string                    `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                 string                    `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                 string                    `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole              string                    `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                string                    `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                 string                    `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate              string                    `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type PassThroughOperationType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this coordinate operation. The first coordinateOperationID, if any, is normally the primary identification code, and any others are aliases.
	CoordinateOperationID []IdentifierType `xml:"http://www.opengis.net/gml coordinateOperationID"`
	// Comments on or information about this coordinate operation, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Version of the coordinate transformation (i.e., instantiation due to the stochastic nature of the parameters). Mandatory when describing a transformation, and should not be supplied for a conversion.
	OperationVersion *string `xml:"http://www.opengis.net/gml operationVersion,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
	// Unordered set of estimates of the impact of this coordinate operation on point position accuracy. Gives position error estimates for target coordinates of this coordinate operation, assuming no errors in source coordinates.
	AbstractPositionalAccuracy         []AbstractPositionalAccuracyType         `xml:"http://www.opengis.net/gml _positionalAccuracy"`
	AbsoluteExternalPositionalAccuracy []AbsoluteExternalPositionalAccuracyType `xml:"http://www.opengis.net/gml absoluteExternalPositionalAccuracy"`
	CovarianceMatrix                   []CovarianceMatrixType                   `xml:"http://www.opengis.net/gml covarianceMatrix"`
	RelativeInternalPositionalAccuracy []RelativeInternalPositionalAccuracyType `xml:"http://www.opengis.net/gml relativeInternalPositionalAccuracy"`
	// Association to the source CRS (coordinate reference system) of this coordinate operation.
	SourceCRS *CRSRefType `xml:"http://www.opengis.net/gml sourceCRS,omitempty"`
	// Association to the target CRS (coordinate reference system) of this coordinate operation. For constraints on multiplicity of "sourceCRS" and "targetCRS", see UML model of Coordinate Operation package in OGC Abstract Specification topic 2.
	TargetCRS *CRSRefType `xml:"http://www.opengis.net/gml targetCRS,omitempty"`
	// Ordered sequence of positive integers defining the positions in a coordinate tuple of the coordinates affected by this pass-through operation.
	ModifiedCoordinate []int `xml:"http://www.opengis.net/gml modifiedCoordinate"`
	// Association to the operation applied to the specified ordinates.
	UsesOperation *OperationRefType `xml:"http://www.opengis.net/gml usesOperation,omitempty"`
}

type PixelInCellType struct {
	Value string `xml:",chardata"`
	// Reference to a source of information specifying the values and meanings of all the allowed string values for this PixelInCellType.
	CodeSpace string `xml:"codeSpace,attr"`
}

type PointArrayPropertyType struct {
	Point []PointType `xml:"http://www.opengis.net/gml Point"`
}

type PointPropertyType struct {
	Point        *PointType `xml:"http://www.opengis.net/gml Point,omitempty"`
	RemoteSchema string     `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string     `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string     `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string     `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string     `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string     `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string     `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string     `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type PointType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string             `xml:"uomLabels,attr,omitempty"`
	Pos       *DirectPositionType `xml:"http://www.opengis.net/gml pos,omitempty"`
	// Deprecated with GML version 3.1.0 for coordinates with ordinate values that are numbers. Use "pos"
	// instead. The "coordinates" element shall only be used for coordinates with ordinates that require a string
	// representation, e.g. DMS representations.
	Coordinates *CoordinatesType `xml:"http://www.opengis.net/gml coordinates,omitempty"`
	// Deprecated with GML version 3.0. Use "pos" instead. The "coord" element is included for
	// backwards compatibility with GML 2.
	Coord *CoordType `xml:"http://www.opengis.net/gml coord,omitempty"`
}

type PolarCSRefType struct {
	PolarCS      *PolarCSType `xml:"http://www.opengis.net/gml PolarCS,omitempty"`
	RemoteSchema string       `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string       `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string       `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string       `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string       `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string       `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string       `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string       `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type PolarCSType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this coordinate system. The first csID, if any, is normally the primary identification code, and any others are aliases.
	CsID []IdentifierType `xml:"http://www.opengis.net/gml csID"`
	// Comments on or information about this coordinate system, including data source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Ordered sequence of associations to the coordinate system axes included in this coordinate system.
	UsesAxis []CoordinateSystemAxisRefType `xml:"http://www.opengis.net/gml usesAxis"`
}

type PolygonPatchArrayPropertyType struct {
	// The "_SurfacePatch" element is the abstract head of the substituition group for all surface pach elements describing a continuous portion of a surface.
	AbstractSurfacePatch           []AbstractSurfacePatchType           `xml:"http://www.opengis.net/gml _SurfacePatch"`
	Cone                           []ConeType                           `xml:"http://www.opengis.net/gml Cone"`
	Cylinder                       []CylinderType                       `xml:"http://www.opengis.net/gml Cylinder"`
	PolygonPatch                   []PolygonPatchType                   `xml:"http://www.opengis.net/gml PolygonPatch"`
	Rectangle                      []RectangleType                      `xml:"http://www.opengis.net/gml Rectangle"`
	Sphere                         []SphereType                         `xml:"http://www.opengis.net/gml Sphere"`
	Triangle                       []TriangleType                       `xml:"http://www.opengis.net/gml Triangle"`
	AbstractGriddedSurface         []AbstractGriddedSurfaceType         `xml:"http://www.opengis.net/gml _GriddedSurface"`
	AbstractParametricCurveSurface []AbstractParametricCurveSurfaceType `xml:"http://www.opengis.net/gml _ParametricCurveSurface"`
}

type PolygonPatchType struct {
	// A boundary of a surface consists of a number of rings. In the normal 2D case, one of these rings is distinguished as being the exterior boundary. In a general manifold this is not always possible, in which case all boundaries shall be listed as interior boundaries, and the exterior will be empty.
	Exterior        *AbstractRingPropertyType `xml:"http://www.opengis.net/gml exterior,omitempty"`
	OuterBoundaryIs *AbstractRingPropertyType `xml:"http://www.opengis.net/gml outerBoundaryIs,omitempty"`
	// A boundary of a surface consists of a number of rings. The "interior" rings seperate the surface / surface patch from the area enclosed by the rings.
	Interior        []AbstractRingPropertyType `xml:"http://www.opengis.net/gml interior"`
	InnerBoundaryIs []AbstractRingPropertyType `xml:"http://www.opengis.net/gml innerBoundaryIs"`
	// The attribute "interpolation" specifies the interpolation mechanism used for this surface patch. Currently only planar surface patches are defined in GML 3, the attribute is fixed to "planar", i.e. the interpolation method shall return points on a single plane. The boundary of the patch shall be contained within that plane.
	Interpolation *string `xml:"interpolation,attr,omitempty"`
}

type PolygonPropertyType struct {
	Polygon      *PolygonType `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	RemoteSchema string       `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string       `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string       `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string       `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string       `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string       `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string       `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string       `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type PolygonType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
	// A boundary of a surface consists of a number of rings. In the normal 2D case, one of these rings is distinguished as being the exterior boundary. In a general manifold this is not always possible, in which case all boundaries shall be listed as interior boundaries, and the exterior will be empty.
	Exterior        *AbstractRingPropertyType `xml:"http://www.opengis.net/gml exterior,omitempty"`
	OuterBoundaryIs *AbstractRingPropertyType `xml:"http://www.opengis.net/gml outerBoundaryIs,omitempty"`
	// A boundary of a surface consists of a number of rings. The "interior" rings seperate the surface / surface patch from the area enclosed by the rings.
	Interior        []AbstractRingPropertyType `xml:"http://www.opengis.net/gml interior"`
	InnerBoundaryIs []AbstractRingPropertyType `xml:"http://www.opengis.net/gml innerBoundaryIs"`
}

type PolyhedralSurfaceType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
	// This element encapsulates the patches of the surface.
	Patches         *SurfacePatchArrayPropertyType  `xml:"http://www.opengis.net/gml patches,omitempty"`
	PolygonPatches  *PolygonPatchArrayPropertyType  `xml:"http://www.opengis.net/gml polygonPatches,omitempty"`
	TrianglePatches *TrianglePatchArrayPropertyType `xml:"http://www.opengis.net/gml trianglePatches,omitempty"`
}

type PrimeMeridianBaseType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
}

type PrimeMeridianRefType struct {
	PrimeMeridian *PrimeMeridianType `xml:"http://www.opengis.net/gml PrimeMeridian,omitempty"`
	RemoteSchema  string             `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField     string             `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href          string             `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role          string             `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole       string             `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title         string             `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show          string             `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate       string             `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type PrimeMeridianType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this prime meridian. The first meridianID, if any, is normally the primary identification code, and any others are aliases.
	MeridianID []IdentifierType `xml:"http://www.opengis.net/gml meridianID"`
	// Comments on or information about this prime meridian, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Longitude of the prime meridian measured from the Greenwich meridian, positive eastward. The greenwichLongitude most common value is zero, and that value shall be used when the meridianName value is Greenwich.
	GreenwichLongitude *AngleChoiceType `xml:"http://www.opengis.net/gml greenwichLongitude,omitempty"`
}

type PriorityLocationPropertyType struct {
	// The "_Geometry" element is the abstract head of the substituition group for all geometry elements of GML 3. This
	// includes pre-defined and user-defined geometry elements. Any geometry element must be a direct or indirect extension/restriction
	// of AbstractGeometryType and must be directly or indirectly in the substitution group of "_Geometry".
	AbstractGeometry           *AbstractGeometryType           `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	CompositeCurve             *CompositeCurveType             `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	CompositeSolid             *CompositeSolidType             `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	CompositeSurface           *CompositeSurfaceType           `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	Curve                      *CurveType                      `xml:"http://www.opengis.net/gml Curve,omitempty"`
	GeometricComplex           *GeometricComplexType           `xml:"http://www.opengis.net/gml GeometricComplex,omitempty"`
	Grid                       *GridType                       `xml:"http://www.opengis.net/gml Grid,omitempty"`
	LineString                 *LineStringType                 `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearRing                 *LinearRingType                 `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MultiCurve                 *MultiCurveType                 `xml:"http://www.opengis.net/gml MultiCurve,omitempty"`
	MultiGeometry              *MultiGeometryType              `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString            *MultiLineStringType            `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                 *MultiPointType                 `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPolygon               *MultiPolygonType               `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	MultiSolid                 *MultiSolidType                 `xml:"http://www.opengis.net/gml MultiSolid,omitempty"`
	MultiSurface               *MultiSurfaceType               `xml:"http://www.opengis.net/gml MultiSurface,omitempty"`
	OrientableCurve            *OrientableCurveType            `xml:"http://www.opengis.net/gml OrientableCurve,omitempty"`
	OrientableSurface          *OrientableSurfaceType          `xml:"http://www.opengis.net/gml OrientableSurface,omitempty"`
	Point                      *PointType                      `xml:"http://www.opengis.net/gml Point,omitempty"`
	Polygon                    *PolygonType                    `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	PolyhedralSurface          *PolyhedralSurfaceType          `xml:"http://www.opengis.net/gml PolyhedralSurface,omitempty"`
	RectifiedGrid              *RectifiedGridType              `xml:"http://www.opengis.net/gml RectifiedGrid,omitempty"`
	Ring                       *RingType                       `xml:"http://www.opengis.net/gml Ring,omitempty"`
	Solid                      *SolidType                      `xml:"http://www.opengis.net/gml Solid,omitempty"`
	Surface                    *SurfaceType                    `xml:"http://www.opengis.net/gml Surface,omitempty"`
	Tin                        *TinType                        `xml:"http://www.opengis.net/gml Tin,omitempty"`
	TriangulatedSurface        *TriangulatedSurfaceType        `xml:"http://www.opengis.net/gml TriangulatedSurface,omitempty"`
	AbstractCurve              *AbstractCurveType              `xml:"http://www.opengis.net/gml _Curve,omitempty"`
	AbstractGeometricAggregate *AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml _GeometricAggregate,omitempty"`
	AbstractGeometricPrimitive *AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml _GeometricPrimitive,omitempty"`
	AbstractImplicitGeometry   *AbstractGeometryType           `xml:"http://www.opengis.net/gml _ImplicitGeometry,omitempty"`
	AbstractRing               *AbstractRingType               `xml:"http://www.opengis.net/gml _Ring,omitempty"`
	AbstractSolid              *AbstractSolidType              `xml:"http://www.opengis.net/gml _Solid,omitempty"`
	AbstractSurface            *AbstractSurfaceType            `xml:"http://www.opengis.net/gml _Surface,omitempty"`
	LocationKeyWord            *CodeType                       `xml:"http://www.opengis.net/gml LocationKeyWord,omitempty"`
	LocationString             *StringOrRefType                `xml:"http://www.opengis.net/gml LocationString,omitempty"`
	Null                       *string                         `xml:"http://www.opengis.net/gml Null,omitempty"`
	RemoteSchema               string                          `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                  string                          `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                       string                          `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                       string                          `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                    string                          `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                      string                          `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                       string                          `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                    string                          `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
	Priority                   *string                         `xml:"priority,attr,omitempty"`
}

type ProjectedCRSRefType struct {
	ProjectedCRS *ProjectedCRSType `xml:"http://www.opengis.net/gml ProjectedCRS,omitempty"`
	RemoteSchema string            `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string            `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string            `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string            `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string            `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string            `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string            `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string            `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type ProjectedCRSType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alterative identifications of this reference system. The first srsID, if any, is normally the primary identification code, and any others are aliases.
	SrsID []IdentifierType `xml:"http://www.opengis.net/gml srsID"`
	// Comments on or information about this reference system, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
	// Association to the coordinate reference system used by this derived CRS.
	BaseCRS *CoordinateReferenceSystemRefType `xml:"http://www.opengis.net/gml baseCRS,omitempty"`
	// Association to the coordinate conversion used to define this derived CRS.
	DefinedByConversion *GeneralConversionRefType `xml:"http://www.opengis.net/gml definedByConversion,omitempty"`
	// Association to the Cartesian coordinate system used by this CRS.
	UsesCartesianCS *CartesianCSRefType `xml:"http://www.opengis.net/gml usesCartesianCS,omitempty"`
}

type QuantityExtentType struct {
	Value string `xml:",chardata"`
}

type QuantityPropertyType struct {
	// This abstract element is the head of a substitutionGroup hierararchy which may contain either simpleContent or complexContent elements.  It is used to assert the model position of "class" elements declared in other GML schemas.
	AbstractObject                    *string                                `xml:"http://www.opengis.net/gml _Object,omitempty"`
	Array                             *ArrayType                             `xml:"http://www.opengis.net/gml Array,omitempty"`
	Bag                               *BagType                               `xml:"http://www.opengis.net/gml Bag,omitempty"`
	BaseUnit                          *BaseUnitType                          `xml:"http://www.opengis.net/gml BaseUnit,omitempty"`
	CartesianCS                       *CartesianCSType                       `xml:"http://www.opengis.net/gml CartesianCS,omitempty"`
	CompositeCurve                    *CompositeCurveType                    `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	CompositeSolid                    *CompositeSolidType                    `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	CompositeSurface                  *CompositeSurfaceType                  `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	CompoundCRS                       *CompoundCRSType                       `xml:"http://www.opengis.net/gml CompoundCRS,omitempty"`
	ConcatenatedOperation             *ConcatenatedOperationType             `xml:"http://www.opengis.net/gml ConcatenatedOperation,omitempty"`
	ConventionalUnit                  *ConventionalUnitType                  `xml:"http://www.opengis.net/gml ConventionalUnit,omitempty"`
	Conversion                        *ConversionType                        `xml:"http://www.opengis.net/gml Conversion,omitempty"`
	CoordinateSystemAxis              *CoordinateSystemAxisType              `xml:"http://www.opengis.net/gml CoordinateSystemAxis,omitempty"`
	Curve                             *CurveType                             `xml:"http://www.opengis.net/gml Curve,omitempty"`
	CylindricalCS                     *CylindricalCSType                     `xml:"http://www.opengis.net/gml CylindricalCS,omitempty"`
	Definition                        *DefinitionType                        `xml:"http://www.opengis.net/gml Definition,omitempty"`
	DefinitionCollection              *DictionaryType                        `xml:"http://www.opengis.net/gml DefinitionCollection,omitempty"`
	DefinitionProxy                   *DefinitionProxyType                   `xml:"http://www.opengis.net/gml DefinitionProxy,omitempty"`
	DerivedCRS                        *DerivedCRSType                        `xml:"http://www.opengis.net/gml DerivedCRS,omitempty"`
	DerivedUnit                       *DerivedUnitType                       `xml:"http://www.opengis.net/gml DerivedUnit,omitempty"`
	Dictionary                        *DictionaryType                        `xml:"http://www.opengis.net/gml Dictionary,omitempty"`
	DirectedObservation               *DirectedObservationType               `xml:"http://www.opengis.net/gml DirectedObservation,omitempty"`
	DirectedObservationAtDistance     *DirectedObservationAtDistanceType     `xml:"http://www.opengis.net/gml DirectedObservationAtDistance,omitempty"`
	Edge                              *EdgeType                              `xml:"http://www.opengis.net/gml Edge,omitempty"`
	Ellipsoid                         *EllipsoidType                         `xml:"http://www.opengis.net/gml Ellipsoid,omitempty"`
	EllipsoidalCS                     *EllipsoidalCSType                     `xml:"http://www.opengis.net/gml EllipsoidalCS,omitempty"`
	EngineeringCRS                    *EngineeringCRSType                    `xml:"http://www.opengis.net/gml EngineeringCRS,omitempty"`
	EngineeringDatum                  *EngineeringDatumType                  `xml:"http://www.opengis.net/gml EngineeringDatum,omitempty"`
	Face                              *FaceType                              `xml:"http://www.opengis.net/gml Face,omitempty"`
	FeatureCollection                 *FeatureCollectionType                 `xml:"http://www.opengis.net/gml FeatureCollection,omitempty"`
	FeatureStyle                      *FeatureStyleType                      `xml:"http://www.opengis.net/gml FeatureStyle,omitempty"`
	GenericMetaData                   *GenericMetaDataType                   `xml:"http://www.opengis.net/gml GenericMetaData,omitempty"`
	GeocentricCRS                     *GeocentricCRSType                     `xml:"http://www.opengis.net/gml GeocentricCRS,omitempty"`
	GeodeticDatum                     *GeodeticDatumType                     `xml:"http://www.opengis.net/gml GeodeticDatum,omitempty"`
	GeographicCRS                     *GeographicCRSType                     `xml:"http://www.opengis.net/gml GeographicCRS,omitempty"`
	GeometricComplex                  *GeometricComplexType                  `xml:"http://www.opengis.net/gml GeometricComplex,omitempty"`
	GeometryStyle                     *GeometryStyleType                     `xml:"http://www.opengis.net/gml GeometryStyle,omitempty"`
	GraphStyle                        *GraphStyleType                        `xml:"http://www.opengis.net/gml GraphStyle,omitempty"`
	Grid                              *GridType                              `xml:"http://www.opengis.net/gml Grid,omitempty"`
	GridCoverage                      *GridCoverageType                      `xml:"http://www.opengis.net/gml GridCoverage,omitempty"`
	ImageCRS                          *ImageCRSType                          `xml:"http://www.opengis.net/gml ImageCRS,omitempty"`
	ImageDatum                        *ImageDatumType                        `xml:"http://www.opengis.net/gml ImageDatum,omitempty"`
	LabelStyle                        *LabelStyleType                        `xml:"http://www.opengis.net/gml LabelStyle,omitempty"`
	LineString                        *LineStringType                        `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearCS                          *LinearCSType                          `xml:"http://www.opengis.net/gml LinearCS,omitempty"`
	LinearRing                        *LinearRingType                        `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MovingObjectStatus                *MovingObjectStatusType                `xml:"http://www.opengis.net/gml MovingObjectStatus,omitempty"`
	MultiCurve                        *MultiCurveType                        `xml:"http://www.opengis.net/gml MultiCurve,omitempty"`
	MultiCurveCoverage                *MultiCurveCoverageType                `xml:"http://www.opengis.net/gml MultiCurveCoverage,omitempty"`
	MultiGeometry                     *MultiGeometryType                     `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString                   *MultiLineStringType                   `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                        *MultiPointType                        `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPointCoverage                *MultiPointCoverageType                `xml:"http://www.opengis.net/gml MultiPointCoverage,omitempty"`
	MultiPolygon                      *MultiPolygonType                      `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	MultiSolid                        *MultiSolidType                        `xml:"http://www.opengis.net/gml MultiSolid,omitempty"`
	MultiSolidCoverage                *MultiSolidCoverageType                `xml:"http://www.opengis.net/gml MultiSolidCoverage,omitempty"`
	MultiSurface                      *MultiSurfaceType                      `xml:"http://www.opengis.net/gml MultiSurface,omitempty"`
	MultiSurfaceCoverage              *MultiSurfaceCoverageType              `xml:"http://www.opengis.net/gml MultiSurfaceCoverage,omitempty"`
	Node                              *NodeType                              `xml:"http://www.opengis.net/gml Node,omitempty"`
	ObliqueCartesianCS                *ObliqueCartesianCSType                `xml:"http://www.opengis.net/gml ObliqueCartesianCS,omitempty"`
	Observation                       *ObservationType                       `xml:"http://www.opengis.net/gml Observation,omitempty"`
	OperationMethod                   *OperationMethodType                   `xml:"http://www.opengis.net/gml OperationMethod,omitempty"`
	OperationParameter                *OperationParameterType                `xml:"http://www.opengis.net/gml OperationParameter,omitempty"`
	OperationParameterGroup           *OperationParameterGroupType           `xml:"http://www.opengis.net/gml OperationParameterGroup,omitempty"`
	OrientableCurve                   *OrientableCurveType                   `xml:"http://www.opengis.net/gml OrientableCurve,omitempty"`
	OrientableSurface                 *OrientableSurfaceType                 `xml:"http://www.opengis.net/gml OrientableSurface,omitempty"`
	PassThroughOperation              *PassThroughOperationType              `xml:"http://www.opengis.net/gml PassThroughOperation,omitempty"`
	Point                             *PointType                             `xml:"http://www.opengis.net/gml Point,omitempty"`
	PolarCS                           *PolarCSType                           `xml:"http://www.opengis.net/gml PolarCS,omitempty"`
	Polygon                           *PolygonType                           `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	PolyhedralSurface                 *PolyhedralSurfaceType                 `xml:"http://www.opengis.net/gml PolyhedralSurface,omitempty"`
	PrimeMeridian                     *PrimeMeridianType                     `xml:"http://www.opengis.net/gml PrimeMeridian,omitempty"`
	ProjectedCRS                      *ProjectedCRSType                      `xml:"http://www.opengis.net/gml ProjectedCRS,omitempty"`
	RectifiedGrid                     *RectifiedGridType                     `xml:"http://www.opengis.net/gml RectifiedGrid,omitempty"`
	RectifiedGridCoverage             *RectifiedGridCoverageType             `xml:"http://www.opengis.net/gml RectifiedGridCoverage,omitempty"`
	Ring                              *RingType                              `xml:"http://www.opengis.net/gml Ring,omitempty"`
	Solid                             *SolidType                             `xml:"http://www.opengis.net/gml Solid,omitempty"`
	SphericalCS                       *SphericalCSType                       `xml:"http://www.opengis.net/gml SphericalCS,omitempty"`
	Style                             *StyleType                             `xml:"http://www.opengis.net/gml Style,omitempty"`
	Surface                           *SurfaceType                           `xml:"http://www.opengis.net/gml Surface,omitempty"`
	TemporalCRS                       *TemporalCRSType                       `xml:"http://www.opengis.net/gml TemporalCRS,omitempty"`
	TemporalCS                        *TemporalCSType                        `xml:"http://www.opengis.net/gml TemporalCS,omitempty"`
	TemporalDatum                     *TemporalDatumType                     `xml:"http://www.opengis.net/gml TemporalDatum,omitempty"`
	TimeCalendar                      *TimeCalendarType                      `xml:"http://www.opengis.net/gml TimeCalendar,omitempty"`
	TimeCalendarEra                   *TimeCalendarEraType                   `xml:"http://www.opengis.net/gml TimeCalendarEra,omitempty"`
	TimeClock                         *TimeClockType                         `xml:"http://www.opengis.net/gml TimeClock,omitempty"`
	TimeCoordinateSystem              *TimeCoordinateSystemType              `xml:"http://www.opengis.net/gml TimeCoordinateSystem,omitempty"`
	TimeEdge                          *TimeEdgeType                          `xml:"http://www.opengis.net/gml TimeEdge,omitempty"`
	TimeInstant                       *TimeInstantType                       `xml:"http://www.opengis.net/gml TimeInstant,omitempty"`
	TimeNode                          *TimeNodeType                          `xml:"http://www.opengis.net/gml TimeNode,omitempty"`
	TimeOrdinalReferenceSystem        *TimeOrdinalReferenceSystemType        `xml:"http://www.opengis.net/gml TimeOrdinalReferenceSystem,omitempty"`
	TimePeriod                        *TimePeriodType                        `xml:"http://www.opengis.net/gml TimePeriod,omitempty"`
	TimeTopologyComplex               *TimeTopologyComplexType               `xml:"http://www.opengis.net/gml TimeTopologyComplex,omitempty"`
	Tin                               *TinType                               `xml:"http://www.opengis.net/gml Tin,omitempty"`
	TopoComplex                       *TopoComplexType                       `xml:"http://www.opengis.net/gml TopoComplex,omitempty"`
	TopoSolid                         *TopoSolidType                         `xml:"http://www.opengis.net/gml TopoSolid,omitempty"`
	TopologyStyle                     *TopologyStyleType                     `xml:"http://www.opengis.net/gml TopologyStyle,omitempty"`
	Transformation                    *TransformationType                    `xml:"http://www.opengis.net/gml Transformation,omitempty"`
	TriangulatedSurface               *TriangulatedSurfaceType               `xml:"http://www.opengis.net/gml TriangulatedSurface,omitempty"`
	UnitDefinition                    *UnitDefinitionType                    `xml:"http://www.opengis.net/gml UnitDefinition,omitempty"`
	UserDefinedCS                     *UserDefinedCSType                     `xml:"http://www.opengis.net/gml UserDefinedCS,omitempty"`
	VerticalCRS                       *VerticalCRSType                       `xml:"http://www.opengis.net/gml VerticalCRS,omitempty"`
	VerticalCS                        *VerticalCSType                        `xml:"http://www.opengis.net/gml VerticalCS,omitempty"`
	VerticalDatum                     *VerticalDatumType                     `xml:"http://www.opengis.net/gml VerticalDatum,omitempty"`
	AbstractCRS                       *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _CRS,omitempty"`
	AbstractContinuousCoverage        *AbstractContinuousCoverageType        `xml:"http://www.opengis.net/gml _ContinuousCoverage,omitempty"`
	AbstractCoordinateOperation       *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _CoordinateOperation,omitempty"`
	AbstractCoordinateReferenceSystem *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _CoordinateReferenceSystem,omitempty"`
	AbstractCoordinateSystem          *AbstractCoordinateSystemType          `xml:"http://www.opengis.net/gml _CoordinateSystem,omitempty"`
	AbstractCoverage                  *AbstractCoverageType                  `xml:"http://www.opengis.net/gml _Coverage,omitempty"`
	AbstractCurve                     *AbstractCurveType                     `xml:"http://www.opengis.net/gml _Curve,omitempty"`
	AbstractDatum                     *AbstractDatumType                     `xml:"http://www.opengis.net/gml _Datum,omitempty"`
	AbstractDiscreteCoverage          *AbstractDiscreteCoverageType          `xml:"http://www.opengis.net/gml _DiscreteCoverage,omitempty"`
	AbstractFeature                   *AbstractFeatureType                   `xml:"http://www.opengis.net/gml _Feature,omitempty"`
	AbstractFeatureCollection         *AbstractFeatureCollectionType         `xml:"http://www.opengis.net/gml _FeatureCollection,omitempty"`
	AbstractGML                       *AbstractGMLType                       `xml:"http://www.opengis.net/gml _GML,omitempty"`
	AbstractGeneralConversion         *AbstractGeneralConversionType         `xml:"http://www.opengis.net/gml _GeneralConversion,omitempty"`
	AbstractGeneralDerivedCRS         *AbstractGeneralDerivedCRSType         `xml:"http://www.opengis.net/gml _GeneralDerivedCRS,omitempty"`
	AbstractGeneralOperationParameter *AbstractGeneralOperationParameterType `xml:"http://www.opengis.net/gml _GeneralOperationParameter,omitempty"`
	AbstractGeneralTransformation     *AbstractGeneralTransformationType     `xml:"http://www.opengis.net/gml _GeneralTransformation,omitempty"`
	AbstractGeometricAggregate        *AbstractGeometricAggregateType        `xml:"http://www.opengis.net/gml _GeometricAggregate,omitempty"`
	AbstractGeometricPrimitive        *AbstractGeometricPrimitiveType        `xml:"http://www.opengis.net/gml _GeometricPrimitive,omitempty"`
	AbstractGeometry                  *AbstractGeometryType                  `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	AbstractImplicitGeometry          *AbstractGeometryType                  `xml:"http://www.opengis.net/gml _ImplicitGeometry,omitempty"`
	AbstractMetaData                  *AbstractMetaDataType                  `xml:"http://www.opengis.net/gml _MetaData,omitempty"`
	AbstractOperation                 *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _Operation,omitempty"`
	AbstractReferenceSystem           *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _ReferenceSystem,omitempty"`
	AbstractRing                      *AbstractRingType                      `xml:"http://www.opengis.net/gml _Ring,omitempty"`
	AbstractSingleOperation           *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _SingleOperation,omitempty"`
	AbstractSolid                     *AbstractSolidType                     `xml:"http://www.opengis.net/gml _Solid,omitempty"`
	AbstractStyle                     *AbstractStyleType                     `xml:"http://www.opengis.net/gml _Style,omitempty"`
	AbstractSurface                   *AbstractSurfaceType                   `xml:"http://www.opengis.net/gml _Surface,omitempty"`
	AbstractTimeComplex               *AbstractTimeComplexType               `xml:"http://www.opengis.net/gml _TimeComplex,omitempty"`
	AbstractTimeGeometricPrimitive    *AbstractTimeGeometricPrimitiveType    `xml:"http://www.opengis.net/gml _TimeGeometricPrimitive,omitempty"`
	AbstractTimeObject                *AbstractTimeObjectType                `xml:"http://www.opengis.net/gml _TimeObject,omitempty"`
	AbstractTimePrimitive             *AbstractTimePrimitiveType             `xml:"http://www.opengis.net/gml _TimePrimitive,omitempty"`
	AbstractTimeReferenceSystem       *AbstractTimeReferenceSystemType       `xml:"http://www.opengis.net/gml _TimeReferenceSystem,omitempty"`
	AbstractTimeSlice                 *AbstractTimeSliceType                 `xml:"http://www.opengis.net/gml _TimeSlice,omitempty"`
	AbstractTimeTopologyPrimitive     *AbstractTimeTopologyPrimitiveType     `xml:"http://www.opengis.net/gml _TimeTopologyPrimitive,omitempty"`
	AbstractTopoPrimitive             *AbstractTopoPrimitiveType             `xml:"http://www.opengis.net/gml _TopoPrimitive,omitempty"`
	AbstractTopology                  *AbstractTopologyType                  `xml:"http://www.opengis.net/gml _Topology,omitempty"`
	Null                              *string                                `xml:"http://www.opengis.net/gml Null,omitempty"`
	// Aggregate value built using the Composite pattern.
	CompositeValue *CompositeValueType `xml:"http://www.opengis.net/gml CompositeValue,omitempty"`
	ValueArray     *ValueArrayType     `xml:"http://www.opengis.net/gml ValueArray,omitempty"`
	// A value from two-valued logic, using the XML Schema boolean type.  An instance may take the values {true, false, 1, 0}.
	Boolean *bool `xml:"http://www.opengis.net/gml Boolean,omitempty"`
	// A term representing a classification.  It has an optional XML attribute codeSpace, whose value is a URI which identifies a dictionary, codelist or authority for the term.
	Category *CodeType `xml:"http://www.opengis.net/gml Category,omitempty"`
	// A numeric value with a scale.  The content of the element is an amount using the XML Schema type double which permits decimal or scientific notation.  An XML attribute uom (unit of measure) is required, whose value is a URI which identifies the definition of the scale or units by which the numeric value must be multiplied.
	Quantity *MeasureType `xml:"http://www.opengis.net/gml Quantity,omitempty"`
	// An integer representing a frequency of occurrence.
	Count *int `xml:"http://www.opengis.net/gml Count,omitempty"`
	// XML List based on XML Schema boolean type.  An element of this type contains a space-separated list of boolean values {0,1,true,false}
	BooleanList *string `xml:"http://www.opengis.net/gml BooleanList,omitempty"`
	// A space-separated list of terms or nulls.  A single XML attribute codeSpace may be provided, which authorises all the terms in the list.
	CategoryList *CodeOrNullListType `xml:"http://www.opengis.net/gml CategoryList,omitempty"`
	// A space separated list of amounts or nulls.  The amounts use the XML Schema type double.  A single XML attribute uom (unit of measure) is required, whose value is a URI which identifies the definition of the scale or units by which all the amounts in the list must be multiplied.
	QuantityList *MeasureOrNullListType `xml:"http://www.opengis.net/gml QuantityList,omitempty"`
	// A space-separated list of integers or nulls.
	CountList *string `xml:"http://www.opengis.net/gml CountList,omitempty"`
	// Utility element to store a 2-point range of ordinal values. If one member is a null, then this is a single ended interval.
	CategoryExtent *CategoryExtentType `xml:"http://www.opengis.net/gml CategoryExtent,omitempty"`
	// Utility element to store a 2-point range of numeric values. If one member is a null, then this is a single ended interval.
	QuantityExtent *QuantityExtentType `xml:"http://www.opengis.net/gml QuantityExtent,omitempty"`
	// Utility element to store a 2-point range of frequency values. If one member is a null, then this is a single ended interval.
	CountExtent  *string `xml:"http://www.opengis.net/gml CountExtent,omitempty"`
	RemoteSchema string  `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string  `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string  `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string  `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string  `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string  `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string  `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string  `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type RangeParametersType struct {
	// Aggregate value built using the Composite pattern.
	CompositeValue *CompositeValueType `xml:"http://www.opengis.net/gml CompositeValue,omitempty"`
	ValueArray     *ValueArrayType     `xml:"http://www.opengis.net/gml ValueArray,omitempty"`
	// A value from two-valued logic, using the XML Schema boolean type.  An instance may take the values {true, false, 1, 0}.
	Boolean *bool `xml:"http://www.opengis.net/gml Boolean,omitempty"`
	// A term representing a classification.  It has an optional XML attribute codeSpace, whose value is a URI which identifies a dictionary, codelist or authority for the term.
	Category *CodeType `xml:"http://www.opengis.net/gml Category,omitempty"`
	// A numeric value with a scale.  The content of the element is an amount using the XML Schema type double which permits decimal or scientific notation.  An XML attribute uom (unit of measure) is required, whose value is a URI which identifies the definition of the scale or units by which the numeric value must be multiplied.
	Quantity *MeasureType `xml:"http://www.opengis.net/gml Quantity,omitempty"`
	// An integer representing a frequency of occurrence.
	Count *int `xml:"http://www.opengis.net/gml Count,omitempty"`
	// XML List based on XML Schema boolean type.  An element of this type contains a space-separated list of boolean values {0,1,true,false}
	BooleanList *string `xml:"http://www.opengis.net/gml BooleanList,omitempty"`
	// A space-separated list of terms or nulls.  A single XML attribute codeSpace may be provided, which authorises all the terms in the list.
	CategoryList *CodeOrNullListType `xml:"http://www.opengis.net/gml CategoryList,omitempty"`
	// A space separated list of amounts or nulls.  The amounts use the XML Schema type double.  A single XML attribute uom (unit of measure) is required, whose value is a URI which identifies the definition of the scale or units by which all the amounts in the list must be multiplied.
	QuantityList *MeasureOrNullListType `xml:"http://www.opengis.net/gml QuantityList,omitempty"`
	// A space-separated list of integers or nulls.
	CountList *string `xml:"http://www.opengis.net/gml CountList,omitempty"`
	// Utility element to store a 2-point range of ordinal values. If one member is a null, then this is a single ended interval.
	CategoryExtent *CategoryExtentType `xml:"http://www.opengis.net/gml CategoryExtent,omitempty"`
	// Utility element to store a 2-point range of numeric values. If one member is a null, then this is a single ended interval.
	QuantityExtent *QuantityExtentType `xml:"http://www.opengis.net/gml QuantityExtent,omitempty"`
	// Utility element to store a 2-point range of frequency values. If one member is a null, then this is a single ended interval.
	CountExtent  *string `xml:"http://www.opengis.net/gml CountExtent,omitempty"`
	RemoteSchema string  `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string  `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string  `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string  `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string  `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string  `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string  `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string  `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type RangeSetType struct {
	// each member _Value holds a tuple or "row" from the equivalent table
	ValueArray []ValueArrayType `xml:"http://www.opengis.net/gml ValueArray"`
	// Its tuple list holds the values as space-separated tuples each of which contains comma-separated components, and the tuple structure is specified using the rangeParameters property.
	DataBlock *DataBlockType `xml:"http://www.opengis.net/gml DataBlock,omitempty"`
	// a reference to an external source for the data, together with a description of how that external source is structured
	File *FileType `xml:"http://www.opengis.net/gml File,omitempty"`
	// XML List based on XML Schema boolean type.  An element of this type contains a space-separated list of boolean values {0,1,true,false}
	BooleanList *string `xml:"http://www.opengis.net/gml BooleanList,omitempty"`
	// A space-separated list of terms or nulls.  A single XML attribute codeSpace may be provided, which authorises all the terms in the list.
	CategoryList *CodeOrNullListType `xml:"http://www.opengis.net/gml CategoryList,omitempty"`
	// A space separated list of amounts or nulls.  The amounts use the XML Schema type double.  A single XML attribute uom (unit of measure) is required, whose value is a URI which identifies the definition of the scale or units by which all the amounts in the list must be multiplied.
	QuantityList *MeasureOrNullListType `xml:"http://www.opengis.net/gml QuantityList,omitempty"`
	// A space-separated list of integers or nulls.
	CountList *string `xml:"http://www.opengis.net/gml CountList,omitempty"`
}

type RectangleType struct {
	// Constraint: The Ring shall be a LinearRing and must form a rectangle; the first and the last position must be co-incident.
	Exterior        *AbstractRingPropertyType `xml:"http://www.opengis.net/gml exterior,omitempty"`
	OuterBoundaryIs *AbstractRingPropertyType `xml:"http://www.opengis.net/gml outerBoundaryIs,omitempty"`
	// The attribute "interpolation" specifies the interpolation mechanism used for this surface patch. Currently only planar surface patches are defined in GML 3, the attribute is fixed to "planar", i.e. the interpolation method shall return points on a single plane. The boundary of the patch shall be contained within that plane.
	Interpolation *string `xml:"interpolation,attr,omitempty"`
}

type RectifiedGridCoverageType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType         `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType         `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType         `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType         `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType         `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType         `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType         `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType         `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType         `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType         `xml:"http://www.opengis.net/gml srsName"`
	Id                      string             `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	BoundedBy               *BoundingShapeType `xml:"http://www.opengis.net/gml boundedBy,omitempty"`
	// deprecated in GML version 3.1
	Location            *LocationPropertyType         `xml:"http://www.opengis.net/gml location,omitempty"`
	PriorityLocation    *PriorityLocationPropertyType `xml:"http://www.opengis.net/gml priorityLocation,omitempty"`
	DomainSet           *DomainSetType                `xml:"http://www.opengis.net/gml domainSet,omitempty"`
	GridDomain          *GridDomainType               `xml:"http://www.opengis.net/gml gridDomain,omitempty"`
	MultiCurveDomain    *MultiCurveDomainType         `xml:"http://www.opengis.net/gml multiCurveDomain,omitempty"`
	MultiPointDomain    *MultiPointDomainType         `xml:"http://www.opengis.net/gml multiPointDomain,omitempty"`
	MultiSolidDomain    *MultiSolidDomainType         `xml:"http://www.opengis.net/gml multiSolidDomain,omitempty"`
	MultiSurfaceDomain  *MultiSurfaceDomainType       `xml:"http://www.opengis.net/gml multiSurfaceDomain,omitempty"`
	RectifiedGridDomain *RectifiedGridDomainType      `xml:"http://www.opengis.net/gml rectifiedGridDomain,omitempty"`
	RangeSet            *RangeSetType                 `xml:"http://www.opengis.net/gml rangeSet,omitempty"`
	Dimension           *int                          `xml:"dimension,attr,omitempty"`
	CoverageFunction    *CoverageFunctionType         `xml:"http://www.opengis.net/gml coverageFunction,omitempty"`
}

type RectifiedGridDomainType struct {
	// The "_Geometry" element is the abstract head of the substituition group for all geometry elements of GML 3. This
	// includes pre-defined and user-defined geometry elements. Any geometry element must be a direct or indirect extension/restriction
	// of AbstractGeometryType and must be directly or indirectly in the substitution group of "_Geometry".
	AbstractGeometry           *AbstractGeometryType           `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	CompositeCurve             *CompositeCurveType             `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	CompositeSolid             *CompositeSolidType             `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	CompositeSurface           *CompositeSurfaceType           `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	Curve                      *CurveType                      `xml:"http://www.opengis.net/gml Curve,omitempty"`
	GeometricComplex           *GeometricComplexType           `xml:"http://www.opengis.net/gml GeometricComplex,omitempty"`
	Grid                       *GridType                       `xml:"http://www.opengis.net/gml Grid,omitempty"`
	LineString                 *LineStringType                 `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearRing                 *LinearRingType                 `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MultiCurve                 *MultiCurveType                 `xml:"http://www.opengis.net/gml MultiCurve,omitempty"`
	MultiGeometry              *MultiGeometryType              `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString            *MultiLineStringType            `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                 *MultiPointType                 `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPolygon               *MultiPolygonType               `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	MultiSolid                 *MultiSolidType                 `xml:"http://www.opengis.net/gml MultiSolid,omitempty"`
	MultiSurface               *MultiSurfaceType               `xml:"http://www.opengis.net/gml MultiSurface,omitempty"`
	OrientableCurve            *OrientableCurveType            `xml:"http://www.opengis.net/gml OrientableCurve,omitempty"`
	OrientableSurface          *OrientableSurfaceType          `xml:"http://www.opengis.net/gml OrientableSurface,omitempty"`
	Point                      *PointType                      `xml:"http://www.opengis.net/gml Point,omitempty"`
	Polygon                    *PolygonType                    `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	PolyhedralSurface          *PolyhedralSurfaceType          `xml:"http://www.opengis.net/gml PolyhedralSurface,omitempty"`
	RectifiedGrid              *RectifiedGridType              `xml:"http://www.opengis.net/gml RectifiedGrid,omitempty"`
	Ring                       *RingType                       `xml:"http://www.opengis.net/gml Ring,omitempty"`
	Solid                      *SolidType                      `xml:"http://www.opengis.net/gml Solid,omitempty"`
	Surface                    *SurfaceType                    `xml:"http://www.opengis.net/gml Surface,omitempty"`
	Tin                        *TinType                        `xml:"http://www.opengis.net/gml Tin,omitempty"`
	TriangulatedSurface        *TriangulatedSurfaceType        `xml:"http://www.opengis.net/gml TriangulatedSurface,omitempty"`
	AbstractCurve              *AbstractCurveType              `xml:"http://www.opengis.net/gml _Curve,omitempty"`
	AbstractGeometricAggregate *AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml _GeometricAggregate,omitempty"`
	AbstractGeometricPrimitive *AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml _GeometricPrimitive,omitempty"`
	AbstractImplicitGeometry   *AbstractGeometryType           `xml:"http://www.opengis.net/gml _ImplicitGeometry,omitempty"`
	AbstractRing               *AbstractRingType               `xml:"http://www.opengis.net/gml _Ring,omitempty"`
	AbstractSolid              *AbstractSolidType              `xml:"http://www.opengis.net/gml _Solid,omitempty"`
	AbstractSurface            *AbstractSurfaceType            `xml:"http://www.opengis.net/gml _Surface,omitempty"`
	// This abstract element acts as the head of the substitution group for temporal primitives and complexes.
	AbstractTimeObject             *AbstractTimeObjectType             `xml:"http://www.opengis.net/gml _TimeObject,omitempty"`
	TimeEdge                       *TimeEdgeType                       `xml:"http://www.opengis.net/gml TimeEdge,omitempty"`
	TimeInstant                    *TimeInstantType                    `xml:"http://www.opengis.net/gml TimeInstant,omitempty"`
	TimeNode                       *TimeNodeType                       `xml:"http://www.opengis.net/gml TimeNode,omitempty"`
	TimePeriod                     *TimePeriodType                     `xml:"http://www.opengis.net/gml TimePeriod,omitempty"`
	TimeTopologyComplex            *TimeTopologyComplexType            `xml:"http://www.opengis.net/gml TimeTopologyComplex,omitempty"`
	AbstractTimeComplex            *AbstractTimeComplexType            `xml:"http://www.opengis.net/gml _TimeComplex,omitempty"`
	AbstractTimeGeometricPrimitive *AbstractTimeGeometricPrimitiveType `xml:"http://www.opengis.net/gml _TimeGeometricPrimitive,omitempty"`
	AbstractTimePrimitive          *AbstractTimePrimitiveType          `xml:"http://www.opengis.net/gml _TimePrimitive,omitempty"`
	AbstractTimeTopologyPrimitive  *AbstractTimeTopologyPrimitiveType  `xml:"http://www.opengis.net/gml _TimeTopologyPrimitive,omitempty"`
	RemoteSchema                   string                              `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                      string                              `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                           string                              `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                           string                              `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                        string                              `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                          string                              `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                           string                              `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                        string                              `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type RectifiedGridType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels    *string            `xml:"uomLabels,attr,omitempty"`
	Limits       *GridLimitsType    `xml:"http://www.opengis.net/gml limits,omitempty"`
	AxisName     []string           `xml:"http://www.opengis.net/gml axisName"`
	Dimension    int                `xml:"dimension,attr"`
	Origin       *PointPropertyType `xml:"http://www.opengis.net/gml origin,omitempty"`
	OffsetVector []VectorType       `xml:"http://www.opengis.net/gml offsetVector"`
}

type ReferenceSystemRefType struct {
	AbstractReferenceSystem           *AbstractReferenceSystemType   `xml:"http://www.opengis.net/gml _ReferenceSystem,omitempty"`
	CompoundCRS                       *CompoundCRSType               `xml:"http://www.opengis.net/gml CompoundCRS,omitempty"`
	DerivedCRS                        *DerivedCRSType                `xml:"http://www.opengis.net/gml DerivedCRS,omitempty"`
	EngineeringCRS                    *EngineeringCRSType            `xml:"http://www.opengis.net/gml EngineeringCRS,omitempty"`
	GeocentricCRS                     *GeocentricCRSType             `xml:"http://www.opengis.net/gml GeocentricCRS,omitempty"`
	GeographicCRS                     *GeographicCRSType             `xml:"http://www.opengis.net/gml GeographicCRS,omitempty"`
	ImageCRS                          *ImageCRSType                  `xml:"http://www.opengis.net/gml ImageCRS,omitempty"`
	ProjectedCRS                      *ProjectedCRSType              `xml:"http://www.opengis.net/gml ProjectedCRS,omitempty"`
	TemporalCRS                       *TemporalCRSType               `xml:"http://www.opengis.net/gml TemporalCRS,omitempty"`
	VerticalCRS                       *VerticalCRSType               `xml:"http://www.opengis.net/gml VerticalCRS,omitempty"`
	AbstractCRS                       *AbstractReferenceSystemType   `xml:"http://www.opengis.net/gml _CRS,omitempty"`
	AbstractCoordinateReferenceSystem *AbstractReferenceSystemType   `xml:"http://www.opengis.net/gml _CoordinateReferenceSystem,omitempty"`
	AbstractGeneralDerivedCRS         *AbstractGeneralDerivedCRSType `xml:"http://www.opengis.net/gml _GeneralDerivedCRS,omitempty"`
	RemoteSchema                      string                         `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                         string                         `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                              string                         `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                              string                         `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                           string                         `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                             string                         `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                              string                         `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                           string                         `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type ReferenceType struct {
	RemoteSchema string `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type RelatedTimeType struct {
	// This abstract element acts as the head of the substitution group for temporal primitives.
	AbstractTimePrimitive          *AbstractTimePrimitiveType          `xml:"http://www.opengis.net/gml _TimePrimitive,omitempty"`
	TimeEdge                       *TimeEdgeType                       `xml:"http://www.opengis.net/gml TimeEdge,omitempty"`
	TimeInstant                    *TimeInstantType                    `xml:"http://www.opengis.net/gml TimeInstant,omitempty"`
	TimeNode                       *TimeNodeType                       `xml:"http://www.opengis.net/gml TimeNode,omitempty"`
	TimePeriod                     *TimePeriodType                     `xml:"http://www.opengis.net/gml TimePeriod,omitempty"`
	AbstractTimeGeometricPrimitive *AbstractTimeGeometricPrimitiveType `xml:"http://www.opengis.net/gml _TimeGeometricPrimitive,omitempty"`
	AbstractTimeTopologyPrimitive  *AbstractTimeTopologyPrimitiveType  `xml:"http://www.opengis.net/gml _TimeTopologyPrimitive,omitempty"`
	RemoteSchema                   string                              `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                      string                              `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                           string                              `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                           string                              `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                        string                              `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                          string                              `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                           string                              `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                        string                              `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
	RelativePosition               *string                             `xml:"relativePosition,attr,omitempty"`
}

type RelativeInternalPositionalAccuracyType struct {
	// A description of the position accuracy parameter(s) provided.
	MeasureDescription *CodeType `xml:"http://www.opengis.net/gml measureDescription,omitempty"`
	// A quantitative result defined by the evaluation procedure used, and identified by the measureDescription.
	Result *MeasureType `xml:"http://www.opengis.net/gml result,omitempty"`
}

type RingPropertyType struct {
	Ring *RingType `xml:"http://www.opengis.net/gml Ring,omitempty"`
}

type RingType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
	// This element references or contains one curve in the composite curve. The curves are contiguous, the collection of curves is ordered.
	// NOTE: This definition allows for a nested structure, i.e. a CompositeCurve may use, for example, another CompositeCurve as a curve member.
	CurveMember []CurvePropertyType `xml:"http://www.opengis.net/gml curveMember"`
}

type ScalarValuePropertyType struct {
	// This abstract element is the head of a substitutionGroup hierararchy which may contain either simpleContent or complexContent elements.  It is used to assert the model position of "class" elements declared in other GML schemas.
	AbstractObject                    *string                                `xml:"http://www.opengis.net/gml _Object,omitempty"`
	Array                             *ArrayType                             `xml:"http://www.opengis.net/gml Array,omitempty"`
	Bag                               *BagType                               `xml:"http://www.opengis.net/gml Bag,omitempty"`
	BaseUnit                          *BaseUnitType                          `xml:"http://www.opengis.net/gml BaseUnit,omitempty"`
	CartesianCS                       *CartesianCSType                       `xml:"http://www.opengis.net/gml CartesianCS,omitempty"`
	CompositeCurve                    *CompositeCurveType                    `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	CompositeSolid                    *CompositeSolidType                    `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	CompositeSurface                  *CompositeSurfaceType                  `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	CompoundCRS                       *CompoundCRSType                       `xml:"http://www.opengis.net/gml CompoundCRS,omitempty"`
	ConcatenatedOperation             *ConcatenatedOperationType             `xml:"http://www.opengis.net/gml ConcatenatedOperation,omitempty"`
	ConventionalUnit                  *ConventionalUnitType                  `xml:"http://www.opengis.net/gml ConventionalUnit,omitempty"`
	Conversion                        *ConversionType                        `xml:"http://www.opengis.net/gml Conversion,omitempty"`
	CoordinateSystemAxis              *CoordinateSystemAxisType              `xml:"http://www.opengis.net/gml CoordinateSystemAxis,omitempty"`
	Curve                             *CurveType                             `xml:"http://www.opengis.net/gml Curve,omitempty"`
	CylindricalCS                     *CylindricalCSType                     `xml:"http://www.opengis.net/gml CylindricalCS,omitempty"`
	Definition                        *DefinitionType                        `xml:"http://www.opengis.net/gml Definition,omitempty"`
	DefinitionCollection              *DictionaryType                        `xml:"http://www.opengis.net/gml DefinitionCollection,omitempty"`
	DefinitionProxy                   *DefinitionProxyType                   `xml:"http://www.opengis.net/gml DefinitionProxy,omitempty"`
	DerivedCRS                        *DerivedCRSType                        `xml:"http://www.opengis.net/gml DerivedCRS,omitempty"`
	DerivedUnit                       *DerivedUnitType                       `xml:"http://www.opengis.net/gml DerivedUnit,omitempty"`
	Dictionary                        *DictionaryType                        `xml:"http://www.opengis.net/gml Dictionary,omitempty"`
	DirectedObservation               *DirectedObservationType               `xml:"http://www.opengis.net/gml DirectedObservation,omitempty"`
	DirectedObservationAtDistance     *DirectedObservationAtDistanceType     `xml:"http://www.opengis.net/gml DirectedObservationAtDistance,omitempty"`
	Edge                              *EdgeType                              `xml:"http://www.opengis.net/gml Edge,omitempty"`
	Ellipsoid                         *EllipsoidType                         `xml:"http://www.opengis.net/gml Ellipsoid,omitempty"`
	EllipsoidalCS                     *EllipsoidalCSType                     `xml:"http://www.opengis.net/gml EllipsoidalCS,omitempty"`
	EngineeringCRS                    *EngineeringCRSType                    `xml:"http://www.opengis.net/gml EngineeringCRS,omitempty"`
	EngineeringDatum                  *EngineeringDatumType                  `xml:"http://www.opengis.net/gml EngineeringDatum,omitempty"`
	Face                              *FaceType                              `xml:"http://www.opengis.net/gml Face,omitempty"`
	FeatureCollection                 *FeatureCollectionType                 `xml:"http://www.opengis.net/gml FeatureCollection,omitempty"`
	FeatureStyle                      *FeatureStyleType                      `xml:"http://www.opengis.net/gml FeatureStyle,omitempty"`
	GenericMetaData                   *GenericMetaDataType                   `xml:"http://www.opengis.net/gml GenericMetaData,omitempty"`
	GeocentricCRS                     *GeocentricCRSType                     `xml:"http://www.opengis.net/gml GeocentricCRS,omitempty"`
	GeodeticDatum                     *GeodeticDatumType                     `xml:"http://www.opengis.net/gml GeodeticDatum,omitempty"`
	GeographicCRS                     *GeographicCRSType                     `xml:"http://www.opengis.net/gml GeographicCRS,omitempty"`
	GeometricComplex                  *GeometricComplexType                  `xml:"http://www.opengis.net/gml GeometricComplex,omitempty"`
	GeometryStyle                     *GeometryStyleType                     `xml:"http://www.opengis.net/gml GeometryStyle,omitempty"`
	GraphStyle                        *GraphStyleType                        `xml:"http://www.opengis.net/gml GraphStyle,omitempty"`
	Grid                              *GridType                              `xml:"http://www.opengis.net/gml Grid,omitempty"`
	GridCoverage                      *GridCoverageType                      `xml:"http://www.opengis.net/gml GridCoverage,omitempty"`
	ImageCRS                          *ImageCRSType                          `xml:"http://www.opengis.net/gml ImageCRS,omitempty"`
	ImageDatum                        *ImageDatumType                        `xml:"http://www.opengis.net/gml ImageDatum,omitempty"`
	LabelStyle                        *LabelStyleType                        `xml:"http://www.opengis.net/gml LabelStyle,omitempty"`
	LineString                        *LineStringType                        `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearCS                          *LinearCSType                          `xml:"http://www.opengis.net/gml LinearCS,omitempty"`
	LinearRing                        *LinearRingType                        `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MovingObjectStatus                *MovingObjectStatusType                `xml:"http://www.opengis.net/gml MovingObjectStatus,omitempty"`
	MultiCurve                        *MultiCurveType                        `xml:"http://www.opengis.net/gml MultiCurve,omitempty"`
	MultiCurveCoverage                *MultiCurveCoverageType                `xml:"http://www.opengis.net/gml MultiCurveCoverage,omitempty"`
	MultiGeometry                     *MultiGeometryType                     `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString                   *MultiLineStringType                   `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                        *MultiPointType                        `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPointCoverage                *MultiPointCoverageType                `xml:"http://www.opengis.net/gml MultiPointCoverage,omitempty"`
	MultiPolygon                      *MultiPolygonType                      `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	MultiSolid                        *MultiSolidType                        `xml:"http://www.opengis.net/gml MultiSolid,omitempty"`
	MultiSolidCoverage                *MultiSolidCoverageType                `xml:"http://www.opengis.net/gml MultiSolidCoverage,omitempty"`
	MultiSurface                      *MultiSurfaceType                      `xml:"http://www.opengis.net/gml MultiSurface,omitempty"`
	MultiSurfaceCoverage              *MultiSurfaceCoverageType              `xml:"http://www.opengis.net/gml MultiSurfaceCoverage,omitempty"`
	Node                              *NodeType                              `xml:"http://www.opengis.net/gml Node,omitempty"`
	ObliqueCartesianCS                *ObliqueCartesianCSType                `xml:"http://www.opengis.net/gml ObliqueCartesianCS,omitempty"`
	Observation                       *ObservationType                       `xml:"http://www.opengis.net/gml Observation,omitempty"`
	OperationMethod                   *OperationMethodType                   `xml:"http://www.opengis.net/gml OperationMethod,omitempty"`
	OperationParameter                *OperationParameterType                `xml:"http://www.opengis.net/gml OperationParameter,omitempty"`
	OperationParameterGroup           *OperationParameterGroupType           `xml:"http://www.opengis.net/gml OperationParameterGroup,omitempty"`
	OrientableCurve                   *OrientableCurveType                   `xml:"http://www.opengis.net/gml OrientableCurve,omitempty"`
	OrientableSurface                 *OrientableSurfaceType                 `xml:"http://www.opengis.net/gml OrientableSurface,omitempty"`
	PassThroughOperation              *PassThroughOperationType              `xml:"http://www.opengis.net/gml PassThroughOperation,omitempty"`
	Point                             *PointType                             `xml:"http://www.opengis.net/gml Point,omitempty"`
	PolarCS                           *PolarCSType                           `xml:"http://www.opengis.net/gml PolarCS,omitempty"`
	Polygon                           *PolygonType                           `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	PolyhedralSurface                 *PolyhedralSurfaceType                 `xml:"http://www.opengis.net/gml PolyhedralSurface,omitempty"`
	PrimeMeridian                     *PrimeMeridianType                     `xml:"http://www.opengis.net/gml PrimeMeridian,omitempty"`
	ProjectedCRS                      *ProjectedCRSType                      `xml:"http://www.opengis.net/gml ProjectedCRS,omitempty"`
	RectifiedGrid                     *RectifiedGridType                     `xml:"http://www.opengis.net/gml RectifiedGrid,omitempty"`
	RectifiedGridCoverage             *RectifiedGridCoverageType             `xml:"http://www.opengis.net/gml RectifiedGridCoverage,omitempty"`
	Ring                              *RingType                              `xml:"http://www.opengis.net/gml Ring,omitempty"`
	Solid                             *SolidType                             `xml:"http://www.opengis.net/gml Solid,omitempty"`
	SphericalCS                       *SphericalCSType                       `xml:"http://www.opengis.net/gml SphericalCS,omitempty"`
	Style                             *StyleType                             `xml:"http://www.opengis.net/gml Style,omitempty"`
	Surface                           *SurfaceType                           `xml:"http://www.opengis.net/gml Surface,omitempty"`
	TemporalCRS                       *TemporalCRSType                       `xml:"http://www.opengis.net/gml TemporalCRS,omitempty"`
	TemporalCS                        *TemporalCSType                        `xml:"http://www.opengis.net/gml TemporalCS,omitempty"`
	TemporalDatum                     *TemporalDatumType                     `xml:"http://www.opengis.net/gml TemporalDatum,omitempty"`
	TimeCalendar                      *TimeCalendarType                      `xml:"http://www.opengis.net/gml TimeCalendar,omitempty"`
	TimeCalendarEra                   *TimeCalendarEraType                   `xml:"http://www.opengis.net/gml TimeCalendarEra,omitempty"`
	TimeClock                         *TimeClockType                         `xml:"http://www.opengis.net/gml TimeClock,omitempty"`
	TimeCoordinateSystem              *TimeCoordinateSystemType              `xml:"http://www.opengis.net/gml TimeCoordinateSystem,omitempty"`
	TimeEdge                          *TimeEdgeType                          `xml:"http://www.opengis.net/gml TimeEdge,omitempty"`
	TimeInstant                       *TimeInstantType                       `xml:"http://www.opengis.net/gml TimeInstant,omitempty"`
	TimeNode                          *TimeNodeType                          `xml:"http://www.opengis.net/gml TimeNode,omitempty"`
	TimeOrdinalReferenceSystem        *TimeOrdinalReferenceSystemType        `xml:"http://www.opengis.net/gml TimeOrdinalReferenceSystem,omitempty"`
	TimePeriod                        *TimePeriodType                        `xml:"http://www.opengis.net/gml TimePeriod,omitempty"`
	TimeTopologyComplex               *TimeTopologyComplexType               `xml:"http://www.opengis.net/gml TimeTopologyComplex,omitempty"`
	Tin                               *TinType                               `xml:"http://www.opengis.net/gml Tin,omitempty"`
	TopoComplex                       *TopoComplexType                       `xml:"http://www.opengis.net/gml TopoComplex,omitempty"`
	TopoSolid                         *TopoSolidType                         `xml:"http://www.opengis.net/gml TopoSolid,omitempty"`
	TopologyStyle                     *TopologyStyleType                     `xml:"http://www.opengis.net/gml TopologyStyle,omitempty"`
	Transformation                    *TransformationType                    `xml:"http://www.opengis.net/gml Transformation,omitempty"`
	TriangulatedSurface               *TriangulatedSurfaceType               `xml:"http://www.opengis.net/gml TriangulatedSurface,omitempty"`
	UnitDefinition                    *UnitDefinitionType                    `xml:"http://www.opengis.net/gml UnitDefinition,omitempty"`
	UserDefinedCS                     *UserDefinedCSType                     `xml:"http://www.opengis.net/gml UserDefinedCS,omitempty"`
	VerticalCRS                       *VerticalCRSType                       `xml:"http://www.opengis.net/gml VerticalCRS,omitempty"`
	VerticalCS                        *VerticalCSType                        `xml:"http://www.opengis.net/gml VerticalCS,omitempty"`
	VerticalDatum                     *VerticalDatumType                     `xml:"http://www.opengis.net/gml VerticalDatum,omitempty"`
	AbstractCRS                       *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _CRS,omitempty"`
	AbstractContinuousCoverage        *AbstractContinuousCoverageType        `xml:"http://www.opengis.net/gml _ContinuousCoverage,omitempty"`
	AbstractCoordinateOperation       *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _CoordinateOperation,omitempty"`
	AbstractCoordinateReferenceSystem *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _CoordinateReferenceSystem,omitempty"`
	AbstractCoordinateSystem          *AbstractCoordinateSystemType          `xml:"http://www.opengis.net/gml _CoordinateSystem,omitempty"`
	AbstractCoverage                  *AbstractCoverageType                  `xml:"http://www.opengis.net/gml _Coverage,omitempty"`
	AbstractCurve                     *AbstractCurveType                     `xml:"http://www.opengis.net/gml _Curve,omitempty"`
	AbstractDatum                     *AbstractDatumType                     `xml:"http://www.opengis.net/gml _Datum,omitempty"`
	AbstractDiscreteCoverage          *AbstractDiscreteCoverageType          `xml:"http://www.opengis.net/gml _DiscreteCoverage,omitempty"`
	AbstractFeature                   *AbstractFeatureType                   `xml:"http://www.opengis.net/gml _Feature,omitempty"`
	AbstractFeatureCollection         *AbstractFeatureCollectionType         `xml:"http://www.opengis.net/gml _FeatureCollection,omitempty"`
	AbstractGML                       *AbstractGMLType                       `xml:"http://www.opengis.net/gml _GML,omitempty"`
	AbstractGeneralConversion         *AbstractGeneralConversionType         `xml:"http://www.opengis.net/gml _GeneralConversion,omitempty"`
	AbstractGeneralDerivedCRS         *AbstractGeneralDerivedCRSType         `xml:"http://www.opengis.net/gml _GeneralDerivedCRS,omitempty"`
	AbstractGeneralOperationParameter *AbstractGeneralOperationParameterType `xml:"http://www.opengis.net/gml _GeneralOperationParameter,omitempty"`
	AbstractGeneralTransformation     *AbstractGeneralTransformationType     `xml:"http://www.opengis.net/gml _GeneralTransformation,omitempty"`
	AbstractGeometricAggregate        *AbstractGeometricAggregateType        `xml:"http://www.opengis.net/gml _GeometricAggregate,omitempty"`
	AbstractGeometricPrimitive        *AbstractGeometricPrimitiveType        `xml:"http://www.opengis.net/gml _GeometricPrimitive,omitempty"`
	AbstractGeometry                  *AbstractGeometryType                  `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	AbstractImplicitGeometry          *AbstractGeometryType                  `xml:"http://www.opengis.net/gml _ImplicitGeometry,omitempty"`
	AbstractMetaData                  *AbstractMetaDataType                  `xml:"http://www.opengis.net/gml _MetaData,omitempty"`
	AbstractOperation                 *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _Operation,omitempty"`
	AbstractReferenceSystem           *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _ReferenceSystem,omitempty"`
	AbstractRing                      *AbstractRingType                      `xml:"http://www.opengis.net/gml _Ring,omitempty"`
	AbstractSingleOperation           *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _SingleOperation,omitempty"`
	AbstractSolid                     *AbstractSolidType                     `xml:"http://www.opengis.net/gml _Solid,omitempty"`
	AbstractStyle                     *AbstractStyleType                     `xml:"http://www.opengis.net/gml _Style,omitempty"`
	AbstractSurface                   *AbstractSurfaceType                   `xml:"http://www.opengis.net/gml _Surface,omitempty"`
	AbstractTimeComplex               *AbstractTimeComplexType               `xml:"http://www.opengis.net/gml _TimeComplex,omitempty"`
	AbstractTimeGeometricPrimitive    *AbstractTimeGeometricPrimitiveType    `xml:"http://www.opengis.net/gml _TimeGeometricPrimitive,omitempty"`
	AbstractTimeObject                *AbstractTimeObjectType                `xml:"http://www.opengis.net/gml _TimeObject,omitempty"`
	AbstractTimePrimitive             *AbstractTimePrimitiveType             `xml:"http://www.opengis.net/gml _TimePrimitive,omitempty"`
	AbstractTimeReferenceSystem       *AbstractTimeReferenceSystemType       `xml:"http://www.opengis.net/gml _TimeReferenceSystem,omitempty"`
	AbstractTimeSlice                 *AbstractTimeSliceType                 `xml:"http://www.opengis.net/gml _TimeSlice,omitempty"`
	AbstractTimeTopologyPrimitive     *AbstractTimeTopologyPrimitiveType     `xml:"http://www.opengis.net/gml _TimeTopologyPrimitive,omitempty"`
	AbstractTopoPrimitive             *AbstractTopoPrimitiveType             `xml:"http://www.opengis.net/gml _TopoPrimitive,omitempty"`
	AbstractTopology                  *AbstractTopologyType                  `xml:"http://www.opengis.net/gml _Topology,omitempty"`
	Null                              *string                                `xml:"http://www.opengis.net/gml Null,omitempty"`
	// Aggregate value built using the Composite pattern.
	CompositeValue *CompositeValueType `xml:"http://www.opengis.net/gml CompositeValue,omitempty"`
	ValueArray     *ValueArrayType     `xml:"http://www.opengis.net/gml ValueArray,omitempty"`
	// A value from two-valued logic, using the XML Schema boolean type.  An instance may take the values {true, false, 1, 0}.
	Boolean *bool `xml:"http://www.opengis.net/gml Boolean,omitempty"`
	// A term representing a classification.  It has an optional XML attribute codeSpace, whose value is a URI which identifies a dictionary, codelist or authority for the term.
	Category *CodeType `xml:"http://www.opengis.net/gml Category,omitempty"`
	// A numeric value with a scale.  The content of the element is an amount using the XML Schema type double which permits decimal or scientific notation.  An XML attribute uom (unit of measure) is required, whose value is a URI which identifies the definition of the scale or units by which the numeric value must be multiplied.
	Quantity *MeasureType `xml:"http://www.opengis.net/gml Quantity,omitempty"`
	// An integer representing a frequency of occurrence.
	Count *int `xml:"http://www.opengis.net/gml Count,omitempty"`
	// XML List based on XML Schema boolean type.  An element of this type contains a space-separated list of boolean values {0,1,true,false}
	BooleanList *string `xml:"http://www.opengis.net/gml BooleanList,omitempty"`
	// A space-separated list of terms or nulls.  A single XML attribute codeSpace may be provided, which authorises all the terms in the list.
	CategoryList *CodeOrNullListType `xml:"http://www.opengis.net/gml CategoryList,omitempty"`
	// A space separated list of amounts or nulls.  The amounts use the XML Schema type double.  A single XML attribute uom (unit of measure) is required, whose value is a URI which identifies the definition of the scale or units by which all the amounts in the list must be multiplied.
	QuantityList *MeasureOrNullListType `xml:"http://www.opengis.net/gml QuantityList,omitempty"`
	// A space-separated list of integers or nulls.
	CountList *string `xml:"http://www.opengis.net/gml CountList,omitempty"`
	// Utility element to store a 2-point range of ordinal values. If one member is a null, then this is a single ended interval.
	CategoryExtent *CategoryExtentType `xml:"http://www.opengis.net/gml CategoryExtent,omitempty"`
	// Utility element to store a 2-point range of numeric values. If one member is a null, then this is a single ended interval.
	QuantityExtent *QuantityExtentType `xml:"http://www.opengis.net/gml QuantityExtent,omitempty"`
	// Utility element to store a 2-point range of frequency values. If one member is a null, then this is a single ended interval.
	CountExtent  *string `xml:"http://www.opengis.net/gml CountExtent,omitempty"`
	RemoteSchema string  `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string  `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string  `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string  `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string  `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string  `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string  `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string  `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type ScaleType struct {
	Value string `xml:",chardata"`
}

type SecondDefiningParameterType struct {
	// Inverse flattening value of the ellipsoid. Value is a scale factor (or ratio) that has no physical unit. Uses the MeasureType with the restriction that the unit of measure referenced by uom must be suitable for a scale factor, such as percent, permil, or parts-per-million.
	InverseFlattening *MeasureType `xml:"http://www.opengis.net/gml inverseFlattening,omitempty"`
	// Length of the semi-minor axis of the ellipsoid. Uses the MeasureType with the restriction that the unit of measure referenced by uom must be suitable for a length, such as metres or feet.
	SemiMinorAxis *MeasureType `xml:"http://www.opengis.net/gml semiMinorAxis,omitempty"`
	// The ellipsoid is degenerate and is actually a sphere. The sphere is completely defined by the semi-major axis, which is the radius of the sphere.
	IsSphere *string `xml:"http://www.opengis.net/gml isSphere,omitempty"`
}

type SequenceRuleType struct {
	Value string  `xml:",chardata"`
	Order *string `xml:"order,attr,omitempty"`
}

type SingleOperationRefType struct {
	// A single (not concatenated) coordinate operation.
	AbstractSingleOperation       *AbstractCoordinateOperationType   `xml:"http://www.opengis.net/gml _SingleOperation,omitempty"`
	Conversion                    *ConversionType                    `xml:"http://www.opengis.net/gml Conversion,omitempty"`
	PassThroughOperation          *PassThroughOperationType          `xml:"http://www.opengis.net/gml PassThroughOperation,omitempty"`
	Transformation                *TransformationType                `xml:"http://www.opengis.net/gml Transformation,omitempty"`
	AbstractGeneralConversion     *AbstractGeneralConversionType     `xml:"http://www.opengis.net/gml _GeneralConversion,omitempty"`
	AbstractGeneralTransformation *AbstractGeneralTransformationType `xml:"http://www.opengis.net/gml _GeneralTransformation,omitempty"`
	AbstractOperation             *AbstractCoordinateOperationType   `xml:"http://www.opengis.net/gml _Operation,omitempty"`
	RemoteSchema                  string                             `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                     string                             `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                          string                             `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                          string                             `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                       string                             `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                         string                             `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                          string                             `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                       string                             `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type SolidArrayPropertyType struct {
	// The "_Solid" element is the abstract head of the substituition group for all (continuous) solid elements.
	AbstractSolid  []AbstractSolidType  `xml:"http://www.opengis.net/gml _Solid"`
	CompositeSolid []CompositeSolidType `xml:"http://www.opengis.net/gml CompositeSolid"`
	Solid          []SolidType          `xml:"http://www.opengis.net/gml Solid"`
}

type SolidPropertyType struct {
	// The "_Solid" element is the abstract head of the substituition group for all (continuous) solid elements.
	AbstractSolid  *AbstractSolidType  `xml:"http://www.opengis.net/gml _Solid,omitempty"`
	CompositeSolid *CompositeSolidType `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	Solid          *SolidType          `xml:"http://www.opengis.net/gml Solid,omitempty"`
	RemoteSchema   string              `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField      string              `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href           string              `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role           string              `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole        string              `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title          string              `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show           string              `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate        string              `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type SolidType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
	// Boundaries of solids are similar to surface boundaries. In normal 3-dimensional Euclidean space, one (composite) surface is distinguished as the exterior. In the more general case, this is not always possible.
	Exterior *SurfacePropertyType `xml:"http://www.opengis.net/gml exterior,omitempty"`
	// Boundaries of solids are similar to surface boundaries.
	Interior []SurfacePropertyType `xml:"http://www.opengis.net/gml interior"`
}

type SpeedType struct {
	Value string `xml:",chardata"`
}

type SphereType struct {
	// The attribute rows gives the number
	// of rows in the parameter grid.
	Rows int `xml:"http://www.opengis.net/gml rows,omitempty"`
	// The attribute columns gives the number
	// of columns in the parameter grid.
	Columns             int      `xml:"http://www.opengis.net/gml columns,omitempty"`
	Row                 []string `xml:"http://www.opengis.net/gml row"`
	HorizontalCurveType *string  `xml:"horizontalCurveType,attr,omitempty"`
	VerticalCurveType   *string  `xml:"verticalCurveType,attr,omitempty"`
}

type SphericalCSRefType struct {
	SphericalCS  *SphericalCSType `xml:"http://www.opengis.net/gml SphericalCS,omitempty"`
	RemoteSchema string           `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string           `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string           `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string           `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string           `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string           `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string           `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string           `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type SphericalCSType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this coordinate system. The first csID, if any, is normally the primary identification code, and any others are aliases.
	CsID []IdentifierType `xml:"http://www.opengis.net/gml csID"`
	// Comments on or information about this coordinate system, including data source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Ordered sequence of associations to the coordinate system axes included in this coordinate system.
	UsesAxis []CoordinateSystemAxisRefType `xml:"http://www.opengis.net/gml usesAxis"`
}

type StringOrRefType struct {
	Value        string `xml:",chardata"`
	RemoteSchema string `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type StyleType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType                 `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType                 `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType                 `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType                 `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType                 `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType                 `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType                 `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType                 `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType                 `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType                 `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	FeatureStyle            []FeatureStylePropertyType `xml:"http://www.opengis.net/gml featureStyle"`
	GraphStyle              *GraphStylePropertyType    `xml:"http://www.opengis.net/gml graphStyle,omitempty"`
}

type StyleVariationType struct {
	Value                string  `xml:",chardata"`
	StyleProperty        string  `xml:"styleProperty,attr"`
	FeaturePropertyRange *string `xml:"featurePropertyRange,attr,omitempty"`
}

type SurfaceArrayPropertyType struct {
	// The "_Surface" element is the abstract head of the substituition group for all (continuous) surface elements.
	AbstractSurface     []AbstractSurfaceType     `xml:"http://www.opengis.net/gml _Surface"`
	CompositeSurface    []CompositeSurfaceType    `xml:"http://www.opengis.net/gml CompositeSurface"`
	OrientableSurface   []OrientableSurfaceType   `xml:"http://www.opengis.net/gml OrientableSurface"`
	Polygon             []PolygonType             `xml:"http://www.opengis.net/gml Polygon"`
	PolyhedralSurface   []PolyhedralSurfaceType   `xml:"http://www.opengis.net/gml PolyhedralSurface"`
	Surface             []SurfaceType             `xml:"http://www.opengis.net/gml Surface"`
	Tin                 []TinType                 `xml:"http://www.opengis.net/gml Tin"`
	TriangulatedSurface []TriangulatedSurfaceType `xml:"http://www.opengis.net/gml TriangulatedSurface"`
}

type SurfacePatchArrayPropertyType struct {
	// The "_SurfacePatch" element is the abstract head of the substituition group for all surface pach elements describing a continuous portion of a surface.
	AbstractSurfacePatch           []AbstractSurfacePatchType           `xml:"http://www.opengis.net/gml _SurfacePatch"`
	Cone                           []ConeType                           `xml:"http://www.opengis.net/gml Cone"`
	Cylinder                       []CylinderType                       `xml:"http://www.opengis.net/gml Cylinder"`
	PolygonPatch                   []PolygonPatchType                   `xml:"http://www.opengis.net/gml PolygonPatch"`
	Rectangle                      []RectangleType                      `xml:"http://www.opengis.net/gml Rectangle"`
	Sphere                         []SphereType                         `xml:"http://www.opengis.net/gml Sphere"`
	Triangle                       []TriangleType                       `xml:"http://www.opengis.net/gml Triangle"`
	AbstractGriddedSurface         []AbstractGriddedSurfaceType         `xml:"http://www.opengis.net/gml _GriddedSurface"`
	AbstractParametricCurveSurface []AbstractParametricCurveSurfaceType `xml:"http://www.opengis.net/gml _ParametricCurveSurface"`
}

type SurfacePropertyType struct {
	// The "_Surface" element is the abstract head of the substituition group for all (continuous) surface elements.
	AbstractSurface     *AbstractSurfaceType     `xml:"http://www.opengis.net/gml _Surface,omitempty"`
	CompositeSurface    *CompositeSurfaceType    `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	OrientableSurface   *OrientableSurfaceType   `xml:"http://www.opengis.net/gml OrientableSurface,omitempty"`
	Polygon             *PolygonType             `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	PolyhedralSurface   *PolyhedralSurfaceType   `xml:"http://www.opengis.net/gml PolyhedralSurface,omitempty"`
	Surface             *SurfaceType             `xml:"http://www.opengis.net/gml Surface,omitempty"`
	Tin                 *TinType                 `xml:"http://www.opengis.net/gml Tin,omitempty"`
	TriangulatedSurface *TriangulatedSurfaceType `xml:"http://www.opengis.net/gml TriangulatedSurface,omitempty"`
	RemoteSchema        string                   `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField           string                   `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                string                   `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                string                   `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole             string                   `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title               string                   `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                string                   `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate             string                   `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type SurfaceType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	Id      string  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
	// This element encapsulates the patches of the surface.
	Patches         *SurfacePatchArrayPropertyType  `xml:"http://www.opengis.net/gml patches,omitempty"`
	PolygonPatches  *PolygonPatchArrayPropertyType  `xml:"http://www.opengis.net/gml polygonPatches,omitempty"`
	TrianglePatches *TrianglePatchArrayPropertyType `xml:"http://www.opengis.net/gml trianglePatches,omitempty"`
}

type SymbolType struct {
	SymbolType   string  `xml:"symbolType,attr"`
	Transform    string  `xml:"http://www.opengis.net/gml transform,attr,omitempty"`
	About        *string `xml:"about,attr,omitempty"`
	RemoteSchema string  `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string  `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string  `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string  `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string  `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string  `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string  `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string  `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type TargetPropertyType struct {
	AbstractFeature               *AbstractFeatureType               `xml:"http://www.opengis.net/gml _Feature,omitempty"`
	DirectedObservation           *DirectedObservationType           `xml:"http://www.opengis.net/gml DirectedObservation,omitempty"`
	DirectedObservationAtDistance *DirectedObservationAtDistanceType `xml:"http://www.opengis.net/gml DirectedObservationAtDistance,omitempty"`
	FeatureCollection             *FeatureCollectionType             `xml:"http://www.opengis.net/gml FeatureCollection,omitempty"`
	GridCoverage                  *GridCoverageType                  `xml:"http://www.opengis.net/gml GridCoverage,omitempty"`
	MultiCurveCoverage            *MultiCurveCoverageType            `xml:"http://www.opengis.net/gml MultiCurveCoverage,omitempty"`
	MultiPointCoverage            *MultiPointCoverageType            `xml:"http://www.opengis.net/gml MultiPointCoverage,omitempty"`
	MultiSolidCoverage            *MultiSolidCoverageType            `xml:"http://www.opengis.net/gml MultiSolidCoverage,omitempty"`
	MultiSurfaceCoverage          *MultiSurfaceCoverageType          `xml:"http://www.opengis.net/gml MultiSurfaceCoverage,omitempty"`
	Observation                   *ObservationType                   `xml:"http://www.opengis.net/gml Observation,omitempty"`
	RectifiedGridCoverage         *RectifiedGridCoverageType         `xml:"http://www.opengis.net/gml RectifiedGridCoverage,omitempty"`
	AbstractContinuousCoverage    *AbstractContinuousCoverageType    `xml:"http://www.opengis.net/gml _ContinuousCoverage,omitempty"`
	AbstractCoverage              *AbstractCoverageType              `xml:"http://www.opengis.net/gml _Coverage,omitempty"`
	AbstractDiscreteCoverage      *AbstractDiscreteCoverageType      `xml:"http://www.opengis.net/gml _DiscreteCoverage,omitempty"`
	AbstractFeatureCollection     *AbstractFeatureCollectionType     `xml:"http://www.opengis.net/gml _FeatureCollection,omitempty"`
	// The "_Geometry" element is the abstract head of the substituition group for all geometry elements of GML 3. This
	// includes pre-defined and user-defined geometry elements. Any geometry element must be a direct or indirect extension/restriction
	// of AbstractGeometryType and must be directly or indirectly in the substitution group of "_Geometry".
	AbstractGeometry           *AbstractGeometryType           `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	CompositeCurve             *CompositeCurveType             `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	CompositeSolid             *CompositeSolidType             `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	CompositeSurface           *CompositeSurfaceType           `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	Curve                      *CurveType                      `xml:"http://www.opengis.net/gml Curve,omitempty"`
	GeometricComplex           *GeometricComplexType           `xml:"http://www.opengis.net/gml GeometricComplex,omitempty"`
	Grid                       *GridType                       `xml:"http://www.opengis.net/gml Grid,omitempty"`
	LineString                 *LineStringType                 `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearRing                 *LinearRingType                 `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MultiCurve                 *MultiCurveType                 `xml:"http://www.opengis.net/gml MultiCurve,omitempty"`
	MultiGeometry              *MultiGeometryType              `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString            *MultiLineStringType            `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                 *MultiPointType                 `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPolygon               *MultiPolygonType               `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	MultiSolid                 *MultiSolidType                 `xml:"http://www.opengis.net/gml MultiSolid,omitempty"`
	MultiSurface               *MultiSurfaceType               `xml:"http://www.opengis.net/gml MultiSurface,omitempty"`
	OrientableCurve            *OrientableCurveType            `xml:"http://www.opengis.net/gml OrientableCurve,omitempty"`
	OrientableSurface          *OrientableSurfaceType          `xml:"http://www.opengis.net/gml OrientableSurface,omitempty"`
	Point                      *PointType                      `xml:"http://www.opengis.net/gml Point,omitempty"`
	Polygon                    *PolygonType                    `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	PolyhedralSurface          *PolyhedralSurfaceType          `xml:"http://www.opengis.net/gml PolyhedralSurface,omitempty"`
	RectifiedGrid              *RectifiedGridType              `xml:"http://www.opengis.net/gml RectifiedGrid,omitempty"`
	Ring                       *RingType                       `xml:"http://www.opengis.net/gml Ring,omitempty"`
	Solid                      *SolidType                      `xml:"http://www.opengis.net/gml Solid,omitempty"`
	Surface                    *SurfaceType                    `xml:"http://www.opengis.net/gml Surface,omitempty"`
	Tin                        *TinType                        `xml:"http://www.opengis.net/gml Tin,omitempty"`
	TriangulatedSurface        *TriangulatedSurfaceType        `xml:"http://www.opengis.net/gml TriangulatedSurface,omitempty"`
	AbstractCurve              *AbstractCurveType              `xml:"http://www.opengis.net/gml _Curve,omitempty"`
	AbstractGeometricAggregate *AbstractGeometricAggregateType `xml:"http://www.opengis.net/gml _GeometricAggregate,omitempty"`
	AbstractGeometricPrimitive *AbstractGeometricPrimitiveType `xml:"http://www.opengis.net/gml _GeometricPrimitive,omitempty"`
	AbstractImplicitGeometry   *AbstractGeometryType           `xml:"http://www.opengis.net/gml _ImplicitGeometry,omitempty"`
	AbstractRing               *AbstractRingType               `xml:"http://www.opengis.net/gml _Ring,omitempty"`
	AbstractSolid              *AbstractSolidType              `xml:"http://www.opengis.net/gml _Solid,omitempty"`
	AbstractSurface            *AbstractSurfaceType            `xml:"http://www.opengis.net/gml _Surface,omitempty"`
	RemoteSchema               string                          `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                  string                          `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                       string                          `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                       string                          `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                    string                          `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                      string                          `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                       string                          `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                    string                          `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type TemporalCRSRefType struct {
	TemporalCRS  *TemporalCRSType `xml:"http://www.opengis.net/gml TemporalCRS,omitempty"`
	RemoteSchema string           `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string           `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string           `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string           `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string           `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string           `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string           `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string           `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type TemporalCRSType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alterative identifications of this reference system. The first srsID, if any, is normally the primary identification code, and any others are aliases.
	SrsID []IdentifierType `xml:"http://www.opengis.net/gml srsID"`
	// Comments on or information about this reference system, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
	// Association to the temporal coordinate system used by this CRS.
	UsesTemporalCS *TemporalCSRefType `xml:"http://www.opengis.net/gml usesTemporalCS,omitempty"`
	// Association to the temporal datum used by this CRS.
	UsesTemporalDatum *TemporalDatumRefType `xml:"http://www.opengis.net/gml usesTemporalDatum,omitempty"`
}

type TemporalCSRefType struct {
	TemporalCS   *TemporalCSType `xml:"http://www.opengis.net/gml TemporalCS,omitempty"`
	RemoteSchema string          `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string          `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string          `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string          `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string          `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string          `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string          `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string          `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type TemporalCSType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this coordinate system. The first csID, if any, is normally the primary identification code, and any others are aliases.
	CsID []IdentifierType `xml:"http://www.opengis.net/gml csID"`
	// Comments on or information about this coordinate system, including data source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Ordered sequence of associations to the coordinate system axes included in this coordinate system.
	UsesAxis []CoordinateSystemAxisRefType `xml:"http://www.opengis.net/gml usesAxis"`
}

type TemporalDatumBaseType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this datum. The first datumID, if any, is normally the primary identification code, and any others are aliases.
	DatumID []IdentifierType `xml:"http://www.opengis.net/gml datumID"`
	// Comments on this reference system, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Description, possibly including coordinates, of the point or points used to anchor the datum to the Earth. Also known as the "origin", especially for engineering and image datums. The codeSpace attribute can be used to reference a source of more detailed on this point or surface, or on a set of such descriptions.
	// - For a geodetic datum, this point is also known as the fundamental point, which is traditionally the point where the relationship between geoid and ellipsoid is defined. In some cases, the "fundamental point" may consist of a number of points. In those cases, the parameters defining the geoid/ellipsoid relationship have been averaged for these points, and the averages adopted as the datum definition.
	// - For an engineering datum, the anchor point may be a physical point, or it may be a point with defined coordinates in another CRS. When appropriate, the coordinates of this anchor point can be referenced in another document, such as referencing a GML feature that references or includes a point position.
	// - For an image datum, the anchor point is usually either the centre of the image or the corner of the image.
	// - For a temporal datum, this attribute is not defined. Instead of the anchor point, a temporal datum carries a separate time origin of type DateTime.
	AnchorPoint *CodeType `xml:"http://www.opengis.net/gml anchorPoint,omitempty"`
	// The time after which this datum definition is valid. This time may be precise (e.g. 1997.0 for IRTF97) or merely a year (e.g. 1983 for NAD83). In the latter case, the epoch usually refers to the year in which a major recalculation of the geodetic control network, underlying the datum, was executed or initiated. An old datum can remain valid after a new datum is defined. Alternatively, a datum may be superseded by a later datum, in which case the realization epoch for the new datum defines the upper limit for the validity of the superseded datum.
	RealizationEpoch *string `xml:"http://www.opengis.net/gml realizationEpoch,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
}

type TemporalDatumRefType struct {
	TemporalDatum *TemporalDatumType `xml:"http://www.opengis.net/gml TemporalDatum,omitempty"`
	RemoteSchema  string             `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField     string             `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href          string             `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role          string             `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole       string             `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title         string             `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show          string             `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate       string             `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type TemporalDatumType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this datum. The first datumID, if any, is normally the primary identification code, and any others are aliases.
	DatumID []IdentifierType `xml:"http://www.opengis.net/gml datumID"`
	// Comments on this reference system, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Description, possibly including coordinates, of the point or points used to anchor the datum to the Earth. Also known as the "origin", especially for engineering and image datums. The codeSpace attribute can be used to reference a source of more detailed on this point or surface, or on a set of such descriptions.
	// - For a geodetic datum, this point is also known as the fundamental point, which is traditionally the point where the relationship between geoid and ellipsoid is defined. In some cases, the "fundamental point" may consist of a number of points. In those cases, the parameters defining the geoid/ellipsoid relationship have been averaged for these points, and the averages adopted as the datum definition.
	// - For an engineering datum, the anchor point may be a physical point, or it may be a point with defined coordinates in another CRS. When appropriate, the coordinates of this anchor point can be referenced in another document, such as referencing a GML feature that references or includes a point position.
	// - For an image datum, the anchor point is usually either the centre of the image or the corner of the image.
	// - For a temporal datum, this attribute is not defined. Instead of the anchor point, a temporal datum carries a separate time origin of type DateTime.
	AnchorPoint *CodeType `xml:"http://www.opengis.net/gml anchorPoint,omitempty"`
	// The time after which this datum definition is valid. This time may be precise (e.g. 1997.0 for IRTF97) or merely a year (e.g. 1983 for NAD83). In the latter case, the epoch usually refers to the year in which a major recalculation of the geodetic control network, underlying the datum, was executed or initiated. An old datum can remain valid after a new datum is defined. Alternatively, a datum may be superseded by a later datum, in which case the realization epoch for the new datum defines the upper limit for the validity of the superseded datum.
	RealizationEpoch *string `xml:"http://www.opengis.net/gml realizationEpoch,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
	// The date and time origin of this temporal datum.
	Origin *string `xml:"http://www.opengis.net/gml origin,omitempty"`
}

type TimeCalendarEraPropertyType struct {
	TimeCalendarEra *TimeCalendarEraType `xml:"http://www.opengis.net/gml TimeCalendarEra,omitempty"`
	RemoteSchema    string               `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField       string               `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href            string               `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role            string               `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole         string               `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title           string               `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show            string               `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate         string               `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type TimeCalendarEraType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Name or description of a mythical or historic event which fixes the position of the base scale of the calendar era.
	ReferenceEvent *StringOrRefType `xml:"http://www.opengis.net/gml referenceEvent,omitempty"`
	// Date of the referenceEvent expressed as a date in the given calendar.
	// In most calendars, this date is the origin (i.e., the first day) of the scale, but this is not always true.
	ReferenceDate string `xml:"http://www.opengis.net/gml referenceDate,omitempty"`
	// Julian date that corresponds to the reference date.
	// The Julian day numbering system is a temporal coordinate system that has an
	// origin earlier than any known calendar,
	// at noon on 1 January 4713 BC in the Julian proleptic calendar.
	// The Julian day number is an integer value;
	// the Julian date is a decimal value that allows greater resolution.
	// Transforming calendar dates to and from Julian dates provides a
	// relatively simple basis for transforming dates from one calendar to another.
	JulianReference float64 `xml:"http://www.opengis.net/gml julianReference,omitempty"`
	// Period for which the calendar era was used as a basis for dating.
	EpochOfUse *TimePeriodPropertyType `xml:"http://www.opengis.net/gml epochOfUse,omitempty"`
}

type TimeCalendarPropertyType struct {
	TimeCalendar *TimeCalendarType `xml:"http://www.opengis.net/gml TimeCalendar,omitempty"`
	RemoteSchema string            `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string            `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string            `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string            `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string            `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string            `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string            `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string            `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type TimeCalendarType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	DomainOfValidity        string     `xml:"http://www.opengis.net/gml domainOfValidity,omitempty"`
	// Link to the CalendarEras that it uses as a reference for dating.
	ReferenceFrame []TimeCalendarEraPropertyType `xml:"http://www.opengis.net/gml referenceFrame"`
}

type TimeClockPropertyType struct {
	TimeClock    *TimeClockType `xml:"http://www.opengis.net/gml TimeClock,omitempty"`
	RemoteSchema string         `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string         `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string         `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string         `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string         `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string         `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string         `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string         `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type TimeClockType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	DomainOfValidity        string     `xml:"http://www.opengis.net/gml domainOfValidity,omitempty"`
	// Name or description of an event, such as solar noon or sunrise,
	// which fixes the position of the base scale of the clock.
	ReferenceEvent *StringOrRefType `xml:"http://www.opengis.net/gml referenceEvent,omitempty"`
	// time of day associated with the reference event expressed as
	// a time of day in the given clock. The reference time is usually the origin of the clock scale.
	ReferenceTime string `xml:"http://www.opengis.net/gml referenceTime,omitempty"`
	// 24 hour local or UTC time that corresponds to the reference time.
	UtcReference string                     `xml:"http://www.opengis.net/gml utcReference,omitempty"`
	DateBasis    []TimeCalendarPropertyType `xml:"http://www.opengis.net/gml dateBasis"`
}

type TimeCoordinateSystemType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType               `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType               `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType               `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType               `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType               `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType               `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType               `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType               `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType               `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType               `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                   `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	DomainOfValidity        string                   `xml:"http://www.opengis.net/gml domainOfValidity,omitempty"`
	Interval                *TimeIntervalLengthType  `xml:"http://www.opengis.net/gml interval,omitempty"`
	OriginPosition          *TimePositionType        `xml:"http://www.opengis.net/gml originPosition,omitempty"`
	Origin                  *TimeInstantPropertyType `xml:"http://www.opengis.net/gml origin,omitempty"`
}

type TimeEdgePropertyType struct {
	// TimeEdge is one dimensional temporal topology primitive,
	// expresses a state in topological time. It has an orientation from its start toward the end,
	// and its boundaries shall associate with two different time nodes.
	TimeEdge     *TimeEdgeType `xml:"http://www.opengis.net/gml TimeEdge,omitempty"`
	RemoteSchema string        `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string        `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string        `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string        `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string        `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string        `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string        `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string        `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type TimeEdgeType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType              `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType              `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType              `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType              `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType              `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType              `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType              `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType              `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType              `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType              `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	RelatedTime             []RelatedTimeType       `xml:"http://www.opengis.net/gml relatedTime"`
	Complex                 *ReferenceType          `xml:"http://www.opengis.net/gml complex,omitempty"`
	Start                   *TimeNodePropertyType   `xml:"http://www.opengis.net/gml start,omitempty"`
	End                     *TimeNodePropertyType   `xml:"http://www.opengis.net/gml end,omitempty"`
	Extent                  *TimePeriodPropertyType `xml:"http://www.opengis.net/gml extent,omitempty"`
}

type TimeGeometricPrimitivePropertyType struct {
	// This abstract element acts as the head of the substitution group for temporal geometric primitives.
	AbstractTimeGeometricPrimitive *AbstractTimeGeometricPrimitiveType `xml:"http://www.opengis.net/gml _TimeGeometricPrimitive,omitempty"`
	TimeInstant                    *TimeInstantType                    `xml:"http://www.opengis.net/gml TimeInstant,omitempty"`
	TimePeriod                     *TimePeriodType                     `xml:"http://www.opengis.net/gml TimePeriod,omitempty"`
	RemoteSchema                   string                              `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                      string                              `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                           string                              `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                           string                              `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                        string                              `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                          string                              `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                           string                              `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                        string                              `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type TimeInstantPropertyType struct {
	TimeInstant  *TimeInstantType `xml:"http://www.opengis.net/gml TimeInstant,omitempty"`
	RemoteSchema string           `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string           `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string           `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string           `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string           `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string           `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string           `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string           `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type TimeInstantType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType        `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType        `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType        `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType        `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType        `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType        `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType        `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType        `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType        `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType        `xml:"http://www.opengis.net/gml srsName"`
	Id                      string            `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	RelatedTime             []RelatedTimeType `xml:"http://www.opengis.net/gml relatedTime"`
	Frame                   *string           `xml:"frame,attr,omitempty"`
	// Direct representation of a temporal position
	TimePosition *TimePositionType `xml:"http://www.opengis.net/gml timePosition,omitempty"`
}

type TimeIntervalLengthType struct {
	Value  float64 `xml:",chardata"`
	Unit   string  `xml:"unit,attr"`
	Radix  *int    `xml:"radix,attr,omitempty"`
	Factor *int    `xml:"factor,attr,omitempty"`
}

type TimeNodePropertyType struct {
	// "TimeNode" is a zero dimensional temporal topology primitive,
	// expresses a position in topological time, and is a start and an end of time edge, which represents states of time.
	// Time node may be isolated. However, it cannot describe the ordering relationships with other primitives.
	// An isolated node may not be an element of any temporal topology complex.
	TimeNode     *TimeNodeType `xml:"http://www.opengis.net/gml TimeNode,omitempty"`
	RemoteSchema string        `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string        `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string        `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string        `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string        `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string        `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string        `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string        `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type TimeNodeType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType               `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType               `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType               `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType               `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType               `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType               `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType               `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType               `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType               `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType               `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                   `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	RelatedTime             []RelatedTimeType        `xml:"http://www.opengis.net/gml relatedTime"`
	Complex                 *ReferenceType           `xml:"http://www.opengis.net/gml complex,omitempty"`
	PreviousEdge            []TimeEdgePropertyType   `xml:"http://www.opengis.net/gml previousEdge"`
	NextEdge                []TimeEdgePropertyType   `xml:"http://www.opengis.net/gml nextEdge"`
	Position                *TimeInstantPropertyType `xml:"http://www.opengis.net/gml position,omitempty"`
}

type TimeOrdinalEraPropertyType struct {
	TimeOrdinalEra *TimeOrdinalEraType `xml:"http://www.opengis.net/gml TimeOrdinalEra,omitempty"`
	RemoteSchema   string              `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField      string              `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href           string              `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role           string              `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole        string              `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title          string              `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show           string              `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate        string              `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type TimeOrdinalEraType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType              `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType              `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType              `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType              `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType              `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType              `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType              `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType              `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType              `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType              `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	RelatedTime             []RelatedTimeType       `xml:"http://www.opengis.net/gml relatedTime"`
	Start                   *TimeNodePropertyType   `xml:"http://www.opengis.net/gml start,omitempty"`
	End                     *TimeNodePropertyType   `xml:"http://www.opengis.net/gml end,omitempty"`
	Extent                  *TimePeriodPropertyType `xml:"http://www.opengis.net/gml extent,omitempty"`
	// An Era may be composed of several member Eras. The "member" element implements the association to the Era at the next level down the hierarchy.  "member" follows the standard GML property pattern whereby its (complex) value may be either described fully inline, or may be the target of a link carried on the member element and described fully elsewhere, either in the same document or from another service.
	Member []TimeOrdinalEraPropertyType `xml:"http://www.opengis.net/gml member"`
	// In a particular Time System, an Era may be a member of a group.  The "group" element implements the back-pointer to the Era at the next level up in the hierarchy.
	//
	// If the hierarchy is represented by describing the nested components fully in the their nested position inside "member" elements, then the parent can be easily inferred, so the group property is unnecessary.
	//
	// However, if the hierarchy is represented by links carried on the "member" property elements, pointing to Eras described fully elsewhere, then it may be useful for a child (member) era to carry an explicit pointer back to its parent (group) Era.
	Group *ReferenceType `xml:"http://www.opengis.net/gml group,omitempty"`
}

type TimeOrdinalReferenceSystemType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType                   `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType                   `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType                   `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType                   `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType                   `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType                   `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType                   `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType                   `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType                   `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType                   `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                       `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	DomainOfValidity        string                       `xml:"http://www.opengis.net/gml domainOfValidity,omitempty"`
	Component               []TimeOrdinalEraPropertyType `xml:"http://www.opengis.net/gml component"`
}

type TimePeriodPropertyType struct {
	TimePeriod   *TimePeriodType `xml:"http://www.opengis.net/gml TimePeriod,omitempty"`
	RemoteSchema string          `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string          `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string          `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string          `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string          `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string          `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string          `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string          `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type TimePeriodType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType               `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType               `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType               `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType               `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType               `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType               `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType               `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType               `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType               `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType               `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                   `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	RelatedTime             []RelatedTimeType        `xml:"http://www.opengis.net/gml relatedTime"`
	Frame                   *string                  `xml:"frame,attr,omitempty"`
	BeginPosition           *TimePositionType        `xml:"http://www.opengis.net/gml beginPosition,omitempty"`
	Begin                   *TimeInstantPropertyType `xml:"http://www.opengis.net/gml begin,omitempty"`
	EndPosition             *TimePositionType        `xml:"http://www.opengis.net/gml endPosition,omitempty"`
	End                     *TimeInstantPropertyType `xml:"http://www.opengis.net/gml end,omitempty"`
	// This element is an instance of the primitive xsd:duration simple type to
	// enable use of the ISO 8601 syntax for temporal length (e.g. P5DT4H30M).
	// It is a valid subtype of TimeDurationType according to section 3.14.6,
	// rule 2.2.4 in XML Schema, Part 1.
	Duration *string `xml:"http://www.opengis.net/gml duration,omitempty"`
	// This element is a valid subtype of TimeDurationType
	// according to section 3.14.6, rule 2.2.4 in XML Schema, Part 1.
	TimeInterval *TimeIntervalLengthType `xml:"http://www.opengis.net/gml timeInterval,omitempty"`
}

type TimePositionType struct {
	Value                 string  `xml:",chardata"`
	Frame                 *string `xml:"frame,attr,omitempty"`
	CalendarEraName       *string `xml:"calendarEraName,attr,omitempty"`
	IndeterminatePosition *string `xml:"indeterminatePosition,attr,omitempty"`
}

type TimePrimitivePropertyType struct {
	// This abstract element acts as the head of the substitution group for temporal primitives.
	AbstractTimePrimitive          *AbstractTimePrimitiveType          `xml:"http://www.opengis.net/gml _TimePrimitive,omitempty"`
	TimeEdge                       *TimeEdgeType                       `xml:"http://www.opengis.net/gml TimeEdge,omitempty"`
	TimeInstant                    *TimeInstantType                    `xml:"http://www.opengis.net/gml TimeInstant,omitempty"`
	TimeNode                       *TimeNodeType                       `xml:"http://www.opengis.net/gml TimeNode,omitempty"`
	TimePeriod                     *TimePeriodType                     `xml:"http://www.opengis.net/gml TimePeriod,omitempty"`
	AbstractTimeGeometricPrimitive *AbstractTimeGeometricPrimitiveType `xml:"http://www.opengis.net/gml _TimeGeometricPrimitive,omitempty"`
	AbstractTimeTopologyPrimitive  *AbstractTimeTopologyPrimitiveType  `xml:"http://www.opengis.net/gml _TimeTopologyPrimitive,omitempty"`
	RemoteSchema                   string                              `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                      string                              `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                           string                              `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                           string                              `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                        string                              `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                          string                              `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                           string                              `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                        string                              `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type TimeTopologyComplexPropertyType struct {
	// This element represents temporal topology complex. It shall be the connected acyclic directed graph composed of time nodes and time edges.
	TimeTopologyComplex *TimeTopologyComplexType `xml:"http://www.opengis.net/gml TimeTopologyComplex,omitempty"`
	RemoteSchema        string                   `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField           string                   `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                string                   `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                string                   `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole             string                   `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title               string                   `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                string                   `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate             string                   `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type TimeTopologyComplexType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType                          `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType                          `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType                          `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType                          `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType                          `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType                          `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType                          `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType                          `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType                          `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType                          `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                              `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	Primitive               []TimeTopologyPrimitivePropertyType `xml:"http://www.opengis.net/gml primitive"`
}

type TimeTopologyPrimitivePropertyType struct {
	// This abstract element acts as the head of the substitution group for temporal topology primitives.
	AbstractTimeTopologyPrimitive *AbstractTimeTopologyPrimitiveType `xml:"http://www.opengis.net/gml _TimeTopologyPrimitive,omitempty"`
	TimeEdge                      *TimeEdgeType                      `xml:"http://www.opengis.net/gml TimeEdge,omitempty"`
	TimeNode                      *TimeNodeType                      `xml:"http://www.opengis.net/gml TimeNode,omitempty"`
	RemoteSchema                  string                             `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                     string                             `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                          string                             `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                          string                             `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                       string                             `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                         string                             `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                          string                             `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                       string                             `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type TimeType struct {
	Value string `xml:",chardata"`
}

type TinType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
	// This element encapsulates the patches of the surface.
	Patches         *SurfacePatchArrayPropertyType  `xml:"http://www.opengis.net/gml patches,omitempty"`
	PolygonPatches  *PolygonPatchArrayPropertyType  `xml:"http://www.opengis.net/gml polygonPatches,omitempty"`
	TrianglePatches *TrianglePatchArrayPropertyType `xml:"http://www.opengis.net/gml trianglePatches,omitempty"`
	// Stoplines are lines where the local
	// continuity or regularity of the surface is questionable.
	// In the area of these pathologies, triangles intersecting
	// a stopline shall be removed from the tin surface, leaving
	// holes in the surface. If coincidence occurs on surface
	// boundary triangles, the result shall be a change of the
	// surface boundary. Stoplines contains all these
	// pathological segments as a set of line strings.
	StopLines []LineStringSegmentArrayPropertyType `xml:"http://www.opengis.net/gml stopLines"`
	// Breaklines are lines of a critical
	// nature to the shape of the surface, representing local
	// ridges, or depressions (such as drainage lines) in the
	// surface. As such their constituent segments must be
	// included in the tin eve if doing so
	// violates the Delauny criterion. Break lines contains these
	// critical segments as a set of line strings.
	BreakLines []LineStringSegmentArrayPropertyType `xml:"http://www.opengis.net/gml breakLines"`
	// Areas of the surface where data is not
	// sufficiently dense to assure reasonable calculation shall be
	// removed by adding a retention criterion for triangles based
	// on the length of their sides. For many triangle sides
	// exceeding maximum length, the adjacent triangles to that
	// triangle side shall be removed from the surface.
	MaxLength *LengthType `xml:"http://www.opengis.net/gml maxLength,omitempty"`
	// The corners of the triangles in the TIN
	// are often referred to as pots. ControlPoint shall contain a
	// set of the GM_Position used as posts for this TIN. Since each
	// TIN contains triangles, there must be at least 3 posts. The
	// order in which these points are given does not affect the
	// surface that is represented. Application schemas may add
	// information based on ordering of control points to facilitate
	// the reconstruction of the TIN from the control points.
	ControlPoint string `xml:"http://www.opengis.net/gml controlPoint,omitempty"`
}

type TopoComplexMemberType struct {
	TopoComplex  *TopoComplexType `xml:"http://www.opengis.net/gml TopoComplex,omitempty"`
	RemoteSchema string           `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string           `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string           `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string           `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string           `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string           `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string           `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string           `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type TopoComplexType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Need schamatron test here that isMaximal attribute value is true
	MaximalComplex       *TopoComplexMemberType             `xml:"http://www.opengis.net/gml maximalComplex,omitempty"`
	SuperComplex         []TopoComplexMemberType            `xml:"http://www.opengis.net/gml superComplex"`
	SubComplex           []TopoComplexMemberType            `xml:"http://www.opengis.net/gml subComplex"`
	TopoPrimitiveMember  []TopoPrimitiveMemberType          `xml:"http://www.opengis.net/gml topoPrimitiveMember"`
	TopoPrimitiveMembers *TopoPrimitiveArrayAssociationType `xml:"http://www.opengis.net/gml topoPrimitiveMembers,omitempty"`
	IsMaximal            *bool                              `xml:"isMaximal,attr,omitempty"`
}

type TopoCurvePropertyType struct {
	TopoCurve *TopoCurveType `xml:"http://www.opengis.net/gml TopoCurve,omitempty"`
}

type TopoCurveType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType                 `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType                 `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType                 `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType                 `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType                 `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType                 `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType                 `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType                 `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType                 `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType                 `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	DirectedEdge            []DirectedEdgePropertyType `xml:"http://www.opengis.net/gml directedEdge"`
}

type TopoPointPropertyType struct {
	TopoPoint *TopoPointType `xml:"http://www.opengis.net/gml TopoPoint,omitempty"`
}

type TopoPointType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType                `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType                `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType                `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType                `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType                `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType                `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType                `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType                `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType                `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType                `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                    `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	DirectedNode            *DirectedNodePropertyType `xml:"http://www.opengis.net/gml directedNode,omitempty"`
}

type TopoPrimitiveArrayAssociationType struct {
	// Substitution group branch for Topo Primitives, used by TopoPrimitiveArrayAssociationType
	AbstractTopoPrimitive []AbstractTopoPrimitiveType `xml:"http://www.opengis.net/gml _TopoPrimitive"`
	Edge                  []EdgeType                  `xml:"http://www.opengis.net/gml Edge"`
	Face                  []FaceType                  `xml:"http://www.opengis.net/gml Face"`
	Node                  []NodeType                  `xml:"http://www.opengis.net/gml Node"`
	TopoSolid             []TopoSolidType             `xml:"http://www.opengis.net/gml TopoSolid"`
}

type TopoPrimitiveMemberType struct {
	// Substitution group branch for Topo Primitives, used by TopoPrimitiveArrayAssociationType
	AbstractTopoPrimitive *AbstractTopoPrimitiveType `xml:"http://www.opengis.net/gml _TopoPrimitive,omitempty"`
	Edge                  *EdgeType                  `xml:"http://www.opengis.net/gml Edge,omitempty"`
	Face                  *FaceType                  `xml:"http://www.opengis.net/gml Face,omitempty"`
	Node                  *NodeType                  `xml:"http://www.opengis.net/gml Node,omitempty"`
	TopoSolid             *TopoSolidType             `xml:"http://www.opengis.net/gml TopoSolid,omitempty"`
	RemoteSchema          string                     `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField             string                     `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                  string                     `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                  string                     `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole               string                     `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                 string                     `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                  string                     `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate               string                     `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type TopoSolidType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType                 `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType                 `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType                 `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType                 `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType                 `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType                 `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType                 `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType                 `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType                 `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType                 `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	Isolated                []IsolatedPropertyType     `xml:"http://www.opengis.net/gml isolated"`
	Container               *ContainerPropertyType     `xml:"http://www.opengis.net/gml container,omitempty"`
	DirectedFace            []DirectedFacePropertyType `xml:"http://www.opengis.net/gml directedFace"`
}

type TopoSurfacePropertyType struct {
	TopoSurface *TopoSurfaceType `xml:"http://www.opengis.net/gml TopoSurface,omitempty"`
}

type TopoSurfaceType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType                 `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType                 `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType                 `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType                 `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType                 `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType                 `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType                 `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType                 `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType                 `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType                 `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	DirectedFace            []DirectedFacePropertyType `xml:"http://www.opengis.net/gml directedFace"`
}

type TopoVolumePropertyType struct {
	TopoVolume *TopoVolumeType `xml:"http://www.opengis.net/gml TopoVolume,omitempty"`
}

type TopoVolumeType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType                      `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType                      `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType                      `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType                      `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType                      `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType                      `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType                      `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType                      `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType                      `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType                      `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                          `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	DirectedTopoSolid       []DirectedTopoSolidPropertyType `xml:"http://www.opengis.net/gml directedTopoSolid"`
}

type TopologyStylePropertyType struct {
	// The style descriptor for topologies of a feature. Describes individual topology elements styles.
	TopologyStyle *TopologyStyleType `xml:"http://www.opengis.net/gml TopologyStyle,omitempty"`
	About         *string            `xml:"about,attr,omitempty"`
	RemoteSchema  string             `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField     string             `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href          string             `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role          string             `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole       string             `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title         string             `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show          string             `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate       string             `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type TopologyStyleType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType              `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType              `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType              `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType              `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType              `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType              `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType              `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType              `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType              `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType              `xml:"http://www.opengis.net/gml srsName"`
	Id                      string                  `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	SpatialResolution       *ScaleType              `xml:"http://www.opengis.net/gml spatialResolution,omitempty"`
	StyleVariation          []StyleVariationType    `xml:"http://www.opengis.net/gml styleVariation"`
	Animate                 []string                `xml:"http://www.w3.org/2001/SMIL20/ animate"`
	AnimateMotion           []string                `xml:"http://www.w3.org/2001/SMIL20/ animateMotion"`
	AnimateColor            []string                `xml:"http://www.w3.org/2001/SMIL20/ animateColor"`
	Set                     []string                `xml:"http://www.w3.org/2001/SMIL20/ set"`
	LabelStyle              *LabelStylePropertyType `xml:"http://www.opengis.net/gml labelStyle,omitempty"`
	// The symbol property. Extends the gml:AssociationType to allow for remote referencing of symbols.
	Symbol *SymbolType `xml:"http://www.opengis.net/gml symbol,omitempty"`
	// Deprecated in GML version 3.1.0. Use symbol with inline content instead.
	Style            string  `xml:"http://www.opengis.net/gml style,omitempty"`
	TopologyProperty *string `xml:"topologyProperty,attr,omitempty"`
	TopologyType     *string `xml:"topologyType,attr,omitempty"`
}

type TrackType struct {
	AbstractTimeSlice  []AbstractTimeSliceType  `xml:"http://www.opengis.net/gml _TimeSlice"`
	MovingObjectStatus []MovingObjectStatusType `xml:"http://www.opengis.net/gml MovingObjectStatus"`
}

type TransformationRefType struct {
	Transformation *TransformationType `xml:"http://www.opengis.net/gml Transformation,omitempty"`
	RemoteSchema   string              `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField      string              `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href           string              `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role           string              `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole        string              `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title          string              `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show           string              `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate        string              `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type TransformationType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this coordinate operation. The first coordinateOperationID, if any, is normally the primary identification code, and any others are aliases.
	CoordinateOperationID []IdentifierType `xml:"http://www.opengis.net/gml coordinateOperationID"`
	// Comments on or information about this coordinate operation, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Version of the coordinate transformation (i.e., instantiation due to the stochastic nature of the parameters). Mandatory when describing a transformation, and should not be supplied for a conversion.
	OperationVersion *string `xml:"http://www.opengis.net/gml operationVersion,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
	// Unordered set of estimates of the impact of this coordinate operation on point position accuracy. Gives position error estimates for target coordinates of this coordinate operation, assuming no errors in source coordinates.
	AbstractPositionalAccuracy         []AbstractPositionalAccuracyType         `xml:"http://www.opengis.net/gml _positionalAccuracy"`
	AbsoluteExternalPositionalAccuracy []AbsoluteExternalPositionalAccuracyType `xml:"http://www.opengis.net/gml absoluteExternalPositionalAccuracy"`
	CovarianceMatrix                   []CovarianceMatrixType                   `xml:"http://www.opengis.net/gml covarianceMatrix"`
	RelativeInternalPositionalAccuracy []RelativeInternalPositionalAccuracyType `xml:"http://www.opengis.net/gml relativeInternalPositionalAccuracy"`
	// Association to the source CRS (coordinate reference system) of this coordinate operation.
	SourceCRS *CRSRefType `xml:"http://www.opengis.net/gml sourceCRS,omitempty"`
	// Association to the target CRS (coordinate reference system) of this coordinate operation. For constraints on multiplicity of "sourceCRS" and "targetCRS", see UML model of Coordinate Operation package in OGC Abstract Specification topic 2.
	TargetCRS *CRSRefType `xml:"http://www.opengis.net/gml targetCRS,omitempty"`
	// Association to the operation method used by this coordinate operation.
	UsesMethod *OperationMethodRefType `xml:"http://www.opengis.net/gml usesMethod,omitempty"`
	// Unordered set of composition associations to the set of parameter values used by this transformation operation.
	UsesValue []ParameterValueType `xml:"http://www.opengis.net/gml usesValue"`
}

type TrianglePatchArrayPropertyType struct {
	// The "_SurfacePatch" element is the abstract head of the substituition group for all surface pach elements describing a continuous portion of a surface.
	AbstractSurfacePatch           []AbstractSurfacePatchType           `xml:"http://www.opengis.net/gml _SurfacePatch"`
	Cone                           []ConeType                           `xml:"http://www.opengis.net/gml Cone"`
	Cylinder                       []CylinderType                       `xml:"http://www.opengis.net/gml Cylinder"`
	PolygonPatch                   []PolygonPatchType                   `xml:"http://www.opengis.net/gml PolygonPatch"`
	Rectangle                      []RectangleType                      `xml:"http://www.opengis.net/gml Rectangle"`
	Sphere                         []SphereType                         `xml:"http://www.opengis.net/gml Sphere"`
	Triangle                       []TriangleType                       `xml:"http://www.opengis.net/gml Triangle"`
	AbstractGriddedSurface         []AbstractGriddedSurfaceType         `xml:"http://www.opengis.net/gml _GriddedSurface"`
	AbstractParametricCurveSurface []AbstractParametricCurveSurfaceType `xml:"http://www.opengis.net/gml _ParametricCurveSurface"`
}

type TriangleType struct {
	// Constraint: The Ring shall be a LinearRing and must form a triangle, the first and the last position must be co-incident.
	Exterior        *AbstractRingPropertyType `xml:"http://www.opengis.net/gml exterior,omitempty"`
	OuterBoundaryIs *AbstractRingPropertyType `xml:"http://www.opengis.net/gml outerBoundaryIs,omitempty"`
	// The attribute "interpolation" specifies the interpolation mechanism used for this surface patch. Currently only planar surface patches are defined in GML 3, the attribute is fixed to "planar", i.e. the interpolation method shall return points on a single plane. The boundary of the patch shall be contained within that plane.
	Interpolation *string `xml:"interpolation,attr,omitempty"`
}

type TriangulatedSurfaceType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// This attribute is included for backward compatibility with GML 2 and is deprecated with GML 3.
	// This identifer is superceded by "gml:id" inherited from AbstractGMLType. The attribute "gid" should not be used
	// anymore and may be deleted in future versions of GML without further notice.
	Gid *string `xml:"gid,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
	// This element encapsulates the patches of the surface.
	Patches         *SurfacePatchArrayPropertyType  `xml:"http://www.opengis.net/gml patches,omitempty"`
	PolygonPatches  *PolygonPatchArrayPropertyType  `xml:"http://www.opengis.net/gml polygonPatches,omitempty"`
	TrianglePatches *TrianglePatchArrayPropertyType `xml:"http://www.opengis.net/gml trianglePatches,omitempty"`
}

type UnitDefinitionType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Informal description of the phenomenon or type of quantity that is measured or observed. For example, "length", "angle", "time", "pressure", or "temperature". When the quantity is the result of an observation or measurement, this term is known as Observable Type or Measurand.
	QuantityType *StringOrRefType `xml:"http://www.opengis.net/gml quantityType,omitempty"`
	// For global understanding of a unit of measure, it is often possible to reference an item in a catalog of units, using a symbol in that catalog. The "codeSpace" attribute in "CodeType" identifies a namespace for the catalog symbol value, and might reference the catalog. The "string" value in "CodeType" contains the value of a symbol that is unique within this catalog namespace. This symbol often appears explicitly in the catalog, but it could be a combination of symbols using a specified algebra of units. For example, the symbol "cm" might indicate that it is the "m" symbol combined with the "c" prefix.
	CatalogSymbol *CodeType `xml:"http://www.opengis.net/gml catalogSymbol,omitempty"`
}

type UnitOfMeasureType struct {
	// Reference to a unit of measure definition, usually within the same XML document but possibly outside the XML document which contains this reference. For a reference within the same XML document, the "#" symbol should be used, followed by a text abbreviation of the unit name. However, the "#" symbol may be optional, and still may be interpreted as a reference.
	Uom string `xml:"uom,attr"`
}

type UserDefinedCSRefType struct {
	UserDefinedCS *UserDefinedCSType `xml:"http://www.opengis.net/gml UserDefinedCS,omitempty"`
	RemoteSchema  string             `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField     string             `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href          string             `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role          string             `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole       string             `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title         string             `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show          string             `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate       string             `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type UserDefinedCSType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this coordinate system. The first csID, if any, is normally the primary identification code, and any others are aliases.
	CsID []IdentifierType `xml:"http://www.opengis.net/gml csID"`
	// Comments on or information about this coordinate system, including data source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Ordered sequence of associations to the coordinate system axes included in this coordinate system.
	UsesAxis []CoordinateSystemAxisRefType `xml:"http://www.opengis.net/gml usesAxis"`
}

type ValueArrayPropertyType struct {
	// This abstract element is the head of a substitutionGroup hierararchy which may contain either simpleContent or complexContent elements.  It is used to assert the model position of "class" elements declared in other GML schemas.
	AbstractObject                    *string                                `xml:"http://www.opengis.net/gml _Object,omitempty"`
	Array                             *ArrayType                             `xml:"http://www.opengis.net/gml Array,omitempty"`
	Bag                               *BagType                               `xml:"http://www.opengis.net/gml Bag,omitempty"`
	BaseUnit                          *BaseUnitType                          `xml:"http://www.opengis.net/gml BaseUnit,omitempty"`
	CartesianCS                       *CartesianCSType                       `xml:"http://www.opengis.net/gml CartesianCS,omitempty"`
	CompositeCurve                    *CompositeCurveType                    `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	CompositeSolid                    *CompositeSolidType                    `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	CompositeSurface                  *CompositeSurfaceType                  `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	CompoundCRS                       *CompoundCRSType                       `xml:"http://www.opengis.net/gml CompoundCRS,omitempty"`
	ConcatenatedOperation             *ConcatenatedOperationType             `xml:"http://www.opengis.net/gml ConcatenatedOperation,omitempty"`
	ConventionalUnit                  *ConventionalUnitType                  `xml:"http://www.opengis.net/gml ConventionalUnit,omitempty"`
	Conversion                        *ConversionType                        `xml:"http://www.opengis.net/gml Conversion,omitempty"`
	CoordinateSystemAxis              *CoordinateSystemAxisType              `xml:"http://www.opengis.net/gml CoordinateSystemAxis,omitempty"`
	Curve                             *CurveType                             `xml:"http://www.opengis.net/gml Curve,omitempty"`
	CylindricalCS                     *CylindricalCSType                     `xml:"http://www.opengis.net/gml CylindricalCS,omitempty"`
	Definition                        *DefinitionType                        `xml:"http://www.opengis.net/gml Definition,omitempty"`
	DefinitionCollection              *DictionaryType                        `xml:"http://www.opengis.net/gml DefinitionCollection,omitempty"`
	DefinitionProxy                   *DefinitionProxyType                   `xml:"http://www.opengis.net/gml DefinitionProxy,omitempty"`
	DerivedCRS                        *DerivedCRSType                        `xml:"http://www.opengis.net/gml DerivedCRS,omitempty"`
	DerivedUnit                       *DerivedUnitType                       `xml:"http://www.opengis.net/gml DerivedUnit,omitempty"`
	Dictionary                        *DictionaryType                        `xml:"http://www.opengis.net/gml Dictionary,omitempty"`
	DirectedObservation               *DirectedObservationType               `xml:"http://www.opengis.net/gml DirectedObservation,omitempty"`
	DirectedObservationAtDistance     *DirectedObservationAtDistanceType     `xml:"http://www.opengis.net/gml DirectedObservationAtDistance,omitempty"`
	Edge                              *EdgeType                              `xml:"http://www.opengis.net/gml Edge,omitempty"`
	Ellipsoid                         *EllipsoidType                         `xml:"http://www.opengis.net/gml Ellipsoid,omitempty"`
	EllipsoidalCS                     *EllipsoidalCSType                     `xml:"http://www.opengis.net/gml EllipsoidalCS,omitempty"`
	EngineeringCRS                    *EngineeringCRSType                    `xml:"http://www.opengis.net/gml EngineeringCRS,omitempty"`
	EngineeringDatum                  *EngineeringDatumType                  `xml:"http://www.opengis.net/gml EngineeringDatum,omitempty"`
	Face                              *FaceType                              `xml:"http://www.opengis.net/gml Face,omitempty"`
	FeatureCollection                 *FeatureCollectionType                 `xml:"http://www.opengis.net/gml FeatureCollection,omitempty"`
	FeatureStyle                      *FeatureStyleType                      `xml:"http://www.opengis.net/gml FeatureStyle,omitempty"`
	GenericMetaData                   *GenericMetaDataType                   `xml:"http://www.opengis.net/gml GenericMetaData,omitempty"`
	GeocentricCRS                     *GeocentricCRSType                     `xml:"http://www.opengis.net/gml GeocentricCRS,omitempty"`
	GeodeticDatum                     *GeodeticDatumType                     `xml:"http://www.opengis.net/gml GeodeticDatum,omitempty"`
	GeographicCRS                     *GeographicCRSType                     `xml:"http://www.opengis.net/gml GeographicCRS,omitempty"`
	GeometricComplex                  *GeometricComplexType                  `xml:"http://www.opengis.net/gml GeometricComplex,omitempty"`
	GeometryStyle                     *GeometryStyleType                     `xml:"http://www.opengis.net/gml GeometryStyle,omitempty"`
	GraphStyle                        *GraphStyleType                        `xml:"http://www.opengis.net/gml GraphStyle,omitempty"`
	Grid                              *GridType                              `xml:"http://www.opengis.net/gml Grid,omitempty"`
	GridCoverage                      *GridCoverageType                      `xml:"http://www.opengis.net/gml GridCoverage,omitempty"`
	ImageCRS                          *ImageCRSType                          `xml:"http://www.opengis.net/gml ImageCRS,omitempty"`
	ImageDatum                        *ImageDatumType                        `xml:"http://www.opengis.net/gml ImageDatum,omitempty"`
	LabelStyle                        *LabelStyleType                        `xml:"http://www.opengis.net/gml LabelStyle,omitempty"`
	LineString                        *LineStringType                        `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearCS                          *LinearCSType                          `xml:"http://www.opengis.net/gml LinearCS,omitempty"`
	LinearRing                        *LinearRingType                        `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MovingObjectStatus                *MovingObjectStatusType                `xml:"http://www.opengis.net/gml MovingObjectStatus,omitempty"`
	MultiCurve                        *MultiCurveType                        `xml:"http://www.opengis.net/gml MultiCurve,omitempty"`
	MultiCurveCoverage                *MultiCurveCoverageType                `xml:"http://www.opengis.net/gml MultiCurveCoverage,omitempty"`
	MultiGeometry                     *MultiGeometryType                     `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString                   *MultiLineStringType                   `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                        *MultiPointType                        `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPointCoverage                *MultiPointCoverageType                `xml:"http://www.opengis.net/gml MultiPointCoverage,omitempty"`
	MultiPolygon                      *MultiPolygonType                      `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	MultiSolid                        *MultiSolidType                        `xml:"http://www.opengis.net/gml MultiSolid,omitempty"`
	MultiSolidCoverage                *MultiSolidCoverageType                `xml:"http://www.opengis.net/gml MultiSolidCoverage,omitempty"`
	MultiSurface                      *MultiSurfaceType                      `xml:"http://www.opengis.net/gml MultiSurface,omitempty"`
	MultiSurfaceCoverage              *MultiSurfaceCoverageType              `xml:"http://www.opengis.net/gml MultiSurfaceCoverage,omitempty"`
	Node                              *NodeType                              `xml:"http://www.opengis.net/gml Node,omitempty"`
	ObliqueCartesianCS                *ObliqueCartesianCSType                `xml:"http://www.opengis.net/gml ObliqueCartesianCS,omitempty"`
	Observation                       *ObservationType                       `xml:"http://www.opengis.net/gml Observation,omitempty"`
	OperationMethod                   *OperationMethodType                   `xml:"http://www.opengis.net/gml OperationMethod,omitempty"`
	OperationParameter                *OperationParameterType                `xml:"http://www.opengis.net/gml OperationParameter,omitempty"`
	OperationParameterGroup           *OperationParameterGroupType           `xml:"http://www.opengis.net/gml OperationParameterGroup,omitempty"`
	OrientableCurve                   *OrientableCurveType                   `xml:"http://www.opengis.net/gml OrientableCurve,omitempty"`
	OrientableSurface                 *OrientableSurfaceType                 `xml:"http://www.opengis.net/gml OrientableSurface,omitempty"`
	PassThroughOperation              *PassThroughOperationType              `xml:"http://www.opengis.net/gml PassThroughOperation,omitempty"`
	Point                             *PointType                             `xml:"http://www.opengis.net/gml Point,omitempty"`
	PolarCS                           *PolarCSType                           `xml:"http://www.opengis.net/gml PolarCS,omitempty"`
	Polygon                           *PolygonType                           `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	PolyhedralSurface                 *PolyhedralSurfaceType                 `xml:"http://www.opengis.net/gml PolyhedralSurface,omitempty"`
	PrimeMeridian                     *PrimeMeridianType                     `xml:"http://www.opengis.net/gml PrimeMeridian,omitempty"`
	ProjectedCRS                      *ProjectedCRSType                      `xml:"http://www.opengis.net/gml ProjectedCRS,omitempty"`
	RectifiedGrid                     *RectifiedGridType                     `xml:"http://www.opengis.net/gml RectifiedGrid,omitempty"`
	RectifiedGridCoverage             *RectifiedGridCoverageType             `xml:"http://www.opengis.net/gml RectifiedGridCoverage,omitempty"`
	Ring                              *RingType                              `xml:"http://www.opengis.net/gml Ring,omitempty"`
	Solid                             *SolidType                             `xml:"http://www.opengis.net/gml Solid,omitempty"`
	SphericalCS                       *SphericalCSType                       `xml:"http://www.opengis.net/gml SphericalCS,omitempty"`
	Style                             *StyleType                             `xml:"http://www.opengis.net/gml Style,omitempty"`
	Surface                           *SurfaceType                           `xml:"http://www.opengis.net/gml Surface,omitempty"`
	TemporalCRS                       *TemporalCRSType                       `xml:"http://www.opengis.net/gml TemporalCRS,omitempty"`
	TemporalCS                        *TemporalCSType                        `xml:"http://www.opengis.net/gml TemporalCS,omitempty"`
	TemporalDatum                     *TemporalDatumType                     `xml:"http://www.opengis.net/gml TemporalDatum,omitempty"`
	TimeCalendar                      *TimeCalendarType                      `xml:"http://www.opengis.net/gml TimeCalendar,omitempty"`
	TimeCalendarEra                   *TimeCalendarEraType                   `xml:"http://www.opengis.net/gml TimeCalendarEra,omitempty"`
	TimeClock                         *TimeClockType                         `xml:"http://www.opengis.net/gml TimeClock,omitempty"`
	TimeCoordinateSystem              *TimeCoordinateSystemType              `xml:"http://www.opengis.net/gml TimeCoordinateSystem,omitempty"`
	TimeEdge                          *TimeEdgeType                          `xml:"http://www.opengis.net/gml TimeEdge,omitempty"`
	TimeInstant                       *TimeInstantType                       `xml:"http://www.opengis.net/gml TimeInstant,omitempty"`
	TimeNode                          *TimeNodeType                          `xml:"http://www.opengis.net/gml TimeNode,omitempty"`
	TimeOrdinalReferenceSystem        *TimeOrdinalReferenceSystemType        `xml:"http://www.opengis.net/gml TimeOrdinalReferenceSystem,omitempty"`
	TimePeriod                        *TimePeriodType                        `xml:"http://www.opengis.net/gml TimePeriod,omitempty"`
	TimeTopologyComplex               *TimeTopologyComplexType               `xml:"http://www.opengis.net/gml TimeTopologyComplex,omitempty"`
	Tin                               *TinType                               `xml:"http://www.opengis.net/gml Tin,omitempty"`
	TopoComplex                       *TopoComplexType                       `xml:"http://www.opengis.net/gml TopoComplex,omitempty"`
	TopoSolid                         *TopoSolidType                         `xml:"http://www.opengis.net/gml TopoSolid,omitempty"`
	TopologyStyle                     *TopologyStyleType                     `xml:"http://www.opengis.net/gml TopologyStyle,omitempty"`
	Transformation                    *TransformationType                    `xml:"http://www.opengis.net/gml Transformation,omitempty"`
	TriangulatedSurface               *TriangulatedSurfaceType               `xml:"http://www.opengis.net/gml TriangulatedSurface,omitempty"`
	UnitDefinition                    *UnitDefinitionType                    `xml:"http://www.opengis.net/gml UnitDefinition,omitempty"`
	UserDefinedCS                     *UserDefinedCSType                     `xml:"http://www.opengis.net/gml UserDefinedCS,omitempty"`
	VerticalCRS                       *VerticalCRSType                       `xml:"http://www.opengis.net/gml VerticalCRS,omitempty"`
	VerticalCS                        *VerticalCSType                        `xml:"http://www.opengis.net/gml VerticalCS,omitempty"`
	VerticalDatum                     *VerticalDatumType                     `xml:"http://www.opengis.net/gml VerticalDatum,omitempty"`
	AbstractCRS                       *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _CRS,omitempty"`
	AbstractContinuousCoverage        *AbstractContinuousCoverageType        `xml:"http://www.opengis.net/gml _ContinuousCoverage,omitempty"`
	AbstractCoordinateOperation       *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _CoordinateOperation,omitempty"`
	AbstractCoordinateReferenceSystem *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _CoordinateReferenceSystem,omitempty"`
	AbstractCoordinateSystem          *AbstractCoordinateSystemType          `xml:"http://www.opengis.net/gml _CoordinateSystem,omitempty"`
	AbstractCoverage                  *AbstractCoverageType                  `xml:"http://www.opengis.net/gml _Coverage,omitempty"`
	AbstractCurve                     *AbstractCurveType                     `xml:"http://www.opengis.net/gml _Curve,omitempty"`
	AbstractDatum                     *AbstractDatumType                     `xml:"http://www.opengis.net/gml _Datum,omitempty"`
	AbstractDiscreteCoverage          *AbstractDiscreteCoverageType          `xml:"http://www.opengis.net/gml _DiscreteCoverage,omitempty"`
	AbstractFeature                   *AbstractFeatureType                   `xml:"http://www.opengis.net/gml _Feature,omitempty"`
	AbstractFeatureCollection         *AbstractFeatureCollectionType         `xml:"http://www.opengis.net/gml _FeatureCollection,omitempty"`
	AbstractGML                       *AbstractGMLType                       `xml:"http://www.opengis.net/gml _GML,omitempty"`
	AbstractGeneralConversion         *AbstractGeneralConversionType         `xml:"http://www.opengis.net/gml _GeneralConversion,omitempty"`
	AbstractGeneralDerivedCRS         *AbstractGeneralDerivedCRSType         `xml:"http://www.opengis.net/gml _GeneralDerivedCRS,omitempty"`
	AbstractGeneralOperationParameter *AbstractGeneralOperationParameterType `xml:"http://www.opengis.net/gml _GeneralOperationParameter,omitempty"`
	AbstractGeneralTransformation     *AbstractGeneralTransformationType     `xml:"http://www.opengis.net/gml _GeneralTransformation,omitempty"`
	AbstractGeometricAggregate        *AbstractGeometricAggregateType        `xml:"http://www.opengis.net/gml _GeometricAggregate,omitempty"`
	AbstractGeometricPrimitive        *AbstractGeometricPrimitiveType        `xml:"http://www.opengis.net/gml _GeometricPrimitive,omitempty"`
	AbstractGeometry                  *AbstractGeometryType                  `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	AbstractImplicitGeometry          *AbstractGeometryType                  `xml:"http://www.opengis.net/gml _ImplicitGeometry,omitempty"`
	AbstractMetaData                  *AbstractMetaDataType                  `xml:"http://www.opengis.net/gml _MetaData,omitempty"`
	AbstractOperation                 *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _Operation,omitempty"`
	AbstractReferenceSystem           *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _ReferenceSystem,omitempty"`
	AbstractRing                      *AbstractRingType                      `xml:"http://www.opengis.net/gml _Ring,omitempty"`
	AbstractSingleOperation           *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _SingleOperation,omitempty"`
	AbstractSolid                     *AbstractSolidType                     `xml:"http://www.opengis.net/gml _Solid,omitempty"`
	AbstractStyle                     *AbstractStyleType                     `xml:"http://www.opengis.net/gml _Style,omitempty"`
	AbstractSurface                   *AbstractSurfaceType                   `xml:"http://www.opengis.net/gml _Surface,omitempty"`
	AbstractTimeComplex               *AbstractTimeComplexType               `xml:"http://www.opengis.net/gml _TimeComplex,omitempty"`
	AbstractTimeGeometricPrimitive    *AbstractTimeGeometricPrimitiveType    `xml:"http://www.opengis.net/gml _TimeGeometricPrimitive,omitempty"`
	AbstractTimeObject                *AbstractTimeObjectType                `xml:"http://www.opengis.net/gml _TimeObject,omitempty"`
	AbstractTimePrimitive             *AbstractTimePrimitiveType             `xml:"http://www.opengis.net/gml _TimePrimitive,omitempty"`
	AbstractTimeReferenceSystem       *AbstractTimeReferenceSystemType       `xml:"http://www.opengis.net/gml _TimeReferenceSystem,omitempty"`
	AbstractTimeSlice                 *AbstractTimeSliceType                 `xml:"http://www.opengis.net/gml _TimeSlice,omitempty"`
	AbstractTimeTopologyPrimitive     *AbstractTimeTopologyPrimitiveType     `xml:"http://www.opengis.net/gml _TimeTopologyPrimitive,omitempty"`
	AbstractTopoPrimitive             *AbstractTopoPrimitiveType             `xml:"http://www.opengis.net/gml _TopoPrimitive,omitempty"`
	AbstractTopology                  *AbstractTopologyType                  `xml:"http://www.opengis.net/gml _Topology,omitempty"`
	Null                              *string                                `xml:"http://www.opengis.net/gml Null,omitempty"`
	// Aggregate value built using the Composite pattern.
	CompositeValue *CompositeValueType `xml:"http://www.opengis.net/gml CompositeValue,omitempty"`
	ValueArray     *ValueArrayType     `xml:"http://www.opengis.net/gml ValueArray,omitempty"`
	// A value from two-valued logic, using the XML Schema boolean type.  An instance may take the values {true, false, 1, 0}.
	Boolean *bool `xml:"http://www.opengis.net/gml Boolean,omitempty"`
	// A term representing a classification.  It has an optional XML attribute codeSpace, whose value is a URI which identifies a dictionary, codelist or authority for the term.
	Category *CodeType `xml:"http://www.opengis.net/gml Category,omitempty"`
	// A numeric value with a scale.  The content of the element is an amount using the XML Schema type double which permits decimal or scientific notation.  An XML attribute uom (unit of measure) is required, whose value is a URI which identifies the definition of the scale or units by which the numeric value must be multiplied.
	Quantity *MeasureType `xml:"http://www.opengis.net/gml Quantity,omitempty"`
	// An integer representing a frequency of occurrence.
	Count *int `xml:"http://www.opengis.net/gml Count,omitempty"`
	// XML List based on XML Schema boolean type.  An element of this type contains a space-separated list of boolean values {0,1,true,false}
	BooleanList *string `xml:"http://www.opengis.net/gml BooleanList,omitempty"`
	// A space-separated list of terms or nulls.  A single XML attribute codeSpace may be provided, which authorises all the terms in the list.
	CategoryList *CodeOrNullListType `xml:"http://www.opengis.net/gml CategoryList,omitempty"`
	// A space separated list of amounts or nulls.  The amounts use the XML Schema type double.  A single XML attribute uom (unit of measure) is required, whose value is a URI which identifies the definition of the scale or units by which all the amounts in the list must be multiplied.
	QuantityList *MeasureOrNullListType `xml:"http://www.opengis.net/gml QuantityList,omitempty"`
	// A space-separated list of integers or nulls.
	CountList *string `xml:"http://www.opengis.net/gml CountList,omitempty"`
	// Utility element to store a 2-point range of ordinal values. If one member is a null, then this is a single ended interval.
	CategoryExtent *CategoryExtentType `xml:"http://www.opengis.net/gml CategoryExtent,omitempty"`
	// Utility element to store a 2-point range of numeric values. If one member is a null, then this is a single ended interval.
	QuantityExtent *QuantityExtentType `xml:"http://www.opengis.net/gml QuantityExtent,omitempty"`
	// Utility element to store a 2-point range of frequency values. If one member is a null, then this is a single ended interval.
	CountExtent *string `xml:"http://www.opengis.net/gml CountExtent,omitempty"`
}

type ValueArrayType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Element which refers to, or contains, a Value.  This version is used in CompositeValues.
	ValueComponent []ValuePropertyType `xml:"http://www.opengis.net/gml valueComponent"`
	// Element which refers to, or contains, a set of homogeneously typed Values.
	ValueComponents *ValueArrayPropertyType `xml:"http://www.opengis.net/gml valueComponents,omitempty"`
	CodeSpace       *string                 `xml:"codeSpace,attr,omitempty"`
	Uom             *string                 `xml:"uom,attr,omitempty"`
}

type ValuePropertyType struct {
	// This abstract element is the head of a substitutionGroup hierararchy which may contain either simpleContent or complexContent elements.  It is used to assert the model position of "class" elements declared in other GML schemas.
	AbstractObject                    *string                                `xml:"http://www.opengis.net/gml _Object,omitempty"`
	Array                             *ArrayType                             `xml:"http://www.opengis.net/gml Array,omitempty"`
	Bag                               *BagType                               `xml:"http://www.opengis.net/gml Bag,omitempty"`
	BaseUnit                          *BaseUnitType                          `xml:"http://www.opengis.net/gml BaseUnit,omitempty"`
	CartesianCS                       *CartesianCSType                       `xml:"http://www.opengis.net/gml CartesianCS,omitempty"`
	CompositeCurve                    *CompositeCurveType                    `xml:"http://www.opengis.net/gml CompositeCurve,omitempty"`
	CompositeSolid                    *CompositeSolidType                    `xml:"http://www.opengis.net/gml CompositeSolid,omitempty"`
	CompositeSurface                  *CompositeSurfaceType                  `xml:"http://www.opengis.net/gml CompositeSurface,omitempty"`
	CompoundCRS                       *CompoundCRSType                       `xml:"http://www.opengis.net/gml CompoundCRS,omitempty"`
	ConcatenatedOperation             *ConcatenatedOperationType             `xml:"http://www.opengis.net/gml ConcatenatedOperation,omitempty"`
	ConventionalUnit                  *ConventionalUnitType                  `xml:"http://www.opengis.net/gml ConventionalUnit,omitempty"`
	Conversion                        *ConversionType                        `xml:"http://www.opengis.net/gml Conversion,omitempty"`
	CoordinateSystemAxis              *CoordinateSystemAxisType              `xml:"http://www.opengis.net/gml CoordinateSystemAxis,omitempty"`
	Curve                             *CurveType                             `xml:"http://www.opengis.net/gml Curve,omitempty"`
	CylindricalCS                     *CylindricalCSType                     `xml:"http://www.opengis.net/gml CylindricalCS,omitempty"`
	Definition                        *DefinitionType                        `xml:"http://www.opengis.net/gml Definition,omitempty"`
	DefinitionCollection              *DictionaryType                        `xml:"http://www.opengis.net/gml DefinitionCollection,omitempty"`
	DefinitionProxy                   *DefinitionProxyType                   `xml:"http://www.opengis.net/gml DefinitionProxy,omitempty"`
	DerivedCRS                        *DerivedCRSType                        `xml:"http://www.opengis.net/gml DerivedCRS,omitempty"`
	DerivedUnit                       *DerivedUnitType                       `xml:"http://www.opengis.net/gml DerivedUnit,omitempty"`
	Dictionary                        *DictionaryType                        `xml:"http://www.opengis.net/gml Dictionary,omitempty"`
	DirectedObservation               *DirectedObservationType               `xml:"http://www.opengis.net/gml DirectedObservation,omitempty"`
	DirectedObservationAtDistance     *DirectedObservationAtDistanceType     `xml:"http://www.opengis.net/gml DirectedObservationAtDistance,omitempty"`
	Edge                              *EdgeType                              `xml:"http://www.opengis.net/gml Edge,omitempty"`
	Ellipsoid                         *EllipsoidType                         `xml:"http://www.opengis.net/gml Ellipsoid,omitempty"`
	EllipsoidalCS                     *EllipsoidalCSType                     `xml:"http://www.opengis.net/gml EllipsoidalCS,omitempty"`
	EngineeringCRS                    *EngineeringCRSType                    `xml:"http://www.opengis.net/gml EngineeringCRS,omitempty"`
	EngineeringDatum                  *EngineeringDatumType                  `xml:"http://www.opengis.net/gml EngineeringDatum,omitempty"`
	Face                              *FaceType                              `xml:"http://www.opengis.net/gml Face,omitempty"`
	FeatureCollection                 *FeatureCollectionType                 `xml:"http://www.opengis.net/gml FeatureCollection,omitempty"`
	FeatureStyle                      *FeatureStyleType                      `xml:"http://www.opengis.net/gml FeatureStyle,omitempty"`
	GenericMetaData                   *GenericMetaDataType                   `xml:"http://www.opengis.net/gml GenericMetaData,omitempty"`
	GeocentricCRS                     *GeocentricCRSType                     `xml:"http://www.opengis.net/gml GeocentricCRS,omitempty"`
	GeodeticDatum                     *GeodeticDatumType                     `xml:"http://www.opengis.net/gml GeodeticDatum,omitempty"`
	GeographicCRS                     *GeographicCRSType                     `xml:"http://www.opengis.net/gml GeographicCRS,omitempty"`
	GeometricComplex                  *GeometricComplexType                  `xml:"http://www.opengis.net/gml GeometricComplex,omitempty"`
	GeometryStyle                     *GeometryStyleType                     `xml:"http://www.opengis.net/gml GeometryStyle,omitempty"`
	GraphStyle                        *GraphStyleType                        `xml:"http://www.opengis.net/gml GraphStyle,omitempty"`
	Grid                              *GridType                              `xml:"http://www.opengis.net/gml Grid,omitempty"`
	GridCoverage                      *GridCoverageType                      `xml:"http://www.opengis.net/gml GridCoverage,omitempty"`
	ImageCRS                          *ImageCRSType                          `xml:"http://www.opengis.net/gml ImageCRS,omitempty"`
	ImageDatum                        *ImageDatumType                        `xml:"http://www.opengis.net/gml ImageDatum,omitempty"`
	LabelStyle                        *LabelStyleType                        `xml:"http://www.opengis.net/gml LabelStyle,omitempty"`
	LineString                        *LineStringType                        `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearCS                          *LinearCSType                          `xml:"http://www.opengis.net/gml LinearCS,omitempty"`
	LinearRing                        *LinearRingType                        `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MovingObjectStatus                *MovingObjectStatusType                `xml:"http://www.opengis.net/gml MovingObjectStatus,omitempty"`
	MultiCurve                        *MultiCurveType                        `xml:"http://www.opengis.net/gml MultiCurve,omitempty"`
	MultiCurveCoverage                *MultiCurveCoverageType                `xml:"http://www.opengis.net/gml MultiCurveCoverage,omitempty"`
	MultiGeometry                     *MultiGeometryType                     `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString                   *MultiLineStringType                   `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                        *MultiPointType                        `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPointCoverage                *MultiPointCoverageType                `xml:"http://www.opengis.net/gml MultiPointCoverage,omitempty"`
	MultiPolygon                      *MultiPolygonType                      `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	MultiSolid                        *MultiSolidType                        `xml:"http://www.opengis.net/gml MultiSolid,omitempty"`
	MultiSolidCoverage                *MultiSolidCoverageType                `xml:"http://www.opengis.net/gml MultiSolidCoverage,omitempty"`
	MultiSurface                      *MultiSurfaceType                      `xml:"http://www.opengis.net/gml MultiSurface,omitempty"`
	MultiSurfaceCoverage              *MultiSurfaceCoverageType              `xml:"http://www.opengis.net/gml MultiSurfaceCoverage,omitempty"`
	Node                              *NodeType                              `xml:"http://www.opengis.net/gml Node,omitempty"`
	ObliqueCartesianCS                *ObliqueCartesianCSType                `xml:"http://www.opengis.net/gml ObliqueCartesianCS,omitempty"`
	Observation                       *ObservationType                       `xml:"http://www.opengis.net/gml Observation,omitempty"`
	OperationMethod                   *OperationMethodType                   `xml:"http://www.opengis.net/gml OperationMethod,omitempty"`
	OperationParameter                *OperationParameterType                `xml:"http://www.opengis.net/gml OperationParameter,omitempty"`
	OperationParameterGroup           *OperationParameterGroupType           `xml:"http://www.opengis.net/gml OperationParameterGroup,omitempty"`
	OrientableCurve                   *OrientableCurveType                   `xml:"http://www.opengis.net/gml OrientableCurve,omitempty"`
	OrientableSurface                 *OrientableSurfaceType                 `xml:"http://www.opengis.net/gml OrientableSurface,omitempty"`
	PassThroughOperation              *PassThroughOperationType              `xml:"http://www.opengis.net/gml PassThroughOperation,omitempty"`
	Point                             *PointType                             `xml:"http://www.opengis.net/gml Point,omitempty"`
	PolarCS                           *PolarCSType                           `xml:"http://www.opengis.net/gml PolarCS,omitempty"`
	Polygon                           *PolygonType                           `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	PolyhedralSurface                 *PolyhedralSurfaceType                 `xml:"http://www.opengis.net/gml PolyhedralSurface,omitempty"`
	PrimeMeridian                     *PrimeMeridianType                     `xml:"http://www.opengis.net/gml PrimeMeridian,omitempty"`
	ProjectedCRS                      *ProjectedCRSType                      `xml:"http://www.opengis.net/gml ProjectedCRS,omitempty"`
	RectifiedGrid                     *RectifiedGridType                     `xml:"http://www.opengis.net/gml RectifiedGrid,omitempty"`
	RectifiedGridCoverage             *RectifiedGridCoverageType             `xml:"http://www.opengis.net/gml RectifiedGridCoverage,omitempty"`
	Ring                              *RingType                              `xml:"http://www.opengis.net/gml Ring,omitempty"`
	Solid                             *SolidType                             `xml:"http://www.opengis.net/gml Solid,omitempty"`
	SphericalCS                       *SphericalCSType                       `xml:"http://www.opengis.net/gml SphericalCS,omitempty"`
	Style                             *StyleType                             `xml:"http://www.opengis.net/gml Style,omitempty"`
	Surface                           *SurfaceType                           `xml:"http://www.opengis.net/gml Surface,omitempty"`
	TemporalCRS                       *TemporalCRSType                       `xml:"http://www.opengis.net/gml TemporalCRS,omitempty"`
	TemporalCS                        *TemporalCSType                        `xml:"http://www.opengis.net/gml TemporalCS,omitempty"`
	TemporalDatum                     *TemporalDatumType                     `xml:"http://www.opengis.net/gml TemporalDatum,omitempty"`
	TimeCalendar                      *TimeCalendarType                      `xml:"http://www.opengis.net/gml TimeCalendar,omitempty"`
	TimeCalendarEra                   *TimeCalendarEraType                   `xml:"http://www.opengis.net/gml TimeCalendarEra,omitempty"`
	TimeClock                         *TimeClockType                         `xml:"http://www.opengis.net/gml TimeClock,omitempty"`
	TimeCoordinateSystem              *TimeCoordinateSystemType              `xml:"http://www.opengis.net/gml TimeCoordinateSystem,omitempty"`
	TimeEdge                          *TimeEdgeType                          `xml:"http://www.opengis.net/gml TimeEdge,omitempty"`
	TimeInstant                       *TimeInstantType                       `xml:"http://www.opengis.net/gml TimeInstant,omitempty"`
	TimeNode                          *TimeNodeType                          `xml:"http://www.opengis.net/gml TimeNode,omitempty"`
	TimeOrdinalReferenceSystem        *TimeOrdinalReferenceSystemType        `xml:"http://www.opengis.net/gml TimeOrdinalReferenceSystem,omitempty"`
	TimePeriod                        *TimePeriodType                        `xml:"http://www.opengis.net/gml TimePeriod,omitempty"`
	TimeTopologyComplex               *TimeTopologyComplexType               `xml:"http://www.opengis.net/gml TimeTopologyComplex,omitempty"`
	Tin                               *TinType                               `xml:"http://www.opengis.net/gml Tin,omitempty"`
	TopoComplex                       *TopoComplexType                       `xml:"http://www.opengis.net/gml TopoComplex,omitempty"`
	TopoSolid                         *TopoSolidType                         `xml:"http://www.opengis.net/gml TopoSolid,omitempty"`
	TopologyStyle                     *TopologyStyleType                     `xml:"http://www.opengis.net/gml TopologyStyle,omitempty"`
	Transformation                    *TransformationType                    `xml:"http://www.opengis.net/gml Transformation,omitempty"`
	TriangulatedSurface               *TriangulatedSurfaceType               `xml:"http://www.opengis.net/gml TriangulatedSurface,omitempty"`
	UnitDefinition                    *UnitDefinitionType                    `xml:"http://www.opengis.net/gml UnitDefinition,omitempty"`
	UserDefinedCS                     *UserDefinedCSType                     `xml:"http://www.opengis.net/gml UserDefinedCS,omitempty"`
	VerticalCRS                       *VerticalCRSType                       `xml:"http://www.opengis.net/gml VerticalCRS,omitempty"`
	VerticalCS                        *VerticalCSType                        `xml:"http://www.opengis.net/gml VerticalCS,omitempty"`
	VerticalDatum                     *VerticalDatumType                     `xml:"http://www.opengis.net/gml VerticalDatum,omitempty"`
	AbstractCRS                       *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _CRS,omitempty"`
	AbstractContinuousCoverage        *AbstractContinuousCoverageType        `xml:"http://www.opengis.net/gml _ContinuousCoverage,omitempty"`
	AbstractCoordinateOperation       *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _CoordinateOperation,omitempty"`
	AbstractCoordinateReferenceSystem *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _CoordinateReferenceSystem,omitempty"`
	AbstractCoordinateSystem          *AbstractCoordinateSystemType          `xml:"http://www.opengis.net/gml _CoordinateSystem,omitempty"`
	AbstractCoverage                  *AbstractCoverageType                  `xml:"http://www.opengis.net/gml _Coverage,omitempty"`
	AbstractCurve                     *AbstractCurveType                     `xml:"http://www.opengis.net/gml _Curve,omitempty"`
	AbstractDatum                     *AbstractDatumType                     `xml:"http://www.opengis.net/gml _Datum,omitempty"`
	AbstractDiscreteCoverage          *AbstractDiscreteCoverageType          `xml:"http://www.opengis.net/gml _DiscreteCoverage,omitempty"`
	AbstractFeature                   *AbstractFeatureType                   `xml:"http://www.opengis.net/gml _Feature,omitempty"`
	AbstractFeatureCollection         *AbstractFeatureCollectionType         `xml:"http://www.opengis.net/gml _FeatureCollection,omitempty"`
	AbstractGML                       *AbstractGMLType                       `xml:"http://www.opengis.net/gml _GML,omitempty"`
	AbstractGeneralConversion         *AbstractGeneralConversionType         `xml:"http://www.opengis.net/gml _GeneralConversion,omitempty"`
	AbstractGeneralDerivedCRS         *AbstractGeneralDerivedCRSType         `xml:"http://www.opengis.net/gml _GeneralDerivedCRS,omitempty"`
	AbstractGeneralOperationParameter *AbstractGeneralOperationParameterType `xml:"http://www.opengis.net/gml _GeneralOperationParameter,omitempty"`
	AbstractGeneralTransformation     *AbstractGeneralTransformationType     `xml:"http://www.opengis.net/gml _GeneralTransformation,omitempty"`
	AbstractGeometricAggregate        *AbstractGeometricAggregateType        `xml:"http://www.opengis.net/gml _GeometricAggregate,omitempty"`
	AbstractGeometricPrimitive        *AbstractGeometricPrimitiveType        `xml:"http://www.opengis.net/gml _GeometricPrimitive,omitempty"`
	AbstractGeometry                  *AbstractGeometryType                  `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	AbstractImplicitGeometry          *AbstractGeometryType                  `xml:"http://www.opengis.net/gml _ImplicitGeometry,omitempty"`
	AbstractMetaData                  *AbstractMetaDataType                  `xml:"http://www.opengis.net/gml _MetaData,omitempty"`
	AbstractOperation                 *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _Operation,omitempty"`
	AbstractReferenceSystem           *AbstractReferenceSystemType           `xml:"http://www.opengis.net/gml _ReferenceSystem,omitempty"`
	AbstractRing                      *AbstractRingType                      `xml:"http://www.opengis.net/gml _Ring,omitempty"`
	AbstractSingleOperation           *AbstractCoordinateOperationType       `xml:"http://www.opengis.net/gml _SingleOperation,omitempty"`
	AbstractSolid                     *AbstractSolidType                     `xml:"http://www.opengis.net/gml _Solid,omitempty"`
	AbstractStyle                     *AbstractStyleType                     `xml:"http://www.opengis.net/gml _Style,omitempty"`
	AbstractSurface                   *AbstractSurfaceType                   `xml:"http://www.opengis.net/gml _Surface,omitempty"`
	AbstractTimeComplex               *AbstractTimeComplexType               `xml:"http://www.opengis.net/gml _TimeComplex,omitempty"`
	AbstractTimeGeometricPrimitive    *AbstractTimeGeometricPrimitiveType    `xml:"http://www.opengis.net/gml _TimeGeometricPrimitive,omitempty"`
	AbstractTimeObject                *AbstractTimeObjectType                `xml:"http://www.opengis.net/gml _TimeObject,omitempty"`
	AbstractTimePrimitive             *AbstractTimePrimitiveType             `xml:"http://www.opengis.net/gml _TimePrimitive,omitempty"`
	AbstractTimeReferenceSystem       *AbstractTimeReferenceSystemType       `xml:"http://www.opengis.net/gml _TimeReferenceSystem,omitempty"`
	AbstractTimeSlice                 *AbstractTimeSliceType                 `xml:"http://www.opengis.net/gml _TimeSlice,omitempty"`
	AbstractTimeTopologyPrimitive     *AbstractTimeTopologyPrimitiveType     `xml:"http://www.opengis.net/gml _TimeTopologyPrimitive,omitempty"`
	AbstractTopoPrimitive             *AbstractTopoPrimitiveType             `xml:"http://www.opengis.net/gml _TopoPrimitive,omitempty"`
	AbstractTopology                  *AbstractTopologyType                  `xml:"http://www.opengis.net/gml _Topology,omitempty"`
	Null                              *string                                `xml:"http://www.opengis.net/gml Null,omitempty"`
	// Aggregate value built using the Composite pattern.
	CompositeValue *CompositeValueType `xml:"http://www.opengis.net/gml CompositeValue,omitempty"`
	ValueArray     *ValueArrayType     `xml:"http://www.opengis.net/gml ValueArray,omitempty"`
	// A value from two-valued logic, using the XML Schema boolean type.  An instance may take the values {true, false, 1, 0}.
	Boolean *bool `xml:"http://www.opengis.net/gml Boolean,omitempty"`
	// A term representing a classification.  It has an optional XML attribute codeSpace, whose value is a URI which identifies a dictionary, codelist or authority for the term.
	Category *CodeType `xml:"http://www.opengis.net/gml Category,omitempty"`
	// A numeric value with a scale.  The content of the element is an amount using the XML Schema type double which permits decimal or scientific notation.  An XML attribute uom (unit of measure) is required, whose value is a URI which identifies the definition of the scale or units by which the numeric value must be multiplied.
	Quantity *MeasureType `xml:"http://www.opengis.net/gml Quantity,omitempty"`
	// An integer representing a frequency of occurrence.
	Count *int `xml:"http://www.opengis.net/gml Count,omitempty"`
	// XML List based on XML Schema boolean type.  An element of this type contains a space-separated list of boolean values {0,1,true,false}
	BooleanList *string `xml:"http://www.opengis.net/gml BooleanList,omitempty"`
	// A space-separated list of terms or nulls.  A single XML attribute codeSpace may be provided, which authorises all the terms in the list.
	CategoryList *CodeOrNullListType `xml:"http://www.opengis.net/gml CategoryList,omitempty"`
	// A space separated list of amounts or nulls.  The amounts use the XML Schema type double.  A single XML attribute uom (unit of measure) is required, whose value is a URI which identifies the definition of the scale or units by which all the amounts in the list must be multiplied.
	QuantityList *MeasureOrNullListType `xml:"http://www.opengis.net/gml QuantityList,omitempty"`
	// A space-separated list of integers or nulls.
	CountList *string `xml:"http://www.opengis.net/gml CountList,omitempty"`
	// Utility element to store a 2-point range of ordinal values. If one member is a null, then this is a single ended interval.
	CategoryExtent *CategoryExtentType `xml:"http://www.opengis.net/gml CategoryExtent,omitempty"`
	// Utility element to store a 2-point range of numeric values. If one member is a null, then this is a single ended interval.
	QuantityExtent *QuantityExtentType `xml:"http://www.opengis.net/gml QuantityExtent,omitempty"`
	// Utility element to store a 2-point range of frequency values. If one member is a null, then this is a single ended interval.
	CountExtent  *string `xml:"http://www.opengis.net/gml CountExtent,omitempty"`
	RemoteSchema string  `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string  `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string  `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string  `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string  `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string  `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string  `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string  `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type VectorType struct {
	Value string `xml:",chardata"`
	// In general this reference points to a CRS instance of gml:CoordinateReferenceSystemType
	// (see coordinateReferenceSystems.xsd). For well known references it is not required that the CRS description exists at the
	// location the URI points to. If no srsName attribute is given, the CRS must be specified as part of the larger context this
	// geometry element is part of, e.g. a geometric element like point, curve, etc. It is expected that this attribute will be specified
	// at the direct position level only in rare cases.
	SrsName *string `xml:"srsName,attr,omitempty"`
	// The "srsDimension" is the length of coordinate sequence (the number of entries in the list). This dimension is
	// specified by the coordinate reference system. When the srsName attribute is omitted, this attribute shall be omitted.
	SrsDimension *int `xml:"srsDimension,attr,omitempty"`
	// Ordered list of labels for all the axes of this CRS. The gml:axisAbbrev value should be used for these axis
	// labels, after spaces and forbiddden characters are removed. When the srsName attribute is included, this attribute is optional.
	// When the srsName attribute is omitted, this attribute shall also be omitted.
	AxisLabels *string `xml:"axisLabels,attr,omitempty"`
	// Ordered list of unit of measure (uom) labels for all the axes of this CRS. The value of the string in the
	// gml:catalogSymbol should be used for this uom labels, after spaces and forbiddden characters are removed. When the
	// axisLabels attribute is included, this attribute shall also be included. When the axisLabels attribute is omitted, this attribute
	// shall also be omitted.
	UomLabels *string `xml:"uomLabels,attr,omitempty"`
}

type VerticalCRSRefType struct {
	VerticalCRS  *VerticalCRSType `xml:"http://www.opengis.net/gml VerticalCRS,omitempty"`
	RemoteSchema string           `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string           `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string           `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string           `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string           `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string           `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string           `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string           `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type VerticalCRSType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alterative identifications of this reference system. The first srsID, if any, is normally the primary identification code, and any others are aliases.
	SrsID []IdentifierType `xml:"http://www.opengis.net/gml srsID"`
	// Comments on or information about this reference system, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope *string `xml:"http://www.opengis.net/gml scope,omitempty"`
	// Association to the vertical coordinate system used by this CRS.
	UsesVerticalCS *VerticalCSRefType `xml:"http://www.opengis.net/gml usesVerticalCS,omitempty"`
	// Association to the vertical datum used by this CRS.
	UsesVerticalDatum *VerticalDatumRefType `xml:"http://www.opengis.net/gml usesVerticalDatum,omitempty"`
}

type VerticalCSRefType struct {
	VerticalCS   *VerticalCSType `xml:"http://www.opengis.net/gml VerticalCS,omitempty"`
	RemoteSchema string          `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string          `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string          `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string          `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string          `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string          `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string          `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string          `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type VerticalCSType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this coordinate system. The first csID, if any, is normally the primary identification code, and any others are aliases.
	CsID []IdentifierType `xml:"http://www.opengis.net/gml csID"`
	// Comments on or information about this coordinate system, including data source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Ordered sequence of associations to the coordinate system axes included in this coordinate system.
	UsesAxis []CoordinateSystemAxisRefType `xml:"http://www.opengis.net/gml usesAxis"`
}

type VerticalDatumRefType struct {
	VerticalDatum *VerticalDatumType `xml:"http://www.opengis.net/gml VerticalDatum,omitempty"`
	RemoteSchema  string             `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField     string             `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href          string             `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role          string             `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole       string             `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title         string             `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show          string             `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate       string             `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type VerticalDatumType struct {
	// Contains or refers to a metadata package that contains metadata properties.
	MetaDataProperty []MetaDataPropertyType `xml:"http://www.opengis.net/gml metaDataProperty"`
	// Contains a simple text description of the object, or refers to an external description.
	Description *StringOrRefType `xml:"http://www.opengis.net/gml description,omitempty"`
	// Multiple names may be provided.  These will often be distinguished by being assigned by different authorities, as indicated by the value of the codeSpace attribute.  In an instance document there will usually only be one name per authority.
	Name                    []CodeType `xml:"http://www.opengis.net/gml name"`
	CoordinateOperationName []CodeType `xml:"http://www.opengis.net/gml coordinateOperationName"`
	CsName                  []CodeType `xml:"http://www.opengis.net/gml csName"`
	DatumName               []CodeType `xml:"http://www.opengis.net/gml datumName"`
	EllipsoidName           []CodeType `xml:"http://www.opengis.net/gml ellipsoidName"`
	GroupName               []CodeType `xml:"http://www.opengis.net/gml groupName"`
	MeridianName            []CodeType `xml:"http://www.opengis.net/gml meridianName"`
	MethodName              []CodeType `xml:"http://www.opengis.net/gml methodName"`
	ParameterName           []CodeType `xml:"http://www.opengis.net/gml parameterName"`
	SrsName                 []CodeType `xml:"http://www.opengis.net/gml srsName"`
	Id                      string     `xml:"http://www.opengis.net/gml id,attr,omitempty"`
	// Set of alternative identifications of this datum. The first datumID, if any, is normally the primary identification code, and any others are aliases.
	DatumID []IdentifierType `xml:"http://www.opengis.net/gml datumID"`
	// Comments on this reference system, including source information.
	Remarks *StringOrRefType `xml:"http://www.opengis.net/gml remarks,omitempty"`
	// Description, possibly including coordinates, of the point or points used to anchor the datum to the Earth. Also known as the "origin", especially for engineering and image datums. The codeSpace attribute can be used to reference a source of more detailed on this point or surface, or on a set of such descriptions.
	// - For a geodetic datum, this point is also known as the fundamental point, which is traditionally the point where the relationship between geoid and ellipsoid is defined. In some cases, the "fundamental point" may consist of a number of points. In those cases, the parameters defining the geoid/ellipsoid relationship have been averaged for these points, and the averages adopted as the datum definition.
	// - For an engineering datum, the anchor point may be a physical point, or it may be a point with defined coordinates in another CRS. When appropriate, the coordinates of this anchor point can be referenced in another document, such as referencing a GML feature that references or includes a point position.
	// - For an image datum, the anchor point is usually either the centre of the image or the corner of the image.
	// - For a temporal datum, this attribute is not defined. Instead of the anchor point, a temporal datum carries a separate time origin of type DateTime.
	AnchorPoint *CodeType `xml:"http://www.opengis.net/gml anchorPoint,omitempty"`
	// The time after which this datum definition is valid. This time may be precise (e.g. 1997.0 for IRTF97) or merely a year (e.g. 1983 for NAD83). In the latter case, the epoch usually refers to the year in which a major recalculation of the geodetic control network, underlying the datum, was executed or initiated. An old datum can remain valid after a new datum is defined. Alternatively, a datum may be superseded by a later datum, in which case the realization epoch for the new datum defines the upper limit for the validity of the superseded datum.
	RealizationEpoch *string `xml:"http://www.opengis.net/gml realizationEpoch,omitempty"`
	// Area or region in which this CRS object is valid.
	ValidArea *ExtentType `xml:"http://www.opengis.net/gml validArea,omitempty"`
	// Description of domain of usage, or limitations of usage, for which this CRS object is valid.
	Scope             *string                `xml:"http://www.opengis.net/gml scope,omitempty"`
	VerticalDatumType *VerticalDatumTypeType `xml:"http://www.opengis.net/gml verticalDatumType,omitempty"`
}

type VerticalDatumTypeType struct {
	Value string `xml:",chardata"`
	// Reference to a source of information specifying the values and meanings of all the allowed string values for this VerticalDatumTypeType.
	CodeSpace string `xml:"codeSpace,attr"`
}

type VolumeType struct {
	Value string `xml:",chardata"`
}
