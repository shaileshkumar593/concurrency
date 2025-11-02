from pydantic import BaseModel, Field, EmailStr
from typing import Optional

class UserCreate(BaseModel):
    username: str = Field(..., min_length=3, max_length=128)
    password: str = Field(..., min_length=6)
    email: Optional[EmailStr] = None

class TokenOut(BaseModel):
    access_token: str
    token_type: str = "bearer"
