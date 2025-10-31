import asyncio

async def task(n):
    print(f"Task {n} started")
    await asyncio.sleep(1)
    print(f"Task {n} done")
    return n * n

async def main():
    results = await asyncio.gather(
        task(1),
        task(2),
        task(3)
    )
    print("Results:", results)

asyncio.run(main())


# Concurrent Async Tasks with gather()
# Runs all tasks concurrently
# ✅ Returns results in order of submission