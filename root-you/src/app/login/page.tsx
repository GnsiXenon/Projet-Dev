import { cookies } from "next/headers"
import { redirect } from "next/navigation"
import Link from "next/link"

export default function Login() {
    if (cookies().get("token")) {
        redirect("/")
    }

    return (
        <main className="flex flex-col justify-center items-center h-screen">
            <form className="flex flex-col justify-center items-start border-[2px] border-white rounded-lg p-7 gap-4" action="/api/login" method="POST">
                <div className="flex flex-col justify-center items-start">
                    <label>
                        Mail
                    </label>
                    <input type="mail" name="mail" className="text-black" />
                </div>
                <div className="flex flex-col justify-center items-start">
                    <label>
                        Password
                    </label>
                    <input type="password" name="password" className="text-black" />
                </div>
                <input type="submit" value="Submit" className="p-4 text-green-400 border-green-400 border-[2px] rounded-lg" />
                <Link href="/register" className="">Don't have an account, let's create one !</Link>
            </form>
        </main>
    )
}