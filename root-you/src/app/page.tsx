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
    <main className="flex flex-col justify-start gap-10 text-2xl">
      <nav className="w-full flex flex-row flex-wrap justify-between items-center py-7 px-6 border-b-white border-b-[2px] bg-gray-500/30 backdrop-blur-sm">
        <h1 className="relative text-7xl font-bold">RootYou</h1>
        <Link href="/scoreboard">Scoreboard</Link>
        <form action="/api/delete-user" method="POST">
          <input type="hidden" name="mail" value={decoded.mail} />
          <input className="text-red-700" type="submit" value="Delete account" />
        </form>
        <Link className="relative right-3 text-red-700 border-[2px] border-red-700 p-4 rounded-xl" href="/api/disconnect">Disconnect</Link>
      </nav>
      <div className="flex flex-col relative">
        <div className="w-full flex flex-col justify-center items-center">
          <h2 className="text-5xl font-bold">Welcome : {decoded.name}</h2>
          <p>You have {decoded.score} points</p>
        </div>
        <div className="flex flex-col justify-center items-center gap-[50px]">
          <form className="w-[90%] bg-green-500/30 flex flex-col justify-center items-start gap-3 p-4 rounded-lg backdrop-blur-sm" action="/api/submit" method="POST">
            <h3 className="text-3xl font-bold">Gauntlet</h3>
            <Link href={`http://${hostname}:5000`}>Start challenge</Link>
            <div className="w-[70vw] flex flex-col justify-center items-start gap-3">
              <input type="hidden" name="chall-id" value="1" />
              <input  type="hidden" name="user-id" value={decoded["id-user"]} />
              <input  type="hidden" name="mail" value={decoded["mail"]} />
              <input type="text" name="flag" className="text-black w-full rounded-lg p-2" />
              <input type="submit" value="Submit" className="p-2 border-2 rounded-lg" />
            </div>
          </form>
          <form className="w-[90%] bg-green-500/30 flex flex-col justify-center items-start gap-3 p-4 rounded-lg backdrop-blur-sm" action="/api/submit" method="POST">
            <h3 className="text-3xl font-bold">FTC</h3>
            <Link href="./ftc.zip" className="hover:underline">ftc.zip</Link>
            <Link href={`http://${hostname}:5001`}>Start challenge</Link>
            <div className="w-[70vw] flex flex-col justify-center items-start gap-3">
              <input type="hidden" name="chall-id" value="2" />
              <input type="hidden" name="user-id" value={decoded["id-user"]} />
              <input  type="hidden" name="mail" value={decoded["mail"]} />
              <input type="text" name="flag" className="text-black w-full rounded-lg p-2" />
              <input type="submit" value="Submit" className="p-2 border-2 rounded-lg" />
            </div>
          </form>
          <form className="w-[90%] bg-green-500/30 flex flex-col justify-center items-start gap-3 p-4 rounded-lg backdrop-blur-sm" action="/api/submit" method="POST">
            <h3 className="text-3xl font-bold">Apero</h3>
            <p>
              Suite à des activités suspectes dans la région, une enquête approfondie a été lancée pour identifier un espion présumé opérant sous le nom de code 'Marion Sinktaw'.
              Des rapports récents suggèrent que cette personne pourrait être liée à plusieurs incidents d'espionnage industriel et de cyberintrusions.
              Les autorités ont déployé des ressources spéciales pour traquer cette mystérieuse figure, tandis que les agences de renseignement internationales ont été alertées pour une coopération accrue dans cette affaire.
              Avec la sécurité nationale en jeu, chaque indice et témoignage sont cruciaux pour démêler le réseau d'intrigues dans lequel 'Marion Sinktaw' est impliqué.
              Nous avons besoin de ces deux informations pour mener à bien notre mission : Le nom de la piscine en lien avec son lieu de travail, et le numéro de rue ou elle habite. (Séparer les deux informations par un _)
            </p>
            <div className="w-[70vw] flex flex-col justify-center items-start gap-3">
              <input type="hidden" name="chall-id" value="3" />
              <input type="hidden" name="user-id" value={decoded["id-user"]} />
              <input  type="hidden" name="mail" value={decoded["mail"]} />
              <input type="text" name="flag" className="text-black w-full rounded-lg p-2" />
              <input type="submit" value="Submit" className="p-2 border-2 rounded-lg" />
            </div>
          </form>
          <form className="w-[90%] bg-green-500/30 flex flex-col justify-center items-start gap-3 p-4 rounded-lg backdrop-blur-sm" action="/api/submit" method="POST">
            <h3 className="text-3xl font-bold">Kyk's bot</h3>
            <Link href="https://discord.gg/VD7T4qjKF3">Start challenge</Link>
            <div className="w-[70vw] flex flex-col justify-center items-start gap-3">
              <input type="hidden" name="chall-id" value="4" />
              <input type="hidden" name="user-id" value={decoded["id-user"]} />
              <input  type="hidden" name="mail" value={decoded["mail"]} />
              <input type="text" name="flag" className="text-black w-full rounded-lg p-2" />
              <input type="submit" value="Submit" className="p-2 border-2 rounded-lg" />
            </div>
          </form>
        </div>
      </div>
    </main>
  );
}
