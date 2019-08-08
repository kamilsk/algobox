# python3

from collections import namedtuple, deque
from unittest import TestCase

request = namedtuple('request', 'arrived_at time_to_process')
response = namedtuple('response', 'was_dropped started_at')
test = namedtuple('test', 'buffer_size requests expected')


class Buffer:
    def __init__(self, cap: int):
        self.queue = deque(maxlen=cap)

    def process(self, req: request) -> response:
        while self.queue:
            timestamp = self.queue.popleft()
            if req.arrived_at >= timestamp:
                continue
            self.queue.appendleft(timestamp)
            break

        if not self.queue:
            self.queue.append(req.arrived_at + req.time_to_process)
            return response(False, req.arrived_at)

        if len(self.queue) == self.queue.maxlen:
            return response(True, -1)

        last = self.queue.pop()
        self.queue.append(last)
        self.queue.append(last + req.time_to_process)
        return response(False, last)


class Test(TestCase):
    def test_network_simulation(self):
        tests = [
            # samples
            test(1, [], []),
            test(1, [request(0, 0)], [0]),
            test(1, [
                request(0, 1),
                request(0, 1),
            ], [0, -1]),
            test(1, [
                request(0, 1),
                request(1, 1),
            ], [0, 1]),

            # acceptance
            test(2, [
                request(0, 2),
                request(1, 4),
                request(5, 3),
            ], [0, 2, 6]),
        ]
        for i, t in enumerate(tests):
            b = Buffer(t.buffer_size)
            obtained = [r.started_at if not r.was_dropped else -1 for r in [b.process(req) for req in t.requests]]
            self.assertEqual(t.expected, obtained, msg='at {} position'.format(i))


if __name__ == '__main__':
    size, n = map(int, input().split())

    requests = []
    for _ in range(n):
        arrived_at, time_to_process = map(int, input().split())
        requests.append(request(arrived_at, time_to_process))

    buf = Buffer(size)
    for resp in [buf.process(req) for req in requests]:
        print(resp.started_at if not resp.was_dropped else -1)
