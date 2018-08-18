import fileinput


def main():
    for line in fileinput.input():
        line = line.strip()
        possible, front, back = True, [], []
        for bracket in line:
            if not possible:
                break
            if bracket in closing:
                back.append(bracket)
                continue
            if not back:
                front.append(bracket)
            else:
                possible = opening.get(bracket) is back.pop()
        if not possible:
            print("IMPOSSIBLE")
        else:
            while front:
                cl = front.pop()
                print(opening.get(cl), end='')
            print(line, end='')
            while back:
                op = back.pop()
                print(closing.get(op), end='')
            print()


opening = {'}': '{', ']': '[', ')': '('}
closing = {'{': '}', '[': ']', '(': ')'}

if __name__ == "__main__":
    main()
