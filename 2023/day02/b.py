file = open("input.txt")

totale = 0

for line in file:
    split = line.split(":")
    
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
    
    totale += dic["red"] * dic["blue"] * dic["green"]

print(totale)