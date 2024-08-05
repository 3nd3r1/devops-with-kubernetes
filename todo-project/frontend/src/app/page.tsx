import TodoForm from "@/components/todo-form";
import TodoList from "@/components/todo-list";
import ImagenatorImage from "@/components/imagenator-image";

const Home = async () => {
  return (
    <main className="max-w-4xl mx-auto py-10">
      <div className="flex flex-col gap-4">
        <ImagenatorImage />
        <TodoForm />
        <TodoList />
      </div>
    </main>
  );
};

export const dynamic = "force-dynamic";
export default Home;
