import asyncio
import json
import os
from app.services.pubsub import RedisPubSub
from app.core.config import settings

async def handle_order_event(message: dict):
    # Placeholder: implement idempotent processing, DB writes, external calls, retries
    print("[worker] Received order event:", message)
    # Simulate processing delay
    await asyncio.sleep(0.1)

async def run_worker():
    print("Worker starting. Connecting to Redis:", settings.REDIS_URL)
    rp = RedisPubSub()
    await rp.init()
    sub = await rp.subscribe("orders")
    print("Subscribed to 'orders' channel. Waiting for messages...")
    try:
        while True:
            msg = await sub.get_message(ignore_subscribe_messages=True, timeout=1.0)
            if msg is None:
                await asyncio.sleep(0.1)
                continue
            # msg example: {'type': 'message', 'pattern': None, 'channel': 'orders', 'data': '<payload>'}
            data = msg.get("data")
            try:
                obj = json.loads(data)
            except Exception:
                print("[worker] Failed to parse message:", data)
                continue
            # Process message safely
            try:
                await handle_order_event(obj)
            except Exception as e:
                print("[worker] Error processing message:", e)
    except asyncio.CancelledError:
        print("Worker cancelled, exiting.")
    finally:
        await sub.unsubscribe("orders")
        await rp._redis.close()

if __name__ == '__main__':
    import asyncio
    asyncio.run(run_worker())
