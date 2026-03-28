package generated

type AbstractBoundarySurfaceType struct {
	CreationDate                                        string                       `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                     string                       `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                                   []ExternalReferenceType      `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                       []GeneralizationRelationType `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                                   string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                     string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject      []string                     `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
	Opening                                             []OpeningPropertyType        `xml:"http://www.opengis.net/citygml/building/2.0 opening"`
	AbstractGenericApplicationPropertyOfBoundarySurface []string                     `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfBoundarySurface"`
}

type AbstractBuildingType struct {
	CreationDate                                         string                                `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                      string                                `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                                    []ExternalReferenceType               `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                        []GeneralizationRelationType          `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                                    string                                `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                      string                                `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject       []string                              `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
	AbstractGenericApplicationPropertyOfSite             []string                              `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfSite"`
	YearOfConstruction                                   string                                `xml:"http://www.opengis.net/citygml/building/2.0 yearOfConstruction,omitempty"`
	YearOfDemolition                                     string                                `xml:"http://www.opengis.net/citygml/building/2.0 yearOfDemolition,omitempty"`
	StoreysAboveGround                                   int                                   `xml:"http://www.opengis.net/citygml/building/2.0 storeysAboveGround,omitempty"`
	StoreysBelowGround                                   int                                   `xml:"http://www.opengis.net/citygml/building/2.0 storeysBelowGround,omitempty"`
	OuterBuildingInstallation                            []BuildingInstallationPropertyType    `xml:"http://www.opengis.net/citygml/building/2.0 outerBuildingInstallation"`
	InteriorBuildingInstallation                         []IntBuildingInstallationPropertyType `xml:"http://www.opengis.net/citygml/building/2.0 interiorBuildingInstallation"`
	InteriorRoom                                         []InteriorRoomPropertyType            `xml:"http://www.opengis.net/citygml/building/2.0 interiorRoom"`
	ConsistsOfBuildingPart                               []BuildingPartPropertyType            `xml:"http://www.opengis.net/citygml/building/2.0 consistsOfBuildingPart"`
	Address                                              []AddressPropertyType                 `xml:"http://www.opengis.net/citygml/building/2.0 address"`
	AbstractGenericApplicationPropertyOfAbstractBuilding []string                              `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfAbstractBuilding"`
}

type AbstractOpeningType struct {
	CreationDate                                   string                              `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                string                              `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                              []ExternalReferenceType             `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                  []GeneralizationRelationType        `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                              string                              `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                string                              `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject []string                            `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
	Lod3ImplicitRepresentation                     *ImplicitRepresentationPropertyType `xml:"http://www.opengis.net/citygml/building/2.0 lod3ImplicitRepresentation,omitempty"`
	Lod4ImplicitRepresentation                     *ImplicitRepresentationPropertyType `xml:"http://www.opengis.net/citygml/building/2.0 lod4ImplicitRepresentation,omitempty"`
	AbstractGenericApplicationPropertyOfOpening    []string                            `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfOpening"`
}

type BoundarySurfacePropertyType struct {
	AbstractBoundarySurface *AbstractBoundarySurfaceType `xml:"http://www.opengis.net/citygml/building/2.0 _BoundarySurface,omitempty"`
	CeilingSurface          *CeilingSurfaceType          `xml:"http://www.opengis.net/citygml/building/2.0 CeilingSurface,omitempty"`
	ClosureSurface          *ClosureSurfaceType          `xml:"http://www.opengis.net/citygml/building/2.0 ClosureSurface,omitempty"`
	FloorSurface            *FloorSurfaceType            `xml:"http://www.opengis.net/citygml/building/2.0 FloorSurface,omitempty"`
	GroundSurface           *GroundSurfaceType           `xml:"http://www.opengis.net/citygml/building/2.0 GroundSurface,omitempty"`
	InteriorWallSurface     *InteriorWallSurfaceType     `xml:"http://www.opengis.net/citygml/building/2.0 InteriorWallSurface,omitempty"`
	OuterCeilingSurface     *OuterCeilingSurfaceType     `xml:"http://www.opengis.net/citygml/building/2.0 OuterCeilingSurface,omitempty"`
	OuterFloorSurface       *OuterFloorSurfaceType       `xml:"http://www.opengis.net/citygml/building/2.0 OuterFloorSurface,omitempty"`
	RoofSurface             *RoofSurfaceType             `xml:"http://www.opengis.net/citygml/building/2.0 RoofSurface,omitempty"`
	WallSurface             *WallSurfaceType             `xml:"http://www.opengis.net/citygml/building/2.0 WallSurface,omitempty"`
	TypeField               string                       `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                    string                       `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                    string                       `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                 string                       `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                   string                       `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                    string                       `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                 string                       `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type BuildingFurnitureType struct {
	CreationDate                                          string                              `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                       string                              `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                                     []ExternalReferenceType             `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                         []GeneralizationRelationType        `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                                     string                              `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                       string                              `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject        []string                            `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
	Lod4ImplicitRepresentation                            *ImplicitRepresentationPropertyType `xml:"http://www.opengis.net/citygml/building/2.0 lod4ImplicitRepresentation,omitempty"`
	AbstractGenericApplicationPropertyOfBuildingFurniture []string                            `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfBuildingFurniture"`
}

