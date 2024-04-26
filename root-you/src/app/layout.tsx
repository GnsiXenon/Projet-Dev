import type { Metadata } from "next"
import "./globals.css"
import Image from "next/image"
import Bg from "@/../public/bg.png"

export const metadata: Metadata = {
  title: "RootYou",
  description: "CTF training plateform",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className="text-white">
        <Image src={Bg} alt="bg.png" className="z-[-1] absolute top-0 left-0" />
        {children}
      </body>
    </html>
  );
}
