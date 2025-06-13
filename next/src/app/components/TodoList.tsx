import { Todo } from "../types/todo";
import { TodoItem } from "./TodoItem";

type Props = {
  todos: Todo[];
  onToggle: (todo: Todo) => void;
  onDelete: (todo: Todo) => void;
};

export const TodoList = ({ todos, onToggle, onDelete }: Props) => (
  <ul className="space-y-4">
    {todos.map((todo) => (
      <TodoItem
        key={todo.id}
        todo={todo}
        onToggle={() => onToggle(todo)}
        onDelete={() => onDelete(todo)}
      />
    ))}
  </ul>
);
