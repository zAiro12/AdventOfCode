file = open('input.txt')
# file = open('prova.txt')
# file = open('2025/day07/prova.txt')

total = 0
expected = 40

dic: dict[tuple[int, int], str] = {}
row = 0
col = 0
start = (0, 0)

for line in file:
    line = line.strip()
    for char in line:
        dic[(row, col)] = char
        
        if char == 'S':
            start = (row, col)
        col += 1

    row += 1
    col = 0

coda = [start]

while len(coda) > 0:
    pos = coda.pop(0)
    pos = (pos[0] + 1, pos[1])
    while dic.get(pos) is not None and dic[pos] != '^':
        pos = (pos[0] + 1, pos[1])
    if dic.get(pos) is not None and dic[pos] == '^':
        coda.insert(0, (pos[0], pos[1] + 1))
        coda.insert(0, (pos[0], pos[1] - 1))
        
    # print(len(coda))
    if pos[0] >= row:
        total += 1

print(total)
# print(total == expected, expected - total)