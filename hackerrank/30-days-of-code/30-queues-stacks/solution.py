class Alias(object):
    def __init__(self, *aliases):
        self.aliases = set(aliases)

    def __call__(self, f):
        f.aliases = self.aliases
        return f


def aliased(aliased_class):
    original_methods = aliased_class.__dict__.copy()
    for name, method in original_methods.items():
        if hasattr(method, 'aliases'):
            for alias in method.aliases - set(original_methods):
                setattr(aliased_class, alias, method)
    return aliased_class


@aliased
class Solution:
    def __init__(self):
        self.stack = []
        self.queue = []

    @Alias('pushCharacter')
    def push_character(self, ch: str):
        self.stack.append(ch)

    @Alias('popCharacter')
    def pop_character(self) -> str:
        return self.stack.pop()

    @Alias('enqueueCharacter')
    def enqueue_character(self, ch: str):
        self.queue.append(ch)

    @Alias('dequeueCharacter')
    def dequeue_character(self) -> str:
        return self.queue.pop(0)
