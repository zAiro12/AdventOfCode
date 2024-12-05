file = open('input.txt')
# file = open('2024/day04/input.txt')
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

def cercaA(pos, matrice):
    if Posizione(pos.y-1, pos.x-1) in matrice and Posizione(pos.y+1, pos.x+1) in matrice and Posizione(pos.y-1, pos.x+1) in matrice and Posizione(pos.y+1, pos.x-1) in matrice:
        if matrice[Posizione(pos.y-1, pos.x-1)] != matrice[Posizione(pos.y+1, pos.x+1)] and matrice[Posizione(pos.y+1, pos.x-1)] != matrice[Posizione(pos.y-1, pos.x+1)]:
            if (matrice[Posizione(pos.y-1, pos.x-1)] == "M" or matrice[Posizione(pos.y-1, pos.x-1)] == "S") and (matrice[Posizione(pos.y+1, pos.x+1)] == "M" or matrice[Posizione(pos.y+1, pos.x+1)] == "S"):
                if (matrice[Posizione(pos.y-1, pos.x-1)] == matrice[Posizione(pos.y+1, pos.x-1)] and matrice[Posizione(pos.y-1, pos.x+1)] == matrice[Posizione(pos.y+1, pos.x+1)]) or (matrice[Posizione(pos.y+1, pos.x+1)] == matrice[Posizione(pos.y+1, pos.x-1)] and matrice[Posizione(pos.y-1, pos.x+1)] == matrice[Posizione(pos.y-1, pos.x-1)]):
                    return True

    

totale = 0
y = 0
matrice = {}

for line in file:
    split = list(line.strip())
    for x in range(len(split)):
        matrice[Posizione(y, x)] = split[x]
    y += 1
for pos in matrice:
    if matrice[pos] == "A":
        if cercaA(pos, matrice):
            totale += 1
print(totale)


