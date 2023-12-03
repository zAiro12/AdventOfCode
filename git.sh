SHELL=/bin/bash

giorno=$(date +'%d')
anno=$(date +'%Y')

cd /Users/zairo/Library/Mobile\ Documents/com~apple~CloudDocs/AdventOfCode/AdventOfCode
git add *
git commit -m"$anno $giorno"
git push