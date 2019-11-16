package thingfulx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCC0(t *testing.T) {

	license := GetDataLicense(CC0V1URL)
	assert.NotNil(t, license)

	assert.Equal(t, license.Name, "CC0 1.0 Universal (CC0 1.0) Public Domain Dedication")
	assert.Equal(t, license.URL, "https://creativecommons.org/publicdomain/zero/1.0/")
	assert.Equal(t, license.LegalCodeURL, "https://creativecommons.org/publicdomain/zero/1.0/legalcode")
	// check Permits
	assert.Contains(t, license.Permits, ReproductionPerm)
	assert.Contains(t, license.Permits, DistributionPerm)
	assert.Contains(t, license.Permits, DerivativeWorksPerm)
	assert.Contains(t, license.Permits, SharingPerm)

	// check Requires
	assert.NotContains(t, license.Requires, NoticeReq)
	assert.NotContains(t, license.Requires, AttributionReq)
	assert.NotContains(t, license.Requires, ShareAlikeReq)
	assert.NotContains(t, license.Requires, SourceCodeReq)
	assert.NotContains(t, license.Requires, CopyleftReq)
	assert.NotContains(t, license.Requires, LesserCopyleftReq)

	// check Prohibits
	assert.NotContains(t, license.Prohibits, CommercialUseProhib)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUseProhib)
}

func TestCCByV3(t *testing.T) {

	license := GetDataLicense(CCByV3URL)

	assert.Equal(t, license.Name, "Attribution 3.0 Unported (CC BY 3.0)")
	assert.Equal(t, license.URL, "https://creativecommons.org/licenses/by/3.0/")
	assert.Equal(t, license.LegalCodeURL, "https://creativecommons.org/licenses/by/3.0/legalcode")

	// check Permits
	assert.Contains(t, license.Permits, ReproductionPerm)
	assert.Contains(t, license.Permits, DistributionPerm)
	assert.Contains(t, license.Permits, DerivativeWorksPerm)
	assert.Contains(t, license.Permits, SharingPerm)

	// check Requires
	assert.Contains(t, license.Requires, AttributionReq)
	assert.Contains(t, license.Requires, NoticeReq)
	assert.NotContains(t, license.Requires, ShareAlikeReq)
	assert.NotContains(t, license.Requires, SourceCodeReq)
	assert.NotContains(t, license.Requires, CopyleftReq)
	assert.NotContains(t, license.Requires, LesserCopyleftReq)

	// check Prohibits
	assert.NotContains(t, license.Prohibits, CommercialUseProhib)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUseProhib)
}

func TestCCBySAV3(t *testing.T) {

	license := GetDataLicense(CCBySAV4URL)

	assert.Equal(t, license.Name, "Attribution-ShareAlike 4.0 International (CC BY-SA 4.0)")
	assert.Equal(t, license.URL, "https://creativecommons.org/licenses/by-sa/4.0/")
	assert.Equal(t, license.LegalCodeURL, "https://creativecommons.org/licenses/by-sa/4.0/legalcode")
	// check Permits
	assert.Contains(t, license.Permits, ReproductionPerm)
	assert.Contains(t, license.Permits, DistributionPerm)
	assert.Contains(t, license.Permits, DerivativeWorksPerm)
	assert.Contains(t, license.Permits, SharingPerm)

	// check Requires
	assert.Contains(t, license.Requires, AttributionReq)
	assert.Contains(t, license.Requires, NoticeReq)
	assert.Contains(t, license.Requires, ShareAlikeReq)
	assert.NotContains(t, license.Requires, SourceCodeReq)
	assert.NotContains(t, license.Requires, CopyleftReq)
	assert.NotContains(t, license.Requires, LesserCopyleftReq)

	// check Prohibits
	assert.NotContains(t, license.Prohibits, CommercialUseProhib)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUseProhib)
}

