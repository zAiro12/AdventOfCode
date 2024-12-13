with open('prova.txt', 'r') as file:
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
print(array)
while sx < dx:
    while sx < len(array) and type(array[sx]) == str:
        sx += 1
    while dx >= 0 and type(array[sx]) != str:
        dx -= 1
    print(sx, dx)
    if sx < dx and abs(array[sx]) >= len(array[dx]):
        array[sx] = array[sx] + len(array[dx])
        if array[sx] == 0:
            array.pop(sx)
        array.insert(sx, array[dx])
        array[dx] = 0

print(array)