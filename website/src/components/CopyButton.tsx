"use client";

import { useState } from "react";
import { Copy, Check } from "lucide-react";

export default function CopyButton({ text }: { text: string }) {
  const [copied, setCopied] = useState(false);

  const handleCopy = () => {
    navigator.clipboard.writeText(text);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  return (
    <button
      onClick={handleCopy}
      className="absolute right-4 top-1/2 -translate-y-1/2 p-2 hover:bg-foreground/10 transition-colors flex items-center gap-2 group-hover:opacity-100 opacity-50"
      title="Copy to clipboard"
    >
      {copied ? <Check className="w-4 h-4 text-accent" /> : <Copy className="w-4 h-4" />}
    </button>
  );
}