func TestCCByNCV3(t *testing.T) {

	license := GetDataLicense(CCByNCV3URL)

	assert.Equal(t, license.Name, "Attribution-NonCommercial 3.0 Unported (CC BY-NC 3.0)")
	assert.Equal(t, license.URL, "https://creativecommons.org/licenses/by-nc/3.0/")
	assert.Equal(t, license.LegalCodeURL, "https://creativecommons.org/licenses/by-nc/3.0/legalcode")
	// check Permits
	assert.Contains(t, license.Permits, ReproductionPerm)
	assert.Contains(t, license.Permits, DistributionPerm)
	assert.Contains(t, license.Permits, DerivativeWorksPerm)
	assert.Contains(t, license.Permits, SharingPerm)

	// check Requires
	assert.Contains(t, license.Requires, AttributionReq)
	assert.Contains(t, license.Requires, NoticeReq)
	assert.NotContains(t, license.Requires, ShareAlikeReq)
	assert.NotContains(t, license.Requires, SourceCodeReq)
	assert.NotContains(t, license.Requires, CopyleftReq)
	assert.NotContains(t, license.Requires, LesserCopyleftReq)

	// check Prohibits
	assert.Contains(t, license.Prohibits, CommercialUseProhib)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUseProhib)
}

func TestCCByNDV3(t *testing.T) {

	license := GetDataLicense(CCByNDV3URL)

	assert.Equal(t, license.Name, "Attribution-NoDerivs 3.0 Unported (CC BY-ND 3.0)")
	assert.Equal(t, license.URL, "https://creativecommons.org/licenses/by-nd/3.0/")
	assert.Equal(t, license.LegalCodeURL, "https://creativecommons.org/licenses/by-nd/3.0/legalcode")
	// check Permits
	assert.Contains(t, license.Permits, ReproductionPerm)
	assert.Contains(t, license.Permits, DistributionPerm)
	assert.NotContains(t, license.Permits, DerivativeWorksPerm)
	assert.NotContains(t, license.Permits, SharingPerm)

	// check Requires
	assert.Contains(t, license.Requires, AttributionReq)
	assert.Contains(t, license.Requires, NoticeReq)
	assert.NotContains(t, license.Requires, ShareAlikeReq)
	assert.NotContains(t, license.Requires, SourceCodeReq)
	assert.NotContains(t, license.Requires, CopyleftReq)
	assert.NotContains(t, license.Requires, LesserCopyleftReq)

	// check Prohibits
	assert.NotContains(t, license.Prohibits, CommercialUseProhib)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUseProhib)
}

func TestCCByNcSaV3(t *testing.T) {

	license := GetDataLicense(CCByNCSAV3URL)

	assert.Equal(t, license.Name, "Attribution-NonCommercial-ShareAlike 3.0 Unported (CC BY-NC-SA 3.0)")
	assert.Equal(t, license.URL, "https://creativecommons.org/licenses/by-nc-sa/3.0/")
	assert.Equal(t, license.LegalCodeURL, "https://creativecommons.org/licenses/by-nc-sa/3.0/legalcode")
	// check Permits
	assert.Contains(t, license.Permits, ReproductionPerm)
	assert.Contains(t, license.Permits, DistributionPerm)
	assert.Contains(t, license.Permits, DerivativeWorksPerm)
	assert.Contains(t, license.Permits, SharingPerm)

	// check Requires
	assert.Contains(t, license.Requires, AttributionReq)
	assert.Contains(t, license.Requires, NoticeReq)
	assert.Contains(t, license.Requires, ShareAlikeReq)
	assert.NotContains(t, license.Requires, SourceCodeReq)
	assert.NotContains(t, license.Requires, CopyleftReq)
	assert.NotContains(t, license.Requires, LesserCopyleftReq)

	// check Prohibits
	assert.Contains(t, license.Prohibits, CommercialUseProhib)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUseProhib)
}

func TestOGLV2(t *testing.T) {

	license := GetDataLicense(OGLV2URL)

	assert.Equal(t, license.Name, "Open Government Licence version 2.0")
	assert.Equal(t, license.URL, "http://www.nationalarchives.gov.uk/doc/open-government-licence/version/2/")
	assert.Equal(t, license.LegalCodeURL, "http://www.nationalarchives.gov.uk/doc/open-government-licence/version/2/")
	// check Permits
	assert.Contains(t, license.Permits, ReproductionPerm)
	assert.Contains(t, license.Permits, DistributionPerm)
	assert.Contains(t, license.Permits, DerivativeWorksPerm)
	assert.Contains(t, license.Permits, SharingPerm)

	// check Requires
	assert.Contains(t, license.Requires, AttributionReq)
	assert.Contains(t, license.Requires, NoticeReq)
	assert.NotContains(t, license.Requires, ShareAlikeReq)
	assert.NotContains(t, license.Requires, SourceCodeReq)
	assert.NotContains(t, license.Requires, CopyleftReq)
	assert.NotContains(t, license.Requires, LesserCopyleftReq)

	// check Prohibits
	assert.NotContains(t, license.Prohibits, CommercialUseProhib)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUseProhib)
}

