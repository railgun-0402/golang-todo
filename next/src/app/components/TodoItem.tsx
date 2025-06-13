import { Todo } from "../types/todo";

type Props = {
  todo: Todo;
  onToggle: () => void;
  onDelete: () => void;
};

export const TodoItem = ({ todo, onToggle, onDelete }: Props) => {
  return (
    <li className="flex justify-between items-center p-4 bg-gray-800 border border-gray-700 rounded-xl shadow-sm hover:shadow-md transition">
      <div className="flex items-center">
        {/* チェックボックス */}
        <input
          type="checkbox"
          checked={todo.done}
          onChange={onToggle}
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
        onClick={onDelete}
        className={`px-3 py-1 rounded-full text-sm font-semibold text-white`}
      >
        削除
      </button>
    </li>
  );
};
