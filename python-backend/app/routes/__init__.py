from fastapi import APIRouter
from . import login
from . import panel


html_router = APIRouter()
html_router.include_router(login.router)
html_router.include_router(panel.router)
