#!/usr/bin/env bash
#PWD$(pwd)
nvim --server /tmp/bmax-nvim.pipe --remote-send ':source ~/bmax/nvim-to-emacs-theme/export-highlights.vim<CR>'
#nvim --server /tmp/bmax-nvim.pipe --remote-send ':source ~/bmax/nvim-to-emacs-theme/export-highlights.vim<CR>'
nvim --server /tmp/bmax-nvim.pipe --remote-send ':call ExportHighlights("/home/bmax/bmax/nvim-to-emacs-theme/theme.txt")<CR>'
