package schema

var (
	// Thing level properties

	// HasCategory is the property for the hasCategory relationship
	HasCategory = Expand("thingful:hasCategory")

	// HasCategoryUID is the property for the hasCategoryUID relationship
	HasCategoryUID = Expand("thingful:hasCategoryUID")

	// IsThingType is the property for the isThingType relationship
	IsThingType = Expand("thingful:isThingType")

	// HasDataLicense is the property fo the hasDataLicense relationship
	HasDataLicense = Expand("thingful:hasDataLicense")

	// Channel level properties

	// MeasuredBy is the property for the measuredBy relationship
	MeasuredBy = Expand("thingful:measuredBy")

	// HasDomainOfInterest represents the domain of interest of a channel
	HasDomainOfInterest = Expand("thingful:hasDomainOfInterest")

	// IsQuantityKind represents the quantity kind of a channel
	IsQuantityKind = Expand("thingful:isQuantityKind")

	// HasUnit represents the unit of a channel
	HasUnit = Expand("thingful:hasUnit")
)
