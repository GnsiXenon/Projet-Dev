import { NextRequest } from "next/server"
import { NextResponse } from "next/server"
import { cookies } from "next/headers"
import { redirect } from "next/navigation"

var jwt = require("jsonwebtoken")

export async function POST(request: NextRequest) {
    let apiHostname = "localhost:3232"
    if (process.env.API_HOSTNAME) {
        apiHostname = process.env.API_HOSTNAME
    }
    const formData = await request.formData()
    let data = {
        mail: formData.get("mail")?.toString(),
        password: formData.get("password")?.toString()
    }
    const res = await fetch(`http://${apiHostname}/login`, {
        method: "POST",
        body: JSON.stringify(data)
    })
    if (res.ok) {
        const resJson = await res.json()
        if (resJson["error"]) {
            redirect("/login")
        }
        cookies().set("martin_session_id", jwt.sign(resJson, "HaCoeur"))
        redirect("/")
    } else {
        return NextResponse.json({message: "wrong creds"})
    }
}