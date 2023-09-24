// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matcolor

import "image/color"

// Scheme contains the colors for one color scheme (ex: light or dark).
// To generate a scheme, use [NewScheme].
type Scheme struct {

	// Primary is the base primary color applied to important elements
	Primary color.RGBA `desc:"Primary is the base primary color applied to important elements"`

	// OnPrimary is the color applied to content on top of Primary. It defaults to the contrast color of Primary.
	OnPrimary color.RGBA `desc:"OnPrimary is the color applied to content on top of Primary. It defaults to the contrast color of Primary."`

	// PrimaryContainer is the color applied to elements with less emphasis than Primary
	PrimaryContainer color.RGBA `desc:"PrimaryContainer is the color applied to elements with less emphasis than Primary"`

	// OnPrimaryContainer is the color applied to content on top of PrimaryContainer. It defaults to the contrast color of PrimaryContainer.
	OnPrimaryContainer color.RGBA `desc:"OnPrimaryContainer is the color applied to content on top of PrimaryContainer. It defaults to the contrast color of PrimaryContainer."`

	// Secondary is the base secondary color applied to less important elements
	Secondary color.RGBA `desc:"Secondary is the base secondary color applied to less important elements"`

	// OnSecondary is the color applied to content on top of Secondary. It defaults to the contrast color of Secondary.
	OnSecondary color.RGBA `desc:"OnSecondary is the color applied to content on top of Secondary. It defaults to the contrast color of Secondary."`

	// SecondaryContainer is the color applied to elements with less emphasis than Secondary
	SecondaryContainer color.RGBA `desc:"SecondaryContainer is the color applied to elements with less emphasis than Secondary"`

	// OnSecondaryContainer is the color applied to content on top of SecondaryContainer. It defaults to the contrast color of SecondaryContainer.
	OnSecondaryContainer color.RGBA `desc:"OnSecondaryContainer is the color applied to content on top of SecondaryContainer. It defaults to the contrast color of SecondaryContainer."`

	// Tertiary is the base tertiary color applied as an accent to highlight elements and screate contrast between other colors
	Tertiary color.RGBA `desc:"Tertiary is the base tertiary color applied as an accent to highlight elements and screate contrast between other colors"`

	// OnTertiary is the color applied to content on top of Tertiary. It defaults to the contrast color of Tertiary.
	OnTertiary color.RGBA `desc:"OnTertiary is the color applied to content on top of Tertiary. It defaults to the contrast color of Tertiary."`

	// TertiaryContainer is the color applied to elements with less emphasis than Tertiary
	TertiaryContainer color.RGBA `desc:"TertiaryContainer is the color applied to elements with less emphasis than Tertiary"`

	// OnTertiaryContainer is the color applied to content on top of TertiaryContainer. It defaults to the contrast color of TertiaryContainer.
	OnTertiaryContainer color.RGBA `desc:"OnTertiaryContainer is the color applied to content on top of TertiaryContainer. It defaults to the contrast color of TertiaryContainer."`

	// Error is the base error color applied to elements that indicate an error or danger
	Error color.RGBA `desc:"Error is the base error color applied to elements that indicate an error or danger"`

	// OnError is the color applied to content on top of Error. It defaults to the contrast color of Error.
	OnError color.RGBA `desc:"OnError is the color applied to content on top of Error. It defaults to the contrast color of Error."`

	// ErrorContainer is the color applied to elements with less emphasis than Error
	ErrorContainer color.RGBA `desc:"ErrorContainer is the color applied to elements with less emphasis than Error"`

	// OnErrorContainer is the color applied to content on top of ErrorContainer. It defaults to the contrast color of ErrorContainer.
	OnErrorContainer color.RGBA `desc:"OnErrorContainer is the color applied to content on top of ErrorContainer. It defaults to the contrast color of ErrorContainer."`

	// SurfaceDim is the color applied to elements that will always have the dimmest surface color (see Surface for more information)
	SurfaceDim color.RGBA `desc:"SurfaceDim is the color applied to elements that will always have the dimmest surface color (see Surface for more information)"`

	// Surface is the color applied to contained areas, like the background of an app
	Surface color.RGBA `desc:"Surface is the color applied to contained areas, like the background of an app"`

	// SurfaceBright is the color applied to elements that will always have the brightest surface color (see Surface for more information)
	SurfaceBright color.RGBA `desc:"SurfaceBright is the color applied to elements that will always have the brightest surface color (see Surface for more information)"`

	// SurfaceContainerLowest is the color applied to surface container elements that have the lowest emphasis (see SurfaceContainer for more information)
	SurfaceContainerLowest color.RGBA `desc:"SurfaceContainerLowest is the color applied to surface container elements that have the lowest emphasis (see SurfaceContainer for more information)"`

	// SurfaceContainerLow is the color applied to surface container elements that have lower emphasis (see SurfaceContainer for more information)
	SurfaceContainerLow color.RGBA `desc:"SurfaceContainerLow is the color applied to surface container elements that have lower emphasis (see SurfaceContainer for more information)"`

	// SurfaceContainer is the color applied to container elements that contrast elements with the surface color
	SurfaceContainer color.RGBA `desc:"SurfaceContainer is the color applied to container elements that contrast elements with the surface color"`

	// SurfaceContainerHigh is the color applied to surface container elements that have higher emphasis (see SurfaceContainer for more information)
	SurfaceContainerHigh color.RGBA `desc:"SurfaceContainerHigh is the color applied to surface container elements that have higher emphasis (see SurfaceContainer for more information)"`

	// SurfaceContainerHighest is the color applied to surface container elements that have the highest emphasis (see SurfaceContainer for more information)
	SurfaceContainerHighest color.RGBA `desc:"SurfaceContainerHighest is the color applied to surface container elements that have the highest emphasis (see SurfaceContainer for more information)"`

	// SurfaceVariant is the color applied to contained areas that contrast standard Surface elements
	SurfaceVariant color.RGBA `desc:"SurfaceVariant is the color applied to contained areas that contrast standard Surface elements"`

	// OnSurface is the color applied to content on top of Surface elements
	OnSurface color.RGBA `desc:"OnSurface is the color applied to content on top of Surface elements"`

	// OnSurfaceVariant is the color applied to content on top of SurfaceVariant elements
	OnSurfaceVariant color.RGBA `desc:"OnSurfaceVariant is the color applied to content on top of SurfaceVariant elements"`

	// InverseSurface is the color applied to elements to make them the reverse color of the surrounding elements and create a contrasting effect
	InverseSurface color.RGBA `desc:"InverseSurface is the color applied to elements to make them the reverse color of the surrounding elements and create a contrasting effect"`

	// InverseOnSurface is the color applied to content on top of InverseSurface
	InverseOnSurface color.RGBA `desc:"InverseOnSurface is the color applied to content on top of InverseSurface"`

	// InversePrimary is the color applied to interactive elements on top of InverseSurface
	InversePrimary color.RGBA `desc:"InversePrimary is the color applied to interactive elements on top of InverseSurface"`

	// Background is the color applied to the background of the app and other low-emphasis areas
	Background color.RGBA `desc:"Background is the color applied to the background of the app and other low-emphasis areas"`

	// OnBackground is the color applied to content on top of Background
	OnBackground color.RGBA `desc:"OnBackground is the color applied to content on top of Background"`

	// Outline is the color applied to borders to create emphasized boundaries that need to have sufficient contrast
	Outline color.RGBA `desc:"Outline is the color applied to borders to create emphasized boundaries that need to have sufficient contrast"`

	// OutlineVariant is the color applied to create decorative boundaries
	OutlineVariant color.RGBA `desc:"OutlineVariant is the color applied to create decorative boundaries"`

	// Shadow is the color applied to shadows
	Shadow color.RGBA `desc:"Shadow is the color applied to shadows"`

	// SurfaceTint is the color applied to tint surfaces
	SurfaceTint color.RGBA `desc:"SurfaceTint is the color applied to tint surfaces"`

	// Scrim is the color applied to scrims (semi-transparent overlays)
	Scrim color.RGBA `desc:"Scrim is the color applied to scrims (semi-transparent overlays)"`

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
func NewLightScheme(p Palette) Scheme {
	return Scheme{
		Primary:            p.Primary.Tone(40),
		OnPrimary:          p.Primary.Tone(100),
		PrimaryContainer:   p.Primary.Tone(90),
		OnPrimaryContainer: p.Primary.Tone(10),

		Secondary:            p.Secondary.Tone(40),
		OnSecondary:          p.Secondary.Tone(100),
		SecondaryContainer:   p.Secondary.Tone(90),
		OnSecondaryContainer: p.Secondary.Tone(10),

		Tertiary:            p.Tertiary.Tone(40),
		OnTertiary:          p.Tertiary.Tone(100),
		TertiaryContainer:   p.Tertiary.Tone(90),
		OnTertiaryContainer: p.Tertiary.Tone(10),

		Error:            p.Error.Tone(40),
		OnError:          p.Error.Tone(100),
		ErrorContainer:   p.Error.Tone(90),
		OnErrorContainer: p.Error.Tone(10),

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
func NewDarkScheme(p Palette) Scheme {
	return Scheme{
		Primary:            p.Primary.Tone(80),
		OnPrimary:          p.Primary.Tone(20),
		PrimaryContainer:   p.Primary.Tone(30),
		OnPrimaryContainer: p.Primary.Tone(90),

		Secondary:            p.Secondary.Tone(80),
		OnSecondary:          p.Secondary.Tone(20),
		SecondaryContainer:   p.Secondary.Tone(30),
		OnSecondaryContainer: p.Secondary.Tone(90),

		Tertiary:            p.Tertiary.Tone(80),
		OnTertiary:          p.Tertiary.Tone(20),
		TertiaryContainer:   p.Tertiary.Tone(30),
		OnTertiaryContainer: p.Tertiary.Tone(90),

		Error:            p.Error.Tone(80),
		OnError:          p.Error.Tone(20),
		ErrorContainer:   p.Error.Tone(30),
		OnErrorContainer: p.Error.Tone(90),

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
