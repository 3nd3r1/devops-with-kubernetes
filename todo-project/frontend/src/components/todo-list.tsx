"use client";

import { getTodos } from "@/services/todo";
import useSWR from "swr";

const TodoList = () => {
  const { data, error, isLoading } = useSWR("todos", getTodos, {
    refreshInterval: 1000,
  });

  if (isLoading) {
    return <div>loading...</div>;
  }

  if (error || typeof data === "undefined") {
    return <div>failed to load</div>;
  }

  return (
    <div className="flex justify-center">
      <ul className="list-disc">
        {data.map((data) => (
          <li key={data.id}>{data.title}</li>
        ))}
      </ul>
    </div>
  );
};

export default TodoList;
