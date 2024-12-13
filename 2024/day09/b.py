with open('input.txt', 'r') as file:
# with open('2024/day09/prova.txt', 'r') as file:
    disk_map = file.readline().strip()

array = []
idNum = 0
isNum = True
for char in disk_map:
    if isNum:
        array.append(str(idNum) * int(char))
        isNum = False
        idNum += 1
    else:
        if char != '0':
            array.append(-1 * int(char))
        isNum = True

sx, dx = 0, len(array) - 1

while sx < dx:
    while dx >= 0 and type(array[dx]) != str:
        dx -= 1
    while sx < len(array):
        if type(array[sx]) is int and abs(array[sx]) >= len(array[dx]):
            break
        sx += 1
    # dx -= 1

    if sx < dx:
        array[sx] = array[sx] + len(array[dx])
        if array[sx] == 0:
            array.pop(sx)
            dx -= 1
        array.insert(sx, array[dx])
        array[dx + 1] = -1 * len(array[sx])
        
    sx, dx = 0, dx - 1
    sx = 0

result = ""
for item in array:
    if isinstance(item, str):
        result += item
    elif isinstance(item, int):
        result += '.' * abs(item)

checksum = 0
for i, char in enumerate(result):
    if char != '.':
        checksum += int(char) * (i)

print("Checksum:", checksum)