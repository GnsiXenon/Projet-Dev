import { cookies } from "next/headers"
import { redirect } from "next/navigation"
import Link from "next/link"

var jwt = require('jsonwebtoken')

export default async function Home() {
  if (!cookies().get("token")) {
    redirect("/login")
  }

  let decoded

  try {
    let token = cookies().get("token")?.value
    decoded = jwt.verify(token, 'HaCoeur');
  } catch(err) {
    redirect("/api/disconnect")
  }

  return (
    <main>
      <h1 className="text-7xl font-bold">RootYou</h1>
      <h2>Welcome {decoded.name}</h2>
      <p>You have {decoded.score} points</p>
      <Link href="/api/disconnect">Disconnect</Link>
      <div className="flex flex-col">
        <form className="flex flex-col" action="/api/submit" method="POST">
          <label>Gauntlet</label>
          <Link href="http://0.0.0.0:5000">Start challenge</Link>
          <div>
            <input type="text" id="1" />
            <input type="submit" value="Submit" />
          </div>
        </form>
        <form className="flex flex-col" action="/api/submit" method="POST">
          <label>FTC</label>
          <Link href="http://0.0.0.0:5001">Start challenge</Link>
          <Link href="./ftc.zip">ftc.zip</Link>
          <div>
            <input type="text" id="2" />
            <input type="submit" value="Submit" />
          </div>
        </form>
      </div>
    </main>
  );
}
