from fastapi import FastAPI
from sqlmodel import SQLModel, create_engine, Session, select

from models.user import User
from routes.auth import hash_password, verify_password

app = FastAPI()

engine = create_engine("sqlite:///database.db")

SQLModel.metadata.create_all(engine)

@app.get("/")
def home():

    return {"message": "Auth Service Running"}

@app.post("/register")
def register(username: str, password: str):

    hashed_password = hash_password(password)

    user = User(
        username=username,
        password=hashed_password
    )

    with Session(engine) as session:

        session.add(user)
        session.commit()

    return {"message": "Usuario registrado"}

@app.post("/login")
def login(username: str, password: str):

    with Session(engine) as session:

        statement = select(User).where(User.username == username)

        user = session.exec(statement).first()

        if not user:
            return {"error": "Usuario no encontrado"}

        if verify_password(password, user.password):

            return {"message": "Login correcto"}

        return {"error": "Contraseña incorrecta"}