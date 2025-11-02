from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from app.core.config import settings
from app.api import router\nfrom app.services.pubsub import pubsub

from app.db.session import engine, Base

app = FastAPI(title=settings.APP_NAME)

origins = settings.BACKEND_CORS_ORIGINS.split(",") if isinstance(settings.BACKEND_CORS_ORIGINS, str) else settings.BACKEND_CORS_ORIGINS
app.add_middleware(
    CORSMiddleware,
    allow_origins=origins or ["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

app.include_router(router, prefix="/api")

@app.on_event("startup")
async def on_startup():
    async with engine.begin() as conn:
        await conn.run_sync(Base.metadata.create_all)

@app.get("/healthz")
async def health():
    return {"status": "ok"}