type BuildingInstallationPropertyType struct {
	BuildingInstallation *BuildingInstallationType `xml:"http://www.opengis.net/citygml/building/2.0 BuildingInstallation,omitempty"`
	TypeField            string                    `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                 string                    `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                 string                    `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole              string                    `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                string                    `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                 string                    `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate              string                    `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type BuildingInstallationType struct {
	CreationDate                                             string                              `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                          string                              `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                                        []ExternalReferenceType             `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                            []GeneralizationRelationType        `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                                        string                              `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                          string                              `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject           []string                            `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
	Lod2ImplicitRepresentation                               *ImplicitRepresentationPropertyType `xml:"http://www.opengis.net/citygml/building/2.0 lod2ImplicitRepresentation,omitempty"`
	Lod3ImplicitRepresentation                               *ImplicitRepresentationPropertyType `xml:"http://www.opengis.net/citygml/building/2.0 lod3ImplicitRepresentation,omitempty"`
	Lod4ImplicitRepresentation                               *ImplicitRepresentationPropertyType `xml:"http://www.opengis.net/citygml/building/2.0 lod4ImplicitRepresentation,omitempty"`
	AbstractGenericApplicationPropertyOfBuildingInstallation []string                            `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfBuildingInstallation"`
}

type BuildingPartPropertyType struct {
	BuildingPart *BuildingPartType `xml:"http://www.opengis.net/citygml/building/2.0 BuildingPart,omitempty"`
	TypeField    string            `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href         string            `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role         string            `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole      string            `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title        string            `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show         string            `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate      string            `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type BuildingPartType struct {
	CreationDate                                         string                                `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                      string                                `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                                    []ExternalReferenceType               `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                        []GeneralizationRelationType          `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                                    string                                `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                      string                                `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject       []string                              `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
	AbstractGenericApplicationPropertyOfSite             []string                              `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfSite"`
	YearOfConstruction                                   string                                `xml:"http://www.opengis.net/citygml/building/2.0 yearOfConstruction,omitempty"`
	YearOfDemolition                                     string                                `xml:"http://www.opengis.net/citygml/building/2.0 yearOfDemolition,omitempty"`
	StoreysAboveGround                                   int                                   `xml:"http://www.opengis.net/citygml/building/2.0 storeysAboveGround,omitempty"`
	StoreysBelowGround                                   int                                   `xml:"http://www.opengis.net/citygml/building/2.0 storeysBelowGround,omitempty"`
	OuterBuildingInstallation                            []BuildingInstallationPropertyType    `xml:"http://www.opengis.net/citygml/building/2.0 outerBuildingInstallation"`
	InteriorBuildingInstallation                         []IntBuildingInstallationPropertyType `xml:"http://www.opengis.net/citygml/building/2.0 interiorBuildingInstallation"`
	InteriorRoom                                         []InteriorRoomPropertyType            `xml:"http://www.opengis.net/citygml/building/2.0 interiorRoom"`
	ConsistsOfBuildingPart                               []BuildingPartPropertyType            `xml:"http://www.opengis.net/citygml/building/2.0 consistsOfBuildingPart"`
	Address                                              []AddressPropertyType                 `xml:"http://www.opengis.net/citygml/building/2.0 address"`
	AbstractGenericApplicationPropertyOfAbstractBuilding []string                              `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfAbstractBuilding"`
	AbstractGenericApplicationPropertyOfBuildingPart     []string                              `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfBuildingPart"`
}

