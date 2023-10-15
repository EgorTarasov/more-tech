import uuid
import random
import httpx
import asyncio
from app.models.rating import DepartmentRatingCreate


def generate_reviews(
    departmentId: str, num_reviews: int
) -> list[DepartmentRatingCreate]:
    reviews = []
    possible_ratings = [1.0, 1.5, 2.0, 2.5, 3.0, 3.5, 4.0, 4.5, 5.0]
    sample_texts = [
        "Отличное обслуживание, но немного долго ждал в очереди.",
        "Средний отдел, персонал вежлив, но процессы могли бы быть быстрее.",
        "Прекрасное место, всегда рад помочь клиентам.",
        "Неприятный опыт, долго не могли решить мою проблему.",
        "Хороший отдел, но парковка всегда переполнена.",
        "Средний сервис, но сотрудники стараются помочь.",
        "Отличное обслуживание, всегда быстро решают вопросы.",
        "Ужасное обслуживание, сотрудники не заботятся о клиентах.",
        "Лучший банк в городе, всегда рад работать с ними.",
        "Плохой опыт, долгое ожидание и невежливый персонал.",
        "Хороший отдел, но курсы обмена невыгодные.",
        "Средний сервис, но кассиры всегда в хорошем настроении.",
        "Отличное обслуживание, никаких нареканий.",
        "Ужасное место, не рекомендую никому.",
        "Превосходный сервис, всегда быстро и эффективно.",
        "Ужасное обслуживание, никто не хочет помочь клиентам.",
        "Средний отдел, средний опыт.",
        "Хороший банк, но надо улучшить организацию очереди.",
        "Ужасное обслуживание, долго ждал и ничего не решили.",
        "Отличное отделение, всегда помогают в любых вопросах.",
    ]

    for _ in range(num_reviews):
        rating = random.choice(possible_ratings)
        text = random.choice(sample_texts)
        # set userId as random uuid

        review = {
            "rating": rating,
            "text": text,
            "departmentId": departmentId,
            "userId": str(uuid.uuid4()),
        }

        reviews.append(DepartmentRatingCreate(**review))

    return reviews


async def upload_reviews(
    department_ids: list[str],
    server_url: str = "http://localhost:9999/v1/departments/rating",
):
    # разбиваем список на части по 20 элементов
    # параллельно генерируем и отправляем отзывы на сервер
    # с помощью asyncio.gather
    batches = [department_ids[i : i + 20] for i in range(0, len(department_ids), 20)]

    async def upload_batch(batch):
        async with httpx.AsyncClient() as client:
            for department_id in batch:
                for review in generate_reviews(department_id, random.randint(1, 20)):
                    await client.post(
                        server_url,
                        json=review.model_dump(),
                    )

    tasks = []
    for batch in batches:
        tasks.append(asyncio.create_task(upload_batch(batch)))

    await asyncio.gather(*tasks)


# Пример использования:
# if __name__ == "__main__":
# department_id = "6529ad1f08712cd8d7df05ec"
# num_reviews = 20
# reviews = generate_reviews(department_id, num_reviews)
# print(reviews)
