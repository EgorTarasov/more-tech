from fastapi import APIRouter, BackgroundTasks, Depends
from app.services import rating
from app.db import MongoDB
from app.dependencies import get_db

router = APIRouter(prefix="/prototype")


@router.post("/upload")
async def upload_reviews(
    url: str = "http://localhost:9999/v1/departments/rating",
    db: MongoDB = Depends(get_db),
):
    department_ids = db.get_departments_ids()

    await rating.upload_reviews(department_ids, url)

    return {"message": department_ids}
