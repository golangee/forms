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

import (
	"github.com/golangee/forms/dom"
	"github.com/golangee/forms/event"
	"sort"
	"syscall/js"
)

type Table struct {
	*absComponent
	table              dom.Element
	thead              dom.Element
	tbody              dom.Element
	alignments         map[int]Alignment
	selectedRows       map[int]bool
	nodeHeadRow        dom.Element
	nodeRows           []dom.Element
	checkboxesRows     []*Checkbox
	checkboxHeader     *Checkbox
	showRowSelection   bool
	onSelectionChanged func(t *Table)
	rowClickListener   func(v View, rowIdx int)
}

func NewTable() *Table {
	t := &Table{}
	t.alignments = make(map[int]Alignment)
	t.selectedRows = make(map[int]bool)
	t.absComponent = newComponent(t, "div")
	t.node().AddClass("mdc-data-table")
	t.table = dom.CreateElement("table").AddClass("mdc-data-table__table")
	t.node().AppendChild(t.table)

	t.thead = dom.CreateElement("thead")
	t.table.AppendChild(t.thead)

	t.tbody = dom.CreateElement("tbody").AddClass("mdc-data-table__content")
	t.table.AppendChild(t.tbody)

	//t.addResource(js2.Attach(js2.DataTable, t.node())) //empty tables crashes the foundation, we do it better on our own
	return t
}

func (t *Table) Align(colIdx int, align Alignment) *Table {
	t.alignments[colIdx] = align
	return t
}

func (t *Table) SetHeader(columns ...View) *Table {
	t.thead.SetText("")
	row := dom.CreateElement("tr").AddClass("mdc-data-table__header-row")
	t.nodeHeadRow = row

	for i, col := range columns {
		th := dom.CreateElement("th").AddClass("mdc-data-table__header-cell")
		if a, ok := t.alignments[i]; ok {
			if a == Trailing {
				th.AddClass("mdc-data-table__header-cell--numeric")
			}
		}
		th.SetRole("columnheader").SetScope("col")
		th.AppendChild(col.node())
		col.attach(t) //TODO vs addView?
		row.AppendChild(th)
	}

	t.thead.AppendChild(row)
	return t
}

func (t *Table) selectRow(idx int, selected bool) *Table {
	t.selectedRows[idx] = selected
	if idx < len(t.nodeRows) && idx >= 0 {
		if selected {
			t.nodeRows[idx].AddClass("mdc-data-table__row--selected")
		} else {
			t.nodeRows[idx].RemoveClass("mdc-data-table__row--selected")
		}
	}
	return t
}

func (t *Table) AddRow(columns ...View) *Table {
	row := dom.CreateElement("tr").AddClass("mdc-data-table__row")

	rowIdx := len(t.nodeRows)
	t.nodeRows = append(t.nodeRows, row)
	if t.selectedRows[len(t.nodeRows)-1] {
		row.AddClass("mdc-data-table__row--selected")
	}
	if t.showRowSelection {
		t.applyRow(row, true, len(t.nodeRows)-1)
	}

	for i, col := range columns {
		th := dom.CreateElement("td").AddClass("mdc-data-table__cell")
		if a, ok := t.alignments[i]; ok {
			if a == Trailing {
				th.AddClass("mdc-data-table__header-cell--numeric")
			}
		}
		th.AppendChild(col.node())
		col.attach(t) //TODO vs addView?
		row.AppendChild(th)
	}

	t.tbody.AppendChild(row)

	t.absComponent.addResource(row.AddEventListener(string(event.Click), func(this js.Value, args []js.Value) interface{} {
		if t.rowClickListener != nil {
			t.rowClickListener(t, rowIdx)
		}

		return nil
	}, false))

	return t
}

func (t *Table) getSelectionCell() (*Checkbox, dom.Element) {
	cb := NewCheckbox()
	cb.attach(t)
	cell := dom.CreateElement("td").SetClassName("mdc-data-table__cell wtk-table-select")
	cell.AppendChild(cb.node())
	return cb, cell
}

func (t *Table) RowSelection() bool {
	return t.showRowSelection
}

func (t *Table) SetRowSelection(b bool) *Table {
	t.showRowSelection = b

	if !t.nodeHeadRow.Unwrap().IsUndefined() {
		t.applyRow(t.nodeHeadRow, b, -1)
	}

	for i, row := range t.nodeRows {
		t.applyRow(row, b, i)
	}
	return t
}

func (t *Table) applyRow(row dom.Element, b bool, idx int) {
	for idx > len(t.checkboxesRows)-1 {
		t.checkboxesRows = append(t.checkboxesRows, nil)
	}

	col := row.FirstChild()
	if !col.Unwrap().IsNull() && col.HasClass("wtk-table-select") {
		if !b {
			row.RemoveChild(col)
		}
	} else {
		if b {
			cb, cell := t.getSelectionCell()
			if col.Unwrap().IsNull() {
				row.AppendChild(cell)
			} else {
				row.InsertBefore(cell, col)
			}
			if idx >= 0 {
				t.checkboxesRows[idx] = cb
				cb.SetChecked(t.selectedRows[idx])
				cb.AddChangeListener(func(v *Checkbox) {
					t.selectRow(idx, v.Checked())
					t.updateCheckboxHeaderState()
					t.notifySelectionChanged()
				})
			} else {
				t.checkboxHeader = cb
				t.updateCheckboxHeaderState()
				cb.AddChangeListener(func(v *Checkbox) {
					checked := v.Checked()
					for i := 0; i < len(t.nodeRows); i++ {
						t.selectRow(i, checked)
					}
					for _, cb := range t.checkboxesRows {
						cb.SetChecked(checked)
					}
					t.notifySelectionChanged()
				})
			}

		}
	}
}

func (t *Table) notifySelectionChanged() {
	if t.onSelectionChanged != nil {
		t.onSelectionChanged(t)
	}
}

func (t *Table) updateCheckboxHeaderState() {
	cb := t.checkboxHeader
	selectedCount := len(t.Selected())
	if selectedCount == len(t.nodeRows) {
		cb.SetChecked(true)
		cb.SetIndeterminate(false)
	} else if selectedCount == 0 {
		cb.SetChecked(false)
		cb.SetIndeterminate(false)
	} else {
		cb.SetChecked(false)
		cb.SetIndeterminate(true)
	}
}

// Selected returns the zero based indices of row which are selected
func (t *Table) Selected() []int {
	var res []int
	for k, v := range t.selectedRows {
		if v {
			res = append(res, k)
		}
	}
	sort.Ints(res)
	return res
}

func (t *Table) SetSelected(rows ...int) *Table {
	t.selectedRows = make(map[int]bool)
	for i := 0; i < len(t.nodeRows); i++ {
		t.selectRow(i, false)
	}
	for _, v := range rows {
		t.selectRow(v, true)
	}
	t.notifySelectionChanged()
	return t
}

func (t *Table) SetSelectionChangeListener(f func(t *Table)) *Table {
	t.onSelectionChanged = f
	return t
}

func (t *Table) Style(style ...Style) *Table {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *Table) Self(ref **Table) *Table {
	*ref = t
	return t
}

// SetRowClickListener registers a row click listener, which has nothing to do with the
// row selection.
func (t *Table) SetRowClickListener(f func(v View, rowIdx int)) *Table {
	t.rowClickListener = f
	return t
}
