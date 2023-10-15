from datetime import timedelta
from fastapi import APIRouter
from fastapi import Depends
from fastapi import HTTPException
from fastapi import Response
from fastapi import status
from fastapi.security import OAuth2PasswordRequestForm
from jose import jwt
from jose import JWTError


from app.models.token import Token
from app.api.utils import OAuth2PasswordBearerWithCookie
from app.utils import Hasher, create_access_token
from app.dependencies import get_db
from app.db import MongoDB


# from fastapi.security import OAuth2PasswordBearer


router = APIRouter()


def authenticate_user(username: str, password: str, db: MongoDB = Depends(get_db)):
    # FIXME MONGO
    user = db.get_user(username)
    print(user)
    if not user:
        return False
    if not Hasher.verify_password(password, user["hashed_password"]):
        return False
    return user


@router.post("/token", response_model=Token)
def login_for_access_token(
    response: Response,
    form_data: OAuth2PasswordRequestForm = Depends(),
    db: MongoDB = Depends(get_db),
):
    user = authenticate_user(form_data.username, form_data.password, db)
    if not user:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Incorrect username or password",
        )
    access_token_expires = timedelta(minutes=60 * 24 * 7)
    access_token = create_access_token(
        data={"sub": user["email"]}, expires_delta=access_token_expires
    )
    response.set_cookie(
        key="access_token", value=f"Bearer {access_token}", httponly=True
    )
    return {"access_token": access_token, "token_type": "bearer"}


oauth2_scheme = OAuth2PasswordBearerWithCookie(tokenUrl="/login/token")


def get_current_user_from_token(
    token: str = Depends(oauth2_scheme), db: MongoDB = Depends(get_db)
):
    credentials_exception = HTTPException(
        status_code=status.HTTP_401_UNAUTHORIZED,
        detail="Could not validate credentials",
    )
    try:
        payload = jwt.decode(token, "blabla", algorithms=["HS256"])
        username: str = payload.get("sub")
        print("username/email extracted is ", username)
        if username is None:
            raise credentials_exception
    except JWTError:
        raise credentials_exception
    user = db.get_user(username)
    if user is None:
        raise credentials_exception
    return user
