import pymongo
from app.config import CONFIG
from app.models.department import Department
from app.models.rating import DepartmentRating
from typing import Type
from bson import ObjectId


class MongoDB:
    # make this class a singleton
    instance = None

    def __init__(self):
        self.client = pymongo.MongoClient(CONFIG.mongo_uri)
        self.db = self.client[CONFIG.mongo_db]
        self.admin_db = self.client["admin"]
        self.users = self.db["users"]
        self.dev = self.client["dev"]

    def __new__(cls):
        if cls.instance is None:
            cls.instance = super().__new__(cls)
        return cls.instance

    def get_user(self, username: str):
        return self.users.find_one({"email": username})

    def get_department_by_id(self, departmentId: str):
        try:
            department = self.dev["departments"].find_one(
                {"_id": ObjectId(departmentId)}
            )
            if department:
                return Department(**department)

            return None
        except Exception as e:
            return None

    def get_departments(self, _ids: list[str] = list()):
        if _ids:
            departments = self.dev["departments"].find(
                {"_id": {"$in": [ObjectId(_id) for _id in _ids]}}
            )
        else:
            departments = self.dev["departments"].find()
        return [Department(**department) for department in departments]

    def get_departments_ids(self) -> list[str]:
        return [str(department["_id"]) for department in self.dev["departments"].find()]

    def get_department_ratings(self, departmentId: str):
        # 6529ad1f08712cd8d7df05ec
        data = []
        for obj in self.dev["rating"].find({"departmentId": departmentId}):
            print(type(obj["_id"]))
            id = str(obj["_id"])
            data.append(DepartmentRating(id=id, **obj))
        return data
