package thingfulx

// Permission is a string type alias used to define a permission that may be
// granted by a license.
type Permission string

// Requirement is a string type alias used to define a requirement that may be
// enforced by a license.
type Requirement string

// Prohibition is a string type alias used to define a prohibtion that may be
// dictated by a license.
type Prohibition string

const (
	// ReproductionPerm is "making multiple copies of a dataset
	ReproductionPerm Permission = "cc:Reproduction"

	// DistributionPerm is "distribution, public display, and publicly performance"
	DistributionPerm Permission = "cc:Distribution"

	// DerivativeWorksPerm is "distribution of derivative works"
	DerivativeWorksPerm Permission = "cc:DerivativeWorks"

	// SharingPerm is "permits commercial derivatives, but only non-commercial
	// distribution"
	SharingPerm Permission = "cc:Sharing"

	// Requirements type

	// NoticeReq is "copyright and license notices be kept intact"
	NoticeReq Requirement = "cc:Notice"

	// AttributionReq is "credit required to give to copyright holder and/or author"
	AttributionReq Requirement = "cc:Attribution"

	// ShareAlikeReq is "derivative works be licensed under the same terms or
	// compatible terms as the original work"
	ShareAlikeReq Requirement = "cc:ShareAlike"

	// SourceCodeReq is "source code (the preferred form for making modifications)
	// must be provided when exercising some rights granted by the license"
	SourceCodeReq Requirement = "cc:SourceCode"

	// CopyleftReq is "derivative and combined works must be licensed under
	// specified terms, similar to those on the original work"
	CopyleftReq Requirement = "cc:Copyleft"

	// LesserCopyleftReq is "derivative works must be licensed under specified
	// terms, with at least the same conditions as the original work;
	// combinations with the work may be licensed under different terms"
	LesserCopyleftReq Requirement = "cc:LesserCopyleft"

	// Prohibitions type

	// CommercialUseProhib is "exercising rights for commercial purposes"
	CommercialUseProhib Prohibition = "cc:CommercialUse"

	// HighIncomeNationUseProhib is "use in a non-developing country"
	HighIncomeNationUseProhib = "cc:HighIncomeNationUse"

	// License URLS

	// CC0V1URL is the URL identifier for CC0 1.0 Universal License
	CC0V1URL = "https://creativecommons.org/publicdomain/zero/1.0/"

	// CCByV3URL is the URL identifier for Creative Commons Attribution 3.0
	// Unported License
	CCByV3URL = "https://creativecommons.org/licenses/by/3.0/"

	// CCBySAV4URL is the URL identifier for Creative Commons
	// Attribution-ShareAlike 4.0 International
	CCBySAV4URL = "https://creativecommons.org/licenses/by-sa/4.0/"

	// CCByV4URL is the URL identifier for the Creative Commons Attribution 4.0
	// License
	CCByV4URL = "https://creativecommons.org/licenses/by/4.0/"

	// CCByNCV3URL is the URL identifier for Creative Commons
	// Atribution-NonCommercial 3.0 License
	CCByNCV3URL = "https://creativecommons.org/licenses/by-nc/3.0/"

	// CCByNDV3URL is the URL identifier for Creative Commons
	// Attribution-NoDerivs 3.0 License.
	CCByNDV3URL = "https://creativecommons.org/licenses/by-nd/3.0/"

	// CCByNCSAV3URL is the URL identifier for Creative Commons
	// Attribution-NonCommercial-Sharealike 3.0 License.
	CCByNCSAV3URL = "https://creativecommons.org/licenses/by-nc-sa/3.0/"

	// OGLV2URL is the URL identifier for the Open Government Licence version
	// 2.0 License.
	OGLV2URL = "http://www.nationalarchives.gov.uk/doc/open-government-licence/version/2/"

	// OGLV3URL is the URL identifier for the Open Government Licence version
	// 3.0 License.
	OGLV3URL = "http://www.nationalarchives.gov.uk/doc/open-government-licence/version/3/"

	// PDDLV1URL is the URL identifier for the Open Data Commons Public Domain
	// Dedication and License.
	PDDLV1URL = "https://opendatacommons.org/licenses/pddl/1.0/"

	// ODCByV1URL is the URL identifier for the Open Data Commons Attribution
	// License.
	ODCByV1URL = "https://opendatacommons.org/licenses/by/1.0/"

	// ODbLV1URL is a the URL identifier for the Open Data Commons Open Database
	// License.
	ODbLV1URL = "https://opendatacommons.org/licenses/odbl/1.0/"

	// IODLV2URL is the URL identifier for the Italian Open Data License
	IODLV2URL = "https://www.dati.gov.it/content/italian-open-data-license-v20"

	// SGODLV1 is the URL identifier for the Singapore Open Data License
	SGODLV1 = "https://data.gov.sg/open-data-licence"
)

