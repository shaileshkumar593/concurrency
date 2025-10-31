"""
    Using asyncio.create_task() for Background Work
    Create background jobs
✅ Continue main flow concurrently

"""

import asyncio

async def background_job(name):
    for i in range(3):
        print(f"{name} - iteration {i}")
        await asyncio.sleep(1)

async def main():
    task = asyncio.create_task(background_job("Worker"))
    print("Main doing other work...")
    await asyncio.sleep(2)
    print("Main waiting for background job")
    await task

asyncio.run(main())
