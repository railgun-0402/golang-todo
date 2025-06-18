import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import { TodoList } from "./TodoList";
import { Todo } from "../types/todo";

jest.mock("./TodoItem", () => ({
  TodoItem: ({
    todo,
    onToggle,
    onDelete,
  }: {
    todo: Todo;
    onToggle: () => void;
    onDelete: () => void;
  }) => (
    <li data-testid={`todo-${todo.id}`}>
      <span>{todo.title}</span>
      <button onClick={onToggle}>Toggle</button>
      <button onClick={onDelete}>Delete</button>
    </li>
  ),
}));

/**
 * Test: TodoListクラスのテストクラス
 */
describe("TodoList", () => {
  const sampleTodos: Todo[] = [
    { id: 1, title: "買い物", done: false },
    { id: 2, title: "勉強", done: true },
  ];

  it("todosをすべて表示する", () => {
    render(
      <TodoList todos={sampleTodos} onToggle={jest.fn()} onDelete={jest.fn()} />
    );

    expect(screen.getByText("買い物")).toBeInTheDocument();
    expect(screen.getByText("勉強")).toBeInTheDocument();
  });

  it("TodoItemのonToggleとonDeleteが呼ばれる", () => {
    const onToggle = jest.fn();
    const onDelete = jest.fn();

    render(
      <TodoList todos={sampleTodos} onToggle={onToggle} onDelete={onDelete} />
    );

    // Toggle ボタンをクリック
    const toggleButtons = screen.getAllByText("Toggle");
    fireEvent.click(toggleButtons[0]);
    expect(onToggle).toHaveBeenCalledWith(sampleTodos[0]);

    // Delete ボタンをクリック
    const deleteButtons = screen.getAllByText("Delete");
    fireEvent.click(deleteButtons[1]);
    expect(onDelete).toHaveBeenCalledWith(sampleTodos[1]);
  });
});
