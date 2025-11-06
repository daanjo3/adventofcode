previousMeasurement = None
countBigger = 0

windowCur = []
windowPrev = []

def getInput():
    with open('input.txt', 'r') as input:
        return input.readlines()

lines = getInput()

depth = 0
pos = 0
aim = 0

for line in lines:
    command, num = line.split(" ")
    
    if command == "forward":
        pos += int(num)
        depth += aim * int(num)
    elif command == "down":
        aim += int(num)
    elif command == "up":
        aim -= int(num)

print(depth, pos)
print(depth * pos)
