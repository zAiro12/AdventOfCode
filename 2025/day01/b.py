file = open('input.txt')
# file = open('prova.txt')
# file = open('2025/day01/prova.txt')

index = 50
risultato = 0
isPrevius0 = False

for line in file:
    line = line.strip()
    letter = line[0]
    number = int(line[1:])

    # print("Index prima:", index, " Istruzione:", line)
    if letter == 'R':
        index += number
    elif letter == 'L':
        index -= number

    if index >= 100 or index <= 0:
        if index <= 0:
            risultato += 1 * abs((abs(index) // 100) + 1)
            if isPrevius0 :
                risultato -= 1
        else:
            risultato += 1 * abs(index // 100)
        index = index % 100


        if index == 0:
            isPrevius0 = True

    if isPrevius0 and index != 0:
        isPrevius0 = False

    # print("Index dopo:", index, " Risultato parziale:", risultato)
    

print("Risultato b:", risultato)