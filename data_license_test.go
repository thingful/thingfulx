package thingfulx

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// what should we test?
// 1. the predefined ones have correct values
// 2. we can add predefined ones to thing
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

func TestLicense(t *testing.T) {
	j, _ := json.Marshal(CCByNcV3)
	fmt.Println(string(j))

	// spew.Dump(CC_0_V1)
	// spew.Dump(CC_BY_NC_V3)
	// spew.Dump(CC_BY_SA_V4)
	// spew.Dump(OGL_V3)
}
