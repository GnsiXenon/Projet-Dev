export default async function ScoreBoard() {
    let apiHostname = "localhost:3232"
    if (process.env.API_HOSTNAME) {
        apiHostname = process.env.API_HOSTNAME
    }
    const resp = await fetch(`http://${apiHostname}/users`, {
        method: "GET",
        cache: "no-cache"
    })
    var data
    if (resp.ok) {
        data = await resp.json()
    }

    return (
        <main className="flex flex-col justify-center items-center gap-4 relative top-4">
            {data.map((user: {"id-user": string, "name": string, "score": 0}) => (<div className="flex flex-row w-[90vw] justify-between p-6 bg-blue-600/50 rounded-lg backdrop-blur-sm" key={user["id-user"]}>
                <p className="text-3xl">
                    {user["name"]}
                </p>
                <p className="text-3xl">
                    {user["score"]}
                </p>
            </div>))}
        </main>
    )
}