import type { Metadata } from "next";
import "./globals.css";

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
      <body className="bg-[#48EFEF]">{children}</body>
    </html>
  );
}
