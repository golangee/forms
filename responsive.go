package forms

import (
	"github.com/golangee/forms/dom"
	"strings"
	"syscall/js"
)

// MediaCriteria describes rules to match against various screen properties.
type MediaCriteria string

// MatchMaxWidth returns a criteria for a max width
func MatchMaxWidth(scalar Scalar) MediaCriteria {
	return MediaCriteria("(max-width: " + scalar + ")")
}

// MatchMaxWidth returns a criteria for a max height
func MatchMaxHeight(scalar Scalar) MediaCriteria {
	return MediaCriteria("(max-height: " + scalar + ")")
}

// MatchMinWidth returns a criteria for a min width
func MatchMinWidth(scalar Scalar) MediaCriteria {
	return MediaCriteria("(min-width: " + scalar + ")")
}

// MatchMinHeight returns a criteria for a min height
func MatchMinHeight(scalar Scalar) MediaCriteria {
	return MediaCriteria("(min-height: " + scalar + ")")
}

// MatchOne concats all given criterias into a big or statement
func MatchOne(criterias ...MediaCriteria) MediaCriteria {
	sb := &strings.Builder{}
	for i, c := range criterias {
		sb.WriteString(string(c))
		if i < len(criterias)-1 {
			sb.WriteString(" , ")
		}
	}
	return MediaCriteria(sb.String())
}

// MatchAll concats all given criterias into a big and statement
func MatchAll(criterias ...MediaCriteria) MediaCriteria {
	sb := &strings.Builder{}
	for i, c := range criterias {
		sb.WriteString(string(c))
		if i < len(criterias)-1 {
			sb.WriteString(" and ")
		}
	}
	return MediaCriteria(sb.String())
}

// MatchLandscape matches for a screen where the width is longer than the height
func MatchLandscape() MediaCriteria {
	return "orientation: landscape"
}

// MatchPortrait matches for a screen where the height is longer than the width
func MatchPortrait() MediaCriteria {
	return "orientation: portrait"
}

type criteriaAndCallback struct {
	criteria MediaCriteria
	onMatch  func(view View)
	active   bool
}

// MediaMatcher registers multiple matching and a non-matching callback a number of criterias to the view.
// If the view is destroyed, the callbacks are unregistered.
type MediaMatcher struct {
	view      View
	criterias []*criteriaAndCallback
	noMatch   func(view View)
}

func NewMediaMatcher(scope View, noMatch func(view View)) *MediaMatcher {
	m := &MediaMatcher{}
	m.view = scope
	m.noMatch = noMatch
	m.view.addResource(dom.GetWindow().AddEventListener("resize", func(this js.Value, args []js.Value) interface{} {
		m.Check()
		return nil
	}, false))

	return m
}

// Add appends another criteria to evaluate. The given callback is invoked, if it matches the criteria. Only
// the first registered and matching criteria is invoked.
func (m *MediaMatcher) Add(criteria MediaCriteria, onMatch func(view View)) {
	m.criterias = append(m.criterias, &criteriaAndCallback{
		criteria: criteria,
		onMatch:  onMatch,
		active:   false,
	})
}

// Check evaluates all registered criterias and applies the first matching callback. It only invokes a callback
// if it has not been active. The noMatch callback is only called, if no active criteria match has been found.
// It will be called internally automatically, if some environment conditions like the screen size has changed.
// It is safe to be invoked on a nil pointer.
func (m *MediaMatcher) Check() {
	if m == nil {
		return
	}

	anyMatch := false
	for idx, c := range m.criterias {
		if dom.GetWindow().MatchesMedia(string(c.criteria)) {
			anyMatch = true
			if !c.active {
				m.setActive(idx)
				c.onMatch(m.view)
				break
			}
		}
	}

	if !anyMatch {
		m.setActive(-1)
		if m.noMatch != nil {
			m.noMatch(m.view)
		}
	}
}

func (m *MediaMatcher) setActive(idx int) {
	for i, c := range m.criterias {
		c.active = i == idx
	}
}
