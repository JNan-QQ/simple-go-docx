package docx

import (
	"fmt"
	"testing"
)

func TestFs(t *testing.T) {
	fileBytes, _ := templateFS.ReadFile("templates/A4/_rels/.rels")
	fmt.Println(string(fileBytes))
}
