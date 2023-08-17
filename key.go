// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matcolor

import "image/color"

// Key contains the set of key colors used to generate
// a [Scheme] and [Palette]
type Key struct {

	// the primary accent key color
	Primary color.RGBA `desc:"the primary accent key color"`

	// the secondary accent key color
	Secondary color.RGBA `desc:"the secondary accent key color"`

	// the tertiary accent key color
	Tertiary color.RGBA `desc:"the tertiary accent key color"`

	// the error accent key color
	Error color.RGBA `desc:"the error accent key color"`

	// the neutral key color used to generate surface and surface container colors
	Neutral color.RGBA `desc:"the neutral key color used to generate surface and surface container colors"`

	// the neutral variant key color used to generate surface variant and outline colors
	NeutralVariant color.RGBA `desc:"the neutral variant key color used to generate surface variant and outline colors"`

	// an optional map of custom accent key colors
	Custom map[string]color.RGBA `desc:"an optional map of custom accent key colors"`
}
