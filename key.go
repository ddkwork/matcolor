// Copyright (c) 2023, The Goki Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on https://github.com/material-foundation/material-color-utilities/blob/main/dart/lib/palettes/core_palette.dart
// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package matcolor

import (
	"image/color"

	"goki.dev/cam/hct"
)

// Key contains the set of key colors used to generate
// a [Scheme] and [Palette]
type Key struct {

	// the primary accent key color
	Primary color.RGBA

	// the secondary accent key color
	Secondary color.RGBA

	// the tertiary accent key color
	Tertiary color.RGBA

	// the error accent key color
	Error color.RGBA

	// the neutral key color used to generate surface and surface container colors
	Neutral color.RGBA

	// the neutral variant key color used to generate surface variant and outline colors
	NeutralVariant color.RGBA

	// an optional map of custom accent key colors
	Custom map[string]color.RGBA
}

// Key returns a new [Key] from the given primary accent key color.
func KeyFromPrimary(primary color.RGBA) *Key {
	k := &Key{}
	p := hct.FromColor(primary)
	p.SetTone(40)

	k.Primary = p.SetChroma(max(p.Chroma, 48)).AsRGBA()
	k.Secondary = p.SetChroma(16).AsRGBA()
	// Material adds 60, but we subtract 60 to get green instead of pink when specifying
	// blue (TODO: is this a good idea, or should we just follow Material?)
	k.Tertiary = p.SetHue(p.Hue - 60).SetChroma(24).AsRGBA()
	k.Error = color.RGBA{179, 38, 30, 255} // #B3261E (Material default error color)
	k.Neutral = p.SetChroma(4).AsRGBA()
	k.NeutralVariant = p.SetChroma(8).AsRGBA()
	return k
}
