package settings

import (
	"github.com/gdamore/tcell/v2"
	"gitlab.com/tslocum/cview"
)

const (
	Submissions    = 0
	CommentSection = 1
)

func GetUnselectableItems() []int {
	return []int{0, 1, 3, 5, 7, 9, 10, 11, 13}
}

func SetSettingsList(list *cview.List, page int) {
	if page == 0 {
		SetToSubmissionsSettings(list)
	}
	if page == 1 {
		SetToCommentSectionSettings(list)
	}
}

func SetToSubmissionsSettings(list *cview.List) {
	list.Clear()
	indentation := "               "

	//0
	li := cview.NewListItem("")
	li.SetSecondaryText("[::b]Submissions")
	list.AddItem(li)

	//1
	li = cview.NewListItem("")
	list.AddItem(li)

	//2
	li = cview.NewListItem("Comment Width:")
	li.SetSecondaryText(indentation + "65     70     75     [[::b]80[::-]]     85     90     Screen ")
	list.AddItem(li)

	//3
	li = cview.NewListItem(" ")
	li.SetSecondaryText(" ")
	list.AddItem(li)

	//4
	li = cview.NewListItem("Indent size: ")
	li.SetSecondaryText(indentation + "1     2     3     [[::b]4[::-]]     5")
	list.AddItem(li)

	//5
	li = cview.NewListItem(" ")
	li.SetSecondaryText(" ")
	list.AddItem(li)

	//6
	li = cview.NewListItem("Show colors: ")
	li.SetSecondaryText(indentation + "no     [[black:#82aaff:]yes[-:-:-]]")
	list.AddItem(li)

	//7
	li = cview.NewListItem(" ")
	li.SetSecondaryText("")
	list.AddItem(li)

	//8
	li = cview.NewListItem("Show labels: ")
	li.SetSecondaryText(indentation + "[[black:orange:]no[-:-:-]]     yes")
	list.AddItem(li)

	//9
	//Space
	li = cview.NewListItem(" ")
	li.SetSecondaryText(" ")
	list.AddItem(li)

	//10
	li = cview.NewListItem(" ")
	li.SetSecondaryText("[::b]Comment Section")
	list.AddItem(li)

	//11
	//Line
	li = cview.NewListItem("")
	list.AddItem(li)

	//12
	li = cview.NewListItem("Show colors: ")
	li.SetSecondaryText(indentation + "[black:#82aaff:]yes")
	list.AddItem(li)

	//13
	li = cview.NewListItem(" ")
	li.SetSecondaryText("")
	list.AddItem(li)

	//14
	li = cview.NewListItem("Show labels: ")
	li.SetSecondaryText(indentation + "[black:orange:]no")
	list.AddItem(li)

	list.SetCurrentItem(2)
}

func SetToCommentSectionSettings(list *cview.List) {
	list.Clear()

	li := cview.NewListItem("")
	list.AddItem(li)

	li = cview.NewListItem("[::d]Press Enter to change")
	li.SetSecondaryText("Comment width: [::b]80")
	list.AddItem(li)

	li = cview.NewListItem(" ")
	li.SetSecondaryText(" ")
	list.AddItem(li)

	li = cview.NewListItem("[::d]Press Enter to change")
	li.SetSecondaryText("Indent size: [::b]4")
	list.AddItem(li)

	li = cview.NewListItem("")
	li.SetSecondaryText("Comment Section")

	//Line
	li = cview.NewListItem("")
	list.AddItem(li)

	list.AddItem(li)

	list.SetCurrentItem(1)
}

func NewDialogueBox() *cview.Modal {
	modal := cview.NewModal()
	modal.SetText("Do you want to quit the application? " +
		"Do you want to quit the application? Do you want to quit the application?")
	modal.AddButtons([]string{"Quit", "Cancel"})
	modal.SetBackgroundColor(tcell.ColorDefault)
	modal.SetTextColor(tcell.ColorDefault)

	return modal
}

func GetHeader(page int) string {
	if page == Submissions {
		return "[::b]Submissions"
	}
	if page == CommentSection {
		return "[::b]Comment Section"
	}

	return ""
}