func TestOGLV3(t *testing.T) {

	license := GetDataLicense(OGLV3URL)

	assert.Equal(t, license.Name, "Open Government Licence version 3.0")
	assert.Equal(t, license.URL, "http://www.nationalarchives.gov.uk/doc/open-government-licence/version/3/")
	assert.Equal(t, license.LegalCodeURL, "http://www.nationalarchives.gov.uk/doc/open-government-licence/version/3/")
	// check Permits
	assert.Contains(t, license.Permits, ReproductionPerm)
	assert.Contains(t, license.Permits, DistributionPerm)
	assert.Contains(t, license.Permits, DerivativeWorksPerm)
	assert.Contains(t, license.Permits, SharingPerm)

	// check Requires
	assert.Contains(t, license.Requires, AttributionReq)
	assert.Contains(t, license.Requires, NoticeReq)
	assert.NotContains(t, license.Requires, ShareAlikeReq)
	assert.NotContains(t, license.Requires, SourceCodeReq)
	assert.NotContains(t, license.Requires, CopyleftReq)
	assert.NotContains(t, license.Requires, LesserCopyleftReq)

	// check Prohibits
	assert.NotContains(t, license.Prohibits, CommercialUseProhib)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUseProhib)
}

func TestPDDLV1(t *testing.T) {

	license := GetDataLicense(PDDLV1URL)

	assert.Equal(t, license.Name, "Open Data Commons Public Domain Dedication and License (PDDL) v1.0")
	assert.Equal(t, license.URL, "https://opendatacommons.org/licenses/pddl/1.0/")
	assert.Equal(t, license.LegalCodeURL, "https://opendatacommons.org/licenses/pddl/1.0/")
	// check Permits
	assert.Contains(t, license.Permits, ReproductionPerm)
	assert.Contains(t, license.Permits, DistributionPerm)
	assert.Contains(t, license.Permits, DerivativeWorksPerm)
	assert.Contains(t, license.Permits, SharingPerm)

	// check Requires
	assert.NotContains(t, license.Requires, AttributionReq)
	assert.NotContains(t, license.Requires, NoticeReq)
	assert.NotContains(t, license.Requires, ShareAlikeReq)
	assert.NotContains(t, license.Requires, SourceCodeReq)
	assert.NotContains(t, license.Requires, CopyleftReq)
	assert.NotContains(t, license.Requires, LesserCopyleftReq)

	// check Prohibits
	assert.NotContains(t, license.Prohibits, CommercialUseProhib)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUseProhib)
}

func TestODCByV1(t *testing.T) {

	license := GetDataLicense(ODCByV1URL)

	assert.Equal(t, license.Name, "Open Data Commons Attribution License (ODC-By) v1.0")
	assert.Equal(t, license.URL, "https://opendatacommons.org/licenses/by/1.0/")
	assert.Equal(t, license.LegalCodeURL, "https://opendatacommons.org/licenses/by/1.0/")
	// check Permits
	assert.Contains(t, license.Permits, ReproductionPerm)
	assert.Contains(t, license.Permits, DistributionPerm)
	assert.Contains(t, license.Permits, DerivativeWorksPerm)
	assert.Contains(t, license.Permits, SharingPerm)

	// check Requires
	assert.Contains(t, license.Requires, AttributionReq)
	assert.Contains(t, license.Requires, NoticeReq)
	assert.NotContains(t, license.Requires, ShareAlikeReq)
	assert.NotContains(t, license.Requires, SourceCodeReq)
	assert.NotContains(t, license.Requires, CopyleftReq)
	assert.NotContains(t, license.Requires, LesserCopyleftReq)

	// check Prohibits
	assert.NotContains(t, license.Prohibits, CommercialUseProhib)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUseProhib)
}

