from sqlalchemy import Column, Integer, ForeignKey, DateTime, Float, JSON
from sqlalchemy.sql import func
from app.db.session import Base

class Order(Base):
    __tablename__ = 'orders'
    id = Column(Integer, primary_key=True, index=True)
    owner_id = Column(Integer, ForeignKey('users.id'), nullable=False)
    created_at = Column(DateTime(timezone=True), server_default=func.now())
    total = Column(Float, default=0.0)
    payload = Column(JSON, nullable=False)
