#!/usr/bin/env bash
rm -rf theme.txt
PIPE=/tmp/bmax-nvim.pipe
# nvim --headless --listen $PIPE
nvim --server $PIPE --remote-send ':source ~/bmax/nvim-to-emacs-theme/export-highlights.vim<CR>'
nvim --server $PIPE --remote-send ':lua vim.cmd.colorscheme("gruvbuddy")<CR>'
nvim --server $PIPE --remote-send ':call ExportHighlights("/home/bmax/bmax/nvim-to-emacs-theme/theme.txt")<CR>'
nvim --server $PIPE --remote-send ':qa!<CR>'
