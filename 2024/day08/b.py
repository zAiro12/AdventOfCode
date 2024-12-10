file = open('input.txt')

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

def stampa_matrice_antinodi(antinodi, max_y, max_x):
    matrice = [['.' for _ in range(max_x)] for _ in range(max_y)]
    for antinodo in antinodi:
        matrice[antinodo.y][antinodo.x] = '#'
    
    for row in matrice:
        print(''.join(row))

def calcola_antinodi(matrice, max_y, max_x):
    antinodi = set()

    # Trova antinodi per ogni coppia di antenne con la stessa frequenza
    for _, posizioni in matrice.items():
        n = len(posizioni)

        # Se c'è una sola antenna di questa frequenza, è un antinodo
        if n == 1:
            antinodi.update(posizioni)
            continue

        # Analizza ogni coppia di antenne
        for i in range(n):
            for j in range(i + 1, n):
                p1 = posizioni[i]
                p2 = posizioni[j]
                dy = p2.y - p1.y
                dx = p2.x - p1.x

                # Estendi la linea tra p1 e p2 in entrambe le direzioni
                y, x = p1.y, p1.x
                while 0 <= y < max_y and 0 <= x < max_x:
                    antinodi.add(Posizione(y, x))
                    y -= dy
                    x -= dx

                y, x = p2.y, p2.x
                while 0 <= y < max_y and 0 <= x < max_x:
                    antinodi.add(Posizione(y, x))
                    y += dy
                    x += dx

        # Aggiungi tutte le posizioni delle antenne come antinodi
        antinodi.update(posizioni)

    # Restituisci il conteggio degli antinodi
    return len(antinodi)

# Leggi la mappa delle antenne
matrice = {}
y = 0
max_x = 0
for line in file:
    line = line.strip()
    max_x = max(max_x, len(line))
    for x, el in enumerate(line):
        if el != '.':
            if el not in matrice:
                matrice[el] = []
            matrice[el].append(Posizione(y, x))
    y += 1
max_y = y

# Calcola e stampa il risultato
print(calcola_antinodi(matrice, max_y, max_x))