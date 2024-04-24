import { cookies } from "next/headers"
import { redirect } from "next/navigation"
import Link from "next/link"

export default function Home() {
  if (!cookies().get("token")) {
    redirect("/login")
  }

  return (
    <main >
      <h1 className="text-7xl font-bold">RootYou</h1>
      <form className="flex flex-col" action="" method="POST">
        <div className="flex flex-col">
          <label>Gauntlet</label>
          <Link href="http://0.0.0.0:5000">Start challenge</Link>
          <div>
            <input type="text" id="1" />
            <input type="submit" value="Submit" />
          </div>
        </div>
        <div className="flex flex-col">
          <label>FTC</label>
          <Link href="http://0.0.0.0:5001">Start challenge</Link>
          <Link href="./ftc.zip">ftc.zip</Link>
          <div>
            <input type="text" id="2" />
            <input type="submit" value="Submit" />
          </div>
        </div>
      </form>
    </main>
  );
}
