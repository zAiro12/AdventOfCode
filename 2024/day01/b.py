file = open('input.txt')
# file = open('2024/day01/prova.txt')

totale = 0
sx = {}
dx = {}

for line in file:
    line = line.strip().split('   ')
    sx[line[0]] = sx.get(line[0], 0) + 1
    dx[line[1]] = dx.get(line[1], 0) + 1

for key in sx:
    if key in dx:
        totale += sx[key] * (int(key) * dx[key])

print(totale)