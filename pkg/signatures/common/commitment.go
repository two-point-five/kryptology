//
// Copyright Coinbase, Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0
//

package common

import (
	"github.com/two-point-five/kryptology/pkg/core/curves"
)

// Commitment represents a point Pedersen commitment of one or more
// points multiplied by scalars
type Commitment = curves.Point
