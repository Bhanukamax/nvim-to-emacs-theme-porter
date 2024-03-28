" Copied from:
" https://gist.github.com/Makaze/c27966b7108f77b280e5a0fce37d727c
function ExportHighlights(file)
  try
    let lines = execute('hi')
    let lines = substitute(lines, '\v(^|\n)+', '\1hi ', 'g')
    let lines = substitute(lines, 'xxx ', '', 'g')
    let lines = substitute(lines, '\v(\S+) +links to +(\S+)', 'link \1 \2', 'g')
    let lines = substitute(lines, '\v(\S+) +cleared', 'clear \1', 'g')
    let lines = split(lines, "\n")
    call writefile(lines, a:file, 'b')
    echo "Highlights exported to " . a:file
  catch
    echo "Error exporting highlights: " . v:exception
  endtry
endfunction