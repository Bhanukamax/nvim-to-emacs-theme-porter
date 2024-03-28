#!/usr/bin/env bash
rm -rf theme.txt
PIPE=/tmp/bmax-nvim.pipe
nvim --server $PIPE --remote-send ':source ./export-highlights.vim<CR>'
nvim --server $PIPE --remote-send ':lua vim.cmd.colorscheme("gruvbuddy")<CR>'
nvim --server $PIPE --remote-send ':call ExportHighlights("./theme.txt")<CR>'
nvim --server $PIPE --remote-send ':qa!<CR>'
