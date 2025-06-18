import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import { TodoInput } from "./TodoInput";

/**
 * Test: TodoInputクラスのテストクラス
 */
describe("TodoInput", () => {
  it("入力欄とボタンが表示される", () => {
    render(<TodoInput newTitle="" setNewTitle={jest.fn()} onAdd={jest.fn()} />);

    expect(
      screen.getByPlaceholderText("新しいTodoを入力...")
    ).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "追加" })).toBeInTheDocument();
  });

  it("テキストを入力すると setNewTitle が呼ばれる", () => {
    const setNewTitleMock = jest.fn();
    render(
      <TodoInput newTitle="" setNewTitle={setNewTitleMock} onAdd={jest.fn()} />
    );

    const input = screen.getByPlaceholderText("新しいTodoを入力...");
    fireEvent.change(input, { target: { value: "買い物" } });

    expect(setNewTitleMock).toHaveBeenCalledWith("買い物");
  });

  it("追加ボタンをクリックすると onAdd が呼ばれる", () => {
    const onAddMock = jest.fn();
    render(
      <TodoInput newTitle="掃除" setNewTitle={jest.fn()} onAdd={onAddMock} />
    );

    const button = screen.getByRole("button", { name: "追加" });
    fireEvent.click(button);

    expect(onAddMock).toHaveBeenCalled();
  });
});
