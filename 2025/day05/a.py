file = open('input.txt')
# file = open('prova.txt')
# file = open('2025/day05/prova.txt')

def is_fresh(ingredient_id: int, ranges: list[tuple[int, int]]):
    for start, stop in ranges:
        if start <= ingredient_id <= stop:
            return True
    return False

totale = 0
exepted = 3
checkIfFresh = False
ranges: list[tuple[int, int]] = []

for line in file:
    if line == '\n':
        checkIfFresh = True
        continue
    line = line.strip()

    if not checkIfFresh:
        start, stop = map(int, line.strip().split('-'))
        ranges.append((start, stop))
    else:
        if is_fresh(int(line), ranges):
            totale += 1

print(totale)
# print(totale == exepted, exepted - totale)