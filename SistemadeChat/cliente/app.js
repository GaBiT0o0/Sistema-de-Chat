const socket = new WebSocket("ws://localhost:8080/ws");

socket.onmessage = function(event) {

    const chat = document.getElementById("chat");

    chat.innerHTML += `<p>${event.data}</p>`;
};

function sendMessage() {

    const input = document.getElementById("message");

    socket.send(input.value);

    input.value = "";
}
