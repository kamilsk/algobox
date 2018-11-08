from unittest import TestCase
from solution import fix


class Test(TestCase):
    def test_fix(self):
        tests = [
            {'name': 'case#01', 'input': '}[[([{[]}', 'expected': '{}[[([{[]}])]]'},
            {'name': 'case#02', 'input': '{[({[({[(', 'expected': '{[({[({[()]})]})]}'},
            {'name': 'case#03', 'input': ')]})]})]}', 'expected': '{[({[({[()]})]})]}'},
            {'name': 'case#04', 'input': '{][[[[{}[]', 'expected': 'IMPOSSIBLE'},
        ]
        for test in tests:
            self.assertEqual(fix(test['input']), test['expected'], test['name'])
