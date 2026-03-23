package generated

type AbstractFeatureCollectionBaseType struct {
	Description *string            `xml:"http://www.opengis.net/gml description,omitempty"`
	Name        *string            `xml:"http://www.opengis.net/gml name,omitempty"`
	BoundedBy   *BoundingShapeType `xml:"http://www.opengis.net/gml boundedBy,omitempty"`
	Fid         *string            `xml:"fid,attr,omitempty"`
}

type AbstractFeatureCollectionType struct {
	Description   *string                  `xml:"http://www.opengis.net/gml description,omitempty"`
	Name          *string                  `xml:"http://www.opengis.net/gml name,omitempty"`
	BoundedBy     *BoundingShapeType       `xml:"http://www.opengis.net/gml boundedBy,omitempty"`
	Fid           *string                  `xml:"fid,attr,omitempty"`
	FeatureMember []FeatureAssociationType `xml:"http://www.opengis.net/gml featureMember"`
}

type AbstractFeatureType struct {
	Description *string            `xml:"http://www.opengis.net/gml description,omitempty"`
	Name        *string            `xml:"http://www.opengis.net/gml name,omitempty"`
	BoundedBy   *BoundingShapeType `xml:"http://www.opengis.net/gml boundedBy,omitempty"`
	Fid         *string            `xml:"fid,attr,omitempty"`
}

type AbstractGeometryCollectionBaseType struct {
	Value   string  `xml:",chardata"`
	Gid     *string `xml:"gid,attr,omitempty"`
	SrsName *string `xml:"srsName,attr,omitempty"`
}

type AbstractGeometryType struct {
	Value   string  `xml:",chardata"`
	Gid     *string `xml:"gid,attr,omitempty"`
	SrsName *string `xml:"srsName,attr,omitempty"`
}

type BoundingShapeType struct {
	Box  *BoxType `xml:"http://www.opengis.net/gml Box,omitempty"`
	Null string   `xml:"http://www.opengis.net/gml null,omitempty"`
}

type BoxType struct {
	Value       string           `xml:",chardata"`
	Gid         *string          `xml:"gid,attr,omitempty"`
	SrsName     *string          `xml:"srsName,attr,omitempty"`
	Coord       []CoordType      `xml:"http://www.opengis.net/gml coord"`
	Coordinates *CoordinatesType `xml:"http://www.opengis.net/gml coordinates,omitempty"`
}

type CoordType struct {
	X float64 `xml:"http://www.opengis.net/gml X,omitempty"`
	Y float64 `xml:"http://www.opengis.net/gml Y,omitempty"`
	Z float64 `xml:"http://www.opengis.net/gml Z,omitempty"`
}

type CoordinatesType struct {
	Value   string  `xml:",chardata"`
	Decimal *string `xml:"decimal,attr,omitempty"`
	Cs      *string `xml:"cs,attr,omitempty"`
	Ts      *string `xml:"ts,attr,omitempty"`
}

