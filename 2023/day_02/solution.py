def check_possible_game_id(line: str, critiria: dict) -> int:
    splits = line.split(":")
    gameid = int(splits[0][len("Game "):])
    sets = splits[1].split(';')
    d = {'red':0,'green':0, 'blue':0}
    for s in sets:
        for num_color in s.split(','):
            if not num_color.strip():
                continue
            sp = num_color.strip().split(' ')
            num, color = int(sp[0]), sp[1]
            if num > critiria[color]:
                return 0
    return gameid


def get_min_possible(line: str) ->int:
    splits = line.split(":")
    sets = splits[1].split(';')
    res = {'red':1,'green':1, 'blue':1}
    for s in sets:
        for num_color in s.split(','):
            if not num_color.strip():
                continue
            sp = num_color.strip().split(' ')
            num, color = int(sp[0]), sp[1]
            res[color] = max(res[color], num)
    return res['green'] * res['red'] * res['blue']  


def part_one(filename: str)-> int:
    critiria = {'red':12, 'green':13, 'blue':14}
    with open(filename, encoding="utf-8") as f:
        lines = f.read().split("\n")
        return sum([check_possible_game_id(line, critiria) for line in lines])
        
def part_two(filename: str)-> int:
    with open(filename, encoding="utf-8") as f:
        lines = f.read().split("\n")
        return sum([get_min_possible(line) for line in lines])

if __name__ == "__main__":
    input_path = "./2023/day_02/input.txt"
    print("---part one---")
    print(part_one(input_path))
    assert part_one(input_path) == 2207

    print("---part two---")
    print(part_two(input_path))
    assert part_two(input_path) == 62241