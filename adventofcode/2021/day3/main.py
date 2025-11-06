previousMeasurement = None
countBigger = 0

windowCur = []
windowPrev = []

def getInput():
    with open('input.txt', 'r') as input:
        return input.readlines()

lines = getInput()

counts = [{"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}]

for line in lines:
    for i in range(0, len(line)):
        bit = line[i]
        if (bit == "1" or bit == "0"):
            counts[i][bit] += 1

gammaBin = ""
epsilonBin = ""

for bitcount in counts:
    if (bitcount["0"] > bitcount["1"]):
        gammaBin += "0"
        epsilonBin += "1"
    else:
        gammaBin += "1"
        epsilonBin += "0"

print(gammaBin)
print(epsilonBin)

gamma = int(gammaBin, 2)
epsilon = int(epsilonBin, 2)

print(gamma)
print(epsilon)

print(gamma * epsilon)
