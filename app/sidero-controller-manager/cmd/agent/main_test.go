// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

//nolint:scopelint,testpackage
package main

import (
	"testing"
)

//nolint:unparam
func TestMapHardwareInformation_DoesNotPanic(t *testing.T) {
	MapHardwareInformation(nil, nil, nil)
}
