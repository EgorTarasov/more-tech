from __future__ import annotations

from typing import List

from pydantic import BaseModel, Field


class Special(BaseModel):
    vipzone: int
    vipoffice: int
    ramp: int
    person: int
    juridical: int
    prime: int


class Coordinates(BaseModel):
    latitude: float
    longitude: float


class Coordinates1(BaseModel):
    latitude: float
    longitude: float


class Location(BaseModel):
    type: str
    coordinates: Coordinates1


class Loadhour(BaseModel):
    hour: str
    load: float


class WorkloadItem(BaseModel):
    day: str
    loadhours: List[Loadhour]


class Department(BaseModel):
    _id: str
    id: int
    biskvitid: str
    shortname: str
    address: str
    city: str
    schedulefl: str
    schedulejurl: str
    special: Special
    coordinates: Coordinates
    location: Location
    workload: List[WorkloadItem]
