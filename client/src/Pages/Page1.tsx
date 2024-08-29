interface Data {
  id: number;
  task: string;
  completed: boolean;
}

const Page1 = ({ data }: { data: Data }) => {
  console.log("data: ", data);
  return (
    <div>
      <h1>Page 1</h1>
      <div>
        <p>id: {data.id}</p>
        <p>task: {data.task}</p>
        <input
          type="checkbox"
          disabled
          checked={data.completed}
          title="completed"
        />
      </div>
      <a
        href="#"
        onClick={() => {
          window.history.back();
        }}
      >
        Back
      </a>
    </div>
  );
};

export default Page1;
