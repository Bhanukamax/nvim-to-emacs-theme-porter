package main

import (
	"fmt"
	"os"
	"strings"
	"themer/lexer"
	"themer/nvim"
	"time"
)

func checkAndPanic(err error, errorMsg string) {
	if err != nil {
		panic(fmt.Sprintf("Bmaxo: Error", errorMsg, err))
	}
}

var shouldExport bool

func main() {
	//	shouldExport = true
	//	if shouldExport {
	n := nvim.New("/tmp/bmax-nvim.pipe")
	go func() {
		if err := n.StartServer(); err != nil {
			fmt.Println("error running neovim headless ", err)
		}
	}()
	fmt.Println("starting neovim")
	time.Sleep(time.Second / 2)
	fmt.Println("exporting theme to theme.vim")
	nvim.ExportTheme()
	fmt.Println("done")
	//	}

	colorMap := makeColorMap()
	colorNameMap := getColorNameMap()
	//	for key, color := range colorMap {
	//		fmt.Println("key: ", key, "color", color)
	//		fmt.Println("key", key, "color: ", color)
	//	}

	makeTheme("bmax-buddy-theme", colorNameMap, colorMap)

	fmt.Println("comment >>>", colorMap["Keyword"])
}

//	colorMap := makeColorMap()
//	fmt.Println(colorMap)

type Color struct {
	Fg string
	Bg string
}

func (c Color) String() string {
	if c.Bg != "" && c.Fg != "" {
		return fmt.Sprintf("{ bg: %s, fg: %s }", c.Bg, c.Fg)
	} else if c.Bg == "" {
		return fmt.Sprintf("{ fg: %s }", c.Fg)
	}
	return fmt.Sprintf("{ bg: %s }", c.Bg)
}

func newColor(fg string, bg string) *Color {
	return &Color{fg, bg}
	//	c := Color{fg, bg}
	//	return &ce
}

type ColorMap map[string]Color

func parseColor(input string) Color {
	color := Color{}

	fmt.Println("part: ", input)
	parts := strings.Fields(input)
	for _, part := range parts {

		pin := "guifg="

		if strings.HasPrefix(part[:], pin) {
			color.Fg = strings.TrimPrefix(part[:], pin)
		}

		pin = "guibg="
		if strings.HasPrefix(part[:], pin) {
			color.Bg = strings.TrimPrefix(part[:], pin)
		}
	}

	return color
}

func getColorNameMap() map[string]string {
	//   '(button ((t (:forground :inherit :underline t))))
	//   '(cursor ((t (:background "#6cb080" ))))
	//   '(secondary-selection ((t (:background "#3e3834" ))))
	//   '(lsp-flycheck-info-unnecessary-face ((t (:foreground "#666666" :bold nil ))))
	colorNameMap := map[string]string{
		"default":                      "Normal",
		"font-lock-comment-face":       "Comment",
		"line-number":                  "LineNr",
		"fringe":                       "LineNr",
		"mode-line":                    "StatusLine",
		"region":                       "Visual",
		"font-lock-builtin-face":       "Keyword",
		"font-lock-function-name-face": "Function",
		"font-lock-keyword-face":       "Keyword",
		"font-lock-string-face":        "String",
		"font-lock-type-face":          "Type",
		"font-lock-constant-face":      "Constant",
		"font-lock-variable-name-face": "variable",
		"minibuffer-prompt":            "commandmode",
		"font-lock-warning-face":       "ErrorMsg",
		"flycheck-info":                "DiagnosticInfo",
		"web-mode-variable-name-face":  "@property",
		"web-mode-html-tag-face":       "Function",
	}

	return colorNameMap
}

func mapHasKey(m map[string]Color, key string) bool {
	_, ok := m[key]

	return ok
}

func makeTheme(themeName string, names map[string]string, colorMap ColorMap) {
	theme := `(deftheme ` + themeName + ` "DOCSTRING for ` + themeName + `")
  (custom-theme-set-faces '` + themeName + `
`

	for key, color := range names {
		//		fmt.Println("key: ", key, "color", color)
		fmt.Println("key", key, "color: ", colorMap[color])
		c := colorMap[color]
		// 		theme += fmt.Sprintf(`   '(%s ((t (:foreground "%s" :background "%s" ))))
		// `, key, c.Fg, c.Bg)
		theme += fmt.Sprintf(`   '(%s ((t (`, key)

		if c.Fg != "" {
			theme += fmt.Sprintf(`:foreground "%s"`, c.Fg)
		}
		if c.Bg != "" {
			theme += fmt.Sprintf(`:background "%s"`, c.Bg)
		}
		theme += `))))
`
	}
	fmt.Println(theme)
}

func makeColorMap() ColorMap {
	colorMap := make(map[string]Color)
	data, err := os.ReadFile("./theme.vim")
	checkAndPanic(err, "reading file")
	//	fmt.Println(string(data))
	lines := strings.Split(string(data), "\n")
	for _, colorLine := range lines {
		//		fmt.Println(i, lines[i])

		parts := strings.Fields(colorLine)
		key := parts[1]
		if !mapHasKey(colorMap, key) {
			colorMap[key] = parseColor(colorLine)
		}
	}

	lexer.New("foo")
	return colorMap
}
