// src/app/todos/components/__tests__/TodoItem.test.tsx
import { fireEvent, render, screen } from "@testing-library/react";
import { TodoItem } from "./TodoItem";
import { Todo } from "../types/todo";
import "@testing-library/jest-dom";

describe("TodoItem", () => {
  const baseTodo: Todo = {
    id: 1,
    title: "テストタスク",
    done: false,
  };

  it("タイトルの表示", () => {
    render(
      <TodoItem todo={baseTodo} onToggle={jest.fn()} onDelete={jest.fn()} />
    );
    expect(screen.getByText("テストタスク")).toBeInTheDocument();
  });

  it("Checkboxの状態に応じて done が変わる", () => {
    const doneTodo = { ...baseTodo, done: true };
    render(
      <TodoItem todo={doneTodo} onToggle={jest.fn()} onDelete={jest.fn()} />
    );
    expect(screen.getByRole("checkbox")).toBeChecked();
  });

  it("Checkboxをクリックすると、 onToggle が呼ばれること", () => {
    const onToggle = jest.fn();
    render(
      <TodoItem todo={baseTodo} onToggle={onToggle} onDelete={jest.fn()} />
    );
    fireEvent.click(screen.getByRole("checkbox"));
    expect(onToggle).toHaveBeenCalled();
  });

  it("削除ボタンをクリックすると、 onDelete が呼ばれること", () => {
    const onDelete = jest.fn();
    render(
      <TodoItem todo={baseTodo} onToggle={jest.fn()} onDelete={onDelete} />
    );
    // fireEvent.click(screen.getByText("削除"));
    fireEvent.click(screen.getByRole("button", { name: "削除" }));
    expect(onDelete).toHaveBeenCalled();
  });
});
