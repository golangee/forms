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

// hstepper contains a horizontal stepper component.
package hstepper

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/theme/material/icon"
	"strconv"
)

// Step is the model for the internal step view.
type Step struct {
	ico        icon.Icon
	numberOnly bool
	title      string
}

// NewStep creates a numbered step entry.
func NewStep(caption string) Step {
	return Step{numberOnly: true, title: caption}
}

// NewIconStep creates a step with an icon instead of a number.
func NewIconStep(ico icon.Icon, caption string) Step {
	return Step{ico: ico, title: caption}
}

// Stepper is a horizontal stepper.
type Stepper struct {
	*HStack
	stepInactiveColor Color
	steps             []*stepView
}

// NewStepper creates a new view with the given steps.
func NewStepper(steps ...Step) *Stepper {
	t := &Stepper{}
	t.HStack = NewHStack()
	t.stepInactiveColor = RGB(0x9e, 0x9e, 0x9e)
	t.SetHorizontalAlign(Center)
	t.SetSteps(steps...)
	return t
}

// SetSteps removes all existing steps and sets new ones.
func (t *Stepper) SetSteps(steps ...Step) *Stepper {
	t.ClearViews()
	t.steps = nil
	for idx, step := range steps {
		myIdx := -1
		if step.numberOnly {
			myIdx = idx + 1
		}
		stepView := newStepView(t, step.ico, step.title, myIdx, idx == len(steps)-1)
		stepView.setActive(false)
		t.steps = append(t.steps, stepView)
		t.HStack.AddViews(stepView)
	}
	return t
}

// SetProgress updates the view state of the stepper steps. Passed steps
// are colorized using the primary color and the active step caption is bolder.
func (t *Stepper) SetProgress(idx int) *Stepper {
	if idx > len(t.steps) {
		idx = len(t.steps)
	}

	for _, step := range t.steps {
		step.setDone(false)
		step.setActive(false)
	}

	for i := 0; i < idx; i++ {
		t.steps[i].setDone(true)
		if i == idx-1 {
			t.steps[i].setActive(true)
		}
	}

	return t
}

// Style applies generic style attributes.
func (t *Stepper) Style(style ...Style) *Stepper {
	t.HStack.Style(style...)
	return t
}

// Self assigns the receiver to the given reference
func (t *Stepper) Self(ref **Stepper) *Stepper {
	*ref = t
	return t
}

type stepView struct {
	parent  *Stepper
	btn     *IconButton
	caption *Text
	sepView *Frame
	*HStack
}

func newStepView(parent *Stepper, ico icon.Icon, text string, num int, last bool) *stepView {
	t := &stepView{}
	t.parent = parent
	t.HStack = NewHStack(
		NewIconButton(ico).Style(
			BackgroundColor(Theme().Color()),
			ForegroundColor(Theme().ForegroundColor()),
		).Self(&t.btn),
		newStepTitle(text).Self(&t.caption),

	)

	if num >= 0 {
		t.btn.SetChar(rune(strconv.Itoa(num)[0]))
	}

	if !last {
		t.HStack.AddViews(newStepSeparator().Self(&t.sepView))
	}

	return t
}

func (t *stepView) setDone(a bool) *stepView {
	if a {
		t.btn.Style(BackgroundColor(Theme().Color()))
		if t.sepView != nil {
			t.sepView.Style(BackgroundColor(Theme().Color()))
		}
	} else {
		t.btn.Style(BackgroundColor(t.parent.stepInactiveColor))
		if t.sepView != nil {
			t.sepView.Style(BackgroundColor(t.parent.stepInactiveColor))
		}
	}
	return t
}

func (t *stepView) setActive(a bool) *stepView {
	if a {
		t.caption.Style(FontWeight(WeightBolder))
	} else {
		t.caption.Style(FontWeight(WeightNormal))
	}

	return t
}

func newStepTitle(s string) *Text {
	return NewText(s).Style(
		MarginTop(Auto()),
		MarginBottom(Auto()),
		PadLeft(DefaultPadding),
		PadRight(DefaultPadding),
	)
}

func newStepSeparator() *Frame {
	return NewFrame().Style(
		BackgroundColor(Theme().Color()),
		Height(Pixel(1)),
		Width(Pixel(80)),
		MarginTop(Auto()),
		MarginBottom(Auto()),
		MarginRight(DefaultPadding),
	)
}
