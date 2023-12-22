// Copyright (c) 2023, The Goki Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matcolor

import "image/color"

// Accent contains the 4 variations of a given Base color
type Accent struct {

	// Primary is the base primary color applied to important elements
	Base color.RGBA

	// OnPrimary is the color applied to content on top of Primary. It defaults to the contrast color of Primary.
	On color.RGBA

	// PrimaryContainer is the color applied to elements with less emphasis than Primary
	Container color.RGBA

	// OnPrimaryContainer is the color applied to content on top of PrimaryContainer. It defaults to the contrast color of PrimaryContainer.
	OnContainer color.RGBA
}

func NewAccentLight(tones Tones) Accent {
	return Accent{
		Base:        tones.Tone(40),
		On:          tones.Tone(100),
		Container:   tones.Tone(90),
		OnContainer: tones.Tone(10),
	}
}

func NewAccentDark(tones Tones) Accent {
	return Accent{
		Base:        tones.Tone(80),
		On:          tones.Tone(20),
		Container:   tones.Tone(30),
		OnContainer: tones.Tone(90),
	}
}

// Scheme contains the colors for one color scheme (ex: light or dark).
// To generate a scheme, use [NewScheme].
type Scheme struct {

	// Primary is the primary color applied to important elements
	Primary Accent

	// Secondary is the secondary color applied to less important elements
	Secondary Accent

	// Tertiary is the tertiary color applied as an accent to highlight elements and create contrast between other colors
	Tertiary Accent

	// Error is the error color applied to elements that indicate an error or danger
	Error Accent

	// SurfaceDim is the color applied to elements that will always have the dimmest surface color (see Surface for more information)
	SurfaceDim color.RGBA

	// Surface is the color applied to contained areas, like the background of an app
	Surface color.RGBA

	// SurfaceBright is the color applied to elements that will always have the brightest surface color (see Surface for more information)
	SurfaceBright color.RGBA

	// SurfaceContainerLowest is the color applied to surface container elements that have the lowest emphasis (see SurfaceContainer for more information)
	SurfaceContainerLowest color.RGBA

	// SurfaceContainerLow is the color applied to surface container elements that have lower emphasis (see SurfaceContainer for more information)
	SurfaceContainerLow color.RGBA

	// SurfaceContainer is the color applied to container elements that contrast elements with the surface color
	SurfaceContainer color.RGBA

	// SurfaceContainerHigh is the color applied to surface container elements that have higher emphasis (see SurfaceContainer for more information)
	SurfaceContainerHigh color.RGBA

	// SurfaceContainerHighest is the color applied to surface container elements that have the highest emphasis (see SurfaceContainer for more information)
	SurfaceContainerHighest color.RGBA

	// SurfaceVariant is the color applied to contained areas that contrast standard Surface elements
	SurfaceVariant color.RGBA

	// OnSurface is the color applied to content on top of Surface elements
	OnSurface color.RGBA

	// OnSurfaceVariant is the color applied to content on top of SurfaceVariant elements
	OnSurfaceVariant color.RGBA

	// InverseSurface is the color applied to elements to make them the reverse color of the surrounding elements and create a contrasting effect
	InverseSurface color.RGBA

	// InverseOnSurface is the color applied to content on top of InverseSurface
	InverseOnSurface color.RGBA

	// InversePrimary is the color applied to interactive elements on top of InverseSurface
	InversePrimary color.RGBA

	// Background is the color applied to the background of the app and other low-emphasis areas
	Background color.RGBA

	// OnBackground is the color applied to content on top of Background
	OnBackground color.RGBA

	// Outline is the color applied to borders to create emphasized boundaries that need to have sufficient contrast
	Outline color.RGBA

	// OutlineVariant is the color applied to create decorative boundaries
	OutlineVariant color.RGBA

	// Shadow is the color applied to shadows
	Shadow color.RGBA

	// SurfaceTint is the color applied to tint surfaces
	SurfaceTint color.RGBA

	// Scrim is the color applied to scrims (semi-transparent overlays)
	Scrim color.RGBA

	// // PrimaryFixed is a primary fill color that stays the same regardless of color scheme type (light/dark)
	// PrimaryFixed color.RGBA `desc:"PrimaryFixed is a primary fill color that stays the same regardless of color scheme type (light/dark)"`

	// // PrimaryFixedDim is a higher-emphasis, dimmer primary fill color that stays the same regardless of color scheme type (light/dark)
	// PrimaryFixedDim color.RGBA `desc:"PrimaryFixedDim is a higher-emphasis, dimmer primary fill color that stays the same regardless of color scheme type (light/dark)"`

	// // OnPrimaryFixed is the color applied to high-emphasis content on top of PrimaryFixed
	// OnPrimaryFixed color.RGBA `desc:"OnPrimaryFixed is the color applied to high-emphasis content on top of PrimaryFixed"`

	// // OnPrimaryFixedVariant is the color applied to low-emphasis content on top of PrimaryFixed
	// OnPrimaryFixedVariant color.RGBA `desc:"OnPrimaryFixedVariant is the color applied to low-emphasis content on top of PrimaryFixed"`

	// // SecondaryFixed is a secondary fill color that stays the same regardless of color scheme type (light/dark)
	// SecondaryFixed color.RGBA `desc:"SecondaryFixed is a secondary fill color that stays the same regardless of color scheme type (light/dark)"`

	// // SecondaryFixedDim is a higher-emphasis, dimmer secondary fill color that stays the same regardless of color scheme type (light/dark)
	// SecondaryFixedDim color.RGBA `desc:"SecondaryFixedDim is a higher-emphasis, dimmer secondary fill color that stays the same regardless of color scheme type (light/dark)"`

	// // OnSecondaryFixed is the color applied to high-emphasis content on top of SecondaryFixed
	// OnSecondaryFixed color.RGBA `desc:"OnSecondaryFixed is the color applied to high-emphasis content on top of SecondaryFixed"`

	// // OnSecondaryFixedVariant is the color applied to low-emphasis content on top of SecondaryFixed
	// OnSecondaryFixedVariant color.RGBA `desc:"OnSecondaryFixedVariant is the color applied to low-emphasis content on top of SecondaryFixed"`

	// // TertiaryFixed is a tertiary fill color that stays the same regardless of color scheme type (light/dark)
	// TertiaryFixed color.RGBA `desc:"TertiaryFixed is a tertiary fill color that stays the same regardless of color scheme type (light/dark)"`

	// // TertiaryFixedDim is a higher-emphasis, dimmer tertiary fill color that stays the same regardless of color scheme type (light/dark)
	// TertiaryFixedDim color.RGBA `desc:"TertiaryFixedDim is a higher-emphasis, dimmer tertiary fill color that stays the same regardless of color scheme type (light/dark)"`

	// // OnTertiaryFixed is the color applied to high-emphasis content on top of TertiaryFixed
	// OnTertiaryFixed color.RGBA `desc:"OnTertiaryFixed is the color applied to high-emphasis content on top of TertiaryFixed"`

	// // OnTertiaryFixedVariant is the color applied to low-emphasis content on top of TertiaryFixed
	// OnTertiaryFixedVariant color.RGBA `desc:"OnTertiaryFixedVariant is the color applied to low-emphasis content on top of TertiaryFixed"`
}

