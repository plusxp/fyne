package widget_test

import (
	"testing"
	"time"

	"fyne.io/fyne"
	internalwidget "fyne.io/fyne/internal/widget"
	"fyne.io/fyne/test"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/stretchr/testify/assert"
)

func TestTree_OpenClose(t *testing.T) {
	t.Run("Exists", func(t *testing.T) {
		data := make(map[string][]string)
		internalwidget.AddTreePath(data, "foo", "foobar")
		tree := widget.NewTreeWithStrings(data)

		assert.False(t, tree.IsBranchOpen("foo"))

		tree.OpenBranch("foo")
		assert.True(t, tree.IsBranchOpen("foo"))

		tree.CloseBranch("foo")
		assert.False(t, tree.IsBranchOpen("foo"))

		tree.ToggleBranch("foo")
		assert.True(t, tree.IsBranchOpen("foo"))

		tree.ToggleBranch("foo")
		assert.False(t, tree.IsBranchOpen("foo"))
	})
	t.Run("Missing", func(t *testing.T) {
		data := make(map[string][]string)
		internalwidget.AddTreePath(data, "foo", "foobar")
		tree := widget.NewTreeWithStrings(data)

		assert.False(t, tree.IsBranchOpen("foo"))

		tree.OpenBranch("bar")
		assert.False(t, tree.IsBranchOpen("foo"))

		tree.CloseBranch("bar")
		assert.False(t, tree.IsBranchOpen("foo"))

		tree.ToggleBranch("bar")
		assert.False(t, tree.IsBranchOpen("foo"))

		tree.ToggleBranch("bar")
		assert.False(t, tree.IsBranchOpen("foo"))
	})
}

func TestTree_OpenCloseAll(t *testing.T) {
	data := make(map[string][]string)
	internalwidget.AddTreePath(data, "foo0", "foobar0")
	internalwidget.AddTreePath(data, "foo1", "foobar1")
	internalwidget.AddTreePath(data, "foo2", "foobar2")
	tree := widget.NewTreeWithStrings(data)

	tree.OpenAllBranches()
	assert.True(t, tree.IsBranchOpen("foo0"))
	assert.True(t, tree.IsBranchOpen("foo1"))
	assert.True(t, tree.IsBranchOpen("foo2"))

	tree.CloseAllBranches()
	assert.False(t, tree.IsBranchOpen("foo0"))
	assert.False(t, tree.IsBranchOpen("foo1"))
	assert.False(t, tree.IsBranchOpen("foo2"))
}

