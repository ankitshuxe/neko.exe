"use client";

import { useEffect, useState } from "react";
import { motion } from "framer-motion";

export default function ASCIIPlayer() {
  const [frame, setFrame] = useState(0);

  useEffect(() => {
    const timer = setInterval(() => {
      setFrame((prev) => (prev + 1) % 4);
    }, 1000);
    return () => clearInterval(timer);
  }, []);

  const zzz = "Z".repeat((frame % 3) + 1);

  return (
    <motion.pre 
      initial={{ opacity: 0 }}
      animate={{ opacity: 1 }}
      transition={{ duration: 1 }}
      className="font-mono text-sm leading-tight text-foreground relative z-10 p-8 border border-foreground bg-page shadow-[8px_8px_0px_0px_rgba(26,26,26,1)]"
    >
{`   +--------------------------+
   |  [SESSION // IN PROGRESS] |
   +--------------------------+
   |                          |
   |        (=^･ω･^=) ${zzz.padEnd(3)}     |
   |                          |
   |      25:00 REMAINING     |
   |                          |
   +--------------------------+`}
    </motion.pre>
  );
}
