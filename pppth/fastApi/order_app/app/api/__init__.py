from fastapi import APIRouter
from app.api import auth, orders
router = APIRouter()
router.include_router(auth.router, prefix="/auth", tags=["auth"])
router.include_router(orders.router, prefix="/orders", tags=["orders"])
