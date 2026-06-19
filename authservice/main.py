import os

from dotenv import load_dotenv
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from fastapi.staticfiles import StaticFiles

load_dotenv()

from database.db import create_db_and_tables  # noqa: E402
from routes.auth import router as auth_router  # noqa: E402

app = FastAPI(title="Sistema de Chat - Auth Service", version="1.0.0")

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.on_event("startup")
def on_startup() -> None:
    """Crea la base de datos y las tablas si no existen."""
    create_db_and_tables()


@app.get("/health")
def health() -> dict:
    """Endpoint simple para verificar que el servicio está activo."""
    return {"status": "ok", "service": "authservice"}

app.include_router(auth_router)

_BASE_DIR = os.path.dirname(os.path.abspath(__file__))
_CLIENT_DIR = os.getenv("CLIENT_DIR") or os.path.join(_BASE_DIR, "..", "cliente")
if os.path.isdir(_CLIENT_DIR):
    app.mount("/", StaticFiles(directory=_CLIENT_DIR, html=True), name="cliente")
