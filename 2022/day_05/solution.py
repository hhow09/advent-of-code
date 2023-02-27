def parseState(filename: str) -> list[list[str]]:
	with open(filename) as f:
		lines = f.read().split('\n\n')
		initState = lines[0]
		rows = initState.split('\n')
		stacks = (len(rows[0]) + 1)//4
		l = [[] for _ in range(stacks)]
		for row in rows:
			for i, char in enumerate(row):
				if char.isalpha():
					l[i//4].append(char)
		for i in range(len(l)):
			l[i].reverse()
		return l

def parseOp(op: str) -> list[int]:
	op = op[5:]
	arr = op.split(' ')
	return map(int, [arr[0], arr[2], arr[4]])

def applyOps(state: list[list[str]], filename: str, rev: bool) -> list[list[str]]:
	with open(filename) as f:
		lines = f.read().split('\n\n')
		ops = lines[1].split('\n')
		for op in ops:
			count, f, t = parseOp(op)
			last = state[f-1][-count:]
			if rev is True:
				state[t-1] += reversed(state[f-1][-count:])
			else:
				state[t-1] += state[f-1][-count:]
			state[f-1] = state[f-1][:-count]


def part_one(file: str) -> str:
	state = parseState(file)
	applyOps(state, file, True)
	return "".join(map(lambda l: l[-1], state))

def part_two(file: str) -> str:
	state = parseState(file)
	applyOps(state, file, False)
	return "".join(map(lambda l: l[-1], state))


if __name__=="__main__":
	input_file = "./2022/day_05/input.txt"
	print("---part one---")
	print(part_one(input_file))
	print("---part two---")
	print(part_two(input_file))