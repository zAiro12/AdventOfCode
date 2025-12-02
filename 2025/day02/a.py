file = open('input.txt')
# file = open('prova.txt')
# file = open('2025/day02/prova.txt')

def split_exact_half(s: str) -> tuple[str, str]:
    if len(s) % 2 != 0:
        raise ValueError("La stringa non ha lunghezza pari")
    mid = len(s) // 2
    return s[:mid], s[mid:]

lines = file.readlines()[0].strip().split(',')
totale = 0
expected = 1227775554

for line in lines:
    start, stop = map(int, line.strip().split('-'))
    
    if len(str(start)) % 2 != 0 and len(str(stop)) % 2 != 0:
        continue

    for index in range(start, stop + 1):
        s = str(index)

        if len(s) % 2 != 0:
            continue
        sx, dx = split_exact_half(str(index))
        if sx == dx:
            totale += index
    

print("Totale:", totale)