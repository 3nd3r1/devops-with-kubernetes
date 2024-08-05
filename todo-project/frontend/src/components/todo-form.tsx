"use client";

import { useState } from "react";

import { createTodo } from "@/services/todo";
import { mutate } from "swr";

const TodoForm = () => {
  const [title, setTitle] = useState<string>("");

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    createTodo({ title }).then(() => mutate("todos"));
    setTitle("");
  };

  return (
    <div className="flex justify-center">
      <form className="flex flex-row gap-4" onSubmit={handleSubmit}>
        <input
          name="title"
          type="text"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          maxLength={140}
          required
          placeholder="Todo Title"
          className="bg-transparent border border-white rounded-lg p-1"
        />
        <button type="submit" className="border border-white rounded-lg p-1">
          Create Todo
        </button>
      </form>
    </div>
  );
};

export default TodoForm;
