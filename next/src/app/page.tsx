"use client";

import { useTodos } from "./hooks/useTodos";
import { TodoInput } from "./components/TodoInput";
import { TodoList } from "./components/TodoList";

export default function TodoPage() {
  const {
    todos,
    newTitle,
    setNewTitle,
    handleAddTodo,
    handleDeleteTodo,
    handleToggleCompleted,
  } = useTodos();

  return (
    <div className="p-6 max-w-2xl mx-auto">
      <h1 className="text-2xl font-semibold text-center text-while mb-6">
        📋 Todo List
      </h1>

      {/* 入力フォーム */}
      <TodoInput
        newTitle={newTitle}
        setNewTitle={setNewTitle}
        onAdd={handleAddTodo}
      />

      {/* Todo一覧 */}
      <TodoList
        todos={todos}
        onToggle={handleToggleCompleted}
        onDelete={handleDeleteTodo}
      />
    </div>
  );
}
