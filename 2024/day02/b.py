file = open('input.txt')

totale = 0

def verificaDiff(nums, i):
    return abs(int(nums[i]) - int(nums[i+1])) <= 3 and abs(int(nums[i]) - int(nums[i+1])) >= 1

def verificaOrdine(nums):
    return all(int(nums[i]) <= int(nums[i+1]) for i in range(len(nums)-1)) or all(int(nums[i]) >= int(nums[i+1]) for i in range(len(nums)-1))

def verificaRimozione(nums):
    for i in range(len(nums)):
        temp = nums[:i] + nums[i+1:]
        if verificaOrdine(temp) and all(verificaDiff(temp, j) for j in range(len(temp)-1)):
            return True
    return False

for line in file:
    nums = line.strip().split(' ')
    if (verificaOrdine(nums) and all(verificaDiff(nums, i) for i in range(len(nums)-1))) or verificaRimozione(nums):
        totale += 1

print(totale)