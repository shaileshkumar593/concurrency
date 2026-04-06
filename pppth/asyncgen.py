import asyncio

# async generator
async def g(x):
    print("Generator started")
    yield x
    print("Generator finished")

# consumer
async def main():
    async for value in g(10):
        print("Received:", value)

# run event loop
if __name__ == "__main__":
    asyncio.run(main())