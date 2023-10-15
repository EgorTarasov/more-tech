from app.api.login import login_for_access_token
from fastapi import APIRouter, Depends, Request, HTTPException
from fastapi.responses import RedirectResponse
from app.dependencies import get_db
from app.db import MongoDB
import app.models
from app.forms.login import LoginForm

from fastapi.templating import Jinja2Templates

templates = Jinja2Templates(directory="templates")

router = APIRouter(prefix="/login")

# auth routes
# login
@router.get("/")
def login(request: Request):
    return templates.TemplateResponse("auth/login.html", {"request": request})


@router.post("/")
async def login(request: Request, db: MongoDB = Depends(get_db)):
    form = LoginForm(request)

    await form.load_data()
    print(form.__dict__)
    if await form.is_valid():
        try:
            form.__dict__.update(msg="Login Successful :)")
            response = templates.TemplateResponse("auth/login.html", form.__dict__)
            login_for_access_token(response=response, form_data=form, db=db)
            return response
        except HTTPException:
            form.__dict__.update(msg="")
            form.__dict__.get("errors").append("Incorrect Email or Password")
            return templates.TemplateResponse("auth/login.html", form.__dict__)
    return templates.TemplateResponse("auth/login.html", form.__dict__)


# stats routes
