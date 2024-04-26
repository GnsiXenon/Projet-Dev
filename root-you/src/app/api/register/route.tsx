import { NextRequest } from "next/server"
import { NextResponse } from "next/server"
import { redirect } from "next/navigation"

export async function POST(request: NextRequest) {
    let apiHostname = "localhost:3232"
    if (process.env.API_HOSTNAME) {
        apiHostname = process.env.API_HOSTNAME
    }
    const formData = await request.formData()
    let data = {
        name: formData.get("name")?.toString(),
        mail: formData.get("mail")?.toString(),
        password: formData.get("password")?.toString()
    }
    const res = await fetch(`http://${apiHostname}/register`, {
        method: "POST",
        body: JSON.stringify(data)
    })
    if (res.ok) {
        redirect("/login")
    } else {
        return NextResponse.json({message: "an error happend"})
    }
}