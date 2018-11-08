def fix(string: str) -> str:
    brackets = {
        '[': {'opening': True,  'reverse': ']'},
        '{': {'opening': True,  'reverse': '}'},
        '(': {'opening': True,  'reverse': ')'},
        ']': {'opening': False, 'reverse': '['},
        '}': {'opening': False, 'reverse': '{'},
        ')': {'opening': False, 'reverse': '('},
    }
    front, back = [], []
    for char in string:
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
    buf.append(string)
    while back:
        buf.append(brackets[back.pop()]['reverse'])
    return ''.join(buf)


def main():
    print(fix(input()))


if __name__ == "__main__":
    main()
