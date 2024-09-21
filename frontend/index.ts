const apiUrl = "/api/messages";

// Function to fetch messages from the server
async function fetchMessages() {
    try {
        const response = await fetch(apiUrl);
        const messages = await response.json();
        const messagesElement = document.getElementById("messages")!;
        messagesElement.textContent = ""; // Clear existing messaget h

        messages.forEach((message: string) => {
            messagesElement.textContent += "Received: " + message + "\n";
        });
    } catch (error) {
        console.error("Error fetching messages:", error);
    }
}

// Function to send a message to the server
async function sendMessage(event: Event) {
    event.preventDefault()
    const message = (document.getElementById("messageInput") as HTMLInputElement).value;
    try {
        await fetch(apiUrl, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ content: message })
        });
        (document.getElementById("messageInput") as HTMLInputElement).value = ''; // Clear the input field after sending
        fetchMessages(); // Refresh the messages list
    } catch (error) {
        console.error("Error sending message:", error);
    }
}

// Expose the sendMessage function to the global scope
(window as any).sendMessage = sendMessage;

// Fetch messages when the page loads
fetchMessages();
