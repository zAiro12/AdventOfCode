file = open('input.txt')
# file = open('prova.txt')
# file = open('2025/day04/prova.txt')

grid = [list(line.strip()) for line in file]
file.close()

rows = len(grid)
cols = len(grid[0]) if rows > 0 else 0

directions = [(-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, -1), (1, 0), (1, 1)]

total_removed = 0

while True:
    accessible: list[tuple[int, int]] = []
    
    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == '@':
                
                adjacent_rolls = 0
                for dr, dc in directions:
                    nr, nc = r + dr, c + dc
                    if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] == '@':
                        adjacent_rolls += 1
                
                if adjacent_rolls < 4:
                    accessible.append((r, c))
    
    if not accessible:
        break
    
    for r, c in accessible:
        grid[r][c] = '.'
    
    total_removed += len(accessible)

print(total_removed)