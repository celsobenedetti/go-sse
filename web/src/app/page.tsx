"use client";

import { useState } from "react";

import { MessagesList } from "@/components/messages";
import { Input } from "@/components/ui/input";
import { useSessionStore } from "@/store/session-store";

export default function Home() {
  const [input, setInput] = useState("");
  const { userId, username } = useSessionStore();

  const handleSubmit = () => {
    fetch("http://localhost:3000/messages", {
      method: "POST",
      body: JSON.stringify({
        message: input,
        sender: userId,
        roomId: "123",
      }),
    })
      .then((res) => {
        console.info("successfully posted message", { res });
        setInput("");
      })
      .catch(console.error);
  };

  return (
    <div className="m-auto flex max-h-[90vh] w-full max-w-[100rem] flex-1 flex-col gap-4 rounded-xl border border-zinc-300 p-4 pb-2 pt-2">
      <MessagesList />
      <Input
        type="text"
        placeholder="Chat in this room"
        value={input}
        onChange={(e) => setInput(e.target.value)}
        onKeyUp={(e) => e.code === "Enter" && handleSubmit()}
        onSubmit={() => console.log({ input })}
      />
    </div>
  );
}
