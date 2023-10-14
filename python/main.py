from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel, Field
import json
import openai
import time
import datetime
import os
import dotenv


dotenv.load_dotenv()

openai.api_key = os.getenv("API_KEY")


async def predict(user_text: str):

    prompt = f"""
Можешь выделить ключевые слова для TfidfVectorizer из текста?
у меня есть следующие фильтры (классы):банкомат, онлайн, привелегия, прайм, маломобильный, физ лицо, юр лицо
Можешь помочь определить к каким фильтрам относиться  предложение. Ответ предоставь в формате json <название фильтра> :  <True или False> 
"""
    response = openai.ChatCompletion.create(
        model="gpt-3.5-turbo",
        messages=[
            {"role": "user", "content": prompt},
            {"role": "user", "content": user_text},
        ],
        temperature=0,
        max_tokens=1024,
    )
    print(response)
    return json.loads(response["choices"][0]["message"]["content"])  # type:ignore


app = FastAPI()
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


class Coordinates(BaseModel):
    latitude: float
    longitude: float


class Request(BaseModel):
    coordinates: Coordinates = Field(
        ..., examples=[{"latitude": 55.755864, "longitude": 37.617698}]
    )
    text: str = Field(..., examples=["Хочу снять наличные с моей кредитной карты."])
    test: bool = Field(..., examples=[True])


class Special(BaseModel):
    Prime: bool
    juridical: bool
    person: bool
    ramp: bool
    vipOffice: bool
    vipZone: bool


class SpecialGpt(BaseModel):
    atm: bool = Field(..., alias="банкомат")
    online: bool = Field(..., alias="онлайн")
    vipZone: bool = Field(..., alias="привелегия")
    vipOffice: bool = False
    Prime: bool = Field(..., alias="прайм")
    ramp: bool = Field(..., alias="маломобильный")
    person: bool = Field(..., alias="физ лицо")
    juridical: bool = Field(..., alias="юр лицо")


class Response(BaseModel):
    _id: str
    atm: bool
    coordinates: Coordinates
    createdAt: str
    online: bool
    special: Special
    text: str
    userId: str


@app.post("/service")
async def get_service(data: Request):
    print(openai.api_key)
    print("start", data)
    if data.test:
        res = {
            "банкомат": True,
            "онлайн": False,
            "привелегия": False,
            "прайм": False,
            "маломобильный": False,
            "физ лицо": True,
            "юр лицо": False,
        }

        res = SpecialGpt(**res).model_dump()
        response = Response(
            _id="",
            atm=res["atm"],
            coordinates=Coordinates(
                latitude=0.0,
                longitude=0.0,
            ),
            createdAt=datetime.datetime.now().isoformat(),
            online=res["online"],
            special=Special(**res),
            userId="",
            text=data.text,
        )
        return response
    else:
        res = await predict(data.text)
        res = SpecialGpt(**res).model_dump()
        response = Response(
            _id="",
            atm=res["atm"],
            coordinates=Coordinates(
                latitude=0.0,
                longitude=0.0,
            ),
            createdAt=datetime.datetime.now().isoformat(),
            # 
            online=res["online"],
            special=Special(**res),
            userId="",
            text=data.text,
        )
        return response
