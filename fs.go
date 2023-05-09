package docx

import "embed"

var (
	//go:embed templates
	//go:embed templates/A4/_rels/.rels
	templateFS embed.FS

	templateFiles = []string{
		"_rels/.rels",
		"docProps/app.xml",
		"docProps/core.xml",
		"word/theme/theme1.xml",
		"word/fontTable.xml",
		"word/styles.xml",
		"[Content_Types].xml",
	}
)
