"""
Docstring for pppth.coroutine.sixth

However, if the coroutine() took forever, you would be stuck waiting for the await statement to finish without obtaining any result. 
Additionally, you would have no way to stop it if you wanted to.

To resolve this, you can cancel the task using the cancel() method of the Task object. 
If you cancel a task, it’ll raise the CancelledError exception when you await"""

import asyncio
from asyncio import CancelledError


async def call_api(message, result=1000, delay=3):
    print(message)
    await asyncio.sleep(delay)
    return result


async def main():
    task = asyncio.create_task(
        call_api('Calling API...', result=2000, delay=5)
    )
    #check if the task is not done by calling the done()
    if not task.done():
        print('Cancelling the task...')
        task.cancel()  #  stopping long running task () i.e. starvation of task 

    try:
        await task
    except CancelledError:
        print('Task has been cancelled.')


asyncio.run(main())