from fastapi import APIRouter
import app.api.login as login
import app.api.prototype as prototype

api_router = APIRouter()
api_router.include_router(login.router, prefix="/login", tags=["login"])
api_router.include_router(prototype.router, prefix="/prototype", tags=["prototype"])
