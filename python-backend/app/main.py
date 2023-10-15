from fastapi import FastAPI
from fastapi.staticfiles import StaticFiles
from fastapi.middleware.cors import CORSMiddleware
from contextlib import asynccontextmanager
from app.api import api_router
from app.routes import html_router
from app.utils import Hasher


@asynccontextmanager
async def lifespan(app: FastAPI):
    # startup

    yield
    # shutdown


def create_app():
    _app = FastAPI(
        title="Vicktor The Bear Admin panel",
        description="Admin panel for Vicktor The Bear",
        version="0.0.1",
        lifespan=lifespan,
    )

    _app.add_middleware(
        CORSMiddleware,
        allow_origins=["*"],
        allow_credentials=True,
        allow_methods=["*"],
        allow_headers=["*"],
    )

    _app.mount("/static", StaticFiles(directory="static"), name="static")
    _app.include_router(api_router, prefix="/api")
    _app.include_router(html_router)

    print(Hasher.get_password_hash("Test123456"))
    return _app


app = create_app()
