from fastapi import APIRouter, Depends, HTTPException, status
from sqlalchemy.future import select
from sqlalchemy.ext.asyncio import AsyncSession
from app.db.session import get_db
from app.schemas.order import OrderCreate, OrderOut
from app.models.order import Order
from app.models.user import User\nfrom app.services.pubsub import pubsub\nimport asyncio
from app.core.config import settings
from fastapi.security import OAuth2PasswordBearer
from jose import JWTError, jwt
from typing import List
from sqlalchemy import select as sselect

oauth2_scheme = OAuth2PasswordBearer(tokenUrl='/api/auth/token')

router = APIRouter()

async def get_current_user(token: str = Depends(oauth2_scheme), db: AsyncSession = Depends(get_db)) -> User:
    credentials_exception = HTTPException(status_code=status.HTTP_401_UNAUTHORIZED, detail='Could not validate')
    try:
        payload = jwt.decode(token, settings.SECRET_KEY, algorithms=['HS256'])
        username: str = payload.get('sub')
        if username is None:
            raise credentials_exception
    except JWTError:
        raise credentials_exception
    q = await db.execute(sselect(User).where(User.username == username))
    user = q.scalars().first()
    if user is None:
        raise credentials_exception
    return user

@router.post('/', response_model=OrderOut, status_code=201)
async def create_order(order_in: OrderCreate, current_user: User = Depends(get_current_user), db: AsyncSession = Depends(get_db)):
    order = Order(owner_id=current_user.id, payload=order_in.payload, total=order_in.total)
    db.add(order)
    await db.commit()
    await db.refresh(order)
    return order

@router.get('/', response_model=List[OrderOut])
async def list_orders(current_user: User = Depends(get_current_user), db: AsyncSession = Depends(get_db)):
    q = await db.execute(select(Order).where(Order.owner_id == current_user.id))
    orders = q.scalars().all()
    return orders