func TestODbLV1(t *testing.T) {

	license := GetDataLicense(ODbLV1URL)

	assert.Equal(t, license.Name, "Open Data Commons Open Database License (ODbL) v1.0")
	assert.Equal(t, license.URL, "https://opendatacommons.org/licenses/odbl/1.0/")
	assert.Equal(t, license.LegalCodeURL, "https://opendatacommons.org/licenses/odbl/1.0/")
	// check Permits
	assert.Contains(t, license.Permits, ReproductionPerm)
	assert.Contains(t, license.Permits, DistributionPerm)
	assert.Contains(t, license.Permits, DerivativeWorksPerm)
	assert.Contains(t, license.Permits, SharingPerm)

	// check Requires
	assert.Contains(t, license.Requires, AttributionReq)
	assert.Contains(t, license.Requires, NoticeReq)
	assert.Contains(t, license.Requires, ShareAlikeReq)
	assert.NotContains(t, license.Requires, SourceCodeReq)
	assert.NotContains(t, license.Requires, CopyleftReq)
	assert.NotContains(t, license.Requires, LesserCopyleftReq)

	// check Prohibits
	assert.NotContains(t, license.Prohibits, CommercialUseProhib)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUseProhib)
}

func TestIODLV2URL(t *testing.T) {

	license := GetDataLicense(IODLV2URL)

	assert.Equal(t, license.Name, "Italian Open Data License v2.0")
	assert.Equal(t, license.URL, "https://www.dati.gov.it/content/italian-open-data-license-v20")
	assert.Equal(t, license.LegalCodeURL, "https://www.dati.gov.it/content/italian-open-data-license-v20")
	// check Permits
	assert.Contains(t, license.Permits, ReproductionPerm)
	assert.Contains(t, license.Permits, DistributionPerm)
	assert.Contains(t, license.Permits, DerivativeWorksPerm)
	assert.Contains(t, license.Permits, SharingPerm)

	// check Requires
	assert.Contains(t, license.Requires, AttributionReq)

	// check Prohibits
	assert.NotContains(t, license.Prohibits, CommercialUseProhib)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUseProhib)
}

func TestSGODLV1(t *testing.T) {

	license := GetDataLicense(SGODLV1)

	assert.Equal(t, license.Name, "Singapore Open Data License v1.0")
	assert.Equal(t, license.URL, "https://data.gov.sg/open-data-licence")
	assert.Equal(t, license.LegalCodeURL, "https://data.gov.sg/open-data-licence")

	// check Permits
	assert.Contains(t, license.Permits, ReproductionPerm)
	assert.Contains(t, license.Permits, DistributionPerm)
	assert.Contains(t, license.Permits, DerivativeWorksPerm)
	assert.Contains(t, license.Permits, SharingPerm)

	// check Requires
	assert.Contains(t, license.Requires, AttributionReq)

	// check Prohibits
	assert.NotContains(t, license.Prohibits, CommercialUseProhib)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUseProhib)
}

func TestCustomLicense(t *testing.T) {
	channel := Channel{
		Title:       "Title",
		Description: "Description",
		DataLicense: &DataLicense{
			Name: "custom made data license",
			URL:  "http://some/url/to/the/license",
			Permits: []Permission{
				ReproductionPerm,
				SharingPerm,
			},
			Requires: []Requirement{
				AttributionReq,
			},
		},
	}

	assert.Equal(t, channel.DataLicense.Name, "custom made data license")
	assert.Equal(t, channel.DataLicense.URL, "http://some/url/to/the/license")
	assert.Contains(t, channel.DataLicense.Permits, ReproductionPerm)
	assert.Contains(t, channel.DataLicense.Permits, SharingPerm)
	assert.Contains(t, channel.DataLicense.Requires, AttributionReq)
	assert.NotContains(t, channel.DataLicense.Prohibits, CommercialUseProhib)
}

func TestGetDataLicense(t *testing.T) {
	license := GetDataLicense("unknown-license")
	assert.Nil(t, license)
}

func TestUnableToModifyReadLicenses(t *testing.T) {
	license := GetDataLicense(PDDLV1URL)
	license.Name = "foo"

	license2 := GetDataLicense(PDDLV1URL)
	assert.NotEqual(t, "foo", license2.Name)
}
