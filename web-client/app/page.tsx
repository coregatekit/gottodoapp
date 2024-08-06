export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div className="w-full items-center">
        <div className="text-center text-4xl font-bold italic text-indigo-700 drop-shadow-xl">
          Todo List
        </div>
        <div className="flex flex-col justify-between">
          <InputTodo />
          <TodoList />
        </div>
      </div>
    </main>
  );
}

function InputTodo() {
  return (
    <div className="mt-8 text-center">
      <input
        type="text"
        className="border rounded border-gray-400 p-2"
        placeholder="Add a new task"
      />
      <button className="bg-indigo-600 rounded text-white p-2 ml-2">Add</button>
    </div>
  );
}

function TodoList() {
  const todos = ['Task 1', 'Task 2', 'Task 3'];

  const handleDeleteTodo = (index: number) => {
    alert(`Delete todo at index ${index}`);
  };
  return (
    <div className="mt-8 text-center">
      <ul className="mt-4 space-y-2">
        {todos.map((todo, index) => (
          <li key={index} className="text-lg flex items-center">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-5 w-5 text-indigo-600 mr-2"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fillRule="evenodd"
                d="M3 5a1 1 0 011-1h12a1 1 0 011 1v10a1 1 0 01-1 1H4a1 1 0 01-1-1V5zm1-2a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V5a2 2 0 00-2-2H4z"
                clipRule="evenodd"
              />
            </svg>
            {todo}
            <button
              className="text-red-600 ml-2"
              // onClick={() => handleDeleteTodo(index)}
            >
              Delete
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
}
