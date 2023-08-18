package styles

import "gitee.com/jn-qq/simple-go-docx/shared"

func ExampleRunProperties_SetFont() {
	run := &RunProperties{}
	run.SetFont("楷体")
	//or run.SetFont("楷体", "楷体", "楷体", "楷体", "eastAsia")
}

func ExampleRunProperties_SetColor() {
	run := &RunProperties{}

	run.SetColor(shared.ColorLib.Fuchsia)

	run.SetColor(shared.RGB{
		R: 154,
		G: 255,
		B: 100,
	})

	run.SetColor("FFFFFF")
}
