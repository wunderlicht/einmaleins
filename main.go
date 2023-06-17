package main

import (
	"fmt"
	rtfdoc "github.com/therox/rtf-doc"
	"math/rand"
	"os"
)

const (
	upperBound   = 10
	linesPerPage = 20
	color        = rtfdoc.ColorBlack
	font         = rtfdoc.FontTimesNewRoman
	fontSize     = 20
)

func main() {
	//setup the document
	doc := rtfdoc.NewDocument()
	doc.SetFormat(rtfdoc.FormatA4)
	doc.SetOrientation(rtfdoc.OrientationPortrait)

	//add tasks to the document
	for i := 0; i < linesPerPage; i++ {
		taskText := fmt.Sprintf("%d * %d =", rand.Intn(upperBound+1), rand.Intn(upperBound+1))
		taskParagraph := doc.AddParagraph()
		taskParagraph.SetAlign(rtfdoc.AlignLeft)
		taskParagraph.AddText(taskText, fontSize, font, color)
		taskParagraph.AddNewLine()
	}

	//write it out to a file
	err := os.WriteFile("einmaleins.rtf", doc.Export(), 0644)
	if err != nil {
		panic(err)
	}
}
