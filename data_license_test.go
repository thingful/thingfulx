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
	assert.Contains(t, license.Permits, Reproduction)
	assert.Contains(t, license.Permits, Distribution)
	assert.Contains(t, license.Permits, DerivativeWorks)
	assert.Contains(t, license.Permits, Sharing)

	// check Requires
	assert.NotContains(t, license.Requires, Notice)
	assert.NotContains(t, license.Requires, Attribution)
	assert.NotContains(t, license.Requires, ShareAlike)
	assert.NotContains(t, license.Requires, SourceCode)
	assert.NotContains(t, license.Requires, Copyleft)
	assert.NotContains(t, license.Requires, LesserCopyleft)

	// check Prohibits
	assert.NotContains(t, license.Prohibits, CommercialUse)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUse)
}

func TestCCByV3(t *testing.T) {

	license := GetDataLicense(CCByV3URL)

	assert.Equal(t, license.Name, "Attribution 3.0 Unported (CC BY 3.0)")
	assert.Equal(t, license.URL, "https://creativecommons.org/licenses/by/3.0/")
	assert.Equal(t, license.LegalCodeURL, "https://creativecommons.org/licenses/by/3.0/legalcode")
	// check Permits
	assert.Contains(t, license.Permits, Reproduction)
	assert.Contains(t, license.Permits, Distribution)
	assert.Contains(t, license.Permits, DerivativeWorks)
	assert.Contains(t, license.Permits, Sharing)

	// check Requires
	assert.Contains(t, license.Requires, Attribution)
	assert.Contains(t, license.Requires, Notice)
	assert.NotContains(t, license.Requires, ShareAlike)
	assert.NotContains(t, license.Requires, SourceCode)
	assert.NotContains(t, license.Requires, Copyleft)
	assert.NotContains(t, license.Requires, LesserCopyleft)

	// check Prohibits
	assert.NotContains(t, license.Prohibits, CommercialUse)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUse)
}

func TestCCBySAV3(t *testing.T) {

	license := GetDataLicense(CCBySAV4URL)

	assert.Equal(t, license.Name, "Attribution-ShareAlike 4.0 International (CC BY-SA 4.0)")
	assert.Equal(t, license.URL, "https://creativecommons.org/licenses/by-sa/4.0/")
	assert.Equal(t, license.LegalCodeURL, "https://creativecommons.org/licenses/by-sa/4.0/legalcode")
	// check Permits
	assert.Contains(t, license.Permits, Reproduction)
	assert.Contains(t, license.Permits, Distribution)
	assert.Contains(t, license.Permits, DerivativeWorks)
	assert.Contains(t, license.Permits, Sharing)

	// check Requires
	assert.Contains(t, license.Requires, Attribution)
	assert.Contains(t, license.Requires, Notice)
	assert.Contains(t, license.Requires, ShareAlike)
	assert.NotContains(t, license.Requires, SourceCode)
	assert.NotContains(t, license.Requires, Copyleft)
	assert.NotContains(t, license.Requires, LesserCopyleft)

	// check Prohibits
	assert.NotContains(t, license.Prohibits, CommercialUse)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUse)
}

func TestCCByNCV3(t *testing.T) {

	license := GetDataLicense(CCByNCV3URL)

	assert.Equal(t, license.Name, "Attribution-NonCommercial 3.0 Unported (CC BY-NC 3.0)")
	assert.Equal(t, license.URL, "https://creativecommons.org/licenses/by-nc/3.0/")
	assert.Equal(t, license.LegalCodeURL, "https://creativecommons.org/licenses/by-nc/3.0/legalcode")
	// check Permits
	assert.Contains(t, license.Permits, Reproduction)
	assert.Contains(t, license.Permits, Distribution)
	assert.Contains(t, license.Permits, DerivativeWorks)
	assert.Contains(t, license.Permits, Sharing)

	// check Requires
	assert.Contains(t, license.Requires, Attribution)
	assert.Contains(t, license.Requires, Notice)
	assert.NotContains(t, license.Requires, ShareAlike)
	assert.NotContains(t, license.Requires, SourceCode)
	assert.NotContains(t, license.Requires, Copyleft)
	assert.NotContains(t, license.Requires, LesserCopyleft)

	// check Prohibits
	assert.Contains(t, license.Prohibits, CommercialUse)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUse)
}

