FastAPI Orders Service (Modular)
=================================
Structure:
- app/
  - main.py
  - core/ (config, security)
  - db/ (session, base)
  - models/ (SQLAlchemy models)
  - schemas/ (Pydantic schemas)
  - api/ (routers: auth, orders)
  - services/ (business logic)
  - tests/
- Dockerfile
- docker-compose.yml
- .env.example

Quick start (with Docker):
  docker-compose up --build

Development (local):
  python -m venv .venv
  .venv/bin/pip install -r requirements.txt
  uvicorn app.main:app --reload --host 0.0.0.0 --port 8000

Notes:
- This scaffold uses Postgres. For quick testing, modify DATABASE_URL to sqlite.

Redis & Worker
--------------
- Redis is added as a service in docker-compose.
- A `worker` service runs the async subscriber `app/worker/worker.py` and listens on the `orders` channel.
- On order creation, the API publishes a JSON event to Redis channel `orders`.
- The worker prints and processes events; replace `handle_order_event` with real business logic (payments, inventory, notifications).
