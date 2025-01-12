const messages = new EventSource("/room/1");

messages.addEventListener("message", (args) => {
  console.log("Handling message", { args });
});

messages.addEventListener("close", (args) => {
  console.log("Handling close", { args });
  messages.close();
});

messages.onerror((event) => {
  console.log("Handling close", { args: event });
});
