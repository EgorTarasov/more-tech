import pymongo
import json


connection_string = ""  # "mongodb://mongouser:mongopass@localhost:27017/?authSource=admin&readPreference=primary&appname=mongodb-vscode+1.3.1&ssl=false&directConnection=true&serverSelectionTimeoutMS=2000"

client = pymongo.MongoClient(connection_string)

# Access a database.
db = client.dev

# load data from json
with open("data/data.json") as f:
    branches = json.load(f)["branches"]

    for doc in branches:

        loc = {
            "type": "Point",
            "coordinates": [
                doc["coordinates"]["longitude"],
                doc["coordinates"]["latitude"],
            ],
        }
        doc["location"] = loc
    db.departments.insert_many(branches)


# Банкомат №399829 Ст. м. «Боровицкая (Серпуховско-Тимирязевская)» м. Боровицкая г. Москва, Вестибюль ст.  Выдает: ₽ Принимает: ₽  Режим работы: ПН: 05:45-23:59 ВТ: 05:45-23:59 СР: 05:45-23:59 ЧТ: 05:45-23:59 ПТ: 05:45-23:59 СБ: 05:45-23:59 ВС: 05:45-23:59  Возможности банкомата: NFC (бесконтактное обслуживание) Платежи по QR‑коду Снятие наличных по QR‑коду Оборудован для слабовидящих  Актуальная ссылка на данный банкомат:
# https://online.vtb.ru/i/ATM
