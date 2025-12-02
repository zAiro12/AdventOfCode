SHELL=/bin/bash

giorno=$(date +'%d')
anno=$(date +'%Y')
cd $anno
mkdir day$giorno
cd day$giorno

touch prova.txt
echo "# file = open('input.txt')" > a.py
echo "file = open('prova.txt')" >> a.py
echo "# file = open('$anno/$giorno/prova.txt')" >> a.py
echo "# file = open('input.txt')" > b.py
echo "file = open('prova.txt')" >> b.py
echo "# file = open('$anno/$giorno/prova.txt')" >> b.py

giorno=$(date +'%e' | xargs)
echo $giorno $anno
curl https://adventofcode.com/$anno/day/$giorno/input -o input.txt -H "cookie: session=53616c7465645f5f05b484988b0356fdd8bc6cbf9187ffe0c2ee9449c6f4d9c30c4dd1fd20cd8d051cdd0a026c6d7f23808c3a2d160e2033d86d902bc69599e3"