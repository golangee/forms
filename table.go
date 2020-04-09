package wtk

import (
	"github.com/worldiety/wtk/dom"
)

type Table struct {
	*absComponent
	table      dom.Element
	thead      dom.Element
	tbody      dom.Element
	alignments map[int]Alignment
}

func NewTable() *Table {
	t := &Table{}
	t.alignments = make(map[int]Alignment)
	t.absComponent = newComponent(t, "div")
	t.node().AddClass("mdc-data-table")
	t.table = dom.CreateElement("table").AddClass("mdc-data-table__table")
	t.node().AppendChild(t.table)

	t.thead = dom.CreateElement("thead")
	t.table.AppendChild(t.thead)

	t.tbody = dom.CreateElement("tbody").AddClass("mdc-data-table__content")
	t.table.AppendChild(t.tbody)

	return t
}

func (t *Table) Align(colIdx int, align Alignment) *Table {
	t.alignments[colIdx] = align
	return t
}

func (t *Table) SetHeader(columns ...View) *Table {
	t.thead.SetText("")
	row := dom.CreateElement("tr").AddClass("mdc-data-table__header-row")

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
func (t *Table) Layout() *Table {
	//t.addResource(js2.Attach(js2.DataTable, t.node())) //TODO this crashes in the foundation
	return t
}

func (t *Table) AddRow(columns ...View) *Table {
	row := dom.CreateElement("tr").AddClass("mdc-data-table__row")

	for i, col := range columns {
		th := dom.CreateElement("th").AddClass("mdc-data-table__cell")
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
