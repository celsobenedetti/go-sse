import { Message, MessageProps } from "@/components/messages/message";
import { Input } from "@/components/ui/input";

const mockMessages: MessageProps[] = [
  { message: "Hello message 1", username: "John Doe", date: new Date() },
  { message: "Hello message 2", username: "John Doe", date: new Date() },
  { message: "Hello message 3", username: "John Doe", date: new Date() },
  { message: "Hello message 4", username: "John Doe", date: new Date() },
  { message: "Hello message 5", username: "John Doe", date: new Date() },
  { message: "Hello message 6", username: "John Doe", date: new Date() },
  { message: "Hello message 7", username: "John Doe", date: new Date() },
  { message: "Hello message 8", username: "John Doe", date: new Date() },
  { message: "Hello message 9", username: "John Doe", date: new Date() },
];

export default function Home() {
  return (
    <div className="m-auto flex w-full max-w-[100rem] flex-1 flex-col gap-4 rounded-xl border border-zinc-300 p-4 pb-2 pt-2">
      <div className="h-full">
        {mockMessages.map(({ message, username, date }) => (
          <Message
            key={message}
            message={message}
            username={username}
            date={date}
          />
        ))}
      </div>
      <Input />
    </div>
  );
}
