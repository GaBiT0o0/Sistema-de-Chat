let socket;
let currentUsername;
let token;

// =========================
// REGISTER
// =========================
async function registerAndLogin() {

    const username = document.getElementById("register-username").value.trim();
    const password = document.getElementById("register-password").value.trim();

    if (!username) return alert("El nombre de usuario es obligatorio");
    if (!password) return alert("La contraseña es obligatoria");

    try {

        const res = await fetch("http://localhost:8000/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ username, password })
        });

        const data = await res.json();

        if (!res.ok) {
            alert(data.detail);
            return;
        }

        alert("Usuario registrado correctamente");

        document.getElementById("login-username").value = username;
        document.getElementById("login-password").value = password;

        await login();

    } catch (error) {

        console.error(error);
        alert("Error al conectar con el servidor");

    }
}

// =========================
// LOGIN
// =========================
async function login() {

    const username = document.getElementById("login-username").value.trim();
    const password = document.getElementById("login-password").value.trim();

    if (!username) return alert("Ingrese un usuario");
    if (!password) return alert("Ingrese una contraseña");

    try {

        const res = await fetch("http://localhost:8000/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ username, password })
        });

        const data = await res.json();

        if (!res.ok) {
            alert(data.detail);
            return;
        }

        token = data.token;
        currentUsername = data.username;

        localStorage.setItem("token", token);
        localStorage.setItem("username", currentUsername);

        document.getElementById("login-screen").style.display = "none";
        document.getElementById("chat-screen").style.display = "flex";

        connectChat(token);

    } catch (error) {

        console.error(error);
        alert("Error al iniciar sesión");

    }
}

// =========================
// WEBSOCKET
// =========================
function connectChat(token) {

    socket = new WebSocket(`ws://localhost:8080/ws?token=${token}`);

    socket.onopen = () => {
        console.log("Conectado como:", currentUsername);
    };

    socket.onmessage = (event) => {

        const data = JSON.parse(event.data);

        const div = document.createElement("div");

        div.className =
            data.user === currentUsername
                ? "message mine"
                : "message other";

        div.textContent = `${data.user}: ${data.text}`;

        const chat = document.getElementById("chat");

        chat.appendChild(div);
        chat.scrollTop = chat.scrollHeight;
    };

    socket.onerror = (err) => {
        console.error("WebSocket error:", err);
    };

    socket.onclose = () => {
        console.log("Desconectado");
    };
}

// =========================
// ENVIAR MENSAJE
// =========================
function sendMessage() {

    const input = document.getElementById("msg");
    const message = input.value.trim();

    if (
        !message ||
        !socket ||
        socket.readyState !== WebSocket.OPEN
    ) {
        return;
    }

    socket.send(message);
    input.value = "";
}

// =========================
// LOGOUT
// =========================
function logout() {

    localStorage.removeItem("token");
    localStorage.removeItem("username");

    if (socket) {
        socket.close();
    }

    location.reload();
}

// =========================
// AUTO LOGIN
// =========================
window.onload = async () => {

    token = localStorage.getItem("token");
    currentUsername = localStorage.getItem("username");

    if (!token) return;

    try {

        const res = await fetch(
            `http://localhost:8000/verify-token?token=${token}`
        );

        if (!res.ok) {

            localStorage.removeItem("token");
            localStorage.removeItem("username");

            return;
        }

        document.getElementById("login-screen").style.display = "none";
        document.getElementById("chat-screen").style.display = "flex";

        connectChat(token);

    } catch {

        localStorage.removeItem("token");
        localStorage.removeItem("username");

    }
};

// =========================
// ENTER
// =========================
document.addEventListener("DOMContentLoaded", () => {

    const input = document.getElementById("msg");

    if (!input) return;

    input.addEventListener("keydown", (e) => {
        if (e.key === "Enter") sendMessage();
    });

});