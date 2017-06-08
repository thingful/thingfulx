package thingfulx

const (

	// Permissions type

	// Reproduction is "making multiple copies"
	Reproduction = "https://creativecommons.org/ns#Reproduction"

	// Distribution is "distribution, public display, and publicly performance"
	Distribution = "https://creativecommons.org/ns#Distribution"

	// DerivativeWorks is "distribution of derivative works"
	DerivativeWorks = "https://creativecommons.org/ns#DerivativeWorks"

	// Sharing is "permits commercial derivatives, but only non-commercial distribution"
	Sharing = "https://creativecommons.org/ns#Sharing"

	// Requirements type

	// Notice is "copyright and license notices be kept intact"
	Notice = "https://creativecommons.org/ns#Notice"

	// Attribution is "credit be given to copyright holder and/or author"
	Attribution = "https://creativecommons.org/ns#Attribution"

	// ShareAlike is "derivative works be licensed under the same terms or compatible terms as the original work"
	ShareAlike = "https://creativecommons.org/ns#ShareAlike"

	// SourceCode is "source code (the preferred form for making modifications) must be provided when exercising some rights granted by the license"
	SourceCode = "https://creativecommons.org/ns#SourceCode"

	// Copyleft is "derivative and combined works must be licensed under specified terms, similar to those on the original work"
	Copyleft = "https://creativecommons.org/ns#Copyleft"

	// LesserCopyleft is "derivative works must be licensed under specified terms, with at least the same conditions as the original work; combinations with the work may be licensed under different terms"
	LesserCopyleft = "https://creativecommons.org/ns#LesserCopyleft"

	// Prohibitions type

	// CommercialUse is "exercising rights for commercial purposes"
	CommercialUse = "http://web.resource.org/cc/CommercialUse"

	// HighIncomeNationUse is "use in a non-developing country"
	HighIncomeNationUse = "https://creativecommons.org/ns#HighIncomeNationUse"
)

// DataLicense is struct used to describe data license
type DataLicense struct {
	// Human readable name for this license
	Name string `json:"name"`

	// url to the short version of the license
	URL string `json:"id"`

	// url to the long or legal version of the license
	LegalCodeURL string `json:"legalCode"`

	Permits []string `json:"permits"`

	Requires []string `json:"requires"`

	Prohibits []string `json:"prohibits"`
}

var (

	// CC0V1 is a predefined Creative Commons CC0 version 1.0 license
	CC0V1 = DataLicense{

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

	// CCByV3 is a predefined Creative Commons CC BY version 3.0 license
	CCByV3 = DataLicense{

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
			Notice,
		},

		Prohibits: []string{},
	}

	// CCBySaV4 is a predefined Creative Commons CC BY-SA version 4.0 license
	CCBySaV4 = DataLicense{

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
			Notice,
			ShareAlike,
		},

		Prohibits: []string{},
	}

	// CCByNcV3 is a predefined Creative Commons CC BY-NC version 3.0 license
	CCByNcV3 = DataLicense{

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
			Notice,
		},

		Prohibits: []string{
			CommercialUse,
		},
	}

	// CCByNdV3 is a predefined Creative Commons CC BY-ND version 3.0 license
	CCByNdV3 = DataLicense{ // i'm a bit confuse with this one how to describe it, need confirmation with CC

		Name: "Attribution-NoDerivs 3.0 Unported (CC BY-ND 3.0)",

		URL: "https://creativecommons.org/licenses/by-nd/3.0/",

		LegalCodeURL: "https://creativecommons.org/licenses/by-nd/3.0/legalcode",

		Permits: []string{
			Reproduction,
			Distribution,
		},

		Requires: []string{
			Attribution,
			Notice,
		},

		Prohibits: []string{},
	}

	// CCByNcSaV3 is a predefined Creative Commons CC BY-NC-SA version 3.0 license
	CCByNcSaV3 = DataLicense{

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
			Notice,
		},

		Prohibits: []string{
			CommercialUse,
		},
	}

	// OGLV3 is a predefined Open Government Licence version 3.0 license
	OGLV3 = DataLicense{

		Name: "Open Government Licence version 3.0",

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
			Notice,
		},

		Prohibits: []string{},
	}

	// PDDLV1 is a predefined Open Data Commons Public Domain Dedication and License
	PDDLV1 = DataLicense{

		Name: "Open Data Commons Public Domain Dedication and License (PDDL) v1.0",

		URL: "https://opendatacommons.org/licenses/pddl/1.0/",

		LegalCodeURL: "https://opendatacommons.org/licenses/pddl/1.0/",

		Permits: []string{
			Reproduction,
			Distribution,
			DerivativeWorks,
			Sharing,
		},

		Requires: []string{},

		Prohibits: []string{},
	}

	//ODCByV1 is a predefined Open Data Commons Attribution License
	ODCByV1 = DataLicense{

		Name: "Open Data Commons Attribution License (ODC-By) v1.0",

		URL: "https://opendatacommons.org/licenses/by/1.0/",

		LegalCodeURL: "https://opendatacommons.org/licenses/by/1.0/",

		Permits: []string{
			Reproduction,
			Distribution,
			DerivativeWorks,
			Sharing,
		},

		Requires: []string{
			Attribution,
			Notice,
		},

		Prohibits: []string{},
	}

	//ODCODbLV1 is a predefined Open Data Commons Open Database License
	ODCODbLV1 = DataLicense{

		Name: "Open Data Commons Open Database License (ODbL) v1.0",

		URL: "https://opendatacommons.org/licenses/odbl/1.0/",

		LegalCodeURL: "https://opendatacommons.org/licenses/odbl/1.0/",

		Permits: []string{
			Reproduction,
			Distribution,
			DerivativeWorks,
			Sharing,
		},

		Requires: []string{
			Attribution,
			Notice,
			ShareAlike,
		},

		Prohibits: []string{},
	}
)