type BuildingType struct {
	CreationDate                                         string                                `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                      string                                `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                                    []ExternalReferenceType               `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                        []GeneralizationRelationType          `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                                    string                                `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                      string                                `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject       []string                              `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
	AbstractGenericApplicationPropertyOfSite             []string                              `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfSite"`
	YearOfConstruction                                   string                                `xml:"http://www.opengis.net/citygml/building/2.0 yearOfConstruction,omitempty"`
	YearOfDemolition                                     string                                `xml:"http://www.opengis.net/citygml/building/2.0 yearOfDemolition,omitempty"`
	StoreysAboveGround                                   int                                   `xml:"http://www.opengis.net/citygml/building/2.0 storeysAboveGround,omitempty"`
	StoreysBelowGround                                   int                                   `xml:"http://www.opengis.net/citygml/building/2.0 storeysBelowGround,omitempty"`
	OuterBuildingInstallation                            []BuildingInstallationPropertyType    `xml:"http://www.opengis.net/citygml/building/2.0 outerBuildingInstallation"`
	InteriorBuildingInstallation                         []IntBuildingInstallationPropertyType `xml:"http://www.opengis.net/citygml/building/2.0 interiorBuildingInstallation"`
	InteriorRoom                                         []InteriorRoomPropertyType            `xml:"http://www.opengis.net/citygml/building/2.0 interiorRoom"`
	ConsistsOfBuildingPart                               []BuildingPartPropertyType            `xml:"http://www.opengis.net/citygml/building/2.0 consistsOfBuildingPart"`
	Address                                              []AddressPropertyType                 `xml:"http://www.opengis.net/citygml/building/2.0 address"`
	AbstractGenericApplicationPropertyOfAbstractBuilding []string                              `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfAbstractBuilding"`
	AbstractGenericApplicationPropertyOfBuilding         []string                              `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfBuilding"`
}

type CeilingSurfaceType struct {
	CreationDate                                        string                       `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                     string                       `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                                   []ExternalReferenceType      `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                       []GeneralizationRelationType `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                                   string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                     string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject      []string                     `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
	Opening                                             []OpeningPropertyType        `xml:"http://www.opengis.net/citygml/building/2.0 opening"`
	AbstractGenericApplicationPropertyOfBoundarySurface []string                     `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfBoundarySurface"`
	AbstractGenericApplicationPropertyOfCeilingSurface  []string                     `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfCeilingSurface"`
}

type ClosureSurfaceType struct {
	CreationDate                                        string                       `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                     string                       `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                                   []ExternalReferenceType      `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                       []GeneralizationRelationType `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                                   string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                     string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject      []string                     `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
	Opening                                             []OpeningPropertyType        `xml:"http://www.opengis.net/citygml/building/2.0 opening"`
	AbstractGenericApplicationPropertyOfBoundarySurface []string                     `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfBoundarySurface"`
	AbstractGenericApplicationPropertyOfClosureSurface  []string                     `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfClosureSurface"`
}

type DoorType struct {
	CreationDate                                   string                              `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                string                              `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                              []ExternalReferenceType             `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                  []GeneralizationRelationType        `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                              string                              `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                string                              `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject []string                            `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
	Lod3ImplicitRepresentation                     *ImplicitRepresentationPropertyType `xml:"http://www.opengis.net/citygml/building/2.0 lod3ImplicitRepresentation,omitempty"`
	Lod4ImplicitRepresentation                     *ImplicitRepresentationPropertyType `xml:"http://www.opengis.net/citygml/building/2.0 lod4ImplicitRepresentation,omitempty"`
	AbstractGenericApplicationPropertyOfOpening    []string                            `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfOpening"`
	Address                                        []AddressPropertyType               `xml:"http://www.opengis.net/citygml/building/2.0 address"`
	AbstractGenericApplicationPropertyOfDoor       []string                            `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfDoor"`
}

