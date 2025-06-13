type Props = {
  newTitle: string;
  setNewTitle: (title: string) => void;
  onAdd: () => void;
};

export const TodoInput = ({ newTitle, setNewTitle, onAdd }: Props) => {
  return (
    <div className="flex items-center gap-2 mb-6">
      <input
        type="text"
        value={newTitle}
        onChange={(e) => setNewTitle(e.target.value)}
        placeholder="新しいTodoを入力..."
        className="flex-1 p-2 rounded bg-gray-700 text-white placeholder-gray-400"
      />
      <button
        onClick={onAdd}
        className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition"
      >
        追加
      </button>
    </div>
  );
};
