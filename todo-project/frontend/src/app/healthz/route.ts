export async function GET() {
  try {
    // This works because backend is in same pod
    const res = await fetch("http://localhost:8080/healthz");
    if (!res.ok) {
      return new Response("not ready", { status: 500 });
    }
  } catch (error) {
    return new Response("not ready", { status: 500 });
  }

  return new Response("ok", { status: 200 });
}
