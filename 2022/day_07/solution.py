from pathlib import Path
from collections import defaultdict
from bisect import bisect_left

def get_dir_sizes_dict(filename: str):
	curr = Path("/")
	dir_sizes = defaultdict(int) # path -> file size
	with open(filename) as f:
		lines = f.read().strip().splitlines()
		for line in lines:
			if line.startswith("$ cd"):
				curr = (curr / line[5:]).resolve()
			elif line.startswith("$ ls"):
				continue
			elif line.startswith("dir"):
				continue
			else:
				size = int(line.split(" ")[0])
				dir_sizes[curr] += size
				# parents: An immutable sequence providing access to the logical ancestors of the path
				for parent in curr.parents:
					dir_sizes[parent] += size
	return dir_sizes

def part_one(filename: str) -> int:
	dir_sizes = get_dir_sizes_dict(filename)
	return sum(v for v in dir_sizes.values() if v <= 100000)

# find out the smallest to delete which larger than "need"
def part_two(filename: str) -> int:
	dir_sizes = get_dir_sizes_dict(filename)
	vals = sorted(dir_sizes.values())
	need = 30000000 - (70000000 - vals[-1])
	idx = bisect_left(vals, need)
	return vals[idx]

if __name__=="__main__":
	input_file = "./2022/day_07/input.txt"
	print("---part one---")
	print(part_one(input_file))
	print("---part two---")
	print(part_two(input_file))

# ref: https://github.com/jmerle/advent-of-code-2022/blob/master/src/day07/part1.py