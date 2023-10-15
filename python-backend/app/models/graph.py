from pydantic import BaseModel, Field


class GraphView(BaseModel):
    id: str = Field(...)
    title: str = Field(...)
    data: list[int | float] = Field(...)
    labels: list[str] = Field(...)
