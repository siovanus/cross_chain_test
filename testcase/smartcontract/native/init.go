package native

import (
	"github.com/ontio/cross_chain_test/testcase/smartcontract/native/side_chain_governance"
)

func TestNative() {
	side_chain_governance.TestGovernanceMethods()
}
