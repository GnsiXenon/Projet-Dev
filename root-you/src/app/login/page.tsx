import { cookies } from "next/headers"
import { redirect } from "next/navigation"
import Link from "next/link"

export default function Login() {
    if (cookies().get("token")) {
        redirect("/")
    }

    return (
        <main>
            <form action="/api/login" method="POST">
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
                <input type="submit" value="submit"/>
            </form>
            <Link href="/register" className="">Don't have an account, let's create one !</Link>
        </main>
    )
}