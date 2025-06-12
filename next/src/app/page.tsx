"use client";

import { useEffect, useState } from "react";

type Todo = {
  id: number;
  title: string;
  done: boolean;
};

// TODO: ECSの環境を作成次第変更
const apiUrl = "http://localhost:8080";

export default function TodoPage() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [newTitle, setNewTitle] = useState("");

  // タスクを取得する
  useEffect(() => {
    fetchTodos();
  }, []);

  // タスク一覧を取得
  const fetchTodos = async () => {
    const res = await fetch(apiUrl + "/get");
    const data = await res.json();
    setTodos(data);
  };

  // タスクの追加
  const handleAddTodo = async () => {
    await fetch(apiUrl + "/create", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        title: newTitle,
        done: false,
      }),
    });

    // 入力された値をリセット
    setNewTitle("");
    // 全タスクを再度取得
    fetchTodos();
  };

  // チェック/チェックを外す
  const handleToggleCompleted = async (todo: Todo) => {
    await fetch(apiUrl + `/update/${todo.id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        id: todo.id,
        title: todo.title,
        done: !todo.done,
      }),
      mode: "cors",
    });

    fetchTodos();
  };

  // タスクの削除
  const handleDeleteTodo = async (todo: Todo) => {
    console.log(todo);
    console.log(todo.id);
    await fetch(apiUrl + `/delete/${todo.id}`, {
      method: "DELETE",
    });

    fetchTodos();
  };

  return (
    <div className="p-6 max-w-2xl mx-auto">
      <h1 className="text-2xl font-semibold text-center text-while mb-6">
        📋 Todo List
      </h1>

      {/* 入力フォーム */}
      <div className="flex items-center gap-2 mb-6">
        <input
          type="text"
          value={newTitle}
          onChange={(e) => setNewTitle(e.target.value)}
          placeholder="新しいTodoを入力..."
          className="flex-1 p-2 rounded bg-gray-700 text-white placeholder-gray-400"
        />
        <button
          onClick={handleAddTodo}
          className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition"
        >
          追加
        </button>
      </div>

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
              <span
                className={`
                  text-lg font-medium
                  ${todo.done ? "text-green-400 line-through" : "text-white"}
                  `}
              >
                {todo.title}
              </span>
            </div>
            {/* 削除ボタン */}
            <button
              onClick={() => handleDeleteTodo(todo)}
              className={`px-3 py-1 rounded-full text-sm font-semibold text-white`}
            >
              削除
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
}
