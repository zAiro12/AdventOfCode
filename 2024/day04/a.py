file = open('input.txt')
# file = open('2024/day04/prova.txt')

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

def cercaM(pos, matrice):
    isxmass = []
    if Posizione(pos.y-1, pos.x-1) in matrice and matrice[Posizione(pos.y-1, pos.x-1)] == "M":
        isxmass.append(cercaAS(Posizione(pos.y-1, pos.x-1), matrice, 1))
    if Posizione(pos.y-1, pos.x) in matrice and matrice[Posizione(pos.y-1, pos.x)] == "M":
        isxmass.append(cercaAS(Posizione(pos.y-1, pos.x), matrice, 2))
    if Posizione(pos.y-1, pos.x+1) in matrice and matrice[Posizione(pos.y-1, pos.x+1)] == "M":
        isxmass.append(cercaAS(Posizione(pos.y-1, pos.x+1), matrice, 3))
    if Posizione(pos.y, pos.x-1) in matrice and matrice[Posizione(pos.y, pos.x-1)] == "M":
        isxmass.append(cercaAS(Posizione(pos.y, pos.x-1), matrice, 4))
    if Posizione(pos.y, pos.x+1) in matrice and matrice[Posizione(pos.y, pos.x+1)] == "M":
        isxmass.append(cercaAS(Posizione(pos.y, pos.x+1), matrice, 6))
    if Posizione(pos.y+1, pos.x-1) in matrice and matrice[Posizione(pos.y+1, pos.x-1)] == "M":
        isxmass.append(cercaAS(Posizione(pos.y+1, pos.x-1), matrice, 7))
    if Posizione(pos.y+1, pos.x) in matrice and matrice[Posizione(pos.y+1, pos.x)] == "M":
        isxmass.append(cercaAS(Posizione(pos.y+1, pos.x), matrice, 8))
    if Posizione(pos.y+1, pos.x+1) in matrice and matrice[Posizione(pos.y+1, pos.x+1)] == "M":
        isxmass.append(cercaAS(Posizione(pos.y+1, pos.x+1), matrice, 9))
    return isxmass.count(True)

def cercaAS(pos, matrice, direzione):
    if direzione == 1:
        if (Posizione(pos.y-1, pos.x-1) in matrice and matrice[Posizione(pos.y-1, pos.x-1)] == "A" and
            Posizione(pos.y-2, pos.x-2) in matrice and matrice[Posizione(pos.y-2, pos.x-2)] == "S"):
            return True

    if direzione == 2:
        if (Posizione(pos.y-1, pos.x) in matrice and matrice[Posizione(pos.y-1, pos.x)] == "A" and
            Posizione(pos.y-2, pos.x) in matrice and matrice[Posizione(pos.y-2, pos.x)] == "S"):
            return True

    if direzione == 3:
        if (Posizione(pos.y-1, pos.x+1) in matrice and matrice[Posizione(pos.y-1, pos.x+1)] == "A" and
            Posizione(pos.y-2, pos.x+2) in matrice and matrice[Posizione(pos.y-2, pos.x+2)] == "S"):
            return True

    if direzione == 4:
        if (Posizione(pos.y, pos.x-1) in matrice and matrice[Posizione(pos.y, pos.x-1)] == "A" and
            Posizione(pos.y, pos.x-2) in matrice and matrice[Posizione(pos.y, pos.x-2)] == "S"):
            return True

    if direzione == 6:
        if (Posizione(pos.y, pos.x+1) in matrice and matrice[Posizione(pos.y, pos.x+1)] == "A" and
            Posizione(pos.y, pos.x+2) in matrice and matrice[Posizione(pos.y, pos.x+2)] == "S"):
            return True

    if direzione == 7:
        if (Posizione(pos.y+1, pos.x-1) in matrice and matrice[Posizione(pos.y+1, pos.x-1)] == "A" and
            Posizione(pos.y+2, pos.x-2) in matrice and matrice[Posizione(pos.y+2, pos.x-2)] == "S"):
            return True

    if direzione == 8:
        if (Posizione(pos.y+1, pos.x) in matrice and matrice[Posizione(pos.y+1, pos.x)] == "A" and
            Posizione(pos.y+2, pos.x) in matrice and matrice[Posizione(pos.y+2, pos.x)] == "S"):
            return True
    
    if direzione == 9:
        if (Posizione(pos.y+1, pos.x+1) in matrice and matrice[Posizione(pos.y+1, pos.x+1)] == "A" and
            Posizione(pos.y+2, pos.x+2) in matrice and matrice[Posizione(pos.y+2, pos.x+2)] == "S"):
            return True

    return False

totale = 0
y = 0
matrice = {}

for line in file:
    split = list(line.strip())
    for x in range(len(split)):
        matrice[Posizione(y, x)] = split[x]
    y += 1
for pos in matrice:
    if matrice[pos] == "X":
        totale += cercaM(pos, matrice)
print(totale)


