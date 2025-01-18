import { MessagesList } from "@/components/messages";
import { Input } from "@/components/ui/input";

export default function Home() {
  return (
    <div className="m-auto flex max-h-[90vh] w-full max-w-[100rem] flex-1 flex-col gap-4 rounded-xl border border-zinc-300 p-4 pb-2 pt-2">
      {/* TODO: FIX UI: should not overflow vertically 
      https://ui.shadcn.com/docs/components/scroll-area
      */}
      <div className="h-full">
        <MessagesList />
      </div>
      <Input />
    </div>
  );
}
