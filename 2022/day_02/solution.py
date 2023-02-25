d1 = {"A": 1, "B": 2, "C": 3}
d2 = {"X": 1, "Y": 2, "Z": 3}

def win_score(a, b: str) -> int:
    if d1[a] == d2[b]:
        return 3
    if (a=="A" and b=="Y") or (a=="B" and b=="Z") or (a=="C" and b=="X"):
        return 6
    return 0

def rawScore(b: str) -> int:
    return d2[b]

def part_one(filename: str) -> int:
    scores = 0
    with open(filename, encoding="utf-8") as f:
        lines = list(map(lambda x: x.split(" "), f.read().split("\n")))
        for line in lines:
            scores += (rawScore(line[1]) + win_score(line[0], line[1]))
    return scores

def win_score_2(x: str) -> int:
    if x == "X":
        return 0
    if x == "Y":
        return 3
    return 6

loss_dict = {"A": 3, "B": 1, "C": 2}
win_dict = {"A": 2, "B": 3, "C": 1}
def rawScore_2(a, x: str) -> int:
    if x == "X":
        return loss_dict[a]
    if x == "Y":
        return d1[a]
    return win_dict[a]
        

def part_two(filename: str) -> int:
    scores = 0
    with open(filename, encoding="utf-8") as f:
        lines = list(map(lambda x: x.split(" "), f.read().split("\n")))
        for line in lines:
             scores += (rawScore_2(line[0], line[1]) + win_score_2(line[1]))
    return scores

if __name__ == "__main__":
    # input_file = "./2022/day_02/test.txt"
    input_file = "./2022/day_02/input.txt"
    print("---part one---")
    print(part_one(input_file))
    print("---part two---")
    print(part_two(input_file))