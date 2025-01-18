"use client";
import { useEffect, useState } from "react";
import { Message, MessageProps } from "./message";

const mockMessage = (m: string): MessageProps => ({
  message: m,
  username: "John Doe",
  date: new Date(),
});

const mockMessages: MessageProps[] = [
  mockMessage("Hello message 1"),
  mockMessage("Hello message 2"),
  mockMessage("Hello message 3"),
];

export function MessagesList() {
  const [messages, setMessages] = useState(mockMessages);

  const addMessage = (msg: MessageProps) =>
    setMessages((prev) => [...prev, msg]);

  useEffect(() => {
    const subscribe = new EventSource(
      "http://localhost:3000/rooms/123/subscribe/123",
    );

    subscribe.addEventListener("message", (ev) => {
      const data = JSON.parse(JSON.parse(ev.data)); //HACK: does this need to be fixed on the server?
      console.log("new message", { data, t: typeof data });
      addMessage(mockMessage(data.message));
    });
  }, []);

  return (
    <div className="h-[inherit] overflow-y-scroll">
      {messages.map(({ message, username, date }, i) => (
        <Message
          key={message + i}
          message={message}
          username={username}
          date={date}
        />
      ))}
    </div>
  );
}