type FloorSurfaceType struct {
	CreationDate                                        string                       `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                     string                       `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                                   []ExternalReferenceType      `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                       []GeneralizationRelationType `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                                   string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                     string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject      []string                     `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
	Opening                                             []OpeningPropertyType        `xml:"http://www.opengis.net/citygml/building/2.0 opening"`
	AbstractGenericApplicationPropertyOfBoundarySurface []string                     `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfBoundarySurface"`
	AbstractGenericApplicationPropertyOfFloorSurface    []string                     `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfFloorSurface"`
}

type GroundSurfaceType struct {
	CreationDate                                        string                       `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                     string                       `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                                   []ExternalReferenceType      `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                       []GeneralizationRelationType `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                                   string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                     string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject      []string                     `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
	Opening                                             []OpeningPropertyType        `xml:"http://www.opengis.net/citygml/building/2.0 opening"`
	AbstractGenericApplicationPropertyOfBoundarySurface []string                     `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfBoundarySurface"`
	AbstractGenericApplicationPropertyOfGroundSurface   []string                     `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfGroundSurface"`
}

type IntBuildingInstallationPropertyType struct {
	IntBuildingInstallation *IntBuildingInstallationType `xml:"http://www.opengis.net/citygml/building/2.0 IntBuildingInstallation,omitempty"`
	TypeField               string                       `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href                    string                       `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role                    string                       `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole                 string                       `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title                   string                       `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show                    string                       `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate                 string                       `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type IntBuildingInstallationType struct {
	CreationDate                                                string                              `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                             string                              `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                                           []ExternalReferenceType             `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                               []GeneralizationRelationType        `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                                           string                              `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                             string                              `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject              []string                            `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
	Lod4ImplicitRepresentation                                  *ImplicitRepresentationPropertyType `xml:"http://www.opengis.net/citygml/building/2.0 lod4ImplicitRepresentation,omitempty"`
	AbstractGenericApplicationPropertyOfIntBuildingInstallation []string                            `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfIntBuildingInstallation"`
}

