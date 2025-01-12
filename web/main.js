const userId = new Date().getMilliseconds();
const messages = new EventSource(`/rooms/123/subscribe/${userId}`);
const messagesList = document.querySelector("#messages-list");

function sendMessage() {
  fetch("/messages", {
    method: "POST",
    body: JSON.stringify({
      roomId: "123",
      senderId: "javascript_client",
      message: "hello from javascript",
    }),
  })
    .then(console.info)
    .catch(console.error);
}

function addMessage(data) {
  const newMsg = document.createElement("li");
  newMsg.textContent = data;
  messagesList.appendChild(newMsg);
}

messages.addEventListener("message", (args) => {
  console.log("Handling message", { args });
  addMessage(args.data);
});

messages.addEventListener("close", (args) => {
  console.log("Handling close", { args });
  messages.close();
});

messages.addEventListener("error", (args) => {
  console.log("Handling error", { args });
});
