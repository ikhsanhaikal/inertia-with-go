import { useForm } from "@inertiajs/react";

const Page2New = () => {
  const { data, setData, post } = useForm({
    title: "",
    completed: false,
  });

  return (
    <div>
      <h1>Page 2 Form</h1>
      <form
        onSubmit={(e) => {
          e.preventDefault();
          post("/page2");
        }}
      >
        <table>
          <tbody>
            <tr>
              <th style={{ textAlign: "start" }}>
                <label
                  htmlFor="title"
                  style={{ paddingRight: 10, paddingBottom: 5 }}
                >
                  task
                </label>
              </th>
              <td>
                <input
                  id="title"
                  type="text"
                  value={data.title}
                  onChange={(e) => setData("title", e.target.value)}
                />
              </td>
            </tr>
            <tr>
              <th style={{ textAlign: "start" }}>
                <label
                  htmlFor="completed"
                  style={{ paddingRight: 10, paddingBottom: 5 }}
                >
                  completed
                </label>
              </th>
              <td>
                <input
                  id="completed"
                  type="checkbox"
                  checked={data.completed}
                  onChange={(e) => setData("completed", e.target.checked)}
                />
              </td>
            </tr>
          </tbody>
        </table>
        <br />
        <button type="submit">submit</button>
      </form>
    </div>
  );
};

export default Page2New;