type InteriorFurniturePropertyType struct {
	BuildingFurniture *BuildingFurnitureType `xml:"http://www.opengis.net/citygml/building/2.0 BuildingFurniture,omitempty"`
	TypeField         string                 `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href              string                 `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role              string                 `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole           string                 `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title             string                 `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show              string                 `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate           string                 `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type InteriorRoomPropertyType struct {
	Room      *RoomType `xml:"http://www.opengis.net/citygml/building/2.0 Room,omitempty"`
	TypeField string    `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href      string    `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role      string    `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole   string    `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title     string    `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show      string    `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate   string    `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type InteriorWallSurfaceType struct {
	CreationDate                                            string                       `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                         string                       `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                                       []ExternalReferenceType      `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                           []GeneralizationRelationType `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                                       string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                         string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject          []string                     `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
	Opening                                                 []OpeningPropertyType        `xml:"http://www.opengis.net/citygml/building/2.0 opening"`
	AbstractGenericApplicationPropertyOfBoundarySurface     []string                     `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfBoundarySurface"`
	AbstractGenericApplicationPropertyOfInteriorWallSurface []string                     `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfInteriorWallSurface"`
}

type OpeningPropertyType struct {
	AbstractOpening *AbstractOpeningType `xml:"http://www.opengis.net/citygml/building/2.0 _Opening,omitempty"`
	Door            *DoorType            `xml:"http://www.opengis.net/citygml/building/2.0 Door,omitempty"`
	Window          *WindowType          `xml:"http://www.opengis.net/citygml/building/2.0 Window,omitempty"`
	TypeField       string               `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Href            string               `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Role            string               `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Arcrole         string               `xml:"http://www.w3.org/1999/xlink arcrole,attr,omitempty"`
	Title           string               `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show            string               `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate         string               `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type OuterCeilingSurfaceType struct {
	CreationDate                                            string                       `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                         string                       `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                                       []ExternalReferenceType      `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                           []GeneralizationRelationType `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                                       string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                         string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject          []string                     `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
	Opening                                                 []OpeningPropertyType        `xml:"http://www.opengis.net/citygml/building/2.0 opening"`
	AbstractGenericApplicationPropertyOfBoundarySurface     []string                     `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfBoundarySurface"`
	AbstractGenericApplicationPropertyOfOuterCeilingSurface []string                     `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfOuterCeilingSurface"`
}

type OuterFloorSurfaceType struct {
	CreationDate                                          string                       `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                       string                       `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                                     []ExternalReferenceType      `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                         []GeneralizationRelationType `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                                     string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                       string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject        []string                     `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
	Opening                                               []OpeningPropertyType        `xml:"http://www.opengis.net/citygml/building/2.0 opening"`
	AbstractGenericApplicationPropertyOfBoundarySurface   []string                     `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfBoundarySurface"`
	AbstractGenericApplicationPropertyOfOuterFloorSurface []string                     `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfOuterFloorSurface"`
}

type RoofSurfaceType struct {
	CreationDate                                        string                       `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                     string                       `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                                   []ExternalReferenceType      `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                       []GeneralizationRelationType `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                                   string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                     string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject      []string                     `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
	Opening                                             []OpeningPropertyType        `xml:"http://www.opengis.net/citygml/building/2.0 opening"`
	AbstractGenericApplicationPropertyOfBoundarySurface []string                     `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfBoundarySurface"`
	AbstractGenericApplicationPropertyOfRoofSurface     []string                     `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfRoofSurface"`
}

type RoomType struct {
	CreationDate                                   string                                `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                string                                `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                              []ExternalReferenceType               `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                  []GeneralizationRelationType          `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                              string                                `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                string                                `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject []string                              `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
	InteriorFurniture                              []InteriorFurniturePropertyType       `xml:"http://www.opengis.net/citygml/building/2.0 interiorFurniture"`
	RoomInstallation                               []IntBuildingInstallationPropertyType `xml:"http://www.opengis.net/citygml/building/2.0 roomInstallation"`
	AbstractGenericApplicationPropertyOfRoom       []string                              `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfRoom"`
}

type WallSurfaceType struct {
	CreationDate                                        string                       `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                     string                       `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                                   []ExternalReferenceType      `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                       []GeneralizationRelationType `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                                   string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                     string                       `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject      []string                     `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
	Opening                                             []OpeningPropertyType        `xml:"http://www.opengis.net/citygml/building/2.0 opening"`
	AbstractGenericApplicationPropertyOfBoundarySurface []string                     `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfBoundarySurface"`
	AbstractGenericApplicationPropertyOfWallSurface     []string                     `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfWallSurface"`
}

type WindowType struct {
	CreationDate                                   string                              `xml:"http://www.opengis.net/citygml/2.0 creationDate,omitempty"`
	TerminationDate                                string                              `xml:"http://www.opengis.net/citygml/2.0 terminationDate,omitempty"`
	ExternalReference                              []ExternalReferenceType             `xml:"http://www.opengis.net/citygml/2.0 externalReference"`
	GeneralizesTo                                  []GeneralizationRelationType        `xml:"http://www.opengis.net/citygml/2.0 generalizesTo"`
	RelativeToTerrain                              string                              `xml:"http://www.opengis.net/citygml/2.0 relativeToTerrain,omitempty"`
	RelativeToWater                                string                              `xml:"http://www.opengis.net/citygml/2.0 relativeToWater,omitempty"`
	AbstractGenericApplicationPropertyOfCityObject []string                            `xml:"http://www.opengis.net/citygml/2.0 _GenericApplicationPropertyOfCityObject"`
	Lod3ImplicitRepresentation                     *ImplicitRepresentationPropertyType `xml:"http://www.opengis.net/citygml/building/2.0 lod3ImplicitRepresentation,omitempty"`
	Lod4ImplicitRepresentation                     *ImplicitRepresentationPropertyType `xml:"http://www.opengis.net/citygml/building/2.0 lod4ImplicitRepresentation,omitempty"`
	AbstractGenericApplicationPropertyOfOpening    []string                            `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfOpening"`
	AbstractGenericApplicationPropertyOfWindow     []string                            `xml:"http://www.opengis.net/citygml/building/2.0 _GenericApplicationPropertyOfWindow"`
}
