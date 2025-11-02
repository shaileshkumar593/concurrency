from pydantic import BaseSettings, Field
from typing import Optional

class Settings(BaseSettings):
    APP_NAME: str = Field("orders-service", env="APP_NAME")
    DATABASE_URL: str = Field(..., env="DATABASE_URL")
    SECRET_KEY: str = Field(..., env="SECRET_KEY")
    ACCESS_TOKEN_EXPIRE_MINUTES: int = Field(60, env="ACCESS_TOKEN_EXPIRE_MINUTES")
    BACKEND_CORS_ORIGINS: Optional[str] = Field("http://localhost:3000", env="BACKEND_CORS_ORIGINS")

    REDIS_URL: str = Field('redis://redis:6379/0', env='REDIS_URL')\n\n    class Config:
        env_file = ".env"
        env_file_encoding = 'utf-8'

settings = Settings()
