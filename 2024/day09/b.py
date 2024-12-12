with open('prova.txt', 'r') as file:
    disk_map = file.readline().strip()

array = []
liberi = {}
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
            liberi[len(array)-1] = int(char)
        isNum = True

sx, dx = 0, len(array) - 1
while sx < dx:
    while sx < len(array) and type(array[sx]) == str:
        sx += 1
    while dx >= 0 and type(array[sx]) != str:
        dx -= 1
    print(sx, dx)
    # if sx < dx:
    #     array[sx], array[dx] = array[dx], array[sx]