file = open('input.txt')
# file = open('prova.txt')
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

class Cancelletto:
    def __init__(self, pos, direzione):
        self.pos = pos
        self.direzione = direzione

    def __str__(self):
        return f"({self.pos}, {self.direzione})"

    def __eq__(self, other):
        return self.pos == other.pos and self.direzione == other.direzione

    def __hash__(self):
        return hash((self.pos, self.direzione))


def stampaMatrice(matrice, larghezza, altezza):
    for y in range(altezza):
        for x in range(larghezza):
            pos = Posizione(y, x)
            if pos in matrice:
                print(matrice[pos], end='')
            else:
                print(' ', end='')
        print()

def cercaLoop(matrice, start, newBlocco):
    cursore = start
    direzione = 1
    visti = {}
    while True:
        try:
            if direzione == 1:
                pos = Posizione(cursore.y - 1, cursore.x)
                if matrice[pos] == '#' or pos == newBlocco:
                    ora = Cancelletto(pos, 1)
                    if ora in visti:
                        return 1
                    visti[ora] = True
                    direzione = 2
                else:
                    cursore = pos
                continue
            
            if direzione == 2:
                pos = Posizione(cursore.y, cursore.x + 1)
                if matrice[pos] == '#' or pos == newBlocco:
                    ora = Cancelletto(pos, 2)
                    if ora in visti:
                        return 1
                    visti[ora] = True
                    direzione = 3
                else:
                    cursore = pos
                continue
                
            if direzione == 3:
                pos = Posizione(cursore.y + 1, cursore.x)
                if matrice[pos] == '#' or pos == newBlocco:
                    ora = Cancelletto(pos, 3)
                    if ora in visti:
                        return 1
                    visti[ora] = True
                    direzione = 4
                else:
                    cursore = pos
                continue
            
            if direzione == 4:
                pos = Posizione(cursore.y, cursore.x - 1)
                if matrice[pos] == '#' or pos == newBlocco:
                    ora = Cancelletto(pos, 4)
                    if ora in visti:
                        return 1
                    visti[ora] = True
                    direzione = 1
                else:
                    cursore = pos
                continue
        except KeyError:
            return 0

totale = 0
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

for pos in matrice:
    print(pos)
    totale += cercaLoop(matrice, start, pos)

print(totale)
        
