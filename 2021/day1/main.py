previousMeasurement = None
countBigger = 0

with open('input.txt', 'r') as input:
    for line in input.readlines():
        measurement = int(line)
        # print(measurement)
        if (previousMeasurement != None and measurement > previousMeasurement):
            print(measurement, previousMeasurement)
            countBigger += 1
        previousMeasurement = measurement

print(countBigger)