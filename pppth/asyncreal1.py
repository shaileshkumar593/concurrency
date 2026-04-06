import time

def count():
    print("One")
    time.sleep(1)
    print("Two")
    time.sleep(1)

def main():
    for _ in range(3):
        count()

if __name__ == "__main__":
    start = time.perf_counter()
    main()
    elapsed = time.perf_counter() - start
    print(f"{__file__} executed in {elapsed:0.2f} seconds.")



"""
    Now, you use the async keyword to turn count() into a coroutine function that prints One, 
    waits for one second, then prints Two, and waits another second. You use the await keyword to 
    await the execution of asyncio.sleep(). This gives the control back to the program’s event loop, 
    saying: I will sleep for one second. Go ahead and run something else in the meantime.

    The main() function is another coroutine function that uses asyncio.gather() to run three instances of count() 
    concurrently. You use the asyncio.run() function to launch the event loop and execute main().


    Cooperative multitasking, also known as non-preemptive multitasking, is a computer multitasking technique in which the operating system never initiates a context switch from a running process to another process. Instead, in order to run multiple applications concurrently, processes voluntarily yield control periodically or when idle or logically blocked. This type of multitasking is called cooperative because all programs must cooperate for the scheduling scheme to work.

    In this scheme, the process scheduler of an operating system is known as a cooperative scheduler whose role is limited to starting the processes and letting them return control back to it voluntarily.[1][2]

    This is related to the asynchronous programming approach.
"""


"""

    While using time.sleep() and asyncio.sleep() may seem banal, they serve as stand-ins for 
    time-intensive processes that involve wait time. A call to time.sleep() can represent a
    time-consuming blocking function call, while asyncio.sleep() is used to stand in for a 
    non-blocking call that also takes some time to complete.


     the benefit of awaiting something, including asyncio.sleep(), is that the surrounding function can temporarily 
     cede control to another function that’s more readily able to do something immediately. In contrast, time.sleep() 
     or any other blocking call is incompatible with asynchronous Python code because it stops everything in its tracks 
     for the duration of the sleep time.
"""