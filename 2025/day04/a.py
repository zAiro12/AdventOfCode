file = open('input.txt')
# file = open('prova.txt')
# file = open('2025/day04/prova.txt')

grid = [line.strip() for line in file]

rows = len(grid)
cols = len(grid[0]) if rows > 0 else 0

directions = [(-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, -1), (1, 0), (1, 1)]

totale = 0
expected = 13

for r in range(rows):
    for c in range(cols):
        if grid[r][c] == '@':
            
            adjacent_rolls = 0
            for dr, dc in directions:
                nr, nc = r + dr, c + dc
                if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] == '@':
                    adjacent_rolls += 1
            
            if adjacent_rolls < 4:
                totale += 1

print(totale)
# print(totale == expected, expected - totale)