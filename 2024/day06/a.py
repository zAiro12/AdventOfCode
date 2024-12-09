file = open('input.txt')
# file = open('2024/day06/prova.txt')
# file = open('2024/day06/input.txt')

class Posizione:
    def __init__(self, y, x):
        self.y = y
        self.x = x

    def __str__(self):
        return f"(y={self.y}, x={self.x})"

    def __eq__(self, other):
        return self.y == other.y and self.x == other.x

    def __hash__(self):
        return hash((self.y, self.x))

totale = 1
y = 0
matrice = {}
start = Posizione(0, 0)
for line in file:
    line = line.strip()
    for x, car in enumerate(line):
        if car == '^':
            start = Posizione(y, x)
        matrice[Posizione(y, x)] = car
    y += 1

fine = False
direzione = 1
matrice[start] = 'X'
while not fine:
    try:
        if direzione == 1:
            if matrice[Posizione(start.y - 1, start.x)] != '#':
                if matrice[Posizione(start.y - 1, start.x)] != 'X':
                    matrice[Posizione(start.y - 1, start.x)] = 'X'
                    totale += 1
                start = Posizione(start.y - 1, start.x)
            else:
                direzione = 2
            continue
        
        if direzione == 2:
            if matrice[Posizione(start.y, start.x + 1)] != '#':
                if matrice[Posizione(start.y, start.x + 1)] != 'X':
                    matrice[Posizione(start.y, start.x + 1)] = 'X'
                    totale += 1
                start = Posizione(start.y, start.x + 1)
            else:
                direzione = 3
            continue
            
        if direzione == 3:
            if matrice[Posizione(start.y + 1, start.x)] != '#':
                if matrice[Posizione(start.y + 1, start.x)] != 'X':
                    matrice[Posizione(start.y + 1, start.x)] = 'X'
                    totale += 1
                start = Posizione(start.y + 1, start.x)
            else:
                direzione = 4
            continue
        
        if direzione == 4:
            if matrice[Posizione(start.y, start.x - 1)] != '#':
                if matrice[Posizione(start.y, start.x - 1)] != 'X':
                    matrice[Posizione(start.y, start.x - 1)] = 'X'
                    totale += 1
                start = Posizione(start.y, start.x - 1)
            else:
                direzione = 1
            continue
    except KeyError:
        fine = True

print(totale)
        
