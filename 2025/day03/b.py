file = open('input.txt')
# file = open('prova.txt')
# file = open('2025/03/prova.txt')

totale = 0
expected = 3121910778619

def max_joltage(line_str: str, k: int = 12) -> int:
    """
    Trova il numero più grande di k cifre mantenendo l'ordine originale.
    Strategia greedy: per ogni posizione, scegli la cifra massima possibile
    lasciando abbastanza cifre per completare le k richieste.
    """
    n: int = len(line_str)
    if k > n:
        return int(line_str)
    
    result: list[str] = []
    start: int = 0
    
    for i in range(k):
        remaining: int = k - i
        end: int = n - remaining + 1
        
        max_digit: str = max(line_str[start:end])
        idx: int = line_str.index(max_digit, start, end)
        
        result.append(max_digit)
        start = idx + 1 
    
    return int(''.join(result))

for line in file:
    line_str = line.strip()
    joltage = max_joltage(line_str, k=12)
    # print(f"{line_str} -> {joltage}")
    totale += joltage

print(f"\nTotale: {totale}")
# print(f"Match: {totale == expected} (Difference: {expected - totale})")