type FeatureAssociationType struct {
	Feature           *AbstractFeatureType           `xml:"http://www.opengis.net/gml _Feature,omitempty"`
	FeatureCollection *AbstractFeatureCollectionType `xml:"http://www.opengis.net/gml _FeatureCollection,omitempty"`
	RemoteSchema      string                         `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField         string                         `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href              string                         `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role              string                         `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole           string                         `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title             string                         `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show              string                         `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate           string                         `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type GeometryAssociationType struct {
	Geometry           *AbstractGeometryType   `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	LineString         *LineStringType         `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearRing         *LinearRingType         `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MultiGeometry      *GeometryCollectionType `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString    *MultiLineStringType    `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint         *MultiPointType         `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPolygon       *MultiPolygonType       `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	Point              *PointType              `xml:"http://www.opengis.net/gml Point,omitempty"`
	Polygon            *PolygonType            `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	GeometryCollection *GeometryCollectionType `xml:"http://www.opengis.net/gml _GeometryCollection,omitempty"`
	RemoteSchema       string                  `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField          string                  `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href               string                  `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role               string                  `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole            string                  `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title              string                  `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show               string                  `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate            string                  `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type GeometryCollectionType struct {
	Value            string                    `xml:",chardata"`
	Gid              *string                   `xml:"gid,attr,omitempty"`
	SrsName          *string                   `xml:"srsName,attr,omitempty"`
	GeometryMember   []GeometryAssociationType `xml:"http://www.opengis.net/gml geometryMember"`
	LineStringMember []LineStringMemberType    `xml:"http://www.opengis.net/gml lineStringMember"`
	PointMember      []PointMemberType         `xml:"http://www.opengis.net/gml pointMember"`
	PolygonMember    []PolygonMemberType       `xml:"http://www.opengis.net/gml polygonMember"`
}

type LinearRingMemberType struct {
	LinearRing *LinearRingType `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
}

type LinearRingType struct {
	Value       string           `xml:",chardata"`
	Gid         *string          `xml:"gid,attr,omitempty"`
	SrsName     *string          `xml:"srsName,attr,omitempty"`
	Coord       []CoordType      `xml:"http://www.opengis.net/gml coord"`
	Coordinates *CoordinatesType `xml:"http://www.opengis.net/gml coordinates,omitempty"`
}

type LineStringMemberType struct {
	LineString *LineStringType `xml:"http://www.opengis.net/gml LineString,omitempty"`
}

type LineStringType struct {
	Value       string           `xml:",chardata"`
	Gid         *string          `xml:"gid,attr,omitempty"`
	SrsName     *string          `xml:"srsName,attr,omitempty"`
	Coord       []CoordType      `xml:"http://www.opengis.net/gml coord"`
	Coordinates *CoordinatesType `xml:"http://www.opengis.net/gml coordinates,omitempty"`
}

type MultiLineStringType struct {
	Value            string                    `xml:",chardata"`
	Gid              *string                   `xml:"gid,attr,omitempty"`
	SrsName          *string                   `xml:"srsName,attr,omitempty"`
	GeometryMember   []GeometryAssociationType `xml:"http://www.opengis.net/gml geometryMember"`
	LineStringMember []LineStringMemberType    `xml:"http://www.opengis.net/gml lineStringMember"`
}

type MultiPointType struct {
	Value       string            `xml:",chardata"`
	Gid         *string           `xml:"gid,attr,omitempty"`
	SrsName     *string           `xml:"srsName,attr,omitempty"`
	PointMember []PointMemberType `xml:"http://www.opengis.net/gml pointMember"`
}

type MultiPolygonType struct {
	Value         string              `xml:",chardata"`
	Gid           *string             `xml:"gid,attr,omitempty"`
	SrsName       *string             `xml:"srsName,attr,omitempty"`
	PolygonMember []PolygonMemberType `xml:"http://www.opengis.net/gml polygonMember"`
}

type PointMemberType struct {
	Point *PointType `xml:"http://www.opengis.net/gml Point,omitempty"`
}

type PointType struct {
	Value       string           `xml:",chardata"`
	Gid         *string          `xml:"gid,attr,omitempty"`
	SrsName     *string          `xml:"srsName,attr,omitempty"`
	Coord       *CoordType       `xml:"http://www.opengis.net/gml coord,omitempty"`
	Coordinates *CoordinatesType `xml:"http://www.opengis.net/gml coordinates,omitempty"`
}

type PolygonMemberType struct {
	Polygon *PolygonType `xml:"http://www.opengis.net/gml Polygon,omitempty"`
}

type PolygonType struct {
	Value           string                 `xml:",chardata"`
	Gid             *string                `xml:"gid,attr,omitempty"`
	SrsName         *string                `xml:"srsName,attr,omitempty"`
	OuterBoundaryIs *LinearRingMemberType  `xml:"http://www.opengis.net/gml outerBoundaryIs,omitempty"`
	InnerBoundaryIs []LinearRingMemberType `xml:"http://www.opengis.net/gml innerBoundaryIs"`
}
