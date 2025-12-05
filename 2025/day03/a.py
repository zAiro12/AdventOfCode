file = open('input.txt')
# file = open('prova.txt')
# file = open('2025/day03/prova.txt')

totale = 0
expected = 357

for line in file:
    line = list(line.strip())
    mappa: dict[int, list[int]] = {}
    for k, v in enumerate(line):
        mappa.setdefault(int(v), []).append(k)
    index = 9
    while index >= 0:
        try:
            maax = -1
            prima_pos = min(mappa[index])
            
            if prima_pos == len(line) - 1:
                index -= 1
                continue
            for val in line[prima_pos + 1:]:
                if int(val) > maax:
                    maax = int(val)
            # print(index, maax)
            totale += int(f"{index}{maax}")
            break
        except KeyError:
            index -= 1
print(totale)
# print(totale == expected, expected-totale)
    