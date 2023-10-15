# {
#   "_id": {
#     "$oid": "652b291b73ff37b271628510"
#   },
#   "rating": 4.5,
#   "departmentId": "6529ad1f08712cd8d7df05ec",
#   "userId": "1546b9db-b95c-44e1-bd2f-efa67b81db8e",
#   "text": "Отличное обслуживание, но немного долго ждал в очереди.",
#   "createdAt": {
#     "$date": "2023-10-14T23:49:47.622Z"
#   }
# }

from typing import Annotated
from datetime import datetime
from pydantic import BaseModel, Field, BeforeValidator
from bson import ObjectId


class DepartmentRatingCreate(BaseModel):
    rating: float = Field(..., alias="rating")
    departmentId: str = Field(..., alias="departmentId")
    userId: str = Field(..., alias="userId")
    text: str = Field(..., alias="text")


class DepartmentRating(BaseModel):
    id: str = Field(..., alias="id")
    rating: float = Field(..., alias="rating")
    departmentId: str = Field(..., alias="departmentId")
    userId: str = Field(..., alias="userId")
    text: str = Field(..., alias="text")
    createdAt: datetime = Field(..., alias="createdAt")