func TestCCByNDV3(t *testing.T) {

	license := GetDataLicense(CCByNDV3URL)

	assert.Equal(t, license.Name, "Attribution-NoDerivs 3.0 Unported (CC BY-ND 3.0)")
	assert.Equal(t, license.URL, "https://creativecommons.org/licenses/by-nd/3.0/")
	assert.Equal(t, license.LegalCodeURL, "https://creativecommons.org/licenses/by-nd/3.0/legalcode")
	// check Permits
	assert.Contains(t, license.Permits, Reproduction)
	assert.Contains(t, license.Permits, Distribution)
	assert.NotContains(t, license.Permits, DerivativeWorks)
	assert.NotContains(t, license.Permits, Sharing)

	// check Requires
	assert.Contains(t, license.Requires, Attribution)
	assert.Contains(t, license.Requires, Notice)
	assert.NotContains(t, license.Requires, ShareAlike)
	assert.NotContains(t, license.Requires, SourceCode)
	assert.NotContains(t, license.Requires, Copyleft)
	assert.NotContains(t, license.Requires, LesserCopyleft)

	// check Prohibits
	assert.NotContains(t, license.Prohibits, CommercialUse)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUse)
}

func TestCCByNcSaV3(t *testing.T) {

	license := GetDataLicense(CCByNCSAV3URL)

	assert.Equal(t, license.Name, "Attribution-NonCommercial-ShareAlike 3.0 Unported (CC BY-NC-SA 3.0)")
	assert.Equal(t, license.URL, "https://creativecommons.org/licenses/by-nc-sa/3.0/")
	assert.Equal(t, license.LegalCodeURL, "https://creativecommons.org/licenses/by-nc-sa/3.0/legalcode")
	// check Permits
	assert.Contains(t, license.Permits, Reproduction)
	assert.Contains(t, license.Permits, Distribution)
	assert.Contains(t, license.Permits, DerivativeWorks)
	assert.Contains(t, license.Permits, Sharing)

	// check Requires
	assert.Contains(t, license.Requires, Attribution)
	assert.Contains(t, license.Requires, Notice)
	assert.Contains(t, license.Requires, ShareAlike)
	assert.NotContains(t, license.Requires, SourceCode)
	assert.NotContains(t, license.Requires, Copyleft)
	assert.NotContains(t, license.Requires, LesserCopyleft)

	// check Prohibits
	assert.Contains(t, license.Prohibits, CommercialUse)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUse)
}

func TestOGLV3(t *testing.T) {

	license := GetDataLicense(OGLV3URL)

	assert.Equal(t, license.Name, "Open Government Licence version 3.0")
	assert.Equal(t, license.URL, "http://www.nationalarchives.gov.uk/doc/open-government-licence/version/3/")
	assert.Equal(t, license.LegalCodeURL, "http://www.nationalarchives.gov.uk/doc/open-government-licence/version/3/")
	// check Permits
	assert.Contains(t, license.Permits, Reproduction)
	assert.Contains(t, license.Permits, Distribution)
	assert.Contains(t, license.Permits, DerivativeWorks)
	assert.Contains(t, license.Permits, Sharing)

	// check Requires
	assert.Contains(t, license.Requires, Attribution)
	assert.Contains(t, license.Requires, Notice)
	assert.NotContains(t, license.Requires, ShareAlike)
	assert.NotContains(t, license.Requires, SourceCode)
	assert.NotContains(t, license.Requires, Copyleft)
	assert.NotContains(t, license.Requires, LesserCopyleft)

	// check Prohibits
	assert.NotContains(t, license.Prohibits, CommercialUse)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUse)
}

