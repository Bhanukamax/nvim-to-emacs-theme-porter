package main

import (
	"fmt"
	"os"
	"strings"
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
	n := nvim.New("/tmp/bmax-nvim.pipe")
	go func() {
		if err := n.StartServer(); err != nil {
			fmt.Println("error running neovim headless ", err)
		}
	}()
	fmt.Println(";; starting neovim")
	time.Sleep(time.Second / 2)
	fmt.Println(";; exporting theme to theme.vim")
	nvim.ExportTheme()
	fmt.Println(";; done exporting from nvim")

	faceMap := makeFaceMap()
	colorNameMap := getColorNameMap()

	makeTheme("bmax-buddy-theme", colorNameMap, faceMap)
}

type Face struct {
	Fg     string
	Bg     string
	Weight string
	Italic bool
}

func (c Face) String() string {
	if c.Bg != "" && c.Fg != "" {
		return fmt.Sprintf("{ bg: %s, fg: %s }", c.Bg, c.Fg)
	} else if c.Bg == "" {
		return fmt.Sprintf("{ fg: %s }", c.Fg)
	}
	return fmt.Sprintf("{ bg: %s }", c.Bg)
}

func newColor(fg string, bg string) *Face {
	return &Face{fg, bg, "", false}
}

type FaceMap map[string]Face

func parseColor(input string) Face {
	color := Face{}

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

		pin = "gui=bold"
		if strings.HasPrefix(part[:], pin) {
			color.Weight = "bold"
		}

		pin = "gui=italic"
		if strings.HasPrefix(part[:], pin) {
			color.Italic = true
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
		"web-mode-type-face":           "@type",
	}

	return colorNameMap
}

func mapHasKey(m map[string]Face, key string) bool {
	_, ok := m[key]

	return ok
}

func makeTheme(themeName string, names map[string]string, faceMap FaceMap) {
	theme := `(deftheme ` + themeName + ` "DOCSTRING for ` + themeName + `")
  (custom-theme-set-faces '` + themeName + `
`

	for key, color := range names {
		c := faceMap[color]
		theme += fmt.Sprintf(`   '(%s ((t (`, key)

		if c.Fg != "" {
			theme += fmt.Sprintf(`:foreground "%s"`, c.Fg)
		}
		if c.Bg != "" {
			theme += fmt.Sprintf(` :background "%s"`, c.Bg)
		}
		if c.Weight != "" {
			theme += fmt.Sprintf(` :weight %s`, c.Weight)
		}
		if c.Italic {
			theme += fmt.Sprintf(" :italic t")
		}

		theme += `))))
`
	}
	theme += `)
;;;###autoload
(and load-file-name
     (boundp 'custom-theme-load-path)
     (add-to-list 'custom-theme-load-path
                  (file-name-as-directory
                   (file-name-directory load-file-name))))
;; Automatically add this theme to the load path

(provide-theme '` + themeName + `)

;;; ` + themeName + `-theme.el ends here
;;; save this to your theme load path as "` + themeName + `-theme.el"

`

	fmt.Println(theme)
}

func makeFaceMap() FaceMap {
	faceMap := make(map[string]Face)
	data, err := os.ReadFile("./theme.vim")
	checkAndPanic(err, "reading file")
	lines := strings.Split(string(data), "\n")
	for _, colorLine := range lines {
		parts := strings.Fields(colorLine)
		key := parts[1]
		if !mapHasKey(faceMap, key) {
			faceMap[key] = parseColor(colorLine)
		}
	}
	return faceMap
}
