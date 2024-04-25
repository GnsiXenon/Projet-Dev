import { cookies } from "next/headers"
import { redirect } from "next/navigation"
import Link from "next/link"
import { headers } from "next/headers"

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

  const hostname = headers().get("host")?.split(":")[0]

  return (
    <main className="flex flex-col justify-start gap-3">
      <nav className="w-full flex flex-row justify-between items-center h-20 border-b-white border-b-[2px] bg-gray-500/30 backdrop-blur-sm">
        <h1 className="relative left-3 text-7xl font-bold">RootYou</h1>
        <Link className="relative right-3 text-red-700 border-[2px] border-red-700 p-4 rounded-xl" href="/api/disconnect">Disconnect</Link>
      </nav>
      <div className="flex flex-col">
        <h2>Welcome {decoded.name}</h2>
        <p>You have {decoded.score} points</p>
        <form className="flex flex-col" action="/api/submit" method="POST">
          <label>Gauntlet</label>
          <Link href={`http://${hostname}:5000`}>Start challenge</Link>
          <div>
            <input type="hidden" name="chall-id" value="1" />
            <input  type="hidden" name="user-id" value={decoded["id-user"]} />
            <input  type="hidden" name="mail" value={decoded["mail"]} />
            <input type="text" name="flag" className="text-black" />
            <input type="submit" value="Submit" />
          </div>
        </form>
        <form className="flex flex-col" action="/api/submit" method="POST">
          <label>FTC</label>
          <Link href={`http://${hostname}:5001`}>Start challenge</Link>
          <Link href="./ftc.zip">ftc.zip</Link>
          <div>
            <input type="hidden" name="chall-id" value="2" />
            <input type="hidden" name="user-id" value={decoded["id-user"]} />
            <input  type="hidden" name="mail" value={decoded["mail"]} />
            <input type="text" name="flag" className="text-black" />
            <input type="submit" value="Submit" />
          </div>
        </form>
      </div>
    </main>
  );
}
