"use server";

import { NewTodo, Todo } from "@/lib/definitions";
import { revalidateTag } from "next/cache";

const apiUrl = process.env.NEXT_PUBLIC_BACKEND_URL ?? "http://localhost:3001/api";

export const getTodos = async (): Promise<Todo[]> => {
  const res = await fetch(apiUrl + "/todos", {
    next: { tags: ["todos"] },
  });
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
  revalidateTag("todos");
  return createdTodo;
};
