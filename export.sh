#!/usr/bin/env bash
#touch -rf theme.vim
#rm -rf theme.vim
PIPE=/tmp/bmax-nvim.pipe
nvim --server $PIPE --remote-send ':source ./export-highlights.vim<CR>'
nvim --server $PIPE --remote-send ':lua vim.cmd.colorscheme("gruvbuddy")<CR>'
nvim --server $PIPE --remote-send ':call ExportHighlights("./theme.vim")<CR>'
nvim --server $PIPE --remote-send ':qa!<CR>'
