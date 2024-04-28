import { cookies } from "next/headers"
import { redirect } from "next/navigation"
import Link from "next/link"
import { headers } from "next/headers"

var jwt = require('jsonwebtoken')

export default async function Home() {
  if (!cookies().get("martin_session_id")) {
    redirect("/login")
  }

  let decoded

  try {
    let token = cookies().get("martin_session_id")?.value
    decoded = jwt.verify(token, 'HaCoeur');
  } catch(err) {
    redirect("/api/disconnect")
  }

  const hostname = headers().get("host")?.split(":")[0]

  return (
    <main className="flex flex-col justify-start gap-3">
      <nav className="w-full flex flex-row justify-between items-center h-20 border-b-white border-b-[2px] bg-gray-500/30 backdrop-blur-sm">
        <h1 className="relative left-3 text-7xl font-bold">RootYou</h1>
        <form action="/api/delete-user" method="POST">
          <input type="hidden" name="mail" value={decoded.mail} />
          <input className="text-red-700" type="submit" value="Delete account" />
        </form>
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
        <form className="flex flex-col" action="/api/submit" method="POST">
          <label>Apero</label>
          <p>
            Suite à des activités suspectes dans la région, une enquête approfondie a été lancée pour identifier un espion présumé opérant sous le nom de code 'Marion Sinktaw'.
            Des rapports récents suggèrent que cette personne pourrait être liée à plusieurs incidents d'espionnage industriel et de cyberintrusions.
            Les autorités ont déployé des ressources spéciales pour traquer cette mystérieuse figure, tandis que les agences de renseignement internationales ont été alertées pour une coopération accrue dans cette affaire.
            Avec la sécurité nationale en jeu, chaque indice et témoignage sont cruciaux pour démêler le réseau d'intrigues dans lequel 'Marion Sinktaw' est impliqué.
            Nous avons besoin de ces deux informations pour mener à bien notre mission : Le nom de la piscine en lien avec son lieu de travail, et le numéro de rue ou elle habite. (Séparer les deux informations par un _)
          </p>
          <div>
            <input type="hidden" name="chall-id" value="3" />
            <input type="hidden" name="user-id" value={decoded["id-user"]} />
            <input  type="hidden" name="mail" value={decoded["mail"]} />
            <input type="text" name="flag" className="text-black" />
            <input type="submit" value="Submit" />
          </div>
        </form>
        <form className="flex flex-col" action="/api/submit" method="POST">
          <label>Kyk's bot</label>
          <Link href="https://discord.gg/VD7T4qjKF3">Start challenge</Link>
          <div>
            <input type="hidden" name="chall-id" value="4" />
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
