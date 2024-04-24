import { cookies } from "next/headers"
import { redirect } from "next/navigation";

export default function Home() {
  if (!cookies().get("token")) {
    redirect("/login")
  }

  return (
    <main >
      
    </main>
  );
}
