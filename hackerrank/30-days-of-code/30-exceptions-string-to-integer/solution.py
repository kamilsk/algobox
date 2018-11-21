S = input().strip()
code = None
try:
    code = int(S)
except ValueError:
    print('Bad String')
    exit(1)
print(code)
