<!-- Generated guidance for AI coding assistants. Keep concise and actionable. -->
# Copilot instructions — fastApi project

This repository is a very small FastAPI project. The guidance below is derived from the current workspace layout and the observable patterns (entrypoint, naming, and conventions). Use these to be immediately productive when adding features, tests, or CI automation.

## Big picture
- Single microservice-style Python project using FastAPI. The immediate entrypoint is `first/main.py` which creates a `FastAPI()` app instance.
- Project structure is minimal: `first/` contains `main.py`. Expect future modules (routers, models, services) to be placed under `first/` or sibling packages.

## Key files and patterns
- `first/main.py`: Application factory/entrypoint. When adding endpoints, register them on the `app` object here or import routers and include with `app.include_router(...)`.

## Development workflows
- Run the app locally with uvicorn (not present in the repo yet). Suggested command:

```powershell
python -m uvicorn first.main:app --reload
```

- If adding dependencies, update a `requirements.txt` or `pyproject.toml` at repo root and include pinned versions.

## Project-specific conventions
- Keep modules under `first/` for now. New functionality should follow package-style layout: `first/routers/*.py`, `first/services/*.py`, `first/models/*.py`.
- Import the FastAPI instance from `first.main` for tests or for mounting sub-apps.

## Testing and quality
- No tests are present. When adding tests, use `pytest` and place tests under `tests/` at the repository root. Example test import:

```python
from first.main import app
from fastapi.testclient import TestClient

client = TestClient(app)

def test_root():
    resp = client.get('/')
    assert resp.status_code in (200, 404)
```

## Integration points and dependencies
- The project uses `fastApi` import in `first/main.py`. Verify whether the project expects `fastapi` (lowercase) or a local module `fastApi`. Current code: `from fastApi import FastAPI`. If this is not a local module, replace with `from fastapi import FastAPI` and add `fastapi` to dependencies.

## When editing code — do this first
1. Run static checks (flake8/ruff/mypy) if you add them to the project. Add a simple `pyproject.toml` if configuring tools.
2. Run the app with uvicorn to smoke-test endpoints.
3. Add at least one unit test for new behavior and run `pytest`.

## Examples to follow
- To add a router: create `first/routers/items.py` with a FastAPI `APIRouter()` and then in `first/main.py` add `app.include_router(items.router)`.
- To add models: create Pydantic models under `first/models.py` or `first/models/*.py`.

## Notes and gotchas
- The import in `first/main.py` uses `fastApi` (capital A). This is likely a typo. Confirm whether `fastApi` is a local package or meant to be the external `fastapi` package. Fixing it will be required to run the app.
- The repository is extremely small; avoid speculative large refactors. Prefer adding files alongside `first/` following the simple package layout.

---
If anything in this file looks incorrect or incomplete, point to the file(s) you want scanned next (for example other packages or a requirements file) and I'll update the guidance.
