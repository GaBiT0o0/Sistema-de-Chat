import bcrypt
import os
from dotenv import load_dotenv

load_dotenv()

PEPPER = os.getenv("PEPPER")


def hash_password(password: str):
    password_peppered = password + PEPPER
    salt = bcrypt.gensalt()
    hashed = bcrypt.hashpw(password_peppered.encode(), salt)
    return hashed.decode()


def verify_password(password: str, hashed_password: str):
    password_peppered = password + PEPPER
    return bcrypt.checkpw(
        password_peppered.encode(),
        hashed_password.encode()
    )