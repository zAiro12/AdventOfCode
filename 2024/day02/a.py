file = open('input.txt')
# file = open('2024/day02/prova.txt')

totale = 0
for line in file:
    nums = line.strip().split(' ')
    parziale = 0
    for i in range(len(nums)-1):
        if abs(int(nums[i]) - int(nums[i+1])) <= 3 and abs(int(nums[i]) - int(nums[i+1])) > 0:
            if int(nums[i]) > int(nums[i+1]):
                parziale -= 1
            else:
                parziale += 1
    if abs(parziale) == len(nums)-1:
        totale += 1
print(totale)