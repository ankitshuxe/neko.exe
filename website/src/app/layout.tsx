import type { Metadata } from "next";
import { Space_Mono } from "next/font/google";
import "./globals.css";

const spaceMono = Space_Mono({
  weight: ["400", "700"],
  subsets: ["latin"],
  variable: "--font-mono",
});

export const metadata: Metadata = {
  title: "neko.exe | Terminal-Native Pomodoro Companion",
  description: "A strictly CLI-based Pomodoro timer featuring a digital ASCII cat. Earn Fish Coins for deep focus, upgrade your environment, and zero browser bloat.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className={`${spaceMono.variable} h-full antialiased`}>
      <body className="min-h-full flex flex-col font-mono text-foreground bg-page selection:bg-accent selection:text-white">
        {children}
      </body>
    </html>
  );
}
