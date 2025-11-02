"""
    Parallel Tasks with async for and async with

"""

# async for
import asyncio

async def countdown(n):
    while n > 0:
        yield n
        await asyncio.sleep(1)
        n -= 1

async def main():
    async for i in countdown(3):
        print(i)

asyncio.run(main())


# async with (for managing async resources)

import asyncio

class AsyncResource:
    async def __aenter__(self):
        print("Connecting...")
        await asyncio.sleep(1)
        return self

    async def __aexit__(self, exc_type, exc, tb):
        print("Disconnecting...")
        await asyncio.sleep(1)

async def main():
    async with AsyncResource() as res:
        print("Using resource...")

asyncio.run(main())

