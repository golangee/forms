// Copyright 2020 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package forms

import "strings"

// GridAlign
type GridAlign string

const (
	// SpaceEvenly distributes cells evenly in height (vertically) or width (horizontally).
	SpaceEvenly GridAlign = "space-evenly"

	// SpaceBetween distributes the free space equally *between* each item in height (vertically) or
	// (horizontally)
	SpaceBetween GridAlign = "space-between"

	// SpaceAround distributes the free space equally *around* each item in height (vertically) or
	// width (horizontally)
	SpaceAround GridAlign = "space-around"

	// Center aligns cells in the center of the grid container, affects rows (vertically) or
	// columns (horizontally)
	Center GridAlign = "center"

	// Start aligns the cells at the left (horizontally) or top (vertically).
	Start GridAlign = "start"

	// End aligns the cells at the right (horizontally) or bottom (vertically).
	End GridAlign = "end"
)

// GridLayoutParams define how a view spans inside the grid container.
type GridLayoutParams struct {
	Area string
}

// A Grid offers a grid-based layout, where you define the area of children with rows and columns.
type Grid struct {
	*absComponent
}

// NewGrid allocates a new view group, ready to insert views.
func NewGrid() *Grid {
	t := &Grid{}
	t.absComponent = newComponent(t, "div")
	t.node().Style().Set("box-sizing", "border-box")
	t.node().Style().Set("display", "grid")
	return t
}

// SetAreas defines the grid layout "visually" in a two dimensional string array. Example:
//  SetAreas([][]string{
//   {"header", "header", "header"},
//   {"menu", "main", "main"},
//   {"menu", "footer", "footer"},
//  })
//
// Use the Area
func (t *Grid) SetAreas(areas [][]string) *Grid {
	sb := &strings.Builder{}
	for _, row := range areas {
		sb.WriteRune('\'')
		for _, col := range row {
			sb.WriteString(col)
			sb.WriteRune(' ')
		}
		sb.WriteRune('\'')
	}
	t.node().Style().Set("grid-template-areas", sb.String())
	return t
}

// SetColumnWidths defines two things: firstly how many columns at all and secondly how they are calculated.
// You can also use the *Auto*-scalar.
func (t *Grid) SetColumnWidths(scalars ...Scalar) {
	t.node().Style().Set("grid-template-columns", strings.Join(scalarSlice(scalars).toStrings(), " "))
}

// SetRowHeights defines two things: firstly how many rows at all and secondly how they are calculated.
// You can also use the *Auto*-scalar.
func (t *Grid) SetRowHeights(scalars ...Scalar) {
	t.node().Style().Set("grid-template-rows", strings.Join(scalarSlice(scalars).toStrings(), " "))
}

// SetGap defines the margin between all children.
func (t *Grid) SetGap(scalar Scalar) *Grid {
	t.SetColumnGap(scalar)
	t.SetRowGap(scalar)
	return t
}

// SetRowGap defines the row margin between children.
func (t *Grid) SetRowGap(scalar Scalar) *Grid {
	t.node().Style().Set("grid-row-gap", string(scalar)) // non-standard draft
	t.node().Style().Set("row-gap", string(scalar))
	return t
}

// SetColumnGap defines the column margin between children.
func (t *Grid) SetColumnGap(scalar Scalar) *Grid {
	t.node().Style().Set("grid-column-gap", string(scalar)) //non-standard draft?
	t.node().Style().Set("column-gap", string(scalar))

	return t
}

// SetHorizontalAlign determines how columns are distributed inside the grid container
func (t *Grid) SetHorizontalAlign(alignment GridAlign) *Grid {
	t.node().Style().Set("justify-content", string(alignment))
	return t
}

// SetVerticalAlign determines how rows are distributed inside the grid container
func (t *Grid) SetVerticalAlign(alignment GridAlign) *Grid {
	t.node().Style().Set("align-content", string(alignment))
	return t
}

// ClearViews removes all views.
func (t *Grid) ClearViews() ViewGroup {
	return t.RemoveAll()
}

// AppendViews adds all views.
func (t *Grid) AppendViews(views ...View) ViewGroup {
	return t.AddViews(views...)
}

// AddViews adds all views.
func (t *Grid) AddViews(views ...View) *Grid {
	for _, v := range views {
		t.addView(v)
	}
	return t
}

func (t *Grid) AddView(view View, opt GridLayoutParams) *Grid {
	view.node().Style().Update("grid-area", opt.Area)
	t.addView(view)
	return t
}

// RemoveAll removes all views.
func (t *Grid) RemoveAll() *Grid {
	t.absComponent.removeAll()
	return t
}

// Style applies generic style attributes.
func (t *Grid) Style(style ...Style) *Grid {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given reference
func (t *Grid) Self(ref **Grid) *Grid {
	*ref = t
	return t
}
