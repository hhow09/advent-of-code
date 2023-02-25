
def process_carolies(filename: str) -> list[list[int]]:
    with open(filename, encoding="utf-8") as f:
        elves = f.read().split("\n\n")
        return [list(map(int, elf.strip().split("\n"))) for elf in elves]

def part_one(filename: str)-> int:
    carolies = process_carolies(filename)
    return max(sum(elf_car) for elf_car in carolies)

def part_two(filename: str)-> int:
    carolies = process_carolies(filename)
    return sum(sorted((sum(elf_car) for elf_car in carolies), reverse=True)[:3])

if __name__ == "__main__":
    input_path = "./2022/day_01/input.txt"
    print("---part one---")
    print(part_one(input_path))

    print("---part two---")
    print(part_two(input_path))