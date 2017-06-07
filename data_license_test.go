package thingfulx

import (
	// "encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// what should we test?
// 1. the predefined ones have correct values
// 2. we can add predefined ones to thing // this is done inside thing_test.go
// 3. we can add non-predefined ones to thing

func TestCC0(t *testing.T) {

	assert.Equal(t, CC0V1.Name, "CC0 1.0 Universal (CC0 1.0) Public Domain Dedication")
	assert.Equal(t, CC0V1.URL, "https://creativecommons.org/publicdomain/zero/1.0/")
	assert.Equal(t, CC0V1.LegalCodeURL, "https://creativecommons.org/publicdomain/zero/1.0/legalcode")
	// check Permits
	assert.Contains(t, CC0V1.Permits, Reproduction)
	assert.Contains(t, CC0V1.Permits, Distribution)
	assert.Contains(t, CC0V1.Permits, DerivativeWorks)
	assert.Contains(t, CC0V1.Permits, Sharing)

	// check Requires
	assert.NotContains(t, CC0V1.Requires, Notice)
	assert.NotContains(t, CC0V1.Requires, Attribution)
	assert.NotContains(t, CC0V1.Requires, ShareAlike)
	assert.NotContains(t, CC0V1.Requires, SourceCode)
	assert.NotContains(t, CC0V1.Requires, Copyleft)
	assert.NotContains(t, CC0V1.Requires, LesserCopyleft)

	// check Prohibits
	assert.NotContains(t, CC0V1.Prohibits, CommercialUse)
	assert.NotContains(t, CC0V1.Prohibits, HighIncomeNationUse)
}

func TestCCCByV3(t *testing.T) {

	assert.Equal(t, CCByV3.Name, "Attribution 3.0 Unported (CC BY 3.0)")
	assert.Equal(t, CCByV3.URL, "https://creativecommons.org/licenses/by/3.0/")
	assert.Equal(t, CCByV3.LegalCodeURL, "https://creativecommons.org/licenses/by/3.0/legalcode")
	// check Permits
	assert.Contains(t, CCByV3.Permits, Reproduction)
	assert.Contains(t, CCByV3.Permits, Distribution)
	assert.Contains(t, CCByV3.Permits, DerivativeWorks)
	assert.Contains(t, CCByV3.Permits, Sharing)

	// check Requires
	assert.Contains(t, CCByV3.Requires, Attribution)
	assert.NotContains(t, CCByV3.Requires, Notice)
	assert.NotContains(t, CCByV3.Requires, ShareAlike)
	assert.NotContains(t, CCByV3.Requires, SourceCode)
	assert.NotContains(t, CCByV3.Requires, Copyleft)
	assert.NotContains(t, CCByV3.Requires, LesserCopyleft)

	// check Prohibits
	assert.NotContains(t, CCByV3.Prohibits, CommercialUse)
	assert.NotContains(t, CCByV3.Prohibits, HighIncomeNationUse)
}

func TestCustomLicense(t *testing.T) {
	thing := Thing{
		Title:       "Title",
		Description: "Description",
		DataLicense: DataLicense{
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
