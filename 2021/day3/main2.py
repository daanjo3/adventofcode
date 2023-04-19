import numpy as np

previousMeasurement = None
countBigger = 0

windowCur = []
windowPrev = []

def getInput():
    with open('input.txt', 'r') as input:
        return input.readlines()

def getInputAsMatrix():
    lines = getInput()
    columns = []
    for line in lines:
        row = []
        for bit in line:
            if (bit != '\n'):
                row.append(bit)
        columns.append(row)
    return columns

lines = getInputAsMatrix()

def getCountBits(matrix, column):
    zeroes = 0
    ones = 0

    for row in matrix:
        if row[column] == '0':
            zeroes += 1
        else:
            ones += 1
    return zeroes, ones

def getMatchingRows(matrix, column, target):
    matches = []
    for row in matrix:
        if row[column] == target:
            matches.append(row)
    return matches

def getOxygenBinary(matrix):
    oxygenMatrix = matrix.copy()
    for i in range(0, len(oxygenMatrix[0])):
        zeroes, ones = getCountBits(oxygenMatrix, i)

        # Only bit that differs
        target = '1' if ones >= zeroes else '0'

        matches = getMatchingRows(oxygenMatrix, i, target)
        
        if len(matches) == 1:
            return ''.join(matches[0])

        oxygenMatrix = matches.copy()

def getCO2Binary(matrix):
    co2Matrix = matrix.copy()
    for i in range(0, len(co2Matrix[0])):
        zeroes, ones = getCountBits(co2Matrix, i)

        # Only bit that differs
        target = '1' if ones < zeroes else '0'

        matches = getMatchingRows(co2Matrix, i, target)
        
        if len(matches) == 1:
            return ''.join(matches[0])

        co2Matrix = matches.copy()

oxygenBin = getOxygenBinary(lines.copy())
co2Bin = getCO2Binary(lines.copy())

print(oxygenBin)
print(co2Bin)

print(int(oxygenBin, 2))
print(int(co2Bin, 2))

print(int(oxygenBin, 2) * int(co2Bin, 2))