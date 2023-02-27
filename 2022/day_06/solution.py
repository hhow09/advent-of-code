import collections

def process(file: str, markerlen: int) -> int:
	with open(file) as f:
		signals = f.read()
		dq = collections.deque()
		s = set()
		for i, char in enumerate(signals):
			while char in s:
				s.remove(dq.popleft())
			dq.append(char)
			s.add(char)
			if len(dq)==markerlen:
				return i+1	

def part_one(file: str) -> int:
	return process(file, 4)


def part_two(file: str) -> str:
	return process(file, 14)

if __name__=="__main__":
	input_file = "./2022/day_06/input.txt"
	print("---part one---")
	print(part_one(input_file)) # 1855
	print("---part two---") 
	print(part_two(input_file)) # 3256