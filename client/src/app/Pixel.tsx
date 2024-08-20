import { Signal } from "solid-js";
import type { Pixel } from "../utils/models";
import { setPixel, store } from "../utils/store"

export function DrawPixel({ signal }: { signal: Signal<Pixel> }) {
  const [pixel] = signal
  return (
    <div 
      onPointerDown={() => setPixel({ index: pixel().index, color: store.selectedColor })}
      class="hover:backdrop-brightness-75 aspect-square"
      style={{ "background-color": pixel().color }}
    />
  )
}
