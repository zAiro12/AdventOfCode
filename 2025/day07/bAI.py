file = open('input.txt')
# file = open('prova.txt')
# file = open('2025/day07/prova.txt')

total = 0
expected = 40

dic: dict[tuple[int, int], str] = {}
row = 0
col = 0
start = (0, 0)
max_row = 0

for line in file:
    line = line.strip()
    if not line:
        continue
    for char in line:
        dic[(row, col)] = char
        
        if char == 'S':
            start = (row, col)
        col += 1

    row += 1
    col = 0

max_row = row

# Memoization: da una posizione, quanti percorsi portano al fondo
memo: dict[tuple[int, int], int] = {}

def count_paths(pos: tuple[int, int]) -> int:
    """Conta quanti percorsi dalla posizione pos portano al fondo"""
    
    if pos in memo:
        return memo[pos]
    
    row, col = pos
    
    # Verifica che la posizione sia valida
    if pos not in dic:
        return 0
    
    # Caso base: siamo arrivati all'ultima riga
    if row == max_row - 1:
        return 1
    
    # Continua verso il basso
    next_row = row + 1
    next_pos = (next_row, col)
    
    # Se la prossima posizione è fuori dalla griglia
    if next_pos not in dic:
        return 0
    
    next_char = dic[next_pos]
    
    result = 0
    
    # Se incontriamo un ^, biforcazione
    if next_char == '^':
        # Biforcazione: vai a sinistra e a destra dalla riga DOPO il ^
        after_split_row = next_row + 1
        
        left = (after_split_row, col - 1)
        right = (after_split_row, col + 1)
        
        # Somma i percorsi da entrambi i rami
        if left in dic:
            result += count_paths(left)
        if right in dic:
            result += count_paths(right)
    else:
        # Continua dritto
        result = count_paths(next_pos)
    
    memo[pos] = result
    return result

# Conta i percorsi dalla posizione di start
total = count_paths(start)

print(total)
# print(total == expected, expected - total)