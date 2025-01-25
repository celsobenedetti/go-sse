import { useEffect, useRef } from "react";

export type MessageProps = {
  /**
   * Sender username
   */
  username: string;
  message: string;
  date: Date;
};

export function Message({ message, username, date }: MessageProps) {
  const elementRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (elementRef.current != null) {
      elementRef.current.scrollIntoView();
    }
  }, []);

  return (
    <div ref={elementRef} className="rounded-xl p-4 hover:bg-zinc-100">
      <div className="flex items-center justify-between gap-4">
        <p className="cursor-pointer font-bold hover:underline">{username}</p>
        <p className="text-sm font-thin text-zinc-500" suppressHydrationWarning>
          {date.toISOString()}
        </p>
      </div>
      <p>{message}</p>
    </div>
  );
}
