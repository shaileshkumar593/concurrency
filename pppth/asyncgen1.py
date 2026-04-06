import asyncio

async def g(n):
    for i in range(n):
        await asyncio.sleep(1)  # simulate async I/O (DB/API call)
        yield i

async def main():
    async for value in g(5):
        print("Streamed:", value)

if __name__ == "__main__":
    asyncio.run(main())


#Example with Multiple Yields + await (Realistic)