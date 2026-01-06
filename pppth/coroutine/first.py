import asyncio


async def square(number: int) -> int:
    return number*number

result = asyncio.run(square(10))
print(result)

"""
    The event loop is the central orchestrator in asyncio that manages and executes asynchronous tasks in a single thread. It is the core mechanism that allows a program to handle multiple I/O-bound operations concurrently without traditional blocking, resulting in efficient and scalable code. 
How the Event Loop Works
The event loop operates on a cooperative multitasking model. Coroutines (async def functions) explicitly yield control back to the event loop when they encounter an await expression for an I/O operation (like a network request or a file read). 

Scheduling Tasks: Coroutines are wrapped in Task objects and scheduled on the event loop.

Execution and Yielding: The event loop runs a task until it hits an await statement. At that point, the task is suspended, and control returns to the loop.

Monitoring Events: The event loop uses low-level operating system mechanisms (like epoll on Linux or IOCP on Windows) to monitor multiple I/O operations simultaneously.

Resuming Tasks: When an I/O operation finishes (e.g., data is received from a socket or a timer expires), the event loop receives a signal and marks the corresponding suspended task as ready to resume.

Cycle Repetition: The loop then picks the next ready task and repeats the process. This cycle continues until all tasks are complete. 
This process effectively switches between tasks during their "wait time," giving the illusion of simultaneous execution within a single thread. 
Key Interactions with the Event Loop
Application developers typically use high-level functions to interact with the event loop, rather than managing the loop manually. 

asyncio.run(main()): This is the primary entry point for asyncio applications. It creates an event loop, runs the main coroutine until it completes, and then automatically closes the loop.

asyncio.create_task(coro): This schedules a coroutine to run concurrently on the event loop in the background. It returns a Task object that can be awaited later.

await: The crucial keyword that signals to the event loop that the current coroutine is waiting for a result from an awaitable (another coroutine, task, or future), allowing the loop to switch to other tasks in the meantime.

loop.run_in_executor(): Used to run blocking I/O or CPU-bound code in a separate thread or process pool without blocking the main event loop thread. 
For more details on the specific methods and behaviors, refer to the Python documentation on the event loop. """