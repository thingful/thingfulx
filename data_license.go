package thingfulx

// Category is a simple type alias for a string

const (

	// Permissions
	// Reproduction refers to "making multiple copies"
	Reproduction = "https://creativecommons.org/ns#Reproduction"

	// Distribution refers to "distribution, public display, and publicly performance"
	Distribution = "https://creativecommons.org/ns#Distribution"

	// DerivativeWorks refers to "distribution of derivative works"
	DerivativeWorks = "https://creativecommons.org/ns#DerivativeWorks"

	// Sharing refers to "permits commercial derivatives, but only non-commercial distribution"
	Sharing = "https://creativecommons.org/ns#Sharing"

	// Requirements
	// Notice refers to "copyright and license notices be kept intact"
	Notice = "https://creativecommons.org/ns#Notice"

	// Attribution refers to "credit be given to copyright holder and/or author"
	Attribution = "https://creativecommons.org/ns#Attribution"

	// ShareAlike refers to "derivative works be licensed under the same terms or compatible terms as the original work"
	ShareAlike = "https://creativecommons.org/ns#ShareAlike"

	// SourceCode refers to "source code (the preferred form for making modifications) must be provided when exercising some rights granted by the license"
	SourceCode = "https://creativecommons.org/ns#SourceCode"

	// Copyleft refers to "derivative and combined works must be licensed under specified terms, similar to those on the original work"
	Copyleft = "https://creativecommons.org/ns#Copyleft"

	// LesserCopyleft refers to "derivative works must be licensed under specified terms, with at least the same conditions as the original work; combinations with the work may be licensed under different terms"
	LesserCopyleft = "https://creativecommons.org/ns#LesserCopyleft"

	// Prohibitions
	// CommercialUse refers to "exercising rights for commercial purposes"
	CommercialUse = "http://web.resource.org/cc/CommercialUse"

	// HighIncomeNationUse refers to "use in a non-developing country"
	HighIncomeNationUse = "https://creativecommons.org/ns#HighIncomeNationUse"
)

type License struct {
	// Human readable name for this license
	Name string `json:"name"`

	// url to the short version of the license
	URL string `json:"id"`

	// url to the long or legal version of the license
	LegalCodeURL string `json:"legalCodeUrl"`

	Permits []string `json:"permits"`

	Requires []string `json:"requires"`

	Prohibits []string `json:"prohibits"`
}

var (
	CC0V1 = License{

		Name: "CC0 1.0 Universal (CC0 1.0) Public Domain Dedication",

		URL: "https://creativecommons.org/publicdomain/zero/1.0/",

		LegalCodeURL: "https://creativecommons.org/publicdomain/zero/1.0/legalcode",

		Permits: []string{
			Reproduction,
			Distribution,
			DerivativeWorks,
			Sharing,
		},

		Requires: []string{},

		Prohibits: []string{},
	}

	CCByV3 = License{

		Name: "Attribution 3.0 Unported (CC BY 3.0)",

		URL: "https://creativecommons.org/licenses/by/3.0/",

		LegalCodeURL: "https://creativecommons.org/licenses/by/3.0/legalcode",

		Permits: []string{
			Reproduction,
			Distribution,
			DerivativeWorks,
			Sharing,
		},

		Requires: []string{
			Attribution,
		},

		Prohibits: []string{},
	}

	CCBySaV4 = License{

		Name: "Attribution-ShareAlike 4.0 International (CC BY-SA 4.0)",

		URL: "https://creativecommons.org/licenses/by-sa/4.0/",

		LegalCodeURL: "https://creativecommons.org/licenses/by-sa/4.0/legalcode",

		Permits: []string{
			Reproduction,
			Distribution,
			DerivativeWorks,
			Sharing,
		},

		Requires: []string{
			Attribution,
			ShareAlike,
		},

		Prohibits: []string{},
	}

	CCByNcV3 = License{

		Name: "Attribution-NonCommercial 3.0 Unported (CC BY-NC 3.0)",

		URL: "https://creativecommons.org/licenses/by-nc/3.0/",

		LegalCodeURL: "https://creativecommons.org/licenses/by-nc/3.0/legalcode",

		Permits: []string{
			Reproduction,
			Distribution,
			DerivativeWorks,
			Sharing,
		},

		Requires: []string{
			Attribution,
		},

		Prohibits: []string{
			CommercialUse,
		},
	}

	CCByNdV3 = License{ // i'm a bit confuse with this one how to describe it, need confirmation with CC

		Name: "Attribution-NoDerivs 3.0 Unported (CC BY-ND 3.0)",

		URL: "https://creativecommons.org/licenses/by-nd/3.0/",

		LegalCodeURL: "https://creativecommons.org/licenses/by-nd/3.0/legalcode",

		Permits: []string{
			Reproduction,
			Distribution,
			Sharing,
		},

		Requires: []string{
			Attribution,
		},

		Prohibits: []string{},
	}

	CCByNcSaV3 = License{

		Name: "Attribution-NonCommercial-ShareAlike 3.0 Unported (CC BY-NC-SA 3.0)",

		URL: "https://creativecommons.org/licenses/by-nc-sa/3.0/",

		LegalCodeURL: "https://creativecommons.org/licenses/by-nc-sa/3.0/legalcode",

		Permits: []string{
			Reproduction,
			Distribution,
			Sharing,
			DerivativeWorks,
		},

		Requires: []string{
			Attribution,
			ShareAlike,
		},

		Prohibits: []string{
			CommercialUse,
		},
	}

	OGLV3 = License{

		Name: "open government licence version 3.0",

		URL: "http://www.nationalarchives.gov.uk/doc/open-government-licence/version/3/",

		LegalCodeURL: "http://www.nationalarchives.gov.uk/doc/open-government-licence/version/3/",

		Permits: []string{
			Reproduction,
			Distribution,
			DerivativeWorks,
			Sharing,
		},

		Requires: []string{
			Attribution,
		},

		Prohibits: []string{},
	}
)
