import { NextRequest } from "next/server"
import { NextResponse } from "next/server"
import { redirect } from "next/navigation"
import { cookies } from "next/headers"

var jwt = require('jsonwebtoken')

export async function POST(request: NextRequest) {
    let apiHostname = "localhost:3232"
    if (process.env.API_HOSTNAME) {
        apiHostname = process.env.API_HOSTNAME
    }
    const formData = await request.formData()
    let userId = formData.get("user-id")?.toString()
    let userIdData = 0
    if (userId) {
        userIdData = parseInt(userId)
    }
    let challId = formData.get("chall-id")?.toString()
    let challIdData = 0
    if (challId) {
        challIdData = parseInt(challId)
    }
    let data = {
        "user-id": userIdData,
        "chall-id": challIdData,
        "flag": formData.get("flag")?.toString(),
        "mail": formData.get("mail")?.toString()
    }
    const res = await fetch(`http://${apiHostname}/submit-flag`, {
        method: "POST",
        body: JSON.stringify(data)
    })
    if (res.ok) {
        const resBody = await res.json()
        cookies().set("token", jwt.sign(resBody, "HaCoeur"))
        redirect("/")
    } else {
        return NextResponse.json({message: "an error happend"})
    }
}