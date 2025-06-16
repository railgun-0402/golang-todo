import { renderHook, waitFor } from "@testing-library/react";
import { useTodos } from "./useTodos";
import { Todo } from "../types/todo";

const mockTodos: Todo[] = [
  { id: 1, title: "Test Todo 1", done: false },
  { id: 2, title: "Test Todo 2", done: true },
];

describe("useTodos hook", () => {
  beforeEach(() => {
    global.fetch = jest.fn();
  });

  afterEach(() => {
    jest.resetAllMocks();
  });

  it("fetches todos on mount", async () => {
    // fetch のモックレスポンス
    (fetch as jest.Mock).mockResolvedValueOnce({
      json: async () => mockTodos,
    });

    const { result } = renderHook(() => useTodos());

    // todos が更新されるのを待つ
    await waitFor(() => {
      expect(result.current.todos).toEqual(mockTodos);
    });

    expect(fetch).toHaveBeenCalledWith("http://localhost:8080/get");
  });
});
