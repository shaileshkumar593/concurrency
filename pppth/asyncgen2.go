# Instead of async for, you can use __anext__():

import asyncio

async def g(x):
    yield x

async def main():
    gen = g(100)

    try:
        value = await gen.__anext__()
        print(value)
    except StopAsyncIteration:
        print("Done")

asyncio.run(main())