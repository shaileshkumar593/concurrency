import asyncio

async def greet(name):
    print(f"Hello, {name}!")
    await asyncio.sleep(1)
    print(f"Goodbye, {name}!")

async def main():
    await greet("Alice")
    await greet("Bob")

asyncio.run(main())
