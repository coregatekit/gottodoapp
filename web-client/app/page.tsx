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
        className="border border-gray-400 p-2"
        placeholder="Add a new task"
      />
      <button className="bg-indigo-600 text-white p-2 ml-2">Add</button>
    </div>
  );
}

function TodoList() {
  const todos = ['Task 1', 'Task 2', 'Task 3'];

  return (
    <div className="mt-8 text-center">
      <h2 className="text-2xl font-bold">Todo List</h2>
      <ul className="mt-4">
        {todos.map((todo, index) => (
          <li key={index} className="text-lg">
            {todo}
          </li>
        ))}
      </ul>
    </div>
  );
}
