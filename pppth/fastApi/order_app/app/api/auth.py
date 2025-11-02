from fastapi import APIRouter, Depends, HTTPException, status
from sqlalchemy.future import select
from sqlalchemy.ext.asyncio import AsyncSession
from app.db.session import get_db
from app.schemas.auth import UserCreate, TokenOut
from app.models.user import User
from app.core.security import hash_password, verify_password, create_access_token
from fastapi.security import OAuth2PasswordRequestForm

router = APIRouter()

@router.post('/register', status_code=201)
async def register(user_in: UserCreate, db: AsyncSession = Depends(get_db)):
    q = await db.execute(select(User).where(User.username == user_in.username))
    existing = q.scalars().first()
    if existing:
        raise HTTPException(status_code=400, detail='username exists')
    user = User(username=user_in.username, email=user_in.email, hashed_password=hash_password(user_in.password))
    db.add(user)
    await db.commit()
    await db.refresh(user)
    return {'username': user.username}

@router.post('/token', response_model=TokenOut)
async def token(form_data: OAuth2PasswordRequestForm = Depends(), db: AsyncSession = Depends(get_db)):
    q = await db.execute(select(User).where(User.username == form_data.username))
    user = q.scalars().first()
    if not user or not verify_password(form_data.password, user.hashed_password):
        raise HTTPException(status_code=status.HTTP_401_UNAUTHORIZED, detail='invalid credentials')
    token = create_access_token(user.username)
    return {'access_token': token, 'token_type': 'bearer'}
