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
            buf = []
            while front:
                buf.append(opening.get(front.pop()))
            buf.append(line)
            while back:
                buf.append(closing.get(back.pop()))
            print(''.join(buf))


opening = {'}': '{', ']': '[', ')': '('}
closing = {v: k for k, v in opening.items()}

if __name__ == "__main__":
    main()
