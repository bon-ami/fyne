package widget_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
)

func TestCheck_Binding(t *testing.T) {
	c := widget.NewCheck("", nil)
	c.SetChecked(true)
	assert.Equal(t, true, c.Checked)

	val := binding.NewBool()
	c.Bind(val)
	waitForBinding()
	assert.Equal(t, false, c.Checked)

	err := val.Set(true)
	assert.Nil(t, err)
	waitForBinding()
	assert.Equal(t, true, c.Checked)

	c.SetChecked(false)
	v, err := val.Get()
	assert.Nil(t, err)
	assert.Equal(t, false, v)

	c.Unbind()
	waitForBinding()
	assert.Equal(t, false, c.Checked)
}

func TestCheck_Layout(t *testing.T) {
	test.NewTempApp(t)

	for name, tt := range map[string]struct {
		text     string
		checked  bool
		disabled bool
	}{
		"checked": {
			text:    "Test",
			checked: true,
		},
		"unchecked": {
			text: "Test",
		},
		"checked_disabled": {
			text:     "Test",
			checked:  true,
			disabled: true,
		},
		"unchecked_disabled": {
			text:     "Test",
			disabled: true,
		},
	} {
		t.Run(name, func(t *testing.T) {
			check := &widget.Check{
				Text:    tt.text,
				Checked: tt.checked,
			}
			if tt.disabled {
				check.Disable()
			}

			window := test.NewTempWindow(t, &fyne.Container{Layout: layout.NewCenterLayout(), Objects: []fyne.CanvasObject{check}})
			window.Resize(check.MinSize().Max(fyne.NewSize(150, 200)))

			test.AssertRendersToMarkup(t, "check/layout_"+name+".xml", window.Canvas())
		})
	}
}

func TestNewCheckWithData(t *testing.T) {
	val := binding.NewBool()
	err := val.Set(true)
	assert.Nil(t, err)

	c := widget.NewCheckWithData("", val)
	waitForBinding()
	assert.Equal(t, true, c.Checked)

	c.SetChecked(false)
	v, err := val.Get()
	assert.Nil(t, err)
	assert.Equal(t, false, v)
}

func TestCheck_SetText(t *testing.T) {
	check := &widget.Check{Text: "Test"}
	check.SetText("New")

	assert.Equal(t, "New", check.Text)
}

func TestCheck_Tapped(t *testing.T) {
	check := &widget.Check{Text: "test"}
	assert.False(t, check.Checked)

	test.Tap(check)
	assert.True(t, check.Checked)
	test.Tap(check)
	assert.False(t, check.Checked)

	// and test the resetting from partial as well
	check.Partial = true
	test.Tap(check)
	assert.True(t, check.Checked)
	assert.False(t, check.Partial)
	test.Tap(check)
	assert.False(t, check.Checked)
	assert.False(t, check.Partial)

}
