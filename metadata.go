package thingfulx

// Metadata is a data structure containing additional metadata about the Thing.
// A key requirement for Metadata at the Thing level is that it must contain a
// 'thingful:Category' property whose value must be a Category.Name as well as
// a 'thingful:CategoryUID' whose value must me a Category.UID. These two
// metadata object are required for correctly displaying Things on the
// thingful.net map.  The Metadata object must also contain information about
// the data visibility - see http://theodi.org/data-spectrum for more info.
type Metadata struct {
	Prop string `json:"prop"`
	Val  string `json:"val"`
}
