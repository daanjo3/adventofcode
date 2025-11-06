previousMeasurement = None
countBigger = 0

windowCur = []
windowPrev = []

with open('input.txt', 'r') as input:
    lines = input.readlines()
    
    for i in range(0, len(lines)):
        if (i+2 > len(lines)-1):
            break

        windowCur = [int(lines[i]), int(lines[i+1]), int(lines[i+2])]

        if (len(windowPrev) != 0 and sum(windowCur) > sum(windowPrev)):
            countBigger += 1
        
        windowPrev = windowCur.copy()

print(countBigger)