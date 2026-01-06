import asyncio


async def square(number: int) -> int:
    return number*number


async def main() -> None:
    x = await square(10)
    print(f'x={x}')

    y = await square(5)
    print(f'y={y}')

    print(f'total={x+y}')

if __name__ == '__main__':
    asyncio.run(main())


"""
    The event loop is the central orchestrator in asyncio that manages and executes asynchronous tasks in a single thread. 
    It is the core mechanism that allows a program to handle multiple I/O-bound operations concurrently without traditional 
    blocking, resulting in efficient and scalable code.

        First, call the square() coroutine using the await keyword. The await keyword will pause the execution of the main() coroutine, wait for the square() coroutine to complete, and return the result:

        x = await square(10)
        print(f'x={x}')
        Code language: Python (python)
        Second, call the square() coroutine a second time using the await keyword:

        y = await square(5)
        print(f'y={y}')
        Code language: Python (python)
        Third, display the total:

        print(f'total={x+y}')
        Code language: Python (python)
        The following statement uses the run() function to execute the main() coroutine and manage the event loop:

        asyncio.run(main())
        Code language: Python (python)
        So far, our program executes like a synchronous program. It doesn’t reveal the power of the asynchronous programming model.

    """