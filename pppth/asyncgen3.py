import asyncio

async def g():
    yield 1
    yield 2
    yield 3

async def main():
    async for x in g():
        print(x)

asyncio.run(main())


"""
    async for x in g():
    print(x)


    is equivalent to:

it = g().__aiter__()

while True:
    try:
        x = await it.__anext__()
        print(x)
    except StopAsyncIteration:
        break

        



        import asyncio

class MyAsyncGen:
    def __init__(self):
        self.data = [1, 2, 3]
        self.index = 0

    # called first
    def __aiter__(self):
        return self

    # called on every iteration
    async def __anext__(self):
        if self.index >= len(self.data):
            raise StopAsyncIteration

        await asyncio.sleep(0.5)  # simulate async work
        value = self.data[self.index]
        self.index += 1
        return value

async def main():
    async for x in MyAsyncGen():
        print(x)

asyncio.run(main())





When this runs:

async for x in MyAsyncGen():
Step 1
it = MyAsyncGen().__aiter__()
Step 2 (loop)
x = await it.__anext__()
Step 3
returns next value → printed

repeats

Step 4
raise StopAsyncIteration
→ loop ends


uvloop is a fast, drop-in replacement of the built-in asyncio event loop. uvloop is implemented in Cython and uses libuv under the hood.

The project documentation can be found here. Please also check out the wiki.


"""