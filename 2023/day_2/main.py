
FILE = 'input.txt'

sum = 0

max_cubes = {'red': 12, 'green': 13, 'blue': 14}

for line in open(FILE):
    parsed_line = line.split(":")
    game_id = int(parsed_line[0].split()[-1])

    valid = True

    for turn in parsed_line[1].split(";"):
        for move in turn.strip().split(","):
            m = move.strip().split(" ")
            if max_cubes[m[1]] < int(m[0]):
                valid = False

    if valid: sum += game_id

print(sum)