// DataLicense is struct used to describe data license
type DataLicense struct {
	// Human readable name for this license
	Name string `json:"name"`

	// url to the short version of the license
	URL string `json:"id"`

	// url to the long or legal version of the license
	LegalCodeURL string `json:"legalCode,omitempty"`

	// array containing list of properties that this license permits
	Permits []Permission `json:"permits,omitempty"`

	// array containing list of properties that this license requires
	Requires []Requirement `json:"requires,omitempty"`

	// array containing list of properties that this license prohibits
	Prohibits []Prohibition `json:"prohibits,omitempty"`
}

var (
	// licenses is an unexported map of license instances, which allows us to
	// expose a convenience function to instantiate new instances of a standard
	// open license when writing a fetcher without exposing all our vars to the
	// outside world.
	licenses = map[string]DataLicense{
		CC0V1URL: DataLicense{
			Name:         "CC0 1.0 Universal (CC0 1.0) Public Domain Dedication",
			URL:          "https://creativecommons.org/publicdomain/zero/1.0/",
			LegalCodeURL: "https://creativecommons.org/publicdomain/zero/1.0/legalcode",
			Permits: []Permission{
				ReproductionPerm,
				DistributionPerm,
				DerivativeWorksPerm,
				SharingPerm,
			},
			Requires:  []Requirement{},
			Prohibits: []Prohibition{},
		},

		CCByV3URL: DataLicense{
			Name:         "Attribution 3.0 Unported (CC BY 3.0)",
			URL:          "https://creativecommons.org/licenses/by/3.0/",
			LegalCodeURL: "https://creativecommons.org/licenses/by/3.0/legalcode",
			Permits: []Permission{
				ReproductionPerm,
				DistributionPerm,
				DerivativeWorksPerm,
				SharingPerm,
			},
			Requires: []Requirement{
				AttributionReq,
				NoticeReq,
			},
			Prohibits: []Prohibition{},
		},

		CCBySAV4URL: DataLicense{
			Name:         "Attribution-ShareAlike 4.0 International (CC BY-SA 4.0)",
			URL:          "https://creativecommons.org/licenses/by-sa/4.0/",
			LegalCodeURL: "https://creativecommons.org/licenses/by-sa/4.0/legalcode",
			Permits: []Permission{
				ReproductionPerm,
				DistributionPerm,
				DerivativeWorksPerm,
				SharingPerm,
			},
			Requires: []Requirement{
				AttributionReq,
				NoticeReq,
				ShareAlikeReq,
			},
			Prohibits: []Prohibition{},
		},

		CCByV4URL: DataLicense{
			Name:         "Attribution 4.0 International (CC BY 4.0)",
			URL:          "https://creativecommons.org/licenses/by/4.0/",
			LegalCodeURL: "https://creativecommons.org/licenses/by/4.0/legalcode",
			Permits: []Permission{
				ReproductionPerm,
				DistributionPerm,
				DerivativeWorksPerm,
				SharingPerm,
			},
			Requires: []Requirement{
				AttributionReq,
				NoticeReq,
			},
			Prohibits: []Prohibition{},
		},

		CCByNCV3URL: DataLicense{
			Name:         "Attribution-NonCommercial 3.0 Unported (CC BY-NC 3.0)",
			URL:          "https://creativecommons.org/licenses/by-nc/3.0/",
			LegalCodeURL: "https://creativecommons.org/licenses/by-nc/3.0/legalcode",
			Permits: []Permission{
				ReproductionPerm,
				DistributionPerm,
				DerivativeWorksPerm,
				SharingPerm,
			},
			Requires: []Requirement{
				AttributionReq,
				NoticeReq,
			},
			Prohibits: []Prohibition{
				CommercialUseProhib,
			},
		},

		CCByNDV3URL: DataLicense{ // i'm a bit confuse with this one how to describe it, need confirmation with CC
			Name:         "Attribution-NoDerivs 3.0 Unported (CC BY-ND 3.0)",
			URL:          "https://creativecommons.org/licenses/by-nd/3.0/",
			LegalCodeURL: "https://creativecommons.org/licenses/by-nd/3.0/legalcode",
			Permits: []Permission{
				ReproductionPerm,
				DistributionPerm,
			},
			Requires: []Requirement{
				AttributionReq,
				NoticeReq,
			},
			Prohibits: []Prohibition{},
		},

		CCByNCSAV3URL: DataLicense{
			Name:         "Attribution-NonCommercial-ShareAlike 3.0 Unported (CC BY-NC-SA 3.0)",
			URL:          "https://creativecommons.org/licenses/by-nc-sa/3.0/",
			LegalCodeURL: "https://creativecommons.org/licenses/by-nc-sa/3.0/legalcode",
			Permits: []Permission{
				ReproductionPerm,
				DistributionPerm,
				SharingPerm,
				DerivativeWorksPerm,
			},
			Requires: []Requirement{
				AttributionReq,
				ShareAlikeReq,
				NoticeReq,
			},
			Prohibits: []Prohibition{
				CommercialUseProhib,
			},
		},

		OGLV2URL: DataLicense{
			Name:         "Open Government Licence version 2.0",
			URL:          "http://www.nationalarchives.gov.uk/doc/open-government-licence/version/2/",
			LegalCodeURL: "http://www.nationalarchives.gov.uk/doc/open-government-licence/version/2/",
			Permits: []Permission{
				ReproductionPerm,
				DistributionPerm,
				DerivativeWorksPerm,
				SharingPerm,
			},
			Requires: []Requirement{
				AttributionReq,
				NoticeReq,
			},
			Prohibits: []Prohibition{},
		},

		OGLV3URL: DataLicense{
			Name:         "Open Government Licence version 3.0",
			URL:          "http://www.nationalarchives.gov.uk/doc/open-government-licence/version/3/",
			LegalCodeURL: "http://www.nationalarchives.gov.uk/doc/open-government-licence/version/3/",
			Permits: []Permission{
				ReproductionPerm,
				DistributionPerm,
				DerivativeWorksPerm,
				SharingPerm,
			},
			Requires: []Requirement{
				AttributionReq,
				NoticeReq,
			},
			Prohibits: []Prohibition{},
		},

		PDDLV1URL: DataLicense{
			Name:         "Open Data Commons Public Domain Dedication and License (PDDL) v1.0",
			URL:          "https://opendatacommons.org/licenses/pddl/1.0/",
			LegalCodeURL: "https://opendatacommons.org/licenses/pddl/1.0/",
			Permits: []Permission{
				ReproductionPerm,
				DistributionPerm,
				DerivativeWorksPerm,
				SharingPerm,
			},
			Requires:  []Requirement{},
			Prohibits: []Prohibition{},
		},

		ODCByV1URL: DataLicense{
			Name:         "Open Data Commons Attribution License (ODC-By) v1.0",
			URL:          "https://opendatacommons.org/licenses/by/1.0/",
			LegalCodeURL: "https://opendatacommons.org/licenses/by/1.0/",
			Permits: []Permission{
				ReproductionPerm,
				DistributionPerm,
				DerivativeWorksPerm,
				SharingPerm,
			},
			Requires: []Requirement{
				AttributionReq,
				NoticeReq,
			},
			Prohibits: []Prohibition{},
		},

		ODbLV1URL: DataLicense{
			Name:         "Open Data Commons Open Database License (ODbL) v1.0",
			URL:          "https://opendatacommons.org/licenses/odbl/1.0/",
			LegalCodeURL: "https://opendatacommons.org/licenses/odbl/1.0/",
			Permits: []Permission{
				ReproductionPerm,
				DistributionPerm,
				DerivativeWorksPerm,
				SharingPerm,
			},
			Requires: []Requirement{
				AttributionReq,
				NoticeReq,
				ShareAlikeReq,
			},
			Prohibits: []Prohibition{},
		},

		IODLV2URL: DataLicense{
			Name:         "Italian Open Data License v2.0",
			URL:          "https://www.dati.gov.it/content/italian-open-data-license-v20",
			LegalCodeURL: "https://www.dati.gov.it/content/italian-open-data-license-v20",
			Permits: []Permission{
				ReproductionPerm,
				DistributionPerm,
				DerivativeWorksPerm,
				SharingPerm,
			},
			Requires: []Requirement{
				AttributionReq,
			},
			Prohibits: []Prohibition{},
		},

		SGODLV1: DataLicense{
			Name:         "Singapore Open Data License v1.0",
			URL:          "https://data.gov.sg/open-data-licence",
			LegalCodeURL: "https://data.gov.sg/open-data-licence",
			Permits: []Permission{
				ReproductionPerm,
				DistributionPerm,
				DerivativeWorksPerm,
				SharingPerm,
			},
			Requires: []Requirement{
				AttributionReq,
			},
			Prohibits: []Prohibition{},
		},
	}
)

// GetDataLicense is a function that returns a copy of a data license instance
// to make it easier to apply a standard license to a Thing. It takes in a
// license URL as a parameter, and returns an instantiated DataLicense instance
// representing that license or nil if the license is not defined in Thingfulx.
func GetDataLicense(licenseURL string) *DataLicense {
	license, ok := licenses[licenseURL]
	if !ok {
		return nil
	}

	return &license
}
