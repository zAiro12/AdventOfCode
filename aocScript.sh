SHELL=/bin/bash

giorno=$(date +'%d')
anno=$(date +'%Y')
cd $anno
mkdir day$giorno
cd day$giorno

touch prova.txt

giorno=$(date +'%e' | xargs)
echo $giorno $anno
curl https://adventofcode.com/$anno/day/$giorno/input -o input.txt -H "cookie: session=53616c7465645f5f3bac7e5afb0c288b901110d6bddc8184363945e1d7bbee41eb4ef3dc2bfebdc119880970ed2854121838ee942abbb75abfcd96d184fbc6ad"

touch a.py
touch b.py