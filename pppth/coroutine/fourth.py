"""
Docstring for pppth.coroutine.fourth

  Task Creation and Management with asyncio.create_task()

    In asyncio, tasks are used to schedule coroutines to run concurrently on the event loop. 
    The asyncio.create_task() function is a high-level API that wraps a coroutine in a Task object and schedules it to run on the event loop.
    Here’s how to use asyncio.create_task() to create and manage tasks:
     
    A task is a wrapper of a coroutine that schedules the coroutine to run on the event loop as soon as possible.

        The scheduling and execution occur in a non-blocking manner. In other words, you can create a task and execute other code instantly while the task is running.

        Notice that the task is different from the await keyword that blocks the entire coroutine until the operation completes with a result.

        It’s important that you can create multiple tasks and schedule them to run instantly on the event loop at the same time.

        To create a task, you pass a coroutine to the create_task() function of the asyncio package. The create_task() function returns a Task object.

        The following program illustrates how to create two tasks that schedule and execute the call_api() coroutine:

"""

import asyncio
import time


async def call_api(message, result=1000, delay=3):
    print(message)
    await asyncio.sleep(delay)
    return result


async def main():
    start = time.perf_counter()

    task_1 = asyncio.create_task(
        call_api('Get stock price of GOOG...', 300)
    )

    task_2 = asyncio.create_task(
        call_api('Get stock price of APPL...', 300)
    )

    price = await task_1
    print(price)

    price = await task_2
    print(price)

    end = time.perf_counter()
    print(f'It took {round(end-start,0)} second(s) to complete.')


asyncio.run(main())