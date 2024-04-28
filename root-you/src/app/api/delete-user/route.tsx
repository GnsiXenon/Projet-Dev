import { redirect } from "next/navigation"
import { NextRequest } from "next/server"

export async function POST(request: NextRequest) {
    let apiHostname = "localhost:3232"
    if (process.env.API_HOSTNAME) {
        apiHostname = process.env.API_HOSTNAME
    }
    const formData = await request.formData()
    const res = await fetch(`http://${apiHostname}/delete-user`, {
        method: "POST",
        body: JSON.stringify({
            mail: formData.get("mail")?.toString()
        })
    })

    redirect("/api/disconnect")
}