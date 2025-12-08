file = open('input.txt')
# file = open('prova.txt')
# file = open('2025/day06/prova.txt')

def print_matrix(matrix: dict[tuple[int, int], str]):
    max_row = max(pos[0] for pos in matrix.keys())
    max_col = max(pos[1] for pos in matrix.keys())
    
    # Calcola la larghezza massima necessaria per i numeri di riga
    row_width = len(str(max_row))
    
    # Calcola la larghezza massima necessaria per ogni cella
    cell_width = max(len(str(val)) for val in matrix.values())
    cell_width = max(cell_width, len(str(max_col)))  # Considera anche i numeri di colonna
    
    # Header con numeri di colonna
    print(" " * (row_width + 2), end="")
    for col in range(max_col + 1):
        print(f"{col:>{cell_width}}", end=" ")
    print()
    print(" " * (row_width + 2) + "-" * ((cell_width + 1) * (max_col + 1)))
    
    # Righe con numeri di riga
    for row in range(max_row + 1):
        print(f"{row:>{row_width}}|", end=" ")
        for col in range(max_col + 1):
            val = matrix.get((row, col), '.')
            print(f"{val:>{cell_width}}", end=" ")
        print()

matrix: dict[tuple[int, int], str] = {}
row = 0
col = 0
max_row = 0
max_col = 0

totale = 0
expected = 4277556

for line in file:
    line = line.strip().split(' ')
    for num in line:
        if num != '':
            matrix[(row, col)] = num
            max_col = max(max_col, col)
            col += 1
    max_row = max(max_row, row)
    row += 1
    col = 0


# print_matrix(matrix)

for col in range(max_col, -1, -1):
    isMolt = False
    delta = 0
    for row in range(max_row, -1, -1):
        if row == max_row and matrix.get((row, col)) == '*':
            isMolt = True
            delta = 1
            continue
        elif row == max_row and matrix.get((row, col)) == '+':
            isMolt = False
            delta = 0
            continue
        val = matrix.get((row, col))
        if val is None:
            continue
        if isMolt:
            delta *= int(val)
        else:
            delta += int(val)
    totale += delta

print(totale)
# print(totale == expected, expected - totale)