
def getIdx(char: str) -> int:
    if char.islower():
        return ord(char) - ord('a')
    else:
        return ord(char) - ord('A')+26


# return the sum of priority given list of string
def priority(arr: list[str]) -> int:
    l = [0]*52
    for s in arr:
        for char in set(s): # prevent duplicate
            l[getIdx(char)] += 1
    return sum(i+1 if v == len(arr) else 0 for i, v in enumerate(l))


def part_one(fileName: str) -> int:
    with open(fileName, encoding="utf-8", mode='r') as f:
        lines = f.read().split('\n')
        p = 0
        for line in lines:
            n = len(line)
            first, second = line[:n//2], line[n//2:]
            p += priority([first, second])
        return p

def part_two(fileName: str) -> int:
    with open(fileName, encoding="utf-8", mode='r') as f:
        lines = f.read().split('\n')
        p = 0
        for i in range(len(lines)):
            if (i+1)%3==0:
                p+=priority(lines[i-2:i+1])
        return p


if __name__ == "__main__":
    # input_file = "./2022/day_03/test.txt"
    input_file = "./2022/day_03/input.txt"
    print("---part one---")
    print(part_one(input_file)) # 7903
    print("---part two---")
    print(part_two(input_file)) # 2548