func TestPDDLV1(t *testing.T) {

	license := GetDataLicense(PDDLV1URL)

	assert.Equal(t, license.Name, "Open Data Commons Public Domain Dedication and License (PDDL) v1.0")
	assert.Equal(t, license.URL, "https://opendatacommons.org/licenses/pddl/1.0/")
	assert.Equal(t, license.LegalCodeURL, "https://opendatacommons.org/licenses/pddl/1.0/")
	// check Permits
	assert.Contains(t, license.Permits, Reproduction)
	assert.Contains(t, license.Permits, Distribution)
	assert.Contains(t, license.Permits, DerivativeWorks)
	assert.Contains(t, license.Permits, Sharing)

	// check Requires
	assert.NotContains(t, license.Requires, Attribution)
	assert.NotContains(t, license.Requires, Notice)
	assert.NotContains(t, license.Requires, ShareAlike)
	assert.NotContains(t, license.Requires, SourceCode)
	assert.NotContains(t, license.Requires, Copyleft)
	assert.NotContains(t, license.Requires, LesserCopyleft)

	// check Prohibits
	assert.NotContains(t, license.Prohibits, CommercialUse)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUse)
}

func TestODCByV1(t *testing.T) {

	license := GetDataLicense(ODCByV1URL)

	assert.Equal(t, license.Name, "Open Data Commons Attribution License (ODC-By) v1.0")
	assert.Equal(t, license.URL, "https://opendatacommons.org/licenses/by/1.0/")
	assert.Equal(t, license.LegalCodeURL, "https://opendatacommons.org/licenses/by/1.0/")
	// check Permits
	assert.Contains(t, license.Permits, Reproduction)
	assert.Contains(t, license.Permits, Distribution)
	assert.Contains(t, license.Permits, DerivativeWorks)
	assert.Contains(t, license.Permits, Sharing)

	// check Requires
	assert.Contains(t, license.Requires, Attribution)
	assert.Contains(t, license.Requires, Notice)
	assert.NotContains(t, license.Requires, ShareAlike)
	assert.NotContains(t, license.Requires, SourceCode)
	assert.NotContains(t, license.Requires, Copyleft)
	assert.NotContains(t, license.Requires, LesserCopyleft)

	// check Prohibits
	assert.NotContains(t, license.Prohibits, CommercialUse)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUse)
}

func TestODbLV1(t *testing.T) {

	license := GetDataLicense(ODbLV1URL)

	assert.Equal(t, license.Name, "Open Data Commons Open Database License (ODbL) v1.0")
	assert.Equal(t, license.URL, "https://opendatacommons.org/licenses/odbl/1.0/")
	assert.Equal(t, license.LegalCodeURL, "https://opendatacommons.org/licenses/odbl/1.0/")
	// check Permits
	assert.Contains(t, license.Permits, Reproduction)
	assert.Contains(t, license.Permits, Distribution)
	assert.Contains(t, license.Permits, DerivativeWorks)
	assert.Contains(t, license.Permits, Sharing)

	// check Requires
	assert.Contains(t, license.Requires, Attribution)
	assert.Contains(t, license.Requires, Notice)
	assert.Contains(t, license.Requires, ShareAlike)
	assert.NotContains(t, license.Requires, SourceCode)
	assert.NotContains(t, license.Requires, Copyleft)
	assert.NotContains(t, license.Requires, LesserCopyleft)

	// check Prohibits
	assert.NotContains(t, license.Prohibits, CommercialUse)
	assert.NotContains(t, license.Prohibits, HighIncomeNationUse)
}

func TestCustomLicense(t *testing.T) {
	thing := Thing{
		Title:       "Title",
		Description: "Description",
		DataLicense: &DataLicense{
			Name: "custom made data license",
			URL:  "http://some/url/to/the/license",
			Permits: []string{
				Reproduction,
				Sharing,
			},
			Requires: []string{
				Attribution,
			},
		},
	}

	assert.Equal(t, thing.DataLicense.Name, "custom made data license")
	assert.Equal(t, thing.DataLicense.URL, "http://some/url/to/the/license")
	assert.Contains(t, thing.DataLicense.Permits, Reproduction)
	assert.Contains(t, thing.DataLicense.Permits, Sharing)
	assert.Contains(t, thing.DataLicense.Requires, Attribution)
	assert.NotContains(t, thing.DataLicense.Prohibits, CommercialUse)
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
