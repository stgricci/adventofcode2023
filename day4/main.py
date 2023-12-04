def common_elements(list1, list2):
    return [element for element in list1 if element in list2]

lines = open('puzzle_input.txt').readlines()

def part1(f):
    score = 0
    for line in f:
        line = line.strip()
        winning_numbers = [x for x in line.split(":")[1].strip().split("|")[0].split(" ") if x != ""]
        numbers_you_got = [x for x in line.split(":")[1].strip().split("|")[1].split(" ") if x != ""]
        common = [element for element in winning_numbers if element in numbers_you_got]
        n = len(common)
        if n == 0:
            continue

        score += 2**(n-1)
    return score
    
def part2(f):
    cards = [1] * len(f) # start with 1 card each
    for i, line in enumerate(lines):
        x, y = map(str.split, line.split("|"))
        matches = len(set(x) & (set(y)))
        for j in range(i + 1, min(i + 1 + matches, len(lines))):
            cards[j] += cards[i]
    return sum(cards)

print(part1(lines))
print(part2(lines))
