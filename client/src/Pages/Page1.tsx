/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable @typescript-eslint/ban-ts-comment */

// interface Data {
//   id: number;
//   task: string;
//   completed: boolean;
// }

//@ts-ignore
const Page1 = ({ data }) => {
  console.log(`props: `, data);
  // console.log("data: ", data);
  // const data: Data = { id: 1, task: "", completed: false };
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
