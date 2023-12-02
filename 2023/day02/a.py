maxRed = 12
maxGreen = 13
maxBlue = 14

totale = 0

file = open("input.txt")

for line in file:
    split = line.split(":")
    
    numId = int(split[0].split(" ")[1])
    
    
    scan = split[1][1:].split(" ")
    dic = {"red": 0, "blue": 0, "green": 0}
    val = 0
    chiave = ""
    isVal = True
    for el in scan:
        if isVal:
            val = int(el)
            isVal = False
        else:
            if el[-1] == "," or el[-1] == ";" or el[-1] == "\n":
                el = el[:-1]
            if dic[el] < val:
                dic[el] = val
            isVal = True
    
    if dic["blue"] <= maxBlue and dic["green"] <= maxGreen and dic["red"] <= maxRed:
        totale += numId

print(totale)