file = open('input.txt')
# file = open('prova.txt')
# file = open('2025/day05/prova.txt')

totale = 0
exepted = 14
ranges: list[tuple[int, int]] = []

def merge_ranges(ranges: list[tuple[int, int]]) -> list[tuple[int, int]]:
    ranges = sorted(ranges)
    merged = [ranges[0]]
    
    for start, stop in ranges[1:]:
        last_start, last_stop = merged[-1]
        
        if start <= last_stop + 1:
            merged[-1] = (last_start, max(last_stop, stop))
        else:
            merged.append((start, stop))
    
    return merged

for line in file:
    if line == '\n':
        break
    line = line.strip()

    start, stop = map(int, line.strip().split('-'))
    ranges.append((start, stop))

ranges = merge_ranges(ranges)
for start, stop in ranges:
    totale += (stop - start + 1) 

print(totale)
# print(totale == exepted, exepted - totale)