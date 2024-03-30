;;; tjderived-theme.el --- tjderived
;;; Version: 1.0
;;; Commentary:
;;; A theme called tjderived
;;; Code:

(deftheme tjderived "DOCSTRING for tjderived")
  (custom-theme-set-faces 'tjderived
   '(default ((t (:foreground "#bebebe" :background "#0f0f0f" ))))
   '(font-lock-comment-face ((t (:foreground "#999999"))))
   '(button ((t (:forground :inherit :underline t))))
   '(cursor ((t (:background "#6cb080" ))))
   '(fringe ((t (:background "#282828" ))))
   '(mode-line ((t (:foreground "#282828" :background "#7c6f64" ))))
   '(region ((t (:background "#504945" ))))
   '(secondary-selection ((t (:background "#3e3834" ))))
   '(font-lock-builtin-face ((t (:foreground "#9994b1" ))))
   '(font-lock-comment-face ((t (:foreground "#7c6f64" ))))
   '(font-lock-function-name-face ((t (:foreground "#f3f9ae" ))))
;;   '(font-lock-keyword-face ((t (:foreground "#7b6c80" ))))
   '(font-lock-string-face ((t (:foreground "#b5c7b6" ))))
   '(font-lock-keyword-face ((t (:foreground "#ac9db1" :weight light))))
   '(font-lock-type-face ((t (:foreground "#8d8a60" ))))
   '(font-lock-constant-face ((t (:foreground "#9892a8" ))))
   '(font-lock-variable-name-face ((t (:foreground "#83a598" ))))
   '(minibuffer-prompt ((t (:foreground "#b8bb26" :bold t ))))
   '(font-lock-warning-face ((t (:foreground "red" :bold t ))))
   '(lsp-flycheck-info-unnecessary-face ((t (:foreground "#666666" :bold nil ))))
   '(flycheck-info ((t (:foreground "#444444" :bold t ))))
   )

;;;###autoload
(and load-file-name
    (boundp 'custom-theme-load-path)
    (add-to-list 'custom-theme-load-path
                 (file-name-as-directory
                  (file-name-directory load-file-name))))
;; Automatically add this theme to the load path

(provide-theme 'tjderived)

;;; tjderived-theme.el ends here
