file = open('input.txt')
# file = open('prova.txt')
# file = open('2024/day05/prova.txt')
# file = open('2024/day05/input.txt')

def isSorted(nums, regole):
    indexs= {page: idx for idx, page in enumerate(nums)}
    
    for x, y in regole:
        if x in indexs and y in indexs:
            if indexs[x] > indexs[y]:
                return False
    return True

def aggiusta(nums, regole):
    indexs= {page: idx for idx, page in enumerate(nums)}
    
    for _ in range(len(nums)):
        for x, y in regole:
            if x in indexs and y in indexs:
                if indexs[x] > indexs[y]:
                    indexs[x], indexs[y] = indexs[y], indexs[x]
    
    sorted_nums = sorted(indexs, key=indexs.get)
    return sorted_nums

totale = 0
lettura = True
regole = []

for line in file:
    if line == '\n':
        lettura = False
        continue

    if lettura:
        num1, num2 = map(int, line.strip().split('|'))
        regole.append((num1, num2))
    else:
        nums = list(map(int, line.strip().split(",")))
        if not isSorted(nums, regole):
            numSort = aggiusta(nums, regole)
            totale += numSort[len(numSort) // 2]

print(totale)