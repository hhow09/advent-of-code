
def fullConatined(a1: int, a2: int, b1: int, b2: int) -> bool:
    if (a1 <= b1 and b2 <= a2) or (a1 >= b1 and b2 >= a2):
        return True
    return False

def overlapped(a1: int, a2: int, b1: int, b2: int) -> bool:
    if a2 <= b2:
        return a2 >=b1
    return b2 >= a1

def parse(l:str) -> list[int]:
    arr = l.split(",")
    a = arr[0].split("-")
    a1, a2 = int(a[0]), int(a[1])
    b = arr[1].split("-")
    b1, b2 = int(b[0]), int(b[1])    
    return [a1,a2,b1,b2]

def part_one(file: str) -> int:
    count = 0
    with open(file) as f:
        lines = f.read().split('\n')
        for l in lines:
            intvs = parse(l)
            if fullConatined(*intvs):
                count += 1
    return count

def part_two(file: str) -> int:
    count = 0
    with open(file) as f:
        lines = f.read().split('\n')
        for l in lines:
            intvs = parse(l)
            if overlapped(*intvs):
                count += 1
    return count

if __name__ == "__main__":
    input_file = "./2022/day_04/input.txt"
    print("---part one---")
    print(part_one(input_file))
    print("---part two---")
    print(part_two(input_file))