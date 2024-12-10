file = open('input.txt')
# file = open('prova.txt')
# file = open('2024/day08/prova.txt')

class Posizione:
    def __init__(self, y, x):
        self.y = y
        self.x = x

    def __str__(self):
        return f"(y={self.y}, x={self.x})"
    
    def __repr__(self):
        return f"(y={self.y}, x={self.x})"

    def __eq__(self, other):
        return self.y == other.y and self.x == other.x

    def __hash__(self):
        return hash((self.y, self.x))

def stampa_matrice(matrice):
    for chiave, posizioni in matrice.items():
        print(f"{chiave}: {posizioni}")

def stampa_matrice_antinodi(antinodi, max_y, max_x):
    matrice = [['.' for _ in range(max_x)] for _ in range(max_y)]
    for antinodo in antinodi:
        matrice[antinodo.y][antinodo.x] = '#'
    
    for row in matrice:
        print(''.join(row))

def trova_antinodi(matrice, max_y, max_x):
    antinodi = set()
    for _, posizioni in matrice.items():
        n = len(posizioni)
        for i in range(n):
            for j in range(i + 1, n):
                p1 = posizioni[i]
                p2 = posizioni[j]
                dy = p2.y - p1.y
                dx = p2.x - p1.x
                
                # Calcola i potenziali antinodi
                antinodo1 = Posizione(p1.y - dy, p1.x - dx)
                if 0 <= antinodo1.y < max_y and 0 <= antinodo1.x < max_x:
                    antinodi.add(antinodo1)

                antinodo2 = Posizione(p2.y + dy, p2.x + dx)
                if 0 <= antinodo2.y < max_y and 0 <= antinodo2.x < max_x:
                    antinodi.add(antinodo2)

    return len(antinodi)

matrice = {}
y = 0
for line in file:
    line = line.strip()
    for x, el in enumerate(line):
        if el != '.':
            if el not in matrice:
                matrice[el] = []
            matrice[el].append(Posizione(y, x))
    y += 1

max_y = y
max_x = max(len(line.strip()) for line in open('input.txt'))
print(trova_antinodi(matrice, max_y, max_x))