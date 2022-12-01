SHELL=/bin/bash

giorno=$(date +'%d')
anno=$(date +'%Y')
echo "$giorno $anno"
cd $anno
mkdir day$giorno
cd day$giorno

touch prova.txt

curl https://adventofcode.com/$anno/day/$giorno/input -o input.txt -H "cookie: session=53616c7465645f5f42ead91e71c09e0e007085c54e5123b2f618ee4da95986febdcf8b23a8edaaefd112dd540a55954bc400a14c8f652acdb6a2c834717fed7d"

echo 'package main ' > $giorno.go
echo '' >> $giorno.go
echo 'import (' >> $giorno.go
echo '  zairo "github.com/zAiro12/AdventOfCode/utile"'>> $giorno.go
echo ')' >> $giorno.go
echo '' >> $giorno.go
echo 'func main(){' >> $giorno.go
echo '' >> $giorno.go
echo '}' >> $giorno.go