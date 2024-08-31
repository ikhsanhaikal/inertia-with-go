import { router } from "@inertiajs/react";
import { useEffect, useState } from "react";

type Todo = {
  id: number;
  userId: number;
  title: number;
  completed: number;
};
type State = {
  todos: Todo[];
  isLoading: boolean;
  page: number;
};

const Page2Index = (props: { inMemoryData: Todo[] }) => {
  console.log("props: ", props);
  const [state, setState] = useState<State>({
    todos: [],
    isLoading: false,
    page: 1,
  });
  const Get = async (newPage: number = 1) => {
    setState((prevState) => ({ ...prevState, isLoading: true, page: newPage }));
    const url = `http://localhost:4040/api/page2?page=${newPage}`;
    const response = await fetch(url);
    const todos = await response.json();
    setState((prevState) => ({
      ...prevState,
      isLoading: false,
      todos: todos,
    }));
  };
  useEffect(() => {
    Get();
  }, []);

  return (
    <div>
      <h1>list of all todos</h1>

      {state.isLoading ? (
        <p>loading...</p>
      ) : (
        <ul>
          {state.todos.map((todo) => {
            return <li>{todo.title}</li>;
          })}
        </ul>
      )}

      <>
        <p>inmemory todos</p>
        <ul>
          {props.inMemoryData.map((todo) => {
            return <li>{todo.title}</li>;
          })}
        </ul>
      </>

      <div>
        <button
          disabled={state.page == 1}
          onClick={async () => {
            Get(state.page - 1);
          }}
        >
          prev
        </button>
        <button
          disabled={state.page == 5} // for now
          onClick={async () => {
            Get(state.page + 1);
          }}
        >
          next
        </button>
      </div>
      <button onClick={() => router.get("/page2/new")}>add a todo</button>
    </div>
  );
};

export default Page2Index;
