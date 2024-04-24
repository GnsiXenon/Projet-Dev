import { cookies } from "next/headers"
import { redirect } from "next/navigation"
import Link from "next/link"

export default function Login() {
    if (cookies().get("token")) {
        redirect("/")
    }

    return (
        <main>
            <form action="/">
                <div>
                    <label>
                        Mail
                    </label>
                    <input type="text" name="mail" />
                </div>
                <div>
                    <label>
                        Password
                    </label>
                    <input type="password" name="password" />
                </div>
            </form>
            <Link href="/login" className="">Already have an account, let's log in !</Link>
        </main>
    )
}