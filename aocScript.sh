SHELL=/bin/bash

giorno=$(date +'%d')
anno=$(date +'%Y')
cd $anno
mkdir day$giorno
cd day$giorno

touch prova.txt
echo "file = open('input.txt')" > a.py
echo "# file = open('$anno/day$giorno/prova.txt')" >> a.py
echo "file = open('input.txt')" > b.py
echo "# file = open('$anno/day$giorno/prova.txt')" >> b.py

giorno=$(date +'%e' | xargs)
echo $giorno $anno
curl https://adventofcode.com/$anno/day/$giorno/input -o input.txt -H "cookie: session=53616c7465645f5f99bea16bf125ba815c1e8250d697729410524d9eb77505227d986a540d7c385f173cb0565175a73e2d1c3e962273cefed378cbcf75d23355"