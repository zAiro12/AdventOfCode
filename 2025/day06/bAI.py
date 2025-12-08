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

matrix: dict[tuple[int, int], str | None] = {}
row = 0
col = 0
max_row = 0
max_col = 0

totale = 0
expected = 3263827

for line in file:
    line = line.rstrip('\n')  # Rimuovi solo il newline, mantieni gli spazi
    col = 0
    for char in line:
        if char == ' ':
            matrix[(row, col)] = None
        else:
            matrix[(row, col)] = char
        max_col = max(max_col, col)
        col += 1
    max_row = max(max_row, row)
    row += 1


# print_matrix(matrix)

# Raccogli tutti i numeri e operatori colonna per colonna
all_numbers: list[int] = []
current_operator = ""

for col in range(max_col, -1, -1):
    current_number = ""
    
    for row in range(max_row + 1):
        val = matrix.get((row, col))
        
        if row == max_row:
            # Ultima riga: operatore
            if val and val in ['+', '*']:
                current_operator = val
        elif val is None:
            # Spazio: salva il numero corrente se esiste
            if current_number:
                all_numbers.append(int(current_number))
                current_number = ""
        else:
            # Cifra: aggiungila al numero corrente
            current_number += val
    
    # Aggiungi l'ultimo numero se presente
    if current_number:
        all_numbers.append(int(current_number))
    
    # Se troviamo un operatore, calcola il risultato per questo problema
    if current_operator:
        
        if all_numbers:
            result = all_numbers[0]
            if current_operator == '*':
                for num in all_numbers[1:]:
                    result *= num
            elif current_operator == '+':
                for num in all_numbers[1:]:
                    result += num
            
            totale += result
        
        # Reset per il prossimo problema
        all_numbers = []
        current_operator = ""

print(totale)
# print(totale == expected, expected - totale)