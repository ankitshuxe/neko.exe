export default function Marquee() {
  const text = " ++ TERMINAL-NATIVE POMODORO ++ EARN FISH COINS FOR DEEP WORK ++ ZERO ELECTRON BLOAT ++ STARTLE PROTOCOLS ENGAGED ++ PURE GO CLI ++ ";
  
  return (
    <div className="w-full bg-accent text-card overflow-hidden border-y border-foreground py-2 flex whitespace-nowrap">
      <div className="animate-marquee font-mono text-xs tracking-widest inline-block uppercase">
        {text}{text}{text}{text}
      </div>
      <div className="animate-marquee font-mono text-xs tracking-widest inline-block uppercase absolute top-2">
        {text}{text}{text}{text}
      </div>
    </div>
  );
}
