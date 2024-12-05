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
        if isSorted(nums, regole):
            totale += nums[len(nums) // 2]

print(totale)