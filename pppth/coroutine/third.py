import asyncio
import time


async def call_api(message, result=1000, delay=3):
    print(message)
    await asyncio.sleep(delay)
    return result


async def main():
    start = time.perf_counter()

    price = await call_api('Get stock price of GOOG...', 300)
    print(price)

    price = await call_api('Get stock price of APPL...', 400)
    print(price)

    end = time.perf_counter()
    print(f'It took {round(end-start,0)} second(s) to complete.')

asyncio.run(main())


"""
    In this example, the two calls to call_api are executed sequentially. 
    The first call to call_api must complete before the second call begins. 
    As a result, the total execution time is approximately the sum of the delays of both calls (3 seconds + 3 seconds = 6 seconds).
    
        In this example, we call a coroutine directly and don’t put it on the event loop to run.
          Instead, we get a coroutine object and use the await keyword to execute it and get a result.

          In other words, we use async and await to write asynchronous code but can’t run it concurrently. 
          To run multiple operations concurrently, we’ll need to use something called tasks.

          
    """