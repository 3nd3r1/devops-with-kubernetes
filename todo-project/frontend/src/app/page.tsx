import TodoForm from "@/components/todo-form";
import TodoList from "@/components/todo-list";

import { Todo } from "@/lib/definitions";

const Home = () => {
  const todos: Todo[] = [
    {
      id: 1,
      title: "Todo 1",
      completed: false,
    },
    {
      id: 2,
      title: "Todo 2",
      completed: true,
    },
  ];

  return (
    <main className="max-w-4xl mx-auto py-10">
      <div className="flex flex-col gap-4">
        <img src="http://localhost:8081/imagenator/image" alt="Imagenator" />
        <TodoForm />
        <TodoList todos={todos} />
      </div>
    </main>
  );
};

export default Home;
