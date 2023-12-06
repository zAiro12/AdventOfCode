file = open("input.txt")

totalPoint = 0
for line in file:
    split = line.split(":")[1]
    nums = split.split(" | ")
    # tolgo /n
    nums[1] = nums[1][:-1]
    ownNum = nums[1].split(" ")
    winNum = nums[0].split(" ")

    dic = {}

    for n in ownNum:
        if not n == "":
            dic[n] = 0
    
    for n in winNum:
        try:    
            dic[n] += 1
        except KeyError:
            continue
    
    tmptot = 0
    for k in dic:
        if dic[k] > 0:
            tmptot+= dic[k]
            
    if not tmptot == 0:
        totalPoint += 2**(tmptot-1)

    
print(totalPoint)
