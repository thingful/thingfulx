package thingfulx

import (
	"fmt"
	// "github.com/davecgh/go-spew/spew"
	"encoding/json"
	"testing"
)

func TestLicense(t *testing.T) {
	j, _ := json.Marshal(CC_BY_NC_V3)
	fmt.Println(string(j))

	// spew.Dump(CC_0_V1)
	// spew.Dump(CC_BY_NC_V3)
	// spew.Dump(CC_BY_SA_V4)
	// spew.Dump(OGL_V3)
}
