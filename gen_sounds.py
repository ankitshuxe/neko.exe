import wave
import math
import random
import struct

SAMPLE_RATE = 44100

def write_wav(filename, samples):
    with wave.open(filename, 'w') as f:
        f.setnchannels(1)
        f.setsampwidth(2)
        f.setframerate(SAMPLE_RATE)
        for s in samples:
            # Clamp and pack
            s = max(-1.0, min(1.0, s))
            f.writeframesraw(struct.pack('<h', int(s * 32767.0)))

def generate_rain(duration_sec):
    samples = []
    # Simple lowpass filter on white noise for rain/pink-ish noise
    prev = 0.0
    for _ in range(int(SAMPLE_RATE * duration_sec)):
        noise = random.uniform(-1.0, 1.0)
        # Lowpass filter: y[i] = a * x[i] + (1-a) * y[i-1]
        val = 0.1 * noise + 0.9 * prev
        prev = val
        samples.append(val * 2.0)
    return samples

def generate_purr(duration_sec):
    samples = []
    # Purr is a low frequency rumble (~25Hz) modulated by a slower breathing wave (~0.3Hz)
    for i in range(int(SAMPLE_RATE * duration_sec)):
        t = i / SAMPLE_RATE
        rumble = math.sin(2.0 * math.pi * 25.0 * t) + 0.5 * math.sin(2.0 * math.pi * 50.0 * t)
        breath = 0.5 + 0.5 * math.sin(2.0 * math.pi * 0.4 * t)
        samples.append((rumble * breath) * 0.4)
    return samples

def generate_lofi(duration_sec):
    samples = []
    # Simple warm chord loop (Cmaj7 -> Fmaj7) with some crackle
    chords = [
        [261.63, 329.63, 392.00, 493.88], # Cmaj7
        [174.61, 220.00, 261.63, 329.63], # Fmaj7
    ]
    for i in range(int(SAMPLE_RATE * duration_sec)):
        t = i / SAMPLE_RATE
        chord_idx = int(t / 2.0) % 2
        chord = chords[chord_idx]
        
        val = 0.0
        for freq in chord:
            # Soft sine with some warmth (harmonics)
            val += math.sin(2.0 * math.pi * freq * t)
            val += 0.2 * math.sin(2.0 * math.pi * freq * 2.0 * t)
            
        val /= len(chord)
        
        # Add crackle (infrequent random spikes)
        crackle = 0.0
        if random.random() < 0.0005:
            crackle = random.uniform(-0.5, 0.5)
            
        # Envelope to prevent clicking on chord change
        t_chord = t % 2.0
        env = min(1.0, t_chord * 10.0) * min(1.0, (2.0 - t_chord) * 10.0)
            
        samples.append(val * env * 0.3 + crackle)
    return samples

write_wav('sounds/rain.wav', generate_rain(5.0))
write_wav('sounds/purr.wav', generate_purr(5.0))
write_wav('sounds/lofi.wav', generate_lofi(8.0))
