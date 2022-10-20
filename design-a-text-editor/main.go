package designatexteditor

import "strings"

type TextEditor struct {
	blocks       *block
	currentBlock *block
	blockOffset  int
	cursorPos    int
	size         int
}

type block struct {
	text string
	prev *block
	next *block
}

func (b *block) delete() *block {
	b.next.prev = b.prev
	b.prev.next = b.next
	return b.prev
}

func Constructor() TextEditor {
	head := new(block)
	head.prev = head
	head.next = head
	return TextEditor{
		blocks:       head,
		currentBlock: head,
	}
}

func (this *TextEditor) AddText(text string) {
	if len(text) == 0 {
		return
	}
	currentBlock := this.currentBlock
	blockOffset := this.blockOffset
	if blockOffset < len(currentBlock.text) {
		remainText := currentBlock.text[blockOffset:]
		currentBlock.text = currentBlock.text[:blockOffset]
		splitBlock := &block{
			text: remainText,
			prev: currentBlock,
			next: currentBlock.next,
		}
		currentBlock.next.prev = splitBlock
		currentBlock.next = splitBlock
		blockOffset = len(currentBlock.text)
	}
	if blockOffset == 0 {
		currentBlock.text = text
	} else {
		newBlock := &block{
			text: text,
			prev: currentBlock,
			next: currentBlock.next,
		}
		currentBlock.next.prev = newBlock
		currentBlock.next = newBlock
		currentBlock = newBlock
		this.currentBlock = currentBlock
	}
	this.blockOffset = len(currentBlock.text)
	this.cursorPos += len(text)
	this.size += len(text)
}

func (this *TextEditor) DeleteText(k int) int {
	if k == 0 {
		return 0
	}
	k = minOf(this.cursorPos, k)
	currentBlock := this.currentBlock
	blockOffset := this.blockOffset
	if blockOffset == 0 {
		currentBlock = currentBlock.prev
		blockOffset = len(currentBlock.text)
	}
	if blockOffset < len(currentBlock.text) {
		remainText := currentBlock.text[blockOffset:]
		currentBlock.text = currentBlock.text[:blockOffset]
		splitBlock := &block{
			text: remainText,
			prev: currentBlock,
			next: currentBlock.next,
		}
		currentBlock.next.prev = splitBlock
		currentBlock.next = splitBlock
	}
	n := k
	for n > 0 {
		if n <= len(currentBlock.text) {
			currentBlock.text = currentBlock.text[:len(currentBlock.text)-n]
			if len(currentBlock.text) == 0 && currentBlock != this.blocks {
				currentBlock = currentBlock.delete()
			}
			break
		}
		n -= len(currentBlock.text)
		currentBlock = currentBlock.delete()
	}
	this.currentBlock = currentBlock
	this.blockOffset = len(currentBlock.text)
	this.cursorPos -= k
	this.size -= k
	return k
}

func (this *TextEditor) CursorLeft(k int) string {
	k = minOf(this.cursorPos, k)
	this.cursorPos -= k
	currentBlock := this.currentBlock
	blockOffset := this.blockOffset
	for k > 0 {
		if k <= blockOffset {
			this.blockOffset = blockOffset - k
			break
		}
		k -= blockOffset
		currentBlock = currentBlock.prev
		blockOffset = len(currentBlock.text)
	}
	this.currentBlock = currentBlock
	return this.lastN(10)
}

func (this *TextEditor) CursorRight(k int) string {
	k = minOf(this.size-this.cursorPos, k)
	this.cursorPos += k
	currentBlock := this.currentBlock
	blockOffset := this.blockOffset
	for k > 0 {
		if blockOffset+k <= len(currentBlock.text) {
			this.blockOffset = blockOffset + k
			break
		}
		k -= len(currentBlock.text) - blockOffset
		currentBlock = currentBlock.prev
		blockOffset = 0
	}
	this.currentBlock = currentBlock
	return this.lastN(10)
}

func (this *TextEditor) lastN(n int) string {
	n = minOf(this.cursorPos, n)
	var segments []string
	currentBlock := this.currentBlock
	blockOffset := this.blockOffset
	for n > 0 {
		if n <= blockOffset {
			segments = append(segments, currentBlock.text[blockOffset-n:blockOffset])
			break
		}
		n -= blockOffset
		segments = append(segments, currentBlock.text[:blockOffset])
		currentBlock = currentBlock.prev
		blockOffset = len(currentBlock.text)
	}
	for i, j := 0, len(segments)-1; i < j; i, j = i+1, j-1 {
		segments[i], segments[j] = segments[j], segments[i]
	}
	return strings.Join(segments, "")
}

func minOf(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
