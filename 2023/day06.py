times, distances = open("input.txt").readlines()

times = [int(x.strip()) for x in times.split(":")[1].split()]
distances = [int(x.strip()) for x in distances.split(":")[1].split()]

acc = 1
for t, d in zip(times, distances):
    possibilities = len([z for z in range(0, t + 1) if z * (t - z) > d])
    acc *= possibilities

print(acc)