// TheScheme is the currently active global color scheme.
var TheScheme *Scheme

// NewLightScheme returns a new light-themed [Scheme]
// based on the given [Palette].
func NewLightScheme(p *Palette) Scheme {
	return Scheme{
		Primary:   NewAccentLight(p.Primary),
		Secondary: NewAccentLight(p.Secondary),
		Tertiary:  NewAccentLight(p.Tertiary),
		Error:     NewAccentLight(p.Error),

		SurfaceDim:    p.Neutral.Tone(87),
		Surface:       p.Neutral.Tone(98),
		SurfaceBright: p.Neutral.Tone(98),

		SurfaceContainerLowest:  p.Neutral.Tone(100),
		SurfaceContainerLow:     p.Neutral.Tone(96),
		SurfaceContainer:        p.Neutral.Tone(94),
		SurfaceContainerHigh:    p.Neutral.Tone(92),
		SurfaceContainerHighest: p.Neutral.Tone(90),

		SurfaceVariant:   p.NeutralVariant.Tone(90),
		OnSurface:        p.NeutralVariant.Tone(10),
		OnSurfaceVariant: p.NeutralVariant.Tone(30),

		InverseSurface:   p.Neutral.Tone(20),
		InverseOnSurface: p.Neutral.Tone(95),
		InversePrimary:   p.Primary.Tone(80),

		Background:   p.Neutral.Tone(98),
		OnBackground: p.Neutral.Tone(10),

		Outline:        p.NeutralVariant.Tone(50),
		OutlineVariant: p.NeutralVariant.Tone(80),

		Shadow:      p.Neutral.Tone(0),
		SurfaceTint: p.Primary.Tone(40),
		Scrim:       p.Neutral.Tone(0),
	}
	// TODO: custom and fixed colors
}

// NewDarkScheme returns a new dark-themed [Scheme]
// based on the given [Palette].
func NewDarkScheme(p *Palette) Scheme {
	return Scheme{
		Primary:   NewAccentDark(p.Primary),
		Secondary: NewAccentDark(p.Secondary),
		Tertiary:  NewAccentDark(p.Tertiary),
		Error:     NewAccentDark(p.Error),

		SurfaceDim:    p.Neutral.Tone(6),
		Surface:       p.Neutral.Tone(6),
		SurfaceBright: p.Neutral.Tone(24),

		SurfaceContainerLowest:  p.Neutral.Tone(4),
		SurfaceContainerLow:     p.Neutral.Tone(10),
		SurfaceContainer:        p.Neutral.Tone(12),
		SurfaceContainerHigh:    p.Neutral.Tone(17),
		SurfaceContainerHighest: p.Neutral.Tone(22),

		SurfaceVariant:   p.NeutralVariant.Tone(30),
		OnSurface:        p.NeutralVariant.Tone(90),
		OnSurfaceVariant: p.NeutralVariant.Tone(80),

		InverseSurface:   p.Neutral.Tone(90),
		InverseOnSurface: p.Neutral.Tone(20),
		InversePrimary:   p.Primary.Tone(40),

		Background:   p.Neutral.Tone(6),
		OnBackground: p.Neutral.Tone(90),

		Outline:        p.NeutralVariant.Tone(60),
		OutlineVariant: p.NeutralVariant.Tone(30),

		Shadow:      p.Neutral.Tone(0),
		SurfaceTint: p.Primary.Tone(80),
		Scrim:       p.Neutral.Tone(0),
	}
	// TODO: custom and fixed colors
}
