import { useEffect, useState } from "react";
import { Todo } from "../types/todo";

// 本番用はECSの方で読み取る
const apiUrl = process.env.NEXT_PUBLIC_API_BASE_URL;

/**
 * 状態管理のロジックをまとめる
 * タスクのCRUD処理を行うクラス
 */
export function useTodos() {
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

  return {
    todos,
    newTitle,
    setNewTitle,
    handleAddTodo,
    handleDeleteTodo,
    handleToggleCompleted,
  };
}
