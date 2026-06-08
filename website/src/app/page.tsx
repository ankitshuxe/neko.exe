"use client";

import { motion } from "framer-motion";
import { ShieldCheck, Settings, Download, GitBranch, BookOpen, Fingerprint } from "lucide-react";
import Marquee from "@/components/Marquee";
import ASCIIPlayer from "@/components/ASCIIPlayer";
import CopyButton from "@/components/CopyButton";

export default function Home() {
  const containerVariants = {
    hidden: { opacity: 0 },
    show: {
      opacity: 1,
      transition: { staggerChildren: 0.2 }
    }
  };

  const itemVariants = {
    hidden: { opacity: 0, y: 20 },
    show: { opacity: 1, y: 0, transition: { duration: 0.6 } }
  };

  return (
    <main className="flex min-h-screen flex-col items-center justify-start bg-page text-foreground selection:bg-accent selection:text-white">
      
      {/* Top Navbar */}
      <nav className="w-full border-b border-foreground bg-page px-6 py-4 flex justify-between items-center z-50 sticky top-0">
        <div className="flex items-center gap-3">
          <span className="font-mono font-extrabold text-2xl tracking-tighter text-accent">[neko.exe]</span>
        </div>
        <div className="hidden md:flex gap-8 font-mono text-xs font-bold tracking-widest uppercase">
          <a href="#features" className="hover:text-accent transition-colors">Features</a>
          <a href="#install" className="hover:text-accent transition-colors">Install</a>
          <a href="https://github.com/ankitshuxe/neko.exe" target="_blank" rel="noreferrer" className="flex items-center gap-2 hover:text-accent transition-colors">
            <GitBranch className="w-4 h-4" /> Source
          </a>
        </div>
      </nav>

      {/* Marquee Ticker */}
      <Marquee />

      {/* Hero Container */}
      <motion.div 
        variants={containerVariants}
        initial="hidden"
        animate="show"
        className="w-full max-w-7xl mx-auto px-4 md:px-8 py-16 md:py-24"
      >
        <div className="crosshair-corners border border-foreground bg-card shadow-[16px_16px_0px_0px_rgba(26,26,26,1)] mb-24 relative">
          
          {/* Top Decorative Bar */}
          <div className="w-full h-8 border-b border-foreground bg-page flex items-center px-4 justify-between">
            <div className="flex gap-2">
              <span className="w-2 h-2 rounded-full bg-foreground block"></span>
              <span className="w-2 h-2 rounded-full border border-foreground block"></span>
            </div>
            <span className="font-mono text-[10px] uppercase tracking-widest text-foreground/50">SYS.VER.01</span>
          </div>

          <div className="p-8 md:p-16">
            <motion.div variants={itemVariants} className="mb-6 flex items-center gap-4">
              <div className="px-3 py-1 bg-accent text-white font-mono text-xs uppercase font-bold tracking-widest">
                Primary Module
              </div>
              <span className="font-mono text-xs opacity-50">#POMO-X1</span>
            </motion.div>

            <motion.h1 variants={itemVariants} className="text-6xl md:text-8xl font-sans text-foreground mb-8 leading-none">
              TERMINAL-NATIVE <br/><span className="text-accent">POMODORO</span> <br/>COMPANION
            </motion.h1>

            <motion.p variants={itemVariants} className="font-mono text-sm md:text-base leading-relaxed max-w-2xl mb-16 border-l-2 border-accent pl-6">
              Abandon the browser. <span className="font-bold text-accent">neko.exe</span> is a strictly CLI-based productivity companion. Focus deeply, earn Fish Coins, and unlock ASCII environments. Break the timer, and you startle the cat.
            </motion.p>

            <motion.div id="features" variants={containerVariants} className="grid grid-cols-1 md:grid-cols-3 gap-8 border-t border-foreground/20 pt-12 scroll-mt-24">
              
              <motion.div variants={itemVariants} className="group relative">
                <Settings className="w-8 h-8 text-accent mb-6 group-hover:rotate-90 transition-transform duration-500" />
                <h3 className="font-sans text-lg mb-3">Strict Interval Protocols</h3>
                <p className="font-mono text-xs text-foreground/80 leading-relaxed">
                  Execute flawless 25-minute focus intervals directly from your shell. No bloated UI, no web tabs. Just you, your code, and a sleeping ASCII cat keeping watch.
                </p>
              </motion.div>

              <motion.div variants={itemVariants} className="group relative">
                <ShieldCheck className="w-8 h-8 text-accent mb-6" />
                <h3 className="font-sans text-lg mb-3">Gamified Accountability</h3>
                <p className="font-mono text-xs text-foreground/80 leading-relaxed">
                  Every completed session mints Fish Coins. But beware: interrupting a session (Ctrl+C) triggers a "startled" penalty, deducting hard-earned currency. Focus is mandatory.
                </p>
              </motion.div>

              <motion.div variants={itemVariants} className="group relative">
                <Fingerprint className="w-8 h-8 text-accent mb-6" />
                <h3 className="font-sans text-lg mb-3">The ASCII Marketplace</h3>
                <p className="font-mono text-xs text-foreground/80 leading-relaxed">
                  Exchange Fish Coins in the CLI shop for tactical upgrades. Unlock new cat breeds, interactive toys, and cozy environmental skids like cardboard castles or window sills.
                </p>
              </motion.div>

            </motion.div>
          </div>
        </div>

        {/* Installation Blueprint Section */}
        <motion.div 
          initial={{ opacity: 0, y: 40 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true, margin: "-100px" }}
          transition={{ duration: 0.8 }}
          id="install" 
          className="crosshair-corners bg-accent text-card flex flex-col lg:flex-row border border-foreground shadow-[16px_16px_0px_0px_rgba(26,26,26,1)]"
        >
          
          {/* Left Illustration Pane */}
          <div className="flex-1 border-b lg:border-b-0 lg:border-r border-foreground p-8 md:p-16 flex items-center justify-center relative overflow-hidden pattern-dots">
             <ASCIIPlayer />
             
             {/* Decorative blueprint lines */}
             <div className="absolute top-0 left-12 w-px h-full bg-foreground/20"></div>
             <div className="absolute top-1/2 left-0 w-full h-px bg-foreground/20"></div>
          </div>

          {/* Right Data Pane */}
          <div className="flex-1 flex flex-col bg-page text-foreground">
            
            <div className="p-8 border-b border-foreground bg-card">
              <div className="flex items-center gap-3 mb-2">
                <Download className="w-5 h-5 text-accent" />
                <h2 className="font-sans text-2xl">DEPLOYMENT VECTORS</h2>
              </div>
              <p className="font-mono text-xs opacity-60">Select your preferred package manager to initiate installation sequence.</p>
            </div>

            <div className="flex border-b border-foreground group hover:bg-card transition-colors">
              <div className="p-6 border-r border-foreground w-20 flex items-center justify-center font-mono font-bold text-accent">01</div>
              <div className="p-6 flex-1 font-mono text-sm relative pr-16 flex items-center">
                <span className="opacity-50 text-[10px] absolute top-2 right-4">GOLANG</span>
                <span>go install github.com/ankitshuxe/neko.exe@latest</span>
                <CopyButton text="go install github.com/ankitshuxe/neko.exe@latest" />
              </div>
            </div>
            
            <div className="flex border-b border-foreground group hover:bg-card transition-colors">
              <div className="p-6 border-r border-foreground w-20 flex items-center justify-center font-mono font-bold text-accent">02</div>
              <div className="p-6 flex-1 font-mono text-sm relative pr-16 flex items-center">
                <span className="opacity-50 text-[10px] absolute top-2 right-4">WINDOWS</span>
                <span>winget install ankitshuxe.Neko</span>
                <CopyButton text="winget install ankitshuxe.Neko" />
              </div>
            </div>
            
            <div className="flex border-b border-foreground group hover:bg-card transition-colors">
              <div className="p-6 border-r border-foreground w-20 flex items-center justify-center font-mono font-bold text-accent">03</div>
              <div className="p-6 flex-1 font-mono text-sm relative pr-16 flex items-center">
                <span className="opacity-50 text-[10px] absolute top-2 right-4">MACOS / LINUX</span>
                <span>brew install ankitshuxe/tap/neko</span>
                <CopyButton text="brew install ankitshuxe/tap/neko" />
              </div>
            </div>
            
            <div className="flex-1 p-8 md:p-12 flex flex-col justify-end">
              <div className="flex items-center gap-4 mb-6">
                <BookOpen className="w-8 h-8 text-accent" />
                <p className="font-mono italic text-xs">Engineered for Deep Work. Built for Terminals.</p>
              </div>
              <h2 className="font-sans text-3xl leading-none">
                DELIVERABLES ALIGNED WITH STRICT DEVELOPER PRODUCTIVITY STANDARDS.
              </h2>
            </div>
            
            <div className="flex border-t border-foreground bg-foreground text-page">
              <div className="p-4 flex-1 font-mono text-xs uppercase tracking-widest">Neko Technologies Corp.</div>
              <div className="p-4 w-32 border-l border-page text-center font-mono text-xs">EST. 2026</div>
            </div>
          </div>

        </motion.div>
      </motion.div>

      {/* Footer */}
      <footer className="w-full border-t border-foreground bg-card py-12 mt-12">
        <div className="max-w-7xl mx-auto px-8 flex flex-col md:flex-row justify-between items-center gap-6">
          <div className="flex items-center gap-4 opacity-10">
            <div className="font-mono font-extrabold text-4xl md:text-6xl tracking-tighter">[neko.exe]</div>
          </div>
          <div className="font-mono text-xs tracking-widest uppercase flex gap-8">
            <a href="https://github.com/ankitshuxe/neko.exe" className="hover:text-accent border-b border-transparent hover:border-accent pb-1">GitHub</a>
            <span className="opacity-30">|</span>
            <span>MIT License</span>
            <span className="opacity-30">|</span>
            <span>All Systems Nominal</span>
          </div>
        </div>
      </footer>
    </main>
  );
}
