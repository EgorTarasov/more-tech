from typing import List
from typing import Optional

from fastapi import Request


class DepartmentForm:
    def __init__(self, request: Request) -> None:
        self.request: Request = request
        self.errors: List = []
        self.departmentId: Optional[str] = None

    async def load_data(self) -> None:
        form = await self.request.form()
        self.departmentId = form.get("departmentId")

    async def is_valid(self) -> bool:
        return True
