import fileinput


def main():
    for line in fileinput.input():
        print(fix(line))


def fix(input):
    brackets = {
        '[': {'opening': True,  'reverse': ']'},
        '{': {'opening': True,  'reverse': '}'},
        '(': {'opening': True,  'reverse': ')'},
        ']': {'opening': False, 'reverse': '['},
        '}': {'opening': False, 'reverse': '{'},
        ')': {'opening': False, 'reverse': '('},
    }
    front, back = [], []
    for char in input:
        bracket = brackets[char]
        if bracket['opening']:
            back.append(char)
            continue
        if not back:
            front.append(char)
            continue
        if not bracket['reverse'] is back.pop():
            return 'IMPOSSIBLE'
    buf = []
    while front:
        buf.append(brackets[front.pop()]['reverse'])
    buf.append(input)
    while back:
        buf.append(brackets[back.pop()]['reverse'])
    return ''.join(buf)


if __name__ == "__main__":
    main()
