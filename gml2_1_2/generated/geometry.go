package gml

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
	Gid     *string `xml:"gid,attr,omitempty"`
	SrsName string  `xml:"srsName,attr"`
}

type AbstractGeometryType struct {
	Gid     *string `xml:"gid,attr,omitempty"`
	SrsName *string `xml:"srsName,attr,omitempty"`
}

type BoundingShapeType struct {
	Box  *BoxType `xml:"http://www.opengis.net/gml Box,omitempty"`
	Null string   `xml:"http://www.opengis.net/gml null,omitempty"`
}

type BoxType struct {
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
	AbstractFeature           *AbstractFeatureType           `xml:"http://www.opengis.net/gml _Feature,omitempty"`
	AbstractFeatureCollection *AbstractFeatureCollectionType `xml:"http://www.opengis.net/gml _FeatureCollection,omitempty"`
	RemoteSchema              string                         `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                 string                         `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                      string                         `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                      string                         `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                   string                         `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                     string                         `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                      string                         `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                   string                         `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type GeometryAssociationType struct {
	AbstractGeometry           *AbstractGeometryType   `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	LineString                 *LineStringType         `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearRing                 *LinearRingType         `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MultiGeometry              *GeometryCollectionType `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString            *MultiLineStringType    `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                 *MultiPointType         `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPolygon               *MultiPolygonType       `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	Point                      *PointType              `xml:"http://www.opengis.net/gml Point,omitempty"`
	Polygon                    *PolygonType            `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	AbstractGeometryCollection *GeometryCollectionType `xml:"http://www.opengis.net/gml _GeometryCollection,omitempty"`
	RemoteSchema               string                  `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                  string                  `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                       string                  `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                       string                  `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                    string                  `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                      string                  `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                       string                  `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                    string                  `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type GeometryCollectionType struct {
	Gid              *string                   `xml:"gid,attr,omitempty"`
	SrsName          string                    `xml:"srsName,attr"`
	GeometryMember   []GeometryAssociationType `xml:"http://www.opengis.net/gml geometryMember"`
	LineStringMember []LineStringMemberType    `xml:"http://www.opengis.net/gml lineStringMember"`
	PointMember      []PointMemberType         `xml:"http://www.opengis.net/gml pointMember"`
	PolygonMember    []PolygonMemberType       `xml:"http://www.opengis.net/gml polygonMember"`
}

type GeometryPropertyType struct {
	AbstractGeometry           *AbstractGeometryType   `xml:"http://www.opengis.net/gml _Geometry,omitempty"`
	LineString                 *LineStringType         `xml:"http://www.opengis.net/gml LineString,omitempty"`
	LinearRing                 *LinearRingType         `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	MultiGeometry              *GeometryCollectionType `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	MultiLineString            *MultiLineStringType    `xml:"http://www.opengis.net/gml MultiLineString,omitempty"`
	MultiPoint                 *MultiPointType         `xml:"http://www.opengis.net/gml MultiPoint,omitempty"`
	MultiPolygon               *MultiPolygonType       `xml:"http://www.opengis.net/gml MultiPolygon,omitempty"`
	Point                      *PointType              `xml:"http://www.opengis.net/gml Point,omitempty"`
	Polygon                    *PolygonType            `xml:"http://www.opengis.net/gml Polygon,omitempty"`
	AbstractGeometryCollection *GeometryCollectionType `xml:"http://www.opengis.net/gml _GeometryCollection,omitempty"`
	RemoteSchema               string                  `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField                  string                  `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                       string                  `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                       string                  `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                    string                  `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                      string                  `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                       string                  `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                    string                  `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type LineStringMemberType struct {
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

type LineStringType struct {
	Gid         *string          `xml:"gid,attr,omitempty"`
	SrsName     *string          `xml:"srsName,attr,omitempty"`
	Coord       []CoordType      `xml:"http://www.opengis.net/gml coord"`
	Coordinates *CoordinatesType `xml:"http://www.opengis.net/gml coordinates,omitempty"`
}

type LinearRingMemberType struct {
	LinearRing   *LinearRingType `xml:"http://www.opengis.net/gml LinearRing,omitempty"`
	RemoteSchema string          `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField    string          `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string          `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string          `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string          `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string          `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string          `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string          `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type LinearRingType struct {
	Gid         *string          `xml:"gid,attr,omitempty"`
	SrsName     *string          `xml:"srsName,attr,omitempty"`
	Coord       []CoordType      `xml:"http://www.opengis.net/gml coord"`
	Coordinates *CoordinatesType `xml:"http://www.opengis.net/gml coordinates,omitempty"`
}

type MultiGeometryPropertyType struct {
	MultiGeometry *GeometryCollectionType `xml:"http://www.opengis.net/gml MultiGeometry,omitempty"`
	RemoteSchema  string                  `xml:"http://www.opengis.net/gml remoteSchema,attr,omitempty"`
	TypeField     string                  `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href          string                  `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role          string                  `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole       string                  `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title         string                  `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show          string                  `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate       string                  `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type MultiLineStringPropertyType struct {
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
	LineStringMember []LineStringMemberType `xml:"http://www.opengis.net/gml lineStringMember"`
	Gid              *string                `xml:"gid,attr,omitempty"`
	SrsName          string                 `xml:"srsName,attr"`
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
	PointMember []PointMemberType `xml:"http://www.opengis.net/gml pointMember"`
	Gid         *string           `xml:"gid,attr,omitempty"`
	SrsName     string            `xml:"srsName,attr"`
}

type MultiPolygonPropertyType struct {
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
	PolygonMember []PolygonMemberType `xml:"http://www.opengis.net/gml polygonMember"`
	Gid           *string             `xml:"gid,attr,omitempty"`
	SrsName       string              `xml:"srsName,attr"`
}

type PointMemberType struct {
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
	Gid         *string          `xml:"gid,attr,omitempty"`
	SrsName     *string          `xml:"srsName,attr,omitempty"`
	Coord       *CoordType       `xml:"http://www.opengis.net/gml coord,omitempty"`
	Coordinates *CoordinatesType `xml:"http://www.opengis.net/gml coordinates,omitempty"`
}

type PolygonMemberType struct {
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
	Gid             *string                `xml:"gid,attr,omitempty"`
	SrsName         *string                `xml:"srsName,attr,omitempty"`
	OuterBoundaryIs *LinearRingMemberType  `xml:"http://www.opengis.net/gml outerBoundaryIs,omitempty"`
	InnerBoundaryIs []LinearRingMemberType `xml:"http://www.opengis.net/gml innerBoundaryIs"`
}
