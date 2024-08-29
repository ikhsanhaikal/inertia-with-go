import { router } from "@inertiajs/react";
import { useEffect, useState } from "react";

export default function Notfound() {
  const [count, setCount] = useState(5);
  const [intervalId, setIntervalId] = useState<number | null>(null);
  console.log("count: ", count);
  useEffect(() => {
    setIntervalId(
      setInterval(() => {
        setCount((prev: number) => prev - 1);
        console.log("oucch");
      }, 1000)
    );
  }, []);

  useEffect(() => {
    if (count == 0) {
      if (intervalId !== null) {
        clearInterval(intervalId);
      }
      router.get("/");
    }
  }, [count, intervalId]);

  return (
    <div>
      <h1>Page does not exist</h1>
      <h2>You will be redirect to home page in {count} </h2>
    </div>
  );
}
