package designatexteditor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextEditor1(t *testing.T) {
	editor := Constructor()

	editor.AddText("leetcode")
	r1 := editor.DeleteText(4)
	assert.Equal(t, 4, r1)

	editor.AddText("practice")
	r2 := editor.CursorRight(3)
	assert.Equal(t, "etpractice", r2)

	r3 := editor.CursorLeft(8)
	assert.Equal(t, "leet", r3)

	r4 := editor.DeleteText(10)
	assert.Equal(t, 4, r4)

	r5 := editor.CursorLeft(2)
	assert.Equal(t, "", r5)

	r6 := editor.CursorRight(6)
	assert.Equal(t, "practi", r6)
}

/*
["TextEditor","addText","deleteText","addText","cursorLeft","addText","deleteText","addText","cursorLeft","deleteText"]
[[],["arnvmumatgmyw"],[5],["zrlufuifuy"],[2],["unh"],[20],["kwwp"],[6],[9]]
[null,null,5,null,"mazrlufuif",null,19,null,"",0]
*/

func TestTextEditor2(t *testing.T) {
	editor := Constructor()

	editor.AddText("arnvmumatgmyw")

	r1 := editor.DeleteText(5)
	assert.Equal(t, 5, r1)

	editor.AddText("zrlufuifuy")

	r2 := editor.CursorLeft(2)
	assert.Equal(t, "mazrlufuif", r2)

	editor.AddText("unh")

	r3 := editor.DeleteText(20)
	assert.Equal(t, 19, r3)

	editor.AddText("kwwp")

	r4 := editor.CursorLeft(6)
	assert.Equal(t, "", r4)

	r5 := editor.DeleteText(10)
	assert.Equal(t, 0, r5)
}