func TestTree_Layout(t *testing.T) {
	test.NewApp()
	test.ApplyTheme(t, theme.LightTheme())

	for name, tt := range map[string]struct {
		items    [][]string
		selected string
		opened   []string
	}{
		"single_leaf": {
			items: [][]string{
				[]string{
					"11111",
				},
			},
		},
		"single_leaf_selected": {
			items: [][]string{
				[]string{
					"11111",
				},
			},
			selected: "11111",
		},
		"single_branch": {
			items: [][]string{
				[]string{
					"A", "11111",
				},
			},
		},
		"single_branch_selected": {
			items: [][]string{
				[]string{
					"A", "11111",
				},
			},
			selected: "A",
		},
		"single_branch_opened": {
			items: [][]string{
				[]string{
					"A", "11111",
				},
			},
			opened: []string{"A"},
		},
		"single_branch_opened_selected": {
			items: [][]string{
				[]string{
					"A", "11111",
				},
			},
			opened:   []string{"A"},
			selected: "A",
		},
		"single_branch_opened_leaf_selected": {
			items: [][]string{
				[]string{
					"A", "11111",
				},
			},
			opened:   []string{"A"},
			selected: "11111",
		},
		"multiple": {
			items: [][]string{
				[]string{
					"A", "11111",
				},
				[]string{
					"B", "2222222222",
				},
				[]string{
					"44444444444444444444",
				},
			},
		},
		"multiple_selected": {
			items: [][]string{
				[]string{
					"A", "11111",
				},
				[]string{
					"B", "2222222222",
				},
				[]string{
					"44444444444444444444",
				},
			},
			selected: "44444444444444444444",
		},
		"multiple_leaf": {
			items: [][]string{
				[]string{
					"11111",
				},
				[]string{
					"2222222222",
				},
				[]string{
					"333333333333333",
				},
				[]string{
					"44444444444444444444",
				},
			},
		},
		"multiple_leaf_selected": {
			items: [][]string{
				[]string{
					"11111",
				},
				[]string{
					"2222222222",
				},
				[]string{
					"333333333333333",
				},
				[]string{
					"44444444444444444444",
				},
			},
			selected: "2222222222",
		},
		"multiple_branch": {
			items: [][]string{
				[]string{
					"A", "11111",
				},
				[]string{
					"B", "2222222222",
				},
				[]string{
					"B", "C", "333333333333333",
				},
			},
		},
		"multiple_branch_selected": {
			items: [][]string{
				[]string{
					"A", "11111",
				},
				[]string{
					"B", "2222222222",
				},
				[]string{
					"B", "C", "333333333333333",
				},
			},
			selected: "B",
		},
		"multiple_branch_opened": {
			items: [][]string{
				[]string{
					"A", "11111",
				},
				[]string{
					"B", "2222222222",
				},
				[]string{
					"B", "C", "333333333333333",
				},
			},
			opened: []string{"A", "B", "C"},
		},
		"multiple_branch_opened_selected": {
			items: [][]string{
				[]string{
					"A", "11111",
				},
				[]string{
					"B", "2222222222",
				},
				[]string{
					"B", "C", "333333333333333",
				},
			},
			opened:   []string{"A", "B", "C"},
			selected: "B",
		},
		"multiple_branch_opened_leaf_selected": {
			items: [][]string{
				[]string{
					"A", "11111",
				},
				[]string{
					"B", "2222222222",
				},
				[]string{
					"B", "C", "333333333333333",
				},
			},
			opened:   []string{"A", "B", "C"},
			selected: "2222222222",
		},
	} {
		t.Run(name, func(t *testing.T) {
			data := make(map[string][]string)
			for _, d := range tt.items {
				internalwidget.AddTreePath(data, d...)
			}
			tree := widget.NewTreeWithStrings(data)
			for _, o := range tt.opened {
				tree.OpenBranch(o)
			}
			tree.SetSelection(tt.selected)

			window := test.NewWindow(tree)
			defer window.Close()
			window.Resize(fyne.NewSize(200, 300))

			tree.Refresh() // Force layout

			test.AssertImageMatches(t, "tree/layout_"+name+".png", window.Canvas().Capture())
		})
	}
}

func TestTree_ChangeTheme(t *testing.T) {
	app := test.NewApp()
	defer test.NewApp()
	app.Settings().SetTheme(theme.LightTheme())

	data := make(map[string][]string)
	internalwidget.AddTreePath(data, "foo", "foobar")
	tree := widget.NewTreeWithStrings(data)
	tree.OpenBranch("foo")

	window := test.NewWindow(tree)
	defer window.Close()
	window.Resize(fyne.NewSize(220, 220))

	tree.Refresh() // Force layout

	test.AssertImageMatches(t, "tree/theme_initial.png", window.Canvas().Capture())

	test.WithTestTheme(t, func() {
		tree.Refresh()
		time.Sleep(100 * time.Millisecond)
		test.AssertImageMatches(t, "tree/theme_changed.png", window.Canvas().Capture())
	})
}

func TestTree_Move(t *testing.T) {
	app := test.NewApp()
	defer test.NewApp()
	app.Settings().SetTheme(theme.LightTheme())

	data := make(map[string][]string)
	internalwidget.AddTreePath(data, "foo", "foobar")
	tree := widget.NewTreeWithStrings(data)
	tree.OpenBranch("foo")

	window := test.NewWindow(tree)
	defer window.Close()
	window.Resize(fyne.NewSize(220, 220))

	tree.Refresh() // Force layout

	test.AssertImageMatches(t, "tree/move_initial.png", window.Canvas().Capture())

	tree.Move(fyne.NewPos(20, 20))
	test.AssertImageMatches(t, "tree/move_moved.png", window.Canvas().Capture())
}