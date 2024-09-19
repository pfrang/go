var socket = new WebSocket("ws://localhost:8080/ws");
socket.onopen = function () {
    console.log("Connected to the WebSocket server");
};
socket.onmessage = function (event) {
    var message = event.data;
    document.getElementById("messages").textContent += "Received: " + message + "\n";
};
function sendMessage() {
    var message = document.getElementById("messageInput").value;
    socket.send(message);
    document.getElementById("messageInput").value = ''; // Clear the input field after sending
}
window.sendMessage = sendMessage;
socket.onclose = function () {
    console.log("Disconnected from the WebSocket server");
};
