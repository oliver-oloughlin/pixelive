import { Signal } from "solid-js";
import { setPixel, store } from "../utils/store"

export function DrawPixel({ id, signal }: { id: number, signal: Signal<string> }) {
  const [color] = signal
  return (
    <div 
      onPointerDown={() => setPixel({ id, color: store.selectedColor })}
      class="hover:backdrop-brightness-75 hover:invert-[50%] aspect-square"
      style={{ "background-color": color() }}
    />
  )
}
