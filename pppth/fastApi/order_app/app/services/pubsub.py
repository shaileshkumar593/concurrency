import asyncio
import json
from typing import Callable, Awaitable, Optional
import aioredis
from app.core.config import settings

class RedisPubSub:
    def __init__(self):
        self._redis = None
        self._pub = None
        self._sub = None
        self._listeners = {}

    async def init(self):
        # lazy init single client for pub/sub patterns
        self._redis = await aioredis.from_url(settings.REDIS_URL, decode_responses=True)
        self._pub = self._redis
        # subscriber will create its own pubsub when worker starts

    async def publish(self, channel: str, message: dict):
        if self._pub is None:
            await self.init()
        payload = json.dumps(message)
        await self._pub.publish(channel, payload)

    async def subscribe(self, channel: str):
        # returns a new dedicated PubSub for consumption
        if self._redis is None:
            await self.init()
        sub = self._redis.pubsub()
        await sub.subscribe(channel)
        return sub

# single instance to import
pubsub = RedisPubSub()
