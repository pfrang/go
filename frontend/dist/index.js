"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
const apiUrl = "/api/messages";
// Function to fetch messages from the server
function fetchMessages() {
    return __awaiter(this, void 0, void 0, function* () {
        try {
            const response = yield fetch(apiUrl);
            const messages = yield response.json();
            const messagesElement = document.getElementById("messages");
            messagesElement.textContent = ""; // Clear existing messaget h
            messages.forEach((message) => {
                messagesElement.textContent += "Received: " + message + "\n";
            });
        }
        catch (error) {
            console.error("Error fetching messages:", error);
        }
    });
}
// Function to send a message to the server
function sendMessage(event) {
    return __awaiter(this, void 0, void 0, function* () {
        event.preventDefault();
        const message = document.getElementById("messageInput").value;
        try {
            yield fetch(apiUrl, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ content: message })
            });
            document.getElementById("messageInput").value = ''; // Clear the input field after sending
            fetchMessages(); // Refresh the messages list
        }
        catch (error) {
            console.error("Error sending message:", error);
        }
    });
}
// Expose the sendMessage function to the global scope
window.sendMessage = sendMessage;
// Fetch messages when the page loads
fetchMessages();
