from app.api.login import login_for_access_token
from fastapi import APIRouter, Depends, Request, HTTPException, status, Path
from fastapi.responses import RedirectResponse
from app.dependencies import get_db
from app.db import MongoDB
from app.models.graph import GraphView
import datetime
from app.forms.department import DepartmentForm
from fastapi.templating import Jinja2Templates
import random

templates = Jinja2Templates(directory="templates")

router = APIRouter(prefix="/info")


@router.get("/")
async def index(request: Request, db: MongoDB = Depends(get_db)):
    departments = db.get_departments()
    return templates.TemplateResponse(
        "panel/panel.html", {"request": request, "departments": departments}
    )


@router.post("/")
async def get_department(request: Request, db: MongoDB = Depends(get_db)):
    # get data from form
    form = DepartmentForm(request)
    await form.load_data()
    print(form.__dict__)
    department = db.get_department_by_id(form.departmentId)
    if department:
        return RedirectResponse(
            url=f"/info/{form.departmentId}", status_code=status.HTTP_302_FOUND
        )
    return templates.TemplateResponse("panel/panel.html", {"request": request})


@router.get("/{departmentId}")
async def department_detail(
    request: Request, departmentId: str = Path(...), db: MongoDB = Depends(get_db)
):
    department = db.get_department_by_id(departmentId)
    ratings = db.get_department_ratings(departmentId)

    months_ratings = {}

    for rating in ratings:
        month = rating.createdAt.month
        if month in months_ratings:
            months_ratings[month].append(rating)
        else:
            months_ratings[month] = [rating]

    for month in months_ratings:
        months_ratings[month] = sum(
            [rating.rating for rating in months_ratings[month]]
        ) / len(months_ratings[month])
    for month in range(1, 13):
        if month not in months_ratings:
            months_ratings[month] = 0
    month_names = [
        "Январь",
        "Февраль",
        "Март",
        "Апрель",
        "Май",
        "Июнь",
        "Июль",
        "Август",
        "Сентябрь",
        "Октябрь",
        "Ноябрь",
        "Декабрь",
    ]
    months_ratings = [months_ratings[month] for month in range(1, 13)]
    current = datetime.datetime.now().month
    months_ratings = months_ratings[current:] + months_ratings[:current]
    month_names = month_names[current:] + month_names[:current]

    tickets_times = [random.choice([5, 10, 15, 30, 50]) for _ in range(12)]
    waiting_times = [random.randint(1, 15) for _ in range(12)]
    line_graphs = [
        GraphView(
            id="department_ratings",
            title="Рейтинг отделения по месяцам",
            data=months_ratings,
            labels=month_names,
        ),
        GraphView(
            id="tickets_times",
            title="Время обработки заявок",
            data=tickets_times,
            labels=month_names,
        ),
        GraphView(
            id="waiting_times",
            title="Время в очереди ",
            data=waiting_times,
            labels=month_names,
        ),
    ]

    bar_graphs = [
        GraphView(
            id="problems",
            title="Количество зафиксированных неисправностей, усложняющих/делающих невозможным, оказание услуг",
            data=[random.randint(1, 10) for _ in range(12)],
            labels=month_names,
        )
    ]
    return templates.TemplateResponse(
        "panel/department.html",
        {
            "request": request,
            "department": department,
            "line_graphs": line_graphs,
            "bar_graphs": bar_graphs,
            "month_names": month_names,
            "graph_title": "Рейтинг отдела по месяцам",
            "ratings": months_ratings,
            "tickets_times": tickets_times,
        },
    )
