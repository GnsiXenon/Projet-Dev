import { cookies } from "next/headers"
import { redirect } from "next/navigation"
import Link from "next/link"

export default function Login() {
    if (cookies().get("token")) {
        redirect("/")
    }

    return (
        <main>
            <form action="/api/register" method="POST">
                <div>
                    <label>
                        Username
                    </label>
                    <input type="text" name="name" />
                </div>
                <div>
                    <label>
                        Mail
                    </label>
                    <input type="mail" name="mail" />
                </div>
                <div>
                    <label>
                        Password
                    </label>
                    <input type="password" name="password" />
                </div>
                <input type="submit" value="Submit" />
            </form>
            <Link href="/login" className="">Already have an account, let's log in !</Link>
        </main>
    )
}