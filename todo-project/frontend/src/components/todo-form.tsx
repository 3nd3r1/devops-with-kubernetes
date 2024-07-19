const TodoForm = () => (
  <div className="flex justify-center">
  <form className="flex flex-row gap-4">
    <input name="title" type="text" maxLength={140} required placeholder="Todo Title" className="bg-transparent border border-white rounded-lg p-1" />
    <button type="submit" className="border border-white rounded-lg p-1">Create Todo</button>
  </form>
  </div>
);

export default TodoForm;
