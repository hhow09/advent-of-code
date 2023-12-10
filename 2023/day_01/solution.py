
def vals(filename: str) -> list[list[int]]:
    with open(filename, encoding="utf-8") as f:
        lines = f.read().split("\n")
        return [val(line) for line in lines]

def val(s: str):
    lres, rres = 0, 0
    for l in range(len(s)):
        if s[l].isnumeric():
            lres = int(s[l])
            break
    for r in range(len(s)-1, -1, -1):
        if s[r].isnumeric():
            rres = int(s[r])
            break
    return 10*lres + rres

def val2(s: str):
    lres, rres = 0, 0
    for l in range(len(s)):
        res = getNum(s[l:min(l+5, len(s))])
        if res > 0:
            lres = res
            break
    for r in range(len(s)-1, -1, -1):
        res = getNum(s[r:min(r+5, len(s))])
        if res > 0:
            rres = res
            break
    return 10*lres + rres

d = {
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
}

def getNum(s:str):
    if len(s)==0:
        return 0
    if s[0].isnumeric():
        return int(s[0])
    maxlen = max([len(key) for key in d.keys()])
    minlen = min([len(key) for key in d.keys()])

    for targetlen in range(minlen, maxlen+1):
        if len(s)>=targetlen and s[:targetlen] in d:
            return d[s[:targetlen]]
    return 0
    

def part_one(filename: str)-> int:
    with open(filename, encoding="utf-8") as f:
        lines = f.read().split("\n")
        return sum([val(line) for line in lines])
def part_two(filename: str)-> int:
    with open(filename, encoding="utf-8") as f:
        lines = f.read().split("\n")
        return sum([val2(line) for line in lines])

if __name__ == "__main__":
    input_path = "./2023/day_01/input.txt"
    print("---part one---")
    print(part_one(input_path))
    assert part_one(input_path) == 54667

    print("---part two---")
    print(part_two(input_path))
    assert part_two(input_path) == 54203