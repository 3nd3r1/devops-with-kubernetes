import { Todo } from "@/lib/definitions";

const TodoList = ({ todos }: { todos: Todo[] }) => (
  <div className="flex justify-center">
    <ul className="list-disc">
      {todos.map((todo) => (
        <li key={todo.id}>{todo.title}</li>
      ))}
    </ul>
  </div>
);

export default TodoList;
