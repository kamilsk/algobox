# python3
import sys


def sum_of_two_digits(reader, writer):
    a, b = map(int, reader.readline().split())
    writer.write(str(a + b) + '\n')


if __name__ == '__main__':
    sum_of_two_digits(sys.stdin, sys.stdout)
