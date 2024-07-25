import { Todo } from "@/lib/definitions";
import { getTodos } from "@/services/todo";

const TodoList = async () => {
  const todos: Todo[] = await getTodos();
  return (
    <div className="flex justify-center">
      <ul className="list-disc">
        {todos.map((todo) => (
          <li key={todo.id}>{todo.title}</li>
        ))}
      </ul>
    </div>
  );
};

export default TodoList;
