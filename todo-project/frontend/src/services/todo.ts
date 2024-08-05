import { NewTodo, Todo } from "@/lib/definitions";

const apiUrl = process.env.NEXT_PUBLIC_BACKEND_URL ?? "/api";

export const getTodos = async (): Promise<Todo[]> => {
  const res = await fetch(apiUrl + "/todos");
  const todos = await res.json();
  return todos;
};

export const createTodo = async (newTodo: NewTodo): Promise<Todo> => {
  const res = await fetch(apiUrl + "/todos", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(newTodo),
  });

  const createdTodo = await res.json();
  return createdTodo;
};
