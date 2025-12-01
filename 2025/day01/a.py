file = open('input.txt')
# file = open('prova.txt')

index = 50
risultato = 0

for line in file:
    line = line.strip()
    letter = line[0]
    number = int(line[1:]) % 100

    if letter == 'R':
        index = (index + number) % 100
    elif letter == 'L':
        index = (index - number) % 100
    
    if index == 0:
        risultato += 1

print("Risultato a:", risultato)