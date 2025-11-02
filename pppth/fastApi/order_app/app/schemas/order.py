from pydantic import BaseModel
from typing import Any, Dict
from datetime import datetime

class OrderCreate(BaseModel):
    payload: Dict[str, Any]
    total: float

class OrderOut(BaseModel):
    id: int
    owner_id: int
    created_at: datetime
    total: float
    payload: Dict[str, Any]

    class Config:
        orm_mode = True
