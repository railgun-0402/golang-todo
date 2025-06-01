"use client";

import { useEffect, useState } from "react";

type Todo = {
  id: number;
  title: string;
  done: boolean;
};

export default function TodoPage() {
  const [todos, setTodos] = useState<Todo[]>([]);

  const apiUrl = "http://localhost:8080";

  // タスクを取得する
  useEffect(() => {
    fetchTodos();
  }, []);

  const fetchTodos = async () => {
    const res = await fetch(apiUrl + "/get");
    const data = await res.json();
    setTodos(data);
  };

  const handleToggleCompleted = async (todo: Todo) => {
    await fetch(apiUrl + `/update/${todo.id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        title: todo.title,
        completed: !todo.done,
      }),
      mode: "cors",
    });

    fetchTodos();
  };

  return (
    <div className="p-6 max-w-2xl mx-auto">
      <h1 className="text-2xl font-semibold text-center text-while mb-6">
        📋 Todo List
      </h1>

      {/* Todo一覧 */}
      <ul className="space-y-4">
        {todos.map((todo) => (
          <li
            key={todo.id}
            className="flex justify-between items-center p-4 bg-gray-800 border border-gray-700 rounded-xl shadow-sm hover:shadow-md transition"
          >
            <div className="flex items-center">
              {/* チェックボックス */}
              <input
                type="checkbox"
                checked={todo.done}
                onChange={() => handleToggleCompleted(todo)}
                className="w-5 h-5 mr-4 accent-green-600"
              />
              <span className="text-lg font-medium text-white">
                {todo.title}
              </span>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
}
