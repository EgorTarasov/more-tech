from app.db import MongoDB


def get_db() -> MongoDB:
    return MongoDB()
