package generated

type AbstractCityObjectType struct {
	CreationDate                                   string                       `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                string                       `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                              []ExternalReferenceType      `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                  []GeneralizationRelationType `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                              string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject []string                     `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
}

type AbstractSiteType struct {
	CreationDate                                   string                       `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                string                       `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                              []ExternalReferenceType      `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                  []GeneralizationRelationType `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                              string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject []string                     `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
	AbstractGenericApplicationPropertyOfSite       []string                     `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfSite"`
}

type AddressPropertyType struct {
	Address   *AddressType `xml:"http://www.opengis.net/citygml/2.0 Address,omitempty"`
	TypeField string       `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href      string       `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role      string       `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole   string       `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title     string       `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show      string       `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate   string       `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type AddressType struct {
	XalAddress                                  *XalAddressPropertyType `xml:"http://www.opengis.net/citygml/2.0 xalAddress,omitempty"`
	AbstractGenericApplicationPropertyOfAddress []string                `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfAddress"`
}

type CityModelType struct {
	AbstractGenericApplicationPropertyOfCityModel []string `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityModel"`
}

type ExternalObjectReferenceType struct {
	Name string `xml:"http://www.opengis.net/citygml/2.0 name,omitempty"`
	Uri  string `xml:"http://www.opengis.net/citygml/2.0 uri,omitempty"`
}

type ExternalReferenceType struct {
	InformationSystem string                       `xml:"http://www.opengis.net/citygml/2.0 informationSystem,omitempty"`
	ExternalObject    *ExternalObjectReferenceType `xml:"http://www.opengis.net/citygml/2.0 externalObject,omitempty"`
}

type GeneralizationRelationType struct {
	AbstractCityObject *AbstractCityObjectType `xml:"http://www.opengis.net/citygml/2.0 _CityObject,omitempty"`
	AbstractSite       *AbstractSiteType       `xml:"http://www.opengis.net/citygml/2.0 _Site,omitempty"`
	TypeField          string                  `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href               string                  `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role               string                  `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole            string                  `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title              string                  `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show               string                  `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate            string                  `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type ImplicitGeometryType struct {
	TransformationMatrix string `xml:"http://www.opengis.net/citygml/2.0 transformationMatrix,omitempty"`
	LibraryObject        string `xml:"http://www.opengis.net/citygml/2.0 libraryObject,omitempty"`
}

type ImplicitRepresentationPropertyType struct {
	ImplicitGeometry *ImplicitGeometryType `xml:"http://www.opengis.net/citygml/2.0 ImplicitGeometry,omitempty"`
	TypeField        string                `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href             string                `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role             string                `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole          string                `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title            string                `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show             string                `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate          string                `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type XalAddressPropertyType struct {
	// This container defines the details of the address. Can define multiple addresses including tracking address history
	AddressDetails *AddressDetails `xml:"urn:oasis:names:tc:ciq:xsdschema:xAL:2.0 AddressDetails,omitempty"`
}
