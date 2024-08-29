import { Link } from "@inertiajs/react";
import { useState } from "react";

function Index() {
  const [count, setCount] = useState(0);
  console.log("Log(App Component): was called");
  return (
    <div>
      <h1>count {count}</h1>
      <button onClick={() => setCount(count + 1)} style={{ padding: 2 }}>
        increment
      </button>
      <button
        onClick={() => setCount(count == 0 ? count : count - 1)}
        style={{ padding: 2 }}
      >
        decrement
      </button>
      <Link href="/page1">Page 1</Link>
      <Link href="/page2">Page 2</Link>
    </div>
  );
}

export default Index;
